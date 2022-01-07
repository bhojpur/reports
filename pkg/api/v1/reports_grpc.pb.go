// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ReportsServiceClient is the client API for ReportsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportsServiceClient interface {
	// StartLocalDocument starts a Report Document on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the reports/config.yaml
	//   3. all bytes constituting the Document YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalDocument(ctx context.Context, opts ...grpc.CallOption) (ReportsService_StartLocalDocumentClient, error)
	// StartFromPreviousDocument starts a new Document based on a previous one.
	// If the previous Document does not have the can-replay condition set this call will result in an error.
	StartFromPreviousDocument(ctx context.Context, in *StartFromPreviousDocumentRequest, opts ...grpc.CallOption) (*StartDocumentResponse, error)
	// StartDocumentRequest starts a new Document based on its specification.
	StartDocument(ctx context.Context, in *StartDocumentRequest, opts ...grpc.CallOption) (*StartDocumentResponse, error)
	// Searches for Document known to this instance
	ListDocument(ctx context.Context, in *ListDocumentRequest, opts ...grpc.CallOption) (*ListDocumentResponse, error)
	// Subscribe listens to new Document updates
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (ReportsService_SubscribeClient, error)
	// GetDocument retrieves details of a single Document
	GetDocument(ctx context.Context, in *GetDocumentRequest, opts ...grpc.CallOption) (*GetDocumentResponse, error)
	// Listen listens to Document updates and log output of a running Document
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (ReportsService_ListenClient, error)
	// StopDocument stops a currently running Document
	StopDocument(ctx context.Context, in *StopDocumentRequest, opts ...grpc.CallOption) (*StopDocumentResponse, error)
}

type reportsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReportsServiceClient(cc grpc.ClientConnInterface) ReportsServiceClient {
	return &reportsServiceClient{cc}
}

func (c *reportsServiceClient) StartLocalDocument(ctx context.Context, opts ...grpc.CallOption) (ReportsService_StartLocalDocumentClient, error) {
	stream, err := c.cc.NewStream(ctx, &ReportsService_ServiceDesc.Streams[0], "/v1.ReportsService/StartLocalDocument", opts...)
	if err != nil {
		return nil, err
	}
	x := &reportsServiceStartLocalDocumentClient{stream}
	return x, nil
}

type ReportsService_StartLocalDocumentClient interface {
	Send(*StartLocalDocumentRequest) error
	CloseAndRecv() (*StartDocumentResponse, error)
	grpc.ClientStream
}

type reportsServiceStartLocalDocumentClient struct {
	grpc.ClientStream
}

func (x *reportsServiceStartLocalDocumentClient) Send(m *StartLocalDocumentRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *reportsServiceStartLocalDocumentClient) CloseAndRecv() (*StartDocumentResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StartDocumentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *reportsServiceClient) StartFromPreviousDocument(ctx context.Context, in *StartFromPreviousDocumentRequest, opts ...grpc.CallOption) (*StartDocumentResponse, error) {
	out := new(StartDocumentResponse)
	err := c.cc.Invoke(ctx, "/v1.ReportsService/StartFromPreviousDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportsServiceClient) StartDocument(ctx context.Context, in *StartDocumentRequest, opts ...grpc.CallOption) (*StartDocumentResponse, error) {
	out := new(StartDocumentResponse)
	err := c.cc.Invoke(ctx, "/v1.ReportsService/StartDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportsServiceClient) ListDocument(ctx context.Context, in *ListDocumentRequest, opts ...grpc.CallOption) (*ListDocumentResponse, error) {
	out := new(ListDocumentResponse)
	err := c.cc.Invoke(ctx, "/v1.ReportsService/ListDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportsServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (ReportsService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &ReportsService_ServiceDesc.Streams[1], "/v1.ReportsService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &reportsServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ReportsService_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type reportsServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *reportsServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *reportsServiceClient) GetDocument(ctx context.Context, in *GetDocumentRequest, opts ...grpc.CallOption) (*GetDocumentResponse, error) {
	out := new(GetDocumentResponse)
	err := c.cc.Invoke(ctx, "/v1.ReportsService/GetDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportsServiceClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (ReportsService_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &ReportsService_ServiceDesc.Streams[2], "/v1.ReportsService/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &reportsServiceListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ReportsService_ListenClient interface {
	Recv() (*ListenResponse, error)
	grpc.ClientStream
}

type reportsServiceListenClient struct {
	grpc.ClientStream
}

func (x *reportsServiceListenClient) Recv() (*ListenResponse, error) {
	m := new(ListenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *reportsServiceClient) StopDocument(ctx context.Context, in *StopDocumentRequest, opts ...grpc.CallOption) (*StopDocumentResponse, error) {
	out := new(StopDocumentResponse)
	err := c.cc.Invoke(ctx, "/v1.ReportsService/StopDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportsServiceServer is the server API for ReportsService service.
// All implementations must embed UnimplementedReportsServiceServer
// for forward compatibility
type ReportsServiceServer interface {
	// StartLocalDocument starts a Report Document on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the reports/config.yaml
	//   3. all bytes constituting the Document YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalDocument(ReportsService_StartLocalDocumentServer) error
	// StartFromPreviousDocument starts a new Document based on a previous one.
	// If the previous Document does not have the can-replay condition set this call will result in an error.
	StartFromPreviousDocument(context.Context, *StartFromPreviousDocumentRequest) (*StartDocumentResponse, error)
	// StartDocumentRequest starts a new Document based on its specification.
	StartDocument(context.Context, *StartDocumentRequest) (*StartDocumentResponse, error)
	// Searches for Document known to this instance
	ListDocument(context.Context, *ListDocumentRequest) (*ListDocumentResponse, error)
	// Subscribe listens to new Document updates
	Subscribe(*SubscribeRequest, ReportsService_SubscribeServer) error
	// GetDocument retrieves details of a single Document
	GetDocument(context.Context, *GetDocumentRequest) (*GetDocumentResponse, error)
	// Listen listens to Document updates and log output of a running Document
	Listen(*ListenRequest, ReportsService_ListenServer) error
	// StopDocument stops a currently running Document
	StopDocument(context.Context, *StopDocumentRequest) (*StopDocumentResponse, error)
	mustEmbedUnimplementedReportsServiceServer()
}

// UnimplementedReportsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReportsServiceServer struct {
}

func (UnimplementedReportsServiceServer) StartLocalDocument(ReportsService_StartLocalDocumentServer) error {
	return status.Errorf(codes.Unimplemented, "method StartLocalDocument not implemented")
}
func (UnimplementedReportsServiceServer) StartFromPreviousDocument(context.Context, *StartFromPreviousDocumentRequest) (*StartDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartFromPreviousDocument not implemented")
}
func (UnimplementedReportsServiceServer) StartDocument(context.Context, *StartDocumentRequest) (*StartDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartDocument not implemented")
}
func (UnimplementedReportsServiceServer) ListDocument(context.Context, *ListDocumentRequest) (*ListDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDocument not implemented")
}
func (UnimplementedReportsServiceServer) Subscribe(*SubscribeRequest, ReportsService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedReportsServiceServer) GetDocument(context.Context, *GetDocumentRequest) (*GetDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDocument not implemented")
}
func (UnimplementedReportsServiceServer) Listen(*ListenRequest, ReportsService_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedReportsServiceServer) StopDocument(context.Context, *StopDocumentRequest) (*StopDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopDocument not implemented")
}
func (UnimplementedReportsServiceServer) mustEmbedUnimplementedReportsServiceServer() {}

// UnsafeReportsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportsServiceServer will
// result in compilation errors.
type UnsafeReportsServiceServer interface {
	mustEmbedUnimplementedReportsServiceServer()
}

func RegisterReportsServiceServer(s grpc.ServiceRegistrar, srv ReportsServiceServer) {
	s.RegisterService(&ReportsService_ServiceDesc, srv)
}

func _ReportsService_StartLocalDocument_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ReportsServiceServer).StartLocalDocument(&reportsServiceStartLocalDocumentServer{stream})
}

type ReportsService_StartLocalDocumentServer interface {
	SendAndClose(*StartDocumentResponse) error
	Recv() (*StartLocalDocumentRequest, error)
	grpc.ServerStream
}

type reportsServiceStartLocalDocumentServer struct {
	grpc.ServerStream
}

func (x *reportsServiceStartLocalDocumentServer) SendAndClose(m *StartDocumentResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *reportsServiceStartLocalDocumentServer) Recv() (*StartLocalDocumentRequest, error) {
	m := new(StartLocalDocumentRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ReportsService_StartFromPreviousDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartFromPreviousDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportsServiceServer).StartFromPreviousDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ReportsService/StartFromPreviousDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportsServiceServer).StartFromPreviousDocument(ctx, req.(*StartFromPreviousDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportsService_StartDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportsServiceServer).StartDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ReportsService/StartDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportsServiceServer).StartDocument(ctx, req.(*StartDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportsService_ListDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportsServiceServer).ListDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ReportsService/ListDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportsServiceServer).ListDocument(ctx, req.(*ListDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportsService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReportsServiceServer).Subscribe(m, &reportsServiceSubscribeServer{stream})
}

type ReportsService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type reportsServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *reportsServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ReportsService_GetDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportsServiceServer).GetDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ReportsService/GetDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportsServiceServer).GetDocument(ctx, req.(*GetDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportsService_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReportsServiceServer).Listen(m, &reportsServiceListenServer{stream})
}

type ReportsService_ListenServer interface {
	Send(*ListenResponse) error
	grpc.ServerStream
}

type reportsServiceListenServer struct {
	grpc.ServerStream
}

func (x *reportsServiceListenServer) Send(m *ListenResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ReportsService_StopDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportsServiceServer).StopDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ReportsService/StopDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportsServiceServer).StopDocument(ctx, req.(*StopDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReportsService_ServiceDesc is the grpc.ServiceDesc for ReportsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReportsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ReportsService",
	HandlerType: (*ReportsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartFromPreviousDocument",
			Handler:    _ReportsService_StartFromPreviousDocument_Handler,
		},
		{
			MethodName: "StartDocument",
			Handler:    _ReportsService_StartDocument_Handler,
		},
		{
			MethodName: "ListDocument",
			Handler:    _ReportsService_ListDocument_Handler,
		},
		{
			MethodName: "GetDocument",
			Handler:    _ReportsService_GetDocument_Handler,
		},
		{
			MethodName: "StopDocument",
			Handler:    _ReportsService_StopDocument_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartLocalDocument",
			Handler:       _ReportsService_StartLocalDocument_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _ReportsService_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Listen",
			Handler:       _ReportsService_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "reports.proto",
}
