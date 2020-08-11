// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: protos/rusprofile.proto

package rusprofile

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Counterparty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	INN   uint64 `protobuf:"varint,1,opt,name=INN,proto3" json:"INN,omitempty"`
	KPP   uint64 `protobuf:"varint,2,opt,name=KPP,proto3" json:"KPP,omitempty"`
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Head  string `protobuf:"bytes,4,opt,name=head,proto3" json:"head,omitempty"`
}

func (x *Counterparty) Reset() {
	*x = Counterparty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_rusprofile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Counterparty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Counterparty) ProtoMessage() {}

func (x *Counterparty) ProtoReflect() protoreflect.Message {
	mi := &file_protos_rusprofile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Counterparty.ProtoReflect.Descriptor instead.
func (*Counterparty) Descriptor() ([]byte, []int) {
	return file_protos_rusprofile_proto_rawDescGZIP(), []int{0}
}

func (x *Counterparty) GetINN() uint64 {
	if x != nil {
		return x.INN
	}
	return 0
}

func (x *Counterparty) GetKPP() uint64 {
	if x != nil {
		return x.KPP
	}
	return 0
}

func (x *Counterparty) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Counterparty) GetHead() string {
	if x != nil {
		return x.Head
	}
	return ""
}

type INN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *INN) Reset() {
	*x = INN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_rusprofile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *INN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*INN) ProtoMessage() {}

func (x *INN) ProtoReflect() protoreflect.Message {
	mi := &file_protos_rusprofile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use INN.ProtoReflect.Descriptor instead.
func (*INN) Descriptor() ([]byte, []int) {
	return file_protos_rusprofile_proto_rawDescGZIP(), []int{1}
}

func (x *INN) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_protos_rusprofile_proto protoreflect.FileDescriptor

var file_protos_rusprofile_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x75, 0x73, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72, 0x75, 0x73, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x0c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x4e, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x03, 0x49, 0x4e, 0x4e, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x50, 0x50, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x4b, 0x50, 0x50, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x65, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x65, 0x61,
	0x64, 0x22, 0x15, 0x0a, 0x03, 0x49, 0x4e, 0x4e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x32, 0x5c, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x11, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x61, 0x72, 0x74, 0x79, 0x12, 0x0f, 0x2e, 0x72, 0x75, 0x73, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x49, 0x4e, 0x4e, 0x1a, 0x18, 0x2e, 0x72, 0x75, 0x73, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x69, 0x6e,
	0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_rusprofile_proto_rawDescOnce sync.Once
	file_protos_rusprofile_proto_rawDescData = file_protos_rusprofile_proto_rawDesc
)

func file_protos_rusprofile_proto_rawDescGZIP() []byte {
	file_protos_rusprofile_proto_rawDescOnce.Do(func() {
		file_protos_rusprofile_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_rusprofile_proto_rawDescData)
	})
	return file_protos_rusprofile_proto_rawDescData
}

var file_protos_rusprofile_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_rusprofile_proto_goTypes = []interface{}{
	(*Counterparty)(nil), // 0: rusprofile.Counterparty
	(*INN)(nil),          // 1: rusprofile.INN
}
var file_protos_rusprofile_proto_depIdxs = []int32{
	1, // 0: rusprofile.Service.FetchCounterparty:input_type -> rusprofile.INN
	0, // 1: rusprofile.Service.FetchCounterparty:output_type -> rusprofile.Counterparty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_rusprofile_proto_init() }
func file_protos_rusprofile_proto_init() {
	if File_protos_rusprofile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_rusprofile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Counterparty); i {
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
		file_protos_rusprofile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*INN); i {
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
			RawDescriptor: file_protos_rusprofile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_rusprofile_proto_goTypes,
		DependencyIndexes: file_protos_rusprofile_proto_depIdxs,
		MessageInfos:      file_protos_rusprofile_proto_msgTypes,
	}.Build()
	File_protos_rusprofile_proto = out.File
	file_protos_rusprofile_proto_rawDesc = nil
	file_protos_rusprofile_proto_goTypes = nil
	file_protos_rusprofile_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	FetchCounterparty(ctx context.Context, in *INN, opts ...grpc.CallOption) (*Counterparty, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) FetchCounterparty(ctx context.Context, in *INN, opts ...grpc.CallOption) (*Counterparty, error) {
	out := new(Counterparty)
	err := c.cc.Invoke(ctx, "/rusprofile.Service/FetchCounterparty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	FetchCounterparty(context.Context, *INN) (*Counterparty, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) FetchCounterparty(context.Context, *INN) (*Counterparty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchCounterparty not implemented")
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_FetchCounterparty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(INN)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).FetchCounterparty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rusprofile.Service/FetchCounterparty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).FetchCounterparty(ctx, req.(*INN))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rusprofile.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchCounterparty",
			Handler:    _Service_FetchCounterparty_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/rusprofile.proto",
}
