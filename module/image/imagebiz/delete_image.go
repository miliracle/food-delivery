package imagebiz

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/component/uploadprovider"
	"fooddelivery/module/image/imagemodel"
	"fooddelivery/module/restaurant/restaurantmodel"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
)

type DeleteImageStorage interface {
	DeleteImage(ctx context.Context, id int) error
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*common.Image, error)
}

type deleteImageBiz struct {
	provider uploadprovider.UploadProvider
	imgStore DeleteImageStorage
}

func NewDeleteImageBiz(provider uploadprovider.UploadProvider, imgStore DeleteImageStorage) *deleteImageBiz {
	return &deleteImageBiz{provider: provider, imgStore: imgStore}
}

func (biz *deleteImageBiz) DeleteImage(ctx context.Context, id int) error {
	image, err := biz.imgStore.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(imagemodel.EntityName, err)
	}
	lenDNSHostUrl := len([]rune(common.CONFIG.GOOGLE_CDN))
	destination := image.Url[lenDNSHostUrl:]

	if err := biz.provider.DeleteUploadedFile(ctx, destination); err != nil {
		log.Fatalln("Failed to delete image", destination)
	}

	if err := biz.imgStore.DeleteImage(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
