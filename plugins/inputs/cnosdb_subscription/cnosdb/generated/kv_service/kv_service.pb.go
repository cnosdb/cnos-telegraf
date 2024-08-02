// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: kv_service.proto

package kv_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// CnosDB subscription v3 message.
type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant   string  `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	User     *string `protobuf:"bytes,2,opt,name=user,proto3,oneof" json:"user,omitempty"`
	Password *string `protobuf:"bytes,3,opt,name=password,proto3,oneof" json:"password,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_kv_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_kv_service_proto_rawDescGZIP(), []int{0}
}

func (x *Meta) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *Meta) GetUser() string {
	if x != nil && x.User != nil {
		return *x.User
	}
	return ""
}

func (x *Meta) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

// CnosDB subscription v3 message.
type WritePointsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version uint64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Meta    *Meta  `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
	Points  []byte `protobuf:"bytes,3,opt,name=points,proto3" json:"points,omitempty"`
}

func (x *WritePointsRequest) Reset() {
	*x = WritePointsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WritePointsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WritePointsRequest) ProtoMessage() {}

func (x *WritePointsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kv_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WritePointsRequest.ProtoReflect.Descriptor instead.
func (*WritePointsRequest) Descriptor() ([]byte, []int) {
	return file_kv_service_proto_rawDescGZIP(), []int{1}
}

func (x *WritePointsRequest) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *WritePointsRequest) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *WritePointsRequest) GetPoints() []byte {
	if x != nil {
		return x.Points
	}
	return nil
}

// CnosDB subscription v3 message.
type WritePointsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PointsNumber uint64 `protobuf:"varint,1,opt,name=points_number,json=pointsNumber,proto3" json:"points_number,omitempty"`
}

func (x *WritePointsResponse) Reset() {
	*x = WritePointsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WritePointsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WritePointsResponse) ProtoMessage() {}

func (x *WritePointsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kv_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WritePointsResponse.ProtoReflect.Descriptor instead.
func (*WritePointsResponse) Descriptor() ([]byte, []int) {
	return file_kv_service_proto_rawDescGZIP(), []int{2}
}

func (x *WritePointsResponse) GetPointsNumber() uint64 {
	if x != nil {
		return x.PointsNumber
	}
	return 0
}

// CnosDB subscription v4 message.
type SubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Marked as deprecated in kv_service.proto.
	Tenant string `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	// Deprecated: Marked as deprecated in kv_service.proto.
	Db          string `protobuf:"bytes,2,opt,name=db,proto3" json:"db,omitempty"`
	TableSchema []byte `protobuf:"bytes,3,opt,name=table_schema,json=tableSchema,proto3" json:"table_schema,omitempty"`
	RecordData  []byte `protobuf:"bytes,4,opt,name=record_data,json=recordData,proto3" json:"record_data,omitempty"`
	Precision   uint32 `protobuf:"varint,5,opt,name=precision,proto3" json:"precision,omitempty"`
}

func (x *SubscriptionRequest) Reset() {
	*x = SubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionRequest) ProtoMessage() {}

func (x *SubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kv_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionRequest.ProtoReflect.Descriptor instead.
func (*SubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_kv_service_proto_rawDescGZIP(), []int{3}
}

// Deprecated: Marked as deprecated in kv_service.proto.
func (x *SubscriptionRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

// Deprecated: Marked as deprecated in kv_service.proto.
func (x *SubscriptionRequest) GetDb() string {
	if x != nil {
		return x.Db
	}
	return ""
}

func (x *SubscriptionRequest) GetTableSchema() []byte {
	if x != nil {
		return x.TableSchema
	}
	return nil
}

func (x *SubscriptionRequest) GetRecordData() []byte {
	if x != nil {
		return x.RecordData
	}
	return nil
}

func (x *SubscriptionRequest) GetPrecision() uint32 {
	if x != nil {
		return x.Precision
	}
	return 0
}

// CnosDB subscription v4 message.
type SubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscriptionResponse) Reset() {
	*x = SubscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionResponse) ProtoMessage() {}

func (x *SubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kv_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionResponse.ProtoReflect.Descriptor instead.
func (*SubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_kv_service_proto_rawDescGZIP(), []int{4}
}

var File_kv_service_proto protoreflect.FileDescriptor

var file_kv_service_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6b, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x6b, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x6e,
	0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x17,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x6c,
	0x0a, 0x12, 0x57, 0x72, 0x69, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x24,
	0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6b,
	0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x3a, 0x0a, 0x13,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xa7, 0x01, 0x0a, 0x13, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x02, 0x18, 0x01, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x02,
	0x64, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x02, 0x64, 0x62,
	0x12, 0x21, 0x0a, 0x0c, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc1, 0x01, 0x0a, 0x0b, 0x54,
	0x53, 0x4b, 0x56, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0b, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x1e, 0x2e, 0x6b, 0x76, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6b, 0x76, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x12, 0x5c, 0x0a, 0x11, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x6b, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6b, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0e,
	0x5a, 0x0c, 0x2e, 0x3b, 0x6b, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kv_service_proto_rawDescOnce sync.Once
	file_kv_service_proto_rawDescData = file_kv_service_proto_rawDesc
)

func file_kv_service_proto_rawDescGZIP() []byte {
	file_kv_service_proto_rawDescOnce.Do(func() {
		file_kv_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_kv_service_proto_rawDescData)
	})
	return file_kv_service_proto_rawDescData
}

var file_kv_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_kv_service_proto_goTypes = []any{
	(*Meta)(nil),                 // 0: kv_service.Meta
	(*WritePointsRequest)(nil),   // 1: kv_service.WritePointsRequest
	(*WritePointsResponse)(nil),  // 2: kv_service.WritePointsResponse
	(*SubscriptionRequest)(nil),  // 3: kv_service.SubscriptionRequest
	(*SubscriptionResponse)(nil), // 4: kv_service.SubscriptionResponse
}
var file_kv_service_proto_depIdxs = []int32{
	0, // 0: kv_service.WritePointsRequest.meta:type_name -> kv_service.Meta
	1, // 1: kv_service.TSKVService.WritePoints:input_type -> kv_service.WritePointsRequest
	3, // 2: kv_service.TSKVService.WriteSubscription:input_type -> kv_service.SubscriptionRequest
	2, // 3: kv_service.TSKVService.WritePoints:output_type -> kv_service.WritePointsResponse
	4, // 4: kv_service.TSKVService.WriteSubscription:output_type -> kv_service.SubscriptionResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kv_service_proto_init() }
func file_kv_service_proto_init() {
	if File_kv_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kv_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Meta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kv_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*WritePointsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kv_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*WritePointsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kv_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SubscriptionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kv_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SubscriptionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_kv_service_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kv_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kv_service_proto_goTypes,
		DependencyIndexes: file_kv_service_proto_depIdxs,
		MessageInfos:      file_kv_service_proto_msgTypes,
	}.Build()
	File_kv_service_proto = out.File
	file_kv_service_proto_rawDesc = nil
	file_kv_service_proto_goTypes = nil
	file_kv_service_proto_depIdxs = nil
}