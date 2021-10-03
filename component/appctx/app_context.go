package appctx

import (
	"fooddelivery/component/uploadprovider"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
}

func NewAppContext(
	db *gorm.DB,
	uploadprovider uploadprovider.UploadProvider,
	secretKey string) *appCtx {
	return &appCtx{
		db:             db,
		uploadProvider: uploadprovider,
		secretKey:      secretKey,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadProvider
}

func (ctx *appCtx) SecretKey() string { return ctx.secretKey }
