package models

type Product struct {
	Base
	ShopID      string  `gorm:"not null" json:"shop_id"`
	CategoryID  string  `gorm:"not null" json:"category_id"`
	Name        string  `gorm:"not null" json:"name"`
	Slug        string  `gorm:"uniqueIndex;not null" json:"slug"`
	Description string  `json:"description"`
	Price       int64   `gorm:"not null" json:"price"`
	Stock       int     `gorm:"default:0" json:"stock"`
	Weight      float64 `json:"weight"`
	IsActive    bool    `gorm:"default:true" json:"is_active"`

	Shop     Shop
	Category Category
	Images   []ProductImage
}

type ProductImage struct {
	Base
	ProductID string `gorm:"not null" json:"product_id"`
	URL       string `gorm:"not null" json:"url"`
	IsPrimary bool   `gorm:"default:false" json:"is_primary"`
	SortOrder int    `json:"sort_order"`

	Product Product
}

type Category struct {
	Base
	Name     string  `gorm:"not null" json:"name"`
	Slug     string  `gorm:"uniqueIndex;not null" json:"slug"`
	ParentID *string `json:"parent_id"`

	Parent   *Category
	Children []Category `gorm:"foreignKey:ParentID"`
	Products []Product
}
