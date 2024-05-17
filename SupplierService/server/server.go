package main

import (
	"context"
	pb "final/SupplierService/proto"
	"final/SupplierService/server/repository"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct {
	pb.UnimplementedSupplierServiceServer
	db repository.Database
}

func (s *server) CreateDelivery(ctx context.Context, req *pb.ProductInfo) (*pb.ShipTime, error) {
	days, err := s.db.GetDeliveryTime(ctx, req.Name)
	if err != nil {
		return &pb.ShipTime{}, err
	}

	arrival := time.Now().AddDate(0, 0, int(days)).Format("02.01.2006")

	err = s.db.CreateDelivery(ctx, req.Name, req.Quantity, arrival)

	return &pb.ShipTime{
		Days: days,
	}, nil
}

func main() {
	ctx := context.Background()
	db, err := repository.Connect(ctx)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSupplierServiceServer(s, &server{db: db})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
