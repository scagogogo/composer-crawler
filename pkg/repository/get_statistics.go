package repository

import (
	"context"
	"fmt"
	"github.com/scagogogo/composer-crawler/pkg/response"
)

// Statistics 获取仓库中的组件下载统计信息
func (x *Repository) Statistics(ctx context.Context) (*response.StatisticsResponse, error) {
	targetUrl := fmt.Sprintf("%s/statistics.json", x.options.ServerUrl)
	return getJson[*response.StatisticsResponse](ctx, x, targetUrl)
}
