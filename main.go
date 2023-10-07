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
    start_grpc(defaultHost, defaultPort)
}


func start_grpc(host string, port int) {
    flag.Parse()
    lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
    if err != nil {
        slog.Error("failed to listen: {}", err)
    }
    var opts []grpc.ServerOption
    grpcServer := grpc.NewServer(opts...)
    ds.RegisterDatastreamServer(grpcServer, &ds.DefaultReceiver{})
    slog.Info("gRPC server running on", host, port)
    grpcServer.Serve(lis)
}



