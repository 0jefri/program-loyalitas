package models

import "time"

type Transaction struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	UserID          uint      `json:"user_id"`
	TransactionID   string    `json:"transaction_id" gorm:"unique"`
	Nominal         float64   `json:"nominal_transaksi"`
	PaymentMethod   string    `json:"payment_method"`
	TransactionDate time.Time `json:"tanggal_transaksi"`
	PointsEarned    float64   `json:"poin_didapat"`
}
