package client

import (
	"context"
	pb "final/InventoryService/SupplierProto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn pb.SupplierServiceClient
	ctx  context.Context
}

func NewClient() (*Client, error) {
	conn, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := pb.NewSupplierServiceClient(conn)

	return &Client{conn: c, ctx: context.Background()}, nil
}

func (c *Client) CreateDelivery(productName string, quantity int32) (*pb.ShipTime, error) {
	return c.conn.CreateDelivery(c.ctx, &pb.ProductInfo{Name: productName, Quantity: quantity})
}
