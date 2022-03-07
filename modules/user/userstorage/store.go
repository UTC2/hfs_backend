package userstorage

import (
  "golang.org/x/net/context"
  "gorm.io/gorm"
  "hfs_backend/modules/user/usermodel"
)

type sqlStore struct {
  db *gorm.DB
}

func (s *sqlStore) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
  panic("implement me")
}

func NewSQLStore(db *gorm.DB) *sqlStore {
  return &sqlStore{db: db}
}
