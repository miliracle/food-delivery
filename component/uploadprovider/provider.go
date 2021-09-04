package uploadprovider

import (
	"context"
	"fooddelivery/common"
	"io"
)

type UploadProvider interface {
	SaveUploadedFile(ctx context.Context, srcData io.Reader, filename string) (*common.Image, error)
}
