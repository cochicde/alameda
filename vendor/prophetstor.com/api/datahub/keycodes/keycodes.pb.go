// This file has messages related to keycode managements

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: datahub/keycodes/keycodes.proto

package keycodes

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

type Keycode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keycode          string               `protobuf:"bytes,1,opt,name=keycode,proto3" json:"keycode,omitempty"`                                      // example: "A5IMH-KBAFI-XTEDK-G4OQM-QMM67-4TEST"`
	KeycodeType      string               `protobuf:"bytes,2,opt,name=keycode_type,json=keycodeType,proto3" json:"keycode_type,omitempty"`           // example: "Regular/Trial"`
	KeycodeVersion   int32                `protobuf:"varint,3,opt,name=keycode_version,json=keycodeVersion,proto3" json:"keycode_version,omitempty"` // example: "2"`
	ApplyTime        *timestamp.Timestamp `protobuf:"bytes,4,opt,name=apply_time,json=applyTime,proto3" json:"apply_time,omitempty"`                 // example: "2018-01-01T00:00:00Z"`
	ExpireTime       *timestamp.Timestamp `protobuf:"bytes,5,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`              // example: "2018-12-31T11:59:59Z"`
	LicenseState     string               `protobuf:"bytes,6,opt,name=license_state,json=licenseState,proto3" json:"license_state,omitempty"`        // example: "Valid/Invalid/Expired"`
	Registered       bool                 `protobuf:"varint,7,opt,name=registered,proto3" json:"registered,omitempty"`                               // example: "false"`
	Capacity         *Capacity            `protobuf:"bytes,8,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Functionality    *Functionality       `protobuf:"bytes,9,opt,name=functionality,proto3" json:"functionality,omitempty"`
	Retention        *Retention           `protobuf:"bytes,10,opt,name=retention,proto3" json:"retention,omitempty"`
	ServiceAgreement *ServiceAgreement    `protobuf:"bytes,11,opt,name=service_agreement,json=serviceAgreement,proto3" json:"service_agreement,omitempty"`
}

func (x *Keycode) Reset() {
	*x = Keycode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_keycodes_keycodes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Keycode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Keycode) ProtoMessage() {}

func (x *Keycode) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_keycodes_keycodes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Keycode.ProtoReflect.Descriptor instead.
func (*Keycode) Descriptor() ([]byte, []int) {
	return file_datahub_keycodes_keycodes_proto_rawDescGZIP(), []int{0}
}

func (x *Keycode) GetKeycode() string {
	if x != nil {
		return x.Keycode
	}
	return ""
}

func (x *Keycode) GetKeycodeType() string {
	if x != nil {
		return x.KeycodeType
	}
	return ""
}

func (x *Keycode) GetKeycodeVersion() int32 {
	if x != nil {
		return x.KeycodeVersion
	}
	return 0
}

func (x *Keycode) GetApplyTime() *timestamp.Timestamp {
	if x != nil {
		return x.ApplyTime
	}
	return nil
}

func (x *Keycode) GetExpireTime() *timestamp.Timestamp {
	if x != nil {
		return x.ExpireTime
	}
	return nil
}

func (x *Keycode) GetLicenseState() string {
	if x != nil {
		return x.LicenseState
	}
	return ""
}

func (x *Keycode) GetRegistered() bool {
	if x != nil {
		return x.Registered
	}
	return false
}

func (x *Keycode) GetCapacity() *Capacity {
	if x != nil {
		return x.Capacity
	}
	return nil
}

func (x *Keycode) GetFunctionality() *Functionality {
	if x != nil {
		return x.Functionality
	}
	return nil
}

func (x *Keycode) GetRetention() *Retention {
	if x != nil {
		return x.Retention
	}
	return nil
}

func (x *Keycode) GetServiceAgreement() *ServiceAgreement {
	if x != nil {
		return x.ServiceAgreement
	}
	return nil
}

var File_datahub_keycodes_keycodes_proto protoreflect.FileDescriptor

var file_datahub_keycodes_keycodes_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x6b, 0x65, 0x79, 0x63, 0x6f, 0x64,
	0x65, 0x73, 0x2f, 0x6b, 0x65, 0x79, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1c, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x6b, 0x65, 0x79, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x1a,
	0x1c, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x6b, 0x65, 0x79, 0x63, 0x6f, 0x64, 0x65,
	0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7,
	0x04, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65,
	0x79, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6b, 0x65, 0x79, 0x63, 0x6f, 0x64, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6b, 0x65, 0x79, 0x63,
	0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6b, 0x65, 0x79, 0x63, 0x6f,
	0x64, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x6b, 0x65, 0x79, 0x63, 0x6f, 0x64, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x39, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12, 0x42, 0x0a,
	0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x6b, 0x65, 0x79, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x43,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74,
	0x79, 0x12, 0x51, 0x0a, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68,
	0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x6b,
	0x65, 0x79, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x45, 0x0a, 0x09, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65,
	0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x6b, 0x65,
	0x79, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x09, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5b, 0x0a, 0x11, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x6b, 0x65, 0x79,
	0x63, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x67, 0x72,
	0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41,
	0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x26, 0x5a, 0x24, 0x70, 0x72, 0x6f, 0x70,
	0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x6b, 0x65, 0x79, 0x63, 0x6f, 0x64, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datahub_keycodes_keycodes_proto_rawDescOnce sync.Once
	file_datahub_keycodes_keycodes_proto_rawDescData = file_datahub_keycodes_keycodes_proto_rawDesc
)

func file_datahub_keycodes_keycodes_proto_rawDescGZIP() []byte {
	file_datahub_keycodes_keycodes_proto_rawDescOnce.Do(func() {
		file_datahub_keycodes_keycodes_proto_rawDescData = protoimpl.X.CompressGZIP(file_datahub_keycodes_keycodes_proto_rawDescData)
	})
	return file_datahub_keycodes_keycodes_proto_rawDescData
}

var file_datahub_keycodes_keycodes_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_datahub_keycodes_keycodes_proto_goTypes = []interface{}{
	(*Keycode)(nil),             // 0: prophetstor.datahub.keycodes.Keycode
	(*timestamp.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*Capacity)(nil),            // 2: prophetstor.datahub.keycodes.Capacity
	(*Functionality)(nil),       // 3: prophetstor.datahub.keycodes.Functionality
	(*Retention)(nil),           // 4: prophetstor.datahub.keycodes.Retention
	(*ServiceAgreement)(nil),    // 5: prophetstor.datahub.keycodes.ServiceAgreement
}
var file_datahub_keycodes_keycodes_proto_depIdxs = []int32{
	1, // 0: prophetstor.datahub.keycodes.Keycode.apply_time:type_name -> google.protobuf.Timestamp
	1, // 1: prophetstor.datahub.keycodes.Keycode.expire_time:type_name -> google.protobuf.Timestamp
	2, // 2: prophetstor.datahub.keycodes.Keycode.capacity:type_name -> prophetstor.datahub.keycodes.Capacity
	3, // 3: prophetstor.datahub.keycodes.Keycode.functionality:type_name -> prophetstor.datahub.keycodes.Functionality
	4, // 4: prophetstor.datahub.keycodes.Keycode.retention:type_name -> prophetstor.datahub.keycodes.Retention
	5, // 5: prophetstor.datahub.keycodes.Keycode.service_agreement:type_name -> prophetstor.datahub.keycodes.ServiceAgreement
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_datahub_keycodes_keycodes_proto_init() }
func file_datahub_keycodes_keycodes_proto_init() {
	if File_datahub_keycodes_keycodes_proto != nil {
		return
	}
	file_datahub_keycodes_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_datahub_keycodes_keycodes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Keycode); i {
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
			RawDescriptor: file_datahub_keycodes_keycodes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datahub_keycodes_keycodes_proto_goTypes,
		DependencyIndexes: file_datahub_keycodes_keycodes_proto_depIdxs,
		MessageInfos:      file_datahub_keycodes_keycodes_proto_msgTypes,
	}.Build()
	File_datahub_keycodes_keycodes_proto = out.File
	file_datahub_keycodes_keycodes_proto_rawDesc = nil
	file_datahub_keycodes_keycodes_proto_goTypes = nil
	file_datahub_keycodes_keycodes_proto_depIdxs = nil
}