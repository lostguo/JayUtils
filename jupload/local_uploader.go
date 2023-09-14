package jfile

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path"
)

type LocalConfig struct {
	Host string
}

type LocalUploader struct {
	Config LocalConfig
}

func newLocalUploader(config interface{}) (LocalUploader, error) {
	localConfig, ok := config.(LocalConfig)
	if !ok {
		return LocalUploader{}, errors.New("invalid config for Aliyun uploader")
	}
	return LocalUploader{Config: localConfig}, nil
}

func (a LocalUploader) UploadFile(file *multipart.FileHeader, directory string) (string, error) {

	fileName, _ := generateFileName(file.Filename)
	filePath := directory + string(os.PathSeparator) + fileName

	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	if _, err := os.Stat(path.Dir(fileName)); os.IsNotExist(err) {
		err = os.MkdirAll(path.Dir(fileName), os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	outFile, err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(outFile, f)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

func (a LocalUploader) UploadFiles(files []*multipart.FileHeader, directory string) ([]string, error) {

	res := []string{}
	for _, item := range files {
		if filePath, err := a.UploadFile(item, directory); err != nil {
			return []string{}, err
		} else {
			res = append(res, filePath)
		}
	}

	return res, nil
}

func (a LocalUploader) GetUploadProgress() float64 {
	return 0
}

// @param in 文件地址
func (a LocalUploader) GetUrl(in string) (string, error) {
	return a.Config.Host + in, nil
}
