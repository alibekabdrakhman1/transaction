package transport

import (
	"context"
	"github.com/google/uuid"
	"time"
	"transactions/internal/models"
	pb "transactions/proto"
)

func (s *Server) Create(ctx context.Context, in *pb.CreateTransRequest) (*pb.CreateTransResponse, error) {
	transaction := models.Transaction{
		ID:                uuid.NewString(),
		Username:          in.Transaction.Username,
		TypeOfTransaction: in.Transaction.Type,
		Amount:            float32(in.Transaction.Amount),
		Description:       in.Transaction.Description,
		PaymentTime:       time.Now(),
	}
	resp, err := s.service.Transactions.Create(ctx, transaction)
	if err != nil {
		return nil, err
	}
	return &pb.CreateTransResponse{Id: resp}, nil
}

func (s *Server) Get(ctx context.Context, in *pb.GetTransRequest) (*pb.GetTransResponse, error) {
	resp, err := s.service.Transactions.Get(ctx, in.TransactionID)
	if err != nil {
		return nil, err
	}
	ans := pb.Transaction{Username: resp.Username, Type: resp.TypeOfTransaction, Amount: int32(resp.Amount), Description: resp.Description}
	return &pb.GetTransResponse{Transaction: &ans}, nil
}

func (s *Server) Delete(ctx context.Context, in *pb.DeleteTransRequest) (*pb.DeleteTransResponse, error) {
	err := s.service.Transactions.Delete(ctx, in.TransactionID)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteTransResponse{}, nil
}
