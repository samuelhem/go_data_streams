package lib

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	ds "github.com/samuelhem/go_data_streams/datastreams"
        util "github.com/samuelhem/go_data_streams/util"
	"google.golang.org/grpc"
)

type DataStreamsClient struct {
	app            *ds.Application
	serviceClient  ds.DataStreamServiceClient
	conn           *grpc.ClientConn
	messageBrokers map[string]MessageBroker
        grpcProxy     *grpc.Server
}

const (
    REVERSE_PROXY_PORT = "10050"
)

func NewDataStreamsClient(clientName string, serverAddr string, serverPort string, clientPort string)  *DataStreamsClient {

	slog.Info("Initizalizing go_data_streams client...")
        conn := util.InitializeGrpcConnection(fmt.Sprintf("%s:%s", serverAddr, serverPort))
	slog.Info("Initiated grpc connection, registering application...")
	client := ds.NewDataStreamServiceClient(conn)
        app := initApplication(clientName, fmt.Sprintf("%s:%s", serverAddr, clientPort) , client)

	slog.Info("Finished initializing!")


        brokers := make(map[string]MessageBroker)
        grpcProxy := util.CreateGrpcServer()  
        ds.RegisterDataStreamServiceServer(grpcProxy, &DefaultClientReceiver{messageBrokers: brokers})


	return &DataStreamsClient{
		app:            app,
		serviceClient:  client,
		conn:           conn,
		messageBrokers: brokers,
                grpcProxy:     grpcProxy,
	}
}

func (client *DataStreamsClient) Start() {
    slog.Info("Starting grpc proxy...")
    go client.grpcProxy.Serve(util.ListenOnTcpStr(client.app.Hostname))

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
