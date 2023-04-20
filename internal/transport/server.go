package transport

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"transactions/internal/service"
	pb "transactions/proto"
)

type Server struct {
	pb.TransactionServiceServer
	service service.Service
}

func Run(service *service.Service) error {
	listen, err := net.Listen("tcp", ":8081")
	fmt.Println("Starting server")
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterTransactionServiceServer(s, &Server{service: *service})
	reflection.Register(s)
	if err := s.Serve(listen); err != nil {
		return err
	}

	return nil
}
