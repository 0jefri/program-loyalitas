package repositories

import (
	"loyalty-program-api/models"

	"gorm.io/gorm"
)

type RedemptionRepository interface {
	Create(redemption *models.Redemption) error
	GetByUserID(userID uint) ([]models.Redemption, error)
}

type redemptionRepository struct {
	DB *gorm.DB
}

func NewRedemptionRepository(db *gorm.DB) RedemptionRepository {
	return &redemptionRepository{DB: db}
}

func (r *redemptionRepository) Create(redemption *models.Redemption) error {
	return r.DB.Create(redemption).Error
}

func (r *redemptionRepository) GetByUserID(userID uint) ([]models.Redemption, error) {
	var redemptions []models.Redemption
	err := r.DB.Where("user_id = ?", userID).Find(&redemptions).Error
	return redemptions, err
}
