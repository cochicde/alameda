// This file has messages related to system events

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: datahub/events/types.proto

package events

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

type EventType int32

const (
	EventType_EVENT_TYPE_UNDEFINED                    EventType = 0
	EventType_EVENT_TYPE_ALAMEDA_SCALER_CREATE        EventType = 1
	EventType_EVENT_TYPE_ALAMEDA_SCALER_DELETE        EventType = 2
	EventType_EVENT_TYPE_NODE_REGISTER                EventType = 3
	EventType_EVENT_TYPE_DEPLOYMENT_REGISTER          EventType = 4
	EventType_EVENT_TYPE_DEPLOYMENT_CONFIG_REGISTER   EventType = 5
	EventType_EVENT_TYPE_POD_REGISTER                 EventType = 6
	EventType_EVENT_TYPE_NODE_DEREGISTER              EventType = 7
	EventType_EVENT_TYPE_DEPLOYMENT_DEREGISTER        EventType = 8
	EventType_EVENT_TYPE_DEPLOYMENT_CONFIG_DEREGISTER EventType = 9
	EventType_EVENT_TYPE_POD_DEREGISTER               EventType = 10
	EventType_EVENT_TYPE_NODE_PREDICTION_CREATE       EventType = 11
	EventType_EVENT_TYPE_POD_PREDICTION_CREATE        EventType = 12
	EventType_EVENT_TYPE_VPA_RECOMMENDATION_CREATE    EventType = 13
	EventType_EVENT_TYPE_HPA_RECOMMENDATION_CREATE    EventType = 14
	EventType_EVENT_TYPE_VPA_RECOMMENDATION_EXECUTE   EventType = 15
	EventType_EVENT_TYPE_HPA_RECOMMENDATION_EXECUTE   EventType = 16
	EventType_EVENT_TYPE_ANOMALY_METRIC_DETECT        EventType = 17
	EventType_EVENT_TYPE_ANOMALY_ANALYSIS_CREATE      EventType = 18
	EventType_EVENT_TYPE_LICENSE                      EventType = 19
	EventType_EVENT_TYPE_EMAIL_NOTIFICATION           EventType = 20
	EventType_EVENT_TYPE_ANOMALY_FORECAST_DETECT      EventType = 21
	EventType_EVENT_TYPE_ANOMALY_REALTIME_DETECT      EventType = 22
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0:  "EVENT_TYPE_UNDEFINED",
		1:  "EVENT_TYPE_ALAMEDA_SCALER_CREATE",
		2:  "EVENT_TYPE_ALAMEDA_SCALER_DELETE",
		3:  "EVENT_TYPE_NODE_REGISTER",
		4:  "EVENT_TYPE_DEPLOYMENT_REGISTER",
		5:  "EVENT_TYPE_DEPLOYMENT_CONFIG_REGISTER",
		6:  "EVENT_TYPE_POD_REGISTER",
		7:  "EVENT_TYPE_NODE_DEREGISTER",
		8:  "EVENT_TYPE_DEPLOYMENT_DEREGISTER",
		9:  "EVENT_TYPE_DEPLOYMENT_CONFIG_DEREGISTER",
		10: "EVENT_TYPE_POD_DEREGISTER",
		11: "EVENT_TYPE_NODE_PREDICTION_CREATE",
		12: "EVENT_TYPE_POD_PREDICTION_CREATE",
		13: "EVENT_TYPE_VPA_RECOMMENDATION_CREATE",
		14: "EVENT_TYPE_HPA_RECOMMENDATION_CREATE",
		15: "EVENT_TYPE_VPA_RECOMMENDATION_EXECUTE",
		16: "EVENT_TYPE_HPA_RECOMMENDATION_EXECUTE",
		17: "EVENT_TYPE_ANOMALY_METRIC_DETECT",
		18: "EVENT_TYPE_ANOMALY_ANALYSIS_CREATE",
		19: "EVENT_TYPE_LICENSE",
		20: "EVENT_TYPE_EMAIL_NOTIFICATION",
		21: "EVENT_TYPE_ANOMALY_FORECAST_DETECT",
		22: "EVENT_TYPE_ANOMALY_REALTIME_DETECT",
	}
	EventType_value = map[string]int32{
		"EVENT_TYPE_UNDEFINED":                    0,
		"EVENT_TYPE_ALAMEDA_SCALER_CREATE":        1,
		"EVENT_TYPE_ALAMEDA_SCALER_DELETE":        2,
		"EVENT_TYPE_NODE_REGISTER":                3,
		"EVENT_TYPE_DEPLOYMENT_REGISTER":          4,
		"EVENT_TYPE_DEPLOYMENT_CONFIG_REGISTER":   5,
		"EVENT_TYPE_POD_REGISTER":                 6,
		"EVENT_TYPE_NODE_DEREGISTER":              7,
		"EVENT_TYPE_DEPLOYMENT_DEREGISTER":        8,
		"EVENT_TYPE_DEPLOYMENT_CONFIG_DEREGISTER": 9,
		"EVENT_TYPE_POD_DEREGISTER":               10,
		"EVENT_TYPE_NODE_PREDICTION_CREATE":       11,
		"EVENT_TYPE_POD_PREDICTION_CREATE":        12,
		"EVENT_TYPE_VPA_RECOMMENDATION_CREATE":    13,
		"EVENT_TYPE_HPA_RECOMMENDATION_CREATE":    14,
		"EVENT_TYPE_VPA_RECOMMENDATION_EXECUTE":   15,
		"EVENT_TYPE_HPA_RECOMMENDATION_EXECUTE":   16,
		"EVENT_TYPE_ANOMALY_METRIC_DETECT":        17,
		"EVENT_TYPE_ANOMALY_ANALYSIS_CREATE":      18,
		"EVENT_TYPE_LICENSE":                      19,
		"EVENT_TYPE_EMAIL_NOTIFICATION":           20,
		"EVENT_TYPE_ANOMALY_FORECAST_DETECT":      21,
		"EVENT_TYPE_ANOMALY_REALTIME_DETECT":      22,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_datahub_events_types_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_datahub_events_types_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_datahub_events_types_proto_rawDescGZIP(), []int{0}
}

type EventVersion int32

const (
	EventVersion_EVENT_VERSION_UNDEFINED EventVersion = 0
	EventVersion_EVENT_VERSION_V1        EventVersion = 1
)

// Enum value maps for EventVersion.
var (
	EventVersion_name = map[int32]string{
		0: "EVENT_VERSION_UNDEFINED",
		1: "EVENT_VERSION_V1",
	}
	EventVersion_value = map[string]int32{
		"EVENT_VERSION_UNDEFINED": 0,
		"EVENT_VERSION_V1":        1,
	}
)

func (x EventVersion) Enum() *EventVersion {
	p := new(EventVersion)
	*p = x
	return p
}

func (x EventVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_datahub_events_types_proto_enumTypes[1].Descriptor()
}

func (EventVersion) Type() protoreflect.EnumType {
	return &file_datahub_events_types_proto_enumTypes[1]
}

func (x EventVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventVersion.Descriptor instead.
func (EventVersion) EnumDescriptor() ([]byte, []int) {
	return file_datahub_events_types_proto_rawDescGZIP(), []int{1}
}

type EventLevel int32

const (
	EventLevel_EVENT_LEVEL_UNDEFINED EventLevel = 0
	EventLevel_EVENT_LEVEL_DEBUG     EventLevel = 1
	EventLevel_EVENT_LEVEL_INFO      EventLevel = 2
	EventLevel_EVENT_LEVEL_WARNING   EventLevel = 3
	EventLevel_EVENT_LEVEL_ERROR     EventLevel = 4
	EventLevel_EVENT_LEVEL_FATAL     EventLevel = 5
)

// Enum value maps for EventLevel.
var (
	EventLevel_name = map[int32]string{
		0: "EVENT_LEVEL_UNDEFINED",
		1: "EVENT_LEVEL_DEBUG",
		2: "EVENT_LEVEL_INFO",
		3: "EVENT_LEVEL_WARNING",
		4: "EVENT_LEVEL_ERROR",
		5: "EVENT_LEVEL_FATAL",
	}
	EventLevel_value = map[string]int32{
		"EVENT_LEVEL_UNDEFINED": 0,
		"EVENT_LEVEL_DEBUG":     1,
		"EVENT_LEVEL_INFO":      2,
		"EVENT_LEVEL_WARNING":   3,
		"EVENT_LEVEL_ERROR":     4,
		"EVENT_LEVEL_FATAL":     5,
	}
)

func (x EventLevel) Enum() *EventLevel {
	p := new(EventLevel)
	*p = x
	return p
}

func (x EventLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_datahub_events_types_proto_enumTypes[2].Descriptor()
}

func (EventLevel) Type() protoreflect.EnumType {
	return &file_datahub_events_types_proto_enumTypes[2]
}

func (x EventLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventLevel.Descriptor instead.
func (EventLevel) EnumDescriptor() ([]byte, []int) {
	return file_datahub_events_types_proto_rawDescGZIP(), []int{2}
}

type EventSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host      string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Component string `protobuf:"bytes,2,opt,name=component,proto3" json:"component,omitempty"`
}

func (x *EventSource) Reset() {
	*x = EventSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_events_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventSource) ProtoMessage() {}

func (x *EventSource) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_events_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventSource.ProtoReflect.Descriptor instead.
func (*EventSource) Descriptor() ([]byte, []int) {
	return file_datahub_events_types_proto_rawDescGZIP(), []int{0}
}

func (x *EventSource) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *EventSource) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

type K8SObjectReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind       string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Namespace  string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ApiVersion string `protobuf:"bytes,4,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
}

func (x *K8SObjectReference) Reset() {
	*x = K8SObjectReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_events_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *K8SObjectReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*K8SObjectReference) ProtoMessage() {}

func (x *K8SObjectReference) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_events_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use K8SObjectReference.ProtoReflect.Descriptor instead.
func (*K8SObjectReference) Descriptor() ([]byte, []int) {
	return file_datahub_events_types_proto_rawDescGZIP(), []int{1}
}

func (x *K8SObjectReference) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *K8SObjectReference) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *K8SObjectReference) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *K8SObjectReference) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

var File_datahub_events_types_proto protoreflect.FileDescriptor

var file_datahub_events_types_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x70, 0x72,
	0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75,
	0x62, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x3f, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x22, 0x7b, 0x0a, 0x12, 0x4b, 0x38, 0x53,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2a, 0xdd, 0x06, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x24,
	0x0a, 0x20, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4c, 0x41,
	0x4d, 0x45, 0x44, 0x41, 0x5f, 0x53, 0x43, 0x41, 0x4c, 0x45, 0x52, 0x5f, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x10, 0x01, 0x12, 0x24, 0x0a, 0x20, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x41, 0x4c, 0x41, 0x4d, 0x45, 0x44, 0x41, 0x5f, 0x53, 0x43, 0x41, 0x4c, 0x45,
	0x52, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x52, 0x45,
	0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x10, 0x03, 0x12, 0x22, 0x0a, 0x1e, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x10, 0x04, 0x12, 0x29, 0x0a, 0x25,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x50, 0x4c, 0x4f,
	0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x52, 0x45, 0x47,
	0x49, 0x53, 0x54, 0x45, 0x52, 0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4f, 0x44, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54,
	0x45, 0x52, 0x10, 0x06, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x44, 0x45, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54,
	0x45, 0x52, 0x10, 0x07, 0x12, 0x24, 0x0a, 0x20, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x44, 0x45,
	0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x10, 0x08, 0x12, 0x2b, 0x0a, 0x27, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x44, 0x45, 0x52, 0x45, 0x47,
	0x49, 0x53, 0x54, 0x45, 0x52, 0x10, 0x09, 0x12, 0x1d, 0x0a, 0x19, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4f, 0x44, 0x5f, 0x44, 0x45, 0x52, 0x45, 0x47, 0x49,
	0x53, 0x54, 0x45, 0x52, 0x10, 0x0a, 0x12, 0x25, 0x0a, 0x21, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x50, 0x52, 0x45, 0x44, 0x49, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x0b, 0x12, 0x24, 0x0a,
	0x20, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4f, 0x44, 0x5f,
	0x50, 0x52, 0x45, 0x44, 0x49, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x10, 0x0c, 0x12, 0x28, 0x0a, 0x24, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x56, 0x50, 0x41, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x44, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x0d, 0x12, 0x28, 0x0a,
	0x24, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x50, 0x41, 0x5f,
	0x52, 0x45, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43,
	0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x0e, 0x12, 0x29, 0x0a, 0x25, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x50, 0x41, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x4d, 0x4d,
	0x45, 0x4e, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x45,
	0x10, 0x0f, 0x12, 0x29, 0x0a, 0x25, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x48, 0x50, 0x41, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x44, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x45, 0x10, 0x10, 0x12, 0x24, 0x0a,
	0x20, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4e, 0x4f, 0x4d,
	0x41, 0x4c, 0x59, 0x5f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x44, 0x45, 0x54, 0x45, 0x43,
	0x54, 0x10, 0x11, 0x12, 0x26, 0x0a, 0x22, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x41, 0x4e, 0x4f, 0x4d, 0x41, 0x4c, 0x59, 0x5f, 0x41, 0x4e, 0x41, 0x4c, 0x59, 0x53,
	0x49, 0x53, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x12, 0x12, 0x16, 0x0a, 0x12, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53,
	0x45, 0x10, 0x13, 0x12, 0x21, 0x0a, 0x1d, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x14, 0x12, 0x26, 0x0a, 0x22, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4e, 0x4f, 0x4d, 0x41, 0x4c, 0x59, 0x5f, 0x46, 0x4f, 0x52,
	0x45, 0x43, 0x41, 0x53, 0x54, 0x5f, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x10, 0x15, 0x12, 0x26,
	0x0a, 0x22, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4e, 0x4f,
	0x4d, 0x41, 0x4c, 0x59, 0x5f, 0x52, 0x45, 0x41, 0x4c, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x44, 0x45,
	0x54, 0x45, 0x43, 0x54, 0x10, 0x16, 0x2a, 0x41, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x56, 0x45, 0x52,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x31, 0x10, 0x01, 0x2a, 0x9b, 0x01, 0x0a, 0x0a, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4c, 0x45, 0x56,
	0x45, 0x4c, 0x5f, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x02,
	0x12, 0x17, 0x0a, 0x13, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f,
	0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04,
	0x12, 0x15, 0x0a, 0x11, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f,
	0x46, 0x41, 0x54, 0x41, 0x4c, 0x10, 0x05, 0x42, 0x24, 0x5a, 0x22, 0x70, 0x72, 0x6f, 0x70, 0x68,
	0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datahub_events_types_proto_rawDescOnce sync.Once
	file_datahub_events_types_proto_rawDescData = file_datahub_events_types_proto_rawDesc
)

func file_datahub_events_types_proto_rawDescGZIP() []byte {
	file_datahub_events_types_proto_rawDescOnce.Do(func() {
		file_datahub_events_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_datahub_events_types_proto_rawDescData)
	})
	return file_datahub_events_types_proto_rawDescData
}

var file_datahub_events_types_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_datahub_events_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_datahub_events_types_proto_goTypes = []interface{}{
	(EventType)(0),             // 0: prophetstor.datahub.events.EventType
	(EventVersion)(0),          // 1: prophetstor.datahub.events.EventVersion
	(EventLevel)(0),            // 2: prophetstor.datahub.events.EventLevel
	(*EventSource)(nil),        // 3: prophetstor.datahub.events.EventSource
	(*K8SObjectReference)(nil), // 4: prophetstor.datahub.events.K8SObjectReference
}
var file_datahub_events_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_datahub_events_types_proto_init() }
func file_datahub_events_types_proto_init() {
	if File_datahub_events_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datahub_events_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventSource); i {
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
		file_datahub_events_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*K8SObjectReference); i {
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
			RawDescriptor: file_datahub_events_types_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datahub_events_types_proto_goTypes,
		DependencyIndexes: file_datahub_events_types_proto_depIdxs,
		EnumInfos:         file_datahub_events_types_proto_enumTypes,
		MessageInfos:      file_datahub_events_types_proto_msgTypes,
	}.Build()
	File_datahub_events_types_proto = out.File
	file_datahub_events_types_proto_rawDesc = nil
	file_datahub_events_types_proto_goTypes = nil
	file_datahub_events_types_proto_depIdxs = nil
}
