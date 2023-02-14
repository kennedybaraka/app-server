package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Store struct {
	ID           string    `json:"id" gorm:"primaryKey;"`
	StoreName    string    `validate:"required" json:"store_name" gorm:"type:string"`
	StoreAddress string    `validate:"required" json:"store_address" gorm:"type:string"`
	StoreEmail   string    `validate:"required" json:"store_email" gorm:"type:string;unique"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Owners       []Owner
	Products     []Product
	Categories   []Category
	Brands       []Brand
}

type Owner struct {
	ID            string    `json:"id" gorm:"primaryKey;"`
	OwnerName     string    `json:"owner_name" gorm:"type:string"`
	OwnerEmail    string    `validate:"required" json:"owner_email" gorm:"type:string;unique"`
	OwnerPassword string    `validate:"required" json:"owner_password" gorm:"type:string"`
	Verified      bool      `json:"verified" gorm:"type:boolean;default:false"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	StoreId       string    `json:"shop_id" gorm:"foreignKey;"`
}

type Product struct {
	ID                 string    `json:"id" gorm:"primaryKey;"`
	ProductName        string    `validate:"required" json:"product_name" gorm:"type:string"`
	ProductImage       string    `json:"product_image" gorm:"type:string"`
	ProductRating      int64     `json:"product_rating" gorm:"type:int;default:0"`
	ProductDescription string    `validate:"required" json:"product_description" gorm:"type:string"`
	ProductPrice       float64   `validate:"required" json:"product_price" gorm:"type:float"`
	ProductStock       int64     `json:"product_stock" gorm:"type:int;default:1"`
	ProductColors      string    `json:"product_colors" gorm:"type:string"`
	ProductSizes       []string  `json:"product_sizes" gorm:"type:string"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	StoreId            string    `json:"shop_id" gorm:"foreignKey;"`
	CategoryId         string    `json:"category_id" gorm:"foreignKey;"`
	BrandId            string    `json:"brand_id" gorm:"foreignKey;"`
	Reviews            []Review
	// ProductColors      pq.StringArray `json:"product_colors" gorm:"type:string[]"`

}
type Review struct {
	ID            string    `json:"id" gorm:"primaryKey;"`
	ReviewTitle   int64     `json:"review_title" gorm:"type:string"`
	ReviewMessage int64     `json:"review_message" gorm:"type:string"`
	ReviewRating  int64     `json:"review_rating" gorm:"type:int"`
	ProductId     string    `json:"product_id" gorm:"foreignKey;"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Category struct {
	ID           string    `json:"id" gorm:"primaryKey;"`
	CategoryName string    `validate:"required" json:"category_name" gorm:"type:string"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	StoreId      string    `json:"shop_id" gorm:"foreignKey;"`
	Products     []Product
}

type Brand struct {
	ID        string    `json:"id" gorm:"primaryKey;"`
	BrandName string    `validate:"required" json:"brand_name" gorm:"type:string"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	StoreId   string    `json:"shop_id" gorm:"foreignKey;"`
	Products  []Product
}

type Client struct {
	ID              string    `json:"id" gorm:"primaryKey;"`
	ClientName      string    ` json:"client_name" gorm:"type:string"`
	ClientEmail     string    `validate:"required" json:"client_email" gorm:"type:string;unique"`
	ClientPassword  string    `validate:"required" json:"client_password" gorm:"type:string"`
	ClientTelephone string    `json:"client_telephone" gorm:"type:string"`
	NewsLetter      string    `json:"new_letter" gorm:"type:bool;default:false"`
	Role            string    `json:"role" gorm:"type:string;default:client"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
type Address struct {
	ID           string    `json:"id" gorm:"primaryKey;"`
	ClientID     string    `json:"client_id" gorm:"foreignKey;"`
	Longitude    float64   `json:"longitude"`
	Latitude     float64   `json:"latitude"`
	AddressName  string    `json:"client_password" gorm:"type:string"`
	AddressPhone string    `json:"client_phone" gorm:"type:string"`
	AddressCity  string    `json:"client_city" gorm:"type:string"`
	Directions   string    `json:"directions" gorm:"type:string"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (s *Store) BeforeCreate(db *gorm.DB) (err error) {
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
	s.ID = uuid.New().String()
	return
}

func (s *Store) BeforeUpdate(tx *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	return
}

func (o *Owner) BeforeCreate(db *gorm.DB) (err error) {
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()
	o.ID = uuid.New().String()
	return
}

func (o *Owner) BeforeUpdate(tx *gorm.DB) (err error) {
	o.UpdatedAt = time.Now()
	return
}

func (p *Product) BeforeCreate(db *gorm.DB) (err error) {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	p.ID = uuid.New().String()
	return
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	p.UpdatedAt = time.Now()
	return
}

func (c *Category) BeforeCreate(db *gorm.DB) (err error) {
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	c.ID = uuid.New().String()
	return
}

func (c *Category) BeforeUpdate(tx *gorm.DB) (err error) {
	c.UpdatedAt = time.Now()
	return
}

func (b *Brand) BeforeCreate(db *gorm.DB) (err error) {
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
	b.ID = uuid.New().String()
	return
}

func (b *Brand) BeforeUpdate(tx *gorm.DB) (err error) {
	b.UpdatedAt = time.Now()
	return
}

func (a *Address) BeforeCreate(db *gorm.DB) (err error) {
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
	a.ID = uuid.New().String()
	return
}

func (a *Address) BeforeUpdate(tx *gorm.DB) (err error) {
	a.UpdatedAt = time.Now()
	return
}

func (c *Client) BeforeCreate(db *gorm.DB) (err error) {
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	c.ID = uuid.New().String()
	return
}

func (c *Client) BeforeUpdate(tx *gorm.DB) (err error) {
	c.UpdatedAt = time.Now()
	return
}
