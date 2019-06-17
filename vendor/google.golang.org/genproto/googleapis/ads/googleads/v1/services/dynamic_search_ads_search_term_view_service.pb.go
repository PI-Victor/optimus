// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/dynamic_search_ads_search_term_view_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v1/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for
// [DynamicSearchAdsSearchTermViewService.GetDynamicSearchAdsSearchTermView][google.ads.googleads.v1.services.DynamicSearchAdsSearchTermViewService.GetDynamicSearchAdsSearchTermView].
type GetDynamicSearchAdsSearchTermViewRequest struct {
	// The resource name of the dynamic search ads search term view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDynamicSearchAdsSearchTermViewRequest) Reset() {
	*m = GetDynamicSearchAdsSearchTermViewRequest{}
}
func (m *GetDynamicSearchAdsSearchTermViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetDynamicSearchAdsSearchTermViewRequest) ProtoMessage()    {}
func (*GetDynamicSearchAdsSearchTermViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dynamic_search_ads_search_term_view_service_818af03bdd020ada, []int{0}
}
func (m *GetDynamicSearchAdsSearchTermViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDynamicSearchAdsSearchTermViewRequest.Unmarshal(m, b)
}
func (m *GetDynamicSearchAdsSearchTermViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDynamicSearchAdsSearchTermViewRequest.Marshal(b, m, deterministic)
}
func (dst *GetDynamicSearchAdsSearchTermViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDynamicSearchAdsSearchTermViewRequest.Merge(dst, src)
}
func (m *GetDynamicSearchAdsSearchTermViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetDynamicSearchAdsSearchTermViewRequest.Size(m)
}
func (m *GetDynamicSearchAdsSearchTermViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDynamicSearchAdsSearchTermViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDynamicSearchAdsSearchTermViewRequest proto.InternalMessageInfo

func (m *GetDynamicSearchAdsSearchTermViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetDynamicSearchAdsSearchTermViewRequest)(nil), "google.ads.googleads.v1.services.GetDynamicSearchAdsSearchTermViewRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DynamicSearchAdsSearchTermViewServiceClient is the client API for DynamicSearchAdsSearchTermViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DynamicSearchAdsSearchTermViewServiceClient interface {
	// Returns the requested dynamic search ads search term view in full detail.
	GetDynamicSearchAdsSearchTermView(ctx context.Context, in *GetDynamicSearchAdsSearchTermViewRequest, opts ...grpc.CallOption) (*resources.DynamicSearchAdsSearchTermView, error)
}

type dynamicSearchAdsSearchTermViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewDynamicSearchAdsSearchTermViewServiceClient(cc *grpc.ClientConn) DynamicSearchAdsSearchTermViewServiceClient {
	return &dynamicSearchAdsSearchTermViewServiceClient{cc}
}

func (c *dynamicSearchAdsSearchTermViewServiceClient) GetDynamicSearchAdsSearchTermView(ctx context.Context, in *GetDynamicSearchAdsSearchTermViewRequest, opts ...grpc.CallOption) (*resources.DynamicSearchAdsSearchTermView, error) {
	out := new(resources.DynamicSearchAdsSearchTermView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.DynamicSearchAdsSearchTermViewService/GetDynamicSearchAdsSearchTermView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DynamicSearchAdsSearchTermViewServiceServer is the server API for DynamicSearchAdsSearchTermViewService service.
type DynamicSearchAdsSearchTermViewServiceServer interface {
	// Returns the requested dynamic search ads search term view in full detail.
	GetDynamicSearchAdsSearchTermView(context.Context, *GetDynamicSearchAdsSearchTermViewRequest) (*resources.DynamicSearchAdsSearchTermView, error)
}

func RegisterDynamicSearchAdsSearchTermViewServiceServer(s *grpc.Server, srv DynamicSearchAdsSearchTermViewServiceServer) {
	s.RegisterService(&_DynamicSearchAdsSearchTermViewService_serviceDesc, srv)
}

func _DynamicSearchAdsSearchTermViewService_GetDynamicSearchAdsSearchTermView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDynamicSearchAdsSearchTermViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DynamicSearchAdsSearchTermViewServiceServer).GetDynamicSearchAdsSearchTermView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.DynamicSearchAdsSearchTermViewService/GetDynamicSearchAdsSearchTermView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DynamicSearchAdsSearchTermViewServiceServer).GetDynamicSearchAdsSearchTermView(ctx, req.(*GetDynamicSearchAdsSearchTermViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DynamicSearchAdsSearchTermViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.DynamicSearchAdsSearchTermViewService",
	HandlerType: (*DynamicSearchAdsSearchTermViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDynamicSearchAdsSearchTermView",
			Handler:    _DynamicSearchAdsSearchTermViewService_GetDynamicSearchAdsSearchTermView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/dynamic_search_ads_search_term_view_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/dynamic_search_ads_search_term_view_service.proto", fileDescriptor_dynamic_search_ads_search_term_view_service_818af03bdd020ada)
}

var fileDescriptor_dynamic_search_ads_search_term_view_service_818af03bdd020ada = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x6a, 0xdb, 0x40,
	0x14, 0x85, 0x91, 0x0a, 0x85, 0x8a, 0x76, 0xa3, 0x55, 0x31, 0x5d, 0xb8, 0xae, 0x0b, 0xc6, 0x8b,
	0x19, 0xd4, 0xee, 0xa6, 0x74, 0x31, 0x6e, 0xc1, 0xfd, 0x81, 0xd6, 0xd8, 0x41, 0x8b, 0x20, 0x10,
	0x13, 0xcd, 0x45, 0x11, 0x58, 0x1a, 0x67, 0xae, 0x2c, 0x13, 0x42, 0x36, 0x79, 0x83, 0x90, 0x37,
	0xc8, 0x26, 0x90, 0x47, 0xc9, 0x36, 0xaf, 0x90, 0x55, 0xde, 0x21, 0x10, 0xe4, 0xf1, 0x08, 0xb2,
	0x90, 0xed, 0xdd, 0x61, 0xe6, 0xf0, 0x9d, 0xb9, 0xf7, 0x8c, 0x37, 0x4d, 0x95, 0x4a, 0xe7, 0x40,
	0x85, 0x44, 0x6a, 0x64, 0xad, 0xaa, 0x80, 0x22, 0xe8, 0x2a, 0x4b, 0x00, 0xa9, 0x3c, 0x2d, 0x44,
	0x9e, 0x25, 0x31, 0x82, 0xd0, 0xc9, 0x71, 0x2c, 0x24, 0x5a, 0x59, 0x82, 0xce, 0xe3, 0x2a, 0x83,
	0x55, 0xbc, 0x31, 0x93, 0x85, 0x56, 0xa5, 0xf2, 0xbb, 0x06, 0x44, 0x84, 0x44, 0xd2, 0x30, 0x49,
	0x15, 0x10, 0xcb, 0xec, 0xfc, 0x6d, 0x4b, 0xd5, 0x80, 0x6a, 0xa9, 0xf7, 0x8c, 0x35, 0x71, 0x9d,
	0x0f, 0x16, 0xb6, 0xc8, 0xa8, 0x28, 0x0a, 0x55, 0x8a, 0x32, 0x53, 0x05, 0x9a, 0xdb, 0xde, 0x7f,
	0x6f, 0x30, 0x86, 0xf2, 0xa7, 0xa1, 0xcd, 0xd6, 0x04, 0x2e, 0xd1, 0x88, 0x03, 0xd0, 0x79, 0x98,
	0xc1, 0x6a, 0x0a, 0x27, 0x4b, 0xc0, 0xd2, 0xff, 0xe4, 0xbd, 0xb3, 0x0f, 0x88, 0x0b, 0x91, 0xc3,
	0x7b, 0xa7, 0xeb, 0x0c, 0xde, 0x4c, 0xdf, 0xda, 0xc3, 0x7f, 0x22, 0x87, 0x2f, 0x37, 0xae, 0xf7,
	0x79, 0x3b, 0x6e, 0x66, 0xc6, 0xf4, 0x9f, 0x1c, 0xef, 0xe3, 0xce, 0x6c, 0xff, 0x0f, 0xd9, 0xb5,
	0x2e, 0xb2, 0xef, 0x00, 0x1d, 0xde, 0xca, 0x6a, 0x16, 0x4b, 0xb6, 0x93, 0x7a, 0xbf, 0x2f, 0xee,
	0x1f, 0xae, 0xdc, 0x1f, 0x3e, 0xaf, 0xeb, 0x38, 0x7b, 0xb1, 0x8e, 0xef, 0xc9, 0x12, 0x4b, 0x95,
	0x83, 0x46, 0x3a, 0xb4, 0xfd, 0xb4, 0x60, 0x90, 0x0e, 0xcf, 0x47, 0x97, 0xae, 0xd7, 0x4f, 0x54,
	0xbe, 0x73, 0xbe, 0xd1, 0x70, 0xaf, 0x7d, 0x4e, 0xea, 0x3e, 0x27, 0xce, 0xe1, 0xaf, 0x0d, 0x2f,
	0x55, 0x73, 0x51, 0xa4, 0x44, 0xe9, 0x94, 0xa6, 0x50, 0xac, 0xdb, 0xb6, 0x9f, 0x69, 0x91, 0x61,
	0xfb, 0x8f, 0xfe, 0x66, 0xc5, 0xb5, 0xfb, 0x6a, 0xcc, 0xf9, 0xad, 0xdb, 0x1d, 0x1b, 0x20, 0x97,
	0x48, 0x8c, 0xac, 0x55, 0x18, 0x90, 0x4d, 0x30, 0xde, 0x59, 0x4b, 0xc4, 0x25, 0x46, 0x8d, 0x25,
	0x0a, 0x83, 0xc8, 0x5a, 0x1e, 0xdd, 0xbe, 0x39, 0x67, 0x8c, 0x4b, 0x64, 0xac, 0x31, 0x31, 0x16,
	0x06, 0x8c, 0x59, 0xdb, 0xd1, 0xeb, 0xf5, 0x3b, 0xbf, 0x3e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x0e,
	0x7a, 0x64, 0x1d, 0x78, 0x03, 0x00, 0x00,
}
