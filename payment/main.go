package main

import (
	"context"
	"fmt"
	"github.com/Sotatek-HungNgo3/be-practical-payment/config"
	"github.com/Sotatek-HungNgo3/be-practical-payment/pb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"time"
)

type server struct {
	pb.UnimplementedPaymentServiceServer
}

func RandBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

func (s server) MakePayment(ctx context.Context, req *pb.MakePaymentRequest) (*pb.MakePaymentResponse, error) {
	return &pb.MakePaymentResponse{Status: pb.PaymentStatus_Confirmed}, nil
	if RandBool() {
		return &pb.MakePaymentResponse{Status: pb.PaymentStatus_Confirmed}, nil
	}
	return &pb.MakePaymentResponse{Status: pb.PaymentStatus_Declined, ErrorMessage: "Something when wrong with payment gateway, please try again"}, nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	grpcPort := config.GetGRPCPort()

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPaymentServiceServer(s, &server{})

	log.Printf("GRPC server listening on port %v", grpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
