package component

import (
	"fooddelivery/component/uploadprovider"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, uploadprovider uploadprovider.UploadProvider) *appCtx {
	return &appCtx{
		db:             db,
		uploadProvider: uploadprovider,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadProvider
}
