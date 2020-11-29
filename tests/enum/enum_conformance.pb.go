// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: tests/enum/enum_conformance.proto

// clang-format off

package enum

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

type BasicEnum int32

const (
	BasicEnum_BASIC_INVALID BasicEnum = 0
	BasicEnum_BASIC_A       BasicEnum = 1
	BasicEnum_BASIC_B       BasicEnum = 2
	BasicEnum_BASIC_C       BasicEnum = 3
)

// Enum value maps for BasicEnum.
var (
	BasicEnum_name = map[int32]string{
		0: "BASIC_INVALID",
		1: "BASIC_A",
		2: "BASIC_B",
		3: "BASIC_C",
	}
	BasicEnum_value = map[string]int32{
		"BASIC_INVALID": 0,
		"BASIC_A":       1,
		"BASIC_B":       2,
		"BASIC_C":       3,
	}
)

func (x BasicEnum) Enum() *BasicEnum {
	p := new(BasicEnum)
	*p = x
	return p
}

func (x BasicEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BasicEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_tests_enum_enum_conformance_proto_enumTypes[0].Descriptor()
}

func (BasicEnum) Type() protoreflect.EnumType {
	return &file_tests_enum_enum_conformance_proto_enumTypes[0]
}

func (x BasicEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BasicEnum.Descriptor instead.
func (BasicEnum) EnumDescriptor() ([]byte, []int) {
	return file_tests_enum_enum_conformance_proto_rawDescGZIP(), []int{0}
}

type OuterMessage_InnerEnum int32

const (
	OuterMessage_INNER_INVALID OuterMessage_InnerEnum = 0
	OuterMessage_INNER_A       OuterMessage_InnerEnum = 1
	OuterMessage_INNER_B       OuterMessage_InnerEnum = 2
	OuterMessage_INNER_C       OuterMessage_InnerEnum = 3
)

// Enum value maps for OuterMessage_InnerEnum.
var (
	OuterMessage_InnerEnum_name = map[int32]string{
		0: "INNER_INVALID",
		1: "INNER_A",
		2: "INNER_B",
		3: "INNER_C",
	}
	OuterMessage_InnerEnum_value = map[string]int32{
		"INNER_INVALID": 0,
		"INNER_A":       1,
		"INNER_B":       2,
		"INNER_C":       3,
	}
)

func (x OuterMessage_InnerEnum) Enum() *OuterMessage_InnerEnum {
	p := new(OuterMessage_InnerEnum)
	*p = x
	return p
}

func (x OuterMessage_InnerEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OuterMessage_InnerEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_tests_enum_enum_conformance_proto_enumTypes[1].Descriptor()
}

func (OuterMessage_InnerEnum) Type() protoreflect.EnumType {
	return &file_tests_enum_enum_conformance_proto_enumTypes[1]
}

func (x OuterMessage_InnerEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OuterMessage_InnerEnum.Descriptor instead.
func (OuterMessage_InnerEnum) EnumDescriptor() ([]byte, []int) {
	return file_tests_enum_enum_conformance_proto_rawDescGZIP(), []int{0, 0}
}

type OuterMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OuterMessage) Reset() {
	*x = OuterMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_enum_enum_conformance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OuterMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OuterMessage) ProtoMessage() {}

func (x *OuterMessage) ProtoReflect() protoreflect.Message {
	mi := &file_tests_enum_enum_conformance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OuterMessage.ProtoReflect.Descriptor instead.
func (*OuterMessage) Descriptor() ([]byte, []int) {
	return file_tests_enum_enum_conformance_proto_rawDescGZIP(), []int{0}
}

var File_tests_enum_enum_conformance_proto protoreflect.FileDescriptor

var file_tests_enum_enum_conformance_proto_rawDesc = []byte{
	0x0a, 0x21, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x22,
	0x55, 0x0a, 0x0c, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x45, 0x0a, 0x09, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x11, 0x0a, 0x0d,
	0x49, 0x4e, 0x4e, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x4e, 0x45, 0x52, 0x5f, 0x41, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x49, 0x4e, 0x4e, 0x45, 0x52, 0x5f, 0x42, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x4e,
	0x45, 0x52, 0x5f, 0x43, 0x10, 0x03, 0x2a, 0x45, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x45,
	0x6e, 0x75, 0x6d, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x41, 0x53, 0x49, 0x43, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x41, 0x53, 0x49, 0x43, 0x5f,
	0x41, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x41, 0x53, 0x49, 0x43, 0x5f, 0x42, 0x10, 0x02,
	0x12, 0x0b, 0x0a, 0x07, 0x42, 0x41, 0x53, 0x49, 0x43, 0x5f, 0x43, 0x10, 0x03, 0x42, 0x27, 0x5a,
	0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x74, 0x61,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_enum_enum_conformance_proto_rawDescOnce sync.Once
	file_tests_enum_enum_conformance_proto_rawDescData = file_tests_enum_enum_conformance_proto_rawDesc
)

func file_tests_enum_enum_conformance_proto_rawDescGZIP() []byte {
	file_tests_enum_enum_conformance_proto_rawDescOnce.Do(func() {
		file_tests_enum_enum_conformance_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_enum_enum_conformance_proto_rawDescData)
	})
	return file_tests_enum_enum_conformance_proto_rawDescData
}

var file_tests_enum_enum_conformance_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_tests_enum_enum_conformance_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tests_enum_enum_conformance_proto_goTypes = []interface{}{
	(BasicEnum)(0),              // 0: tests.enum.BasicEnum
	(OuterMessage_InnerEnum)(0), // 1: tests.enum.OuterMessage.InnerEnum
	(*OuterMessage)(nil),        // 2: tests.enum.OuterMessage
}
var file_tests_enum_enum_conformance_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tests_enum_enum_conformance_proto_init() }
func file_tests_enum_enum_conformance_proto_init() {
	if File_tests_enum_enum_conformance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_enum_enum_conformance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OuterMessage); i {
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
			RawDescriptor: file_tests_enum_enum_conformance_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_enum_enum_conformance_proto_goTypes,
		DependencyIndexes: file_tests_enum_enum_conformance_proto_depIdxs,
		EnumInfos:         file_tests_enum_enum_conformance_proto_enumTypes,
		MessageInfos:      file_tests_enum_enum_conformance_proto_msgTypes,
	}.Build()
	File_tests_enum_enum_conformance_proto = out.File
	file_tests_enum_enum_conformance_proto_rawDesc = nil
	file_tests_enum_enum_conformance_proto_goTypes = nil
	file_tests_enum_enum_conformance_proto_depIdxs = nil
}
