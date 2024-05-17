package main

import (
	"context"
	"final/ClientService/client"
	pb "final/ClientService/proto"
	"final/ClientService/server/repository"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

type server struct {
	pb.UnsafeClientServiceServer
	db repository.Database
	c  *client.Client
}

func (s *server) MakeOrder(ctx context.Context, req *pb.MakeOrderRequest) (*pb.Status, error) {
	arrivalDate, err := s.c.PackOrder(req.ProductName, req.Quantity)
	if err != nil {
		fmt.Println(err)
		return &pb.Status{Message: "Fail"}, nil
	}

	fmt.Println(arrivalDate)

	err = s.db.MakeOrder(ctx, req.ProductName, req.Quantity, arrivalDate.ArrivalDate)
	return &pb.Status{Message: "Success"}, nil
}

func (s *server) GetOrders(ctx context.Context, req *emptypb.Empty) (*pb.Orders, error) {
	orders, err := s.db.GetOrders(ctx)
	if err != nil {
		fmt.Println(err)
		return &pb.Orders{}, nil
	}
	fmt.Println(orders)
	list := []*pb.Order{}
	for _, order := range orders {
		list = append(list, &pb.Order{Id: order.Id, ProductName: order.ProductName, Quantity: order.Quantity, Created: order.Created, Arrival: order.Arrival})
	}

	return &pb.Orders{Orders: list}, nil
}

//func (s *server) GetMessage(ctx context.Context, req *emptypb.Empty) (*pb.Test, error) {
//	if true {
//		c, err := client.NewClient()
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		message, err := c.GetInventoryMessage()
//
//		if err != nil {
//			fmt.Sprintf("%s Error while getting message", message)
//		}
//
//		return &pb.Test{
//			Message: message.Message,
//		}, nil
//	}
//
//	return &pb.Test{
//		Message: "Hello from the Client Service!",
//	}, nil
//}

func main() {
	ctx := context.Background()
	db, err := repository.Connect(ctx)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	c, err := client.NewClient()
	if err != nil {
		log.Fatalf("could not connect to client: %v", err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterClientServiceServer(s, &server{db: db, c: c})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
