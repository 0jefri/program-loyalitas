package models

type User struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Name   string  `json:"name" gorm:"not null"`
	Email  string  `json:"email" gorm:"unique;not null"`
	Points float64 `json:"points" gorm:"default:0"`
}
