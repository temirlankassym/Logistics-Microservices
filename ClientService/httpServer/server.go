package main

import (
	"context"
	"encoding/json"
	pb "final/ClientService/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
	"net/http"
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

func (c *Client) GetOrders() (*pb.Orders, error) {
	return c.conn.GetOrders(c.ctx, &emptypb.Empty{})
}

func main() {
	http.HandleFunc("/orders/show", ShowOrders)
	http.HandleFunc("/orders/create", MakeOrder)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ShowOrders(writer http.ResponseWriter, request *http.Request) {
	c, err := NewClient()
	message, err := c.GetOrders()
	if err != nil {
		log.Fatal("Can't get orders")
	}

	resp, err := json.MarshalIndent(message, "", " ")
	if err != nil {
		log.Fatal("Can't parse json")
	}

	_, err = writer.Write(resp)
	if err != nil {
		log.Fatal("Can't write")
	}
}

type CreateOrderRequest struct {
	Name     string `json:"name"`
	Quantity int32  `json:"quantity"`
}

func MakeOrder(writer http.ResponseWriter, request *http.Request) {
	c, err := NewClient()
	req := CreateOrderRequest{}

	body, _ := io.ReadAll(request.Body)

	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Fatal("Can't read request body")
	}

	message, err := c.MakeOrder(req.Name, req.Quantity)
	if err != nil {
		log.Fatal("Can't create order")
	}

	resp, err := json.MarshalIndent(message, "", " ")
	if err != nil {
		log.Fatal("Can't parse json")
	}

	_, err = writer.Write(resp)
	if err != nil {
		log.Fatal("Can't write")
	}
}
