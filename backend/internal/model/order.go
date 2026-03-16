package models

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusPaid       OrderStatus = "paid"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

// Order is one checkout session per buyer.
// Splits into SubOrders per shop for fulfillment.
type Order struct {
	Base
	UserID          string      `gorm:"not null" json:"user_id"`
	Status          OrderStatus `gorm:"type:varchar(20);default:'pending'" json:"status"`
	TotalAmount     int64       `gorm:"not null" json:"total_amount"`
	ShippingName    string      `json:"shipping_name"`
	ShippingPhone   string      `json:"shipping_phone"`
	ShippingAddress string      `json:"shipping_address"`
	ShippingCity    string      `json:"shipping_city"`
	ShippingNote    string      `json:"shipping_note"`

	User      User
	SubOrders []SubOrder
	Payment   *Payment
}

type SubOrder struct {
	Base
	OrderID  string      `gorm:"not null" json:"order_id"`
	ShopID   string      `gorm:"not null" json:"shop_id"`
	Status   OrderStatus `gorm:"type:varchar(20);default:'pending'" json:"status"`
	Subtotal int64       `gorm:"not null" json:"subtotal"`

	Order Order
	Shop  Shop
	Items []OrderItem
}

type OrderItem struct {
	Base
	SubOrderID  string `gorm:"not null" json:"sub_order_id"`
	ProductID   string `gorm:"not null" json:"product_id"`
	ProductName string `json:"product_name"`          // snapshot
	Price       int64  `gorm:"not null" json:"price"` // snapshot
	Quantity    int    `gorm:"not null" json:"quantity"`
	Subtotal    int64  `gorm:"not null" json:"subtotal"`

	SubOrder SubOrder
	Product  Product
}
