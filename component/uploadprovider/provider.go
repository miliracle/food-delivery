package uploadprovider

import (
	"context"
	"fooddelivery/common"
	"io"
)

type UploadProvider interface {
	SaveUploadedFile(ctx context.Context, srcData io.Reader, filename string) (*common.Image, error)
	DeleteUploadedFile(ctx context.Context, destination string) error
}
