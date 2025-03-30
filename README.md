# 🎼 Composer Crawler

<div align="center">

![Version](https://img.shields.io/badge/version-1.0.0-blue)
![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go&logoColor=white)
![License](https://img.shields.io/badge/license-MIT-green)
![Test Status](https://img.shields.io/badge/tests-passing-brightgreen)

</div>

Composer Crawler 是一个强大的 Go 语言库，专为与 Composer 包管理器仓库交互而设计。它提供了简洁而全面的 API，允许您轻松获取包信息、统计数据和安全公告等。

<p align="center">
<img src="https://getcomposer.org/img/logo-composer-transparent.png" width="200" alt="Composer Logo">
</p>

<div align="center">

```bash
$ go get github.com/scagogogo/composer-crawler
```

</div>

<p align="center">
  <img src="docs/assets/composer-crawler-demo.png" alt="Composer Crawler Demo">
</p>

## 📋 目录

- [功能特点](#-功能特点)
- [为什么选择 Composer Crawler?](#-为什么选择-composer-crawler)
- [快速开始](#-快速开始)
- [安装指南](#-安装指南)
- [使用方法](#-使用方法)
- [API 文档](#-api-文档)
  - [初始化仓库](#初始化仓库)
  - [下载索引](#下载索引)
  - [列出包](#列出包)
  - [获取统计数据](#获取统计数据)
  - [安全公告](#安全公告)
- [项目结构](#-项目结构)
- [示例代码](#-示例代码)
- [自动化测试](#-自动化测试)
- [贡献指南](#-贡献指南)
- [许可证](#-许可证)

## ✨ 功能特点

- 📦 **下载包索引** - 快速获取完整的 Composer 包列表
- 📋 **列出包信息** - 便捷访问所有可用的包及其详细信息
- 📊 **获取统计数据** - 接收仓库下载量、包数量等关键指标
- 🛡️ **安全公告检测** - 获取并分析 Composer 包的安全漏洞信息
- 🔄 **代理支持** - 通过配置代理在受限网络环境中使用
- 🧩 **泛型支持** - 利用 Go 泛型提供类型安全的 API

## 🤔 为什么选择 Composer Crawler?

在构建需要与 Composer 包仓库交互的 Go 应用时，您可能会遇到以下挑战：

- 🔍 **直接使用 HTTP 客户端**调用 Packagist API 繁琐且易出错
- 📊 **手动解析 JSON 响应**需要大量重复代码
- 🛠️ **缺乏类型安全**让开发体验不佳
- 🌐 **网络限制**在某些环境中阻碍直接访问

Composer Crawler 通过以下方式解决这些问题：

- ✅ **简洁的 API** - 用几行代码完成复杂任务
- ✅ **完整的类型定义** - 提供类型安全的响应模型
- ✅ **内置错误处理** - 优雅处理网络和解析错误
- ✅ **代理支持** - 在受限网络环境中也能工作
- ✅ **经过全面测试** - 确保可靠性和稳定性

无论您是需要监控 PHP 依赖的安全漏洞，还是构建 Composer 包的统计分析工具，Composer Crawler 都能帮助您更快速、更可靠地完成任务。

## 🚀 快速开始

只需几行代码即可获取 Composer 仓库的统计信息：

```go
package main

import (
    "context"
    "fmt"
    "github.com/scagogogo/composer-crawler/pkg/repository"
)

func main() {
    // 初始化仓库客户端
    repo := &repository.Repository{}
    
    // 获取统计数据
    stats, err := repo.Statistics(context.Background())
    if err != nil {
        fmt.Printf("获取统计数据失败: %v\n", err)
        return
    }
    
    // 打印统计信息
    fmt.Printf("Composer 统计数据:\n")
    fmt.Printf("总下载量: %d\n", stats.Totals.Downloads)
    fmt.Printf("包数量: %d\n", stats.Totals.Packages)
    fmt.Printf("版本数量: %d\n", stats.Totals.Versions)
}
```

## 💻 安装指南

使用 Go 模块安装 Composer Crawler:

```bash
go get github.com/scagogogo/composer-crawler
```

## 🚀 使用方法

Composer Crawler 设计简洁明了，以下是基本使用流程：

```go
import (
    "context"
    "github.com/scagogogo/composer-crawler/pkg/repository"
)

func main() {
    // 初始化仓库客户端
    options := &repository.Options{
        ServerUrl: "https://packagist.org",
    }
    repo := &repository.Repository{}
    
    // 使用客户端访问 API
    ctx := context.Background()
    
    // 示例：获取统计数据
    stats, err := repo.Statistics(ctx)
    // 处理数据...
}
```

更多详细用法请查看 [示例代码](#-示例代码) 和 [API 文档](#-api-文档)。

## 📘 API 文档

> **注意**: Composer Crawler 基于 [Packagist API](https://packagist.org/apidoc) 构建，提供了更易用、类型安全的 Go 语言接口。

### 初始化仓库

创建一个仓库客户端实例来与 Composer 仓库交互：

```go
// 创建选项
options := &repository.Options{
    ServerUrl: "https://packagist.org", // Composer 仓库 URL
    Proxy: "http://your-proxy:port",    // 可选：设置代理
}

// 创建仓库客户端
repo := &repository.Repository{}
```

### 下载索引

获取完整的 Composer 包索引：

```go
// 方法 1：直接下载索引
indexBytes, err := repository.DownloadIndex(ctx)
if err != nil {
    // 处理错误
}

// 方法 2：下载索引到文件
err = repository.DownloadIndexToFile(ctx, "/path/to/save/index.json")
if err != nil {
    // 处理错误
}
```

### 列出包

获取 Composer 仓库中所有可用的包：

```go
// 获取所有包列表
packages, err := repo.List(ctx)
if err != nil {
    // 处理错误
}

// 遍历包
for _, pkg := range packages {
    fmt.Println(pkg.Name) // 输出包名
}
```

### 获取统计数据

获取 Composer 仓库的统计数据，包括下载量、包数量和版本数量：

```go
// 获取统计数据
stats, err := repo.Statistics(ctx)
if err != nil {
    // 处理错误
}

// 访问统计信息
fmt.Printf("总下载量: %d\n", stats.Totals.Downloads)
fmt.Printf("包数量: %d\n", stats.Totals.Packages)
fmt.Printf("版本数量: %d\n", stats.Totals.Versions)
```

### 安全公告

获取包的安全漏洞公告信息：

```go
// 方法 1：获取某个时间点后的所有安全公告
// 例如：获取一个月前至今的公告
oneMonthAgo := time.Now().AddDate(0, -1, 0)
advisories, err := repo.ListSecurityAdvisories(ctx, oneMonthAgo)
if err != nil {
    // 处理错误
}

// 遍历安全公告
for packageName, pkgAdvisories := range advisories.Advisories {
    for _, advisory := range pkgAdvisories {
        fmt.Printf("包: %s, 漏洞: %s, CVE: %s\n", 
            packageName, advisory.Title, advisory.Cve)
    }
}

// 方法 2：获取特定包的安全公告
packageName := "symfony/http-kernel"
packageAdvisories, err := repo.ListAdvisories(ctx, packageName)
if err != nil {
    // 处理错误
}

// 遍历特定包的安全公告
for _, advisory := range packageAdvisories {
    fmt.Printf("漏洞: %s, CVE: %s, 报告时间: %s\n", 
        advisory.Title, advisory.Cve, advisory.ReportedAt)
}
```

## 📁 项目结构

```
composer-crawler/
├── .github/              # GitHub 配置
│   └── workflows/        # GitHub Actions 工作流程
├── docs/                 # 详细文档
├── examples/             # 使用示例
│   ├── 01_basic_setup/   # 基本设置示例
│   ├── 02_download_index/# 下载索引示例
│   ├── 03_list_packages/ # 列出包示例
│   ├── 04_get_statistics/# 获取统计示例
│   └── 05_security_advisories/ # 安全公告示例
├── pkg/                  # 包目录
│   ├── repository/       # 仓库交互实现
│   └── response/         # API 响应模型
└── run-act.sh            # 用于本地测试 GitHub Actions
```

## 📝 示例代码

项目包含一系列完整示例，展示如何使用各种功能：

| 示例 | 描述 |
|------|------|
| [01_basic_setup](examples/01_basic_setup/) | 演示如何初始化仓库客户端 |
| [02_download_index](examples/02_download_index/) | 展示如何下载完整的包索引 |
| [03_list_packages](examples/03_list_packages/) | 说明如何获取和处理包列表 |
| [04_get_statistics](examples/04_get_statistics/) | 展示如何获取仓库统计信息 |
| [05_security_advisories](examples/05_security_advisories/) | 演示如何获取安全漏洞信息 |

运行示例：

```bash
cd examples/01_basic_setup
go run main.go
```

## 🧪 自动化测试

### 运行单元测试

```bash
go test -v ./pkg/...
```

### GitHub Actions

本项目使用 GitHub Actions 进行自动化测试。当代码推送到主分支或创建 Pull Request 时，测试会自动运行。

要在本地测试 GitHub Actions 工作流程，可以使用 [act](https://github.com/nektos/act)：

```bash
./run-act.sh
```

更多详情请参阅 [GitHub Actions 文档](docs/github-actions.md)。

## 👥 贡献指南

欢迎贡献代码、报告问题或提出改进建议！请按照以下步骤：

1. Fork 本仓库
2. 创建您的特性分支: `git checkout -b feature/amazing-feature`
3. 提交您的更改: `git commit -m 'Add some amazing feature'`
4. 推送到分支: `git push origin feature/amazing-feature`
5. 创建 Pull Request

在提交 PR 前，请确保通过所有测试。

## 📄 许可证

本项目采用 [MIT License](LICENSE) 许可证 - 详情请见 LICENSE 文件。

---

<div align="center">
  <sub>用 ❤️ 构建 | 反馈问题请创建 <a href="https://github.com/scagogogo/composer-crawler/issues">Issue</a></sub>
</div> 