package jfile

import "io"

type AliyunConfig struct {
	AccessKey  string
	SecretKey  string
	BucketName string
}

type AliyunUploader struct {
	Config AliyunConfig
}

func (a AliyunUploader) UploadFile(file io.Reader, filename string) (string, error) {

	return "", nil
}

func (a AliyunUploader) UploadFiles(files []io.Reader, filenames []string) ([]string, error) {
	return []string{}, nil
}

func (a AliyunUploader) GetUploadProgress() float64 {

	return 0
}
