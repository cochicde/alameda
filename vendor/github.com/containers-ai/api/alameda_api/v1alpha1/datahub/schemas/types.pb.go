// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/schemas/types.proto

package schemas

import (
	fmt "fmt"
	common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	proto "github.com/golang/protobuf/proto"
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

type Scope int32

const (
	Scope_SCOPE_UNDEFINED      Scope = 0
	Scope_SCOPE_APPLICATION    Scope = 1
	Scope_SCOPE_METRIC         Scope = 2
	Scope_SCOPE_PLANNING       Scope = 3
	Scope_SCOPE_PREDICTION     Scope = 4
	Scope_SCOPE_RECOMMENDATION Scope = 5
	Scope_SCOPE_RESOURCE       Scope = 6
)

var Scope_name = map[int32]string{
	0: "SCOPE_UNDEFINED",
	1: "SCOPE_APPLICATION",
	2: "SCOPE_METRIC",
	3: "SCOPE_PLANNING",
	4: "SCOPE_PREDICTION",
	5: "SCOPE_RECOMMENDATION",
	6: "SCOPE_RESOURCE",
}

var Scope_value = map[string]int32{
	"SCOPE_UNDEFINED":      0,
	"SCOPE_APPLICATION":    1,
	"SCOPE_METRIC":         2,
	"SCOPE_PLANNING":       3,
	"SCOPE_PREDICTION":     4,
	"SCOPE_RECOMMENDATION": 5,
	"SCOPE_RESOURCE":       6,
}

func (x Scope) String() string {
	return proto.EnumName(Scope_name, int32(x))
}

func (Scope) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_01194bf23370964f, []int{0}
}

type SchemaMeta struct {
	Scope                Scope    `protobuf:"varint,1,opt,name=scope,proto3,enum=containersai.alameda.v1alpha1.datahub.schemas.Scope" json:"scope,omitempty"`
	Category             string   `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SchemaMeta) Reset()         { *m = SchemaMeta{} }
func (m *SchemaMeta) String() string { return proto.CompactTextString(m) }
func (*SchemaMeta) ProtoMessage()    {}
func (*SchemaMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_01194bf23370964f, []int{0}
}

func (m *SchemaMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchemaMeta.Unmarshal(m, b)
}
func (m *SchemaMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchemaMeta.Marshal(b, m, deterministic)
}
func (m *SchemaMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchemaMeta.Merge(m, src)
}
func (m *SchemaMeta) XXX_Size() int {
	return xxx_messageInfo_SchemaMeta.Size(m)
}
func (m *SchemaMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_SchemaMeta.DiscardUnknown(m)
}

var xxx_messageInfo_SchemaMeta proto.InternalMessageInfo

func (m *SchemaMeta) GetScope() Scope {
	if m != nil {
		return m.Scope
	}
	return Scope_SCOPE_UNDEFINED
}

func (m *SchemaMeta) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *SchemaMeta) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type Measurement struct {
	Name                 string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MetricType           common.MetricType       `protobuf:"varint,2,opt,name=metric_type,json=metricType,proto3,enum=containersai.alameda.v1alpha1.datahub.common.MetricType" json:"metric_type,omitempty"`
	ResourceBoundary     common.ResourceBoundary `protobuf:"varint,3,opt,name=resource_boundary,json=resourceBoundary,proto3,enum=containersai.alameda.v1alpha1.datahub.common.ResourceBoundary" json:"resource_boundary,omitempty"`
	ResourceQuota        common.ResourceQuota    `protobuf:"varint,4,opt,name=resource_quota,json=resourceQuota,proto3,enum=containersai.alameda.v1alpha1.datahub.common.ResourceQuota" json:"resource_quota,omitempty"`
	Columns              []*Column               `protobuf:"bytes,5,rep,name=columns,proto3" json:"columns,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Measurement) Reset()         { *m = Measurement{} }
func (m *Measurement) String() string { return proto.CompactTextString(m) }
func (*Measurement) ProtoMessage()    {}
func (*Measurement) Descriptor() ([]byte, []int) {
	return fileDescriptor_01194bf23370964f, []int{1}
}

func (m *Measurement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Measurement.Unmarshal(m, b)
}
func (m *Measurement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Measurement.Marshal(b, m, deterministic)
}
func (m *Measurement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Measurement.Merge(m, src)
}
func (m *Measurement) XXX_Size() int {
	return xxx_messageInfo_Measurement.Size(m)
}
func (m *Measurement) XXX_DiscardUnknown() {
	xxx_messageInfo_Measurement.DiscardUnknown(m)
}

var xxx_messageInfo_Measurement proto.InternalMessageInfo

func (m *Measurement) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Measurement) GetMetricType() common.MetricType {
	if m != nil {
		return m.MetricType
	}
	return common.MetricType_METRICS_TYPE_UNDEFINED
}

func (m *Measurement) GetResourceBoundary() common.ResourceBoundary {
	if m != nil {
		return m.ResourceBoundary
	}
	return common.ResourceBoundary_RESOURCE_BOUNDARY_UNDEFINED
}

func (m *Measurement) GetResourceQuota() common.ResourceQuota {
	if m != nil {
		return m.ResourceQuota
	}
	return common.ResourceQuota_RESOURCE_QUOTA_UNDEFINED
}

func (m *Measurement) GetColumns() []*Column {
	if m != nil {
		return m.Columns
	}
	return nil
}

type Column struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Required             bool              `protobuf:"varint,2,opt,name=required,proto3" json:"required,omitempty"`
	ColumnType           common.ColumnType `protobuf:"varint,3,opt,name=column_type,json=columnType,proto3,enum=containersai.alameda.v1alpha1.datahub.common.ColumnType" json:"column_type,omitempty"`
	DataType             common.DataType   `protobuf:"varint,4,opt,name=data_type,json=dataType,proto3,enum=containersai.alameda.v1alpha1.datahub.common.DataType" json:"data_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Column) Reset()         { *m = Column{} }
func (m *Column) String() string { return proto.CompactTextString(m) }
func (*Column) ProtoMessage()    {}
func (*Column) Descriptor() ([]byte, []int) {
	return fileDescriptor_01194bf23370964f, []int{2}
}

func (m *Column) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Column.Unmarshal(m, b)
}
func (m *Column) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Column.Marshal(b, m, deterministic)
}
func (m *Column) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Column.Merge(m, src)
}
func (m *Column) XXX_Size() int {
	return xxx_messageInfo_Column.Size(m)
}
func (m *Column) XXX_DiscardUnknown() {
	xxx_messageInfo_Column.DiscardUnknown(m)
}

var xxx_messageInfo_Column proto.InternalMessageInfo

func (m *Column) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Column) GetRequired() bool {
	if m != nil {
		return m.Required
	}
	return false
}

func (m *Column) GetColumnType() common.ColumnType {
	if m != nil {
		return m.ColumnType
	}
	return common.ColumnType_COLUMNTYPE_UDEFINED
}

func (m *Column) GetDataType() common.DataType {
	if m != nil {
		return m.DataType
	}
	return common.DataType_DATATYPE_UNDEFINED
}

func init() {
	proto.RegisterEnum("containersai.alameda.v1alpha1.datahub.schemas.Scope", Scope_name, Scope_value)
	proto.RegisterType((*SchemaMeta)(nil), "containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta")
	proto.RegisterType((*Measurement)(nil), "containersai.alameda.v1alpha1.datahub.schemas.Measurement")
	proto.RegisterType((*Column)(nil), "containersai.alameda.v1alpha1.datahub.schemas.Column")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/schemas/types.proto", fileDescriptor_01194bf23370964f)
}

var fileDescriptor_01194bf23370964f = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdb, 0x8a, 0xd3, 0x40,
	0x18, 0x36, 0x3d, 0xd9, 0xfe, 0xd5, 0x9a, 0x1d, 0x57, 0x08, 0xbd, 0x2a, 0xbd, 0x2a, 0xc2, 0x26,
	0xb6, 0x1e, 0x10, 0x04, 0xa1, 0x9b, 0x46, 0x89, 0x6c, 0xd2, 0x3a, 0xed, 0x5e, 0xe8, 0x4d, 0x99,
	0xa6, 0xc3, 0x36, 0xd8, 0x64, 0xb2, 0x93, 0x89, 0xd0, 0x37, 0xf0, 0x25, 0xf4, 0xf9, 0x7c, 0x0c,
	0xc9, 0x4c, 0x36, 0xbb, 0x88, 0x48, 0xb3, 0x77, 0xff, 0x7c, 0xcc, 0x77, 0xf8, 0xfb, 0x35, 0x03,
	0x2f, 0xc8, 0x9e, 0x44, 0x74, 0x4b, 0xd6, 0x24, 0x09, 0xad, 0xef, 0x63, 0xb2, 0x4f, 0x76, 0x64,
	0x6c, 0x6d, 0x89, 0x20, 0xbb, 0x6c, 0x63, 0xa5, 0xc1, 0x8e, 0x46, 0x24, 0xb5, 0xc4, 0x21, 0xa1,
	0xa9, 0x99, 0x70, 0x26, 0x18, 0x3a, 0x0b, 0x58, 0x2c, 0x48, 0x18, 0x53, 0x9e, 0x92, 0xd0, 0x2c,
	0xe8, 0xe6, 0x0d, 0xd5, 0x2c, 0xa8, 0x66, 0x41, 0xed, 0x8f, 0xff, 0x6b, 0x10, 0xb0, 0x28, 0x62,
	0xb1, 0x15, 0x51, 0xc1, 0xc3, 0xa0, 0x70, 0xe8, 0x5b, 0xc7, 0x50, 0xee, 0x44, 0x1a, 0xfe, 0xd0,
	0x00, 0x96, 0xd2, 0xcf, 0xa3, 0x82, 0xa0, 0x4f, 0xd0, 0x4c, 0x03, 0x96, 0x50, 0x43, 0x1b, 0x68,
	0xa3, 0xde, 0xe4, 0x95, 0x59, 0x29, 0xb1, 0xb9, 0xcc, 0xb9, 0x58, 0x49, 0xa0, 0x3e, 0xb4, 0x03,
	0x22, 0xe8, 0x15, 0xe3, 0x07, 0xa3, 0x36, 0xd0, 0x46, 0x1d, 0x5c, 0x9e, 0x11, 0x82, 0x46, 0x9e,
	0xc2, 0xa8, 0x4b, 0x5c, 0xce, 0xc3, 0x5f, 0x75, 0xe8, 0x7a, 0x94, 0xa4, 0x19, 0xa7, 0x11, 0x8d,
	0x45, 0x7e, 0x27, 0x26, 0x91, 0x8a, 0xd2, 0xc1, 0x72, 0x46, 0x5f, 0xa0, 0xab, 0x16, 0x5e, 0x4b,
	0x7a, 0x4d, 0xa6, 0x7c, 0x7b, 0x64, 0x4a, 0xb5, 0xbe, 0xe9, 0x49, 0x81, 0xd5, 0x21, 0xa1, 0x18,
	0xa2, 0x72, 0x46, 0xdf, 0xe0, 0x84, 0xd3, 0x94, 0x65, 0x3c, 0xa0, 0xeb, 0x0d, 0xcb, 0xe2, 0x2d,
	0xe1, 0x07, 0x99, 0xaf, 0x37, 0x79, 0x5f, 0xcd, 0x00, 0x17, 0x32, 0xe7, 0x85, 0x0a, 0xd6, 0xf9,
	0x5f, 0x08, 0xda, 0x40, 0xaf, 0x34, 0xbb, 0xce, 0x98, 0x20, 0x46, 0x43, 0x3a, 0xbd, 0xbb, 0x9f,
	0xd3, 0xe7, 0x5c, 0x02, 0x3f, 0xe6, 0x77, 0x8f, 0x68, 0x0e, 0x0f, 0x03, 0xb6, 0xcf, 0xa2, 0x38,
	0x35, 0x9a, 0x83, 0xfa, 0xa8, 0x3b, 0x79, 0x5d, 0xb1, 0x4d, 0x5b, 0xb2, 0xf1, 0x8d, 0xca, 0xf0,
	0xb7, 0x06, 0x2d, 0x85, 0xfd, 0xb3, 0x9b, 0x3e, 0xb4, 0x39, 0xbd, 0xce, 0x42, 0x4e, 0xb7, 0xb2,
	0x98, 0x36, 0x2e, 0xcf, 0x79, 0x6f, 0x4a, 0x65, 0x5d, 0xd6, 0x5e, 0xb9, 0x37, 0x65, 0xad, 0x7a,
	0x0b, 0xca, 0x19, 0x2d, 0xa1, 0x93, 0x5f, 0x54, 0xc2, 0xea, 0x57, 0x7c, 0x53, 0x4d, 0x78, 0x46,
	0x04, 0x91, 0xb2, 0xed, 0x6d, 0x31, 0x3d, 0xff, 0xa9, 0x41, 0x53, 0xfe, 0x99, 0xd1, 0x53, 0x78,
	0xb2, 0xb4, 0xe7, 0x0b, 0x67, 0x7d, 0xe9, 0xcf, 0x9c, 0x0f, 0xae, 0xef, 0xcc, 0xf4, 0x07, 0xe8,
	0x19, 0x9c, 0x28, 0x70, 0xba, 0x58, 0x5c, 0xb8, 0xf6, 0x74, 0xe5, 0xce, 0x7d, 0x5d, 0x43, 0x3a,
	0x3c, 0x52, 0xb0, 0xe7, 0xac, 0xb0, 0x6b, 0xeb, 0x35, 0x84, 0xa0, 0xa7, 0x90, 0xc5, 0xc5, 0xd4,
	0xf7, 0x5d, 0xff, 0xa3, 0x5e, 0x47, 0xa7, 0xa0, 0x17, 0x18, 0x76, 0x66, 0xae, 0x2d, 0xb9, 0x0d,
	0x64, 0xc0, 0xa9, 0x42, 0xb1, 0x63, 0xcf, 0x3d, 0xcf, 0xf1, 0x67, 0x4a, 0xb5, 0x79, 0xab, 0x81,
	0x9d, 0xe5, 0xfc, 0x12, 0xdb, 0x8e, 0xde, 0x3a, 0xb7, 0xbf, 0x4e, 0xaf, 0x42, 0x51, 0xec, 0x60,
	0xdd, 0x6e, 0x7b, 0x46, 0x42, 0x2b, 0xff, 0xf4, 0x8f, 0x79, 0x9a, 0x36, 0x2d, 0xf9, 0x04, 0xbc,
	0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0x53, 0xc9, 0xf0, 0x3c, 0xc9, 0x04, 0x00, 0x00,
}