// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kafka/grpc/kafka.proto

package kafka

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type ProduceRequest struct {
	Isolation            string   `protobuf:"bytes,1,opt,name=isolation,proto3" json:"isolation,omitempty"`
	Topic                string   `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Message              []byte   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProduceRequest) Reset()         { *m = ProduceRequest{} }
func (m *ProduceRequest) String() string { return proto.CompactTextString(m) }
func (*ProduceRequest) ProtoMessage()    {}
func (*ProduceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f475d2b3bc19ab6, []int{0}
}

func (m *ProduceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProduceRequest.Unmarshal(m, b)
}
func (m *ProduceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProduceRequest.Marshal(b, m, deterministic)
}
func (m *ProduceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProduceRequest.Merge(m, src)
}
func (m *ProduceRequest) XXX_Size() int {
	return xxx_messageInfo_ProduceRequest.Size(m)
}
func (m *ProduceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProduceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProduceRequest proto.InternalMessageInfo

func (m *ProduceRequest) GetIsolation() string {
	if m != nil {
		return m.Isolation
	}
	return ""
}

func (m *ProduceRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *ProduceRequest) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

type ConsumeRequest struct {
	Isolation            string   `protobuf:"bytes,1,opt,name=isolation,proto3" json:"isolation,omitempty"`
	Topic                string   `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumeRequest) Reset()         { *m = ConsumeRequest{} }
func (m *ConsumeRequest) String() string { return proto.CompactTextString(m) }
func (*ConsumeRequest) ProtoMessage()    {}
func (*ConsumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f475d2b3bc19ab6, []int{1}
}

func (m *ConsumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumeRequest.Unmarshal(m, b)
}
func (m *ConsumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumeRequest.Marshal(b, m, deterministic)
}
func (m *ConsumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumeRequest.Merge(m, src)
}
func (m *ConsumeRequest) XXX_Size() int {
	return xxx_messageInfo_ConsumeRequest.Size(m)
}
func (m *ConsumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumeRequest proto.InternalMessageInfo

func (m *ConsumeRequest) GetIsolation() string {
	if m != nil {
		return m.Isolation
	}
	return ""
}

func (m *ConsumeRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

type ConsumeRes struct {
	Partition            int32    `protobuf:"varint,1,opt,name=partition,proto3" json:"partition,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Value                []byte   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumeRes) Reset()         { *m = ConsumeRes{} }
func (m *ConsumeRes) String() string { return proto.CompactTextString(m) }
func (*ConsumeRes) ProtoMessage()    {}
func (*ConsumeRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f475d2b3bc19ab6, []int{2}
}

func (m *ConsumeRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumeRes.Unmarshal(m, b)
}
func (m *ConsumeRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumeRes.Marshal(b, m, deterministic)
}
func (m *ConsumeRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumeRes.Merge(m, src)
}
func (m *ConsumeRes) XXX_Size() int {
	return xxx_messageInfo_ConsumeRes.Size(m)
}
func (m *ConsumeRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumeRes.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumeRes proto.InternalMessageInfo

func (m *ConsumeRes) GetPartition() int32 {
	if m != nil {
		return m.Partition
	}
	return 0
}

func (m *ConsumeRes) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ConsumeRes) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f475d2b3bc19ab6, []int{3}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ProduceRequest)(nil), "ProduceRequest")
	proto.RegisterType((*ConsumeRequest)(nil), "ConsumeRequest")
	proto.RegisterType((*ConsumeRes)(nil), "ConsumeRes")
	proto.RegisterType((*Empty)(nil), "Empty")
}

func init() {
	proto.RegisterFile("kafka/grpc/kafka.proto", fileDescriptor_1f475d2b3bc19ab6)
}

var fileDescriptor_1f475d2b3bc19ab6 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x17, 0x47, 0x3b, 0xf6, 0xee, 0x8f, 0x10, 0x64, 0x96, 0xe1, 0x61, 0xe4, 0x54, 0x41,
	0x32, 0xff, 0x7c, 0x02, 0xa7, 0x9e, 0xbc, 0x8c, 0xee, 0xe2, 0x49, 0x88, 0x31, 0x1d, 0x61, 0x6b,
	0x53, 0xf3, 0xa6, 0xe2, 0x3e, 0x87, 0x5f, 0x58, 0x9a, 0x75, 0xad, 0xbb, 0x08, 0xde, 0xf2, 0x3c,
	0x6f, 0xf2, 0xe3, 0xc9, 0xf3, 0xc2, 0x64, 0x23, 0xd2, 0x8d, 0x98, 0xaf, 0x6d, 0x21, 0xe7, 0xfe,
	0xc8, 0x0b, 0x6b, 0x9c, 0x61, 0xaf, 0x30, 0x5e, 0x5a, 0xf3, 0x5e, 0x4a, 0x95, 0xa8, 0x8f, 0x52,
	0xa1, 0xa3, 0x17, 0xd0, 0xd7, 0x68, 0xb6, 0xc2, 0x69, 0x93, 0x47, 0x64, 0x46, 0xe2, 0x7e, 0xd2,
	0x1a, 0xf4, 0x0c, 0x02, 0x67, 0x0a, 0x2d, 0xa3, 0x13, 0x3f, 0xd9, 0x0b, 0x1a, 0x41, 0x2f, 0x53,
	0x88, 0x62, 0xad, 0xa2, 0xee, 0x8c, 0xc4, 0xc3, 0xe4, 0x20, 0xd9, 0x23, 0x8c, 0x1f, 0x4c, 0x8e,
	0x65, 0xf6, 0x5f, 0x7e, 0xf7, 0x17, 0x9f, 0xbd, 0x00, 0x34, 0x14, 0xac, 0x08, 0x85, 0xb0, 0x4e,
	0x37, 0x84, 0x20, 0x69, 0x0d, 0x3a, 0x81, 0xd0, 0xa4, 0x29, 0x2a, 0xe7, 0x23, 0x76, 0x93, 0x5a,
	0x55, 0xe4, 0x4f, 0xb1, 0x2d, 0x0f, 0x09, 0xf7, 0x82, 0xf5, 0x20, 0x78, 0xca, 0x0a, 0xb7, 0xbb,
	0xfd, 0x26, 0x10, 0x3c, 0x57, 0xc5, 0xd0, 0x2b, 0x18, 0xd5, 0x95, 0xac, 0x9c, 0x55, 0x22, 0xa3,
	0xa7, 0xfc, 0xb8, 0xa2, 0x69, 0xc8, 0xfd, 0x1b, 0xd6, 0x89, 0x09, 0xbd, 0x84, 0x61, 0x3d, 0xbd,
	0xc7, 0x5d, 0x2e, 0xff, 0xb8, 0x4c, 0x6f, 0x60, 0x54, 0xff, 0xa2, 0x01, 0x1f, 0x77, 0x33, 0x1d,
	0xb4, 0x06, 0xb2, 0xce, 0x35, 0x59, 0xc4, 0x70, 0x2e, 0x4d, 0xc6, 0x85, 0x95, 0x06, 0xbf, 0x38,
	0x6a, 0xab, 0x4b, 0xe4, 0x7e, 0x7f, 0x8b, 0xc1, 0xca, 0x2b, 0x9f, 0x79, 0x49, 0xde, 0x42, 0xbf,
	0xcf, 0xbb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x97, 0xd0, 0x2a, 0xe9, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KafkaClient is the client API for Kafka service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KafkaClient interface {
	ProduceStream(ctx context.Context, opts ...grpc.CallOption) (Kafka_ProduceStreamClient, error)
	ProduceAsync(ctx context.Context, in *ProduceRequest, opts ...grpc.CallOption) (*Empty, error)
	ConsumeStream(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (Kafka_ConsumeStreamClient, error)
}

type kafkaClient struct {
	cc grpc.ClientConnInterface
}

func NewKafkaClient(cc grpc.ClientConnInterface) KafkaClient {
	return &kafkaClient{cc}
}

func (c *kafkaClient) ProduceStream(ctx context.Context, opts ...grpc.CallOption) (Kafka_ProduceStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Kafka_serviceDesc.Streams[0], "/Kafka/ProduceStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &kafkaProduceStreamClient{stream}
	return x, nil
}

type Kafka_ProduceStreamClient interface {
	Send(*ProduceRequest) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type kafkaProduceStreamClient struct {
	grpc.ClientStream
}

func (x *kafkaProduceStreamClient) Send(m *ProduceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kafkaProduceStreamClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kafkaClient) ProduceAsync(ctx context.Context, in *ProduceRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Kafka/ProduceAsync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaClient) ConsumeStream(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (Kafka_ConsumeStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Kafka_serviceDesc.Streams[1], "/Kafka/ConsumeStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &kafkaConsumeStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Kafka_ConsumeStreamClient interface {
	Recv() (*ConsumeRes, error)
	grpc.ClientStream
}

type kafkaConsumeStreamClient struct {
	grpc.ClientStream
}

func (x *kafkaConsumeStreamClient) Recv() (*ConsumeRes, error) {
	m := new(ConsumeRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KafkaServer is the server API for Kafka service.
type KafkaServer interface {
	ProduceStream(Kafka_ProduceStreamServer) error
	ProduceAsync(context.Context, *ProduceRequest) (*Empty, error)
	ConsumeStream(*ConsumeRequest, Kafka_ConsumeStreamServer) error
}

// UnimplementedKafkaServer can be embedded to have forward compatible implementations.
type UnimplementedKafkaServer struct {
}

func (*UnimplementedKafkaServer) ProduceStream(srv Kafka_ProduceStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ProduceStream not implemented")
}
func (*UnimplementedKafkaServer) ProduceAsync(ctx context.Context, req *ProduceRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProduceAsync not implemented")
}
func (*UnimplementedKafkaServer) ConsumeStream(req *ConsumeRequest, srv Kafka_ConsumeStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ConsumeStream not implemented")
}

func RegisterKafkaServer(s *grpc.Server, srv KafkaServer) {
	s.RegisterService(&_Kafka_serviceDesc, srv)
}

func _Kafka_ProduceStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KafkaServer).ProduceStream(&kafkaProduceStreamServer{stream})
}

type Kafka_ProduceStreamServer interface {
	SendAndClose(*Empty) error
	Recv() (*ProduceRequest, error)
	grpc.ServerStream
}

type kafkaProduceStreamServer struct {
	grpc.ServerStream
}

func (x *kafkaProduceStreamServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *kafkaProduceStreamServer) Recv() (*ProduceRequest, error) {
	m := new(ProduceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Kafka_ProduceAsync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProduceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaServer).ProduceAsync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Kafka/ProduceAsync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaServer).ProduceAsync(ctx, req.(*ProduceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kafka_ConsumeStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConsumeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KafkaServer).ConsumeStream(m, &kafkaConsumeStreamServer{stream})
}

type Kafka_ConsumeStreamServer interface {
	Send(*ConsumeRes) error
	grpc.ServerStream
}

type kafkaConsumeStreamServer struct {
	grpc.ServerStream
}

func (x *kafkaConsumeStreamServer) Send(m *ConsumeRes) error {
	return x.ServerStream.SendMsg(m)
}

var _Kafka_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Kafka",
	HandlerType: (*KafkaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProduceAsync",
			Handler:    _Kafka_ProduceAsync_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ProduceStream",
			Handler:       _Kafka_ProduceStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ConsumeStream",
			Handler:       _Kafka_ConsumeStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "kafka/grpc/kafka.proto",
}
