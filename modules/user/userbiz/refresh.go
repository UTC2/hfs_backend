package userbiz

import (
	"context"

	"hfs_backend/common"
	"hfs_backend/component/tokenprovider"
	"hfs_backend/modules/user/usermodel"
)

type refreshBusiness struct {
	tokenProvider tokenprovider.Provider
	accessExpiry  int
	refreshExpiry int
}

func NewRefreshBusiness(tp tokenprovider.Provider, accessExp, refreshExp int) *refreshBusiness {
	return &refreshBusiness{tokenProvider: tp, accessExpiry: accessExp, refreshExpiry: refreshExp}
}

func (b *refreshBusiness) Refresh(_ context.Context, refreshToken string) (*usermodel.Account, error) {
	payload, err := b.tokenProvider.Validate(refreshToken)
	if err != nil {
		return nil, err
	}
	access, err := b.tokenProvider.Generate(*payload, b.accessExpiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	refresh, err := b.tokenProvider.Generate(*payload, b.refreshExpiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return usermodel.NewAccount(access, refresh), nil
}
