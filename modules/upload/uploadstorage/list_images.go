package uploadstorage

import (
	"context"
	"hfs_backend/common"
)

func (store *sqlStore) ListImages(ctx context.Context, ids []int, morekeys ...string) ([]common.Image, error) {
	db := store.db
	var result []common.Image
	db = db.Table(common.Image{}.TableName())
	if error := db.Where("id in (?)", ids).Find(&result).Error; error != nil {
		return nil, common.ErrDB(error)
	}
	return result, nil
}
