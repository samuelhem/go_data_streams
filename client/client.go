package main

import (
	"flag"
        cl  "github.com/samuelhem/go_data_streams/lib"
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

    client := cl.NewDataStreamsClient(*applicationName, serverAddr)
    broker := client.createMessageBroker("test_channel")

    broker.Publish("test message")

}

