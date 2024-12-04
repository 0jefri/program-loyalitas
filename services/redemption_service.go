package services

import (
	"errors"
	"loyalty-program-api/models"
	"loyalty-program-api/repositories"
	"time"
)

type RedemptionService interface {
	RedemptionPoints(userID uint, reward string) (*models.Redemption, error)
	GetRedemptionHistory(userID uint) ([]models.Redemption, error)
}

type redemptionService struct {
	redemptionRepo repositories.RedemptionRepository
	userRepo       repositories.UserRepository
}

func NewRedemptionService(redemptionRepo repositories.RedemptionRepository, userRepo repositories.UserRepository) RedemptionService {
	return &redemptionService{redemptionRepo: redemptionRepo, userRepo: userRepo}
}

func (s *redemptionService) RedemptionPoints(userID uint, reward string) (*models.Redemption, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	var requiredPoints int
	switch reward {
	case "Voucher Diskon Rp50.000":
		requiredPoints = 500
	case "Voucher Diskon Rp100.000":
		requiredPoints = 1000
	default:
		return nil, errors.New("invalid reward")
	}

	if user.Points < float64(requiredPoints) {
		return nil, errors.New("not enough points")
	}

	// Create redemption record
	redemption := &models.Redemption{
		UserID:         userID,
		Reward:         reward,
		RedemptionDate: time.Now(),
	}
	err = s.redemptionRepo.Create(redemption)
	if err != nil {
		return nil, err
	}

	// Deduct points
	user.Points -= float64(requiredPoints)
	err = s.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return redemption, nil
}

func (s *redemptionService) GetRedemptionHistory(userID uint) ([]models.Redemption, error) {
	redemptions, err := s.redemptionRepo.GetByUserID(userID)
	if err != nil {
		return nil, errors.New("redemption history not found")
	}
	return redemptions, nil
}
