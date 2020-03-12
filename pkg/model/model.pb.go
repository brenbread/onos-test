// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: model/model.proto

package model

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// ModelCheckState is the state if a model check
type ModelCheckerState int32

const (
	ModelCheckerState_RUNNING ModelCheckerState = 0
	ModelCheckerState_PASSED  ModelCheckerState = 1
	ModelCheckerState_FAILED  ModelCheckerState = 2
)

var ModelCheckerState_name = map[int32]string{
	0: "RUNNING",
	1: "PASSED",
	2: "FAILED",
}

var ModelCheckerState_value = map[string]int32{
	"RUNNING": 0,
	"PASSED":  1,
	"FAILED":  2,
}

func (x ModelCheckerState) String() string {
	return proto.EnumName(ModelCheckerState_name, int32(x))
}

func (ModelCheckerState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_312ac5bcab6cbb43, []int{0}
}

// Trace is a trace entry
type Trace struct {
	// bytes is the trace bytes
	Bytes []byte `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (m *Trace) Reset()         { *m = Trace{} }
func (m *Trace) String() string { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()    {}
func (*Trace) Descriptor() ([]byte, []int) {
	return fileDescriptor_312ac5bcab6cbb43, []int{0}
}
func (m *Trace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Trace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Trace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Trace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trace.Merge(m, src)
}
func (m *Trace) XXX_Size() int {
	return m.Size()
}
func (m *Trace) XXX_DiscardUnknown() {
	xxx_messageInfo_Trace.DiscardUnknown(m)
}

var xxx_messageInfo_Trace proto.InternalMessageInfo

func (m *Trace) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

// ModelCheckRequest is a model check request
type ModelCheckRequest struct {
	// model is the model to check
	Model string `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	// trace is the trace to check
	Trace string `protobuf:"bytes,2,opt,name=trace,proto3" json:"trace,omitempty"`
	// timeout is the model checker timeout
	Timeout *types.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (m *ModelCheckRequest) Reset()         { *m = ModelCheckRequest{} }
func (m *ModelCheckRequest) String() string { return proto.CompactTextString(m) }
func (*ModelCheckRequest) ProtoMessage()    {}
func (*ModelCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_312ac5bcab6cbb43, []int{1}
}
func (m *ModelCheckRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ModelCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ModelCheckRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ModelCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelCheckRequest.Merge(m, src)
}
func (m *ModelCheckRequest) XXX_Size() int {
	return m.Size()
}
func (m *ModelCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ModelCheckRequest proto.InternalMessageInfo

func (m *ModelCheckRequest) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *ModelCheckRequest) GetTrace() string {
	if m != nil {
		return m.Trace
	}
	return ""
}

func (m *ModelCheckRequest) GetTimeout() *types.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

// ModelCheckResponse is a model check response
type ModelCheckResponse struct {
	// state is the model checker state
	State ModelCheckerState `protobuf:"varint,1,opt,name=state,proto3,enum=onos.test.model.ModelCheckerState" json:"state,omitempty"`
	// message is the model checker message
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *ModelCheckResponse) Reset()         { *m = ModelCheckResponse{} }
func (m *ModelCheckResponse) String() string { return proto.CompactTextString(m) }
func (*ModelCheckResponse) ProtoMessage()    {}
func (*ModelCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_312ac5bcab6cbb43, []int{2}
}
func (m *ModelCheckResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ModelCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ModelCheckResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ModelCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelCheckResponse.Merge(m, src)
}
func (m *ModelCheckResponse) XXX_Size() int {
	return m.Size()
}
func (m *ModelCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ModelCheckResponse proto.InternalMessageInfo

func (m *ModelCheckResponse) GetState() ModelCheckerState {
	if m != nil {
		return m.State
	}
	return ModelCheckerState_RUNNING
}

func (m *ModelCheckResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("onos.test.model.ModelCheckerState", ModelCheckerState_name, ModelCheckerState_value)
	proto.RegisterType((*Trace)(nil), "onos.test.model.Trace")
	proto.RegisterType((*ModelCheckRequest)(nil), "onos.test.model.ModelCheckRequest")
	proto.RegisterType((*ModelCheckResponse)(nil), "onos.test.model.ModelCheckResponse")
}

func init() { proto.RegisterFile("model/model.proto", fileDescriptor_312ac5bcab6cbb43) }

var fileDescriptor_312ac5bcab6cbb43 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x41, 0x4f, 0x32, 0x31,
	0x10, 0xdd, 0xf2, 0x05, 0xc8, 0x37, 0x7c, 0xf9, 0x84, 0xca, 0x61, 0x25, 0xb1, 0x21, 0xeb, 0x85,
	0x78, 0xe8, 0x1a, 0xb8, 0x70, 0x45, 0x41, 0x43, 0xa2, 0xc4, 0x2c, 0x1a, 0xcf, 0x0b, 0x8c, 0x0b,
	0x11, 0xb6, 0xb8, 0xed, 0x9a, 0xf8, 0x2f, 0xfc, 0x59, 0x1e, 0x39, 0x7a, 0x34, 0xf0, 0x47, 0x4c,
	0x5b, 0x36, 0xa2, 0x07, 0x2e, 0x4d, 0xdf, 0x9b, 0xf7, 0x66, 0xde, 0x0c, 0x54, 0x16, 0x62, 0x82,
	0x73, 0xdf, 0xbc, 0x7c, 0x99, 0x08, 0x25, 0xe8, 0x81, 0x88, 0x85, 0xe4, 0x0a, 0xa5, 0xe2, 0x86,
	0xae, 0x55, 0x23, 0x11, 0x09, 0x53, 0xf3, 0xf5, 0xcf, 0xca, 0x6a, 0x2c, 0x12, 0x22, 0x9a, 0xa3,
	0x6f, 0xd0, 0x28, 0x7d, 0xf4, 0x27, 0x69, 0x12, 0xaa, 0x99, 0x88, 0x6d, 0xdd, 0x3b, 0x86, 0xfc,
	0x5d, 0x12, 0x8e, 0x91, 0x56, 0x21, 0x3f, 0x7a, 0x55, 0x28, 0x5d, 0x52, 0x27, 0x8d, 0x7f, 0x81,
	0x05, 0x9e, 0x82, 0xca, 0x8d, 0xee, 0x7e, 0x31, 0xc5, 0xf1, 0x53, 0x80, 0xcf, 0x29, 0x4a, 0xa5,
	0xa5, 0x66, 0xa4, 0x91, 0xfe, 0x0d, 0x2c, 0xd0, 0xac, 0xd2, 0x9d, 0xdc, 0x9c, 0x65, 0x0d, 0xa0,
	0x2d, 0x28, 0xaa, 0xd9, 0x02, 0x45, 0xaa, 0xdc, 0x3f, 0x75, 0xd2, 0x28, 0x35, 0x8f, 0xb8, 0x4d,
	0xc4, 0xb3, 0x44, 0xbc, 0xbb, 0x4d, 0x14, 0x64, 0x4a, 0x6f, 0x0a, 0x74, 0x77, 0xaa, 0x5c, 0x8a,
	0x58, 0x22, 0x6d, 0x43, 0x5e, 0xaa, 0x50, 0xa1, 0x19, 0xfb, 0xbf, 0xe9, 0xf1, 0x5f, 0x17, 0xe0,
	0xdf, 0x1e, 0x4c, 0x86, 0x5a, 0x19, 0x58, 0x03, 0x75, 0xa1, 0xb8, 0x40, 0x29, 0xc3, 0x28, 0x0b,
	0x97, 0xc1, 0xd3, 0xf6, 0xee, 0x7e, 0x5b, 0x17, 0x2d, 0x41, 0x31, 0xb8, 0x1f, 0x0c, 0xfa, 0x83,
	0xab, 0xb2, 0x43, 0x01, 0x0a, 0xb7, 0x9d, 0xe1, 0xb0, 0xd7, 0x2d, 0x13, 0xfd, 0xbf, 0xec, 0xf4,
	0xaf, 0x7b, 0xdd, 0x72, 0xae, 0x19, 0xc3, 0xe1, 0x0f, 0x27, 0x26, 0x2f, 0xb3, 0x31, 0xd2, 0x07,
	0x00, 0xc3, 0x98, 0x1a, 0xdd, 0x97, 0x71, 0x7b, 0xcd, 0xda, 0xc9, 0x5e, 0x8d, 0xdd, 0xfd, 0x8c,
	0x9c, 0xbb, 0xef, 0x6b, 0x46, 0x56, 0x6b, 0x46, 0x3e, 0xd7, 0x8c, 0xbc, 0x6d, 0x98, 0xb3, 0xda,
	0x30, 0xe7, 0x63, 0xc3, 0x9c, 0x51, 0xc1, 0x5c, 0xb2, 0xf5, 0x15, 0x00, 0x00, 0xff, 0xff, 0x21,
	0x95, 0xe4, 0x40, 0x25, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ModelCheckerServiceClient is the client API for ModelCheckerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ModelCheckerServiceClient interface {
	CheckModel(ctx context.Context, in *ModelCheckRequest, opts ...grpc.CallOption) (ModelCheckerService_CheckModelClient, error)
}

type modelCheckerServiceClient struct {
	cc *grpc.ClientConn
}

func NewModelCheckerServiceClient(cc *grpc.ClientConn) ModelCheckerServiceClient {
	return &modelCheckerServiceClient{cc}
}

func (c *modelCheckerServiceClient) CheckModel(ctx context.Context, in *ModelCheckRequest, opts ...grpc.CallOption) (ModelCheckerService_CheckModelClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ModelCheckerService_serviceDesc.Streams[0], "/onos.test.model.ModelCheckerService/CheckModel", opts...)
	if err != nil {
		return nil, err
	}
	x := &modelCheckerServiceCheckModelClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ModelCheckerService_CheckModelClient interface {
	Recv() (*ModelCheckResponse, error)
	grpc.ClientStream
}

type modelCheckerServiceCheckModelClient struct {
	grpc.ClientStream
}

func (x *modelCheckerServiceCheckModelClient) Recv() (*ModelCheckResponse, error) {
	m := new(ModelCheckResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ModelCheckerServiceServer is the server API for ModelCheckerService service.
type ModelCheckerServiceServer interface {
	CheckModel(*ModelCheckRequest, ModelCheckerService_CheckModelServer) error
}

// UnimplementedModelCheckerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedModelCheckerServiceServer struct {
}

func (*UnimplementedModelCheckerServiceServer) CheckModel(req *ModelCheckRequest, srv ModelCheckerService_CheckModelServer) error {
	return status.Errorf(codes.Unimplemented, "method CheckModel not implemented")
}

func RegisterModelCheckerServiceServer(s *grpc.Server, srv ModelCheckerServiceServer) {
	s.RegisterService(&_ModelCheckerService_serviceDesc, srv)
}

func _ModelCheckerService_CheckModel_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ModelCheckRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ModelCheckerServiceServer).CheckModel(m, &modelCheckerServiceCheckModelServer{stream})
}

type ModelCheckerService_CheckModelServer interface {
	Send(*ModelCheckResponse) error
	grpc.ServerStream
}

type modelCheckerServiceCheckModelServer struct {
	grpc.ServerStream
}

func (x *modelCheckerServiceCheckModelServer) Send(m *ModelCheckResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _ModelCheckerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "onos.test.model.ModelCheckerService",
	HandlerType: (*ModelCheckerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CheckModel",
			Handler:       _ModelCheckerService_CheckModel_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "model/model.proto",
}

func (m *Trace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Trace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Trace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bytes) > 0 {
		i -= len(m.Bytes)
		copy(dAtA[i:], m.Bytes)
		i = encodeVarintModel(dAtA, i, uint64(len(m.Bytes)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ModelCheckRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ModelCheckRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ModelCheckRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timeout != nil {
		{
			size, err := m.Timeout.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintModel(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Trace) > 0 {
		i -= len(m.Trace)
		copy(dAtA[i:], m.Trace)
		i = encodeVarintModel(dAtA, i, uint64(len(m.Trace)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Model) > 0 {
		i -= len(m.Model)
		copy(dAtA[i:], m.Model)
		i = encodeVarintModel(dAtA, i, uint64(len(m.Model)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ModelCheckResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ModelCheckResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ModelCheckResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintModel(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x12
	}
	if m.State != 0 {
		i = encodeVarintModel(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintModel(dAtA []byte, offset int, v uint64) int {
	offset -= sovModel(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Trace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Bytes)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	return n
}

func (m *ModelCheckRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Model)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	l = len(m.Trace)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	if m.Timeout != nil {
		l = m.Timeout.Size()
		n += 1 + l + sovModel(uint64(l))
	}
	return n
}

func (m *ModelCheckResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.State != 0 {
		n += 1 + sovModel(uint64(m.State))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	return n
}

func sovModel(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModel(x uint64) (n int) {
	return sovModel(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Trace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModel
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
			return fmt.Errorf("proto: Trace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Trace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthModel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bytes = append(m.Bytes[:0], dAtA[iNdEx:postIndex]...)
			if m.Bytes == nil {
				m.Bytes = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthModel
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthModel
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
func (m *ModelCheckRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModel
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
			return fmt.Errorf("proto: ModelCheckRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ModelCheckRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Model", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Model = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Trace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timeout == nil {
				m.Timeout = &types.Duration{}
			}
			if err := m.Timeout.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthModel
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthModel
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
func (m *ModelCheckResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModel
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
			return fmt.Errorf("proto: ModelCheckResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ModelCheckResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= ModelCheckerState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthModel
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthModel
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
func skipModel(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModel
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
					return 0, ErrIntOverflowModel
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
					return 0, ErrIntOverflowModel
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
				return 0, ErrInvalidLengthModel
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModel
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModel
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModel        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModel          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModel = fmt.Errorf("proto: unexpected end of group")
)