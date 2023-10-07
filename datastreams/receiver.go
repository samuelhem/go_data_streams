package datastreams

import (
	"context"
	"log/slog"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var (
	applicationPool = NewApplicationPool(10)
)

type DefaultReceiver struct {
	UnimplementedDatastreamServer
}

func (s *DefaultReceiver) Exchange(ctx context.Context, message *Message) (*emptypb.Empty, error) {
	slog.Info("Incoming message from sender")
	return &emptypb.Empty{}, nil
}

func (s *DefaultReceiver) Register(ctx context.Context, application *Application) (*Application, error) {
	slog.Info("Application registering to the pool")

	if err := applicationPool.RegisterApplication(application); err != nil {
		slog.Error("Application registration failed", err)
		return nil, err
	}

	slog.Info("Application registered successfully, id: ", application.Id.Value)
	return application, nil
}
