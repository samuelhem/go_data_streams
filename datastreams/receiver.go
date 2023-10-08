package datastreams

import (
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type DefaultReceiver struct {
	ApplicationPool *ApplicationPool
	MessageQueue    chan *Message
	UnimplementedDataStreamServiceServer
}

func (s *DefaultReceiver) Exchange(ctx context.Context, message *Message) (*emptypb.Empty, error) {

	id, err := uuid.Parse(message.Sender.Value)
	if err != nil {
		return nil, errors.New("ApplicationId invalid")
		//TODO: put application not registed message back to the sender
	}

	isRegistered := s.ApplicationPool.IsRegistered(id)
	if !isRegistered {
		return nil, errors.New("Application needs to be subscribed to the server")
	}

	switch message.Type {
	case MessageType_BROADCAST:
		s.processBroadcastMessage(message)

	case MessageType_SUBSCRIBE:
		s.ApplicationPool.GetApplication(id).AddEvent(message.Stream, message.Event)

	case MessageType_UNSUBSCRIBE:

	}

	return &emptypb.Empty{}, nil
}

func (s *DefaultReceiver) Register(ctx context.Context, application *Application) (*Application, error) {
	slog.Info("Application registering to the pool")

	if err := s.ApplicationPool.RegisterApplication(application); err != nil {
		slog.Error("Application registration failed", "error", err)
		return nil, err
	}

	slog.Info("Application registered successfully", "id", application.Id.Value)
	return application, nil
}

func (s *DefaultReceiver) processBroadcastMessage(message *Message) {
	s.MessageQueue <- message

}

