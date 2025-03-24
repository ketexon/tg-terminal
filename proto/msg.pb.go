// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0
// source: proto/msg.proto

package pb

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

type PlayAudio struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Clip          string                 `protobuf:"bytes,1,opt,name=clip,proto3" json:"clip,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlayAudio) Reset() {
	*x = PlayAudio{}
	mi := &file_proto_msg_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlayAudio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayAudio) ProtoMessage() {}

func (x *PlayAudio) ProtoReflect() protoreflect.Message {
	mi := &file_proto_msg_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayAudio.ProtoReflect.Descriptor instead.
func (*PlayAudio) Descriptor() ([]byte, []int) {
	return file_proto_msg_proto_rawDescGZIP(), []int{1}
}

func (x *PlayAudio) GetClip() string {
	if x != nil {
		return x.Clip
	}
	return ""
}

type OpenDoor struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OpenDoor) Reset() {
	*x = OpenDoor{}
	mi := &file_proto_msg_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OpenDoor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenDoor) ProtoMessage() {}

func (x *OpenDoor) ProtoReflect() protoreflect.Message {
	mi := &file_proto_msg_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenDoor.ProtoReflect.Descriptor instead.
func (*OpenDoor) Descriptor() ([]byte, []int) {
	return file_proto_msg_proto_rawDescGZIP(), []int{2}
}

type TerminalMessage struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Content:
	//
	//	*TerminalMessage_PlayAudio
	//	*TerminalMessage_OpenDoor
	//	*TerminalMessage_Ping
	//	*TerminalMessage_Error
	Content       isTerminalMessage_Content `protobuf_oneof:"Content"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TerminalMessage) Reset() {
	*x = TerminalMessage{}
	mi := &file_proto_msg_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TerminalMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminalMessage) ProtoMessage() {}

func (x *TerminalMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_msg_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminalMessage.ProtoReflect.Descriptor instead.
func (*TerminalMessage) Descriptor() ([]byte, []int) {
	return file_proto_msg_proto_rawDescGZIP(), []int{3}
}

func (x *TerminalMessage) GetContent() isTerminalMessage_Content {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *TerminalMessage) GetPlayAudio() *PlayAudio {
	if x != nil {
		if x, ok := x.Content.(*TerminalMessage_PlayAudio); ok {
			return x.PlayAudio
		}
	}
	return nil
}

func (x *TerminalMessage) GetOpenDoor() *OpenDoor {
	if x != nil {
		if x, ok := x.Content.(*TerminalMessage_OpenDoor); ok {
			return x.OpenDoor
		}
	}
	return nil
}

func (x *TerminalMessage) GetPing() *Ping {
	if x != nil {
		if x, ok := x.Content.(*TerminalMessage_Ping); ok {
			return x.Ping
		}
	}
	return nil
}

func (x *TerminalMessage) GetError() *Error {
	if x != nil {
		if x, ok := x.Content.(*TerminalMessage_Error); ok {
			return x.Error
		}
	}
	return nil
}

type isTerminalMessage_Content interface {
	isTerminalMessage_Content()
}

type TerminalMessage_PlayAudio struct {
	PlayAudio *PlayAudio `protobuf:"bytes,1,opt,name=play_audio,json=playAudio,proto3,oneof"`
}

type TerminalMessage_OpenDoor struct {
	OpenDoor *OpenDoor `protobuf:"bytes,2,opt,name=open_door,json=openDoor,proto3,oneof"`
}

type TerminalMessage_Ping struct {
	Ping *Ping `protobuf:"bytes,3,opt,name=ping,proto3,oneof"`
}

type TerminalMessage_Error struct {
	Error *Error `protobuf:"bytes,4,opt,name=error,proto3,oneof"`
}

func (*TerminalMessage_PlayAudio) isTerminalMessage_Content() {}

func (*TerminalMessage_OpenDoor) isTerminalMessage_Content() {}

func (*TerminalMessage_Ping) isTerminalMessage_Content() {}

func (*TerminalMessage_Error) isTerminalMessage_Content() {}

type Login struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ComputerId    string                 `protobuf:"bytes,1,opt,name=computer_id,json=computerId,proto3" json:"computer_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Login) Reset() {
	*x = Login{}
	mi := &file_proto_msg_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Login) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Login) ProtoMessage() {}

func (x *Login) ProtoReflect() protoreflect.Message {
	mi := &file_proto_msg_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Login.ProtoReflect.Descriptor instead.
func (*Login) Descriptor() ([]byte, []int) {
	return file_proto_msg_proto_rawDescGZIP(), []int{4}
}

func (x *Login) GetComputerId() string {
	if x != nil {
		return x.ComputerId
	}
	return ""
}

type Error struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Error) Reset() {
	*x = Error{}
	mi := &file_proto_msg_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_proto_msg_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_proto_msg_proto_rawDescGZIP(), []int{5}
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Ping struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Timestamp     int64                  `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Ping) Reset() {
	*x = Ping{}
	mi := &file_proto_msg_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_proto_msg_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_proto_msg_proto_rawDescGZIP(), []int{6}
}

func (x *Ping) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type GameMessage struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Content:
	//
	//	*GameMessage_Login
	//	*GameMessage_Ping
	Content       isGameMessage_Content `protobuf_oneof:"Content"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GameMessage) Reset() {
	*x = GameMessage{}
	mi := &file_proto_msg_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GameMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameMessage) ProtoMessage() {}

func (x *GameMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_msg_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameMessage.ProtoReflect.Descriptor instead.
func (*GameMessage) Descriptor() ([]byte, []int) {
	return file_proto_msg_proto_rawDescGZIP(), []int{7}
}

func (x *GameMessage) GetContent() isGameMessage_Content {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *GameMessage) GetLogin() *Login {
	if x != nil {
		if x, ok := x.Content.(*GameMessage_Login); ok {
			return x.Login
		}
	}
	return nil
}

func (x *GameMessage) GetPing() *Ping {
	if x != nil {
		if x, ok := x.Content.(*GameMessage_Ping); ok {
			return x.Ping
		}
	}
	return nil
}

type isGameMessage_Content interface {
	isGameMessage_Content()
}

type GameMessage_Login struct {
	Login *Login `protobuf:"bytes,1,opt,name=login,proto3,oneof"`
}

type GameMessage_Ping struct {
	Ping *Ping `protobuf:"bytes,2,opt,name=ping,proto3,oneof"`
}

func (*GameMessage_Login) isGameMessage_Content() {}

func (*GameMessage_Ping) isGameMessage_Content() {}

var File_proto_msg_proto protoreflect.FileDescriptor

var file_proto_msg_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x1a, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x73, 0x74, 0x22, 0x1f, 0x0a,
	0x09, 0x50, 0x6c, 0x61, 0x79, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6c,
	0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6c, 0x69, 0x70, 0x22, 0x0a,
	0x0a, 0x08, 0x4f, 0x70, 0x65, 0x6e, 0x44, 0x6f, 0x6f, 0x72, 0x22, 0xb0, 0x01, 0x0a, 0x0f, 0x54,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b,
	0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x48, 0x00,
	0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x28, 0x0a, 0x09, 0x6f,
	0x70, 0x65, 0x6e, 0x5f, 0x64, 0x6f, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x44, 0x6f, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x08, 0x6f, 0x70, 0x65,
	0x6e, 0x44, 0x6f, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x04, 0x70, 0x69,
	0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x06, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x28, 0x0a,
	0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x21, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x24, 0x0a, 0x04, 0x50, 0x69,
	0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x22, 0x55, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1e, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x1b, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e,
	0x50, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x42, 0x09, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x0f, 0x5a, 0x0d, 0x6b, 0x65, 0x74, 0x65, 0x78,
	0x6f, 0x6e, 0x2f, 0x74, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_proto_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_msg_proto_goTypes = []any{
	(*Test)(nil),            // 0: Test
	(*PlayAudio)(nil),       // 1: PlayAudio
	(*OpenDoor)(nil),        // 2: OpenDoor
	(*TerminalMessage)(nil), // 3: TerminalMessage
	(*Login)(nil),           // 4: Login
	(*Error)(nil),           // 5: Error
	(*Ping)(nil),            // 6: Ping
	(*GameMessage)(nil),     // 7: GameMessage
}
var file_proto_msg_proto_depIdxs = []int32{
	1, // 0: TerminalMessage.play_audio:type_name -> PlayAudio
	2, // 1: TerminalMessage.open_door:type_name -> OpenDoor
	6, // 2: TerminalMessage.ping:type_name -> Ping
	5, // 3: TerminalMessage.error:type_name -> Error
	4, // 4: GameMessage.login:type_name -> Login
	6, // 5: GameMessage.ping:type_name -> Ping
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_msg_proto_init() }
func file_proto_msg_proto_init() {
	if File_proto_msg_proto != nil {
		return
	}
	file_proto_msg_proto_msgTypes[3].OneofWrappers = []any{
		(*TerminalMessage_PlayAudio)(nil),
		(*TerminalMessage_OpenDoor)(nil),
		(*TerminalMessage_Ping)(nil),
		(*TerminalMessage_Error)(nil),
	}
	file_proto_msg_proto_msgTypes[7].OneofWrappers = []any{
		(*GameMessage_Login)(nil),
		(*GameMessage_Ping)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_msg_proto_rawDesc), len(file_proto_msg_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
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
