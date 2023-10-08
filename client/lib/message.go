package lib

import (
	"context"
	"encoding/json"
	ds "github.com/samuelhem/go_data_streams/datastreams"
	"log/slog"
	"time"

	"google.golang.org/protobuf/types/known/anypb"
)

type MessageBroker interface {
	Publish(event string, message any) error
	Subscribe(event string) error
        Receive() (chan *ds.Message, error)
}

type DefaultMessageBroker struct {
	App        *ds.Application
	DataStream *ds.DataStream
	Client     ds.DataStreamServiceClient
        receiver   chan *ds.Message
}

func (dsc *DataStreamsClient) NewDefaultMessageBroker(data_stream *ds.DataStream) *DefaultMessageBroker {
	if dsc.serviceClient == nil {
		panic("Default Message Broker needs a grpc client to communicate with the server")
	}
	return &DefaultMessageBroker{
		App:        dsc.app,
		Client:     dsc.serviceClient,
		DataStream: data_stream,
                receiver:   make(chan *ds.Message, 10), 
	}
}

func (broker *DefaultMessageBroker) publish(message *ds.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := broker.Client.Exchange(ctx, message)
	if err != nil {
		slog.Info("failed to exchange message with server:", err)
		return err
	}
	return nil
}

func (broker *DefaultMessageBroker) Receive()  (chan *ds.Message, error) {
	return broker.receiver, nil
}

func (broker *DefaultMessageBroker) Publish(event string, message any) error {
	data, err := json.Marshal(message)
	if err != nil {
		slog.Error("failed to marshal message:", err)
	}

	return broker.publish(broker.NewMessage(event, ds.MessageType_BROADCAST, data))
}

func (broker *DefaultMessageBroker) Subscribe(event string) error {
	return broker.publish(broker.NewMessage(event, ds.MessageType_SUBSCRIBE, nil))

}

func (broker *DefaultMessageBroker) NewMessage(event string, messageType ds.MessageType, data []byte) *ds.Message {
	return &ds.Message{
		Sender:     broker.App.Id,
		Stream:     broker.DataStream.Name,
		Type:       messageType,
		Data:       &anypb.Any{Value: data},
                Event:      &ds.Event{Name: event},
	}
}
