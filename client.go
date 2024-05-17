package main

import (
	"context"
	pb "final/ClientService/proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
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

	//ctx, _ := context.WithTimeout(context.Background(), time.Second)

	return &Client{conn: c, ctx: context.Background()}, nil
}

func (c *Client) MakeOrder(productName string, quantity int32) (*pb.Status, error) {
	return c.conn.MakeOrder(c.ctx, &pb.MakeOrderRequest{ProductName: productName, Quantity: quantity})
}

func (c *Client) GetOrders() (*pb.Orders, error) {
	return c.conn.GetOrders(c.ctx, &emptypb.Empty{})
}

func main() {
	//http.HandleFunc("/test", TestHandler)
	//log.Fatal(http.ListenAndServe(":8080", nil))

	c, err := NewClient()
	if err != nil {

		log.Fatal("could not connect")
	}

	message, err := c.MakeOrder("Coca-Cola", 5)
	fmt.Println(message)
}

//func TestHandler(writer http.ResponseWriter, request *http.Request) {
//	c, err := NewClient()
//	message, err := c.GetMessage()
//
//	resp, err := json.Marshal(message)
//	if err != nil {
//		log.Fatal("Can't parse json")
//	}
//
//	_, err = writer.Write(resp)
//	if err != nil {
//		log.Fatal("Can't write")
//	}
//}
