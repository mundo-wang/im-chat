package main

import (
	"github.com/mundo-wang/wtool/wlog"
	"im-chat/utils"
)

func main() {
	err := utils.InitConfig()
	if err != nil {
		wlog.Fatal("call utils.InitConfig failed").Err(err).Log()
	}
	bucketName := utils.Config.Minio.Bucket
	minioClient, err := utils.NewMinioClient(bucketName)
	if err != nil {
		wlog.Fatal("call utils.NewMinioClient failed").Err(err).Log()
	}
	_ = minioClient
}
