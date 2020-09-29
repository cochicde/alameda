// This file has messages related to compute resources

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: datahub/resources/resources.proto

package resources

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
// Represents a container and its containing limit and requeset configurations.
type Container struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Resources *ResourceRequirements `protobuf:"bytes,2,opt,name=resources,proto3" json:"resources,omitempty"`
	Status    *ContainerStatus      `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Container) Reset() {
	*x = Container{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_resources_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Container) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Container) ProtoMessage() {}

func (x *Container) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_resources_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Container.ProtoReflect.Descriptor instead.
func (*Container) Descriptor() ([]byte, []int) {
	return file_datahub_resources_resources_proto_rawDescGZIP(), []int{0}
}

func (x *Container) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Container) GetResources() *ResourceRequirements {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *Container) GetStatus() *ContainerStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

//*
// Represents a Kubernetes pod.
type Pod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectMeta     *ObjectMeta          `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	StartTime      *timestamp.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	ResourceLink   string               `protobuf:"bytes,3,opt,name=resource_link,json=resourceLink,proto3" json:"resource_link,omitempty"`
	AppName        string               `protobuf:"bytes,4,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	AppPartOf      string               `protobuf:"bytes,5,opt,name=app_part_of,json=appPartOf,proto3" json:"app_part_of,omitempty"`
	AlamedaPodSpec *AlamedaPodSpec      `protobuf:"bytes,6,opt,name=alameda_pod_spec,json=alamedaPodSpec,proto3" json:"alameda_pod_spec,omitempty"`
	TopController  *Controller          `protobuf:"bytes,7,opt,name=top_controller,json=topController,proto3" json:"top_controller,omitempty"`
	Status         *PodStatus           `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	Containers     []*Container         `protobuf:"bytes,9,rep,name=containers,proto3" json:"containers,omitempty"`
}

func (x *Pod) Reset() {
	*x = Pod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_resources_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pod) ProtoMessage() {}

func (x *Pod) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_resources_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pod.ProtoReflect.Descriptor instead.
func (*Pod) Descriptor() ([]byte, []int) {
	return file_datahub_resources_resources_proto_rawDescGZIP(), []int{1}
}

func (x *Pod) GetObjectMeta() *ObjectMeta {
	if x != nil {
		return x.ObjectMeta
	}
	return nil
}

func (x *Pod) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Pod) GetResourceLink() string {
	if x != nil {
		return x.ResourceLink
	}
	return ""
}

func (x *Pod) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *Pod) GetAppPartOf() string {
	if x != nil {
		return x.AppPartOf
	}
	return ""
}

func (x *Pod) GetAlamedaPodSpec() *AlamedaPodSpec {
	if x != nil {
		return x.AlamedaPodSpec
	}
	return nil
}

func (x *Pod) GetTopController() *Controller {
	if x != nil {
		return x.TopController
	}
	return nil
}

func (x *Pod) GetStatus() *PodStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *Pod) GetContainers() []*Container {
	if x != nil {
		return x.Containers
	}
	return nil
}

//*
// Represents a Kubernetes namespace.
type Controller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectMeta            *ObjectMeta            `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	Kind                  Kind                   `protobuf:"varint,2,opt,name=kind,proto3,enum=prophetstor.datahub.resources.Kind" json:"kind,omitempty"`
	Replicas              int32                  `protobuf:"varint,3,opt,name=replicas,proto3" json:"replicas,omitempty"`
	SpecReplicas          int32                  `protobuf:"varint,4,opt,name=spec_replicas,json=specReplicas,proto3" json:"spec_replicas,omitempty"`
	AlamedaControllerSpec *AlamedaControllerSpec `protobuf:"bytes,5,opt,name=alameda_controller_spec,json=alamedaControllerSpec,proto3" json:"alameda_controller_spec,omitempty"`
}

func (x *Controller) Reset() {
	*x = Controller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_resources_resources_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Controller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Controller) ProtoMessage() {}

func (x *Controller) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_resources_resources_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Controller.ProtoReflect.Descriptor instead.
func (*Controller) Descriptor() ([]byte, []int) {
	return file_datahub_resources_resources_proto_rawDescGZIP(), []int{2}
}

func (x *Controller) GetObjectMeta() *ObjectMeta {
	if x != nil {
		return x.ObjectMeta
	}
	return nil
}

func (x *Controller) GetKind() Kind {
	if x != nil {
		return x.Kind
	}
	return Kind_KIND_UNDEFINED
}

func (x *Controller) GetReplicas() int32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

func (x *Controller) GetSpecReplicas() int32 {
	if x != nil {
		return x.SpecReplicas
	}
	return 0
}

func (x *Controller) GetAlamedaControllerSpec() *AlamedaControllerSpec {
	if x != nil {
		return x.AlamedaControllerSpec
	}
	return nil
}

//*
// Represents a alameda scaler.
type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectMeta             *ObjectMeta             `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	AlamedaApplicationSpec *AlamedaApplicationSpec `protobuf:"bytes,2,opt,name=alameda_application_spec,json=alamedaApplicationSpec,proto3" json:"alameda_application_spec,omitempty"`
	Controllers            []*Controller           `protobuf:"bytes,3,rep,name=controllers,proto3" json:"controllers,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_resources_resources_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_resources_resources_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_datahub_resources_resources_proto_rawDescGZIP(), []int{3}
}

func (x *Application) GetObjectMeta() *ObjectMeta {
	if x != nil {
		return x.ObjectMeta
	}
	return nil
}

func (x *Application) GetAlamedaApplicationSpec() *AlamedaApplicationSpec {
	if x != nil {
		return x.AlamedaApplicationSpec
	}
	return nil
}

func (x *Application) GetControllers() []*Controller {
	if x != nil {
		return x.Controllers
	}
	return nil
}

//*
// Represents a Kubernetes namespace.
type Namespace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectMeta *ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
}

func (x *Namespace) Reset() {
	*x = Namespace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_resources_resources_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Namespace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Namespace) ProtoMessage() {}

func (x *Namespace) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_resources_resources_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Namespace.ProtoReflect.Descriptor instead.
func (*Namespace) Descriptor() ([]byte, []int) {
	return file_datahub_resources_resources_proto_rawDescGZIP(), []int{4}
}

func (x *Namespace) GetObjectMeta() *ObjectMeta {
	if x != nil {
		return x.ObjectMeta
	}
	return nil
}

//*
// Represents a Kubernetes node.
type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectMeta      *ObjectMeta          `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	StartTime       *timestamp.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Capacity        *Capacity            `protobuf:"bytes,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
	AlamedaNodeSpec *AlamedaNodeSpec     `protobuf:"bytes,4,opt,name=alameda_node_spec,json=alamedaNodeSpec,proto3" json:"alameda_node_spec,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_resources_resources_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_resources_resources_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_datahub_resources_resources_proto_rawDescGZIP(), []int{5}
}

func (x *Node) GetObjectMeta() *ObjectMeta {
	if x != nil {
		return x.ObjectMeta
	}
	return nil
}

func (x *Node) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Node) GetCapacity() *Capacity {
	if x != nil {
		return x.Capacity
	}
	return nil
}

func (x *Node) GetAlamedaNodeSpec() *AlamedaNodeSpec {
	if x != nil {
		return x.AlamedaNodeSpec
	}
	return nil
}

//*
// Represents a Kubernetes cluster.
type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectMeta *ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_resources_resources_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_resources_resources_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_datahub_resources_resources_proto_rawDescGZIP(), []int{6}
}

func (x *Cluster) GetObjectMeta() *ObjectMeta {
	if x != nil {
		return x.ObjectMeta
	}
	return nil
}

var File_datahub_resources_resources_proto protoreflect.FileDescriptor

var file_datahub_resources_resources_proto_rawDesc = []byte{
	0x0a, 0x21, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x1a, 0x20, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x51, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x70,
	0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x09,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x46, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x70,
	0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0xa3, 0x04, 0x0a, 0x03, 0x50, 0x6f, 0x64, 0x12, 0x4a, 0x0a, 0x0b, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x6e,
	0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x5f, 0x6f, 0x66, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x50, 0x61, 0x72, 0x74, 0x4f, 0x66,
	0x12, 0x57, 0x0a, 0x10, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x70, 0x6f, 0x64, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x72, 0x6f,
	0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x6c, 0x61, 0x6d, 0x65,
	0x64, 0x61, 0x50, 0x6f, 0x64, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0e, 0x61, 0x6c, 0x61, 0x6d, 0x65,
	0x64, 0x61, 0x50, 0x6f, 0x64, 0x53, 0x70, 0x65, 0x63, 0x12, 0x50, 0x0a, 0x0e, 0x74, 0x6f, 0x70,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x0d, 0x74, 0x6f,
	0x70, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72,
	0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75,
	0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x48, 0x0a,
	0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x22, 0xc0, 0x02, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72,
	0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75,
	0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x12, 0x37, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x70, 0x65, 0x63, 0x5f,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x73, 0x70, 0x65, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x6c, 0x0a, 0x17,
	0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e,
	0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x6c,
	0x61, 0x6d, 0x65, 0x64, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x15, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x22, 0x97, 0x02, 0x0a, 0x0b, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x0b, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x6f, 0x0a, 0x18, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64,
	0x61, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68,
	0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x52,
	0x16, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x4b, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70,
	0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68,
	0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x73, 0x22, 0x57, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74,
	0x61, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x22, 0xae, 0x02,
	0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72,
	0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75,
	0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x43, 0x0a,
	0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x5a, 0x0a, 0x11, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x6c,
	0x61, 0x6d, 0x65, 0x64, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0f, 0x61,
	0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x70, 0x65, 0x63, 0x22, 0x55,
	0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x0b, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x42, 0x27, 0x5a, 0x25, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x68, 0x75, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datahub_resources_resources_proto_rawDescOnce sync.Once
	file_datahub_resources_resources_proto_rawDescData = file_datahub_resources_resources_proto_rawDesc
)

func file_datahub_resources_resources_proto_rawDescGZIP() []byte {
	file_datahub_resources_resources_proto_rawDescOnce.Do(func() {
		file_datahub_resources_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_datahub_resources_resources_proto_rawDescData)
	})
	return file_datahub_resources_resources_proto_rawDescData
}

var file_datahub_resources_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_datahub_resources_resources_proto_goTypes = []interface{}{
	(*Container)(nil),              // 0: prophetstor.datahub.resources.Container
	(*Pod)(nil),                    // 1: prophetstor.datahub.resources.Pod
	(*Controller)(nil),             // 2: prophetstor.datahub.resources.Controller
	(*Application)(nil),            // 3: prophetstor.datahub.resources.Application
	(*Namespace)(nil),              // 4: prophetstor.datahub.resources.Namespace
	(*Node)(nil),                   // 5: prophetstor.datahub.resources.Node
	(*Cluster)(nil),                // 6: prophetstor.datahub.resources.Cluster
	(*ResourceRequirements)(nil),   // 7: prophetstor.datahub.resources.ResourceRequirements
	(*ContainerStatus)(nil),        // 8: prophetstor.datahub.resources.ContainerStatus
	(*ObjectMeta)(nil),             // 9: prophetstor.datahub.resources.ObjectMeta
	(*timestamp.Timestamp)(nil),    // 10: google.protobuf.Timestamp
	(*AlamedaPodSpec)(nil),         // 11: prophetstor.datahub.resources.AlamedaPodSpec
	(*PodStatus)(nil),              // 12: prophetstor.datahub.resources.PodStatus
	(Kind)(0),                      // 13: prophetstor.datahub.resources.Kind
	(*AlamedaControllerSpec)(nil),  // 14: prophetstor.datahub.resources.AlamedaControllerSpec
	(*AlamedaApplicationSpec)(nil), // 15: prophetstor.datahub.resources.AlamedaApplicationSpec
	(*Capacity)(nil),               // 16: prophetstor.datahub.resources.Capacity
	(*AlamedaNodeSpec)(nil),        // 17: prophetstor.datahub.resources.AlamedaNodeSpec
}
var file_datahub_resources_resources_proto_depIdxs = []int32{
	7,  // 0: prophetstor.datahub.resources.Container.resources:type_name -> prophetstor.datahub.resources.ResourceRequirements
	8,  // 1: prophetstor.datahub.resources.Container.status:type_name -> prophetstor.datahub.resources.ContainerStatus
	9,  // 2: prophetstor.datahub.resources.Pod.object_meta:type_name -> prophetstor.datahub.resources.ObjectMeta
	10, // 3: prophetstor.datahub.resources.Pod.start_time:type_name -> google.protobuf.Timestamp
	11, // 4: prophetstor.datahub.resources.Pod.alameda_pod_spec:type_name -> prophetstor.datahub.resources.AlamedaPodSpec
	2,  // 5: prophetstor.datahub.resources.Pod.top_controller:type_name -> prophetstor.datahub.resources.Controller
	12, // 6: prophetstor.datahub.resources.Pod.status:type_name -> prophetstor.datahub.resources.PodStatus
	0,  // 7: prophetstor.datahub.resources.Pod.containers:type_name -> prophetstor.datahub.resources.Container
	9,  // 8: prophetstor.datahub.resources.Controller.object_meta:type_name -> prophetstor.datahub.resources.ObjectMeta
	13, // 9: prophetstor.datahub.resources.Controller.kind:type_name -> prophetstor.datahub.resources.Kind
	14, // 10: prophetstor.datahub.resources.Controller.alameda_controller_spec:type_name -> prophetstor.datahub.resources.AlamedaControllerSpec
	9,  // 11: prophetstor.datahub.resources.Application.object_meta:type_name -> prophetstor.datahub.resources.ObjectMeta
	15, // 12: prophetstor.datahub.resources.Application.alameda_application_spec:type_name -> prophetstor.datahub.resources.AlamedaApplicationSpec
	2,  // 13: prophetstor.datahub.resources.Application.controllers:type_name -> prophetstor.datahub.resources.Controller
	9,  // 14: prophetstor.datahub.resources.Namespace.object_meta:type_name -> prophetstor.datahub.resources.ObjectMeta
	9,  // 15: prophetstor.datahub.resources.Node.object_meta:type_name -> prophetstor.datahub.resources.ObjectMeta
	10, // 16: prophetstor.datahub.resources.Node.start_time:type_name -> google.protobuf.Timestamp
	16, // 17: prophetstor.datahub.resources.Node.capacity:type_name -> prophetstor.datahub.resources.Capacity
	17, // 18: prophetstor.datahub.resources.Node.alameda_node_spec:type_name -> prophetstor.datahub.resources.AlamedaNodeSpec
	9,  // 19: prophetstor.datahub.resources.Cluster.object_meta:type_name -> prophetstor.datahub.resources.ObjectMeta
	20, // [20:20] is the sub-list for method output_type
	20, // [20:20] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_datahub_resources_resources_proto_init() }
func file_datahub_resources_resources_proto_init() {
	if File_datahub_resources_resources_proto != nil {
		return
	}
	file_datahub_resources_metadata_proto_init()
	file_datahub_resources_status_proto_init()
	file_datahub_resources_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_datahub_resources_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Container); i {
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
		file_datahub_resources_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pod); i {
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
		file_datahub_resources_resources_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Controller); i {
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
		file_datahub_resources_resources_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_datahub_resources_resources_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Namespace); i {
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
		file_datahub_resources_resources_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_datahub_resources_resources_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster); i {
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
			RawDescriptor: file_datahub_resources_resources_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datahub_resources_resources_proto_goTypes,
		DependencyIndexes: file_datahub_resources_resources_proto_depIdxs,
		MessageInfos:      file_datahub_resources_resources_proto_msgTypes,
	}.Build()
	File_datahub_resources_resources_proto = out.File
	file_datahub_resources_resources_proto_rawDesc = nil
	file_datahub_resources_resources_proto_goTypes = nil
	file_datahub_resources_resources_proto_depIdxs = nil
}