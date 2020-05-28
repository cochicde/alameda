// This file has messages related to create alameda scaler

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v4.0.0
// source: alameda_api/v1alpha1/datahub/applications/applications.proto

package applications

import (
	common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//*
// Represents the data of alameda scaler which is to be created.
type WriteApplication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Measurement string            `protobuf:"bytes,1,opt,name=measurement,proto3" json:"measurement,omitempty"`
	WriteData   *common.WriteData `protobuf:"bytes,2,opt,name=write_data,json=writeData,proto3" json:"write_data,omitempty"`
}

func (x *WriteApplication) Reset() {
	*x = WriteApplication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_applications_applications_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteApplication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteApplication) ProtoMessage() {}

func (x *WriteApplication) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_applications_applications_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteApplication.ProtoReflect.Descriptor instead.
func (*WriteApplication) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_applications_applications_proto_rawDescGZIP(), []int{0}
}

func (x *WriteApplication) GetMeasurement() string {
	if x != nil {
		return x.Measurement
	}
	return ""
}

func (x *WriteApplication) GetWriteData() *common.WriteData {
	if x != nil {
		return x.WriteData
	}
	return nil
}

//*
// Represents the condition of reading alameda scalers.
type ReadApplication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Measurement    string              `protobuf:"bytes,1,opt,name=measurement,proto3" json:"measurement,omitempty"`
	WhereCondition []*common.Condition `protobuf:"bytes,2,rep,name=where_condition,json=whereCondition,proto3" json:"where_condition,omitempty"`
}

func (x *ReadApplication) Reset() {
	*x = ReadApplication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_applications_applications_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadApplication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadApplication) ProtoMessage() {}

func (x *ReadApplication) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_applications_applications_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadApplication.ProtoReflect.Descriptor instead.
func (*ReadApplication) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_applications_applications_proto_rawDescGZIP(), []int{1}
}

func (x *ReadApplication) GetMeasurement() string {
	if x != nil {
		return x.Measurement
	}
	return ""
}

func (x *ReadApplication) GetWhereCondition() []*common.Condition {
	if x != nil {
		return x.WhereCondition
	}
	return nil
}

//*
// Represents the condition of deleting alameda scalers.
type DeleteApplication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Measurement    string              `protobuf:"bytes,1,opt,name=measurement,proto3" json:"measurement,omitempty"`
	WhereCondition []*common.Condition `protobuf:"bytes,2,rep,name=where_condition,json=whereCondition,proto3" json:"where_condition,omitempty"`
}

func (x *DeleteApplication) Reset() {
	*x = DeleteApplication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_applications_applications_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteApplication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteApplication) ProtoMessage() {}

func (x *DeleteApplication) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_applications_applications_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteApplication.ProtoReflect.Descriptor instead.
func (*DeleteApplication) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_applications_applications_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteApplication) GetMeasurement() string {
	if x != nil {
		return x.Measurement
	}
	return ""
}

func (x *DeleteApplication) GetWhereCondition() []*common.Condition {
	if x != nil {
		return x.WhereCondition
	}
	return nil
}

var File_alameda_api_v1alpha1_datahub_applications_applications_proto protoreflect.FileDescriptor

var file_alameda_api_v1alpha1_datahub_applications_applications_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x32,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61,
	0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x1a, 0x31, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x68, 0x75, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x61, 0x77, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x10, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a,
	0x0b, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x56, 0x0a, 0x0a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73,
	0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x95, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x6d,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x60, 0x0a,
	0x0f, 0x77, 0x68, 0x65, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0e, 0x77, 0x68, 0x65, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x97, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x60, 0x0a, 0x0f, 0x77, 0x68, 0x65, 0x72, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e,
	0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x77, 0x68, 0x65, 0x72, 0x65,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x73, 0x2d, 0x61, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64,
	0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alameda_api_v1alpha1_datahub_applications_applications_proto_rawDescOnce sync.Once
	file_alameda_api_v1alpha1_datahub_applications_applications_proto_rawDescData = file_alameda_api_v1alpha1_datahub_applications_applications_proto_rawDesc
)

func file_alameda_api_v1alpha1_datahub_applications_applications_proto_rawDescGZIP() []byte {
	file_alameda_api_v1alpha1_datahub_applications_applications_proto_rawDescOnce.Do(func() {
		file_alameda_api_v1alpha1_datahub_applications_applications_proto_rawDescData = protoimpl.X.CompressGZIP(file_alameda_api_v1alpha1_datahub_applications_applications_proto_rawDescData)
	})
	return file_alameda_api_v1alpha1_datahub_applications_applications_proto_rawDescData
}

var file_alameda_api_v1alpha1_datahub_applications_applications_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_alameda_api_v1alpha1_datahub_applications_applications_proto_goTypes = []interface{}{
	(*WriteApplication)(nil),  // 0: containersai.alameda.v1alpha1.datahub.applications.WriteApplication
	(*ReadApplication)(nil),   // 1: containersai.alameda.v1alpha1.datahub.applications.ReadApplication
	(*DeleteApplication)(nil), // 2: containersai.alameda.v1alpha1.datahub.applications.DeleteApplication
	(*common.WriteData)(nil),  // 3: containersai.alameda.v1alpha1.datahub.common.WriteData
	(*common.Condition)(nil),  // 4: containersai.alameda.v1alpha1.datahub.common.Condition
}
var file_alameda_api_v1alpha1_datahub_applications_applications_proto_depIdxs = []int32{
	3, // 0: containersai.alameda.v1alpha1.datahub.applications.WriteApplication.write_data:type_name -> containersai.alameda.v1alpha1.datahub.common.WriteData
	4, // 1: containersai.alameda.v1alpha1.datahub.applications.ReadApplication.where_condition:type_name -> containersai.alameda.v1alpha1.datahub.common.Condition
	4, // 2: containersai.alameda.v1alpha1.datahub.applications.DeleteApplication.where_condition:type_name -> containersai.alameda.v1alpha1.datahub.common.Condition
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_alameda_api_v1alpha1_datahub_applications_applications_proto_init() }
func file_alameda_api_v1alpha1_datahub_applications_applications_proto_init() {
	if File_alameda_api_v1alpha1_datahub_applications_applications_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_alameda_api_v1alpha1_datahub_applications_applications_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteApplication); i {
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
		file_alameda_api_v1alpha1_datahub_applications_applications_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadApplication); i {
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
		file_alameda_api_v1alpha1_datahub_applications_applications_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteApplication); i {
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
			RawDescriptor: file_alameda_api_v1alpha1_datahub_applications_applications_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_alameda_api_v1alpha1_datahub_applications_applications_proto_goTypes,
		DependencyIndexes: file_alameda_api_v1alpha1_datahub_applications_applications_proto_depIdxs,
		MessageInfos:      file_alameda_api_v1alpha1_datahub_applications_applications_proto_msgTypes,
	}.Build()
	File_alameda_api_v1alpha1_datahub_applications_applications_proto = out.File
	file_alameda_api_v1alpha1_datahub_applications_applications_proto_rawDesc = nil
	file_alameda_api_v1alpha1_datahub_applications_applications_proto_goTypes = nil
	file_alameda_api_v1alpha1_datahub_applications_applications_proto_depIdxs = nil
}