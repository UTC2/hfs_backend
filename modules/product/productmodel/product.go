package productmodel

import (
	"errors"
	"hfs_backend/common"
	"time"
)

const EntityName = "Product"

type Product struct {
	common.SQLModel `json:",inline"`
	Sku             string         `json:"sku" gorm:"column:sku;"`
	Name            string         `json:"name" gorm:"column:name;"`
	Description     string         `json:"description" gorm:"description;"`
	ProductStatusId int            `json:"product_status_id" gorm:"product_status_id;"`
	RegularPrice    float64        `json:"regular_price" gorm:"regular_price;"`
	DiscountPrice   float64        `json:"discount_price" gorm:"discount_price;"`
	Quantity        int            `json:"quantity" gorm:"quantity;"`
	Taxable         int            `json:"taxable" gorm:"taxable;"`
	Image           *common.Images `json:"image" gorm:"colum:image;"`
}

func (Product) TableName() string {
	return "products"
}

type ProductUpdate struct {
	Description     *string        `json:"description" gorm:"description;"`
	ProductStatusId *int           `json:"product_status_id" gorm:"product_status_id;"`
	RegularPrice    *float64       `json:"regular_price" gorm:"regular_price;"`
	DiscountPrice   *float64       `json:"discount_price" gorm:"discount_price;"`
	Quantity        *int           `json:"quantity" gorm:"quantity;"`
	Taxable         *int           `json:"taxable" gorm:"taxable;"`
	Image           *common.Images `json:"image" gorm:"colum:image;"`
}

func (ProductUpdate) TableName() string {
	return Product{}.TableName()
}

type ProductCreate struct {
	Description     string    `json:"description" gorm:"description;"`
	ProductStatusId int       `json:"product_status_id" gorm:"product_status_id;"`
	RegularPrice    float64   `json:"regular_price" gorm:"regular_price;"`
	DiscountPrice   float64   `json:"discount_price" gorm:"discount_price;"`
	Quantity        int       `json:"quantity" gorm:"quantity;"`
	Taxable         int       `json:"taxable" gorm:"taxable;"`
	InsertAt        time.Time `json:"insert_at" gorm:"insert_at;"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"updated_at;"`
}

func (ProductCreate) TableName() string {
	return Product{}.TableName()
}
func (res *ProductCreate) Validate() error {

	if res.Quantity == 0 {
		return errors.New("quantity must be positive")
	}
	return nil
}
