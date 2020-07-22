// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        (unknown)
// source: proto/prase.proto

package server

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EncodedHTML string `protobuf:"bytes,1,opt,name=EncodedHTML,json=encodedHTML,proto3" json:"EncodedHTML,omitempty"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_prase_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_proto_prase_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_proto_prase_proto_rawDescGZIP(), []int{0}
}

func (x *Input) GetEncodedHTML() string {
	if x != nil {
		return x.EncodedHTML
	}
	return ""
}

type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title     []string `protobuf:"bytes,1,rep,name=Title,json=title,proto3" json:"Title,omitempty"`
	Canonical []string `protobuf:"bytes,2,rep,name=Canonical,json=canonical,proto3" json:"Canonical,omitempty"`
}

func (x *Output) Reset() {
	*x = Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_prase_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_proto_prase_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_proto_prase_proto_rawDescGZIP(), []int{1}
}

func (x *Output) GetTitle() []string {
	if x != nil {
		return x.Title
	}
	return nil
}

func (x *Output) GetCanonical() []string {
	if x != nil {
		return x.Canonical
	}
	return nil
}

var File_proto_prase_proto protoreflect.FileDescriptor

var file_proto_prase_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x69, 0x6f, 0x22, 0x29, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x48, 0x54, 0x4d, 0x4c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x48, 0x54,
	0x4d, 0x4c, 0x22, 0x3c, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c,
	0x32, 0x2c, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x49, 0x4f, 0x12, 0x20, 0x0a, 0x05,
	0x50, 0x61, 0x72, 0x73, 0x65, 0x12, 0x09, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x0a, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x42, 0x0f,
	0x5a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_prase_proto_rawDescOnce sync.Once
	file_proto_prase_proto_rawDescData = file_proto_prase_proto_rawDesc
)

func file_proto_prase_proto_rawDescGZIP() []byte {
	file_proto_prase_proto_rawDescOnce.Do(func() {
		file_proto_prase_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_prase_proto_rawDescData)
	})
	return file_proto_prase_proto_rawDescData
}

var file_proto_prase_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_prase_proto_goTypes = []interface{}{
	(*Input)(nil),  // 0: io.Input
	(*Output)(nil), // 1: io.Output
}
var file_proto_prase_proto_depIdxs = []int32{
	0, // 0: io.ParserIO.Parse:input_type -> io.Input
	1, // 1: io.ParserIO.Parse:output_type -> io.Output
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_prase_proto_init() }
func file_proto_prase_proto_init() {
	if File_proto_prase_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_prase_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Input); i {
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
		file_proto_prase_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Output); i {
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
			RawDescriptor: file_proto_prase_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_prase_proto_goTypes,
		DependencyIndexes: file_proto_prase_proto_depIdxs,
		MessageInfos:      file_proto_prase_proto_msgTypes,
	}.Build()
	File_proto_prase_proto = out.File
	file_proto_prase_proto_rawDesc = nil
	file_proto_prase_proto_goTypes = nil
	file_proto_prase_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ParserIOClient is the client API for ParserIO service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ParserIOClient interface {
	Parse(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Output, error)
}

type parserIOClient struct {
	cc grpc.ClientConnInterface
}

func NewParserIOClient(cc grpc.ClientConnInterface) ParserIOClient {
	return &parserIOClient{cc}
}

func (c *parserIOClient) Parse(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Output, error) {
	out := new(Output)
	err := c.cc.Invoke(ctx, "/io.ParserIO/Parse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParserIOServer is the server API for ParserIO service.
type ParserIOServer interface {
	Parse(context.Context, *Input) (*Output, error)
}

// UnimplementedParserIOServer can be embedded to have forward compatible implementations.
type UnimplementedParserIOServer struct {
}

func (*UnimplementedParserIOServer) Parse(context.Context, *Input) (*Output, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Parse not implemented")
}

func RegisterParserIOServer(s *grpc.Server, srv ParserIOServer) {
	s.RegisterService(&_ParserIO_serviceDesc, srv)
}

func _ParserIO_Parse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParserIOServer).Parse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.ParserIO/Parse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParserIOServer).Parse(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

var _ParserIO_serviceDesc = grpc.ServiceDesc{
	ServiceName: "io.ParserIO",
	HandlerType: (*ParserIOServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Parse",
			Handler:    _ParserIO_Parse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/prase.proto",
}