# GitHub Actions 工作流程

本项目使用 GitHub Actions 来自动化测试和验证示例代码。当代码推送到主分支或创建 Pull Request 时，工作流程会自动运行。

## 工作流程内容

GitHub Actions 工作流程 (`.github/workflows/go-tests.yml`) 执行以下任务：

1. **设置环境** - 安装 Go 1.20 和所需依赖
2. **运行单元测试** - 执行 `pkg` 目录下的所有测试
3. **生成覆盖率报告** - 创建测试覆盖率报告并上传为工件
4. **验证示例代码** - 编译所有示例代码以确保它们没有语法或编译错误

## 本地测试 GitHub Actions

您可以在本地测试 GitHub Actions 工作流程，无需将代码推送到 GitHub。这可以通过 [act](https://github.com/nektos/act) 工具实现。

### 安装 act

**macOS**:
```bash
brew install act
```

**Linux**:
```bash
curl -s https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash
```

**Windows (使用 Chocolatey)**:
```bash
choco install act-cli
```

### 使用预配置脚本运行 act

本项目包含一个预配置脚本 `run-act.sh`，它设置了正确的环境并运行 act：

```bash
./run-act.sh
```

### 手动运行 act

如果您想手动运行 act，可以使用以下命令：

```bash
# 使用 .actrc 中的配置
act
```

或者指定具体的配置：

```bash
act -P ubuntu-latest=catthehacker/ubuntu:act-latest -j test
```

### 配置文件

本项目包含以下 act 配置文件：

- `.actrc` - act 的基本配置
- `.github/act.json` - 环境配置

## 故障排除

如果在本地运行 GitHub Actions 时遇到问题：

1. **Docker 相关错误** - 确保 Docker 已安装且正在运行
2. **内存不足** - 增加分配给 Docker 的内存
3. **权限问题** - 在类 Unix 系统上可能需要使用 `sudo` 运行 act
4. **路径问题** - 确保从项目根目录运行 act

## CI/CD 集成注意事项

- GitHub Actions 配置为在 `main` 和 `master` 分支上的推送和拉取请求时运行
- 测试覆盖率报告保存在每次构建的工件中
- 示例代码验证只检查编译，不实际运行示例

## 自定义工作流程

如果需要自定义工作流程：

1. 编辑 `.github/workflows/go-tests.yml` 文件
2. 添加新步骤或修改现有步骤
3. 使用 act 在本地测试更改
4. 提交并推送更改 