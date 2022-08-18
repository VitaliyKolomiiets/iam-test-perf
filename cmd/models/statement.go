package models

type Statement struct {
	ID         uint        `gorm:"primaryKey"`
	Type       string      `gorm:"column:type;type:string;size:256"        json:"-"`
	Principles []Principle `gorm:"foreignKey:StatementId;references:ID"`
	Actions    []Action    `gorm:"foreignKey:StatementId;references:ID"`
	Resources  []Resource  `gorm:"foreignKey:StatementId;references:ID"`
}
