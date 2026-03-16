package models

import "time"

type PaymentStatus string

const (
	PaymentStatusPending PaymentStatus = "pending"
	PaymentStatusPaid    PaymentStatus = "paid"
	PaymentStatusFailed  PaymentStatus = "failed"
	PaymentStatusExpired PaymentStatus = "expired"
)

type Payment struct {
	Base
	OrderID         string        `gorm:"uniqueIndex;not null" json:"order_id"`
	MidtransOrderID string        `gorm:"uniqueIndex;not null" json:"midtrans_order_id"`
	SnapToken       string        `json:"snap_token"`
	Status          PaymentStatus `gorm:"type:varchar(20);default:'pending'" json:"status"`
	PaymentType     string        `json:"payment_type"`
	Amount          int64         `gorm:"not null" json:"amount"`
	PaidAt          *time.Time    `json:"paid_at"`

	Order Order
}
