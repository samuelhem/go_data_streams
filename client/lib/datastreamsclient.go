package lib 

import (
	"context"
	"errors"
	"log/slog"
	"time"

	ds "github.com/samuelhem/go_data_streams/datastreams"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DataStreamsClient struct {
	serviceClient  ds.DataStreamServiceClient
	messageBrokers map[string]ds.MessageBroker
}

func NewDataStreamsClient(clientName string, serverAddr *string) *DataStreamsClient {

	slog.Info("Initizalizing go_data_streams client...")
	client, conn := initializeGrpcClient(clientName, serverAddr)
	defer conn.Close()
	slog.Info("Finished initializing!")

	return &DataStreamsClient{
		serviceClient:  client,
		messageBrokers: make(map[string]ds.MessageBroker),
	}
}

func (client *DataStreamsClient) createMessageBroker(streamName string) ds.MessageBroker {
	if client.messageBrokers[streamName] == nil {
		slog.Info("Creating message broker for stream", "stream", streamName)
		stream := &ds.DataStream{Name: streamName}
		client.messageBrokers[streamName] = ds.NewDefaultMessageBroker(stream, &client.serviceClient)
	} else {
            slog.Info("Message broker already exists for stream", "stream", streamName)
            return nil
        }
	return client.messageBrokers[streamName]

}

func initializeGrpcClient(appName string, serverAddr *string) (ds.DataStreamServiceClient, *grpc.ClientConn) {
	var opts []grpc.DialOption
	/* if *tls {
	    if *caFile == "" {
	        *caFile = testdata.Path("ca.pem")
	    }
	    creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
	    if err != nil {
	        log.Fatalf("Failed to create TLS credentials %v", err)
	    }
	    opts = append(opts, grpc.WithTransportCredentials(creds))
	} else { */
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//}

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		slog.Error("fail to dial:", err)
                panic("Failed to connect to DataStreams server")
	}

	client := ds.NewDataStreamServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	slog.Info("Initiated grpc connection, registering application...")
        app := ds.NewApplication(appName, *serverAddr)
	register_with_server(ctx, app, client)
	return client, conn
}

func register_with_server(ctx context.Context, app *ds.Application, client ds.DataStreamServiceClient) error {
	app, err := client.Register(ctx, app)
	if err != nil {
		slog.Error("failed to register application:", err)
                return errors.New("Failed to register application")
	}

	slog.Info("Application registered successfully", "id", app.Id.Value)
        return nil
}
