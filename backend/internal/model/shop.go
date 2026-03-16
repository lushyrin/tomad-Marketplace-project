package models

type ShopStatus string

const (
	ShopStatusPending   ShopStatus = "pending"
	ShopStatusApproved  ShopStatus = "approved"
	ShopStatusRejected  ShopStatus = "rejected"
	ShopStatusSuspended ShopStatus = "suspended"
)

type Shop struct {
	Base
	UserID      string     `gorm:"uniqueIndex;not null" json:"user_id"`
	Name        string     `gorm:"not null" json:"name"`
	Slug        string     `gorm:"uniqueIndex;not null" json:"slug"`
	Description string     `json:"description"`
	LogoURL     string     `json:"logo_url"`
	City        string     `json:"city"`
	Status      ShopStatus `gorm:"type:varchar(20);default:'pending'" json:"status"`
	Rating      float64    `gorm:"default:0" json:"rating"`
	TotalSales  int        `gorm:"default:0" json:"total_sales"`
	BankName    string     `json:"bank_name"`
	BankAccount string     `json:"bank_account"`
	BankHolder  string     `json:"bank_holder"`

	User     User
	Products []Product
}
