# ![tiger](assets/TelegrafTigerSmall.png "tiger") Telegraf

[![GoDoc](https://img.shields.io/badge/doc-reference-00ADD8.svg?logo=go)](https://godoc.org/github.com/influxdata/telegraf)  [![Docker pulls](https://img.shields.io/docker/pulls/library/telegraf.svg)](https://hub.docker.com/_/telegraf/) [![Go Report Card](https://goreportcard.com/badge/github.com/influxdata/telegraf)](https://goreportcard.com/report/github.com/influxdata/telegraf) [![Circle CI](https://circleci.com/gh/influxdata/telegraf.svg?style=svg)](https://circleci.com/gh/influxdata/telegraf)
# Telegraf

Telegraf is an agent for collecting, processing, aggregating, and writing
metrics, logs, and other arbitrary data.
## 介绍 Telegraf

* Offers a comprehensive suite of over 300 plugins, covering a wide range of
  functionalities including system monitoring, cloud services, and message
  passing
* Enables the integration of user-defined code to collect, transform, and
  transmit data efficiently
* Compiles into a standalone static binary without any external dependencies,
  ensuring a streamlined deployment process
* Utilizes TOML for configuration, providing a user-friendly and unambiguous
  setup experience
* Developed with contributions from a diverse community of over 1,200
  contributors
[README.md](./README.telegraf.md)

Users can choose plugins from a wide range of topics, including but not limited
to:
## 介绍改动

* Devices: [OPC UA][], [Modbus][]
* Logs: [File][], [Tail][], [Directory Monitor][]
* Messaging: [AMQP][], [Kafka][], [MQTT][]
* Monitoring: [OpenTelemetry][], [Prometheus][]
* Networking: [Cisco TelemetryMDT][], [gNMI][]
* System monitoring: [CPU][], [Memory][], [Disk][], [Network][], [SMART][],
  [Docker][], [Nvidia SMI][], etc.
* Universal: [Exec][], [HTTP][], [HTTP Listener][], [SNMP][], [SQL][]
* Windows: [Event Log][], [Management Instrumentation][],
  [Performance Counters][]
### Parser

## 🔨 Installation
增加 Parser 插件 OpenTSDB 和 OpenTSDB-Telnet，用于采集 OpenTSDB 的写入请求。

For binary builds, Docker images, RPM & DEB packages, and other builds of
Telegraf, please see the [install guide](/docs/INSTALL_GUIDE.md).
**OpenTSDB**

See the [releases documentation](/docs/RELEASES.md) for details on versioning
and when releases are made.
通过使用 Input 插件 http_listener_v2 并配置 `data_format` 为 `"opentsdb"`，将能够解析 OpenTSDB 格式的写入请求。

## 💻 Usage

Users define a TOML configuration with the plugins and settings they wish to
use, then pass that configuration to Telegraf. The Telegraf agent then
collects data from inputs at each interval and sends data to outputs at each
flush interval.

For a basic walkthrough see [quick start](/docs/QUICK_START.md).

## 📖 Documentation

For a full list of documentation including tutorials, reference and other
material, start with the [/docs directory](/docs/README.md).

Additionally, each plugin has its own README that includes details about how to
configure, use, and sometimes debug or troubleshoot. Look under the
[/plugins directory](/plugins/) for specific plugins.

Here are some commonly used documents:

* [Changelog](/CHANGELOG.md)
* [Configuration](/docs/CONFIGURATION.md)
* [FAQ](/docs/FAQ.md)
* [Releases](https://github.com/influxdata/telegraf/releases)
* [Security](/SECURITY.md)

## ❤️ Contribute

[![Contribute](https://img.shields.io/badge/contribute-to_telegraf-blue.svg?logo=influxdb)](https://github.com/influxdata/telegraf/blob/master/CONTRIBUTING.md)
```toml
[[inputs.http_listener_v2]]
service_address = ":8080"
paths = ["/api/put"]
methods = ["POST", "PUT"]
data_format = "opentsdb"
```

We love our community of over 1,200 contributors! Many of the plugins included
in Telegraf were originally contributed by community members. Check out
our [contributing guide](CONTRIBUTING.md) if you are interested in helping out.
Also, join us on our [Community Slack](https://influxdata.com/slack) or
[Community Forums](https://community.influxdata.com/) if you have questions or
comments for our engineering teams.
**OpenTSDB-Telnet**

If you are completely new to Telegraf and InfluxDB, you can also enroll for free at
[InfluxDB university](https://www.influxdata.com/university/) to take courses to
learn more.
通过使用 Input 插件 socket_listener，并配置 `data_format` 为 opentsdbtelnet，将能够解析 OpenTSDB-Telnet 格式的写入请求。

```toml
[[inputs.socket_listener]]
service_address = "tcp://:8081"
data_format = "opentsdbtelnet"
```

## ℹ️ Support
### Output

[![Slack](https://img.shields.io/badge/slack-join_chat-blue.svg?logo=slack)](https://www.influxdata.com/slack) [![Forums](https://img.shields.io/badge/discourse-join_forums-blue.svg?logo=discourse)](https://community.influxdata.com/)
增加 Output 插件 CnosDBG，用于将指标输出到 CnosDB。

Please use the [Community Slack](https://influxdata.com/slack) or
[Community Forums](https://community.influxdata.com/) if you have questions or
comments for our engineering teams. GitHub issues are limited to actual issues
and feature requests only.

## 📜 License

[![MIT](https://img.shields.io/badge/license-MIT-blue)](https://github.com/influxdata/telegraf/blob/master/LICENSE)

[OPC UA]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/opcua
[Modbus]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/modbus
[File]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/file
[Tail]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/tail
[Directory Monitor]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/directory_monitor
[AMQP]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/amqp_consumer
[Kafka]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/kafka_consumer
[MQTT]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/mqtt_consumer
[OpenTelemetry]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/opentelemetry
[Prometheus]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/prometheus
[Cisco TelemetryMDT]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/cisco_telemetry_mdt
[gNMI]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/gnmi
[CPU]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/cpu
[Memory]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/mem
[Disk]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/disk
[Network]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/net
[SMART]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/smartctl
[Docker]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/docker
[Nvidia SMI]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/nvidia_smi
[Exec]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/exec
[HTTP]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/http
[HTTP Listener]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/http_listener_v2
[SNMP]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/snmp
[SQL]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/sql
[Event Log]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/win_eventlog
[Management Instrumentation]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/win_wmi
[Performance Counters]: https://github.com/influxdata/telegraf/tree/master/plugins/inputs/win_perf_counters
```toml
[[outputs.cnosdb]]
url = "localhost:31006"
user = "user"
password = "pass"
database = "telegraf"
```

**配置介绍**

| 参数       | 说明               |
|----------|------------------|
| url      | CnosDB GRpc 服务地址 |
| user     | 用户名              |
| password | 密码               |
| database | CnosDB 数据库       |

### Input

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

以上配置与在 [Output](#output) 章节中的配置相比，增加了 `high_priority_io = true` 配置项。