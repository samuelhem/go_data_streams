package main

import (
	"flag"
	"fmt"
	"log/slog"
	"time"

	cl "github.com/samuelhem/go_data_streams/client/lib"
)

var (
	applicationName    = flag.String("name", "", "The name of the application")
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "localhost", "The server address")
	serverPort         = flag.String("server_port", "10000", "The server port")
	clientPort         = flag.String("client_port", "10050", "The client port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func main() {

	flag.Parse()

	client := cl.NewDataStreamsClient(*applicationName, *serverAddr, *serverPort, *clientPort)
	client.Start()

        time.Sleep(5 * time.Second)

	fmt.Println("Starting to send messages")
	broker := client.CreateMessageBroker("task_channel")

        go listenToEvents(broker)

	broker.Subscribe("create_task_event")
	broker.Subscribe("update_task_event")
	broker.Publish("create_task_event", "Hello World1")
	broker.Publish("create_task_event", "Hello World2")
	broker.Publish("create_task_event", "Hello World3")
	broker.Publish("update_task_event", "Hello World4")

	<-make(chan bool)

}

func listenToEvents(broker cl.MessageBroker) {

	stream, err := broker.Receive()
	if err != nil {
		slog.Info("Error receiving messages from channel")
	}
	for v := range stream {
		slog.Info("Received message: ", "msg", v)
	}

}
