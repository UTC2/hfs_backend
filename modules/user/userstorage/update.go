package userstorage

import (
	"context"
	"hfs_backend/common"
)

func (s *sqlStore) UpdateUser(ctx context.Context, conditions map[string]interface{}, updates map[string]interface{}) error {
	if err := s.db.Table("users").Where(conditions).Updates(updates).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
