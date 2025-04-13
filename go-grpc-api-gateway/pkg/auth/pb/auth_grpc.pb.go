// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: pkg/auth/pb/auth.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AuthServiceP_Register_FullMethodName = "/auth.AuthServiceP/Register"
	AuthServiceP_Login_FullMethodName    = "/auth.AuthServiceP/Login"
	AuthServiceP_Validate_FullMethodName = "/auth.AuthServiceP/Validate"
)

// AuthServicePClient is the client API for AuthServiceP service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServicePClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
}

type authServicePClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServicePClient(cc grpc.ClientConnInterface) AuthServicePClient {
	return &authServicePClient{cc}
}

func (c *authServicePClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, AuthServiceP_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServicePClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, AuthServiceP_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServicePClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidateResponse)
	err := c.cc.Invoke(ctx, AuthServiceP_Validate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServicePServer is the server API for AuthServiceP service.
// All implementations must embed UnimplementedAuthServicePServer
// for forward compatibility.
type AuthServicePServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Validate(context.Context, *ValidateRequest) (*ValidateResponse, error)
	mustEmbedUnimplementedAuthServicePServer()
}

// UnimplementedAuthServicePServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServicePServer struct{}

func (UnimplementedAuthServicePServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthServicePServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServicePServer) Validate(context.Context, *ValidateRequest) (*ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedAuthServicePServer) mustEmbedUnimplementedAuthServicePServer() {}
func (UnimplementedAuthServicePServer) testEmbeddedByValue()                      {}

// UnsafeAuthServicePServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServicePServer will
// result in compilation errors.
type UnsafeAuthServicePServer interface {
	mustEmbedUnimplementedAuthServicePServer()
}

func RegisterAuthServicePServer(s grpc.ServiceRegistrar, srv AuthServicePServer) {
	// If the following call pancis, it indicates UnimplementedAuthServicePServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthServiceP_ServiceDesc, srv)
}

func _AuthServiceP_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServicePServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthServiceP_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServicePServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServiceP_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServicePServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthServiceP_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServicePServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServiceP_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServicePServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthServiceP_Validate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServicePServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthServiceP_ServiceDesc is the grpc.ServiceDesc for AuthServiceP service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthServiceP_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthServiceP",
	HandlerType: (*AuthServicePServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _AuthServiceP_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthServiceP_Login_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _AuthServiceP_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/auth/pb/auth.proto",
}
