package item

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID        int
	Name      string
	Price     string
	Cost      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
