package models

import "time"

type PayoutStatus string

const (
	PayoutStatusPending     PayoutStatus = "pending"
	PayoutStatusTransferred PayoutStatus = "transferred"
)

type Payout struct {
	Base
	ShopID        string       `gorm:"not null" json:"shop_id"`
	Amount        int64        `gorm:"not null" json:"amount"`
	Status        PayoutStatus `gorm:"type:varchar(20);default:'pending'" json:"status"`
	Notes         string       `json:"notes"`
	TransferredAt *time.Time   `json:"transferred_at"`

	Shop Shop
}
