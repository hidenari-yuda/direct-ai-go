// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: media.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MediaServiceClient is the client API for MediaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MediaServiceClient interface {
	// create
	Create(ctx context.Context, in *Media, opts ...grpc.CallOption) (*Media, error)
	// update
	Update(ctx context.Context, in *Media, opts ...grpc.CallOption) (*MediaBoolResponse, error)
	// delete
	Delete(ctx context.Context, in *MediaIdRequest, opts ...grpc.CallOption) (*MediaBoolResponse, error)
	// get by id
	GetById(ctx context.Context, in *MediaIdRequest, opts ...grpc.CallOption) (*Media, error)
	// get by media_type
	GetListByType(ctx context.Context, in *MediaTypeRequest, opts ...grpc.CallOption) (*MediaList, error)
	GetListByUserId(ctx context.Context, in *MediaIdRequest, opts ...grpc.CallOption) (*MediaList, error)
	// get list by id list
	GetMediaListByIdList(ctx context.Context, opts ...grpc.CallOption) (MediaService_GetMediaListByIdListClient, error)
	// get all
	GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*MediaList, error)
}

type mediaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMediaServiceClient(cc grpc.ClientConnInterface) MediaServiceClient {
	return &mediaServiceClient{cc}
}

func (c *mediaServiceClient) Create(ctx context.Context, in *Media, opts ...grpc.CallOption) (*Media, error) {
	out := new(Media)
	err := c.cc.Invoke(ctx, "/media.MediaService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) Update(ctx context.Context, in *Media, opts ...grpc.CallOption) (*MediaBoolResponse, error) {
	out := new(MediaBoolResponse)
	err := c.cc.Invoke(ctx, "/media.MediaService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) Delete(ctx context.Context, in *MediaIdRequest, opts ...grpc.CallOption) (*MediaBoolResponse, error) {
	out := new(MediaBoolResponse)
	err := c.cc.Invoke(ctx, "/media.MediaService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) GetById(ctx context.Context, in *MediaIdRequest, opts ...grpc.CallOption) (*Media, error) {
	out := new(Media)
	err := c.cc.Invoke(ctx, "/media.MediaService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) GetListByType(ctx context.Context, in *MediaTypeRequest, opts ...grpc.CallOption) (*MediaList, error) {
	out := new(MediaList)
	err := c.cc.Invoke(ctx, "/media.MediaService/GetListByType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) GetListByUserId(ctx context.Context, in *MediaIdRequest, opts ...grpc.CallOption) (*MediaList, error) {
	out := new(MediaList)
	err := c.cc.Invoke(ctx, "/media.MediaService/GetListByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) GetMediaListByIdList(ctx context.Context, opts ...grpc.CallOption) (MediaService_GetMediaListByIdListClient, error) {
	stream, err := c.cc.NewStream(ctx, &MediaService_ServiceDesc.Streams[0], "/media.MediaService/GetMediaListByIdList", opts...)
	if err != nil {
		return nil, err
	}
	x := &mediaServiceGetMediaListByIdListClient{stream}
	return x, nil
}

type MediaService_GetMediaListByIdListClient interface {
	Send(*MediaIdListRequest) error
	CloseAndRecv() (*MediaList, error)
	grpc.ClientStream
}

type mediaServiceGetMediaListByIdListClient struct {
	grpc.ClientStream
}

func (x *mediaServiceGetMediaListByIdListClient) Send(m *MediaIdListRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mediaServiceGetMediaListByIdListClient) CloseAndRecv() (*MediaList, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MediaList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mediaServiceClient) GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*MediaList, error) {
	out := new(MediaList)
	err := c.cc.Invoke(ctx, "/media.MediaService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MediaServiceServer is the server API for MediaService service.
// All implementations should embed UnimplementedMediaServiceServer
// for forward compatibility
type MediaServiceServer interface {
	// create
	Create(context.Context, *Media) (*Media, error)
	// update
	Update(context.Context, *Media) (*MediaBoolResponse, error)
	// delete
	Delete(context.Context, *MediaIdRequest) (*MediaBoolResponse, error)
	// get by id
	GetById(context.Context, *MediaIdRequest) (*Media, error)
	// get by media_type
	GetListByType(context.Context, *MediaTypeRequest) (*MediaList, error)
	GetListByUserId(context.Context, *MediaIdRequest) (*MediaList, error)
	// get list by id list
	GetMediaListByIdList(MediaService_GetMediaListByIdListServer) error
	// get all
	GetAll(context.Context, *emptypb.Empty) (*MediaList, error)
}

// UnimplementedMediaServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMediaServiceServer struct {
}

func (UnimplementedMediaServiceServer) Create(context.Context, *Media) (*Media, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMediaServiceServer) Update(context.Context, *Media) (*MediaBoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMediaServiceServer) Delete(context.Context, *MediaIdRequest) (*MediaBoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedMediaServiceServer) GetById(context.Context, *MediaIdRequest) (*Media, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedMediaServiceServer) GetListByType(context.Context, *MediaTypeRequest) (*MediaList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListByType not implemented")
}
func (UnimplementedMediaServiceServer) GetListByUserId(context.Context, *MediaIdRequest) (*MediaList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListByUserId not implemented")
}
func (UnimplementedMediaServiceServer) GetMediaListByIdList(MediaService_GetMediaListByIdListServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMediaListByIdList not implemented")
}
func (UnimplementedMediaServiceServer) GetAll(context.Context, *emptypb.Empty) (*MediaList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}

// UnsafeMediaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MediaServiceServer will
// result in compilation errors.
type UnsafeMediaServiceServer interface {
	mustEmbedUnimplementedMediaServiceServer()
}

func RegisterMediaServiceServer(s grpc.ServiceRegistrar, srv MediaServiceServer) {
	s.RegisterService(&MediaService_ServiceDesc, srv)
}

func _MediaService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Media)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).Create(ctx, req.(*Media))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Media)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).Update(ctx, req.(*Media))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MediaIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).Delete(ctx, req.(*MediaIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MediaIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).GetById(ctx, req.(*MediaIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_GetListByType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MediaTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).GetListByType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/GetListByType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).GetListByType(ctx, req.(*MediaTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_GetListByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MediaIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).GetListByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/GetListByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).GetListByUserId(ctx, req.(*MediaIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_GetMediaListByIdList_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MediaServiceServer).GetMediaListByIdList(&mediaServiceGetMediaListByIdListServer{stream})
}

type MediaService_GetMediaListByIdListServer interface {
	SendAndClose(*MediaList) error
	Recv() (*MediaIdListRequest, error)
	grpc.ServerStream
}

type mediaServiceGetMediaListByIdListServer struct {
	grpc.ServerStream
}

func (x *mediaServiceGetMediaListByIdListServer) SendAndClose(m *MediaList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mediaServiceGetMediaListByIdListServer) Recv() (*MediaIdListRequest, error) {
	m := new(MediaIdListRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MediaService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).GetAll(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// MediaService_ServiceDesc is the grpc.ServiceDesc for MediaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MediaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "media.MediaService",
	HandlerType: (*MediaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _MediaService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _MediaService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _MediaService_Delete_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _MediaService_GetById_Handler,
		},
		{
			MethodName: "GetListByType",
			Handler:    _MediaService_GetListByType_Handler,
		},
		{
			MethodName: "GetListByUserId",
			Handler:    _MediaService_GetListByUserId_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _MediaService_GetAll_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMediaListByIdList",
			Handler:       _MediaService_GetMediaListByIdList_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "media.proto",
}