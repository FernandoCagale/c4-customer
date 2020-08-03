package grpc

import (
	"context"
	"google.golang.org/grpc"
	"time"
)

type NotifyGRPC struct {
	client *grpc.ClientConn
}

func NewNotifyGRPC(client *grpc.ClientConn) *NotifyGRPC {
	return &NotifyGRPC{
		client: client,
	}
}

func (notify *NotifyGRPC) GetNotify() ([]string, error) {
	client := NewNotifyClient(notify.client)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	result, err := client.GetNotify(ctx, &Request{})
	if err != nil {
		return nil, err
	}

	return result.Types, nil
}
