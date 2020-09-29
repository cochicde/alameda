// This file has messages related to recommendations of containers, pods, and nodes

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: datahub/recommendations/types.proto

package recommendations

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	common "prophetstor.com/api/datahub/common"
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

type ControllerRecommendedType int32

const (
	ControllerRecommendedType_CRT_UNDEFINED ControllerRecommendedType = 0
	ControllerRecommendedType_PRIMITIVE     ControllerRecommendedType = 1
	ControllerRecommendedType_K8S           ControllerRecommendedType = 2
)

// Enum value maps for ControllerRecommendedType.
var (
	ControllerRecommendedType_name = map[int32]string{
		0: "CRT_UNDEFINED",
		1: "PRIMITIVE",
		2: "K8S",
	}
	ControllerRecommendedType_value = map[string]int32{
		"CRT_UNDEFINED": 0,
		"PRIMITIVE":     1,
		"K8S":           2,
	}
)

func (x ControllerRecommendedType) Enum() *ControllerRecommendedType {
	p := new(ControllerRecommendedType)
	*p = x
	return p
}

func (x ControllerRecommendedType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ControllerRecommendedType) Descriptor() protoreflect.EnumDescriptor {
	return file_datahub_recommendations_types_proto_enumTypes[0].Descriptor()
}

func (ControllerRecommendedType) Type() protoreflect.EnumType {
	return &file_datahub_recommendations_types_proto_enumTypes[0]
}

func (x ControllerRecommendedType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ControllerRecommendedType.Descriptor instead.
func (ControllerRecommendedType) EnumDescriptor() ([]byte, []int) {
	return file_datahub_recommendations_types_proto_rawDescGZIP(), []int{0}
}

//*
// Represents a private spec of a controller recommendation.
type ControllerRecommendedSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentReplicas    int32                `protobuf:"varint,1,opt,name=current_replicas,json=currentReplicas,proto3" json:"current_replicas,omitempty"`
	DesiredReplicas    int32                `protobuf:"varint,2,opt,name=desired_replicas,json=desiredReplicas,proto3" json:"desired_replicas,omitempty"`
	Time               *timestamp.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	CreateTime         *timestamp.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	CurrentCpuRequests float64              `protobuf:"fixed64,5,opt,name=current_cpu_requests,json=currentCpuRequests,proto3" json:"current_cpu_requests,omitempty"`
	CurrentMemRequests float64              `protobuf:"fixed64,6,opt,name=current_mem_requests,json=currentMemRequests,proto3" json:"current_mem_requests,omitempty"`
	CurrentCpuLimits   float64              `protobuf:"fixed64,7,opt,name=current_cpu_limits,json=currentCpuLimits,proto3" json:"current_cpu_limits,omitempty"`
	CurrentMemLimits   float64              `protobuf:"fixed64,8,opt,name=current_mem_limits,json=currentMemLimits,proto3" json:"current_mem_limits,omitempty"`
	DesiredCpuLimits   float64              `protobuf:"fixed64,9,opt,name=desired_cpu_limits,json=desiredCpuLimits,proto3" json:"desired_cpu_limits,omitempty"`
	DesiredMemLimits   float64              `protobuf:"fixed64,10,opt,name=desired_mem_limits,json=desiredMemLimits,proto3" json:"desired_mem_limits,omitempty"`
	TotalCost          float64              `protobuf:"fixed64,11,opt,name=total_cost,json=totalCost,proto3" json:"total_cost,omitempty"`
}

func (x *ControllerRecommendedSpec) Reset() {
	*x = ControllerRecommendedSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_recommendations_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControllerRecommendedSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControllerRecommendedSpec) ProtoMessage() {}

func (x *ControllerRecommendedSpec) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_recommendations_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControllerRecommendedSpec.ProtoReflect.Descriptor instead.
func (*ControllerRecommendedSpec) Descriptor() ([]byte, []int) {
	return file_datahub_recommendations_types_proto_rawDescGZIP(), []int{0}
}

func (x *ControllerRecommendedSpec) GetCurrentReplicas() int32 {
	if x != nil {
		return x.CurrentReplicas
	}
	return 0
}

func (x *ControllerRecommendedSpec) GetDesiredReplicas() int32 {
	if x != nil {
		return x.DesiredReplicas
	}
	return 0
}

func (x *ControllerRecommendedSpec) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *ControllerRecommendedSpec) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ControllerRecommendedSpec) GetCurrentCpuRequests() float64 {
	if x != nil {
		return x.CurrentCpuRequests
	}
	return 0
}

func (x *ControllerRecommendedSpec) GetCurrentMemRequests() float64 {
	if x != nil {
		return x.CurrentMemRequests
	}
	return 0
}

func (x *ControllerRecommendedSpec) GetCurrentCpuLimits() float64 {
	if x != nil {
		return x.CurrentCpuLimits
	}
	return 0
}

func (x *ControllerRecommendedSpec) GetCurrentMemLimits() float64 {
	if x != nil {
		return x.CurrentMemLimits
	}
	return 0
}

func (x *ControllerRecommendedSpec) GetDesiredCpuLimits() float64 {
	if x != nil {
		return x.DesiredCpuLimits
	}
	return 0
}

func (x *ControllerRecommendedSpec) GetDesiredMemLimits() float64 {
	if x != nil {
		return x.DesiredMemLimits
	}
	return 0
}

func (x *ControllerRecommendedSpec) GetTotalCost() float64 {
	if x != nil {
		return x.TotalCost
	}
	return 0
}

type ControllerRecommendedSpecK8S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentReplicas int32                `protobuf:"varint,1,opt,name=current_replicas,json=currentReplicas,proto3" json:"current_replicas,omitempty"`
	DesiredReplicas int32                `protobuf:"varint,2,opt,name=desired_replicas,json=desiredReplicas,proto3" json:"desired_replicas,omitempty"`
	Time            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	CreateTime      *timestamp.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *ControllerRecommendedSpecK8S) Reset() {
	*x = ControllerRecommendedSpecK8S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_recommendations_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControllerRecommendedSpecK8S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControllerRecommendedSpecK8S) ProtoMessage() {}

func (x *ControllerRecommendedSpecK8S) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_recommendations_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControllerRecommendedSpecK8S.ProtoReflect.Descriptor instead.
func (*ControllerRecommendedSpecK8S) Descriptor() ([]byte, []int) {
	return file_datahub_recommendations_types_proto_rawDescGZIP(), []int{1}
}

func (x *ControllerRecommendedSpecK8S) GetCurrentReplicas() int32 {
	if x != nil {
		return x.CurrentReplicas
	}
	return 0
}

func (x *ControllerRecommendedSpecK8S) GetDesiredReplicas() int32 {
	if x != nil {
		return x.DesiredReplicas
	}
	return 0
}

func (x *ControllerRecommendedSpecK8S) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *ControllerRecommendedSpecK8S) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

//*
// Represents a set of container resource configuration recommendations of a pod.
type Recommendation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaMeta         *schemas.SchemaMeta   `protobuf:"bytes,1,opt,name=schema_meta,json=schemaMeta,proto3" json:"schema_meta,omitempty"`
	RecommendationData []*RecommendationData `protobuf:"bytes,2,rep,name=recommendation_data,json=recommendationData,proto3" json:"recommendation_data,omitempty"`
}

func (x *Recommendation) Reset() {
	*x = Recommendation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_recommendations_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recommendation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recommendation) ProtoMessage() {}

func (x *Recommendation) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_recommendations_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recommendation.ProtoReflect.Descriptor instead.
func (*Recommendation) Descriptor() ([]byte, []int) {
	return file_datahub_recommendations_types_proto_rawDescGZIP(), []int{2}
}

func (x *Recommendation) GetSchemaMeta() *schemas.SchemaMeta {
	if x != nil {
		return x.SchemaMeta
	}
	return nil
}

func (x *Recommendation) GetRecommendationData() []*RecommendationData {
	if x != nil {
		return x.RecommendationData
	}
	return nil
}

//*
// Represents a piece of recommendation data.
type RecommendationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetricType    common.MetricType    `protobuf:"varint,1,opt,name=metric_type,json=metricType,proto3,enum=prophetstor.datahub.common.MetricType" json:"metric_type,omitempty"`
	ResourceQuota common.ResourceQuota `protobuf:"varint,2,opt,name=resource_quota,json=resourceQuota,proto3,enum=prophetstor.datahub.common.ResourceQuota" json:"resource_quota,omitempty"`
	ReadData      *common.ReadData     `protobuf:"bytes,3,opt,name=read_data,json=readData,proto3" json:"read_data,omitempty"`
}

func (x *RecommendationData) Reset() {
	*x = RecommendationData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_recommendations_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendationData) ProtoMessage() {}

func (x *RecommendationData) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_recommendations_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendationData.ProtoReflect.Descriptor instead.
func (*RecommendationData) Descriptor() ([]byte, []int) {
	return file_datahub_recommendations_types_proto_rawDescGZIP(), []int{3}
}

func (x *RecommendationData) GetMetricType() common.MetricType {
	if x != nil {
		return x.MetricType
	}
	return common.MetricType_METRICS_TYPE_UNDEFINED
}

func (x *RecommendationData) GetResourceQuota() common.ResourceQuota {
	if x != nil {
		return x.ResourceQuota
	}
	return common.ResourceQuota_RESOURCE_QUOTA_UNDEFINED
}

func (x *RecommendationData) GetReadData() *common.ReadData {
	if x != nil {
		return x.ReadData
	}
	return nil
}

var File_datahub_recommendations_types_proto protoreflect.FileDescriptor

var file_datahub_recommendations_types_proto_rawDesc = []byte{
	0x0a, 0x23, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x1c, 0x64, 0x61, 0x74, 0x61,
	0x68, 0x75, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75,
	0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x99, 0x04, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x53, 0x70, 0x65, 0x63, 0x12, 0x29,
	0x0a, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x73,
	0x69, 0x72, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0f, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x70, 0x75,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x12, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x70, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6d,
	0x65, 0x6d, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x12, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x63, 0x70, 0x75, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x70, 0x75, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6d,
	0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x6d, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x73, 0x12, 0x2c, 0x0a, 0x12, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x63, 0x70, 0x75,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x64,
	0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x43, 0x70, 0x75, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12,
	0x2c, 0x0a, 0x12, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x6d, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x64, 0x65, 0x73,
	0x69, 0x72, 0x65, 0x64, 0x4d, 0x65, 0x6d, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x73, 0x74, 0x22, 0xe1, 0x01, 0x0a,
	0x1c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x53, 0x70, 0x65, 0x63, 0x4b, 0x38, 0x73, 0x12, 0x29, 0x0a,
	0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x73, 0x69,
	0x72, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0f, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0xc4, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x48, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68,
	0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4d, 0x65, 0x74,
	0x61, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x68, 0x0a,
	0x13, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x70, 0x72, 0x6f,
	0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62,
	0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x12, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x22, 0xf2, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x47,
	0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f,
	0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x50, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x29, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x41, 0x0a, 0x09, 0x72, 0x65, 0x61,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70,
	0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x2a, 0x46, 0x0a, 0x19,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x52, 0x54,
	0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x50, 0x52, 0x49, 0x4d, 0x49, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4b,
	0x38, 0x53, 0x10, 0x02, 0x42, 0x2d, 0x5a, 0x2b, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73,
	0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x68, 0x75, 0x62, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datahub_recommendations_types_proto_rawDescOnce sync.Once
	file_datahub_recommendations_types_proto_rawDescData = file_datahub_recommendations_types_proto_rawDesc
)

func file_datahub_recommendations_types_proto_rawDescGZIP() []byte {
	file_datahub_recommendations_types_proto_rawDescOnce.Do(func() {
		file_datahub_recommendations_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_datahub_recommendations_types_proto_rawDescData)
	})
	return file_datahub_recommendations_types_proto_rawDescData
}

var file_datahub_recommendations_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_datahub_recommendations_types_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_datahub_recommendations_types_proto_goTypes = []interface{}{
	(ControllerRecommendedType)(0),       // 0: prophetstor.datahub.recommendations.ControllerRecommendedType
	(*ControllerRecommendedSpec)(nil),    // 1: prophetstor.datahub.recommendations.ControllerRecommendedSpec
	(*ControllerRecommendedSpecK8S)(nil), // 2: prophetstor.datahub.recommendations.ControllerRecommendedSpecK8s
	(*Recommendation)(nil),               // 3: prophetstor.datahub.recommendations.Recommendation
	(*RecommendationData)(nil),           // 4: prophetstor.datahub.recommendations.RecommendationData
	(*timestamp.Timestamp)(nil),          // 5: google.protobuf.Timestamp
	(*schemas.SchemaMeta)(nil),           // 6: prophetstor.datahub.schemas.SchemaMeta
	(common.MetricType)(0),               // 7: prophetstor.datahub.common.MetricType
	(common.ResourceQuota)(0),            // 8: prophetstor.datahub.common.ResourceQuota
	(*common.ReadData)(nil),              // 9: prophetstor.datahub.common.ReadData
}
var file_datahub_recommendations_types_proto_depIdxs = []int32{
	5, // 0: prophetstor.datahub.recommendations.ControllerRecommendedSpec.time:type_name -> google.protobuf.Timestamp
	5, // 1: prophetstor.datahub.recommendations.ControllerRecommendedSpec.create_time:type_name -> google.protobuf.Timestamp
	5, // 2: prophetstor.datahub.recommendations.ControllerRecommendedSpecK8s.time:type_name -> google.protobuf.Timestamp
	5, // 3: prophetstor.datahub.recommendations.ControllerRecommendedSpecK8s.create_time:type_name -> google.protobuf.Timestamp
	6, // 4: prophetstor.datahub.recommendations.Recommendation.schema_meta:type_name -> prophetstor.datahub.schemas.SchemaMeta
	4, // 5: prophetstor.datahub.recommendations.Recommendation.recommendation_data:type_name -> prophetstor.datahub.recommendations.RecommendationData
	7, // 6: prophetstor.datahub.recommendations.RecommendationData.metric_type:type_name -> prophetstor.datahub.common.MetricType
	8, // 7: prophetstor.datahub.recommendations.RecommendationData.resource_quota:type_name -> prophetstor.datahub.common.ResourceQuota
	9, // 8: prophetstor.datahub.recommendations.RecommendationData.read_data:type_name -> prophetstor.datahub.common.ReadData
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_datahub_recommendations_types_proto_init() }
func file_datahub_recommendations_types_proto_init() {
	if File_datahub_recommendations_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datahub_recommendations_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControllerRecommendedSpec); i {
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
		file_datahub_recommendations_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControllerRecommendedSpecK8S); i {
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
		file_datahub_recommendations_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recommendation); i {
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
		file_datahub_recommendations_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendationData); i {
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
			RawDescriptor: file_datahub_recommendations_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datahub_recommendations_types_proto_goTypes,
		DependencyIndexes: file_datahub_recommendations_types_proto_depIdxs,
		EnumInfos:         file_datahub_recommendations_types_proto_enumTypes,
		MessageInfos:      file_datahub_recommendations_types_proto_msgTypes,
	}.Build()
	File_datahub_recommendations_types_proto = out.File
	file_datahub_recommendations_types_proto_rawDesc = nil
	file_datahub_recommendations_types_proto_goTypes = nil
	file_datahub_recommendations_types_proto_depIdxs = nil
}