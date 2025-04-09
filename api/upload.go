package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mundo-wang/wtool/wlog"
	"im-chat/code"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func Upload(c *gin.Context) (interface{}, error) {
	file, err := c.FormFile("avatar")
	if err != nil {
		wlog.Error("call c.FormFile failed").Err(err).Log()
		return nil, err
	}
	ext := filepath.Ext(file.Filename)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fileName := fmt.Sprintf("%d%04d%s", time.Now().Unix(), r.Int31(), ext)
	filePath := fmt.Sprintf("asset/upload/%s", fileName)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		wlog.Error("call c.SaveUploadedFile failed").Err(err).Field("filePath", filePath).Log()
		return nil, err
	}
	return fileName, nil
}

func Download(c *gin.Context) (string, error) {
	fileName := c.Query("fileName")
	filePath := fmt.Sprintf("asset/upload/%s", fileName)
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", code.FileNotExist
		}
		wlog.Error("call os.Stat failed").Err(err).Field("fileName", fileName).Log()
		return "", err
	}
	return filePath, nil
}
