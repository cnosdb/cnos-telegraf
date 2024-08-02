# CnosDB Subscription Input Plugin

## Usages

To listen on port 8803:

```toml
[[inputs.cnosdb_subscription]]
service_address = ":8803"
```

If you want to get response from output plugins immediately, add config `high_priority_io = true`:

```toml
[[inputs.cnosdb_subscription]]
service_address = ":8803"
high_priority_io = true
```

## Develop

If you want to edit protocol files (in directory `protocol/`),
you need **protoc-gen-go** and **protoc-gen-go-grpc** to re-compile them.
See [generate.go](./generate.go).

```shell
# install protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# install protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
