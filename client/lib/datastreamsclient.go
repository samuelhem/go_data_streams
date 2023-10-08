package lib

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	ds "github.com/samuelhem/go_data_streams/datastreams"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DataStreamsClient struct {
	app            *ds.Application
	serviceClient  ds.DataStreamServiceClient
	conn           *grpc.ClientConn
	messageBrokers map[string]MessageBroker
}

func NewDataStreamsClient(clientName string, serverAddr string) *DataStreamsClient {

	slog.Info("Initizalizing go_data_streams client...")
	conn := initializeGrpcConnection(clientName, serverAddr)
        slog.Info("Initiated grpc connection, registering application...")
	client := ds.NewDataStreamServiceClient(conn)
        app := initApplication(clientName, serverAddr, client)

	slog.Info("Finished initializing!")

	return &DataStreamsClient{
		app:            app,
		serviceClient:  client,
		conn:           conn,
		messageBrokers: make(map[string]MessageBroker),
	}
}

func (client *DataStreamsClient) Close() {
	client.conn.Close()
}

func (client *DataStreamsClient) CreateMessageBroker(streamName string) MessageBroker {
	if client.messageBrokers[streamName] == nil {
		slog.Info("Creating message broker for stream", "stream", streamName)
		stream := &ds.DataStream{Name: streamName}
		client.messageBrokers[streamName] = client.NewDefaultMessageBroker(stream)
	} else {
		slog.Info("Message broker already exists for stream", "stream", streamName)
		return nil
	}
	return client.messageBrokers[streamName]

}

func initializeGrpcConnection(appName string, serverAddr string) *grpc.ClientConn {
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

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		slog.Error("fail to dial:", err)
		panic("Failed to connect to DataStreams server")
	}
        return conn
}

func initApplication(appName string, serverAddr string, client ds.DataStreamServiceClient) *ds.Application {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, app := ds.NewApplication(appName, serverAddr)
        app, err := registerApplication(ctx, app, client)
	if err != nil {
		panic("Failed to register application")
	}
	return app 
}

func registerApplication(ctx context.Context, app *ds.Application, client ds.DataStreamServiceClient) (*ds.Application, error) {
	slog.Info(fmt.Sprintf("Registering application %+v", app))
	app, err := client.Register(ctx, app)
	if err != nil {
		slog.Error("failed to register application:", err)
		return nil, errors.New("Failed to register application")
	}

	slog.Info("Application registered successfully", "id", app.Id.Value)
	return app, nil
}
