package models

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Config struct {
	Entrypoint string
	Port       string
	Bucket     string
	AccessKey  string
	SecretKey  string
}

var CurrentConfig Config

// MinioConnection func for opening minio connection.
func MinioConnection() (*minio.Client, error) {
	useSSL := true
	// Initialize minio client object.
	minioClient, err := minio.New(CurrentConfig.Entrypoint, &minio.Options{
		Creds:  credentials.NewStaticV4(CurrentConfig.AccessKey, CurrentConfig.SecretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Print(err)
	}

	return minioClient, err
}
