package main

import (
	"context"
	"flag"
	"log/slog"
	"time"

	ds "github.com/samuelhem/go_data_streams/datastreams"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	applicationName    = flag.String("name", "", "The name of the application")
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func main() {
	flag.Parse()
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
	}

	defer conn.Close()

	client := ds.NewDatastreamClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

        slog.Info("Initiated connection, registering application")
        register_with_server(ctx, client)
}

func register_with_server(ctx context.Context, client ds.DatastreamClient) {
	app := ds.NewApplication("test", "localhost:10000")
        app, err := client.Register(ctx, app)
        if err != nil {
            slog.Error("failed to register application:", err)
        }

        slog.Info("Application registered successfully, id: ", app.Id.Value)

}
