package lib

import (
	"context"
        ds "github.com/samuelhem/go_data_streams/datastreams"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DefaultClientReceiver struct {
    messageBrokers map[string]MessageBroker
    ds.UnimplementedDataStreamServiceServer
}

func (s* DefaultClientReceiver) Exchange(ctx context.Context, message *ds.Message) (*emptypb.Empty, error) {
    chn, err := s.messageBrokers[message.Stream].Receive()
    if err != nil {
        return nil, err
    }
    chn <- message

    return &emptypb.Empty{}, nil
}
