package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"os"
)

func NewClient() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(os.Getenv("GRPC"), grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect: %v", err)
		return nil, err
	}

	return conn, nil
}
