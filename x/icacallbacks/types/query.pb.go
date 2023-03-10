// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stayking/icacallbacks/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a85d006041c036b7, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a85d006041c036b7, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetCallbackDataRequest struct {
	CallbackKey string `protobuf:"bytes,1,opt,name=callback_key,json=callbackKey,proto3" json:"callback_key,omitempty"`
}

func (m *QueryGetCallbackDataRequest) Reset()         { *m = QueryGetCallbackDataRequest{} }
func (m *QueryGetCallbackDataRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetCallbackDataRequest) ProtoMessage()    {}
func (*QueryGetCallbackDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a85d006041c036b7, []int{2}
}
func (m *QueryGetCallbackDataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCallbackDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCallbackDataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCallbackDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCallbackDataRequest.Merge(m, src)
}
func (m *QueryGetCallbackDataRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCallbackDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCallbackDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCallbackDataRequest proto.InternalMessageInfo

func (m *QueryGetCallbackDataRequest) GetCallbackKey() string {
	if m != nil {
		return m.CallbackKey
	}
	return ""
}

type QueryGetCallbackDataResponse struct {
	CallbackData CallbackData `protobuf:"bytes,1,opt,name=callback_data,json=callbackData,proto3" json:"callback_data"`
}

func (m *QueryGetCallbackDataResponse) Reset()         { *m = QueryGetCallbackDataResponse{} }
func (m *QueryGetCallbackDataResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetCallbackDataResponse) ProtoMessage()    {}
func (*QueryGetCallbackDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a85d006041c036b7, []int{3}
}
func (m *QueryGetCallbackDataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCallbackDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCallbackDataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCallbackDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCallbackDataResponse.Merge(m, src)
}
func (m *QueryGetCallbackDataResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCallbackDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCallbackDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCallbackDataResponse proto.InternalMessageInfo

func (m *QueryGetCallbackDataResponse) GetCallbackData() CallbackData {
	if m != nil {
		return m.CallbackData
	}
	return CallbackData{}
}

type QueryAllCallbackDataRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllCallbackDataRequest) Reset()         { *m = QueryAllCallbackDataRequest{} }
func (m *QueryAllCallbackDataRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllCallbackDataRequest) ProtoMessage()    {}
func (*QueryAllCallbackDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a85d006041c036b7, []int{4}
}
func (m *QueryAllCallbackDataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCallbackDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCallbackDataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCallbackDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCallbackDataRequest.Merge(m, src)
}
func (m *QueryAllCallbackDataRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCallbackDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCallbackDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCallbackDataRequest proto.InternalMessageInfo

func (m *QueryAllCallbackDataRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllCallbackDataResponse struct {
	CallbackData []CallbackData      `protobuf:"bytes,1,rep,name=callback_data,json=callbackData,proto3" json:"callback_data"`
	Pagination   *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllCallbackDataResponse) Reset()         { *m = QueryAllCallbackDataResponse{} }
func (m *QueryAllCallbackDataResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllCallbackDataResponse) ProtoMessage()    {}
func (*QueryAllCallbackDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a85d006041c036b7, []int{5}
}
func (m *QueryAllCallbackDataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCallbackDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCallbackDataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCallbackDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCallbackDataResponse.Merge(m, src)
}
func (m *QueryAllCallbackDataResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCallbackDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCallbackDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCallbackDataResponse proto.InternalMessageInfo

func (m *QueryAllCallbackDataResponse) GetCallbackData() []CallbackData {
	if m != nil {
		return m.CallbackData
	}
	return nil
}

func (m *QueryAllCallbackDataResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "stayking.icacallbacks.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "stayking.icacallbacks.QueryParamsResponse")
	proto.RegisterType((*QueryGetCallbackDataRequest)(nil), "stayking.icacallbacks.QueryGetCallbackDataRequest")
	proto.RegisterType((*QueryGetCallbackDataResponse)(nil), "stayking.icacallbacks.QueryGetCallbackDataResponse")
	proto.RegisterType((*QueryAllCallbackDataRequest)(nil), "stayking.icacallbacks.QueryAllCallbackDataRequest")
	proto.RegisterType((*QueryAllCallbackDataResponse)(nil), "stayking.icacallbacks.QueryAllCallbackDataResponse")
}

func init() { proto.RegisterFile("stayking/icacallbacks/query.proto", fileDescriptor_a85d006041c036b7) }

var fileDescriptor_a85d006041c036b7 = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xc1, 0x6a, 0x13, 0x41,
	0x18, 0xc7, 0x33, 0xad, 0x06, 0x9c, 0x56, 0x84, 0xb1, 0x82, 0xc4, 0x76, 0x6b, 0x47, 0xad, 0xb6,
	0xc8, 0x8c, 0x4d, 0xf0, 0xd4, 0x8b, 0xad, 0x62, 0x41, 0x41, 0xea, 0x1e, 0xbd, 0xc8, 0xb7, 0xeb,
	0x30, 0x5d, 0xb2, 0xd9, 0xd9, 0x66, 0x26, 0xe2, 0x22, 0x5e, 0x3c, 0x78, 0x16, 0x7c, 0x09, 0x1f,
	0x40, 0xf0, 0x15, 0x7a, 0x2c, 0x78, 0xf1, 0x24, 0x92, 0xf8, 0x04, 0x3e, 0x81, 0x64, 0x66, 0xd2,
	0x6e, 0xe8, 0x26, 0x46, 0xbc, 0x2d, 0x93, 0xff, 0xf7, 0xff, 0xff, 0xbe, 0xf9, 0xbe, 0x09, 0x5e,
	0xd3, 0x06, 0x8a, 0x76, 0x92, 0x49, 0x9e, 0xc4, 0x10, 0x43, 0x9a, 0x46, 0x10, 0xb7, 0x35, 0x3f,
	0xec, 0x89, 0x6e, 0xc1, 0xf2, 0xae, 0x32, 0x8a, 0x5c, 0x19, 0x49, 0x58, 0x59, 0xd2, 0x58, 0x92,
	0x4a, 0x2a, 0xab, 0xe0, 0xc3, 0x2f, 0x27, 0x6e, 0x2c, 0x4b, 0xa5, 0x64, 0x2a, 0x38, 0xe4, 0x09,
	0x87, 0x2c, 0x53, 0x06, 0x4c, 0xa2, 0x32, 0xed, 0x7f, 0xdd, 0x8c, 0x95, 0xee, 0x28, 0xcd, 0x23,
	0xd0, 0xc2, 0x65, 0xf0, 0xd7, 0x5b, 0x91, 0x30, 0xb0, 0xc5, 0x73, 0x90, 0x49, 0x66, 0xc5, 0x5e,
	0x4b, 0xab, 0xc9, 0x72, 0xe8, 0x42, 0x67, 0xe4, 0xb7, 0x51, 0xad, 0x19, 0x7d, 0xbd, 0x7c, 0x05,
	0x06, 0x9c, 0x94, 0x2e, 0x61, 0xf2, 0x7c, 0x18, 0xb8, 0x6f, 0xeb, 0x43, 0x71, 0xd8, 0x13, 0xda,
	0xd0, 0x10, 0x5f, 0x1e, 0x3b, 0xd5, 0xb9, 0xca, 0xb4, 0x20, 0xdb, 0xb8, 0xee, 0x72, 0xae, 0xa2,
	0xeb, 0xe8, 0xce, 0x42, 0x73, 0x85, 0x55, 0xde, 0x01, 0x73, 0x65, 0xbb, 0xe7, 0x8e, 0x7e, 0xac,
	0xd6, 0x42, 0x5f, 0x42, 0x1f, 0xe0, 0x6b, 0xd6, 0x73, 0x4f, 0x98, 0x87, 0x5e, 0xf9, 0x08, 0x0c,
	0xf8, 0x48, 0xb2, 0x86, 0x17, 0x4f, 0xf8, 0xda, 0xa2, 0xb0, 0x09, 0x17, 0xc2, 0x85, 0xd1, 0xd9,
	0x53, 0x51, 0xd0, 0x0c, 0x2f, 0x57, 0x3b, 0x78, 0xbc, 0x67, 0xf8, 0xe2, 0x58, 0x8b, 0x9e, 0xf2,
	0xc6, 0x04, 0xca, 0xb2, 0x87, 0x67, 0x3d, 0x41, 0x18, 0x9e, 0x51, 0xe1, 0x89, 0x77, 0xd2, 0xb4,
	0x8a, 0xf8, 0x31, 0xc6, 0xa7, 0xd3, 0xf1, 0x59, 0xeb, 0xcc, 0x8d, 0x92, 0x0d, 0x47, 0xc9, 0xdc,
	0xba, 0xf8, 0x51, 0xb2, 0x7d, 0x90, 0xc2, 0xd7, 0x86, 0xa5, 0x4a, 0xfa, 0x15, 0xf9, 0xbe, 0xce,
	0xe4, 0x4c, 0xee, 0x6b, 0xfe, 0x3f, 0xfa, 0x22, 0x7b, 0x63, 0xe0, 0x73, 0x16, 0xfc, 0xf6, 0x5f,
	0xc1, 0x1d, 0x4c, 0x99, 0xbc, 0xf9, 0x7b, 0x1e, 0x9f, 0xb7, 0xe4, 0xe4, 0x03, 0xc2, 0x75, 0x37,
	0x75, 0xb2, 0x31, 0x01, 0xeb, 0xec, 0x9a, 0x35, 0x36, 0x67, 0x91, 0xba, 0x5c, 0x7a, 0xeb, 0xfd,
	0xb7, 0x5f, 0x9f, 0xe6, 0x56, 0xc9, 0x0a, 0x9f, 0xf6, 0x00, 0xc8, 0x17, 0x84, 0x17, 0xcb, 0x17,
	0x40, 0x9a, 0xd3, 0x32, 0xaa, 0x77, 0xb1, 0xd1, 0xfa, 0xa7, 0x1a, 0x0f, 0xb8, 0x6d, 0x01, 0xef,
	0x93, 0x16, 0x9f, 0xe1, 0xf5, 0xf1, 0xb7, 0xe5, 0x65, 0x7f, 0x47, 0x3e, 0x23, 0x7c, 0xa9, 0xec,
	0xba, 0x93, 0xa6, 0xd3, 0xc9, 0xab, 0x77, 0x72, 0x3a, 0xf9, 0x84, 0xfd, 0xa2, 0x77, 0x2d, 0xf9,
	0x3a, 0xb9, 0x39, 0x0b, 0xf9, 0xee, 0x93, 0xa3, 0x7e, 0x80, 0x8e, 0xfb, 0x01, 0xfa, 0xd9, 0x0f,
	0xd0, 0xc7, 0x41, 0x50, 0x3b, 0x1e, 0x04, 0xb5, 0xef, 0x83, 0xa0, 0xf6, 0xe2, 0x9e, 0x4c, 0xcc,
	0x41, 0x2f, 0x62, 0xb1, 0xea, 0x70, 0xad, 0xd4, 0x81, 0x4a, 0xd4, 0xa9, 0xe3, 0x9b, 0x71, 0x4f,
	0x53, 0xe4, 0x42, 0x47, 0x75, 0xfb, 0x27, 0xd4, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x8e,
	0x92, 0x0a, 0x6f, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a CallbackData by index.
	CallbackData(ctx context.Context, in *QueryGetCallbackDataRequest, opts ...grpc.CallOption) (*QueryGetCallbackDataResponse, error)
	// Queries a list of CallbackData items.
	CallbackDataAll(ctx context.Context, in *QueryAllCallbackDataRequest, opts ...grpc.CallOption) (*QueryAllCallbackDataResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/stayking.icacallbacks.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CallbackData(ctx context.Context, in *QueryGetCallbackDataRequest, opts ...grpc.CallOption) (*QueryGetCallbackDataResponse, error) {
	out := new(QueryGetCallbackDataResponse)
	err := c.cc.Invoke(ctx, "/stayking.icacallbacks.Query/CallbackData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CallbackDataAll(ctx context.Context, in *QueryAllCallbackDataRequest, opts ...grpc.CallOption) (*QueryAllCallbackDataResponse, error) {
	out := new(QueryAllCallbackDataResponse)
	err := c.cc.Invoke(ctx, "/stayking.icacallbacks.Query/CallbackDataAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a CallbackData by index.
	CallbackData(context.Context, *QueryGetCallbackDataRequest) (*QueryGetCallbackDataResponse, error)
	// Queries a list of CallbackData items.
	CallbackDataAll(context.Context, *QueryAllCallbackDataRequest) (*QueryAllCallbackDataResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) CallbackData(ctx context.Context, req *QueryGetCallbackDataRequest) (*QueryGetCallbackDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallbackData not implemented")
}
func (*UnimplementedQueryServer) CallbackDataAll(ctx context.Context, req *QueryAllCallbackDataRequest) (*QueryAllCallbackDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallbackDataAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stayking.icacallbacks.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CallbackData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetCallbackDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CallbackData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stayking.icacallbacks.Query/CallbackData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CallbackData(ctx, req.(*QueryGetCallbackDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CallbackDataAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllCallbackDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CallbackDataAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stayking.icacallbacks.Query/CallbackDataAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CallbackDataAll(ctx, req.(*QueryAllCallbackDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stayking.icacallbacks.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "CallbackData",
			Handler:    _Query_CallbackData_Handler,
		},
		{
			MethodName: "CallbackDataAll",
			Handler:    _Query_CallbackDataAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stayking/icacallbacks/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetCallbackDataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCallbackDataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCallbackDataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CallbackKey) > 0 {
		i -= len(m.CallbackKey)
		copy(dAtA[i:], m.CallbackKey)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.CallbackKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetCallbackDataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCallbackDataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCallbackDataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.CallbackData.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllCallbackDataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCallbackDataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCallbackDataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllCallbackDataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCallbackDataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCallbackDataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.CallbackData) > 0 {
		for iNdEx := len(m.CallbackData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CallbackData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetCallbackDataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CallbackKey)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetCallbackDataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CallbackData.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllCallbackDataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllCallbackDataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CallbackData) > 0 {
		for _, e := range m.CallbackData {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetCallbackDataRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetCallbackDataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCallbackDataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallbackKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallbackKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetCallbackDataResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetCallbackDataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCallbackDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallbackData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CallbackData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllCallbackDataRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllCallbackDataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCallbackDataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllCallbackDataResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllCallbackDataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCallbackDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallbackData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallbackData = append(m.CallbackData, CallbackData{})
			if err := m.CallbackData[len(m.CallbackData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
