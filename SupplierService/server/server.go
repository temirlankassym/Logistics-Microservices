package main

import (
	"context"
	pb "final/SupplierService/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedSupplierServiceServer
}

func (s *server) GetSupplierMessage(ctx context.Context, req *emptypb.Empty) (*pb.SupplierMessage, error) {
	return &pb.SupplierMessage{
		Message: "Hello from the Supplier!",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSupplierServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
