package utils

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/mundo-wang/wtool/wlog"
	"sync"
)

var (
	minioClient *minio.Client
	minioOnce   sync.Once
)

func NewMinioClient(bucketName string) (*minio.Client, error) {
	minioOnce.Do(func() {
		var err error
		minioClient, err = minio.New(Config.Minio.Endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(Config.Minio.AccessKey, Config.Minio.SecretKey, ""),
			Secure: Config.Minio.UseSSL,
		})
		if err != nil {
			wlog.Fatal("call minio.New failed").Err(err).Log()
		}
	})
	exists, err := minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		wlog.Error("call minioClient.BucketExists failed").Err(err).Log()
		return nil, err
	}
	if !exists {
		err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			wlog.Error("call minioClient.MakeBucket failed").Err(err).Log()
			return nil, err
		}
	}
	return minioClient, nil
}
