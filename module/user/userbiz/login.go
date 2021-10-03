package userbiz

import (
	"context"
	"fooddelivery/common"
	component "fooddelivery/component/appctx"
	"fooddelivery/component/tokenprovider"
	"fooddelivery/module/user/usermodel"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type LoginBiz struct {
	appCtx        component.AppContext
	store         LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
}

func NewLoginBiz(store LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher) *LoginBiz {
	return &LoginBiz{store: store, tokenProvider: tokenProvider, hasher: hasher}
}

func (biz *LoginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*usermodel.Account, error) {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, common.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	passHashed := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, common.CONFIG.ACCESS_TOKEN_EXPIRE_TIME)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, err := biz.tokenProvider.Generate(payload, common.CONFIG.REFRESH_TOKEN_EXPIRE_TIME)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := usermodel.NewAccount(accessToken, refreshToken)

	return account, nil
}
