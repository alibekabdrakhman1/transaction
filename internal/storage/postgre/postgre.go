package postgre

import (
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"transactions/internal/models"
)

func Dial(ctx context.Context, url string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if db != nil {
		err := db.AutoMigrate(&models.Transaction{})
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
