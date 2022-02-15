package productstorage

import (
	"context"
	"hfs_backend/common"
	"hfs_backend/modules/product/productmodel"
)

func (s *sqlStore) UpdateData(ctx context.Context, id int, data *productmodel.ProductUpdate) error {
	db := s.db
	if error := db.Where("id = ?", id).Updates(data).Error; error != nil {
		return common.ErrDB(error)
	}
	return nil
}
