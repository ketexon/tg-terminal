// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/msg.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Test struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Test          string                 `protobuf:"bytes,1,opt,name=test,proto3" json:"test,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Test) Reset() {
	*x = Test{}
	mi := &file_proto_msg_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_proto_msg_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_proto_msg_proto_rawDescGZIP(), []int{0}
}

func (x *Test) GetTest() string {
	if x != nil {
		return x.Test
	}
	return ""
}

var File_proto_msg_proto protoreflect.FileDescriptor

var file_proto_msg_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x1a, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x73, 0x74, 0x42, 0x12, 0x5a,
	0x10, 0x6b, 0x65, 0x74, 0x65, 0x78, 0x6f, 0x6e, 0x2f, 0x74, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_msg_proto_rawDescOnce sync.Once
	file_proto_msg_proto_rawDescData []byte
)

func file_proto_msg_proto_rawDescGZIP() []byte {
	file_proto_msg_proto_rawDescOnce.Do(func() {
		file_proto_msg_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_msg_proto_rawDesc), len(file_proto_msg_proto_rawDesc)))
	})
	return file_proto_msg_proto_rawDescData
}

var file_proto_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_msg_proto_goTypes = []any{
	(*Test)(nil), // 0: Test
}
var file_proto_msg_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_msg_proto_init() }
func file_proto_msg_proto_init() {
	if File_proto_msg_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_msg_proto_rawDesc), len(file_proto_msg_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_msg_proto_goTypes,
		DependencyIndexes: file_proto_msg_proto_depIdxs,
		MessageInfos:      file_proto_msg_proto_msgTypes,
	}.Build()
	File_proto_msg_proto = out.File
	file_proto_msg_proto_goTypes = nil
	file_proto_msg_proto_depIdxs = nil
}
