package uploadstorage

import (
	"context"
	"hfs_backend/common"
)

func (store *sqlStore) createImage(ctx context.Context, image *common.Image) error {
	db := store.db
	if err := db.Table(image.TableName()).Create(image).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
