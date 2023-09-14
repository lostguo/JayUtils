package jfile

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
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

	GetUrl(in string) (string, error)
}

func NewUploader(uploaderType UploaderType, config interface{}) (Uploader, error) {
	switch uploaderType {
	case UploaderTypeLocal:
		return newLocalUploader(config)
	case UploaderTypeAliyun:
		return newAliyunUploader(config)
	case UploaderTypeTencent:
		return newTencentUploader(config)
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
