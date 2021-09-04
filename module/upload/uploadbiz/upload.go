package uploadbiz

import (
	"context"
	"fmt"
	"fooddelivery/common"
	"fooddelivery/component/uploadprovider"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type CreateImageStorage interface {
	CreateImage(context context.Context, data *common.Image) error
}

type uploadBiz struct {
	provider uploadprovider.UploadProvider
	imgStore CreateImageStorage
}

func NewUploadBiz(provider uploadprovider.UploadProvider, imgStore CreateImageStorage) *uploadBiz {
	return &uploadBiz{provider: provider, imgStore: imgStore}
}

func (biz *uploadBiz) UploadImage(ctx context.Context, srcData io.Reader, folder string, filename string) (*common.Image, error) {
	width, height, err := getImageDimension(srcData)

	if err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "image"
	}

	fileExt := filepath.Ext(filename)
	subFolderName := fmt.Sprintf("%d", time.Now().Nanosecond())
	destination := fmt.Sprintf("%s/%s/%s", folder, subFolderName, filename)

	img, err := biz.provider.SaveUploadedFile(ctx, srcData, destination)

	img.Width = width
	img.Height = height
	img.Extension = fileExt
	img.Url = ""

	if err != nil {
		return nil, err
	}

	if err := biz.imgStore.CreateImage(ctx, img); err != nil {
		return nil, common.ErrCannotCreateEntity(img.TableName(), err)
	}

	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("getImageDimension err: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
