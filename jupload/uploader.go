package jfile

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"mime/multipart"
	"path"
	"strings"
	"time"
)

type UploaderType int32

const (
	// UploaderTypeLocal 本地
	UploaderTypeLocal UploaderType = 1
	// UploaderTypeAliyun 阿里云
	UploaderTypeAliyun UploaderType = 2
	// UploaderTypeTencent 腾讯云
	UploaderTypeTencent UploaderType = 3
)

// Uploader represents the interface for file upload functionality.
type Uploader interface {
	UploadFile(file *multipart.FileHeader, directory string) (string, error)

	UploadFiles(files []*multipart.FileHeader, directory string) ([]string, error)

	GetUploadProgress() float64

	GetUrl() string
}

func NewUploader(uploaderType UploaderType, config interface{}) (Uploader, error) {
	switch uploaderType {
	case UploaderTypeLocal:
		return nil, errors.New("local uploader did not realize")
	case UploaderTypeAliyun:
		aliyunConfig, ok := config.(AliyunConfig)
		if !ok {
			return nil, fmt.Errorf("invalid config for Aliyun uploader")
		}
		return AliyunUploader{Config: aliyunConfig}, nil
	case UploaderTypeTencent:
		return nil, errors.New("local uploader did not realize")
	default:
		return nil, errors.New("uploader type")
	}
}

func generateFileName(originalName string) (fileName string, fileExt string) {
	ext := strings.ToLower(path.Ext(originalName))

	name := strings.TrimSuffix(originalName, ext)

	h := md5.New()
	h.Write([]byte(name))
	newName := hex.EncodeToString(h.Sum(nil))

	return newName + "_" + time.Now().Format("20060102150405.000000000") + ext, strings.TrimLeft(ext, ".")
}
