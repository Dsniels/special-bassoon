package storage

import (
	"context"
	"fmt"
	"io"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/bloberror"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
)

type BlobStorage struct {
	*azblob.Client
	ContainerName string
}

type IBlobStorage interface {
	UploadImage(ctx context.Context, image io.Reader, filename string, contentType string) (string, error)
}

func (b *BlobStorage) UploadImage(ctx context.Context, image io.Reader, filename string, contentType string) (string, error) {
	err := b.ensureContainer()
	if err != nil {
		return "", err
	}
	opts := &azblob.UploadStreamOptions{
		BlockSize:   1024 * 2,
		Concurrency: 4,
		HTTPHeaders: &blob.HTTPHeaders{
			BlobContentType: &contentType,
		},
	}
	_, err = b.UploadStream(ctx, b.ContainerName, filename, image, opts)
	if err != nil {
		return "", err
	}
	endpoint := b.URL()
	url := fmt.Sprintf("%s%s/%s", endpoint, b.ContainerName, filename)
	return url, nil
}

func (b *BlobStorage) getContainer() (*container.Client, error) {

	if err := b.ensureContainer(); err != nil {
		return nil, err
	}
	client := b.ServiceClient().NewContainerClient(b.ContainerName)
	return client, nil
}

func (b *BlobStorage) ensureContainer() error {
	_, err := b.CreateContainer(context.TODO(), b.ContainerName, &azblob.CreateContainerOptions{})
	if err != nil {
		if bloberror.HasCode(err, bloberror.ContainerAlreadyExists) {
			return nil
		}
		return err
	}
	client := b.ServiceClient().NewContainerClient(b.ContainerName)
	access := container.PublicAccessTypeContainer
	client.SetAccessPolicy(context.TODO(), &container.SetAccessPolicyOptions{Access: &access})
	return nil
}

func NewBlobStorage(client *azblob.Client) IBlobStorage {
	return &BlobStorage{
		ContainerName: "fixable",
		Client:        client,
	}
}
