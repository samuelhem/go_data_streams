package util

import (
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitializeGrpcConnection(serverAddr string) *grpc.ClientConn {
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

func CreateGrpcServer() *grpc.Server {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	return grpcServer
}

func ListenOnTcp(host string, port int) net.Listener {
	return ListenOnTcpStr(fmt.Sprintf("%s:%d", host, port))
}

func ListenOnTcpStr(serverAddr string) net.Listener {
	lis, err := net.Listen("tcp", serverAddr)
	if err != nil {
		slog.Error("failed to listen", "error", err)
		panic(err)
	}
	slog.Info("Listening on TCP", "address", serverAddr)
	return lis
}
