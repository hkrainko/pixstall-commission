// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServiceClient interface {
	SaveFile(ctx context.Context, opts ...grpc.CallOption) (FileService_SaveFileClient, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) SaveFile(ctx context.Context, opts ...grpc.CallOption) (FileService_SaveFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileService_serviceDesc.Streams[0], "/FileService/SaveFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceSaveFileClient{stream}
	return x, nil
}

type FileService_SaveFileClient interface {
	Send(*SaveFileRequest) error
	CloseAndRecv() (*SaveFileResponse, error)
	grpc.ClientStream
}

type fileServiceSaveFileClient struct {
	grpc.ClientStream
}

func (x *fileServiceSaveFileClient) Send(m *SaveFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileServiceSaveFileClient) CloseAndRecv() (*SaveFileResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SaveFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileServiceServer is the server API for FileService service.
// All implementations must embed UnimplementedFileServiceServer
// for forward compatibility
type FileServiceServer interface {
	SaveFile(FileService_SaveFileServer) error
	mustEmbedUnimplementedFileServiceServer()
}

// UnimplementedFileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (UnimplementedFileServiceServer) SaveFile(FileService_SaveFileServer) error {
	return status.Errorf(codes.Unimplemented, "method SaveFile not implemented")
}
func (UnimplementedFileServiceServer) mustEmbedUnimplementedFileServiceServer() {}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	s.RegisterService(&_FileService_serviceDesc, srv)
}

func _FileService_SaveFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServiceServer).SaveFile(&fileServiceSaveFileServer{stream})
}

type FileService_SaveFileServer interface {
	SendAndClose(*SaveFileResponse) error
	Recv() (*SaveFileRequest, error)
	grpc.ServerStream
}

type fileServiceSaveFileServer struct {
	grpc.ServerStream
}

func (x *fileServiceSaveFileServer) SendAndClose(m *SaveFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileServiceSaveFileServer) Recv() (*SaveFileRequest, error) {
	m := new(SaveFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _FileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SaveFile",
			Handler:       _FileService_SaveFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/file.proto",
}