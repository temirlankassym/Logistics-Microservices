package client

import (
	"context"
	pb "final/InventoryService/SupplierProto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
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

	ctx, _ := context.WithTimeout(context.Background(), time.Second)

	return &Client{conn: c, ctx: ctx}, nil
}

func (c *Client) GetSupplierMessage() (*pb.SupplierMessage, error) {
	return c.conn.GetSupplierMessage(c.ctx, &emptypb.Empty{})
}
