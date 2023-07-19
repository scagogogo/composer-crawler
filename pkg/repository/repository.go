package repository

import (
	"context"
	"encoding/json"
	"github.com/crawler-go-go-go/go-requests"
)

// Repository 表示一个 composer 仓库
// 仓库API文档：
// https://packagist.org/apidoc#best-practices
type Repository struct {
	options *Options
}


func getJson[T any](ctx context.Context, repository *Repository, targetUrl string) (T, error) {
	bytes, err := repository.getBytes(ctx, targetUrl)
	if err != nil {
		var zero T
		return zero, err
	}
	return unmarshalJson[T](bytes)
}

func unmarshalJson[T any](bytes []byte) (T, error) {
	var r T
	err := json.Unmarshal(bytes, &r)
	if err != nil {
		var zero T
		return zero, err
	}
	return r, nil
}

// 内部使用统一的方法来请求
func (x *Repository) getBytes(ctx context.Context, targetUrl string) ([]byte, error) {
	options := requests.NewOptions[any, []byte](targetUrl, requests.BytesResponseHandler())
	if x.options.Proxy != "" {
		options.AppendRequestSetting(requests.RequestSettingProxy(x.options.Proxy))
	}
	return requests.SendRequest[any, []byte](ctx, options)
}


