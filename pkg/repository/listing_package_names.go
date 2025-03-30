package repository

import (
	"context"
	"fmt"
)

// PackageListResponse 表示返回的包列表
type PackageListResponse struct {
	PackageNames []string `json:"packageNames"`
}

// GET https://packagist.org/packages/list.json
//
//	{
//	 "packageNames": [
//	   "[vendor]/[package]",
//	   ...
//	 ]
//	}
func (x *Repository) List(ctx context.Context) ([]*Package, error) {
	targetUrl := fmt.Sprintf("%s/packages/list.json", x.options.ServerUrl)
	response, err := getJson[*PackageListResponse](ctx, x, targetUrl)
	if err != nil {
		return nil, err
	}

	result := make([]*Package, 0, len(response.PackageNames))
	for _, name := range response.PackageNames {
		result = append(result, &Package{Name: name})
	}
	return result, nil
}
