package jfile

import (
	"context"
	"errors"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type TencentConfig struct {
	Zone          string `json:"zone" yaml:"zone"`         // oss 地区
	Bucket        string `json:"bucket" yaml:"bucket"`     // bucket
	Domain        string `json:"imgPath" yaml:"imgPath"`   // 域名
	UseHTTPS      bool   `json:"useHttps" yaml:"useHttps"` // 是有https
	AccessKey     string `json:"accessKey" yaml:"accessKey"`
	SecretKey     string `json:"secretKey" yaml:"secretKey"`
	UseCname      bool   `json:"useCname" yaml:"useCname"` // 使用自定义域名
	UseCdnDomains bool   `json:"useCdnDomains" yaml:"useCdnDomains"`
}
type TencentUploader struct {
	Config TencentConfig
	client *cos.Client
}

func newTencentUploader(config interface{}) (TencentUploader, error) {
	conf, ok := config.(TencentConfig)
	if !ok {
		return TencentUploader{}, errors.New("invalid config for Tencent uploader")
	}

	// 创建OSSClient实例
	scheme := "http://"
	if conf.UseHTTPS {
		scheme = "https://"
	}
	u, _ := url.Parse(scheme + conf.Bucket + "." + conf.Zone)
	su, _ := url.Parse(scheme + conf.Zone)
	client := cos.NewClient(
		&cos.BaseURL{
			BucketURL:  u,
			ServiceURL: su,
		},
		&http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  conf.AccessKey,
				SecretKey: conf.SecretKey,
			},
		},
	)

	return TencentUploader{Config: conf, client: client}, nil
}

func (t TencentUploader) UploadFile(file *multipart.FileHeader, directory string) (string, error) {

	f, openError := file.Open()
	if openError != nil {
		return "", openError
	}
	defer f.Close()

	fileName, _ := generateFileName(file.Filename)
	filePath := directory + "/" + fileName
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: "text/html",
		},
	}
	_, err := t.client.Object.Put(context.Background(), filePath, f, opt)
	if err != nil {
		panic(err)
	}

	return filePath, nil
}

func (t TencentUploader) UploadFiles(files []*multipart.FileHeader, directory string) ([]string, error) {
	return []string{}, nil
}

func (t TencentUploader) GetUploadProgress() float64 {

	return 0
}

func (t TencentUploader) GetUrl(in string) (string, error) {
	oUrl, err := t.client.Object.GetPresignedURL(context.Background(), http.MethodGet, in, t.Config.AccessKey, t.Config.SecretKey, time.Hour, nil, false)
	if err != nil {
		return in, err
	}
	fileUrl := oUrl.String()

	return fileUrl, nil
}
