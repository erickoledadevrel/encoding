// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: fixtures.proto

package fixtures

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int64  `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B uint32 `protobuf:"fixed32,2,opt,name=b,proto3" json:"b,omitempty"`
	C uint64 `protobuf:"fixed64,3,opt,name=c,proto3" json:"c,omitempty"`
	D string `protobuf:"bytes,4,opt,name=d,proto3" json:"d,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fixtures_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_fixtures_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_fixtures_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetA() int64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *Message) GetB() uint32 {
	if x != nil {
		return x.B
	}
	return 0
}

func (x *Message) GetC() uint64 {
	if x != nil {
		return x.C
	}
	return 0
}

func (x *Message) GetD() string {
	if x != nil {
		return x.D
	}
	return ""
}

var File_fixtures_proto protoreflect.FileDescriptor

var file_fixtures_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x66, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x66, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0x41, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x07, 0x52, 0x01,
	0x62, 0x12, 0x0c, 0x0a, 0x01, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x06, 0x52, 0x01, 0x63, 0x12,
	0x0c, 0x0a, 0x01, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x64, 0x42, 0x2e, 0x5a,
	0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x69, 0x6f, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fixtures_proto_rawDescOnce sync.Once
	file_fixtures_proto_rawDescData = file_fixtures_proto_rawDesc
)

func file_fixtures_proto_rawDescGZIP() []byte {
	file_fixtures_proto_rawDescOnce.Do(func() {
		file_fixtures_proto_rawDescData = protoimpl.X.CompressGZIP(file_fixtures_proto_rawDescData)
	})
	return file_fixtures_proto_rawDescData
}

var file_fixtures_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_fixtures_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: fixtures.Message
}
var file_fixtures_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fixtures_proto_init() }
func file_fixtures_proto_init() {
	if File_fixtures_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fixtures_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_fixtures_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fixtures_proto_goTypes,
		DependencyIndexes: file_fixtures_proto_depIdxs,
		MessageInfos:      file_fixtures_proto_msgTypes,
	}.Build()
	File_fixtures_proto = out.File
	file_fixtures_proto_rawDesc = nil
	file_fixtures_proto_goTypes = nil
	file_fixtures_proto_depIdxs = nil
}
