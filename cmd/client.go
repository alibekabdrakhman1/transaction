package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "transactions/proto"
)

var addr string = "localhost:8081"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewTransactionServiceClient(conn)
	transactionReq := &pb.CreateTransRequest{
		Transaction: &pb.Transaction{
			Username:    "f",
			Type:        "+",
			Amount:      100,
			Description: "dada",
		}}

	resp, err := c.Create(context.Background(), transactionReq)
	if err != nil {
		log.Println("error", err)
	}

	fmt.Println(resp)
}
