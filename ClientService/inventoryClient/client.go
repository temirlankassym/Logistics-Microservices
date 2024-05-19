package client

import (
	pb "client/inventoryProto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn pb.InventoryServiceClient
	ctx  context.Context
}

func NewClient() (*Client, error) {
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := pb.NewInventoryServiceClient(conn)

	//ctx, _ := context.WithTimeout(context.Background(), time.Second)

	return &Client{conn: c, ctx: context.Background()}, nil
}

func (c *Client) PackOrder(productName string, quantity int32) (*pb.ArrivalDate, error) {
	return c.conn.PackOrder(c.ctx, &pb.OrderDetails{ProductName: productName, Quantity: quantity})
}
