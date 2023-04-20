package postgre

import (
	"context"
	"gorm.io/gorm"
	"transactions/internal/models"
)

type TransactionRepository struct {
	db *gorm.DB
}
type User struct {
	ID       string
	Name     string
	Surname  string
	Login    string
	Password string
	Balance  float32
}

func (r *TransactionRepository) Create(ctx context.Context, transaction models.Transaction) (string, error) {
	if err := r.db.WithContext(ctx).Create(transaction).Error; err != nil {
		return "", err
	}
	var user User
	if err := r.db.WithContext(ctx).Find(&user, "login = ?", transaction.Username).Error; err != nil {
		return "", err
	}
	if transaction.TypeOfTransaction == "-" {
		user.Balance -= transaction.Amount
	} else {
		user.Balance += transaction.Amount
	}
	if err := r.db.WithContext(ctx).Model(&User{}).Where("login = ?", transaction.Username).Updates(&user).Error; err != nil {
		return "", err
	}
	return transaction.ID, nil

}

func (r *TransactionRepository) Get(ctx context.Context, ID string) (models.Transaction, error) {
	var ans models.Transaction
	err := r.db.WithContext(ctx).Where("id = ?", ID).First(&ans).Error
	if err != nil {
		return models.Transaction{}, err
	}
	return ans, nil
}

func (r *TransactionRepository) Delete(ctx context.Context, ID string) error {
	return r.db.WithContext(ctx).Delete(&models.Transaction{}, ID).Error
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}
