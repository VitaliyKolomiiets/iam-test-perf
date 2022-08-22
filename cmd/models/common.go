package models

import (
	"time"
)

type DBModel struct {
	ID        uint      `gorm:"primarykey" json:"-"`
	CreatedAt time.Time `                  json:"-"`
	UpdatedAt time.Time `                  json:"-"`
}

type IDModel struct {
	ID        string    `gorm:"-" json:"id"`
	CreatedAt time.Time `         json:"created"`
	UpdatedAt time.Time `         json:"updated"`
}
