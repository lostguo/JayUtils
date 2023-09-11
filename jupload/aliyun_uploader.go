package jfile

import "mime/multipart"

type AliyunConfig struct {
	AccessKey  string
	SecretKey  string
	BucketName string
}

type AliyunUploader struct {
	Config AliyunConfig
}

func (a AliyunUploader) UploadFile(file *multipart.FileHeader, directory string) (string, error) {

	return "", nil
}

func (a AliyunUploader) UploadFiles(files []*multipart.FileHeader, directory string) ([]string, error) {
	return []string{}, nil
}

func (a AliyunUploader) GetUploadProgress() float64 {

	return 0
}

func (a AliyunUploader) GetUrl() string {
	return ""
}
