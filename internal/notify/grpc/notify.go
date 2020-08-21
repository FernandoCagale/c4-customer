package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type NotifyGRPC struct {
	client *grpc.ClientConn
}

func NewNotifyGRPC(client *grpc.ClientConn) *NotifyGRPC {
	return &NotifyGRPC{
		client: client,
	}
}

func (notify *NotifyGRPC) GetNotify(headers map[string]string,) ([]string, error) {
	client := NewNotifyClient(notify.client)

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(headers))

	result, err := client.GetNotify(ctx, &Request{})
	if err != nil {
		return nil, err
	}

	return result.Types, nil
}
