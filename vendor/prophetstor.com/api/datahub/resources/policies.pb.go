// This file has messages related general definitions

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: datahub/resources/policies.proto

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
// Recommendation policy. A policy may be either stable or compact.
type RecommendationPolicy int32

const (
	RecommendationPolicy_RECOMMENDATION_POLICY_UNDEFINED RecommendationPolicy = 0
	RecommendationPolicy_STABLE                          RecommendationPolicy = 1
	RecommendationPolicy_COMPACT                         RecommendationPolicy = 2
)

// Enum value maps for RecommendationPolicy.
var (
	RecommendationPolicy_name = map[int32]string{
		0: "RECOMMENDATION_POLICY_UNDEFINED",
		1: "STABLE",
		2: "COMPACT",
	}
	RecommendationPolicy_value = map[string]int32{
		"RECOMMENDATION_POLICY_UNDEFINED": 0,
		"STABLE":                          1,
		"COMPACT":                         2,
	}
)

func (x RecommendationPolicy) Enum() *RecommendationPolicy {
	p := new(RecommendationPolicy)
	*p = x
	return p
}

func (x RecommendationPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RecommendationPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_datahub_resources_policies_proto_enumTypes[0].Descriptor()
}

func (RecommendationPolicy) Type() protoreflect.EnumType {
	return &file_datahub_resources_policies_proto_enumTypes[0]
}

func (x RecommendationPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecommendationPolicy.Descriptor instead.
func (RecommendationPolicy) EnumDescriptor() ([]byte, []int) {
	return file_datahub_resources_policies_proto_rawDescGZIP(), []int{0}
}

//*
// Represents the priority of a node.
type NodePriority struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []string `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *NodePriority) Reset() {
	*x = NodePriority{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_resources_policies_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodePriority) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodePriority) ProtoMessage() {}

func (x *NodePriority) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_resources_policies_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodePriority.ProtoReflect.Descriptor instead.
func (*NodePriority) Descriptor() ([]byte, []int) {
	return file_datahub_resources_policies_proto_rawDescGZIP(), []int{0}
}

func (x *NodePriority) GetNodes() []string {
	if x != nil {
		return x.Nodes
	}
	return nil
}

//*
// Represents a Kubernetes label selector.
type Selector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Selector map[string]string `protobuf:"bytes,1,rep,name=selector,proto3" json:"selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Selector) Reset() {
	*x = Selector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_resources_policies_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Selector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Selector) ProtoMessage() {}

func (x *Selector) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_resources_policies_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Selector.ProtoReflect.Descriptor instead.
func (*Selector) Descriptor() ([]byte, []int) {
	return file_datahub_resources_policies_proto_rawDescGZIP(), []int{1}
}

func (x *Selector) GetSelector() map[string]string {
	if x != nil {
		return x.Selector
	}
	return nil
}

//*
// Represents a recommended pod-to-node assignment. (i.e. pod placement)
type AssignPodPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	// Types that are assignable to Policy:
	//	*AssignPodPolicy_NodePriority
	//	*AssignPodPolicy_NodeSelector
	//	*AssignPodPolicy_NodeName
	Policy isAssignPodPolicy_Policy `protobuf_oneof:"policy"`
}

func (x *AssignPodPolicy) Reset() {
	*x = AssignPodPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_resources_policies_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignPodPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignPodPolicy) ProtoMessage() {}

func (x *AssignPodPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_resources_policies_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignPodPolicy.ProtoReflect.Descriptor instead.
func (*AssignPodPolicy) Descriptor() ([]byte, []int) {
	return file_datahub_resources_policies_proto_rawDescGZIP(), []int{2}
}

func (x *AssignPodPolicy) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (m *AssignPodPolicy) GetPolicy() isAssignPodPolicy_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (x *AssignPodPolicy) GetNodePriority() *NodePriority {
	if x, ok := x.GetPolicy().(*AssignPodPolicy_NodePriority); ok {
		return x.NodePriority
	}
	return nil
}

func (x *AssignPodPolicy) GetNodeSelector() *Selector {
	if x, ok := x.GetPolicy().(*AssignPodPolicy_NodeSelector); ok {
		return x.NodeSelector
	}
	return nil
}

func (x *AssignPodPolicy) GetNodeName() string {
	if x, ok := x.GetPolicy().(*AssignPodPolicy_NodeName); ok {
		return x.NodeName
	}
	return ""
}

type isAssignPodPolicy_Policy interface {
	isAssignPodPolicy_Policy()
}

type AssignPodPolicy_NodePriority struct {
	NodePriority *NodePriority `protobuf:"bytes,2,opt,name=node_priority,json=nodePriority,proto3,oneof"`
}

type AssignPodPolicy_NodeSelector struct {
	NodeSelector *Selector `protobuf:"bytes,3,opt,name=node_selector,json=nodeSelector,proto3,oneof"`
}

type AssignPodPolicy_NodeName struct {
	NodeName string `protobuf:"bytes,4,opt,name=node_name,json=nodeName,proto3,oneof"`
}

func (*AssignPodPolicy_NodePriority) isAssignPodPolicy_Policy() {}

func (*AssignPodPolicy_NodeSelector) isAssignPodPolicy_Policy() {}

func (*AssignPodPolicy_NodeName) isAssignPodPolicy_Policy() {}

var File_datahub_resources_policies_proto protoreflect.FileDescriptor

var file_datahub_resources_policies_proto_rawDesc = []byte{
	0x0a, 0x20, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1d, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0c, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x9a, 0x01, 0x0a, 0x08, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x51, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65,
	0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x1a, 0x3b, 0x0a, 0x0d, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8e, 0x02, 0x0a, 0x0f, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x50, 0x6f, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x48, 0x00, 0x52,
	0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x4e, 0x0a,
	0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x48, 0x00, 0x52,
	0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1d, 0x0a,
	0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2a, 0x54, 0x0a, 0x14, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x23,
	0x0a, 0x1f, 0x52, 0x45, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x43, 0x54, 0x10, 0x02, 0x42, 0x27, 0x5a, 0x25,
	0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datahub_resources_policies_proto_rawDescOnce sync.Once
	file_datahub_resources_policies_proto_rawDescData = file_datahub_resources_policies_proto_rawDesc
)

func file_datahub_resources_policies_proto_rawDescGZIP() []byte {
	file_datahub_resources_policies_proto_rawDescOnce.Do(func() {
		file_datahub_resources_policies_proto_rawDescData = protoimpl.X.CompressGZIP(file_datahub_resources_policies_proto_rawDescData)
	})
	return file_datahub_resources_policies_proto_rawDescData
}

var file_datahub_resources_policies_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_datahub_resources_policies_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_datahub_resources_policies_proto_goTypes = []interface{}{
	(RecommendationPolicy)(0),   // 0: prophetstor.datahub.resources.RecommendationPolicy
	(*NodePriority)(nil),        // 1: prophetstor.datahub.resources.NodePriority
	(*Selector)(nil),            // 2: prophetstor.datahub.resources.Selector
	(*AssignPodPolicy)(nil),     // 3: prophetstor.datahub.resources.AssignPodPolicy
	nil,                         // 4: prophetstor.datahub.resources.Selector.SelectorEntry
	(*timestamp.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_datahub_resources_policies_proto_depIdxs = []int32{
	4, // 0: prophetstor.datahub.resources.Selector.selector:type_name -> prophetstor.datahub.resources.Selector.SelectorEntry
	5, // 1: prophetstor.datahub.resources.AssignPodPolicy.time:type_name -> google.protobuf.Timestamp
	1, // 2: prophetstor.datahub.resources.AssignPodPolicy.node_priority:type_name -> prophetstor.datahub.resources.NodePriority
	2, // 3: prophetstor.datahub.resources.AssignPodPolicy.node_selector:type_name -> prophetstor.datahub.resources.Selector
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_datahub_resources_policies_proto_init() }
func file_datahub_resources_policies_proto_init() {
	if File_datahub_resources_policies_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datahub_resources_policies_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodePriority); i {
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
		file_datahub_resources_policies_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Selector); i {
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
		file_datahub_resources_policies_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignPodPolicy); i {
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
	file_datahub_resources_policies_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*AssignPodPolicy_NodePriority)(nil),
		(*AssignPodPolicy_NodeSelector)(nil),
		(*AssignPodPolicy_NodeName)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_datahub_resources_policies_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datahub_resources_policies_proto_goTypes,
		DependencyIndexes: file_datahub_resources_policies_proto_depIdxs,
		EnumInfos:         file_datahub_resources_policies_proto_enumTypes,
		MessageInfos:      file_datahub_resources_policies_proto_msgTypes,
	}.Build()
	File_datahub_resources_policies_proto = out.File
	file_datahub_resources_policies_proto_rawDesc = nil
	file_datahub_resources_policies_proto_goTypes = nil
	file_datahub_resources_policies_proto_depIdxs = nil
}