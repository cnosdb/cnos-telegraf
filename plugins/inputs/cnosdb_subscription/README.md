# CnosDB Subscription Input Plugin

## Build

To compile this plugin it requires protoc-gen-go and protoc-gen-go-grpc

```shell
# install protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# install protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Usages

To listen on port 8803:

```toml
[[inputs.cnosdb_subscription]]
service_address = ":8803"
```
