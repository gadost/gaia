// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: icaauth/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
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

// QueryInterchainAccountFromAddressRequest is the request type for the Query/InterchainAccountAddress RPC
type QueryInterchainAccountFromAddressRequest struct {
	Owner        string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	ConnectionId string `protobuf:"bytes,2,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty" yaml:"connection_id"`
}

func (m *QueryInterchainAccountFromAddressRequest) Reset() {
	*m = QueryInterchainAccountFromAddressRequest{}
}
func (m *QueryInterchainAccountFromAddressRequest) String() string { return proto.CompactTextString(m) }
func (*QueryInterchainAccountFromAddressRequest) ProtoMessage()    {}
func (*QueryInterchainAccountFromAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4ef18af0a32af09, []int{0}
}
func (m *QueryInterchainAccountFromAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryInterchainAccountFromAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryInterchainAccountFromAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryInterchainAccountFromAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryInterchainAccountFromAddressRequest.Merge(m, src)
}
func (m *QueryInterchainAccountFromAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryInterchainAccountFromAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryInterchainAccountFromAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryInterchainAccountFromAddressRequest proto.InternalMessageInfo

// QueryInterchainAccountFromAddressResponse the response type for the Query/InterchainAccountAddress RPC
type QueryInterchainAccountFromAddressResponse struct {
	InterchainAccountAddress string `protobuf:"bytes,1,opt,name=interchain_account_address,json=interchainAccountAddress,proto3" json:"interchain_account_address,omitempty" yaml:"interchain_account_address"`
}

func (m *QueryInterchainAccountFromAddressResponse) Reset() {
	*m = QueryInterchainAccountFromAddressResponse{}
}
func (m *QueryInterchainAccountFromAddressResponse) String() string {
	return proto.CompactTextString(m)
}
func (*QueryInterchainAccountFromAddressResponse) ProtoMessage() {}
func (*QueryInterchainAccountFromAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4ef18af0a32af09, []int{1}
}
func (m *QueryInterchainAccountFromAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryInterchainAccountFromAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryInterchainAccountFromAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryInterchainAccountFromAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryInterchainAccountFromAddressResponse.Merge(m, src)
}
func (m *QueryInterchainAccountFromAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryInterchainAccountFromAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryInterchainAccountFromAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryInterchainAccountFromAddressResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryInterchainAccountFromAddressRequest)(nil), "icaauth.v1.QueryInterchainAccountFromAddressRequest")
	proto.RegisterType((*QueryInterchainAccountFromAddressResponse)(nil), "icaauth.v1.QueryInterchainAccountFromAddressResponse")
}

func init() { proto.RegisterFile("icaauth/v1/query.proto", fileDescriptor_b4ef18af0a32af09) }

var fileDescriptor_b4ef18af0a32af09 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcb, 0x4c, 0x4e, 0x4c,
	0x2c, 0x2d, 0xc9, 0xd0, 0x2f, 0x33, 0xd4, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x82, 0x8a, 0xeb, 0x95, 0x19, 0x4a, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83,
	0x85, 0xf5, 0x41, 0x2c, 0x88, 0x0a, 0x29, 0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0xfd, 0xc4,
	0x82, 0x4c, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88, 0xac,
	0x52, 0x3d, 0x97, 0x46, 0x20, 0xc8, 0x38, 0xcf, 0xbc, 0x92, 0xd4, 0xa2, 0xe4, 0x8c, 0xc4, 0xcc,
	0x3c, 0xc7, 0xe4, 0xe4, 0xfc, 0xd2, 0xbc, 0x12, 0xb7, 0xa2, 0xfc, 0x5c, 0xc7, 0x94, 0x94, 0xa2,
	0xd4, 0xe2, 0xe2, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x11, 0x2e, 0xd6, 0xfc, 0xf2,
	0xbc, 0xd4, 0x22, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0xc8, 0x96, 0x8b, 0x37,
	0x39, 0x3f, 0x2f, 0x2f, 0x35, 0x19, 0x64, 0x6c, 0x7c, 0x66, 0x8a, 0x04, 0x13, 0x48, 0xd6, 0x49,
	0xe2, 0xd3, 0x3d, 0x79, 0x91, 0xca, 0xc4, 0xdc, 0x1c, 0x2b, 0x25, 0x14, 0x69, 0xa5, 0x20, 0x1e,
	0x04, 0xdf, 0x33, 0x45, 0x69, 0x02, 0x23, 0x97, 0x26, 0x11, 0x2e, 0x28, 0x2e, 0xc8, 0xcf, 0x2b,
	0x4e, 0x15, 0x4a, 0xe6, 0x92, 0xca, 0x84, 0xab, 0x8b, 0x4f, 0x84, 0x28, 0x8c, 0x4f, 0x84, 0xa8,
	0x82, 0xb8, 0xcb, 0x49, 0xf5, 0xd3, 0x3d, 0x79, 0x45, 0x88, 0xcd, 0xb8, 0xd5, 0x2a, 0x05, 0x49,
	0x64, 0xa2, 0x5b, 0x08, 0xb5, 0xcc, 0xe8, 0x19, 0x23, 0x17, 0x2b, 0xd8, 0x49, 0x42, 0x77, 0x18,
	0xb9, 0x64, 0xf0, 0xb9, 0x4b, 0xc8, 0x44, 0x0f, 0x11, 0xfe, 0x7a, 0xc4, 0x06, 0xa4, 0x94, 0x29,
	0x89, 0xba, 0x20, 0x9e, 0x57, 0xf2, 0x6f, 0xba, 0xfc, 0x64, 0x32, 0x93, 0xa7, 0x90, 0xbb, 0x3e,
	0x52, 0x62, 0xc0, 0xf4, 0xa2, 0x3e, 0x38, 0x5e, 0xf4, 0xab, 0xc1, 0x54, 0xad, 0x3e, 0x22, 0xb4,
	0xf5, 0xab, 0x51, 0x62, 0xa2, 0xd6, 0xc9, 0xfb, 0xc4, 0x43, 0x39, 0x86, 0x15, 0x8f, 0xe4, 0x18,
	0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5,
	0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x33, 0x3d, 0xb3, 0x24, 0xa3,
	0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x39, 0xbf, 0x38, 0x37, 0xbf, 0x58, 0x3f, 0x3d, 0x31,
	0x33, 0x51, 0xbf, 0xcc, 0x5c, 0xbf, 0x02, 0xee, 0x8a, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36,
	0x70, 0x82, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbc, 0xaf, 0xbc, 0xb0, 0xaa, 0x02, 0x00,
	0x00,
}

func (this *QueryInterchainAccountFromAddressRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryInterchainAccountFromAddressRequest)
	if !ok {
		that2, ok := that.(QueryInterchainAccountFromAddressRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Owner != that1.Owner {
		return false
	}
	if this.ConnectionId != that1.ConnectionId {
		return false
	}
	return true
}
func (this *QueryInterchainAccountFromAddressResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryInterchainAccountFromAddressResponse)
	if !ok {
		that2, ok := that.(QueryInterchainAccountFromAddressResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.InterchainAccountAddress != that1.InterchainAccountAddress {
		return false
	}
	return true
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
	// QueryInterchainAccountFromAddress returns the interchain account for given owner address on a given connection pair
	InterchainAccountFromAddress(ctx context.Context, in *QueryInterchainAccountFromAddressRequest, opts ...grpc.CallOption) (*QueryInterchainAccountFromAddressResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) InterchainAccountFromAddress(ctx context.Context, in *QueryInterchainAccountFromAddressRequest, opts ...grpc.CallOption) (*QueryInterchainAccountFromAddressResponse, error) {
	out := new(QueryInterchainAccountFromAddressResponse)
	err := c.cc.Invoke(ctx, "/icaauth.v1.Query/InterchainAccountFromAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// QueryInterchainAccountFromAddress returns the interchain account for given owner address on a given connection pair
	InterchainAccountFromAddress(context.Context, *QueryInterchainAccountFromAddressRequest) (*QueryInterchainAccountFromAddressResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) InterchainAccountFromAddress(ctx context.Context, req *QueryInterchainAccountFromAddressRequest) (*QueryInterchainAccountFromAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InterchainAccountFromAddress not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_InterchainAccountFromAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryInterchainAccountFromAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).InterchainAccountFromAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/icaauth.v1.Query/InterchainAccountFromAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).InterchainAccountFromAddress(ctx, req.(*QueryInterchainAccountFromAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "icaauth.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InterchainAccountFromAddress",
			Handler:    _Query_InterchainAccountFromAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "icaauth/v1/query.proto",
}

func (m *QueryInterchainAccountFromAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryInterchainAccountFromAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryInterchainAccountFromAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryInterchainAccountFromAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryInterchainAccountFromAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryInterchainAccountFromAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.InterchainAccountAddress) > 0 {
		i -= len(m.InterchainAccountAddress)
		copy(dAtA[i:], m.InterchainAccountAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.InterchainAccountAddress)))
		i--
		dAtA[i] = 0xa
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
func (m *QueryInterchainAccountFromAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryInterchainAccountFromAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.InterchainAccountAddress)
	if l > 0 {
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
func (m *QueryInterchainAccountFromAddressRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryInterchainAccountFromAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryInterchainAccountFromAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
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
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
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
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
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
func (m *QueryInterchainAccountFromAddressResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryInterchainAccountFromAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryInterchainAccountFromAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterchainAccountAddress", wireType)
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
			m.InterchainAccountAddress = string(dAtA[iNdEx:postIndex])
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
