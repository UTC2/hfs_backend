package productbiz

import (
	"context"
	"hfs_backend/common"
	"hfs_backend/modules/product/productmodel"
  "log"
)

type ListProductStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *productmodel.ProductFilter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]productmodel.Product, error)
}
type LikeStore interface {
  GetProductLikes(ctx context.Context,
    ids []int,
  )(map[int]int,error)
}

type listProductBiz struct {
	store ListProductStore
	like LikeStore
}

func NewListProductBiz(store ListProductStore,like LikeStore) *listProductBiz {
	return &listProductBiz{store: store,like: like}
}

func (biz *listProductBiz) ListProduct(ctx context.Context,
  filter *productmodel.ProductFilter,
  paging *common.Paging) ([]productmodel.Product, error) {

	result, err := biz.store.ListDataByCondition(ctx, nil, filter, paging)

	if err != nil {
    return nil, common.ErrCannotDeleteEntity(productmodel.EntityName,err)
  }

  ids := make([]int,len(result))

  for i :=range result{
    ids[i] = result[i].Id
  }
  mapProductLike,err:=biz.like.GetProductLikes(ctx,ids)
  if err != nil {
    log.Println("Cannot get product likes",err)
  }
  if v :=mapProductLike;v!=nil{
    for i,item := range result{
        result[i].LikeCount = mapProductLike[item.Id]
    }
  }
	return result, err
}
