package productlikemodel

import "time"

type Like struct {
  ProductId     int         `json:"product_id" gorm:"column:product_id;"`
  UserId        int         `json:"user_id" gorm:"column:user_id;"`
  CreatedAt    *time.Time   `json:"created_at" gorm:"column:created_at;"`
}

func (like Like) TableName() string { return "product_likes" }
