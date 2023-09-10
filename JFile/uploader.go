package jfile

import (
	"errors"
	"fmt"
	"io"
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
	UploadFile(file io.Reader, filename string) (string, error)

	UploadFiles(files []io.Reader, filenames []string) ([]string, error)

	GetUploadProgress() float64
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
