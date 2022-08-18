package models

type Action struct {
	StatementId uint   `gorm:"primaryKey"`
	Action      string `gorm:"primaryKey"`
}
