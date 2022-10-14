// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.7
// source: proto/brr.proto

package proto

import (
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

type AskForTimeMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId int64 `protobuf:"varint,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
}

func (x *AskForTimeMsg) Reset() {
	*x = AskForTimeMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_brr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskForTimeMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskForTimeMsg) ProtoMessage() {}

func (x *AskForTimeMsg) ProtoReflect() protoreflect.Message {
	mi := &file_proto_brr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AskForTimeMsg.ProtoReflect.Descriptor instead.
func (*AskForTimeMsg) Descriptor() ([]byte, []int) {
	return file_proto_brr_proto_rawDescGZIP(), []int{0}
}

func (x *AskForTimeMsg) GetClientId() int64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

type TimeMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time       string `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	ServerName string `protobuf:"bytes,2,opt,name=serverName,proto3" json:"serverName,omitempty"`
}

func (x *TimeMsg) Reset() {
	*x = TimeMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_brr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeMsg) ProtoMessage() {}

func (x *TimeMsg) ProtoReflect() protoreflect.Message {
	mi := &file_proto_brr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeMsg.ProtoReflect.Descriptor instead.
func (*TimeMsg) Descriptor() ([]byte, []int) {
	return file_proto_brr_proto_rawDescGZIP(), []int{1}
}

func (x *TimeMsg) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *TimeMsg) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

var File_proto_brr_proto protoreflect.FileDescriptor

var file_proto_brr_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x72, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x67, 0x72, 0x70, 0x63, 0x42, 0x72, 0x72, 0x22, 0x2b, 0x0a, 0x0d, 0x41, 0x73,
	0x6b, 0x46, 0x6f, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x4d,
	0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x45, 0x0a, 0x0e, 0x54, 0x69, 0x6d, 0x65, 0x41, 0x73,
	0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x42, 0x72, 0x72, 0x2e, 0x41, 0x73,
	0x6b, 0x46, 0x6f, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x67, 0x1a, 0x10, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x42, 0x72, 0x72, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x67, 0x42, 0x0d, 0x5a,
	0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_brr_proto_rawDescOnce sync.Once
	file_proto_brr_proto_rawDescData = file_proto_brr_proto_rawDesc
)

func file_proto_brr_proto_rawDescGZIP() []byte {
	file_proto_brr_proto_rawDescOnce.Do(func() {
		file_proto_brr_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_brr_proto_rawDescData)
	})
	return file_proto_brr_proto_rawDescData
}

var file_proto_brr_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_brr_proto_goTypes = []interface{}{
	(*AskForTimeMsg)(nil), // 0: grpcBrr.AskForTimeMsg
	(*TimeMsg)(nil),       // 1: grpcBrr.TimeMsg
}
var file_proto_brr_proto_depIdxs = []int32{
	0, // 0: grpcBrr.TimeAskService.GetTime:input_type -> grpcBrr.AskForTimeMsg
	1, // 1: grpcBrr.TimeAskService.GetTime:output_type -> grpcBrr.TimeMsg
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_brr_proto_init() }
func file_proto_brr_proto_init() {
	if File_proto_brr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_brr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AskForTimeMsg); i {
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
		file_proto_brr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeMsg); i {
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
			RawDescriptor: file_proto_brr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_brr_proto_goTypes,
		DependencyIndexes: file_proto_brr_proto_depIdxs,
		MessageInfos:      file_proto_brr_proto_msgTypes,
	}.Build()
	File_proto_brr_proto = out.File
	file_proto_brr_proto_rawDesc = nil
	file_proto_brr_proto_goTypes = nil
	file_proto_brr_proto_depIdxs = nil
}