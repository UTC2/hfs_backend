package productlikemodel

import "time"

type Like struct {
  RestaurantId int        `json:"restaurant_id" gorm:"column:product_id;"`
  UserId       int        `json:"user_id" gorm:"column:user_id;"`
  CreatedAt    *time.Time `json:"created_at" gorm:"column:created_at;"`
}

func (Like) TableName() string { return "product_likes" }
