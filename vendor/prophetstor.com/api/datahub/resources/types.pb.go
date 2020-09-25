// This file has messages related to compute resources

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: datahub/resources/types.proto

package resources

import (
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
// Represents the private alameda pod specification.
type AlamedaPodSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlamedaScaler          *ObjectMeta           `protobuf:"bytes,1,opt,name=alameda_scaler,json=alamedaScaler,proto3" json:"alameda_scaler,omitempty"`
	ScalingTool            ScalingTool           `protobuf:"varint,2,opt,name=scaling_tool,json=scalingTool,proto3,enum=prophetstor.datahub.resources.ScalingTool" json:"scaling_tool,omitempty"`
	UsedRecommendationId   string                `protobuf:"bytes,3,opt,name=used_recommendation_id,json=usedRecommendationId,proto3" json:"used_recommendation_id,omitempty"`
	Policy                 RecommendationPolicy  `protobuf:"varint,4,opt,name=policy,proto3,enum=prophetstor.datahub.resources.RecommendationPolicy" json:"policy,omitempty"`
	AlamedaScalerResources *ResourceRequirements `protobuf:"bytes,5,opt,name=alameda_scaler_resources,json=alamedaScalerResources,proto3" json:"alameda_scaler_resources,omitempty"`
}

func (x *AlamedaPodSpec) Reset() {
	*x = AlamedaPodSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_resources_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlamedaPodSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlamedaPodSpec) ProtoMessage() {}

func (x *AlamedaPodSpec) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_resources_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlamedaPodSpec.ProtoReflect.Descriptor instead.
func (*AlamedaPodSpec) Descriptor() ([]byte, []int) {
	return file_datahub_resources_types_proto_rawDescGZIP(), []int{0}
}

func (x *AlamedaPodSpec) GetAlamedaScaler() *ObjectMeta {
	if x != nil {
		return x.AlamedaScaler
	}
	return nil
}

func (x *AlamedaPodSpec) GetScalingTool() ScalingTool {
	if x != nil {
		return x.ScalingTool
	}
	return ScalingTool_SCALING_TOOL_UNDEFINED
}

func (x *AlamedaPodSpec) GetUsedRecommendationId() string {
	if x != nil {
		return x.UsedRecommendationId
	}
	return ""
}

func (x *AlamedaPodSpec) GetPolicy() RecommendationPolicy {
	if x != nil {
		return x.Policy
	}
	return RecommendationPolicy_RECOMMENDATION_POLICY_UNDEFINED
}

func (x *AlamedaPodSpec) GetAlamedaScalerResources() *ResourceRequirements {
	if x != nil {
		return x.AlamedaScalerResources
	}
	return nil
}

//*
// Represents the private alameda controller specification.
type AlamedaControllerSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlamedaScaler                 *ObjectMeta          `protobuf:"bytes,1,opt,name=alameda_scaler,json=alamedaScaler,proto3" json:"alameda_scaler,omitempty"`
	ScalingTool                   ScalingTool          `protobuf:"varint,2,opt,name=scaling_tool,json=scalingTool,proto3,enum=prophetstor.datahub.resources.ScalingTool" json:"scaling_tool,omitempty"`
	Policy                        RecommendationPolicy `protobuf:"varint,3,opt,name=policy,proto3,enum=prophetstor.datahub.resources.RecommendationPolicy" json:"policy,omitempty"`
	MinReplicas                   int32                `protobuf:"varint,4,opt,name=min_replicas,json=minReplicas,proto3" json:"min_replicas,omitempty"`
	MaxReplicas                   int32                `protobuf:"varint,5,opt,name=max_replicas,json=maxReplicas,proto3" json:"max_replicas,omitempty"`
	EnableRecommendationExecution bool                 `protobuf:"varint,6,opt,name=enable_recommendation_execution,json=enableRecommendationExecution,proto3" json:"enable_recommendation_execution,omitempty"`
}

func (x *AlamedaControllerSpec) Reset() {
	*x = AlamedaControllerSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_resources_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlamedaControllerSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlamedaControllerSpec) ProtoMessage() {}

func (x *AlamedaControllerSpec) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_resources_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlamedaControllerSpec.ProtoReflect.Descriptor instead.
func (*AlamedaControllerSpec) Descriptor() ([]byte, []int) {
	return file_datahub_resources_types_proto_rawDescGZIP(), []int{1}
}

func (x *AlamedaControllerSpec) GetAlamedaScaler() *ObjectMeta {
	if x != nil {
		return x.AlamedaScaler
	}
	return nil
}

func (x *AlamedaControllerSpec) GetScalingTool() ScalingTool {
	if x != nil {
		return x.ScalingTool
	}
	return ScalingTool_SCALING_TOOL_UNDEFINED
}

func (x *AlamedaControllerSpec) GetPolicy() RecommendationPolicy {
	if x != nil {
		return x.Policy
	}
	return RecommendationPolicy_RECOMMENDATION_POLICY_UNDEFINED
}

func (x *AlamedaControllerSpec) GetMinReplicas() int32 {
	if x != nil {
		return x.MinReplicas
	}
	return 0
}

func (x *AlamedaControllerSpec) GetMaxReplicas() int32 {
	if x != nil {
		return x.MaxReplicas
	}
	return 0
}

func (x *AlamedaControllerSpec) GetEnableRecommendationExecution() bool {
	if x != nil {
		return x.EnableRecommendationExecution
	}
	return false
}

//*
// Represents the private alameda applcation specification.
type AlamedaApplicationSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AlamedaApplicationSpec) Reset() {
	*x = AlamedaApplicationSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_resources_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlamedaApplicationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlamedaApplicationSpec) ProtoMessage() {}

func (x *AlamedaApplicationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_resources_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlamedaApplicationSpec.ProtoReflect.Descriptor instead.
func (*AlamedaApplicationSpec) Descriptor() ([]byte, []int) {
	return file_datahub_resources_types_proto_rawDescGZIP(), []int{2}
}

//*
// Represents the private alameda node specification.
type AlamedaNodeSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider *Provider `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *AlamedaNodeSpec) Reset() {
	*x = AlamedaNodeSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_resources_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlamedaNodeSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlamedaNodeSpec) ProtoMessage() {}

func (x *AlamedaNodeSpec) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_resources_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlamedaNodeSpec.ProtoReflect.Descriptor instead.
func (*AlamedaNodeSpec) Descriptor() ([]byte, []int) {
	return file_datahub_resources_types_proto_rawDescGZIP(), []int{3}
}

func (x *AlamedaNodeSpec) GetProvider() *Provider {
	if x != nil {
		return x.Provider
	}
	return nil
}

//*
// Represents the capacity of a Kubernetes node.
type Capacity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CpuCores                 int64 `protobuf:"varint,1,opt,name=cpu_cores,json=cpuCores,proto3" json:"cpu_cores,omitempty"`
	MemoryBytes              int64 `protobuf:"varint,2,opt,name=memory_bytes,json=memoryBytes,proto3" json:"memory_bytes,omitempty"`
	NetworkMegabitsPerSecond int64 `protobuf:"varint,3,opt,name=network_megabits_per_second,json=networkMegabitsPerSecond,proto3" json:"network_megabits_per_second,omitempty"`
}

func (x *Capacity) Reset() {
	*x = Capacity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_resources_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Capacity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Capacity) ProtoMessage() {}

func (x *Capacity) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_resources_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Capacity.ProtoReflect.Descriptor instead.
func (*Capacity) Descriptor() ([]byte, []int) {
	return file_datahub_resources_types_proto_rawDescGZIP(), []int{4}
}

func (x *Capacity) GetCpuCores() int64 {
	if x != nil {
		return x.CpuCores
	}
	return 0
}

func (x *Capacity) GetMemoryBytes() int64 {
	if x != nil {
		return x.MemoryBytes
	}
	return 0
}

func (x *Capacity) GetNetworkMegabitsPerSecond() int64 {
	if x != nil {
		return x.NetworkMegabitsPerSecond
	}
	return 0
}

//*
// The information of cloud service provider.
type Provider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider     string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	InstanceType string `protobuf:"bytes,2,opt,name=instance_type,json=instanceType,proto3" json:"instance_type,omitempty"`
	Region       string `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	Zone         string `protobuf:"bytes,4,opt,name=zone,proto3" json:"zone,omitempty"`
	Os           string `protobuf:"bytes,5,opt,name=os,proto3" json:"os,omitempty"`
	Role         string `protobuf:"bytes,6,opt,name=role,proto3" json:"role,omitempty"`
	InstanceId   string `protobuf:"bytes,7,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	StorageSize  int64  `protobuf:"varint,8,opt,name=storage_size,json=storageSize,proto3" json:"storage_size,omitempty"`
}

func (x *Provider) Reset() {
	*x = Provider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_resources_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Provider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Provider) ProtoMessage() {}

func (x *Provider) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_resources_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Provider.ProtoReflect.Descriptor instead.
func (*Provider) Descriptor() ([]byte, []int) {
	return file_datahub_resources_types_proto_rawDescGZIP(), []int{5}
}

func (x *Provider) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *Provider) GetInstanceType() string {
	if x != nil {
		return x.InstanceType
	}
	return ""
}

func (x *Provider) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Provider) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *Provider) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *Provider) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *Provider) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *Provider) GetStorageSize() int64 {
	if x != nil {
		return x.StorageSize
	}
	return 0
}

//*
// ResourceRequirements describes the compute resource requirements
type ResourceRequirements struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// limits describes the maximum amount of compute resources allowed
	// use enum ResourceName as key of the map which defined in common
	Limits map[int32]string `protobuf:"bytes,1,rep,name=limits,proto3" json:"limits,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// requests describes the minimum amount of compute resources required
	// use enum ResourceName as key of the map which defined in common
	Requests map[int32]string `protobuf:"bytes,2,rep,name=requests,proto3" json:"requests,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ResourceRequirements) Reset() {
	*x = ResourceRequirements{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_resources_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceRequirements) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceRequirements) ProtoMessage() {}

func (x *ResourceRequirements) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_resources_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceRequirements.ProtoReflect.Descriptor instead.
func (*ResourceRequirements) Descriptor() ([]byte, []int) {
	return file_datahub_resources_types_proto_rawDescGZIP(), []int{6}
}

func (x *ResourceRequirements) GetLimits() map[int32]string {
	if x != nil {
		return x.Limits
	}
	return nil
}

func (x *ResourceRequirements) GetRequests() map[int32]string {
	if x != nil {
		return x.Requests
	}
	return nil
}

var File_datahub_resources_types_proto protoreflect.FileDescriptor

var file_datahub_resources_types_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1d, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x20,
	0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa3, 0x03, 0x0a, 0x0e, 0x41, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x50, 0x6f,
	0x64, 0x53, 0x70, 0x65, 0x63, 0x12, 0x50, 0x0a, 0x0e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61,
	0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0d, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64,
	0x61, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x0c, 0x73, 0x63, 0x61, 0x6c, 0x69,
	0x6e, 0x67, 0x5f, 0x74, 0x6f, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e,
	0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x63,
	0x61, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6f, 0x6c, 0x52, 0x0b, 0x73, 0x63, 0x61, 0x6c, 0x69,
	0x6e, 0x67, 0x54, 0x6f, 0x6f, 0x6c, 0x12, 0x34, 0x0a, 0x16, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x72,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x75, 0x73, 0x65, 0x64, 0x52, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x06,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x70,
	0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68,
	0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x6d, 0x0a, 0x18, 0x61, 0x6c, 0x61,
	0x6d, 0x65, 0x64, 0x61, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x70, 0x72,
	0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75,
	0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x16, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x93, 0x03, 0x0a, 0x15, 0x41, 0x6c, 0x61,
	0x6d, 0x65, 0x64, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x50, 0x0a, 0x0e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f,
	0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0d, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x53, 0x63,
	0x61, 0x6c, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x0c, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x5f,
	0x74, 0x6f, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f,
	0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x69,
	0x6e, 0x67, 0x54, 0x6f, 0x6f, 0x6c, 0x52, 0x0b, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x54,
	0x6f, 0x6f, 0x6c, 0x12, 0x4b, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f,
	0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x46, 0x0a, 0x1f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x1d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x18,
	0x0a, 0x16, 0x41, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x22, 0x56, 0x0a, 0x0f, 0x41, 0x6c, 0x61, 0x6d,
	0x65, 0x64, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x43, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x22, 0x89, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x70, 0x75, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x63, 0x70, 0x75, 0x43, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x3d, 0x0a,
	0x1b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6d, 0x65, 0x67, 0x61, 0x62, 0x69, 0x74,
	0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x18, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4d, 0x65, 0x67, 0x61, 0x62,
	0x69, 0x74, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x22, 0xdf, 0x01, 0x0a,
	0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xc6,
	0x02, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x57, 0x0a, 0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65,
	0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73,
	0x12, 0x5d, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x1a,
	0x39, 0x0a, 0x0b, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x27, 0x5a, 0x25, 0x70, 0x72, 0x6f, 0x70, 0x68,
	0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datahub_resources_types_proto_rawDescOnce sync.Once
	file_datahub_resources_types_proto_rawDescData = file_datahub_resources_types_proto_rawDesc
)

func file_datahub_resources_types_proto_rawDescGZIP() []byte {
	file_datahub_resources_types_proto_rawDescOnce.Do(func() {
		file_datahub_resources_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_datahub_resources_types_proto_rawDescData)
	})
	return file_datahub_resources_types_proto_rawDescData
}

var file_datahub_resources_types_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_datahub_resources_types_proto_goTypes = []interface{}{
	(*AlamedaPodSpec)(nil),         // 0: prophetstor.datahub.resources.AlamedaPodSpec
	(*AlamedaControllerSpec)(nil),  // 1: prophetstor.datahub.resources.AlamedaControllerSpec
	(*AlamedaApplicationSpec)(nil), // 2: prophetstor.datahub.resources.AlamedaApplicationSpec
	(*AlamedaNodeSpec)(nil),        // 3: prophetstor.datahub.resources.AlamedaNodeSpec
	(*Capacity)(nil),               // 4: prophetstor.datahub.resources.Capacity
	(*Provider)(nil),               // 5: prophetstor.datahub.resources.Provider
	(*ResourceRequirements)(nil),   // 6: prophetstor.datahub.resources.ResourceRequirements
	nil,                            // 7: prophetstor.datahub.resources.ResourceRequirements.LimitsEntry
	nil,                            // 8: prophetstor.datahub.resources.ResourceRequirements.RequestsEntry
	(*ObjectMeta)(nil),             // 9: prophetstor.datahub.resources.ObjectMeta
	(ScalingTool)(0),               // 10: prophetstor.datahub.resources.ScalingTool
	(RecommendationPolicy)(0),      // 11: prophetstor.datahub.resources.RecommendationPolicy
}
var file_datahub_resources_types_proto_depIdxs = []int32{
	9,  // 0: prophetstor.datahub.resources.AlamedaPodSpec.alameda_scaler:type_name -> prophetstor.datahub.resources.ObjectMeta
	10, // 1: prophetstor.datahub.resources.AlamedaPodSpec.scaling_tool:type_name -> prophetstor.datahub.resources.ScalingTool
	11, // 2: prophetstor.datahub.resources.AlamedaPodSpec.policy:type_name -> prophetstor.datahub.resources.RecommendationPolicy
	6,  // 3: prophetstor.datahub.resources.AlamedaPodSpec.alameda_scaler_resources:type_name -> prophetstor.datahub.resources.ResourceRequirements
	9,  // 4: prophetstor.datahub.resources.AlamedaControllerSpec.alameda_scaler:type_name -> prophetstor.datahub.resources.ObjectMeta
	10, // 5: prophetstor.datahub.resources.AlamedaControllerSpec.scaling_tool:type_name -> prophetstor.datahub.resources.ScalingTool
	11, // 6: prophetstor.datahub.resources.AlamedaControllerSpec.policy:type_name -> prophetstor.datahub.resources.RecommendationPolicy
	5,  // 7: prophetstor.datahub.resources.AlamedaNodeSpec.provider:type_name -> prophetstor.datahub.resources.Provider
	7,  // 8: prophetstor.datahub.resources.ResourceRequirements.limits:type_name -> prophetstor.datahub.resources.ResourceRequirements.LimitsEntry
	8,  // 9: prophetstor.datahub.resources.ResourceRequirements.requests:type_name -> prophetstor.datahub.resources.ResourceRequirements.RequestsEntry
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_datahub_resources_types_proto_init() }
func file_datahub_resources_types_proto_init() {
	if File_datahub_resources_types_proto != nil {
		return
	}
	file_datahub_resources_metadata_proto_init()
	file_datahub_resources_policies_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_datahub_resources_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlamedaPodSpec); i {
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
		file_datahub_resources_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlamedaControllerSpec); i {
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
		file_datahub_resources_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlamedaApplicationSpec); i {
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
		file_datahub_resources_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlamedaNodeSpec); i {
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
		file_datahub_resources_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Capacity); i {
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
		file_datahub_resources_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Provider); i {
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
		file_datahub_resources_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceRequirements); i {
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
			RawDescriptor: file_datahub_resources_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datahub_resources_types_proto_goTypes,
		DependencyIndexes: file_datahub_resources_types_proto_depIdxs,
		MessageInfos:      file_datahub_resources_types_proto_msgTypes,
	}.Build()
	File_datahub_resources_types_proto = out.File
	file_datahub_resources_types_proto_rawDesc = nil
	file_datahub_resources_types_proto_goTypes = nil
	file_datahub_resources_types_proto_depIdxs = nil
}
