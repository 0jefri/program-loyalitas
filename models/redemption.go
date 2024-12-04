package models

import "time"

type Redemption struct {
	ID             uint      `json:"redemption_id" gorm:"primaryKey"`
	UserID         uint      `json:"user_id"`
	Reward         string    `json:"reward"`
	RedemptionDate time.Time `json:"tanggal_penukaran"`
}
