// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/topic_constant_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for [TopicConstantService.GetTopicConstant][google.ads.googleads.v2.services.TopicConstantService.GetTopicConstant].
type GetTopicConstantRequest struct {
	// Required. Resource name of the Topic to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTopicConstantRequest) Reset()         { *m = GetTopicConstantRequest{} }
func (m *GetTopicConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetTopicConstantRequest) ProtoMessage()    {}
func (*GetTopicConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fa68112b91a9250, []int{0}
}

func (m *GetTopicConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopicConstantRequest.Unmarshal(m, b)
}
func (m *GetTopicConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopicConstantRequest.Marshal(b, m, deterministic)
}
func (m *GetTopicConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopicConstantRequest.Merge(m, src)
}
func (m *GetTopicConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetTopicConstantRequest.Size(m)
}
func (m *GetTopicConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopicConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopicConstantRequest proto.InternalMessageInfo

func (m *GetTopicConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetTopicConstantRequest)(nil), "google.ads.googleads.v2.services.GetTopicConstantRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/topic_constant_service.proto", fileDescriptor_8fa68112b91a9250)
}

var fileDescriptor_8fa68112b91a9250 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x25, 0x11, 0x04, 0x83, 0x82, 0x04, 0xa1, 0x6d, 0x14, 0x2c, 0xa5, 0x94, 0xe2, 0x62, 0x46,
	0x22, 0x08, 0x8e, 0x76, 0x31, 0x75, 0x51, 0x57, 0x52, 0xac, 0x74, 0x21, 0x81, 0x30, 0x4d, 0xc6,
	0x38, 0x90, 0xcc, 0xc4, 0xcc, 0x34, 0x1b, 0x11, 0xc4, 0x5f, 0xf0, 0x0f, 0x5c, 0xfa, 0x0f, 0xfe,
	0x40, 0xb7, 0xee, 0x5c, 0xb9, 0x70, 0xe5, 0xd6, 0x9d, 0x2b, 0x49, 0x27, 0x93, 0x26, 0xef, 0xbd,
	0xd0, 0xdd, 0x61, 0xce, 0x39, 0xf7, 0x9e, 0x7b, 0xef, 0x38, 0x8b, 0x44, 0x88, 0x24, 0xa5, 0x90,
	0xc4, 0x12, 0x6a, 0x58, 0xa1, 0xd2, 0x87, 0x92, 0x16, 0x25, 0x8b, 0xa8, 0x84, 0x4a, 0xe4, 0x2c,
	0x0a, 0x23, 0xc1, 0xa5, 0x22, 0x5c, 0x85, 0xf5, 0x3b, 0xc8, 0x0b, 0xa1, 0x84, 0x3b, 0xd6, 0x1e,
	0x40, 0x62, 0x09, 0x1a, 0x3b, 0x28, 0x7d, 0x60, 0xec, 0xde, 0xe3, 0xbe, 0x06, 0x05, 0x95, 0x62,
	0x5f, 0x5c, 0xee, 0xa0, 0x2b, 0x7b, 0xf7, 0x8c, 0x2f, 0x67, 0x90, 0x70, 0x2e, 0x14, 0x51, 0x4c,
	0x70, 0x59, 0xb3, 0x83, 0x16, 0x1b, 0xa5, 0x8c, 0x36, 0xb6, 0xfb, 0x2d, 0xe2, 0x2d, 0xa3, 0x69,
	0x1c, 0xee, 0xe8, 0x3b, 0x52, 0x32, 0x51, 0xd4, 0x82, 0x51, 0x4b, 0x60, 0x22, 0x68, 0x6a, 0xc2,
	0x9d, 0xc1, 0x8a, 0xaa, 0xd7, 0x55, 0x9a, 0xe7, 0x75, 0x98, 0x57, 0xf4, 0xfd, 0x9e, 0x4a, 0xe5,
	0x6e, 0x9c, 0x5b, 0x46, 0x1c, 0x72, 0x92, 0xd1, 0xa1, 0x35, 0xb6, 0xe6, 0x37, 0x96, 0xe0, 0x17,
	0xb6, 0xff, 0xe1, 0xb9, 0x33, 0x3b, 0xcd, 0x5e, 0xa3, 0x9c, 0x49, 0x10, 0x89, 0x0c, 0x76, 0xab,
	0xdd, 0x34, 0x45, 0x5e, 0x92, 0x8c, 0xfa, 0x7f, 0x2d, 0xe7, 0x4e, 0x87, 0xdf, 0xe8, 0xa5, 0xb9,
	0xdf, 0x2d, 0xe7, 0xf6, 0xc5, 0x24, 0xee, 0x13, 0x70, 0x6e, 0xd7, 0xa0, 0x27, 0xbd, 0xf7, 0xb0,
	0xd7, 0xda, 0x1c, 0x01, 0x74, 0x8c, 0x93, 0x67, 0x3f, 0x71, 0x77, 0xe0, 0xcf, 0x3f, 0x7e, 0x7f,
	0xb1, 0x67, 0xee, 0xb4, 0xba, 0xdc, 0x87, 0x0e, 0xb3, 0x50, 0x6d, 0xa7, 0x84, 0x0f, 0x3e, 0x7a,
	0x77, 0x0f, 0x78, 0xd8, 0xb7, 0x91, 0xe5, 0x27, 0xdb, 0x99, 0x46, 0x22, 0x3b, 0x3b, 0xcd, 0x72,
	0x74, 0xd5, 0x6e, 0xd6, 0xd5, 0xa5, 0xd6, 0xd6, 0x9b, 0x17, 0xb5, 0x3d, 0x11, 0x29, 0xe1, 0x09,
	0x10, 0x45, 0x02, 0x13, 0xca, 0x8f, 0x77, 0x84, 0xa7, 0x86, 0xfd, 0xdf, 0xfa, 0xa9, 0x01, 0x5f,
	0xed, 0x6b, 0x2b, 0x8c, 0xbf, 0xd9, 0xe3, 0x95, 0x2e, 0x88, 0x63, 0x09, 0x34, 0xac, 0xd0, 0xd6,
	0x07, 0x75, 0x63, 0x79, 0x30, 0x92, 0x00, 0xc7, 0x32, 0x68, 0x24, 0xc1, 0xd6, 0x0f, 0x8c, 0xe4,
	0x8f, 0x3d, 0xd5, 0xef, 0x08, 0xe1, 0x58, 0x22, 0xd4, 0x88, 0x10, 0xda, 0xfa, 0x08, 0x19, 0xd9,
	0xee, 0xfa, 0x31, 0xe7, 0xa3, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x20, 0xe6, 0xcd, 0xea, 0x7d,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TopicConstantServiceClient is the client API for TopicConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TopicConstantServiceClient interface {
	// Returns the requested topic constant in full detail.
	GetTopicConstant(ctx context.Context, in *GetTopicConstantRequest, opts ...grpc.CallOption) (*resources.TopicConstant, error)
}

type topicConstantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTopicConstantServiceClient(cc grpc.ClientConnInterface) TopicConstantServiceClient {
	return &topicConstantServiceClient{cc}
}

func (c *topicConstantServiceClient) GetTopicConstant(ctx context.Context, in *GetTopicConstantRequest, opts ...grpc.CallOption) (*resources.TopicConstant, error) {
	out := new(resources.TopicConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.TopicConstantService/GetTopicConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopicConstantServiceServer is the server API for TopicConstantService service.
type TopicConstantServiceServer interface {
	// Returns the requested topic constant in full detail.
	GetTopicConstant(context.Context, *GetTopicConstantRequest) (*resources.TopicConstant, error)
}

// UnimplementedTopicConstantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTopicConstantServiceServer struct {
}

func (*UnimplementedTopicConstantServiceServer) GetTopicConstant(ctx context.Context, req *GetTopicConstantRequest) (*resources.TopicConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopicConstant not implemented")
}

func RegisterTopicConstantServiceServer(s *grpc.Server, srv TopicConstantServiceServer) {
	s.RegisterService(&_TopicConstantService_serviceDesc, srv)
}

func _TopicConstantService_GetTopicConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicConstantServiceServer).GetTopicConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.TopicConstantService/GetTopicConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicConstantServiceServer).GetTopicConstant(ctx, req.(*GetTopicConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TopicConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.TopicConstantService",
	HandlerType: (*TopicConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTopicConstant",
			Handler:    _TopicConstantService_GetTopicConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/topic_constant_service.proto",
}