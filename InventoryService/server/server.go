package main

import (
	"context"
	pb "final/InventoryService/proto"
	"final/InventoryService/server/repository"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct {
	pb.UnimplementedInventoryServiceServer
	db repository.Database
	//c  *client.Client
}

func (s *server) PackOrder(ctx context.Context, req *pb.OrderDetails) (*pb.ArrivalDate, error) {
	arrivalDate := time.Now()

	productCount, err := s.db.DecrementProduct(ctx, req.ProductName, req.Quantity)
	if err != nil {
		return &pb.ArrivalDate{}, err
	}

	if productCount > 0 {
		// need to deliver new products, stock is not enough, 14 days
		return &pb.ArrivalDate{ArrivalDate: arrivalDate.AddDate(0, 0, 14).Format("02.01.2006")}, nil
	}

	// 7 days to deliver if product present in stock
	return &pb.ArrivalDate{ArrivalDate: arrivalDate.AddDate(0, 0, 7).Format("02.01.2006")}, nil
}

func main() {
	ctx := context.Background()
	db, err := repository.Connect(ctx)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterInventoryServiceServer(s, &server{db: db})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
