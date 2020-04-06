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

type ProduceRes struct {
	Partition            int32    `protobuf:"varint,1,opt,name=partition,proto3" json:"partition,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProduceRes) Reset()         { *m = ProduceRes{} }
func (m *ProduceRes) String() string { return proto.CompactTextString(m) }
func (*ProduceRes) ProtoMessage()    {}
func (*ProduceRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f475d2b3bc19ab6, []int{1}
}

func (m *ProduceRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProduceRes.Unmarshal(m, b)
}
func (m *ProduceRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProduceRes.Marshal(b, m, deterministic)
}
func (m *ProduceRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProduceRes.Merge(m, src)
}
func (m *ProduceRes) XXX_Size() int {
	return xxx_messageInfo_ProduceRes.Size(m)
}
func (m *ProduceRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ProduceRes.DiscardUnknown(m)
}

var xxx_messageInfo_ProduceRes proto.InternalMessageInfo

func (m *ProduceRes) GetPartition() int32 {
	if m != nil {
		return m.Partition
	}
	return 0
}

func (m *ProduceRes) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ConsumeRequest struct {
	Isolation            string   `protobuf:"bytes,1,opt,name=isolation,proto3" json:"isolation,omitempty"`
	Group                string   `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Topic                string   `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumeRequest) Reset()         { *m = ConsumeRequest{} }
func (m *ConsumeRequest) String() string { return proto.CompactTextString(m) }
func (*ConsumeRequest) ProtoMessage()    {}
func (*ConsumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f475d2b3bc19ab6, []int{2}
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

func (m *ConsumeRequest) GetGroup() string {
	if m != nil {
		return m.Group
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
	return fileDescriptor_1f475d2b3bc19ab6, []int{3}
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
	return fileDescriptor_1f475d2b3bc19ab6, []int{4}
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
	proto.RegisterType((*ProduceRes)(nil), "ProduceRes")
	proto.RegisterType((*ConsumeRequest)(nil), "ConsumeRequest")
	proto.RegisterType((*ConsumeRes)(nil), "ConsumeRes")
	proto.RegisterType((*Empty)(nil), "Empty")
}

func init() {
	proto.RegisterFile("kafka/grpc/kafka.proto", fileDescriptor_1f475d2b3bc19ab6)
}

var fileDescriptor_1f475d2b3bc19ab6 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0x4e, 0x02, 0x31,
	0x14, 0x85, 0xa9, 0x64, 0x20, 0x5c, 0x7e, 0x4c, 0x1a, 0x83, 0x13, 0xe2, 0x82, 0x74, 0x35, 0x6e,
	0x06, 0x94, 0x27, 0x10, 0xe3, 0xca, 0x0d, 0x19, 0x36, 0xc6, 0x85, 0x49, 0xad, 0x85, 0x34, 0x30,
	0xb4, 0xf6, 0xb6, 0x46, 0x5e, 0xc5, 0xa7, 0x35, 0x94, 0xf9, 0x91, 0xb8, 0x31, 0xee, 0x7a, 0x4e,
	0xd3, 0xaf, 0xe7, 0x9e, 0x5c, 0x18, 0x6e, 0xf8, 0x6a, 0xc3, 0x27, 0x6b, 0x6b, 0xc4, 0x24, 0x1c,
	0x53, 0x63, 0xb5, 0xd3, 0xec, 0x05, 0x06, 0x0b, 0xab, 0xdf, 0xbc, 0x90, 0x99, 0x7c, 0xf7, 0x12,
	0x1d, 0xbd, 0x82, 0x8e, 0x42, 0xbd, 0xe5, 0x4e, 0xe9, 0x5d, 0x4c, 0xc6, 0x24, 0xe9, 0x64, 0xb5,
	0x41, 0x2f, 0x20, 0x72, 0xda, 0x28, 0x11, 0x9f, 0x85, 0x9b, 0xa3, 0xa0, 0x31, 0xb4, 0x73, 0x89,
	0xc8, 0xd7, 0x32, 0x6e, 0x8e, 0x49, 0xd2, 0xcb, 0x4a, 0xc9, 0xe6, 0x00, 0x15, 0x1f, 0x0f, 0x6c,
	0xc3, 0xad, 0x53, 0x15, 0x3b, 0xca, 0x6a, 0x83, 0x0e, 0xa1, 0xa5, 0x57, 0x2b, 0x94, 0x2e, 0xc0,
	0x9b, 0x59, 0xa1, 0xd8, 0x33, 0x0c, 0xee, 0xf5, 0x0e, 0x7d, 0xfe, 0xf7, 0x8c, 0x6b, 0xab, 0xbd,
	0x29, 0x33, 0x06, 0x51, 0x27, 0x6f, 0xfe, 0x48, 0xce, 0x9e, 0x00, 0x2a, 0xf6, 0x3f, 0xf3, 0x1d,
	0xc8, 0x1f, 0x7c, 0xeb, 0xcb, 0xd9, 0x8f, 0x82, 0xb5, 0x21, 0x7a, 0xc8, 0x8d, 0xdb, 0xdf, 0x7e,
	0x11, 0x88, 0x1e, 0x0f, 0x95, 0xd3, 0x19, 0xf4, 0x8b, 0x32, 0x96, 0xce, 0x4a, 0x9e, 0xd3, 0xf3,
	0xf4, 0xb4, 0xfc, 0x51, 0xb7, 0x36, 0x90, 0x35, 0x12, 0x32, 0x25, 0xf4, 0x1a, 0x7a, 0x85, 0x77,
	0x87, 0xfb, 0x9d, 0xf8, 0xfd, 0xa6, 0x95, 0x86, 0x7f, 0x58, 0x83, 0xde, 0x40, 0xbf, 0x18, 0xa6,
	0xe2, 0x9f, 0x16, 0x37, 0xea, 0xd6, 0x06, 0xb2, 0xc6, 0x94, 0xcc, 0x13, 0xb8, 0x14, 0x3a, 0x4f,
	0xb9, 0x15, 0x1a, 0x3f, 0x53, 0x54, 0x56, 0x79, 0x4c, 0xc3, 0x82, 0xcc, 0xbb, 0xcb, 0xa0, 0x42,
	0xf4, 0x05, 0x79, 0x6d, 0x85, 0x85, 0x99, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x24, 0x64, 0x2d,
	0x2d, 0x4a, 0x02, 0x00, 0x00,
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
	Recv() (*ProduceRes, error)
	grpc.ClientStream
}

type kafkaProduceStreamClient struct {
	grpc.ClientStream
}

func (x *kafkaProduceStreamClient) Send(m *ProduceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kafkaProduceStreamClient) Recv() (*ProduceRes, error) {
	m := new(ProduceRes)
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
	Send(*ProduceRes) error
	Recv() (*ProduceRequest, error)
	grpc.ServerStream
}

type kafkaProduceStreamServer struct {
	grpc.ServerStream
}

func (x *kafkaProduceStreamServer) Send(m *ProduceRes) error {
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
			ServerStreams: true,
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
