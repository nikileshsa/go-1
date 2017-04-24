// Code generated by protoc-gen-go.
// source: examples/route_guide/routeguide/route_guide.proto
// DO NOT EDIT!

/*
Package routeguide is a generated protocol buffer package.

It is generated from these files:
	examples/route_guide/routeguide/route_guide.proto

It has these top-level messages:
	Point
	Rectangle
	Feature
	RouteNote
	RouteSummary
*/
package routeguide

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

// Points are represented as latitude-longitude pairs in the E7 representation
// (degrees multiplied by 10**7 and rounded to the nearest integer).
// Latitudes should be in the range +/- 90 degrees and longitude should be in
// the range +/- 180 degrees (inclusive).
type Point struct {
	Latitude  int32 `protobuf:"varint,1,opt,name=latitude" json:"latitude,omitempty"`
	Longitude int32 `protobuf:"varint,2,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Point) GetLatitude() int32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Point) GetLongitude() int32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

// A latitude-longitude rectangle, represented as two diagonally opposite
// points "lo" and "hi".
type Rectangle struct {
	// One corner of the rectangle.
	Lo *Point `protobuf:"bytes,1,opt,name=lo" json:"lo,omitempty"`
	// The other corner of the rectangle.
	Hi *Point `protobuf:"bytes,2,opt,name=hi" json:"hi,omitempty"`
}

func (m *Rectangle) Reset()                    { *m = Rectangle{} }
func (m *Rectangle) String() string            { return proto.CompactTextString(m) }
func (*Rectangle) ProtoMessage()               {}
func (*Rectangle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Rectangle) GetLo() *Point {
	if m != nil {
		return m.Lo
	}
	return nil
}

func (m *Rectangle) GetHi() *Point {
	if m != nil {
		return m.Hi
	}
	return nil
}

// A feature names something at a given point.
//
// If a feature could not be named, the name is empty.
type Feature struct {
	// The name of the feature.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The point where the feature is detected.
	Location *Point `protobuf:"bytes,2,opt,name=location" json:"location,omitempty"`
}

func (m *Feature) Reset()                    { *m = Feature{} }
func (m *Feature) String() string            { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()               {}
func (*Feature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Feature) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Feature) GetLocation() *Point {
	if m != nil {
		return m.Location
	}
	return nil
}

// A RouteNote is a message sent while at a given point.
type RouteNote struct {
	// The location from which the message is sent.
	Location *Point `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	// The message to be sent.
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *RouteNote) Reset()                    { *m = RouteNote{} }
func (m *RouteNote) String() string            { return proto.CompactTextString(m) }
func (*RouteNote) ProtoMessage()               {}
func (*RouteNote) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RouteNote) GetLocation() *Point {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *RouteNote) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// A RouteSummary is received in response to a RecordRoute rpc.
//
// It contains the number of individual points received, the number of
// detected features, and the total distance covered as the cumulative sum of
// the distance between each point.
type RouteSummary struct {
	// The number of points received.
	PointCount int32 `protobuf:"varint,1,opt,name=point_count,json=pointCount" json:"point_count,omitempty"`
	// The number of known features passed while traversing the route.
	FeatureCount int32 `protobuf:"varint,2,opt,name=feature_count,json=featureCount" json:"feature_count,omitempty"`
	// The distance covered in metres.
	Distance int32 `protobuf:"varint,3,opt,name=distance" json:"distance,omitempty"`
	// The duration of the traversal in seconds.
	ElapsedTime int32 `protobuf:"varint,4,opt,name=elapsed_time,json=elapsedTime" json:"elapsed_time,omitempty"`
}

func (m *RouteSummary) Reset()                    { *m = RouteSummary{} }
func (m *RouteSummary) String() string            { return proto.CompactTextString(m) }
func (*RouteSummary) ProtoMessage()               {}
func (*RouteSummary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RouteSummary) GetPointCount() int32 {
	if m != nil {
		return m.PointCount
	}
	return 0
}

func (m *RouteSummary) GetFeatureCount() int32 {
	if m != nil {
		return m.FeatureCount
	}
	return 0
}

func (m *RouteSummary) GetDistance() int32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *RouteSummary) GetElapsedTime() int32 {
	if m != nil {
		return m.ElapsedTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Point)(nil), "routeguide.Point")
	proto.RegisterType((*Rectangle)(nil), "routeguide.Rectangle")
	proto.RegisterType((*Feature)(nil), "routeguide.Feature")
	proto.RegisterType((*RouteNote)(nil), "routeguide.RouteNote")
	proto.RegisterType((*RouteSummary)(nil), "routeguide.RouteSummary")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RouteGuide service

type RouteGuideClient interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error)
	// A server-to-client streaming RPC.
	//
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	ListFeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteGuide_ListFeaturesClient, error)
	// A client-to-server streaming RPC.
	//
	// Accepts a stream of Points on a route being traversed, returning a
	// RouteSummary when traversal is completed.
	RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecordRouteClient, error)
	// A Bidirectional streaming RPC.
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RouteChatClient, error)
}

type routeGuideClient struct {
	cc *grpc.ClientConn
}

func NewRouteGuideClient(cc *grpc.ClientConn) RouteGuideClient {
	return &routeGuideClient{cc}
}

func (c *routeGuideClient) GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error) {
	out := new(Feature)
	err := grpc.Invoke(ctx, "/routeguide.RouteGuide/GetFeature", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeGuideClient) ListFeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteGuide_ListFeaturesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RouteGuide_serviceDesc.Streams[0], c.cc, "/routeguide.RouteGuide/ListFeatures", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideListFeaturesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RouteGuide_ListFeaturesClient interface {
	Recv() (*Feature, error)
	grpc.ClientStream
}

type routeGuideListFeaturesClient struct {
	grpc.ClientStream
}

func (x *routeGuideListFeaturesClient) Recv() (*Feature, error) {
	m := new(Feature)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecordRouteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RouteGuide_serviceDesc.Streams[1], c.cc, "/routeguide.RouteGuide/RecordRoute", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideRecordRouteClient{stream}
	return x, nil
}

type RouteGuide_RecordRouteClient interface {
	Send(*Point) error
	CloseAndRecv() (*RouteSummary, error)
	grpc.ClientStream
}

type routeGuideRecordRouteClient struct {
	grpc.ClientStream
}

func (x *routeGuideRecordRouteClient) Send(m *Point) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideRecordRouteClient) CloseAndRecv() (*RouteSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RouteSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RouteChatClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RouteGuide_serviceDesc.Streams[2], c.cc, "/routeguide.RouteGuide/RouteChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideRouteChatClient{stream}
	return x, nil
}

type RouteGuide_RouteChatClient interface {
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
	grpc.ClientStream
}

type routeGuideRouteChatClient struct {
	grpc.ClientStream
}

func (x *routeGuideRouteChatClient) Send(m *RouteNote) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideRouteChatClient) Recv() (*RouteNote, error) {
	m := new(RouteNote)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for RouteGuide service

type RouteGuideServer interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetFeature(context.Context, *Point) (*Feature, error)
	// A server-to-client streaming RPC.
	//
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	ListFeatures(*Rectangle, RouteGuide_ListFeaturesServer) error
	// A client-to-server streaming RPC.
	//
	// Accepts a stream of Points on a route being traversed, returning a
	// RouteSummary when traversal is completed.
	RecordRoute(RouteGuide_RecordRouteServer) error
	// A Bidirectional streaming RPC.
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	RouteChat(RouteGuide_RouteChatServer) error
}

func RegisterRouteGuideServer(s *grpc.Server, srv RouteGuideServer) {
	s.RegisterService(&_RouteGuide_serviceDesc, srv)
}

func _RouteGuide_GetFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).GetFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeguide.RouteGuide/GetFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).GetFeature(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteGuide_ListFeatures_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Rectangle)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouteGuideServer).ListFeatures(m, &routeGuideListFeaturesServer{stream})
}

type RouteGuide_ListFeaturesServer interface {
	Send(*Feature) error
	grpc.ServerStream
}

type routeGuideListFeaturesServer struct {
	grpc.ServerStream
}

func (x *routeGuideListFeaturesServer) Send(m *Feature) error {
	return x.ServerStream.SendMsg(m)
}

func _RouteGuide_RecordRoute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).RecordRoute(&routeGuideRecordRouteServer{stream})
}

type RouteGuide_RecordRouteServer interface {
	SendAndClose(*RouteSummary) error
	Recv() (*Point, error)
	grpc.ServerStream
}

type routeGuideRecordRouteServer struct {
	grpc.ServerStream
}

func (x *routeGuideRecordRouteServer) SendAndClose(m *RouteSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideRecordRouteServer) Recv() (*Point, error) {
	m := new(Point)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteGuide_RouteChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).RouteChat(&routeGuideRouteChatServer{stream})
}

type RouteGuide_RouteChatServer interface {
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
	grpc.ServerStream
}

type routeGuideRouteChatServer struct {
	grpc.ServerStream
}

func (x *routeGuideRouteChatServer) Send(m *RouteNote) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideRouteChatServer) Recv() (*RouteNote, error) {
	m := new(RouteNote)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RouteGuide_serviceDesc = grpc.ServiceDesc{
	ServiceName: "routeguide.RouteGuide",
	HandlerType: (*RouteGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeature",
			Handler:    _RouteGuide_GetFeature_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListFeatures",
			Handler:       _RouteGuide_ListFeatures_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RecordRoute",
			Handler:       _RouteGuide_RecordRoute_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "RouteChat",
			Handler:       _RouteGuide_RouteChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "examples/route_guide/routeguide/route_guide.proto",
}

func init() { proto.RegisterFile("examples/route_guide/routeguide/route_guide.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x5f, 0x8b, 0xd3, 0x40,
	0x10, 0xbf, 0x8d, 0x77, 0x5e, 0x33, 0x89, 0x88, 0x23, 0x42, 0x88, 0x82, 0x5e, 0x7c, 0xb9, 0x17,
	0x63, 0x3d, 0xc1, 0xc7, 0x13, 0xaf, 0x60, 0x5f, 0x8a, 0xd4, 0xd8, 0xf7, 0xb2, 0x26, 0x63, 0xba,
	0xb0, 0xc9, 0x86, 0x64, 0x03, 0xfa, 0x01, 0xfc, 0x04, 0x7e, 0x61, 0xc9, 0x6e, 0xd2, 0xa4, 0xda,
	0x72, 0x6f, 0x33, 0xbf, 0xf9, 0xfd, 0xe6, 0x2f, 0x03, 0xef, 0xe8, 0x27, 0x2f, 0x2a, 0x49, 0xcd,
	0xdb, 0x5a, 0xb5, 0x9a, 0xb6, 0x79, 0x2b, 0x32, 0xb2, 0xf6, 0xc4, 0xb4, 0x70, 0x5c, 0xd5, 0x4a,
	0x2b, 0x84, 0x31, 0x1a, 0x7d, 0x82, 0x8b, 0xb5, 0x12, 0xa5, 0xc6, 0x10, 0x66, 0x92, 0x6b, 0xa1,
	0xdb, 0x8c, 0x02, 0xf6, 0x8a, 0x5d, 0x5f, 0x24, 0x7b, 0x1f, 0x5f, 0x80, 0x2b, 0x55, 0x99, 0xdb,
	0xa0, 0x63, 0x82, 0x23, 0x10, 0x7d, 0x05, 0x37, 0xa1, 0x54, 0xf3, 0x32, 0x97, 0x84, 0x57, 0xe0,
	0x48, 0x65, 0x12, 0x78, 0x37, 0x4f, 0xe2, 0xb1, 0x50, 0x6c, 0xaa, 0x24, 0x8e, 0x54, 0x1d, 0x65,
	0x27, 0x4c, 0x9a, 0xe3, 0x94, 0x9d, 0x88, 0x56, 0x70, 0xf9, 0x99, 0xb8, 0x6e, 0x6b, 0x42, 0x84,
	0xf3, 0x92, 0x17, 0xb6, 0x27, 0x37, 0x31, 0x36, 0xbe, 0x81, 0x99, 0x54, 0x29, 0xd7, 0x42, 0x95,
	0xa7, 0xf3, 0xec, 0x29, 0xd1, 0x06, 0xdc, 0xa4, 0x8b, 0x7e, 0x51, 0xfa, 0x50, 0xcb, 0xee, 0xd5,
	0x62, 0x00, 0x97, 0x05, 0x35, 0x0d, 0xcf, 0xed, 0xe0, 0x6e, 0x32, 0xb8, 0xd1, 0x1f, 0x06, 0xbe,
	0x49, 0xfb, 0xad, 0x2d, 0x0a, 0x5e, 0xff, 0xc2, 0x97, 0xe0, 0x55, 0x9d, 0x7a, 0x9b, 0xaa, 0xb6,
	0xd4, 0xfd, 0x12, 0xc1, 0x40, 0x8b, 0x0e, 0xc1, 0xd7, 0xf0, 0xe8, 0x87, 0x9d, 0xaa, 0xa7, 0xd8,
	0x55, 0xfa, 0x3d, 0x68, 0x49, 0x21, 0xcc, 0x32, 0xd1, 0x68, 0x5e, 0xa6, 0x14, 0x3c, 0xb0, 0x77,
	0x18, 0x7c, 0xbc, 0x02, 0x9f, 0x24, 0xaf, 0x1a, 0xca, 0xb6, 0x5a, 0x14, 0x14, 0x9c, 0x9b, 0xb8,
	0xd7, 0x63, 0x1b, 0x51, 0xd0, 0xcd, 0x6f, 0x07, 0xc0, 0x74, 0xb5, 0xec, 0xc6, 0xc1, 0x0f, 0x00,
	0x4b, 0xd2, 0xc3, 0x2e, 0xff, 0x9f, 0x34, 0x7c, 0x3a, 0x85, 0x7a, 0x5e, 0x74, 0x86, 0xb7, 0xe0,
	0xaf, 0x44, 0x33, 0x08, 0x1b, 0x7c, 0x36, 0xa5, 0xed, 0xaf, 0x7d, 0x42, 0x3d, 0x67, 0x78, 0x0b,
	0x5e, 0x42, 0xa9, 0xaa, 0x33, 0xd3, 0xcb, 0xb1, 0xc2, 0xc1, 0x41, 0xc6, 0xc9, 0x1e, 0xa3, 0xb3,
	0x6b, 0x86, 0x1f, 0xfb, 0x93, 0x2d, 0x76, 0x5c, 0xff, 0x53, 0x7c, 0xb8, 0x64, 0x78, 0x1c, 0xee,
	0xe4, 0x73, 0x76, 0x37, 0x87, 0xe7, 0x42, 0xc5, 0x79, 0x5d, 0xa5, 0xf1, 0xf0, 0x20, 0x13, 0xfa,
	0xdd, 0xe3, 0x71, 0x47, 0xeb, 0xee, 0x27, 0xd6, 0xec, 0xfb, 0x43, 0xf3, 0x1c, 0xef, 0xff, 0x06,
	0x00, 0x00, 0xff, 0xff, 0xca, 0xca, 0x23, 0x46, 0x51, 0x03, 0x00, 0x00,
}