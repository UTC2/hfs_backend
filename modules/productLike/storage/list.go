package productlikestorage

import (
  "context"
  "hfs_backend/common"
  productlikemodel "hfs_backend/modules/productLike/model"
)

func (s *sqlStore) GetProductLikes(ctx context.Context,ids []int)(map[int]int,error)  {
  result := make(map[int]int)

  type sqlData struct {
    ProductId int `gorm:"column:product_id;"`
    LikeCount int `gorm:"column:count;"`
  }

  var listLike []sqlData

  if error := s.db.Table(productlikemodel.Like{}.TableName()).
    Select("product_id, count(product_id) as count").
    Where("product_id in (?)", ids).
    Group("product_id").
    Find(&listLike).Error; error != nil {
      return nil, common.ErrDB(error)
  }

  for _,item := range listLike{
    result[item.ProductId]=item.LikeCount
  }

  return result, nil
}
