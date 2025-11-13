package config

import (
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)



func NewBlobClient() (*azblob.Client, error) {
	connectionString := os.Getenv("blob_connection")
	client, err := azblob.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}
