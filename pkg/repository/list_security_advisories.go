package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/scagogogo/composer-crawler/pkg/response"
)

// List security advisories
// https://packagist.org/apidoc#list-security-advisories

// https://packagist.org/api/security-advisories/?updatedSince=[timestamp]&packages[]=[vendor/package]

// ListSecurityAdvisories 查询给定时间之后被报告的漏洞
// https://packagist.org/api/security-advisories/?updatedSince=2023-05-22 19:49:11
func (x *Repository) ListSecurityAdvisories(ctx context.Context, updatedSince time.Time) (*response.AdvisoriesResponse, error) {
	targetUrl := fmt.Sprintf("%s/api/security-advisories/?updatedSince=%d", x.options.ServerUrl, updatedSince.UnixMilli())
	return getJson[*response.AdvisoriesResponse](ctx, x, targetUrl)
}

// ListAdvisories 获取给定包上的所有漏洞
// https://packagist.org/api/security-advisories/?packages=craftcms/cms
func (x *Repository) ListAdvisories(ctx context.Context, packageName string) ([]*response.Advisory, error) {
	targetUrl := fmt.Sprintf("%s/api/security-advisories/?packages=%s", x.options.ServerUrl, packageName)
	json, err := getJson[*response.AdvisoriesResponse](ctx, x, targetUrl)
	if err != nil {
		return nil, err
	}
	return json.Advisories[packageName], nil
}
