package transaction

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID                int
	Number            string
	Date              time.Time
	PriceTotal        int
	CostTotal         int
	DetailTransaction []DetailTransaction
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeleteAt          gorm.DeletedAt
}

type DetailTransaction struct {
	ID            int
	TransactionId int
	ItemId        string
	ItemQuantity  int
	ItemPrice     int
	ItemCost      int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeleteAt      gorm.DeletedAt
}

type Item struct {
	ID        int
	Name      string
	Price     int
	Cost      int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
