package housemodel
import (
  "hfs_backend/common"
  "strings"
)
const EntityName = "House"
type House struct {
	common.SQLModel `json:",inline"`
	Sku             string          `json:"sku" gorm:"column:sku;"`
	Name            string          `json:"name" gorm:"column:name;"`
	Description     string          `json:"description" gorm:"description;"`
	ProductStatusId int             `json:"product_status_id" gorm:"product_status_id;"`
	RegularPrice    float64         `json:"regular_price" gorm:"regular_price;"`
	DiscountPrice   float64         `json:"discount_price" gorm:"discount_price;"`
	Quantity        int             `json:"quantity" gorm:"quantity;"`
	Taxable         int             `json:"taxable" gorm:"taxable;"`
	productImages   *common.Images  `json:"product_images" gorm:"column:product_images;"`
	likedCount      int             `json:"like_count" gorm:"-"`
}

type HouseUpdate struct {
  Sku             string         `json:"sku" gorm:"column:sku;"`
  Description     *string        `json:"description" gorm:"description;"`
	ProductStatusId *int           `json:"product_status_id" gorm:"product_status_id;"`
	RegularPrice    *float64       `json:"regular_price" gorm:"regular_price;"`
	DiscountPrice   *float64       `json:"discount_price" gorm:"discount_price;"`
	Quantity        *int           `json:"quantity" gorm:"quantity;"`
	Taxable         *int           `json:"taxable" gorm:"taxable;"`
  productImages      *common.Images `json:"foodimages" gorm:"colum:foodimages;"`
}
type HouseCreate struct {
  common.SQLModel `json:",inline"`
  Sku             string         `json:"sku" gorm:"column:sku;"`
  Name            string         `json:"name" gorm:"column:name;"`
  Description     string         `json:"description" gorm:"description;"`
  ProductStatusId int            `json:"product_status_id" gorm:"product_status_id;"`
  RegularPrice    float64        `json:"regular_price" gorm:"regular_price;"`
  DiscountPrice   float64        `json:"discount_price" gorm:"discount_price;"`
  Quantity        int            `json:"quantity" gorm:"quantity;"`
  Taxable         int            `json:"taxable" gorm:"taxable;"`
  productImages      *common.Images `json:"foodimages" gorm:"colum:foodimages;"`
}

func (HouseUpdate) TableName() string {
	return House{}.TableName()
}
func (House) TableName() string {
  return "houses"
}
func (HouseCreate) TableName() string {
	return House{}.TableName()
}
func (res *HouseCreate) Validate() error {
  res.Name = strings.TrimSpace(res.Name)
  if res.Quantity == 0 {
		return ErrQuantityMustBePositive
	}
  if len(res.Name) == 0 {
    return ErrNameCannotBeEmpty
  }
  return nil
}
var (
  ErrNameCannotBeEmpty = common.NewCustomError(nil, "product's name can't be blank", "ErrNameCannotBeEmpty")
  ErrQuantityMustBePositive = common.NewCustomError(nil, "quantity's name must be positive", "ErrQuantityMustBePositive")
)
func (product *House) Mask(isAdminOrOwner bool)  {
  product.GenerateUID(common.DbTypeRestaurant)
}
