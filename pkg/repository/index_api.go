package repository

import (
	"context"
	"os"

	"github.com/crawler-go-go-go/go-requests"
)

// DownloadIndex 下载整个的索引文件
// https://packagist.org/packages/list.json
func DownloadIndex(ctx context.Context) ([]byte, error) {
	targetUrl := "https://packagist.org/packages/list.json"
	return requests.GetBytes(ctx, targetUrl)
}

// DownloadIndexToFile 下载索引文件到本地文件
func DownloadIndexToFile(ctx context.Context, filepath string) error {
	indexBytes, err := DownloadIndex(ctx)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath, indexBytes, os.ModePerm)
}
