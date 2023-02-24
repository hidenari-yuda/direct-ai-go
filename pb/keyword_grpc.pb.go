// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: keyword.proto

package pb

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

// KeywordServiceClient is the client API for KeywordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeywordServiceClient interface {
	// Create
	Create(ctx context.Context, in *Keyword, opts ...grpc.CallOption) (*Keyword, error)
	// Update
	Update(ctx context.Context, in *Keyword, opts ...grpc.CallOption) (*Keyword, error)
	// Get
	GetById(ctx context.Context, in *KeywordIdRequest, opts ...grpc.CallOption) (*Keyword, error)
	GetListByChannel(ctx context.Context, in *KeywordIdRequest, opts ...grpc.CallOption) (*KeywordList, error)
}

type keywordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeywordServiceClient(cc grpc.ClientConnInterface) KeywordServiceClient {
	return &keywordServiceClient{cc}
}

func (c *keywordServiceClient) Create(ctx context.Context, in *Keyword, opts ...grpc.CallOption) (*Keyword, error) {
	out := new(Keyword)
	err := c.cc.Invoke(ctx, "/keyword.KeywordService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keywordServiceClient) Update(ctx context.Context, in *Keyword, opts ...grpc.CallOption) (*Keyword, error) {
	out := new(Keyword)
	err := c.cc.Invoke(ctx, "/keyword.KeywordService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keywordServiceClient) GetById(ctx context.Context, in *KeywordIdRequest, opts ...grpc.CallOption) (*Keyword, error) {
	out := new(Keyword)
	err := c.cc.Invoke(ctx, "/keyword.KeywordService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keywordServiceClient) GetListByChannel(ctx context.Context, in *KeywordIdRequest, opts ...grpc.CallOption) (*KeywordList, error) {
	out := new(KeywordList)
	err := c.cc.Invoke(ctx, "/keyword.KeywordService/GetListByChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordServiceServer is the server API for KeywordService service.
// All implementations should embed UnimplementedKeywordServiceServer
// for forward compatibility
type KeywordServiceServer interface {
	// Create
	Create(context.Context, *Keyword) (*Keyword, error)
	// Update
	Update(context.Context, *Keyword) (*Keyword, error)
	// Get
	GetById(context.Context, *KeywordIdRequest) (*Keyword, error)
	GetListByChannel(context.Context, *KeywordIdRequest) (*KeywordList, error)
}

// UnimplementedKeywordServiceServer should be embedded to have forward compatible implementations.
type UnimplementedKeywordServiceServer struct {
}

func (UnimplementedKeywordServiceServer) Create(context.Context, *Keyword) (*Keyword, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedKeywordServiceServer) Update(context.Context, *Keyword) (*Keyword, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedKeywordServiceServer) GetById(context.Context, *KeywordIdRequest) (*Keyword, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedKeywordServiceServer) GetListByChannel(context.Context, *KeywordIdRequest) (*KeywordList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListByChannel not implemented")
}

// UnsafeKeywordServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeywordServiceServer will
// result in compilation errors.
type UnsafeKeywordServiceServer interface {
	mustEmbedUnimplementedKeywordServiceServer()
}

func RegisterKeywordServiceServer(s grpc.ServiceRegistrar, srv KeywordServiceServer) {
	s.RegisterService(&KeywordService_ServiceDesc, srv)
}

func _KeywordService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Keyword)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyword.KeywordService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordServiceServer).Create(ctx, req.(*Keyword))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeywordService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Keyword)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyword.KeywordService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordServiceServer).Update(ctx, req.(*Keyword))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeywordService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeywordIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyword.KeywordService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordServiceServer).GetById(ctx, req.(*KeywordIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeywordService_GetListByChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeywordIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordServiceServer).GetListByChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyword.KeywordService/GetListByChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordServiceServer).GetListByChannel(ctx, req.(*KeywordIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeywordService_ServiceDesc is the grpc.ServiceDesc for KeywordService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeywordService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "keyword.KeywordService",
	HandlerType: (*KeywordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _KeywordService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _KeywordService_Update_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _KeywordService_GetById_Handler,
		},
		{
			MethodName: "GetListByChannel",
			Handler:    _KeywordService_GetListByChannel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "keyword.proto",
}