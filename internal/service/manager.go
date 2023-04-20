package service

import (
	"transactions/internal/storage"
)

type Service struct {
	Transactions ITransactionsService
}

func NewManager(storage *storage.Storage) (*Service, error) {
	tSrv := NewTransactionsService(storage)
	return &Service{
		Transactions: tSrv,
	}, nil
}
