// This file has messages related to metric data of containers, pods, and nodes

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: alameda_api/v1alpha1/datahub/common/metrics.proto

package common

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
// Metric type. A metric may be CPU, memory and etc.
type MetricType int32

const (
	MetricType_METRICS_TYPE_UNDEFINED   MetricType = 0
	MetricType_CPU_SECONDS_TOTAL        MetricType = 1
	MetricType_CPU_CORES_ALLOCATABLE    MetricType = 2
	MetricType_CPU_MILLICORES_TOTAL     MetricType = 3
	MetricType_CPU_MILLICORES_AVAIL     MetricType = 4
	MetricType_CPU_MILLICORES_USAGE     MetricType = 5
	MetricType_CPU_MILLICORES_USAGE_PCT MetricType = 6
	MetricType_MEMORY_BYTES_ALLOCATABLE MetricType = 7
	MetricType_MEMORY_BYTES_TOTAL       MetricType = 8
	MetricType_MEMORY_BYTES_AVAIL       MetricType = 9
	MetricType_MEMORY_BYTES_USAGE       MetricType = 10
	MetricType_MEMORY_BYTES_USAGE_PCT   MetricType = 11
	MetricType_FS_BYTES_TOTAL           MetricType = 12
	MetricType_FS_BYTES_AVAIL           MetricType = 13
	MetricType_FS_BYTES_USAGE           MetricType = 14
	MetricType_FS_BYTES_USAGE_PCT       MetricType = 15
	MetricType_HTTP_REQUESTS_COUNT      MetricType = 16
	MetricType_HTTP_REQUESTS_TOTAL      MetricType = 17
	MetricType_HTTP_RESPONSE_COUNT      MetricType = 18
	MetricType_HTTP_RESPONSE_TOTAL      MetricType = 19
	MetricType_DISK_IO_SECONDS_TOTAL    MetricType = 20
	MetricType_DISK_IO_UTILIZATION      MetricType = 21
	MetricType_RESTARTS_TOTAL           MetricType = 22
	MetricType_UNSCHEDULABLE            MetricType = 23
	MetricType_HEALTH                   MetricType = 24
	MetricType_POWER_USAGE_WATTS        MetricType = 25
	MetricType_TEMPERATURE_CELSIUS      MetricType = 26
	MetricType_DUTY_CYCLE               MetricType = 27
	MetricType_CURRENT_OFFSET           MetricType = 28
	MetricType_LAG                      MetricType = 29
	MetricType_LATENCY                  MetricType = 30
	MetricType_NUMBER                   MetricType = 31
)

// Enum value maps for MetricType.
var (
	MetricType_name = map[int32]string{
		0:  "METRICS_TYPE_UNDEFINED",
		1:  "CPU_SECONDS_TOTAL",
		2:  "CPU_CORES_ALLOCATABLE",
		3:  "CPU_MILLICORES_TOTAL",
		4:  "CPU_MILLICORES_AVAIL",
		5:  "CPU_MILLICORES_USAGE",
		6:  "CPU_MILLICORES_USAGE_PCT",
		7:  "MEMORY_BYTES_ALLOCATABLE",
		8:  "MEMORY_BYTES_TOTAL",
		9:  "MEMORY_BYTES_AVAIL",
		10: "MEMORY_BYTES_USAGE",
		11: "MEMORY_BYTES_USAGE_PCT",
		12: "FS_BYTES_TOTAL",
		13: "FS_BYTES_AVAIL",
		14: "FS_BYTES_USAGE",
		15: "FS_BYTES_USAGE_PCT",
		16: "HTTP_REQUESTS_COUNT",
		17: "HTTP_REQUESTS_TOTAL",
		18: "HTTP_RESPONSE_COUNT",
		19: "HTTP_RESPONSE_TOTAL",
		20: "DISK_IO_SECONDS_TOTAL",
		21: "DISK_IO_UTILIZATION",
		22: "RESTARTS_TOTAL",
		23: "UNSCHEDULABLE",
		24: "HEALTH",
		25: "POWER_USAGE_WATTS",
		26: "TEMPERATURE_CELSIUS",
		27: "DUTY_CYCLE",
		28: "CURRENT_OFFSET",
		29: "LAG",
		30: "LATENCY",
		31: "NUMBER",
	}
	MetricType_value = map[string]int32{
		"METRICS_TYPE_UNDEFINED":   0,
		"CPU_SECONDS_TOTAL":        1,
		"CPU_CORES_ALLOCATABLE":    2,
		"CPU_MILLICORES_TOTAL":     3,
		"CPU_MILLICORES_AVAIL":     4,
		"CPU_MILLICORES_USAGE":     5,
		"CPU_MILLICORES_USAGE_PCT": 6,
		"MEMORY_BYTES_ALLOCATABLE": 7,
		"MEMORY_BYTES_TOTAL":       8,
		"MEMORY_BYTES_AVAIL":       9,
		"MEMORY_BYTES_USAGE":       10,
		"MEMORY_BYTES_USAGE_PCT":   11,
		"FS_BYTES_TOTAL":           12,
		"FS_BYTES_AVAIL":           13,
		"FS_BYTES_USAGE":           14,
		"FS_BYTES_USAGE_PCT":       15,
		"HTTP_REQUESTS_COUNT":      16,
		"HTTP_REQUESTS_TOTAL":      17,
		"HTTP_RESPONSE_COUNT":      18,
		"HTTP_RESPONSE_TOTAL":      19,
		"DISK_IO_SECONDS_TOTAL":    20,
		"DISK_IO_UTILIZATION":      21,
		"RESTARTS_TOTAL":           22,
		"UNSCHEDULABLE":            23,
		"HEALTH":                   24,
		"POWER_USAGE_WATTS":        25,
		"TEMPERATURE_CELSIUS":      26,
		"DUTY_CYCLE":               27,
		"CURRENT_OFFSET":           28,
		"LAG":                      29,
		"LATENCY":                  30,
		"NUMBER":                   31,
	}
)

func (x MetricType) Enum() *MetricType {
	p := new(MetricType)
	*p = x
	return p
}

func (x MetricType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricType) Descriptor() protoreflect.EnumDescriptor {
	return file_alameda_api_v1alpha1_datahub_common_metrics_proto_enumTypes[0].Descriptor()
}

func (MetricType) Type() protoreflect.EnumType {
	return &file_alameda_api_v1alpha1_datahub_common_metrics_proto_enumTypes[0]
}

func (x MetricType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricType.Descriptor instead.
func (MetricType) EnumDescriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_common_metrics_proto_rawDescGZIP(), []int{0}
}

//*
// Represents Kubernetes resources which will be allocated to pods.
type ResourceName int32

const (
	ResourceName_RESOURCE_NAME_UNDEFINED ResourceName = 0
	ResourceName_CPU                     ResourceName = 1
	ResourceName_MEMORY                  ResourceName = 2
)

// Enum value maps for ResourceName.
var (
	ResourceName_name = map[int32]string{
		0: "RESOURCE_NAME_UNDEFINED",
		1: "CPU",
		2: "MEMORY",
	}
	ResourceName_value = map[string]int32{
		"RESOURCE_NAME_UNDEFINED": 0,
		"CPU":                     1,
		"MEMORY":                  2,
	}
)

func (x ResourceName) Enum() *ResourceName {
	p := new(ResourceName)
	*p = x
	return p
}

func (x ResourceName) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceName) Descriptor() protoreflect.EnumDescriptor {
	return file_alameda_api_v1alpha1_datahub_common_metrics_proto_enumTypes[1].Descriptor()
}

func (ResourceName) Type() protoreflect.EnumType {
	return &file_alameda_api_v1alpha1_datahub_common_metrics_proto_enumTypes[1]
}

func (x ResourceName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceName.Descriptor instead.
func (ResourceName) EnumDescriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_common_metrics_proto_rawDescGZIP(), []int{1}
}

//*
// Represents a data point of time-series metric data.
type Sample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time     *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	EndTime  *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	NumValue string               `protobuf:"bytes,3,opt,name=num_value,json=numValue,proto3" json:"num_value,omitempty"`
}

func (x *Sample) Reset() {
	*x = Sample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_common_metrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sample) ProtoMessage() {}

func (x *Sample) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_common_metrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sample.ProtoReflect.Descriptor instead.
func (*Sample) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_common_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *Sample) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Sample) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *Sample) GetNumValue() string {
	if x != nil {
		return x.NumValue
	}
	return ""
}

//*
// Represents a piece of metreic data.
type MetricData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetricType  MetricType `protobuf:"varint,1,opt,name=metric_type,json=metricType,proto3,enum=containersai.alameda.v1alpha1.datahub.common.MetricType" json:"metric_type,omitempty"`
	Data        []*Sample  `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"` // Data can be time series or non-time series.
	Granularity int64      `protobuf:"varint,3,opt,name=granularity,proto3" json:"granularity,omitempty"`
}

func (x *MetricData) Reset() {
	*x = MetricData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_common_metrics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricData) ProtoMessage() {}

func (x *MetricData) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_common_metrics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricData.ProtoReflect.Descriptor instead.
func (*MetricData) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_common_metrics_proto_rawDescGZIP(), []int{1}
}

func (x *MetricData) GetMetricType() MetricType {
	if x != nil {
		return x.MetricType
	}
	return MetricType_METRICS_TYPE_UNDEFINED
}

func (x *MetricData) GetData() []*Sample {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *MetricData) GetGranularity() int64 {
	if x != nil {
		return x.Granularity
	}
	return 0
}

var File_alameda_api_v1alpha1_datahub_common_metrics_proto protoreflect.FileDescriptor

var file_alameda_api_v1alpha1_datahub_common_metrics_proto_rawDesc = []byte{
	0x0a, 0x31, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x2c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61,
	0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x06, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x2e, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0xd3, 0x01, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x59, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x48, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x67, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x67, 0x72, 0x61, 0x6e,
	0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x2a, 0xdd, 0x05, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43,
	0x53, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x50, 0x55, 0x5f, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44,
	0x53, 0x5f, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x50, 0x55,
	0x5f, 0x43, 0x4f, 0x52, 0x45, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x41, 0x42,
	0x4c, 0x45, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x50, 0x55, 0x5f, 0x4d, 0x49, 0x4c, 0x4c,
	0x49, 0x43, 0x4f, 0x52, 0x45, 0x53, 0x5f, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x10, 0x03, 0x12, 0x18,
	0x0a, 0x14, 0x43, 0x50, 0x55, 0x5f, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x43, 0x4f, 0x52, 0x45, 0x53,
	0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x50, 0x55, 0x5f,
	0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x43, 0x4f, 0x52, 0x45, 0x53, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45,
	0x10, 0x05, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x50, 0x55, 0x5f, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x43,
	0x4f, 0x52, 0x45, 0x53, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x50, 0x43, 0x54, 0x10, 0x06,
	0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59, 0x5f, 0x42, 0x59, 0x54, 0x45, 0x53,
	0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x07, 0x12, 0x16,
	0x0a, 0x12, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59, 0x5f, 0x42, 0x59, 0x54, 0x45, 0x53, 0x5f, 0x54,
	0x4f, 0x54, 0x41, 0x4c, 0x10, 0x08, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59,
	0x5f, 0x42, 0x59, 0x54, 0x45, 0x53, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x10, 0x09, 0x12, 0x16,
	0x0a, 0x12, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59, 0x5f, 0x42, 0x59, 0x54, 0x45, 0x53, 0x5f, 0x55,
	0x53, 0x41, 0x47, 0x45, 0x10, 0x0a, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59,
	0x5f, 0x42, 0x59, 0x54, 0x45, 0x53, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x50, 0x43, 0x54,
	0x10, 0x0b, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x53, 0x5f, 0x42, 0x59, 0x54, 0x45, 0x53, 0x5f, 0x54,
	0x4f, 0x54, 0x41, 0x4c, 0x10, 0x0c, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x53, 0x5f, 0x42, 0x59, 0x54,
	0x45, 0x53, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x10, 0x0d, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x53,
	0x5f, 0x42, 0x59, 0x54, 0x45, 0x53, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x10, 0x0e, 0x12, 0x16,
	0x0a, 0x12, 0x46, 0x53, 0x5f, 0x42, 0x59, 0x54, 0x45, 0x53, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45,
	0x5f, 0x50, 0x43, 0x54, 0x10, 0x0f, 0x12, 0x17, 0x0a, 0x13, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x52,
	0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x53, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x10, 0x12,
	0x17, 0x0a, 0x13, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x53,
	0x5f, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x10, 0x11, 0x12, 0x17, 0x0a, 0x13, 0x48, 0x54, 0x54, 0x50,
	0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10,
	0x12, 0x12, 0x17, 0x0a, 0x13, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e,
	0x53, 0x45, 0x5f, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x10, 0x13, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x49,
	0x53, 0x4b, 0x5f, 0x49, 0x4f, 0x5f, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x53, 0x5f, 0x54, 0x4f,
	0x54, 0x41, 0x4c, 0x10, 0x14, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x49, 0x53, 0x4b, 0x5f, 0x49, 0x4f,
	0x5f, 0x55, 0x54, 0x49, 0x4c, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x15, 0x12, 0x12,
	0x0a, 0x0e, 0x52, 0x45, 0x53, 0x54, 0x41, 0x52, 0x54, 0x53, 0x5f, 0x54, 0x4f, 0x54, 0x41, 0x4c,
	0x10, 0x16, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x41,
	0x42, 0x4c, 0x45, 0x10, 0x17, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x10,
	0x18, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x4f, 0x57, 0x45, 0x52, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45,
	0x5f, 0x57, 0x41, 0x54, 0x54, 0x53, 0x10, 0x19, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x45, 0x4d, 0x50,
	0x45, 0x52, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x43, 0x45, 0x4c, 0x53, 0x49, 0x55, 0x53, 0x10,
	0x1a, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x55, 0x54, 0x59, 0x5f, 0x43, 0x59, 0x43, 0x4c, 0x45, 0x10,
	0x1b, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x54, 0x5f, 0x4f, 0x46, 0x46,
	0x53, 0x45, 0x54, 0x10, 0x1c, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x41, 0x47, 0x10, 0x1d, 0x12, 0x0b,
	0x0a, 0x07, 0x4c, 0x41, 0x54, 0x45, 0x4e, 0x43, 0x59, 0x10, 0x1e, 0x12, 0x0a, 0x0a, 0x06, 0x4e,
	0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x1f, 0x2a, 0x40, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x45, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x50, 0x55, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59, 0x10, 0x02, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x73, 0x2d, 0x61, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64,
	0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alameda_api_v1alpha1_datahub_common_metrics_proto_rawDescOnce sync.Once
	file_alameda_api_v1alpha1_datahub_common_metrics_proto_rawDescData = file_alameda_api_v1alpha1_datahub_common_metrics_proto_rawDesc
)

func file_alameda_api_v1alpha1_datahub_common_metrics_proto_rawDescGZIP() []byte {
	file_alameda_api_v1alpha1_datahub_common_metrics_proto_rawDescOnce.Do(func() {
		file_alameda_api_v1alpha1_datahub_common_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_alameda_api_v1alpha1_datahub_common_metrics_proto_rawDescData)
	})
	return file_alameda_api_v1alpha1_datahub_common_metrics_proto_rawDescData
}

var file_alameda_api_v1alpha1_datahub_common_metrics_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_alameda_api_v1alpha1_datahub_common_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_alameda_api_v1alpha1_datahub_common_metrics_proto_goTypes = []interface{}{
	(MetricType)(0),             // 0: containersai.alameda.v1alpha1.datahub.common.MetricType
	(ResourceName)(0),           // 1: containersai.alameda.v1alpha1.datahub.common.ResourceName
	(*Sample)(nil),              // 2: containersai.alameda.v1alpha1.datahub.common.Sample
	(*MetricData)(nil),          // 3: containersai.alameda.v1alpha1.datahub.common.MetricData
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_alameda_api_v1alpha1_datahub_common_metrics_proto_depIdxs = []int32{
	4, // 0: containersai.alameda.v1alpha1.datahub.common.Sample.time:type_name -> google.protobuf.Timestamp
	4, // 1: containersai.alameda.v1alpha1.datahub.common.Sample.end_time:type_name -> google.protobuf.Timestamp
	0, // 2: containersai.alameda.v1alpha1.datahub.common.MetricData.metric_type:type_name -> containersai.alameda.v1alpha1.datahub.common.MetricType
	2, // 3: containersai.alameda.v1alpha1.datahub.common.MetricData.data:type_name -> containersai.alameda.v1alpha1.datahub.common.Sample
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_alameda_api_v1alpha1_datahub_common_metrics_proto_init() }
func file_alameda_api_v1alpha1_datahub_common_metrics_proto_init() {
	if File_alameda_api_v1alpha1_datahub_common_metrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_alameda_api_v1alpha1_datahub_common_metrics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sample); i {
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
		file_alameda_api_v1alpha1_datahub_common_metrics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricData); i {
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
			RawDescriptor: file_alameda_api_v1alpha1_datahub_common_metrics_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_alameda_api_v1alpha1_datahub_common_metrics_proto_goTypes,
		DependencyIndexes: file_alameda_api_v1alpha1_datahub_common_metrics_proto_depIdxs,
		EnumInfos:         file_alameda_api_v1alpha1_datahub_common_metrics_proto_enumTypes,
		MessageInfos:      file_alameda_api_v1alpha1_datahub_common_metrics_proto_msgTypes,
	}.Build()
	File_alameda_api_v1alpha1_datahub_common_metrics_proto = out.File
	file_alameda_api_v1alpha1_datahub_common_metrics_proto_rawDesc = nil
	file_alameda_api_v1alpha1_datahub_common_metrics_proto_goTypes = nil
	file_alameda_api_v1alpha1_datahub_common_metrics_proto_depIdxs = nil
}
