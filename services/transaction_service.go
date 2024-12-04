package services

import (
	"errors"
	"loyalty-program-api/models"
	"loyalty-program-api/repositories"
	"time"
)

type TransactionService interface {
	RecordTransaction(userID uint, transactionID string, nominal float64, paymentMethod string) (*models.Transaction, error)
	GetTransactionHistory(userID uint) ([]models.Transaction, error)
}

type transactionService struct {
	transactionRepo repositories.TransactionRepository
	userRepo        repositories.UserRepository
}

func NewTransactionService(transactionRepo repositories.TransactionRepository, userRepo repositories.UserRepository) TransactionService {
	return &transactionService{transactionRepo: transactionRepo, userRepo: userRepo}
}

func (s *transactionService) RecordTransaction(userID uint, transactionID string, nominal float64, paymentMethod string) (*models.Transaction, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	// Calculate points
	var pointsEarned float64
	if paymentMethod == "COD" {
		pointsEarned = nominal * 0.01
	} else {
		pointsEarned = nominal * 0.001
	}

	// Create transaction record
	transaction := &models.Transaction{
		UserID:          userID,
		TransactionID:   transactionID,
		Nominal:         nominal,
		PaymentMethod:   paymentMethod,
		TransactionDate: time.Now(),
		PointsEarned:    pointsEarned,
	}
	err = s.transactionRepo.Create(transaction)
	if err != nil {
		return nil, err
	}

	// Update user points
	user.Points += pointsEarned
	err = s.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *transactionService) GetTransactionHistory(userID uint) ([]models.Transaction, error) {
	transactions, err := s.transactionRepo.GetByUserID(userID)
	if err != nil {
		return nil, errors.New("transaction history not found")
	}
	return transactions, nil
}
