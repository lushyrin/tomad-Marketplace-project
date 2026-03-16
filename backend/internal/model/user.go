package models

type Role string

const (
	RoleBuyer  Role = "buyer"
	RoleSeller Role = "seller"
	RoleAdmin  Role = "admin"
)

type User struct {
	Base
	Email     string `gorm:"uniqueIndex;not null" json:"email"`
	Password  string `gorm:"not null" json:"-"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	AvatarURL string `json:"avatar_url"`
	IsSeller  bool   `gorm:"default:false" json:"is_seller"`
	IsAdmin   bool   `gorm:"default:false" json:"is_admin"`

	Shop   *Shop
	Orders []Order
}
