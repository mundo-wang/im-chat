package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/mundo-wang/wtool/wlog"
	"im-chat/conf"
	"im-chat/utils"
	"im-chat/utils/wresp"
	"io"
	"math/rand"
	"mime"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	err := conf.InitConfig()
	if err != nil {
		wlog.Fatal("call utils.InitConfig failed").Err(err).Log()
	}
	r := gin.Default()
	r.POST("/put-file", PutFile)
	r.POST("/fput-file", FPutFile)
	r.GET("/get-file", GetFile)
	r.GET("/delete-file", DeleteFile)
	r.POST("/copy-file", CopyFile)
	r.GET("/copy-files-name", ListFilesName)
	r.GET("/get-file-stat", GetFileStat)
	r.GET("/presign-get-url", PresignGetURL)
	r.GET("/presign-put-url", PresignPutURL)
	r.Run(":8280")
}

func PutFile(c *gin.Context) {
	file, _ := c.FormFile("my-file")
	// 获取文件后缀，并生成唯一对象名（该文件对应的MinIO的Key）
	ext := filepath.Ext(file.Filename)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	objectName := fmt.Sprintf("images/%d%05d%s", time.Now().UnixMilli(), rnd.Intn(100000), ext)
	// 打开上传的文件，获取文件流，其中File是io.Reader的子接口
	srcFile, _ := file.Open()
	defer srcFile.Close()
	// 获取minio客户端对象
	bucketName := conf.Config.Minio.Bucket
	minioClient, _ := utils.NewMinioClient(bucketName)
	// 上传文件到MinIO
	fmt.Println(file.Header.Get("Content-Type"))
	_, _ = minioClient.PutObject(c.Request.Context(), bucketName, objectName,
		srcFile, file.Size, minio.PutObjectOptions{
			ContentType: file.Header.Get("Content-Type"), // 选填项，默认为application/octet-stream
		})
	// 返回成功响应，以及对象名
	wresp.OK(c, objectName)
}

func FPutFile(c *gin.Context) {
	filePath := c.Query("file-path")
	ext := filepath.Ext(filePath)
	// 根据文件后缀推测Content-Type，不一定准确
	contentType := mime.TypeByExtension(ext)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	objectName := fmt.Sprintf("images/%d%05d%s", time.Now().UnixMilli(), rnd.Intn(100000), ext)
	bucketName := conf.Config.Minio.Bucket
	minioClient, _ := utils.NewMinioClient(bucketName)
	_, _ = minioClient.FPutObject(c.Request.Context(), bucketName, objectName, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	wresp.OK(c, objectName)
}

func GetFile(c *gin.Context) {
	// 获取请求参数中的对象名
	objectName := c.Query("object-name")
	bucketName := conf.Config.Minio.Bucket
	minioClient, _ := utils.NewMinioClient(bucketName)
	object, err := minioClient.GetObject(c.Request.Context(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
	}
	defer object.Close()
	// 读取对象为字节数组
	data, _ := io.ReadAll(object)
	// 从对象名中提取扩展名，并尝试根据扩展名获取Content-Type
	ext := filepath.Ext(objectName)
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	//c.Header("Content-Type", "application/octet-stream")
	//c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", path.Base(objectName)))
	c.Data(http.StatusOK, contentType, data)
}

func DeleteFile(c *gin.Context) {
	objectName := c.Query("object-name")
	bucketName := conf.Config.Minio.Bucket
	minioClient, _ := utils.NewMinioClient(bucketName)
	// 删除对象为幂等操作，即使objectName不存在，也不会返回error
	_ = minioClient.RemoveObject(c.Request.Context(), bucketName, objectName, minio.RemoveObjectOptions{})
	wresp.OK(c, fmt.Sprintf("文件 %s 已删除", objectName))
}

func CopyFile(c *gin.Context) {
	// 获取源对象和目标对象名称
	srcObjectName := c.Query("src-object-name")
	dstObjectName := c.Query("dst-object-name")
	bucketName := conf.Config.Minio.Bucket
	minioClient, _ := utils.NewMinioClient(bucketName)
	// 构建源对象的配置信息
	src := minio.CopySrcOptions{
		Bucket: bucketName,
		Object: srcObjectName,
	}
	// 构建目标对象的配置信息
	dst := minio.CopyDestOptions{
		Bucket: bucketName,
		Object: dstObjectName,
	}
	// 执行复制操作，如果srcObjectName在minio不存在，会返回error：The specified key does not exist.
	_, _ = minioClient.CopyObject(c.Request.Context(), dst, src)
	wresp.OK(c, fmt.Sprintf("已将对象从%s复制为%s", srcObjectName, dstObjectName))
}

func ListFilesName(c *gin.Context) {
	// 获取筛选对象Key的前缀，可为空，空则列出所有对象
	prefix := c.Query("prefix")
	bucketName := conf.Config.Minio.Bucket
	minioClient, _ := utils.NewMinioClient(bucketName)
	// 构建对象列举配置项
	opts := minio.ListObjectsOptions{
		Prefix:    prefix,
		Recursive: true, // 设置为true以递归列出所有子目录下的对象
	}
	var objectKeys []string
	for object := range minioClient.ListObjects(c.Request.Context(), bucketName, opts) {
		objectKeys = append(objectKeys, object.Key)
	}
	wresp.OK(c, objectKeys)
}

func GetFileStat(c *gin.Context) {
	objectName := c.Query("object-name")
	bucketName := conf.Config.Minio.Bucket
	minioClient, _ := utils.NewMinioClient(bucketName)
	// 获取对象的元信息
	info, err := minioClient.StatObject(c.Request.Context(), bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		fmt.Println(err)
	}
	// 收集info中重要数据，构造响应数据
	meta := map[string]interface{}{
		"object_name":   info.Key,
		"size":          info.Size,
		"etag":          info.ETag,
		"content_type":  info.ContentType,
		"last_modified": info.LastModified.Format(time.DateTime),
		"user_metadata": info.UserMetadata, // 用户自定义元数据
	}
	wresp.OK(c, meta)
}

func PresignGetURL(c *gin.Context) {
	objectName := c.Query("object-name")
	expireStr := c.DefaultQuery("expire-seconds", "3600") // 默认过期时间：1小时
	expireSeconds, _ := strconv.Atoi(expireStr)
	bucketName := conf.Config.Minio.Bucket
	minioClient, _ := utils.NewMinioClient(bucketName)
	reqParams := make(url.Values)
	// 如果想在访问预签名URL时触发浏览器下载，设置该响应头参数
	//reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=%s", path.Base(objectName)))
	// 生成预签名URL
	presignedURL, _ := minioClient.PresignedGetObject(
		c.Request.Context(),
		bucketName,
		objectName,
		time.Duration(expireSeconds)*time.Second,
		reqParams,
	)
	wresp.OK(c, presignedURL.String())
}

func PresignPutURL(c *gin.Context) {
	objectName := c.Query("object-name")
	expireStr := c.DefaultQuery("expire-seconds", "3600") // 默认过期时间：1小时
	expireSeconds, _ := strconv.Atoi(expireStr)
	bucketName := conf.Config.Minio.Bucket
	minioClient, _ := utils.NewMinioClient(bucketName)
	// 生成预签名上传URL（PUT方式）
	presignedURL, _ := minioClient.PresignedPutObject(
		c.Request.Context(),
		bucketName,
		objectName,
		time.Duration(expireSeconds)*time.Second,
	)
	wresp.OK(c, presignedURL.String())
}
