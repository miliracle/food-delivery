package userbiz

import (
	"context"
	"errors"
	"fmt"
	"fooddelivery/common"
	"fooddelivery/module/user/usermodel"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(context context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type RegisterBiz struct {
	store  RegisterStorage
	hasher Hasher
}

func NewRegisterBiz(store RegisterStorage, hasher Hasher) *RegisterBiz {
	return &RegisterBiz{store: store, hasher: hasher}
}

func (biz *RegisterBiz) RegisterUser(ctx context.Context, data *usermodel.UserCreate) error {
	if data.Email == "" {
		return errors.New("email should not empty")
	}

	userFound, _ := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if userFound != nil {
		return usermodel.ErrEmailExisted
	}
	salt := common.GenSalt(50)

	fmt.Print("salt: ", salt)
	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user" // hard code

	if err := biz.store.CreateUser(ctx, data); err != nil {
		return err
	}
	return nil
}
