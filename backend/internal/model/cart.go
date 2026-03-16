package models

type Cart struct {
	Base
	UserID string `gorm:"uniqueIndex;not null" json:"user_id"`

	User  User
	Items []CartItem
}

type CartItem struct {
	Base
	CartID    string `gorm:"not null" json:"cart_id"`
	ProductID string `gorm:"not null" json:"product_id"`
	Quantity  int    `gorm:"not null" json:"quantity"`

	Cart    Cart
	Product Product
}
