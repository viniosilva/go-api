package model

import (
	"time"
)

type Cat struct {
	ID       int
	Name     string
	Birthday time.Time
}
