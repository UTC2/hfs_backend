package uploadstorage

import (
	"context"
	"hfs_backend/common"
)

func (store *sqlStore) DeleteImages(ctx context.Context, ids []int) error {
	db := store.db
	if error := db.Table(common.Image{}.TableName()).Where("id in ?", ids).Delete(nil).Error; error != nil {
		return error
	}
	return nil
}
