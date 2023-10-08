package main

import (
	"fmt"
	"log/slog"
	"net"

	ds "github.com/samuelhem/go_data_streams/datastreams"
	"google.golang.org/grpc"
)

const (
	defaultHost                = "localhost"
	defaultPort                = 8080
	defaultWorkerCount         = 5
	defaultQueueSize           = 100
	defaultApplicationPoolSize = 10
)

func main() {

	messagingQueue := make(chan *ds.Message, defaultQueueSize)
	appPool := ds.NewApplicationPool(defaultApplicationPoolSize)
	grpcServer := createGrpcServer(defaultHost, defaultPort)

	dataStreamService := newDataStreamService(grpcServer, appPool, messagingQueue)
	ds.RegisterDataStreamServiceServer(grpcServer, dataStreamService.createDataStreamGrpcService())

	dataStreamService.spawnWorkers(defaultWorkerCount)

	slog.Info("gRPC server running on", defaultHost, defaultPort)
	grpcServer.Serve(listenOnTcp(defaultHost, defaultPort))
}

type DataStreamService struct {
	grpcServer   *grpc.Server
	appPool      *ds.ApplicationPool
	messageQueue chan *ds.Message
}

func newDataStreamService(grpcServer *grpc.Server, appPool *ds.ApplicationPool, messageQueue chan *ds.Message) *DataStreamService {
	return &DataStreamService{grpcServer: grpcServer, appPool: appPool, messageQueue: messageQueue}
}

func (dsi *DataStreamService) createWorker() *ds.Worker {
	return ds.NewWorker(dsi.appPool, dsi.messageQueue)
}

func (dsi *DataStreamService) createDataStreamGrpcService() ds.DataStreamServiceServer {
	return &ds.DefaultReceiver{ApplicationPool: dsi.appPool, MessageQueue: dsi.messageQueue}
}

func (dsi *DataStreamService) spawnWorkers(n int) {
	for i := 0; i < n; i++ {
		go dsi.createWorker().Start()
	}
}

func createGrpcServer(host string, port int) *grpc.Server {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	return grpcServer
}

func listenOnTcp(host string, port int) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		slog.Error("failed to listen", "error", err)
		panic(err)
	}
	return lis
}
