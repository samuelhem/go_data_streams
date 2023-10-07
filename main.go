package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net"

	ds "github.com/samuelhem/go_data_streams/datastreams"
	"google.golang.org/grpc"
)

const (
	defaultHost = "localhost"
	defaultPort = 8080
)

func main() {

	global_messaging_queue := make(chan *ds.Message, 100)
	appPool := ds.NewApplicationPool(10)
	spawn_workers(5, appPool, global_messaging_queue)

	start_grpc(defaultHost, defaultPort, appPool, global_messaging_queue)
}

func start_grpc(host string, port int, appPool *ds.ApplicationPool, messageQueue chan *ds.Message) {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		slog.Error("failed to listen: {}", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	ds.RegisterDataStreamServiceServer(grpcServer, &ds.DefaultReceiver{ApplicationPool: appPool, MessageQueue: messageQueue})
	slog.Info("gRPC server running on", host, port)
	grpcServer.Serve(lis)
}

func spawn_workers(n int, appPool *ds.ApplicationPool, messageQueue chan *ds.Message) {
	for i := 0; i < n; i++ {
		go ds.NewWorker(appPool, messageQueue).Start()
	}
}
