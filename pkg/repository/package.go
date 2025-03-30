package repository

// 拼接包名后面加json的形式，可以直接拿到包的json格式的信息
// https://packagist.org/packages/symfony/console.json

// Package 表示一个 composer 包
type Package struct {
	Name string
}
