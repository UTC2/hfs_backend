package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type hfs_sale struct {
	Id               int       `json:"id" gorm:"column:id;"`
	OrderId          string    `json:"order_id" gorm:"column:order_id;"`
	ProductId        int       `json:"product_id" gorm:"column:product_id;"`
	ModelTransaction string    `json:"modelTransaction" gorm:"column:modelTransaction;"`
	FeeRental        string    `json:"feeRental" gorm:"column:feeRental;"`
	StartDateRental  time.Time `json:"startDateRental" gorm:"column:startDateRental;"`
	EndDateRental    time.Time `json:"endDateRental" gorm:"column:endDateRental;"`
	Description      string    `json:"description" gorm:"column:description;"`
}

type hfs_saleUpdate struct {
	Id               *int       `json:"id" gorm:"column:id;"`
	OrderId          *string    `json:"order_id" gorm:"column:order_id;"`
	ProductId        *int       `json:"product_id" gorm:"column:product_id;"`
	ModelTransaction *string    `json:"modelTransaction" gorm:"column:modelTransaction;"`
	FeeRental        *string    `json:"feeRental" gorm:"column:feeRental;"`
	StartDateRental  *time.Time `json:"startDateRental" gorm:"column:startDateRental;"`
	EndDateRental    *time.Time `json:"endDateRental" gorm:"column:endDateRental;"`
	Description      *string    `json:"description" gorm:"column:description;"`
}

func (hfs_sale) TableName() string {
	return "hfs_sale"
}
func (hfs_saleUpdate) TableName() string {
	return hfs_sale{}.TableName()
}

func main() {
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db, err)
}
