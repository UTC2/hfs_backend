package userstorage

import (
  "context"
  "hfs_backend/common"
  "hfs_backend/modules/user/usermodel"
)

func (s *sqlStore) Create(ctx context.Context, data *usermodel.UserCreate) error {
	db := s.db.Begin()
  if err := db.Table(data.TableName()).Create(data).Error; err != nil {
    db.Rollback()
    return common.ErrDB(err)
  }
  if err := db.Commit().Error; err != nil {
    db.Rollback()
    return common.ErrDB(err)
  }
	return nil
}
