package main

import (
	"context"
	pb "final/InventoryService/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedInventoryServiceServer
}

func (s *server) PackOrder(ctx context.Context, req *pb.OrderDetails) (*pb.ArrivalDate, error) {
	return &pb.ArrivalDate{ArrivalDate: "2024.05.22"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterInventoryServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
