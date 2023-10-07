package datastreams

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"google.golang.org/protobuf/types/known/anypb"
)

type MessageBroker interface {
	Publish(event string, message any) error
        Subscribe(event string) error
	Receive() (*Message, error)
}

type DefaultMessageBroker struct {
        AppId *UUID
	DataStream *DataStream
	Client     *DataStreamServiceClient
}

func NewDefaultMessageBroker(appId *UUID, data_stream *DataStream, client *DataStreamServiceClient) *DefaultMessageBroker {
	if client == nil {
		panic("Default Message Broker needs a grpc client to communicate with the server")
	}
	return &DefaultMessageBroker{
                AppId: appId,
		DataStream: data_stream,
		Client:     client,
	}
}

func (broker *DefaultMessageBroker) publish(message *Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := (*broker.Client).Exchange(ctx, message)
	if err != nil {
		slog.Info("failed to exchange message with server:", err)
		return err
	}
	return nil
}

func (broker *DefaultMessageBroker) Receive() (*Message, error) {
	return nil, nil
}

func (broker *DefaultMessageBroker) Publish(event string, message any) error {
	data, err := json.Marshal(message)
	if err != nil {
		slog.Error("failed to marshal message:", err)
	}

	return broker.publish(NewMessage(broker.AppId, broker.DataStream, event, MessageType_BROADCAST, data))
}

func (broker *DefaultMessageBroker) Subscribe(event string) error {
    return broker.publish(NewMessage(broker.AppId, broker.DataStream, event, MessageType_SUBSCRIBE, nil))


}

func NewMessage(appId *UUID, dataStream *DataStream, event string, messageType MessageType, data []byte) *Message {
	return &Message{
                Sender: appId,
 		DataStream: dataStream,
		Type:       messageType,
		Data:       &anypb.Any{Value: data},
		Event:      event,
	}
}
