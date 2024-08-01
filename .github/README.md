# Cnos-Telegraf

CnosDB-Telegraf 基于 Telegraf 进行开发，增加了一些功能与插件。

## 原版 Telegraf 文档

[README.md](../README.md)

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
