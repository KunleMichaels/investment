// Code generated by protoc-gen-go.
// source: protos/broker.proto
// DO NOT EDIT!

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	protos/broker.proto

It has these top-level messages:
	QuoteID
	Quote
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type QuoteID struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *QuoteID) Reset()                    { *m = QuoteID{} }
func (m *QuoteID) String() string            { return proto.CompactTextString(m) }
func (*QuoteID) ProtoMessage()               {}
func (*QuoteID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Quote struct {
	Name string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Date int64   `protobuf:"varint,2,opt,name=date" json:"date,omitempty"`
	Bid  float64 `protobuf:"fixed64,3,opt,name=bid" json:"bid,omitempty"`
	Ask  float64 `protobuf:"fixed64,4,opt,name=ask" json:"ask,omitempty"`
}

func (m *Quote) Reset()                    { *m = Quote{} }
func (m *Quote) String() string            { return proto.CompactTextString(m) }
func (*Quote) ProtoMessage()               {}
func (*Quote) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*QuoteID)(nil), "protos.QuoteID")
	proto.RegisterType((*Quote)(nil), "protos.Quote")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Broker service

type BrokerClient interface {
	GetQuote(ctx context.Context, in *QuoteID, opts ...grpc.CallOption) (*Quote, error)
	StreamQuotes(ctx context.Context, in *QuoteID, opts ...grpc.CallOption) (Broker_StreamQuotesClient, error)
}

type brokerClient struct {
	cc *grpc.ClientConn
}

func NewBrokerClient(cc *grpc.ClientConn) BrokerClient {
	return &brokerClient{cc}
}

func (c *brokerClient) GetQuote(ctx context.Context, in *QuoteID, opts ...grpc.CallOption) (*Quote, error) {
	out := new(Quote)
	err := grpc.Invoke(ctx, "/protos.Broker/GetQuote", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerClient) StreamQuotes(ctx context.Context, in *QuoteID, opts ...grpc.CallOption) (Broker_StreamQuotesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Broker_serviceDesc.Streams[0], c.cc, "/protos.Broker/StreamQuotes", opts...)
	if err != nil {
		return nil, err
	}
	x := &brokerStreamQuotesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Broker_StreamQuotesClient interface {
	Recv() (*Quote, error)
	grpc.ClientStream
}

type brokerStreamQuotesClient struct {
	grpc.ClientStream
}

func (x *brokerStreamQuotesClient) Recv() (*Quote, error) {
	m := new(Quote)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Broker service

type BrokerServer interface {
	GetQuote(context.Context, *QuoteID) (*Quote, error)
	StreamQuotes(*QuoteID, Broker_StreamQuotesServer) error
}

func RegisterBrokerServer(s *grpc.Server, srv BrokerServer) {
	s.RegisterService(&_Broker_serviceDesc, srv)
}

func _Broker_GetQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuoteID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).GetQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Broker/GetQuote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).GetQuote(ctx, req.(*QuoteID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broker_StreamQuotes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QuoteID)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BrokerServer).StreamQuotes(m, &brokerStreamQuotesServer{stream})
}

type Broker_StreamQuotesServer interface {
	Send(*Quote) error
	grpc.ServerStream
}

type brokerStreamQuotesServer struct {
	grpc.ServerStream
}

func (x *brokerStreamQuotesServer) Send(m *Quote) error {
	return x.ServerStream.SendMsg(m)
}

var _Broker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Broker",
	HandlerType: (*BrokerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuote",
			Handler:    _Broker_GetQuote_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamQuotes",
			Handler:       _Broker_StreamQuotes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("protos/broker.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x4f, 0x2a, 0xca, 0xcf, 0x4e, 0x2d, 0xd2, 0x03, 0xf3, 0x84, 0xd8, 0x20, 0x82,
	0x4a, 0xb2, 0x5c, 0xec, 0x81, 0xa5, 0xf9, 0x25, 0xa9, 0x9e, 0x2e, 0x42, 0x42, 0x5c, 0x2c, 0x79,
	0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x52, 0x30, 0x17, 0x2b,
	0x58, 0x1a, 0x9b, 0x24, 0x48, 0x2c, 0x25, 0xb1, 0x24, 0x55, 0x82, 0x09, 0x28, 0xc6, 0x1c, 0x04,
	0x66, 0x0b, 0x09, 0x70, 0x31, 0x27, 0x65, 0xa6, 0x48, 0x30, 0x03, 0x85, 0x18, 0x83, 0x40, 0x4c,
	0x90, 0x48, 0x62, 0x71, 0xb6, 0x04, 0x0b, 0x44, 0x04, 0xc8, 0x34, 0xca, 0xe2, 0x62, 0x73, 0x02,
	0xbb, 0x45, 0x48, 0x87, 0x8b, 0xc3, 0x3d, 0xb5, 0x04, 0x62, 0x03, 0x3f, 0xc4, 0x65, 0xc5, 0x7a,
	0x50, 0xf7, 0x48, 0xf1, 0xa2, 0x08, 0x28, 0x31, 0x08, 0x19, 0x71, 0xf1, 0x04, 0x97, 0x14, 0xa5,
	0x26, 0xe6, 0x82, 0x05, 0x8a, 0x09, 0xeb, 0x30, 0x60, 0x4c, 0x82, 0xf8, 0xd3, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0xd2, 0x05, 0xfb, 0x9c, 0x05, 0x01, 0x00, 0x00,
}