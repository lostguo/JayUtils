package jfile

import (
	"errors"
	"mime/multipart"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliyunConfig struct {
	Zone          string `json:"zone" yaml:"zone"`
	Bucket        string `json:"bucket" yaml:"bucket"`
	ImgPath       string `json:"imgPath" yaml:"imgPath"`
	UseHTTPS      bool   `json:"useHttps" yaml:"useHttps"`
	AccessKey     string `json:"accessKey" yaml:"accessKey"`
	SecretKey     string `json:"secretKey" yaml:"secretKey"`
	UseCname      bool   `json:"useCname" yaml:"useCname"`
	UseCdnDomains bool   `json:"useCdnDomains" yaml:"useCdnDomains"`
}

type AliyunUploader struct {
	Config AliyunConfig
	client *oss.Client
}

func newAliyunUploader(config interface{}) (AliyunUploader, error) {
	conf, ok := config.(AliyunConfig)
	if !ok {
		return AliyunUploader{}, errors.New("invalid config for Tencent uploader")
	}

	client, err := oss.New(conf.Zone, conf.AccessKey, conf.SecretKey, oss.UseCname(conf.UseCname))
	if err != nil {
		return AliyunUploader{}, err
	}

	return AliyunUploader{Config: conf, client: client}, nil
}

func (a AliyunUploader) UploadFile(file *multipart.FileHeader, directory string) (string, error) {
	fileName, _ := generateFileName(file.Filename)
	filePath := directory + string(os.PathSeparator) + fileName
	bucket, err := a.client.Bucket(a.Config.Bucket)
	if err != nil {
		return "", err
	}

	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	if err := bucket.PutObject(filePath, f); err != nil {
		return "", err
	}

	return filePath, nil
}

func (a AliyunUploader) UploadFiles(files []*multipart.FileHeader, directory string) ([]string, error) {
	return []string{}, nil
}

func (a AliyunUploader) GetUploadProgress() float64 {

	return 0
}

func (a AliyunUploader) GetUrl(in string) (string, error) {

	bucket, err := a.client.Bucket(a.Config.Bucket)
	if err != nil {
		return "", err
	}
	signedUrl, err := bucket.SignURL(in, oss.HTTPGet, 600, oss.Process(""))
	if err != nil {
		return "", err
	}

	return signedUrl, nil
}
