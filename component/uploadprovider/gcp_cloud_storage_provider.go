package uploadprovider

import (
	"context"
	"fooddelivery/common"
	"io"
	"net/url"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type GCPCloudStorageProvider struct {
	BucketName         string
	CredentialFilePath string
	StorageClient      *storage.Client
}

func NewGCPCloudStorageProvider(bucketName string, credentialFilePath string) (*GCPCloudStorageProvider, error) {
	provider := &GCPCloudStorageProvider{
		BucketName:         bucketName,
		CredentialFilePath: credentialFilePath,
	}

	ctx := context.Background()

	newStorageClient, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))

	if err != nil {
		panic(err)
	}

	// defer newStorageClient.Close()

	provider.StorageClient = newStorageClient

	return provider, nil
}

func (provider *GCPCloudStorageProvider) SaveUploadedFile(ctx context.Context, srcData io.Reader, fileName string) (*common.Image, error) {

	writer := provider.StorageClient.Bucket(provider.BucketName).Object(fileName).NewWriter(ctx)

	if _, err := io.Copy(writer, srcData); err != nil {
		return nil, err
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	u, err := url.Parse("/" + writer.Attrs().Name)
	if err != nil {
		return nil, err
	}

	img := common.Image{
		Url: common.CONFIG.GOOGLE_CDN + u.EscapedPath(),
	}

	return &img, nil

}
