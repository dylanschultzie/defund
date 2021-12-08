// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: etf/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
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

type QueryGetFundRequest struct {
	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *QueryGetFundRequest) Reset()         { *m = QueryGetFundRequest{} }
func (m *QueryGetFundRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetFundRequest) ProtoMessage()    {}
func (*QueryGetFundRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b29f902a0b14e0e, []int{0}
}
func (m *QueryGetFundRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetFundRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetFundRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetFundRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetFundRequest.Merge(m, src)
}
func (m *QueryGetFundRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetFundRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetFundRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetFundRequest proto.InternalMessageInfo

func (m *QueryGetFundRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type QueryGetFundResponse struct {
	Fund Fund `protobuf:"bytes,1,opt,name=fund,proto3" json:"fund"`
}

func (m *QueryGetFundResponse) Reset()         { *m = QueryGetFundResponse{} }
func (m *QueryGetFundResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetFundResponse) ProtoMessage()    {}
func (*QueryGetFundResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b29f902a0b14e0e, []int{1}
}
func (m *QueryGetFundResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetFundResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetFundResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetFundResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetFundResponse.Merge(m, src)
}
func (m *QueryGetFundResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetFundResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetFundResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetFundResponse proto.InternalMessageInfo

func (m *QueryGetFundResponse) GetFund() Fund {
	if m != nil {
		return m.Fund
	}
	return Fund{}
}

type QueryAllFundRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllFundRequest) Reset()         { *m = QueryAllFundRequest{} }
func (m *QueryAllFundRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllFundRequest) ProtoMessage()    {}
func (*QueryAllFundRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b29f902a0b14e0e, []int{2}
}
func (m *QueryAllFundRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllFundRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllFundRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllFundRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllFundRequest.Merge(m, src)
}
func (m *QueryAllFundRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllFundRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllFundRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllFundRequest proto.InternalMessageInfo

func (m *QueryAllFundRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllFundResponse struct {
	Fund       []Fund              `protobuf:"bytes,1,rep,name=fund,proto3" json:"fund"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllFundResponse) Reset()         { *m = QueryAllFundResponse{} }
func (m *QueryAllFundResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllFundResponse) ProtoMessage()    {}
func (*QueryAllFundResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b29f902a0b14e0e, []int{3}
}
func (m *QueryAllFundResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllFundResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllFundResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllFundResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllFundResponse.Merge(m, src)
}
func (m *QueryAllFundResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllFundResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllFundResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllFundResponse proto.InternalMessageInfo

func (m *QueryAllFundResponse) GetFund() []Fund {
	if m != nil {
		return m.Fund
	}
	return nil
}

func (m *QueryAllFundResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetFundRequest)(nil), "defundhub.defund.etf.QueryGetFundRequest")
	proto.RegisterType((*QueryGetFundResponse)(nil), "defundhub.defund.etf.QueryGetFundResponse")
	proto.RegisterType((*QueryAllFundRequest)(nil), "defundhub.defund.etf.QueryAllFundRequest")
	proto.RegisterType((*QueryAllFundResponse)(nil), "defundhub.defund.etf.QueryAllFundResponse")
}

func init() { proto.RegisterFile("etf/query.proto", fileDescriptor_9b29f902a0b14e0e) }

var fileDescriptor_9b29f902a0b14e0e = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x8a, 0x13, 0x31,
	0x1c, 0xc7, 0x27, 0xb5, 0x55, 0x8c, 0xa0, 0x10, 0xe7, 0x20, 0x43, 0x19, 0x25, 0x88, 0xd5, 0x0a,
	0x09, 0xad, 0xbe, 0x40, 0x7b, 0x68, 0x2f, 0x1e, 0x74, 0x8e, 0x82, 0x87, 0x99, 0x4e, 0x3a, 0x0e,
	0x4c, 0x93, 0x69, 0x93, 0x91, 0x16, 0xf1, 0x22, 0x5e, 0xbc, 0x09, 0xe2, 0xc5, 0x27, 0xea, 0xb1,
	0xb0, 0x97, 0x3d, 0x2d, 0x4b, 0xbb, 0x0f, 0xb2, 0x24, 0x99, 0xed, 0xb6, 0xdb, 0x6e, 0xb7, 0xb7,
	0xdf, 0x0c, 0xdf, 0x3f, 0x9f, 0x5f, 0x12, 0xf8, 0x84, 0xa9, 0x21, 0x1d, 0x17, 0x6c, 0x32, 0x23,
	0xf9, 0x44, 0x28, 0x81, 0xdc, 0x98, 0x0d, 0x0b, 0x1e, 0x7f, 0x2d, 0x22, 0x62, 0x27, 0xc2, 0xd4,
	0xd0, 0xab, 0x27, 0x42, 0x24, 0x19, 0xa3, 0x61, 0x9e, 0xd2, 0x90, 0x73, 0xa1, 0x42, 0x95, 0x0a,
	0x2e, 0xad, 0xc7, 0x6b, 0x0e, 0x84, 0x1c, 0x09, 0x49, 0xa3, 0x50, 0x32, 0x1b, 0x46, 0xbf, 0xb5,
	0x22, 0xa6, 0xc2, 0x16, 0xcd, 0xc3, 0x24, 0xe5, 0x46, 0x5c, 0x6a, 0x1f, 0xeb, 0x42, 0x93, 0x6b,
	0xbf, 0xdd, 0x44, 0x24, 0xc2, 0x8c, 0x54, 0x4f, 0xf6, 0x2f, 0x7e, 0x0b, 0x9f, 0x7e, 0xd2, 0x39,
	0x7d, 0xa6, 0x7a, 0x05, 0x8f, 0x03, 0x36, 0x2e, 0x98, 0x54, 0xc8, 0x85, 0xb5, 0x94, 0xc7, 0x6c,
	0xfa, 0x0c, 0xbc, 0x00, 0xaf, 0x1f, 0x06, 0xf6, 0x03, 0x7f, 0x80, 0xee, 0xb6, 0x58, 0xe6, 0x82,
	0x4b, 0x86, 0xde, 0xc3, 0xaa, 0x2e, 0x32, 0xe2, 0x47, 0x6d, 0x8f, 0xec, 0xdb, 0x8c, 0x68, 0x47,
	0xb7, 0x3a, 0x3f, 0x7b, 0xee, 0x04, 0x46, 0x8d, 0xbf, 0x94, 0xd5, 0x9d, 0x2c, 0xdb, 0xac, 0xee,
	0x41, 0x78, 0xbd, 0x4b, 0x19, 0xf9, 0x8a, 0xd8, 0xc5, 0x89, 0x5e, 0x9c, 0xd8, 0x53, 0x2c, 0x17,
	0x27, 0x1f, 0xc3, 0x84, 0x95, 0xde, 0x60, 0xc3, 0x89, 0xff, 0x81, 0x92, 0x76, 0x9d, 0xbf, 0x43,
	0x7b, 0xef, 0x78, 0x5a, 0xd4, 0xdf, 0xc2, 0xaa, 0x18, 0xac, 0xc6, 0x9d, 0x58, 0xb6, 0x72, 0x93,
	0xab, 0xfd, 0xbf, 0x02, 0x6b, 0x86, 0x0b, 0xfd, 0x06, 0xb0, 0xaa, 0x7b, 0xd0, 0x9b, 0xfd, 0x0c,
	0x7b, 0x2e, 0xc6, 0x6b, 0x1e, 0x23, 0xb5, 0xad, 0xb8, 0xf9, 0xf3, 0xe4, 0xe2, 0x6f, 0xe5, 0x25,
	0xc2, 0x74, 0xed, 0x29, 0x27, 0x7a, 0xf5, 0x36, 0xe8, 0x77, 0x73, 0xb3, 0x3f, 0xd0, 0x2f, 0x00,
	0x1f, 0x68, 0x73, 0x27, 0xcb, 0x0e, 0xe2, 0x6c, 0x5f, 0xd6, 0x41, 0x9c, 0x1b, 0xe7, 0x8e, 0xb1,
	0xc1, 0xa9, 0x23, 0xef, 0x76, 0x9c, 0x6e, 0x67, 0xbe, 0xf4, 0xc1, 0x62, 0xe9, 0x83, 0xf3, 0xa5,
	0x0f, 0xfe, 0xac, 0x7c, 0x67, 0xb1, 0xf2, 0x9d, 0xd3, 0x95, 0xef, 0x7c, 0x6e, 0x24, 0xa9, 0xd2,
	0x2d, 0x03, 0x31, 0xda, 0xf5, 0x4f, 0x4d, 0x82, 0x9a, 0xe5, 0x4c, 0x46, 0xf7, 0xcd, 0xc3, 0x7e,
	0x77, 0x19, 0x00, 0x00, 0xff, 0xff, 0x63, 0x1b, 0x5c, 0x4f, 0x71, 0x03, 0x00, 0x00,
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
	// Queries a fund by index.
	Fund(ctx context.Context, in *QueryGetFundRequest, opts ...grpc.CallOption) (*QueryGetFundResponse, error)
	// Queries a list of fund items.
	FundAll(ctx context.Context, in *QueryAllFundRequest, opts ...grpc.CallOption) (*QueryAllFundResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Fund(ctx context.Context, in *QueryGetFundRequest, opts ...grpc.CallOption) (*QueryGetFundResponse, error) {
	out := new(QueryGetFundResponse)
	err := c.cc.Invoke(ctx, "/defundhub.defund.etf.Query/Fund", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) FundAll(ctx context.Context, in *QueryAllFundRequest, opts ...grpc.CallOption) (*QueryAllFundResponse, error) {
	out := new(QueryAllFundResponse)
	err := c.cc.Invoke(ctx, "/defundhub.defund.etf.Query/FundAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a fund by index.
	Fund(context.Context, *QueryGetFundRequest) (*QueryGetFundResponse, error)
	// Queries a list of fund items.
	FundAll(context.Context, *QueryAllFundRequest) (*QueryAllFundResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Fund(ctx context.Context, req *QueryGetFundRequest) (*QueryGetFundResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fund not implemented")
}
func (*UnimplementedQueryServer) FundAll(ctx context.Context, req *QueryAllFundRequest) (*QueryAllFundResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FundAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Fund_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetFundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Fund(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/defundhub.defund.etf.Query/Fund",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Fund(ctx, req.(*QueryGetFundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_FundAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllFundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).FundAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/defundhub.defund.etf.Query/FundAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).FundAll(ctx, req.(*QueryAllFundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "defundhub.defund.etf.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fund",
			Handler:    _Query_Fund_Handler,
		},
		{
			MethodName: "FundAll",
			Handler:    _Query_FundAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "etf/query.proto",
}

func (m *QueryGetFundRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetFundRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetFundRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetFundResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetFundResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetFundResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Fund.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryAllFundRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllFundRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllFundRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QueryAllFundResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllFundResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllFundResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Fund) > 0 {
		for iNdEx := len(m.Fund) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fund[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryGetFundRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetFundResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Fund.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllFundRequest) Size() (n int) {
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

func (m *QueryAllFundResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Fund) > 0 {
		for _, e := range m.Fund {
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
func (m *QueryGetFundRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetFundRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetFundRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
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
			m.Index = string(dAtA[iNdEx:postIndex])
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
func (m *QueryGetFundResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetFundResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetFundResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fund", wireType)
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
			if err := m.Fund.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllFundRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllFundRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllFundRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryAllFundResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllFundResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllFundResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fund", wireType)
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
			m.Fund = append(m.Fund, Fund{})
			if err := m.Fund[len(m.Fund)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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