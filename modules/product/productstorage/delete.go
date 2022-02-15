package productstorage

import (
	"context"
	"hfs_backend/common"
	"hfs_backend/modules/product/productmodel"
)

func (s *sqlStore) SoftDeleteData(ctx context.Context,
	id int,
) error {
	db := s.db
	if error := db.Table(productmodel.Product{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": 0,
		}).Error; error != nil {
		return common.ErrDB(error)
	}
	return nil
}
