// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: bookstore/bookstore.proto

package bookstore

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

// BookStoreClient is the client API for BookStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookStoreClient interface {
	ListBooks(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BookList, error)
	GetBook(ctx context.Context, in *BookId, opts ...grpc.CallOption) (*Book, error)
	CreateBook(ctx context.Context, in *CreateBookInput, opts ...grpc.CallOption) (*Book, error)
	UpdateBook(ctx context.Context, in *UpdateBookInput, opts ...grpc.CallOption) (*Book, error)
	DeleteBook(ctx context.Context, in *BookId, opts ...grpc.CallOption) (*Empty, error)
}

type bookStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewBookStoreClient(cc grpc.ClientConnInterface) BookStoreClient {
	return &bookStoreClient{cc}
}

func (c *bookStoreClient) ListBooks(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BookList, error) {
	out := new(BookList)
	err := c.cc.Invoke(ctx, "/BookStore/ListBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookStoreClient) GetBook(ctx context.Context, in *BookId, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/BookStore/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookStoreClient) CreateBook(ctx context.Context, in *CreateBookInput, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/BookStore/CreateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookStoreClient) UpdateBook(ctx context.Context, in *UpdateBookInput, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/BookStore/UpdateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookStoreClient) DeleteBook(ctx context.Context, in *BookId, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/BookStore/DeleteBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookStoreServer is the server API for BookStore service.
// All implementations must embed UnimplementedBookStoreServer
// for forward compatibility
type BookStoreServer interface {
	ListBooks(context.Context, *Empty) (*BookList, error)
	GetBook(context.Context, *BookId) (*Book, error)
	CreateBook(context.Context, *CreateBookInput) (*Book, error)
	UpdateBook(context.Context, *UpdateBookInput) (*Book, error)
	DeleteBook(context.Context, *BookId) (*Empty, error)
	mustEmbedUnimplementedBookStoreServer()
}

// UnimplementedBookStoreServer must be embedded to have forward compatible implementations.
type UnimplementedBookStoreServer struct {
}

func (UnimplementedBookStoreServer) ListBooks(context.Context, *Empty) (*BookList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBooks not implemented")
}
func (UnimplementedBookStoreServer) GetBook(context.Context, *BookId) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedBookStoreServer) CreateBook(context.Context, *CreateBookInput) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (UnimplementedBookStoreServer) UpdateBook(context.Context, *UpdateBookInput) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (UnimplementedBookStoreServer) DeleteBook(context.Context, *BookId) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (UnimplementedBookStoreServer) mustEmbedUnimplementedBookStoreServer() {}

// UnsafeBookStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookStoreServer will
// result in compilation errors.
type UnsafeBookStoreServer interface {
	mustEmbedUnimplementedBookStoreServer()
}

func RegisterBookStoreServer(s grpc.ServiceRegistrar, srv BookStoreServer) {
	s.RegisterService(&BookStore_ServiceDesc, srv)
}

func _BookStore_ListBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookStoreServer).ListBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BookStore/ListBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookStoreServer).ListBooks(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookStore_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookStoreServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BookStore/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookStoreServer).GetBook(ctx, req.(*BookId))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookStore_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookStoreServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BookStore/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookStoreServer).CreateBook(ctx, req.(*CreateBookInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookStore_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookStoreServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BookStore/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookStoreServer).UpdateBook(ctx, req.(*UpdateBookInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookStore_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookStoreServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BookStore/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookStoreServer).DeleteBook(ctx, req.(*BookId))
	}
	return interceptor(ctx, in, info, handler)
}

// BookStore_ServiceDesc is the grpc.ServiceDesc for BookStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BookStore",
	HandlerType: (*BookStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBooks",
			Handler:    _BookStore_ListBooks_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _BookStore_GetBook_Handler,
		},
		{
			MethodName: "CreateBook",
			Handler:    _BookStore_CreateBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _BookStore_UpdateBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _BookStore_DeleteBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bookstore/bookstore.proto",
}
