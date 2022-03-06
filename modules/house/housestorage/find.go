package housestorage

import (
  "context"
  "gorm.io/gorm"
  "hfs_backend/common"
  "hfs_backend/modules/house/housemodel"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*housemodel.House, error) {
	var result housemodel.House

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).
		First(&result).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &result, nil
}
