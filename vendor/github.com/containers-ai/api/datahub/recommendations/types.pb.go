// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datahub/recommendations/types.proto

package recommendations

import (
	fmt "fmt"
	resources "github.com/containers-ai/api/datahub/resources"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ControllerRecommendationType int32

const (
	ControllerRecommendationType_CRT_Undefined ControllerRecommendationType = 0
	ControllerRecommendationType_CRT_Primitive ControllerRecommendationType = 1
	ControllerRecommendationType_CRT_K8S       ControllerRecommendationType = 2
)

var ControllerRecommendationType_name = map[int32]string{
	0: "CRT_Undefined",
	1: "CRT_Primitive",
	2: "CRT_K8S",
}

var ControllerRecommendationType_value = map[string]int32{
	"CRT_Undefined": 0,
	"CRT_Primitive": 1,
	"CRT_K8S":       2,
}

func (x ControllerRecommendationType) String() string {
	return proto.EnumName(ControllerRecommendationType_name, int32(x))
}

func (ControllerRecommendationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d6cae026dcf9e7d2, []int{0}
}

// Represents the priority of a node
type NodePriority struct {
	Nodes                []string `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodePriority) Reset()         { *m = NodePriority{} }
func (m *NodePriority) String() string { return proto.CompactTextString(m) }
func (*NodePriority) ProtoMessage()    {}
func (*NodePriority) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6cae026dcf9e7d2, []int{0}
}

func (m *NodePriority) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodePriority.Unmarshal(m, b)
}
func (m *NodePriority) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodePriority.Marshal(b, m, deterministic)
}
func (m *NodePriority) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodePriority.Merge(m, src)
}
func (m *NodePriority) XXX_Size() int {
	return xxx_messageInfo_NodePriority.Size(m)
}
func (m *NodePriority) XXX_DiscardUnknown() {
	xxx_messageInfo_NodePriority.DiscardUnknown(m)
}

var xxx_messageInfo_NodePriority proto.InternalMessageInfo

func (m *NodePriority) GetNodes() []string {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type Selector struct {
	Selector             map[string]string `protobuf:"bytes,1,rep,name=selector,proto3" json:"selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Selector) Reset()         { *m = Selector{} }
func (m *Selector) String() string { return proto.CompactTextString(m) }
func (*Selector) ProtoMessage()    {}
func (*Selector) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6cae026dcf9e7d2, []int{1}
}

func (m *Selector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Selector.Unmarshal(m, b)
}
func (m *Selector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Selector.Marshal(b, m, deterministic)
}
func (m *Selector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Selector.Merge(m, src)
}
func (m *Selector) XXX_Size() int {
	return xxx_messageInfo_Selector.Size(m)
}
func (m *Selector) XXX_DiscardUnknown() {
	xxx_messageInfo_Selector.DiscardUnknown(m)
}

var xxx_messageInfo_Selector proto.InternalMessageInfo

func (m *Selector) GetSelector() map[string]string {
	if m != nil {
		return m.Selector
	}
	return nil
}

// Represents a recommended pod-to-node assignment (i.e. pod placement)
type AssignPodPolicy struct {
	Time *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	// Types that are valid to be assigned to Policy:
	//	*AssignPodPolicy_NodePriority
	//	*AssignPodPolicy_NodeSelector
	//	*AssignPodPolicy_NodeName
	Policy               isAssignPodPolicy_Policy `protobuf_oneof:"policy"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *AssignPodPolicy) Reset()         { *m = AssignPodPolicy{} }
func (m *AssignPodPolicy) String() string { return proto.CompactTextString(m) }
func (*AssignPodPolicy) ProtoMessage()    {}
func (*AssignPodPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6cae026dcf9e7d2, []int{2}
}

func (m *AssignPodPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignPodPolicy.Unmarshal(m, b)
}
func (m *AssignPodPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignPodPolicy.Marshal(b, m, deterministic)
}
func (m *AssignPodPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignPodPolicy.Merge(m, src)
}
func (m *AssignPodPolicy) XXX_Size() int {
	return xxx_messageInfo_AssignPodPolicy.Size(m)
}
func (m *AssignPodPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignPodPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_AssignPodPolicy proto.InternalMessageInfo

func (m *AssignPodPolicy) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
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

func (m *AssignPodPolicy) GetPolicy() isAssignPodPolicy_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (m *AssignPodPolicy) GetNodePriority() *NodePriority {
	if x, ok := m.GetPolicy().(*AssignPodPolicy_NodePriority); ok {
		return x.NodePriority
	}
	return nil
}

func (m *AssignPodPolicy) GetNodeSelector() *Selector {
	if x, ok := m.GetPolicy().(*AssignPodPolicy_NodeSelector); ok {
		return x.NodeSelector
	}
	return nil
}

func (m *AssignPodPolicy) GetNodeName() string {
	if x, ok := m.GetPolicy().(*AssignPodPolicy_NodeName); ok {
		return x.NodeName
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AssignPodPolicy) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AssignPodPolicy_NodePriority)(nil),
		(*AssignPodPolicy_NodeSelector)(nil),
		(*AssignPodPolicy_NodeName)(nil),
	}
}

type ControllerRecommendationSpec struct {
	NamespacedName       *resources.NamespacedName `protobuf:"bytes,1,opt,name=namespaced_name,json=namespacedName,proto3" json:"namespaced_name,omitempty"`
	Kind                 resources.Kind            `protobuf:"varint,2,opt,name=kind,proto3,enum=containersai.datahub.resources.Kind" json:"kind,omitempty"`
	CurrentReplicas      int32                     `protobuf:"varint,3,opt,name=current_replicas,json=currentReplicas,proto3" json:"current_replicas,omitempty"`
	DesiredReplicas      int32                     `protobuf:"varint,4,opt,name=desired_replicas,json=desiredReplicas,proto3" json:"desired_replicas,omitempty"`
	CreateTime           *timestamp.Timestamp      `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ControllerRecommendationSpec) Reset()         { *m = ControllerRecommendationSpec{} }
func (m *ControllerRecommendationSpec) String() string { return proto.CompactTextString(m) }
func (*ControllerRecommendationSpec) ProtoMessage()    {}
func (*ControllerRecommendationSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6cae026dcf9e7d2, []int{3}
}

func (m *ControllerRecommendationSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControllerRecommendationSpec.Unmarshal(m, b)
}
func (m *ControllerRecommendationSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControllerRecommendationSpec.Marshal(b, m, deterministic)
}
func (m *ControllerRecommendationSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControllerRecommendationSpec.Merge(m, src)
}
func (m *ControllerRecommendationSpec) XXX_Size() int {
	return xxx_messageInfo_ControllerRecommendationSpec.Size(m)
}
func (m *ControllerRecommendationSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ControllerRecommendationSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ControllerRecommendationSpec proto.InternalMessageInfo

func (m *ControllerRecommendationSpec) GetNamespacedName() *resources.NamespacedName {
	if m != nil {
		return m.NamespacedName
	}
	return nil
}

func (m *ControllerRecommendationSpec) GetKind() resources.Kind {
	if m != nil {
		return m.Kind
	}
	return resources.Kind_KIND_UNDEFINED
}

func (m *ControllerRecommendationSpec) GetCurrentReplicas() int32 {
	if m != nil {
		return m.CurrentReplicas
	}
	return 0
}

func (m *ControllerRecommendationSpec) GetDesiredReplicas() int32 {
	if m != nil {
		return m.DesiredReplicas
	}
	return 0
}

func (m *ControllerRecommendationSpec) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func init() {
	proto.RegisterEnum("containersai.datahub.recommendations.ControllerRecommendationType", ControllerRecommendationType_name, ControllerRecommendationType_value)
	proto.RegisterType((*NodePriority)(nil), "containersai.datahub.recommendations.NodePriority")
	proto.RegisterType((*Selector)(nil), "containersai.datahub.recommendations.Selector")
	proto.RegisterMapType((map[string]string)(nil), "containersai.datahub.recommendations.Selector.SelectorEntry")
	proto.RegisterType((*AssignPodPolicy)(nil), "containersai.datahub.recommendations.AssignPodPolicy")
	proto.RegisterType((*ControllerRecommendationSpec)(nil), "containersai.datahub.recommendations.ControllerRecommendationSpec")
}

func init() {
	proto.RegisterFile("datahub/recommendations/types.proto", fileDescriptor_d6cae026dcf9e7d2)
}

var fileDescriptor_d6cae026dcf9e7d2 = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x5d, 0x6b, 0xdb, 0x3e,
	0x14, 0xc6, 0xeb, 0xbc, 0xf4, 0x9f, 0x28, 0x4d, 0x93, 0xbf, 0xd9, 0x45, 0x08, 0x2b, 0x2b, 0x59,
	0x2e, 0xba, 0xc1, 0x64, 0xc8, 0xc6, 0x08, 0xeb, 0x6e, 0xd6, 0x32, 0x28, 0x14, 0x42, 0xa6, 0xa4,
	0xec, 0xe5, 0x26, 0x28, 0xd2, 0x69, 0x26, 0x6a, 0x4b, 0x46, 0x52, 0x0a, 0xfe, 0x2e, 0xbb, 0xd9,
	0x77, 0xd9, 0x07, 0x1b, 0x92, 0x1d, 0xa7, 0x19, 0x74, 0xed, 0xee, 0xce, 0x79, 0xfc, 0x9c, 0xc7,
	0x47, 0x3f, 0x0b, 0xa3, 0xe7, 0x9c, 0x5a, 0xfa, 0x7d, 0xbd, 0x8c, 0x34, 0x30, 0x95, 0x24, 0x20,
	0x39, 0xb5, 0x42, 0x49, 0x13, 0xd9, 0x2c, 0x05, 0x83, 0x53, 0xad, 0xac, 0x0a, 0x87, 0x4c, 0x49,
	0x4b, 0x85, 0x04, 0x6d, 0xa8, 0xc0, 0xc5, 0x04, 0xfe, 0x63, 0xa2, 0x7f, 0xb4, 0x8d, 0x32, 0x6a,
	0xad, 0x19, 0xec, 0x84, 0xf4, 0x9f, 0xad, 0x94, 0x5a, 0xc5, 0x10, 0xf9, 0x6e, 0xb9, 0xbe, 0x8e,
	0xac, 0x48, 0xc0, 0x58, 0x9a, 0xa4, 0xb9, 0x61, 0x30, 0x44, 0x07, 0x13, 0xc5, 0x61, 0xaa, 0x85,
	0xd2, 0xc2, 0x66, 0xe1, 0x13, 0x54, 0x97, 0x8a, 0x83, 0xe9, 0x05, 0xc7, 0xd5, 0x93, 0x26, 0xc9,
	0x9b, 0xc1, 0xcf, 0x00, 0x35, 0x66, 0x10, 0x03, 0xb3, 0x4a, 0x87, 0x5f, 0x50, 0xc3, 0x14, 0xb5,
	0x77, 0xb5, 0x46, 0xef, 0xf1, 0x63, 0x76, 0xc5, 0x9b, 0x84, 0xb2, 0xf8, 0x28, 0xad, 0xce, 0x48,
	0x99, 0xd6, 0x3f, 0x45, 0xed, 0x9d, 0x47, 0x61, 0x17, 0x55, 0x6f, 0x20, 0xeb, 0x05, 0xc7, 0xc1,
	0x49, 0x93, 0xb8, 0xd2, 0xed, 0x77, 0x4b, 0xe3, 0x35, 0xf4, 0x2a, 0x5e, 0xcb, 0x9b, 0x77, 0x95,
	0x71, 0x30, 0xf8, 0x51, 0x41, 0x9d, 0x0f, 0xc6, 0x88, 0x95, 0x9c, 0x2a, 0x3e, 0x55, 0xb1, 0x60,
	0x59, 0x88, 0x51, 0xcd, 0x1d, 0xd8, 0x07, 0xb4, 0x46, 0x7d, 0x9c, 0xd3, 0xc0, 0x1b, 0x1a, 0x78,
	0xbe, 0xa1, 0x41, 0xbc, 0x2f, 0xfc, 0x8a, 0xda, 0xee, 0xc0, 0x8b, 0xb4, 0xc0, 0xe1, 0xdf, 0xd2,
	0x1a, 0x8d, 0x1e, 0x77, 0xbe, 0xbb, 0x20, 0x2f, 0xf6, 0xc8, 0x81, 0xbc, 0x0b, 0xf6, 0xaa, 0x88,
	0x2e, 0xd1, 0x55, 0x7d, 0x34, 0xfe, 0x37, 0x74, 0x9b, 0xd8, 0xf2, 0x63, 0x1c, 0xa1, 0xa6, 0x8f,
	0x95, 0x34, 0x81, 0x5e, 0xcd, 0x31, 0xb9, 0xd8, 0x23, 0x0d, 0x27, 0x4d, 0x68, 0x02, 0x67, 0x0d,
	0xb4, 0x9f, 0x7a, 0x14, 0x83, 0x5f, 0x15, 0xf4, 0xf4, 0x5c, 0x49, 0xab, 0x55, 0x1c, 0x83, 0x26,
	0x3b, 0xf9, 0xb3, 0x14, 0x58, 0xf8, 0x19, 0x75, 0x5c, 0x88, 0x49, 0x29, 0x03, 0x9e, 0xe7, 0x05,
	0x7f, 0x5f, 0xb1, 0xb8, 0x70, 0x78, 0x52, 0x8e, 0xb9, 0x8a, 0x1c, 0xca, 0x9d, 0x3e, 0x1c, 0xa3,
	0xda, 0x8d, 0x90, 0xdc, 0xb3, 0x3c, 0x1c, 0x0d, 0x1f, 0x4a, 0xbb, 0x14, 0x92, 0x13, 0x3f, 0x11,
	0xbe, 0x40, 0x5d, 0xb6, 0xd6, 0x1a, 0xa4, 0x5d, 0x68, 0x48, 0x63, 0xc1, 0xa8, 0xf1, 0xd8, 0xea,
	0xa4, 0x53, 0xe8, 0xa4, 0x90, 0x9d, 0x95, 0x83, 0x11, 0x1a, 0xf8, 0xd6, 0x5a, 0xcb, 0xad, 0x85,
	0x5e, 0x5a, 0x4f, 0x51, 0x8b, 0x69, 0xa0, 0x16, 0x16, 0xfe, 0x6e, 0xd4, 0x1f, 0xbc, 0x1b, 0x28,
	0xb7, 0x3b, 0xe1, 0xe5, 0xa7, 0xfb, 0x29, 0xce, 0xb3, 0x14, 0xc2, 0xff, 0x51, 0xfb, 0x9c, 0xcc,
	0x17, 0x57, 0x92, 0xc3, 0xb5, 0x90, 0xc0, 0xbb, 0x7b, 0x1b, 0x69, 0xaa, 0x45, 0x22, 0xac, 0xb8,
	0x85, 0x6e, 0x10, 0xb6, 0xd0, 0x7f, 0x4e, 0xba, 0x1c, 0xcf, 0xba, 0x95, 0xb3, 0xb7, 0xdf, 0xde,
	0xac, 0x84, 0x75, 0x10, 0x98, 0x4a, 0xa2, 0x2d, 0x9d, 0x57, 0x54, 0x44, 0x34, 0x15, 0xd1, 0x3d,
	0x3f, 0x8b, 0xe5, 0xbe, 0x5f, 0xf5, 0xf5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0x29, 0x87,
	0xde, 0x4e, 0x04, 0x00, 0x00,
}