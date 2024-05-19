package main

import (
	"client/client"
	"client/grpcServer/repository"
	pb "client/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type service struct {
	pb.UnsafeClientServiceServer
	db repository.Database
	c  *client.Client
}

func (s *service) MakeOrder(ctx context.Context, req *pb.MakeOrderRequest) (*pb.Status, error) {
	arrivalDate, err := s.c.PackOrder(req.ProductName, req.Quantity)
	if err != nil {
		fmt.Println(err)
		return &pb.Status{Message: err.Error()}, nil
	}

	err = s.db.MakeOrder(ctx, req.ProductName, req.Quantity, arrivalDate.ArrivalDate)
	if err != nil {
		fmt.Println(err)
		return &pb.Status{Message: "can't make order"}, nil
	}
	return &pb.Status{Message: fmt.Sprintf("Success. Arrival Date: %s", arrivalDate.ArrivalDate)}, nil
}

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
	pb.RegisterClientServiceServer(s, &service{db: db, c: c})

	log.Printf("grpcServer listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
