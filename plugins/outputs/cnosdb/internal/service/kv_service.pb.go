// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: internal/service/kv_service.proto

package service

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

type FieldType int32

const (
	FieldType_Float    FieldType = 0
	FieldType_Integer  FieldType = 1
	FieldType_Unsigned FieldType = 2
	FieldType_Boolean  FieldType = 3
	FieldType_String   FieldType = 5
)

// Enum value maps for FieldType.
var (
	FieldType_name = map[int32]string{
		0: "Float",
		1: "Integer",
		2: "Unsigned",
		3: "Boolean",
		5: "String",
	}
	FieldType_value = map[string]int32{
		"Float":    0,
		"Integer":  1,
		"Unsigned": 2,
		"Boolean":  3,
		"String":   5,
	}
)

func (x FieldType) Enum() *FieldType {
	p := new(FieldType)
	*p = x
	return p
}

func (x FieldType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FieldType) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_service_kv_service_proto_enumTypes[0].Descriptor()
}

func (FieldType) Type() protoreflect.EnumType {
	return &file_internal_service_kv_service_proto_enumTypes[0]
}

func (x FieldType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FieldType.Descriptor instead.
func (FieldType) EnumDescriptor() ([]byte, []int) {
	return file_internal_service_kv_service_proto_rawDescGZIP(), []int{0}
}

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version uint64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Body    []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_service_kv_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_service_kv_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_internal_service_kv_service_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *PingRequest) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version uint64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Body    []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_service_kv_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_service_kv_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_internal_service_kv_service_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *PingResponse) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`     // tag key utf-8 bytes
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"` // tag value utf-8 bytes
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_service_kv_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_internal_service_kv_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_internal_service_kv_service_proto_rawDescGZIP(), []int{2}
}

func (x *Tag) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Tag) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type FieldInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldType FieldType `protobuf:"varint,1,opt,name=field_type,json=fieldType,proto3,enum=kv_service.FieldType" json:"field_type,omitempty"`
	Name      []byte    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // field name utf-8 bytes
	Id        uint64    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FieldInfo) Reset() {
	*x = FieldInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_service_kv_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldInfo) ProtoMessage() {}

func (x *FieldInfo) ProtoReflect() protoreflect.Message {
	mi := &file_internal_service_kv_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldInfo.ProtoReflect.Descriptor instead.
func (*FieldInfo) Descriptor() ([]byte, []int) {
	return file_internal_service_kv_service_proto_rawDescGZIP(), []int{3}
}

func (x *FieldInfo) GetFieldType() FieldType {
	if x != nil {
		return x.FieldType
	}
	return FieldType_Float
}

func (x *FieldInfo) GetName() []byte {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *FieldInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AddSeriesRpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version  uint64       `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	SeriesId uint64       `protobuf:"varint,2,opt,name=series_id,json=seriesId,proto3" json:"series_id,omitempty"`
	Tags     []*Tag       `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	Fields   []*FieldInfo `protobuf:"bytes,4,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *AddSeriesRpcRequest) Reset() {
	*x = AddSeriesRpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_service_kv_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSeriesRpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSeriesRpcRequest) ProtoMessage() {}

func (x *AddSeriesRpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_service_kv_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSeriesRpcRequest.ProtoReflect.Descriptor instead.
func (*AddSeriesRpcRequest) Descriptor() ([]byte, []int) {
	return file_internal_service_kv_service_proto_rawDescGZIP(), []int{4}
}

func (x *AddSeriesRpcRequest) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *AddSeriesRpcRequest) GetSeriesId() uint64 {
	if x != nil {
		return x.SeriesId
	}
	return 0
}

func (x *AddSeriesRpcRequest) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *AddSeriesRpcRequest) GetFields() []*FieldInfo {
	if x != nil {
		return x.Fields
	}
	return nil
}

type AddSeriesRpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version  uint64       `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	SeriesId uint64       `protobuf:"varint,2,opt,name=series_id,json=seriesId,proto3" json:"series_id,omitempty"`
	Tags     []*Tag       `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	Fields   []*FieldInfo `protobuf:"bytes,4,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *AddSeriesRpcResponse) Reset() {
	*x = AddSeriesRpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_service_kv_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSeriesRpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSeriesRpcResponse) ProtoMessage() {}

func (x *AddSeriesRpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_service_kv_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSeriesRpcResponse.ProtoReflect.Descriptor instead.
func (*AddSeriesRpcResponse) Descriptor() ([]byte, []int) {
	return file_internal_service_kv_service_proto_rawDescGZIP(), []int{5}
}

func (x *AddSeriesRpcResponse) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *AddSeriesRpcResponse) GetSeriesId() uint64 {
	if x != nil {
		return x.SeriesId
	}
	return 0
}

func (x *AddSeriesRpcResponse) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *AddSeriesRpcResponse) GetFields() []*FieldInfo {
	if x != nil {
		return x.Fields
	}
	return nil
}

type GetSeriesInfoRpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtocolVersion uint64 `protobuf:"varint,1,opt,name=protocol_version,json=protocolVersion,proto3" json:"protocol_version,omitempty"`
	SeriesId        uint64 `protobuf:"varint,2,opt,name=series_id,json=seriesId,proto3" json:"series_id,omitempty"`
}

func (x *GetSeriesInfoRpcRequest) Reset() {
	*x = GetSeriesInfoRpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_service_kv_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSeriesInfoRpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSeriesInfoRpcRequest) ProtoMessage() {}

func (x *GetSeriesInfoRpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_service_kv_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSeriesInfoRpcRequest.ProtoReflect.Descriptor instead.
func (*GetSeriesInfoRpcRequest) Descriptor() ([]byte, []int) {
	return file_internal_service_kv_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetSeriesInfoRpcRequest) GetProtocolVersion() uint64 {
	if x != nil {
		return x.ProtocolVersion
	}
	return 0
}

func (x *GetSeriesInfoRpcRequest) GetSeriesId() uint64 {
	if x != nil {
		return x.SeriesId
	}
	return 0
}

type GetSeriesInfoRpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version  uint64       `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	SeriesId uint64       `protobuf:"varint,2,opt,name=series_id,json=seriesId,proto3" json:"series_id,omitempty"`
	Tags     []*Tag       `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	Fields   []*FieldInfo `protobuf:"bytes,4,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *GetSeriesInfoRpcResponse) Reset() {
	*x = GetSeriesInfoRpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_service_kv_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSeriesInfoRpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSeriesInfoRpcResponse) ProtoMessage() {}

func (x *GetSeriesInfoRpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_service_kv_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSeriesInfoRpcResponse.ProtoReflect.Descriptor instead.
func (*GetSeriesInfoRpcResponse) Descriptor() ([]byte, []int) {
	return file_internal_service_kv_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetSeriesInfoRpcResponse) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *GetSeriesInfoRpcResponse) GetSeriesId() uint64 {
	if x != nil {
		return x.SeriesId
	}
	return 0
}

func (x *GetSeriesInfoRpcResponse) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *GetSeriesInfoRpcResponse) GetFields() []*FieldInfo {
	if x != nil {
		return x.Fields
	}
	return nil
}

type WriteRowsRpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version uint64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Rows    []byte `protobuf:"bytes,2,opt,name=rows,proto3" json:"rows,omitempty"` // flatbuffers bytes ( models::Rows )
}

func (x *WriteRowsRpcRequest) Reset() {
	*x = WriteRowsRpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_service_kv_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRowsRpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRowsRpcRequest) ProtoMessage() {}

func (x *WriteRowsRpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_service_kv_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRowsRpcRequest.ProtoReflect.Descriptor instead.
func (*WriteRowsRpcRequest) Descriptor() ([]byte, []int) {
	return file_internal_service_kv_service_proto_rawDescGZIP(), []int{8}
}

func (x *WriteRowsRpcRequest) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *WriteRowsRpcRequest) GetRows() []byte {
	if x != nil {
		return x.Rows
	}
	return nil
}

type WriteRowsRpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version uint64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Rows    []byte `protobuf:"bytes,2,opt,name=rows,proto3" json:"rows,omitempty"` // flatbuffers bytes ( models::Rows )
}

func (x *WriteRowsRpcResponse) Reset() {
	*x = WriteRowsRpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_service_kv_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRowsRpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRowsRpcResponse) ProtoMessage() {}

func (x *WriteRowsRpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_service_kv_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRowsRpcResponse.ProtoReflect.Descriptor instead.
func (*WriteRowsRpcResponse) Descriptor() ([]byte, []int) {
	return file_internal_service_kv_service_proto_rawDescGZIP(), []int{9}
}

func (x *WriteRowsRpcResponse) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *WriteRowsRpcResponse) GetRows() []byte {
	if x != nil {
		return x.Rows
	}
	return nil
}

type WritePointsRpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version uint64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Points  []byte `protobuf:"bytes,2,opt,name=points,proto3" json:"points,omitempty"` // flatbuffers bytes ( models::Points )
}

func (x *WritePointsRpcRequest) Reset() {
	*x = WritePointsRpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_service_kv_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WritePointsRpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WritePointsRpcRequest) ProtoMessage() {}

func (x *WritePointsRpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_service_kv_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WritePointsRpcRequest.ProtoReflect.Descriptor instead.
func (*WritePointsRpcRequest) Descriptor() ([]byte, []int) {
	return file_internal_service_kv_service_proto_rawDescGZIP(), []int{10}
}

func (x *WritePointsRpcRequest) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *WritePointsRpcRequest) GetPoints() []byte {
	if x != nil {
		return x.Points
	}
	return nil
}

type WritePointsRpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version uint64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Points  []byte `protobuf:"bytes,2,opt,name=points,proto3" json:"points,omitempty"` // flatbuffers bytes ( models::Points )
}

func (x *WritePointsRpcResponse) Reset() {
	*x = WritePointsRpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_service_kv_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WritePointsRpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WritePointsRpcResponse) ProtoMessage() {}

func (x *WritePointsRpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_service_kv_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WritePointsRpcResponse.ProtoReflect.Descriptor instead.
func (*WritePointsRpcResponse) Descriptor() ([]byte, []int) {
	return file_internal_service_kv_service_proto_rawDescGZIP(), []int{11}
}

func (x *WritePointsRpcResponse) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *WritePointsRpcResponse) GetPoints() []byte {
	if x != nil {
		return x.Points
	}
	return nil
}

var File_internal_service_kv_service_proto protoreflect.FileDescriptor

var file_internal_service_kv_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x6b, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6b, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22,
	0x3b, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x3c, 0x0a, 0x0c,
	0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x2d, 0x0a, 0x03, 0x54, 0x61,
	0x67, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x65, 0x0a, 0x09, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6b, 0x76, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xa0, 0x01, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x70,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x49, 0x64, 0x12,
	0x23, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x6b, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x12, 0x2d, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6b, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x22, 0xa1, 0x01, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x53, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x6b, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54,
	0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x2d, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6b, 0x76, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x61, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x49, 0x64, 0x22, 0xa5, 0x01, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x70, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6b,
	0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x12, 0x2d, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6b, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x22, 0x43, 0x0a, 0x13, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x6f, 0x77, 0x73, 0x52,
	0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x22, 0x44, 0x0a, 0x14, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x6f, 0x77, 0x73, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x77,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x22, 0x49, 0x0a,
	0x15, 0x57, 0x72, 0x69, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x70, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x4a, 0x0a, 0x16, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x2a, 0x4a, 0x0a, 0x09, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x6e, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x6f, 0x6f, 0x6c, 0x65,
	0x61, 0x6e, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x10, 0x05,
	0x32, 0xaa, 0x03, 0x0a, 0x0b, 0x54, 0x53, 0x4b, 0x56, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x39, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x2e, 0x6b, 0x76, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x6b, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x09, 0x41,
	0x64, 0x64, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x6b, 0x76, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6b, 0x76, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23,
	0x2e, 0x6b, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6b, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x70,
	0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x09, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x52, 0x6f, 0x77, 0x73, 0x12, 0x1f, 0x2e, 0x6b, 0x76, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x6f, 0x77, 0x73, 0x52,
	0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6b, 0x76, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x6f, 0x77, 0x73,
	0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x12, 0x5a, 0x0a, 0x0b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x12, 0x21, 0x2e, 0x6b, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6b, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x70, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0b, 0x5a,
	0x09, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_internal_service_kv_service_proto_rawDescOnce sync.Once
	file_internal_service_kv_service_proto_rawDescData = file_internal_service_kv_service_proto_rawDesc
)

func file_internal_service_kv_service_proto_rawDescGZIP() []byte {
	file_internal_service_kv_service_proto_rawDescOnce.Do(func() {
		file_internal_service_kv_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_service_kv_service_proto_rawDescData)
	})
	return file_internal_service_kv_service_proto_rawDescData
}

var file_internal_service_kv_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_service_kv_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_internal_service_kv_service_proto_goTypes = []interface{}{
	(FieldType)(0),                   // 0: kv_service.FieldType
	(*PingRequest)(nil),              // 1: kv_service.PingRequest
	(*PingResponse)(nil),             // 2: kv_service.PingResponse
	(*Tag)(nil),                      // 3: kv_service.Tag
	(*FieldInfo)(nil),                // 4: kv_service.FieldInfo
	(*AddSeriesRpcRequest)(nil),      // 5: kv_service.AddSeriesRpcRequest
	(*AddSeriesRpcResponse)(nil),     // 6: kv_service.AddSeriesRpcResponse
	(*GetSeriesInfoRpcRequest)(nil),  // 7: kv_service.GetSeriesInfoRpcRequest
	(*GetSeriesInfoRpcResponse)(nil), // 8: kv_service.GetSeriesInfoRpcResponse
	(*WriteRowsRpcRequest)(nil),      // 9: kv_service.WriteRowsRpcRequest
	(*WriteRowsRpcResponse)(nil),     // 10: kv_service.WriteRowsRpcResponse
	(*WritePointsRpcRequest)(nil),    // 11: kv_service.WritePointsRpcRequest
	(*WritePointsRpcResponse)(nil),   // 12: kv_service.WritePointsRpcResponse
}
var file_internal_service_kv_service_proto_depIdxs = []int32{
	0,  // 0: kv_service.FieldInfo.field_type:type_name -> kv_service.FieldType
	3,  // 1: kv_service.AddSeriesRpcRequest.tags:type_name -> kv_service.Tag
	4,  // 2: kv_service.AddSeriesRpcRequest.fields:type_name -> kv_service.FieldInfo
	3,  // 3: kv_service.AddSeriesRpcResponse.tags:type_name -> kv_service.Tag
	4,  // 4: kv_service.AddSeriesRpcResponse.fields:type_name -> kv_service.FieldInfo
	3,  // 5: kv_service.GetSeriesInfoRpcResponse.tags:type_name -> kv_service.Tag
	4,  // 6: kv_service.GetSeriesInfoRpcResponse.fields:type_name -> kv_service.FieldInfo
	1,  // 7: kv_service.TSKVService.Ping:input_type -> kv_service.PingRequest
	5,  // 8: kv_service.TSKVService.AddSeries:input_type -> kv_service.AddSeriesRpcRequest
	7,  // 9: kv_service.TSKVService.GetSeriesInfo:input_type -> kv_service.GetSeriesInfoRpcRequest
	9,  // 10: kv_service.TSKVService.WriteRows:input_type -> kv_service.WriteRowsRpcRequest
	11, // 11: kv_service.TSKVService.WritePoints:input_type -> kv_service.WritePointsRpcRequest
	2,  // 12: kv_service.TSKVService.Ping:output_type -> kv_service.PingResponse
	6,  // 13: kv_service.TSKVService.AddSeries:output_type -> kv_service.AddSeriesRpcResponse
	8,  // 14: kv_service.TSKVService.GetSeriesInfo:output_type -> kv_service.GetSeriesInfoRpcResponse
	10, // 15: kv_service.TSKVService.WriteRows:output_type -> kv_service.WriteRowsRpcResponse
	12, // 16: kv_service.TSKVService.WritePoints:output_type -> kv_service.WritePointsRpcResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_internal_service_kv_service_proto_init() }
func file_internal_service_kv_service_proto_init() {
	if File_internal_service_kv_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_service_kv_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_internal_service_kv_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_internal_service_kv_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_internal_service_kv_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldInfo); i {
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
		file_internal_service_kv_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSeriesRpcRequest); i {
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
		file_internal_service_kv_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSeriesRpcResponse); i {
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
		file_internal_service_kv_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSeriesInfoRpcRequest); i {
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
		file_internal_service_kv_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSeriesInfoRpcResponse); i {
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
		file_internal_service_kv_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRowsRpcRequest); i {
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
		file_internal_service_kv_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRowsRpcResponse); i {
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
		file_internal_service_kv_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WritePointsRpcRequest); i {
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
		file_internal_service_kv_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WritePointsRpcResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_service_kv_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_service_kv_service_proto_goTypes,
		DependencyIndexes: file_internal_service_kv_service_proto_depIdxs,
		EnumInfos:         file_internal_service_kv_service_proto_enumTypes,
		MessageInfos:      file_internal_service_kv_service_proto_msgTypes,
	}.Build()
	File_internal_service_kv_service_proto = out.File
	file_internal_service_kv_service_proto_rawDesc = nil
	file_internal_service_kv_service_proto_goTypes = nil
	file_internal_service_kv_service_proto_depIdxs = nil
}