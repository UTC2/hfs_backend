package productstorage

import (
	"context"
	"hfs_backend/common"
	"hfs_backend/modules/product/productmodel"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context, conditions map[string]interface{}, filter *productmodel.ProductFilter, paging *common.Paging, moreKeys ...string) ([]productmodel.Product, error) {

	var result []productmodel.Product
	db := s.db
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(productmodel.Product{}.TableName()).
		Where(conditions).
		Where("status in (1)")

	if v := filter; v != nil {
		if v.Id > 0 {
			db = db.Where("id = ?", v.Id)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

  if v := paging.FakeCursor; v!= ""{
    if uid,error := common.FromBase58(v);error != nil{
      db = db.Where("id < ?",uid.GetlocalID())
    }else {
      db = db.Offset((paging.Page - 1) * paging.Limit)
    }
  }

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
