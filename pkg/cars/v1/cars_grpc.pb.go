// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: cars.proto

package v1_cars

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

// CarsServiceClient is the client API for CarsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CarsServiceClient interface {
	CreateCar(ctx context.Context, in *Car, opts ...grpc.CallOption) (*Car, error)
	ListCars(ctx context.Context, in *ListCarsReq, opts ...grpc.CallOption) (*Cars, error)
	GetCar(ctx context.Context, in *CarReq, opts ...grpc.CallOption) (*Car, error)
	DeleteCar(ctx context.Context, in *CarReq, opts ...grpc.CallOption) (*Success, error)
}

type carsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCarsServiceClient(cc grpc.ClientConnInterface) CarsServiceClient {
	return &carsServiceClient{cc}
}

func (c *carsServiceClient) CreateCar(ctx context.Context, in *Car, opts ...grpc.CallOption) (*Car, error) {
	out := new(Car)
	err := c.cc.Invoke(ctx, "/cars.CarsService/CreateCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carsServiceClient) ListCars(ctx context.Context, in *ListCarsReq, opts ...grpc.CallOption) (*Cars, error) {
	out := new(Cars)
	err := c.cc.Invoke(ctx, "/cars.CarsService/ListCars", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carsServiceClient) GetCar(ctx context.Context, in *CarReq, opts ...grpc.CallOption) (*Car, error) {
	out := new(Car)
	err := c.cc.Invoke(ctx, "/cars.CarsService/GetCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carsServiceClient) DeleteCar(ctx context.Context, in *CarReq, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/cars.CarsService/DeleteCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CarsServiceServer is the server API for CarsService service.
// All implementations must embed UnimplementedCarsServiceServer
// for forward compatibility
type CarsServiceServer interface {
	CreateCar(context.Context, *Car) (*Car, error)
	ListCars(context.Context, *ListCarsReq) (*Cars, error)
	GetCar(context.Context, *CarReq) (*Car, error)
	DeleteCar(context.Context, *CarReq) (*Success, error)
	mustEmbedUnimplementedCarsServiceServer()
}

// UnimplementedCarsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCarsServiceServer struct {
}

func (UnimplementedCarsServiceServer) CreateCar(context.Context, *Car) (*Car, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCar not implemented")
}
func (UnimplementedCarsServiceServer) ListCars(context.Context, *ListCarsReq) (*Cars, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCars not implemented")
}
func (UnimplementedCarsServiceServer) GetCar(context.Context, *CarReq) (*Car, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCar not implemented")
}
func (UnimplementedCarsServiceServer) DeleteCar(context.Context, *CarReq) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCar not implemented")
}
func (UnimplementedCarsServiceServer) mustEmbedUnimplementedCarsServiceServer() {}

// UnsafeCarsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CarsServiceServer will
// result in compilation errors.
type UnsafeCarsServiceServer interface {
	mustEmbedUnimplementedCarsServiceServer()
}

func RegisterCarsServiceServer(s grpc.ServiceRegistrar, srv CarsServiceServer) {
	s.RegisterService(&CarsService_ServiceDesc, srv)
}

func _CarsService_CreateCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Car)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarsServiceServer).CreateCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cars.CarsService/CreateCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarsServiceServer).CreateCar(ctx, req.(*Car))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarsService_ListCars_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCarsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarsServiceServer).ListCars(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cars.CarsService/ListCars",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarsServiceServer).ListCars(ctx, req.(*ListCarsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarsService_GetCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CarReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarsServiceServer).GetCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cars.CarsService/GetCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarsServiceServer).GetCar(ctx, req.(*CarReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarsService_DeleteCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CarReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarsServiceServer).DeleteCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cars.CarsService/DeleteCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarsServiceServer).DeleteCar(ctx, req.(*CarReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CarsService_ServiceDesc is the grpc.ServiceDesc for CarsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CarsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cars.CarsService",
	HandlerType: (*CarsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCar",
			Handler:    _CarsService_CreateCar_Handler,
		},
		{
			MethodName: "ListCars",
			Handler:    _CarsService_ListCars_Handler,
		},
		{
			MethodName: "GetCar",
			Handler:    _CarsService_GetCar_Handler,
		},
		{
			MethodName: "DeleteCar",
			Handler:    _CarsService_DeleteCar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cars.proto",
}
