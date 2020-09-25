// This file has messages related to read & write data

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: datahub/data/services.proto

package data

import (
	proto "github.com/golang/protobuf/proto"
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	schemas "prophetstor.com/api/datahub/schemas"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//*
// Represents a request for writing data to datahub.
type WriteDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaMeta *schemas.SchemaMeta `protobuf:"bytes,1,opt,name=schema_meta,json=schemaMeta,proto3" json:"schema_meta,omitempty"`
	WriteData  []*WriteData        `protobuf:"bytes,2,rep,name=write_data,json=writeData,proto3" json:"write_data,omitempty"`
}

func (x *WriteDataRequest) Reset() {
	*x = WriteDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_data_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteDataRequest) ProtoMessage() {}

func (x *WriteDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_data_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteDataRequest.ProtoReflect.Descriptor instead.
func (*WriteDataRequest) Descriptor() ([]byte, []int) {
	return file_datahub_data_services_proto_rawDescGZIP(), []int{0}
}

func (x *WriteDataRequest) GetSchemaMeta() *schemas.SchemaMeta {
	if x != nil {
		return x.SchemaMeta
	}
	return nil
}

func (x *WriteDataRequest) GetWriteData() []*WriteData {
	if x != nil {
		return x.WriteData
	}
	return nil
}

//*
// Represents a request for reading data from datahub.
type ReadDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaMeta *schemas.SchemaMeta `protobuf:"bytes,1,opt,name=schema_meta,json=schemaMeta,proto3" json:"schema_meta,omitempty"`
	ReadData   []*ReadData         `protobuf:"bytes,2,rep,name=read_data,json=readData,proto3" json:"read_data,omitempty"`
}

func (x *ReadDataRequest) Reset() {
	*x = ReadDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_data_services_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadDataRequest) ProtoMessage() {}

func (x *ReadDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_data_services_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadDataRequest.ProtoReflect.Descriptor instead.
func (*ReadDataRequest) Descriptor() ([]byte, []int) {
	return file_datahub_data_services_proto_rawDescGZIP(), []int{1}
}

func (x *ReadDataRequest) GetSchemaMeta() *schemas.SchemaMeta {
	if x != nil {
		return x.SchemaMeta
	}
	return nil
}

func (x *ReadDataRequest) GetReadData() []*ReadData {
	if x != nil {
		return x.ReadData
	}
	return nil
}

//*
// Represents a response for a reading data request.
type ReadDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Data   *Data          `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ReadDataResponse) Reset() {
	*x = ReadDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_data_services_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadDataResponse) ProtoMessage() {}

func (x *ReadDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_data_services_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadDataResponse.ProtoReflect.Descriptor instead.
func (*ReadDataResponse) Descriptor() ([]byte, []int) {
	return file_datahub_data_services_proto_rawDescGZIP(), []int{2}
}

func (x *ReadDataResponse) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ReadDataResponse) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

//*
// Represents a request for deleting data in datahub.
type DeleteDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaMeta *schemas.SchemaMeta `protobuf:"bytes,1,opt,name=schema_meta,json=schemaMeta,proto3" json:"schema_meta,omitempty"`
	DeleteData []*DeleteData       `protobuf:"bytes,2,rep,name=delete_data,json=deleteData,proto3" json:"delete_data,omitempty"`
}

func (x *DeleteDataRequest) Reset() {
	*x = DeleteDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_data_services_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDataRequest) ProtoMessage() {}

func (x *DeleteDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_data_services_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDataRequest.ProtoReflect.Descriptor instead.
func (*DeleteDataRequest) Descriptor() ([]byte, []int) {
	return file_datahub_data_services_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteDataRequest) GetSchemaMeta() *schemas.SchemaMeta {
	if x != nil {
		return x.SchemaMeta
	}
	return nil
}

func (x *DeleteDataRequest) GetDeleteData() []*DeleteData {
	if x != nil {
		return x.DeleteData
	}
	return nil
}

//*
// Represents a request for writing data(none time-series) to datahub.
type WriteMetaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaMeta *schemas.SchemaMeta `protobuf:"bytes,1,opt,name=schema_meta,json=schemaMeta,proto3" json:"schema_meta,omitempty"`
	WriteMeta  []*WriteMeta        `protobuf:"bytes,2,rep,name=write_meta,json=writeMeta,proto3" json:"write_meta,omitempty"`
}

func (x *WriteMetaRequest) Reset() {
	*x = WriteMetaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_data_services_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteMetaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteMetaRequest) ProtoMessage() {}

func (x *WriteMetaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_data_services_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteMetaRequest.ProtoReflect.Descriptor instead.
func (*WriteMetaRequest) Descriptor() ([]byte, []int) {
	return file_datahub_data_services_proto_rawDescGZIP(), []int{4}
}

func (x *WriteMetaRequest) GetSchemaMeta() *schemas.SchemaMeta {
	if x != nil {
		return x.SchemaMeta
	}
	return nil
}

func (x *WriteMetaRequest) GetWriteMeta() []*WriteMeta {
	if x != nil {
		return x.WriteMeta
	}
	return nil
}

var File_datahub_data_services_proto protoreflect.FileDescriptor

var file_datahub_data_services_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x70,
	0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68,
	0x75, 0x62, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x17, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x64, 0x61, 0x74, 0x61,
	0x68, 0x75, 0x62, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa0, 0x01, 0x0a, 0x10, 0x57, 0x72, 0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f,
	0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f,
	0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x12,
	0x42, 0x0a, 0x0a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f,
	0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x77, 0x72, 0x69, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x9c, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70,
	0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68,
	0x75, 0x62, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4d, 0x65, 0x74,
	0x61, 0x12, 0x3f, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x72, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa4, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x0b,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x45, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72,
	0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75,
	0x62, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0xa0, 0x01,
	0x0a, 0x10, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x48, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65,
	0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x42, 0x0a, 0x0a,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x09, 0x77, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x42, 0x22, 0x5a, 0x20, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datahub_data_services_proto_rawDescOnce sync.Once
	file_datahub_data_services_proto_rawDescData = file_datahub_data_services_proto_rawDesc
)

func file_datahub_data_services_proto_rawDescGZIP() []byte {
	file_datahub_data_services_proto_rawDescOnce.Do(func() {
		file_datahub_data_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_datahub_data_services_proto_rawDescData)
	})
	return file_datahub_data_services_proto_rawDescData
}

var file_datahub_data_services_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_datahub_data_services_proto_goTypes = []interface{}{
	(*WriteDataRequest)(nil),   // 0: prophetstor.datahub.data.WriteDataRequest
	(*ReadDataRequest)(nil),    // 1: prophetstor.datahub.data.ReadDataRequest
	(*ReadDataResponse)(nil),   // 2: prophetstor.datahub.data.ReadDataResponse
	(*DeleteDataRequest)(nil),  // 3: prophetstor.datahub.data.DeleteDataRequest
	(*WriteMetaRequest)(nil),   // 4: prophetstor.datahub.data.WriteMetaRequest
	(*schemas.SchemaMeta)(nil), // 5: prophetstor.datahub.schemas.SchemaMeta
	(*WriteData)(nil),          // 6: prophetstor.datahub.data.WriteData
	(*ReadData)(nil),           // 7: prophetstor.datahub.data.ReadData
	(*status.Status)(nil),      // 8: google.rpc.Status
	(*Data)(nil),               // 9: prophetstor.datahub.data.Data
	(*DeleteData)(nil),         // 10: prophetstor.datahub.data.DeleteData
	(*WriteMeta)(nil),          // 11: prophetstor.datahub.data.WriteMeta
}
var file_datahub_data_services_proto_depIdxs = []int32{
	5,  // 0: prophetstor.datahub.data.WriteDataRequest.schema_meta:type_name -> prophetstor.datahub.schemas.SchemaMeta
	6,  // 1: prophetstor.datahub.data.WriteDataRequest.write_data:type_name -> prophetstor.datahub.data.WriteData
	5,  // 2: prophetstor.datahub.data.ReadDataRequest.schema_meta:type_name -> prophetstor.datahub.schemas.SchemaMeta
	7,  // 3: prophetstor.datahub.data.ReadDataRequest.read_data:type_name -> prophetstor.datahub.data.ReadData
	8,  // 4: prophetstor.datahub.data.ReadDataResponse.status:type_name -> google.rpc.Status
	9,  // 5: prophetstor.datahub.data.ReadDataResponse.data:type_name -> prophetstor.datahub.data.Data
	5,  // 6: prophetstor.datahub.data.DeleteDataRequest.schema_meta:type_name -> prophetstor.datahub.schemas.SchemaMeta
	10, // 7: prophetstor.datahub.data.DeleteDataRequest.delete_data:type_name -> prophetstor.datahub.data.DeleteData
	5,  // 8: prophetstor.datahub.data.WriteMetaRequest.schema_meta:type_name -> prophetstor.datahub.schemas.SchemaMeta
	11, // 9: prophetstor.datahub.data.WriteMetaRequest.write_meta:type_name -> prophetstor.datahub.data.WriteMeta
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_datahub_data_services_proto_init() }
func file_datahub_data_services_proto_init() {
	if File_datahub_data_services_proto != nil {
		return
	}
	file_datahub_data_data_proto_init()
	file_datahub_data_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_datahub_data_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteDataRequest); i {
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
		file_datahub_data_services_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadDataRequest); i {
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
		file_datahub_data_services_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadDataResponse); i {
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
		file_datahub_data_services_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDataRequest); i {
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
		file_datahub_data_services_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteMetaRequest); i {
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
			RawDescriptor: file_datahub_data_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datahub_data_services_proto_goTypes,
		DependencyIndexes: file_datahub_data_services_proto_depIdxs,
		MessageInfos:      file_datahub_data_services_proto_msgTypes,
	}.Build()
	File_datahub_data_services_proto = out.File
	file_datahub_data_services_proto_rawDesc = nil
	file_datahub_data_services_proto_goTypes = nil
	file_datahub_data_services_proto_depIdxs = nil
}
