package model

import (
	"time"
)

type Cat struct {
	ID       int `gorm:"primaryKey"`
	Birthday time.Time
	Name     string
}
