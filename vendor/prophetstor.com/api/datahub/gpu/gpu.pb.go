// This file has messages related to gpu info, metrics and predictions

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: datahub/gpu/gpu.proto

package gpu

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	common "prophetstor.com/api/datahub/common"
	predictions "prophetstor.com/api/datahub/predictions"
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
// Represents a graphics processing unit.
type Gpu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uuid     string       `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Metadata *GpuMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec     *GpuSpec     `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *Gpu) Reset() {
	*x = Gpu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_gpu_gpu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gpu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gpu) ProtoMessage() {}

func (x *Gpu) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_gpu_gpu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gpu.ProtoReflect.Descriptor instead.
func (*Gpu) Descriptor() ([]byte, []int) {
	return file_datahub_gpu_gpu_proto_rawDescGZIP(), []int{0}
}

func (x *Gpu) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Gpu) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Gpu) GetMetadata() *GpuMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Gpu) GetSpec() *GpuSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

//*
// Represents metric data of a graphics processing unit.
type GpuMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uuid       string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Metadata   *GpuMetadata         `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	MetricData []*common.MetricData `protobuf:"bytes,4,rep,name=metric_data,json=metricData,proto3" json:"metric_data,omitempty"`
}

func (x *GpuMetric) Reset() {
	*x = GpuMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_gpu_gpu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GpuMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GpuMetric) ProtoMessage() {}

func (x *GpuMetric) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_gpu_gpu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GpuMetric.ProtoReflect.Descriptor instead.
func (*GpuMetric) Descriptor() ([]byte, []int) {
	return file_datahub_gpu_gpu_proto_rawDescGZIP(), []int{1}
}

func (x *GpuMetric) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GpuMetric) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GpuMetric) GetMetadata() *GpuMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *GpuMetric) GetMetricData() []*common.MetricData {
	if x != nil {
		return x.MetricData
	}
	return nil
}

//*
// Represents a list of predicted metrics data of a graphics processing unit.
type GpuPrediction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                    string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uuid                    string                    `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Metadata                *GpuMetadata              `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	PredictedRawData        []*predictions.MetricData `protobuf:"bytes,4,rep,name=predicted_raw_data,json=predictedRawData,proto3" json:"predicted_raw_data,omitempty"`
	PredictedUpperboundData []*predictions.MetricData `protobuf:"bytes,5,rep,name=predicted_upperbound_data,json=predictedUpperboundData,proto3" json:"predicted_upperbound_data,omitempty"`
	PredictedLowerboundData []*predictions.MetricData `protobuf:"bytes,6,rep,name=predicted_lowerbound_data,json=predictedLowerboundData,proto3" json:"predicted_lowerbound_data,omitempty"`
}

func (x *GpuPrediction) Reset() {
	*x = GpuPrediction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_gpu_gpu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GpuPrediction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GpuPrediction) ProtoMessage() {}

func (x *GpuPrediction) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_gpu_gpu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GpuPrediction.ProtoReflect.Descriptor instead.
func (*GpuPrediction) Descriptor() ([]byte, []int) {
	return file_datahub_gpu_gpu_proto_rawDescGZIP(), []int{2}
}

func (x *GpuPrediction) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GpuPrediction) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GpuPrediction) GetMetadata() *GpuMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *GpuPrediction) GetPredictedRawData() []*predictions.MetricData {
	if x != nil {
		return x.PredictedRawData
	}
	return nil
}

func (x *GpuPrediction) GetPredictedUpperboundData() []*predictions.MetricData {
	if x != nil {
		return x.PredictedUpperboundData
	}
	return nil
}

func (x *GpuPrediction) GetPredictedLowerboundData() []*predictions.MetricData {
	if x != nil {
		return x.PredictedLowerboundData
	}
	return nil
}

var File_datahub_gpu_gpu_proto protoreflect.FileDescriptor

var file_datahub_gpu_gpu_proto_rawDesc = []byte{
	0x0a, 0x15, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x67, 0x70, 0x75, 0x2f, 0x67, 0x70,
	0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x67, 0x70, 0x75,
	0x1a, 0x1c, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x67, 0x70, 0x75, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62,
	0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x03, 0x47, 0x70, 0x75,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x40, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f,
	0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62,
	0x2e, 0x67, 0x70, 0x75, 0x2e, 0x47, 0x70, 0x75, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68,
	0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x67,
	0x70, 0x75, 0x2e, 0x47, 0x70, 0x75, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x22, 0xbe, 0x01, 0x0a, 0x09, 0x47, 0x70, 0x75, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x40, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68,
	0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x67,
	0x70, 0x75, 0x2e, 0x47, 0x70, 0x75, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x47, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74,
	0x61, 0x22, 0xa6, 0x03, 0x0a, 0x0d, 0x47, 0x70, 0x75, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x40, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x68, 0x75, 0x62, 0x2e, 0x67, 0x70, 0x75, 0x2e, 0x47, 0x70, 0x75, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x59, 0x0a,
	0x12, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x61, 0x77, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x70,
	0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e,
	0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x52, 0x10, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x65,
	0x64, 0x52, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x67, 0x0a, 0x19, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x70, 0x70, 0x65, 0x72, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72,
	0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75,
	0x62, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x52, 0x17, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x74, 0x65, 0x64, 0x55, 0x70, 0x70, 0x65, 0x72, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x67, 0x0a, 0x19, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x17, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x65, 0x64, 0x4c, 0x6f, 0x77, 0x65,
	0x72, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x42, 0x21, 0x5a, 0x1f, 0x70, 0x72,
	0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x67, 0x70, 0x75, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datahub_gpu_gpu_proto_rawDescOnce sync.Once
	file_datahub_gpu_gpu_proto_rawDescData = file_datahub_gpu_gpu_proto_rawDesc
)

func file_datahub_gpu_gpu_proto_rawDescGZIP() []byte {
	file_datahub_gpu_gpu_proto_rawDescOnce.Do(func() {
		file_datahub_gpu_gpu_proto_rawDescData = protoimpl.X.CompressGZIP(file_datahub_gpu_gpu_proto_rawDescData)
	})
	return file_datahub_gpu_gpu_proto_rawDescData
}

var file_datahub_gpu_gpu_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_datahub_gpu_gpu_proto_goTypes = []interface{}{
	(*Gpu)(nil),                    // 0: prophetstor.datahub.gpu.Gpu
	(*GpuMetric)(nil),              // 1: prophetstor.datahub.gpu.GpuMetric
	(*GpuPrediction)(nil),          // 2: prophetstor.datahub.gpu.GpuPrediction
	(*GpuMetadata)(nil),            // 3: prophetstor.datahub.gpu.GpuMetadata
	(*GpuSpec)(nil),                // 4: prophetstor.datahub.gpu.GpuSpec
	(*common.MetricData)(nil),      // 5: prophetstor.datahub.common.MetricData
	(*predictions.MetricData)(nil), // 6: prophetstor.datahub.predictions.MetricData
}
var file_datahub_gpu_gpu_proto_depIdxs = []int32{
	3, // 0: prophetstor.datahub.gpu.Gpu.metadata:type_name -> prophetstor.datahub.gpu.GpuMetadata
	4, // 1: prophetstor.datahub.gpu.Gpu.spec:type_name -> prophetstor.datahub.gpu.GpuSpec
	3, // 2: prophetstor.datahub.gpu.GpuMetric.metadata:type_name -> prophetstor.datahub.gpu.GpuMetadata
	5, // 3: prophetstor.datahub.gpu.GpuMetric.metric_data:type_name -> prophetstor.datahub.common.MetricData
	3, // 4: prophetstor.datahub.gpu.GpuPrediction.metadata:type_name -> prophetstor.datahub.gpu.GpuMetadata
	6, // 5: prophetstor.datahub.gpu.GpuPrediction.predicted_raw_data:type_name -> prophetstor.datahub.predictions.MetricData
	6, // 6: prophetstor.datahub.gpu.GpuPrediction.predicted_upperbound_data:type_name -> prophetstor.datahub.predictions.MetricData
	6, // 7: prophetstor.datahub.gpu.GpuPrediction.predicted_lowerbound_data:type_name -> prophetstor.datahub.predictions.MetricData
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_datahub_gpu_gpu_proto_init() }
func file_datahub_gpu_gpu_proto_init() {
	if File_datahub_gpu_gpu_proto != nil {
		return
	}
	file_datahub_gpu_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_datahub_gpu_gpu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gpu); i {
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
		file_datahub_gpu_gpu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GpuMetric); i {
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
		file_datahub_gpu_gpu_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GpuPrediction); i {
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
			RawDescriptor: file_datahub_gpu_gpu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datahub_gpu_gpu_proto_goTypes,
		DependencyIndexes: file_datahub_gpu_gpu_proto_depIdxs,
		MessageInfos:      file_datahub_gpu_gpu_proto_msgTypes,
	}.Build()
	File_datahub_gpu_gpu_proto = out.File
	file_datahub_gpu_gpu_proto_rawDesc = nil
	file_datahub_gpu_gpu_proto_goTypes = nil
	file_datahub_gpu_gpu_proto_depIdxs = nil
}
