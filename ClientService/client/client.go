package client

import (
	pb "client/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn pb.ClientServiceClient
	ctx  context.Context
}

func NewClient() (*Client, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := pb.NewClientServiceClient(conn)

	return &Client{conn: c, ctx: context.Background()}, nil
}

func (c *Client) MakeOrder(productName string, quantity int32) (*pb.Status, error) {
	return c.conn.MakeOrder(c.ctx, &pb.MakeOrderRequest{ProductName: productName, Quantity: quantity})
}
