package service

import (
	"context"
	"github.com/google/uuid"
	"time"
	"transactions/internal/models"
	"transactions/internal/storage"
)

type ITransactionsService interface {
	Create(ctx context.Context, transaction models.Transaction) (string, error)
	Get(ctx context.Context, ID string) (models.Transaction, error)
	Delete(ctx context.Context, ID string) error
}

type TransactionsService struct {
	repo *storage.Storage
}

func (s *TransactionsService) Create(ctx context.Context, transaction models.Transaction) (string, error) {
	transaction.PaymentTime = time.Now()
	transaction.ID = uuid.NewString()
	return s.repo.Transaction.Create(ctx, transaction)
}

func (s *TransactionsService) Get(ctx context.Context, ID string) (models.Transaction, error) {
	return s.repo.Transaction.Get(ctx, ID)
}

func (s *TransactionsService) Delete(ctx context.Context, ID string) error {
	return s.repo.Transaction.Delete(ctx, ID)
}

func NewTransactionsService(r *storage.Storage) *TransactionsService {
	return &TransactionsService{
		repo: r,
	}
}
