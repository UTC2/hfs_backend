package productstorage

import (
	"context"
	"hfs_backend/modules/product/productmodel"
)

func (s *sqlStore) Create(ctx context.Context, data *productmodel.ProductCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
