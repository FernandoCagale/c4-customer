package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"os"
)

func NewClient() (*grpc.ClientConn, error) {
	conn, err := grpc.DialContext(context.Background(), os.Getenv("ADDRESS_TYPE_GRPC"), grpc.WithInsecure())
	//conn, err := grpc.Dial(os.Getenv("ADDRESS_TYPE_GRPC"), grpc.WithInsecure(),  )
	if err != nil {
		fmt.Println("did not connect: %v", err)
		return nil, err
	}

	return conn, nil
}
