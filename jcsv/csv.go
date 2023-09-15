package jcsv

import (
	"encoding/csv"
	"os"
	"time"
)

type csvExport struct {
}

func NewCSVExport() csvExport {
	return csvExport{}
}

func (c csvExport) WriteData(outDir string, outName string, data [][]string) (string, error) {

	fileName := outName + "_" + time.Now().Format("20060102150405.000000000") + ".csv"
	filePath := outDir + string(os.PathSeparator) + fileName
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// 写入UTF-8 BOM，避免使用Microsoft Excel打开乱码
	if _, err := f.WriteString("\xEF\xBB\xBF"); err != nil {
		return "", err
	}

	writer := csv.NewWriter(f)
	if err := writer.WriteAll(data); err != nil {
		return "", err
	}

	writer.Flush()

	return filePath, nil
}

func (c csvExport) Destroy(filePath string) error {
	return os.Remove(filePath)
}
