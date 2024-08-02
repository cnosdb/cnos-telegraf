# Cnos-Telegraf

> 提交 PR 至 telegraf 官方仓库时，请一定注意，不要带上本文件。

CnosDB-Telegraf 基于 Telegraf 进行开发，增加了一些功能与插件。

## 原版 Telegraf 文档

[README.md](../README.md)

## 分支介绍

### 新增功能

- `feature/doc` - 文档。
- `feature/docker` - 构建 telegraf 为 docker 镜像。
- `feature/high_priority_channel` - 使 input 插件能够将 output 插件的处理结果返回给调用者。
- `feature/cnosdb_subscription` - 添加 input 插件用于处理 CnosDB 订阅消息的报文。
- `feature/opentsdb_json` - 添加 input 插件用于处理 JSON 格式的 OpenTSDB 写请求的报文。

### cnosdb

## 如何更新版本

确保 git 记录了官方仓库的信息。

```shell
git remote add -t master telegraf git@github.com:influxdata/telegraf.git
```

更新 master 分支。

```shell
git fetch --all --tags --prune --jobs=10
## 按提交时间获得 tag 列表，找到最新的 tag，假设最新的 tag 为 v1.32.2
git tag --sort=-committerdate
## 获得最新的 tag 对应的 commit id
#git rev-list -n 1 v1.32.2
git checkout master
git reset --hard v1.32.2
```

对所有的 feature 分支执行变基。

```shell
git checkout feature/docker
git rebase master

git checkout feature/opentsdb_json
git rebase master

git checkout feature/high_priority_channel
git rebase master

## feature/cnosdb_subscription 依赖 feature/high_priority_channel
git checkout feature/cnosdb_subscription
git rebase feature/high_priority_channel
```

使用 `git cherry-pick` 将所有的 `feature/` 的功能合并至 `cnosdb/`

```shell
git checkout master
git checkout -b cnosdb/main/$(date +"%Y%m%d_%H%M%S")
git cherry-pick $(git rev-list -n 1 feature/docker)
git cherry-pick $(git rev-list -n 1 feature/high_priority_channel)
git cherry-pick $(git rev-list -n 1 feature/opentsdb_json)
git cherry-pick $(git rev-list -n 1 feature/cnosdb_subscription)
git cherry-pick $(git rev-list -n 1 feature/doc)
```

## Cnos-Telegraf 的改动说明

### Parser Plugin

增加 Parser 插件 opentsdb_json，用于采集 OpenTSDB 的 JSON 格式的写入请求。

#### OpenTSDB JSON

通过使用 Input 插件 http_listener_v2 并配置 `data_format` 为 `"opentsdb"`，将能够解析 OpenTSDB 格式的写入请求。

```toml
[[inputs.http_listener_v2]]
service_address = ":8080"
paths = ["/api/put"]
methods = ["POST", "PUT"]
data_format = "opentsdb_json"
```

### Input Plugin

#### CnosDB 订阅

增加 Input 插件 cnosdb_subscription，用于搜集 CnosDB 订阅功能推送的数据，搭配 Output 插件即可实现异构数据同步。

```toml
[[inputs.cnosdb_subscription]]
service_address = ":8803"
```

- **配置介绍**

| 参数              | 说明   |
|-----------------|------|
| service_address | 监听端口 |

#### 通用参数

增加配置参数 high_priority_io，用于开启端到端模式。

当设置为 true 时，写入的数据将立即发送到 Output 插件，并根据 Output 插件的返回参数来决定返回值。

```toml
[[inputs.http_listener_v2]]
service_address = ":8080"
paths = ["/api/put"]
methods = ["POST", "PUT"]
data_format = "opentsdb"
high_priority_io = true
```

## 构建

1. [安装 Go](https://golang.org/doc/install)，版本要求 >=1.22
2. 从 Github 克隆仓库:

   ```shell
   git clone https://github.com/cnosdb/cnos-telegraf.git
   ```

3. 检出与 CnosDB 对应的版本，在仓库目录下执行 `make build`

   ```shell
   cd cnos-telegraf
   # 如果为了适配 cnosdb/2.4 版本
   git checkout cnosdb/2.4
   make build
   ```

## 启动

执行以下指令，查看用例:

```shell
telegraf --help
```

### 生成一份标准的 telegraf 配置文件

```shell
telegraf config > telegraf.conf
```

### 生成一份 telegraf 配置文件，仅包含 cpu 指标采集 & influxdb 输出两个插件

```shell
telegraf config --section-filter agent:inputs:outputs --input-filter cpu --output-filter influxdb
```

### 运行 telegraf 但是将采集指标输出到标准输出

```shell
telegraf --config telegraf.conf --test
```

### 运行 telegraf 并通过配置文件来管理加载的插件

```shell
telegraf --config telegraf.conf
```

### 运行 telegraf，仅加载 cpu & memory 指标采集，和 influxdb 输出插件

```shell
telegraf --config telegraf.conf --input-filter cpu:mem --output-filter influxdb
```
