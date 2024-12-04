package repositories

import (
	"loyalty-program-api/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction *models.Transaction) error
	GetByUserID(userID uint) ([]models.Transaction, error)
}

type transactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{DB: db}
}

func (r *transactionRepository) Create(transaction *models.Transaction) error {
	return r.DB.Create(transaction).Error
}

func (r *transactionRepository) GetByUserID(userID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.DB.Where("user_id = ?", userID).Find(&transactions).Error
	return transactions, err
}
