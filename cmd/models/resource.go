package models

type Resource struct {
	StatementId uint   `gorm:"primaryKey"`
	Krn         string `gorm:"primaryKey"`
}
