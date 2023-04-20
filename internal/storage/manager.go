package storage

import (
	"context"
	"transactions/config"
	"transactions/internal/models"
	"transactions/internal/storage/postgre"
)

type ITransactionsRepository interface {
	Create(ctx context.Context, transaction models.Transaction) (string, error)
	Get(ctx context.Context, ID string) (models.Transaction, error)
	Delete(ctx context.Context, ID string) error
}
type Storage struct {
	Transaction ITransactionsRepository
}

func New(ctx context.Context, cfg *config.Config) (*Storage, error) {
	DB, err := postgre.Dial(ctx, cfg.Database.PgUrl)
	if err != nil {
		return nil, err
	}

	transactionRepo := postgre.NewTransactionRepository(DB)

	storage := Storage{
		Transaction: transactionRepo,
	}
	return &storage, nil
}
