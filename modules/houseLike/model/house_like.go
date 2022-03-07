package houselikemodel

import "time"

type Like struct {
  HouseId       int         `json:"house_id" gorm:"column:house_id;"`
  UserId        int         `json:"user_id" gorm:"column:user_id;"`
  CreatedAt    *time.Time   `json:"created_at" gorm:"column:created_at;"`
}

func (like Like) TableName() string { return "house_likes" }
