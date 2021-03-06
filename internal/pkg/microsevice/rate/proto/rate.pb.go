// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: rate.proto

package proto

import (
	context "context"
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

type Rating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rating  int64  `protobuf:"varint,1,opt,name=rating,proto3" json:"rating,omitempty"`
	FilmId  int64  `protobuf:"varint,2,opt,name=filmId,proto3" json:"filmId,omitempty"`
	Session string `protobuf:"bytes,3,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *Rating) Reset() {
	*x = Rating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rating) ProtoMessage() {}

func (x *Rating) ProtoReflect() protoreflect.Message {
	mi := &file_rate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rating.ProtoReflect.Descriptor instead.
func (*Rating) Descriptor() ([]byte, []int) {
	return file_rate_proto_rawDescGZIP(), []int{0}
}

func (x *Rating) GetRating() int64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Rating) GetFilmId() int64 {
	if x != nil {
		return x.FilmId
	}
	return 0
}

func (x *Rating) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

type Review struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body    string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	FilmId  int64  `protobuf:"varint,2,opt,name=filmId,proto3" json:"filmId,omitempty"`
	Session string `protobuf:"bytes,3,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *Review) Reset() {
	*x = Review{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Review) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Review) ProtoMessage() {}

func (x *Review) ProtoReflect() protoreflect.Message {
	mi := &file_rate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Review.ProtoReflect.Descriptor instead.
func (*Review) Descriptor() ([]byte, []int) {
	return file_rate_proto_rawDescGZIP(), []int{1}
}

func (x *Review) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Review) GetFilmId() int64 {
	if x != nil {
		return x.FilmId
	}
	return 0
}

func (x *Review) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

type RateEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RateEmpty) Reset() {
	*x = RateEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateEmpty) ProtoMessage() {}

func (x *RateEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_rate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateEmpty.ProtoReflect.Descriptor instead.
func (*RateEmpty) Descriptor() ([]byte, []int) {
	return file_rate_proto_rawDescGZIP(), []int{2}
}

var File_rate_proto protoreflect.FileDescriptor

var file_rate_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x6d, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x6d, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x4e, 0x0a, 0x06, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x6d, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x6d, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x0b, 0x0a, 0x09, 0x52, 0x61, 0x74, 0x65, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x32, 0x61, 0x0a, 0x04, 0x52, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x04,
	0x52, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x61, 0x74, 0x65,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x61, 0x74, 0x65,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rate_proto_rawDescOnce sync.Once
	file_rate_proto_rawDescData = file_rate_proto_rawDesc
)

func file_rate_proto_rawDescGZIP() []byte {
	file_rate_proto_rawDescOnce.Do(func() {
		file_rate_proto_rawDescData = protoimpl.X.CompressGZIP(file_rate_proto_rawDescData)
	})
	return file_rate_proto_rawDescData
}

var file_rate_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rate_proto_goTypes = []interface{}{
	(*Rating)(nil),    // 0: proto.Rating
	(*Review)(nil),    // 1: proto.Review
	(*RateEmpty)(nil), // 2: proto.RateEmpty
}
var file_rate_proto_depIdxs = []int32{
	0, // 0: proto.Rate.Rate:input_type -> proto.Rating
	1, // 1: proto.Rate.AddReview:input_type -> proto.Review
	2, // 2: proto.Rate.Rate:output_type -> proto.RateEmpty
	2, // 3: proto.Rate.AddReview:output_type -> proto.RateEmpty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rate_proto_init() }
func file_rate_proto_init() {
	if File_rate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rating); i {
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
		file_rate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Review); i {
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
		file_rate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateEmpty); i {
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
			RawDescriptor: file_rate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rate_proto_goTypes,
		DependencyIndexes: file_rate_proto_depIdxs,
		MessageInfos:      file_rate_proto_msgTypes,
	}.Build()
	File_rate_proto = out.File
	file_rate_proto_rawDesc = nil
	file_rate_proto_goTypes = nil
	file_rate_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RateClient is the client API for Rate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RateClient interface {
	Rate(ctx context.Context, in *Rating, opts ...grpc.CallOption) (*RateEmpty, error)
	AddReview(ctx context.Context, in *Review, opts ...grpc.CallOption) (*RateEmpty, error)
}

type rateClient struct {
	cc grpc.ClientConnInterface
}

func NewRateClient(cc grpc.ClientConnInterface) RateClient {
	return &rateClient{cc}
}

func (c *rateClient) Rate(ctx context.Context, in *Rating, opts ...grpc.CallOption) (*RateEmpty, error) {
	out := new(RateEmpty)
	err := c.cc.Invoke(ctx, "/proto.Rate/Rate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rateClient) AddReview(ctx context.Context, in *Review, opts ...grpc.CallOption) (*RateEmpty, error) {
	out := new(RateEmpty)
	err := c.cc.Invoke(ctx, "/proto.Rate/AddReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateServer is the server API for Rate service.
type RateServer interface {
	Rate(context.Context, *Rating) (*RateEmpty, error)
	AddReview(context.Context, *Review) (*RateEmpty, error)
}

// UnimplementedRateServer can be embedded to have forward compatible implementations.
type UnimplementedRateServer struct {
}

func (*UnimplementedRateServer) Rate(context.Context, *Rating) (*RateEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rate not implemented")
}
func (*UnimplementedRateServer) AddReview(context.Context, *Review) (*RateEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddReview not implemented")
}

func RegisterRateServer(s *grpc.Server, srv RateServer) {
	s.RegisterService(&_Rate_serviceDesc, srv)
}

func _Rate_Rate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Rating)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateServer).Rate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Rate/Rate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateServer).Rate(ctx, req.(*Rating))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rate_AddReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Review)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateServer).AddReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Rate/AddReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateServer).AddReview(ctx, req.(*Review))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Rate",
	HandlerType: (*RateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Rate",
			Handler:    _Rate_Rate_Handler,
		},
		{
			MethodName: "AddReview",
			Handler:    _Rate_AddReview_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rate.proto",
}
