package cnosdb_subscription

//go:generate flatc -o ./cnosdb/generated/models/v3 --go --go-namespace v3 --gen-onefile ./protocol/models.v3.fbs
//go:generate flatc -o ./cnosdb/generated/models/v4 --go --go-namespace v4 --gen-onefile ./protocol/models.v4.fbs
//go:generate protoc --go_out=./cnosdb/generated/kv_service --go-grpc_out=./cnosdb/generated/kv_service --proto_path ./protocol ./protocol/kv_service.proto
