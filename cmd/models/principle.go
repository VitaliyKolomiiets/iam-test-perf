package models

type Principle struct {
	StatementId uint   `gorm:"primaryKey"`
	Krn         string `gorm:"primaryKey"`
}
