package productstorage

import (
	"context"
	"gorm.io/gorm"
	"hfs_backend/modules/product/productmodel"
)

type sqlStore struct {
	db *gorm.DB
}

func (s *sqlStore) FindDataByCondition(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*productmodel.Product, error) {
	panic("implement me")
}
func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
