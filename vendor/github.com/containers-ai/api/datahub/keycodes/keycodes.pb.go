// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datahub/keycodes/keycodes.proto

package keycodes

import (
	fmt "fmt"
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

type Keycode struct {
	Keycode              string               `protobuf:"bytes,1,opt,name=keycode,proto3" json:"keycode,omitempty"`
	KeycodeType          string               `protobuf:"bytes,2,opt,name=keycode_type,json=keycodeType,proto3" json:"keycode_type,omitempty"`
	KeycodeVersion       int32                `protobuf:"varint,3,opt,name=keycode_version,json=keycodeVersion,proto3" json:"keycode_version,omitempty"`
	ApplyTime            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=apply_time,json=applyTime,proto3" json:"apply_time,omitempty"`
	ExpireTime           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	LicenseState         string               `protobuf:"bytes,6,opt,name=license_state,json=licenseState,proto3" json:"license_state,omitempty"`
	Registered           bool                 `protobuf:"varint,7,opt,name=registered,proto3" json:"registered,omitempty"`
	Capacity             *Capacity            `protobuf:"bytes,8,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Functionality        *Functionality       `protobuf:"bytes,9,opt,name=functionality,proto3" json:"functionality,omitempty"`
	Retention            *Retention           `protobuf:"bytes,10,opt,name=retention,proto3" json:"retention,omitempty"`
	ServiceAgreement     *ServiceAgreement    `protobuf:"bytes,11,opt,name=service_agreement,json=serviceAgreement,proto3" json:"service_agreement,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Keycode) Reset()         { *m = Keycode{} }
func (m *Keycode) String() string { return proto.CompactTextString(m) }
func (*Keycode) ProtoMessage()    {}
func (*Keycode) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a20bf7958a5601e, []int{0}
}

func (m *Keycode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Keycode.Unmarshal(m, b)
}
func (m *Keycode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Keycode.Marshal(b, m, deterministic)
}
func (m *Keycode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keycode.Merge(m, src)
}
func (m *Keycode) XXX_Size() int {
	return xxx_messageInfo_Keycode.Size(m)
}
func (m *Keycode) XXX_DiscardUnknown() {
	xxx_messageInfo_Keycode.DiscardUnknown(m)
}

var xxx_messageInfo_Keycode proto.InternalMessageInfo

func (m *Keycode) GetKeycode() string {
	if m != nil {
		return m.Keycode
	}
	return ""
}

func (m *Keycode) GetKeycodeType() string {
	if m != nil {
		return m.KeycodeType
	}
	return ""
}

func (m *Keycode) GetKeycodeVersion() int32 {
	if m != nil {
		return m.KeycodeVersion
	}
	return 0
}

func (m *Keycode) GetApplyTime() *timestamp.Timestamp {
	if m != nil {
		return m.ApplyTime
	}
	return nil
}

func (m *Keycode) GetExpireTime() *timestamp.Timestamp {
	if m != nil {
		return m.ExpireTime
	}
	return nil
}

func (m *Keycode) GetLicenseState() string {
	if m != nil {
		return m.LicenseState
	}
	return ""
}

func (m *Keycode) GetRegistered() bool {
	if m != nil {
		return m.Registered
	}
	return false
}

func (m *Keycode) GetCapacity() *Capacity {
	if m != nil {
		return m.Capacity
	}
	return nil
}

func (m *Keycode) GetFunctionality() *Functionality {
	if m != nil {
		return m.Functionality
	}
	return nil
}

func (m *Keycode) GetRetention() *Retention {
	if m != nil {
		return m.Retention
	}
	return nil
}

func (m *Keycode) GetServiceAgreement() *ServiceAgreement {
	if m != nil {
		return m.ServiceAgreement
	}
	return nil
}

type Capacity struct {
	Users                int32    `protobuf:"varint,1,opt,name=users,proto3" json:"users,omitempty"`
	Hosts                int32    `protobuf:"varint,2,opt,name=hosts,proto3" json:"hosts,omitempty"`
	Disks                int32    `protobuf:"varint,3,opt,name=disks,proto3" json:"disks,omitempty"`
	Cpus                 int32    `protobuf:"varint,4,opt,name=cpus,proto3" json:"cpus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Capacity) Reset()         { *m = Capacity{} }
func (m *Capacity) String() string { return proto.CompactTextString(m) }
func (*Capacity) ProtoMessage()    {}
func (*Capacity) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a20bf7958a5601e, []int{1}
}

func (m *Capacity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Capacity.Unmarshal(m, b)
}
func (m *Capacity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Capacity.Marshal(b, m, deterministic)
}
func (m *Capacity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Capacity.Merge(m, src)
}
func (m *Capacity) XXX_Size() int {
	return xxx_messageInfo_Capacity.Size(m)
}
func (m *Capacity) XXX_DiscardUnknown() {
	xxx_messageInfo_Capacity.DiscardUnknown(m)
}

var xxx_messageInfo_Capacity proto.InternalMessageInfo

func (m *Capacity) GetUsers() int32 {
	if m != nil {
		return m.Users
	}
	return 0
}

func (m *Capacity) GetHosts() int32 {
	if m != nil {
		return m.Hosts
	}
	return 0
}

func (m *Capacity) GetDisks() int32 {
	if m != nil {
		return m.Disks
	}
	return 0
}

func (m *Capacity) GetCpus() int32 {
	if m != nil {
		return m.Cpus
	}
	return 0
}

type Functionality struct {
	DiskProphet          bool     `protobuf:"varint,1,opt,name=disk_prophet,json=diskProphet,proto3" json:"disk_prophet,omitempty"`
	Workload             bool     `protobuf:"varint,2,opt,name=workload,proto3" json:"workload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Functionality) Reset()         { *m = Functionality{} }
func (m *Functionality) String() string { return proto.CompactTextString(m) }
func (*Functionality) ProtoMessage()    {}
func (*Functionality) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a20bf7958a5601e, []int{2}
}

func (m *Functionality) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Functionality.Unmarshal(m, b)
}
func (m *Functionality) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Functionality.Marshal(b, m, deterministic)
}
func (m *Functionality) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Functionality.Merge(m, src)
}
func (m *Functionality) XXX_Size() int {
	return xxx_messageInfo_Functionality.Size(m)
}
func (m *Functionality) XXX_DiscardUnknown() {
	xxx_messageInfo_Functionality.DiscardUnknown(m)
}

var xxx_messageInfo_Functionality proto.InternalMessageInfo

func (m *Functionality) GetDiskProphet() bool {
	if m != nil {
		return m.DiskProphet
	}
	return false
}

func (m *Functionality) GetWorkload() bool {
	if m != nil {
		return m.Workload
	}
	return false
}

type Retention struct {
	ValidMonth           int32    `protobuf:"varint,1,opt,name=valid_month,json=validMonth,proto3" json:"valid_month,omitempty"`
	Years                int32    `protobuf:"varint,2,opt,name=years,proto3" json:"years,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Retention) Reset()         { *m = Retention{} }
func (m *Retention) String() string { return proto.CompactTextString(m) }
func (*Retention) ProtoMessage()    {}
func (*Retention) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a20bf7958a5601e, []int{3}
}

func (m *Retention) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Retention.Unmarshal(m, b)
}
func (m *Retention) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Retention.Marshal(b, m, deterministic)
}
func (m *Retention) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Retention.Merge(m, src)
}
func (m *Retention) XXX_Size() int {
	return xxx_messageInfo_Retention.Size(m)
}
func (m *Retention) XXX_DiscardUnknown() {
	xxx_messageInfo_Retention.DiscardUnknown(m)
}

var xxx_messageInfo_Retention proto.InternalMessageInfo

func (m *Retention) GetValidMonth() int32 {
	if m != nil {
		return m.ValidMonth
	}
	return 0
}

func (m *Retention) GetYears() int32 {
	if m != nil {
		return m.Years
	}
	return 0
}

type ServiceAgreement struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceAgreement) Reset()         { *m = ServiceAgreement{} }
func (m *ServiceAgreement) String() string { return proto.CompactTextString(m) }
func (*ServiceAgreement) ProtoMessage()    {}
func (*ServiceAgreement) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a20bf7958a5601e, []int{4}
}

func (m *ServiceAgreement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceAgreement.Unmarshal(m, b)
}
func (m *ServiceAgreement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceAgreement.Marshal(b, m, deterministic)
}
func (m *ServiceAgreement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceAgreement.Merge(m, src)
}
func (m *ServiceAgreement) XXX_Size() int {
	return xxx_messageInfo_ServiceAgreement.Size(m)
}
func (m *ServiceAgreement) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceAgreement.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceAgreement proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Keycode)(nil), "containersai.datahub.keycodes.Keycode")
	proto.RegisterType((*Capacity)(nil), "containersai.datahub.keycodes.Capacity")
	proto.RegisterType((*Functionality)(nil), "containersai.datahub.keycodes.Functionality")
	proto.RegisterType((*Retention)(nil), "containersai.datahub.keycodes.Retention")
	proto.RegisterType((*ServiceAgreement)(nil), "containersai.datahub.keycodes.ServiceAgreement")
}

func init() { proto.RegisterFile("datahub/keycodes/keycodes.proto", fileDescriptor_7a20bf7958a5601e) }

var fileDescriptor_7a20bf7958a5601e = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xc1, 0x6f, 0xd3, 0x3e,
	0x14, 0xc7, 0xd5, 0xdf, 0xaf, 0x59, 0xd3, 0x97, 0x15, 0x86, 0xc5, 0xc1, 0xaa, 0x04, 0x2d, 0xe5,
	0xb0, 0x1e, 0x58, 0x22, 0xc1, 0x09, 0x71, 0x62, 0x93, 0x76, 0x41, 0x20, 0xe4, 0x4d, 0x1c, 0x10,
	0x52, 0x70, 0x93, 0xb7, 0xd6, 0x6a, 0x1a, 0x5b, 0xb6, 0x5b, 0xc8, 0xdf, 0xcc, 0x3f, 0x81, 0x6c,
	0x27, 0xdd, 0xd6, 0x03, 0xbd, 0xf9, 0x7d, 0xf2, 0xfd, 0x7e, 0xfd, 0xf2, 0xf2, 0x02, 0x93, 0x92,
	0x5b, 0xbe, 0xda, 0x2e, 0xb2, 0x35, 0x36, 0x85, 0x2c, 0xd1, 0xec, 0x0f, 0xa9, 0xd2, 0xd2, 0x4a,
	0xf2, 0xa2, 0x90, 0xb5, 0xe5, 0xa2, 0x46, 0x6d, 0xb8, 0x48, 0x5b, 0x75, 0xda, 0x89, 0xc6, 0x93,
	0xa5, 0x94, 0xcb, 0x0a, 0x33, 0x2f, 0x5e, 0x6c, 0xef, 0x32, 0x2b, 0x36, 0x68, 0x2c, 0xdf, 0xa8,
	0xe0, 0x9f, 0xfd, 0xe9, 0xc3, 0xe0, 0x53, 0x50, 0x13, 0x0a, 0x83, 0xd6, 0x48, 0x7b, 0xd3, 0xde,
	0x7c, 0xc8, 0xba, 0x92, 0xbc, 0x82, 0xd3, 0xf6, 0x98, 0xdb, 0x46, 0x21, 0xfd, 0xcf, 0x3f, 0x4e,
	0x5a, 0x76, 0xdb, 0x28, 0x24, 0xe7, 0xf0, 0xb4, 0x93, 0xec, 0x50, 0x1b, 0x21, 0x6b, 0xfa, 0xff,
	0xb4, 0x37, 0x8f, 0xd8, 0x93, 0x16, 0x7f, 0x0b, 0x94, 0xbc, 0x07, 0xe0, 0x4a, 0x55, 0x4d, 0xee,
	0x5a, 0xa1, 0xfd, 0x69, 0x6f, 0x9e, 0xbc, 0x1d, 0xa7, 0xa1, 0xcf, 0xb4, 0xeb, 0x33, 0xbd, 0xed,
	0xfa, 0x64, 0x43, 0xaf, 0x76, 0x35, 0xf9, 0x00, 0x09, 0xfe, 0x56, 0x42, 0x63, 0xf0, 0x46, 0x47,
	0xbd, 0x10, 0xe4, 0xde, 0xfc, 0x1a, 0x46, 0x95, 0x28, 0xb0, 0x36, 0x98, 0x1b, 0xcb, 0x2d, 0xd2,
	0x13, 0xff, 0x12, 0xa7, 0x2d, 0xbc, 0x71, 0x8c, 0xbc, 0x04, 0xd0, 0xb8, 0x14, 0xc6, 0xa2, 0xc6,
	0x92, 0x0e, 0xa6, 0xbd, 0x79, 0xcc, 0x1e, 0x10, 0x72, 0x05, 0x71, 0xc1, 0x15, 0x2f, 0x84, 0x6d,
	0x68, 0xec, 0xaf, 0x3f, 0x4f, 0xff, 0xf9, 0x05, 0xd2, 0xab, 0x56, 0xce, 0xf6, 0x46, 0xc2, 0x60,
	0x74, 0xb7, 0xad, 0x0b, 0x2b, 0x64, 0xcd, 0x2b, 0x97, 0x34, 0xf4, 0x49, 0x6f, 0x8e, 0x24, 0x5d,
	0x3f, 0xf4, 0xb0, 0xc7, 0x11, 0xe4, 0x1a, 0x86, 0x1a, 0x2d, 0xd6, 0x8e, 0x50, 0xf0, 0x79, 0xf3,
	0x23, 0x79, 0xac, 0xd3, 0xb3, 0x7b, 0x2b, 0xf9, 0x01, 0xcf, 0x0c, 0xea, 0x9d, 0x28, 0x30, 0xe7,
	0x4b, 0x8d, 0xb8, 0xc1, 0xda, 0xd2, 0xc4, 0xe7, 0x65, 0x47, 0xf2, 0x6e, 0x82, 0xef, 0x63, 0x67,
	0x63, 0x67, 0xe6, 0x80, 0xcc, 0x7e, 0x42, 0xdc, 0xcd, 0x83, 0x3c, 0x87, 0x68, 0x6b, 0x50, 0x1b,
	0xbf, 0x6b, 0x11, 0x0b, 0x85, 0xa3, 0x2b, 0x69, 0xac, 0xf1, 0x2b, 0x16, 0xb1, 0x50, 0x38, 0x5a,
	0x0a, 0xb3, 0x36, 0xed, 0x4a, 0x85, 0x82, 0x10, 0xe8, 0x17, 0x6a, 0x6b, 0xfc, 0x0e, 0x45, 0xcc,
	0x9f, 0x67, 0x5f, 0x60, 0xf4, 0x68, 0x4e, 0x6e, 0x75, 0x9d, 0x3a, 0x57, 0x5a, 0xaa, 0x15, 0x5a,
	0x7f, 0x5b, 0xcc, 0x12, 0xc7, 0xbe, 0x06, 0x44, 0xc6, 0x10, 0xff, 0x92, 0x7a, 0x5d, 0x49, 0x5e,
	0xfa, 0x6b, 0x63, 0xb6, 0xaf, 0x67, 0x97, 0x30, 0xdc, 0xcf, 0x89, 0x4c, 0x20, 0xd9, 0xf1, 0x4a,
	0x94, 0xf9, 0x46, 0xd6, 0x76, 0xd5, 0x36, 0x0e, 0x1e, 0x7d, 0x76, 0xc4, 0xf5, 0xd9, 0x20, 0xd7,
	0xfb, 0xee, 0x7d, 0x31, 0x23, 0x70, 0x76, 0x38, 0x9b, 0xcb, 0xec, 0xfb, 0xc5, 0x52, 0x58, 0x37,
	0xbf, 0x42, 0x6e, 0xb2, 0xfb, 0xc1, 0x5e, 0x70, 0x91, 0x71, 0x25, 0xb2, 0xc3, 0xff, 0x7e, 0x71,
	0xe2, 0xd7, 0xfb, 0xdd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x31, 0x19, 0x0e, 0x16, 0x12, 0x04,
	0x00, 0x00,
}
