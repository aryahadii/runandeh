// Code generated by protoc-gen-go. DO NOT EDIT.
// source: saman.proto

/*
Package saman is a generated protocol buffer package.

It is generated from these files:
	saman.proto

It has these top-level messages:
	ReviewIDsBySortTypeRequest
	ReviewIDsBySortTypeReply
	HasReviewBySortTypeRequest
	HasReviewBySortTypeReply
	DeleteReviewRequest
	DeleteReviewReply
*/
package saman

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

type ReviewIDsBySortTypeRequest struct {
	AppID    int32  `protobuf:"varint,1,opt,name=appID" json:"appID,omitempty"`
	SortType string `protobuf:"bytes,2,opt,name=sortType" json:"sortType,omitempty"`
	Limit    int32  `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Offset   int32  `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
}

func (m *ReviewIDsBySortTypeRequest) Reset()                    { *m = ReviewIDsBySortTypeRequest{} }
func (m *ReviewIDsBySortTypeRequest) String() string            { return proto.CompactTextString(m) }
func (*ReviewIDsBySortTypeRequest) ProtoMessage()               {}
func (*ReviewIDsBySortTypeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ReviewIDsBySortTypeRequest) GetAppID() int32 {
	if m != nil {
		return m.AppID
	}
	return 0
}

func (m *ReviewIDsBySortTypeRequest) GetSortType() string {
	if m != nil {
		return m.SortType
	}
	return ""
}

func (m *ReviewIDsBySortTypeRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ReviewIDsBySortTypeRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ReviewIDsBySortTypeReply struct {
	ReviewIDs []int32 `protobuf:"varint,1,rep,packed,name=reviewIDs" json:"reviewIDs,omitempty"`
}

func (m *ReviewIDsBySortTypeReply) Reset()                    { *m = ReviewIDsBySortTypeReply{} }
func (m *ReviewIDsBySortTypeReply) String() string            { return proto.CompactTextString(m) }
func (*ReviewIDsBySortTypeReply) ProtoMessage()               {}
func (*ReviewIDsBySortTypeReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReviewIDsBySortTypeReply) GetReviewIDs() []int32 {
	if m != nil {
		return m.ReviewIDs
	}
	return nil
}

type HasReviewBySortTypeRequest struct {
	AppID    int32  `protobuf:"varint,1,opt,name=appID" json:"appID,omitempty"`
	SortType string `protobuf:"bytes,2,opt,name=sortType" json:"sortType,omitempty"`
}

func (m *HasReviewBySortTypeRequest) Reset()                    { *m = HasReviewBySortTypeRequest{} }
func (m *HasReviewBySortTypeRequest) String() string            { return proto.CompactTextString(m) }
func (*HasReviewBySortTypeRequest) ProtoMessage()               {}
func (*HasReviewBySortTypeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HasReviewBySortTypeRequest) GetAppID() int32 {
	if m != nil {
		return m.AppID
	}
	return 0
}

func (m *HasReviewBySortTypeRequest) GetSortType() string {
	if m != nil {
		return m.SortType
	}
	return ""
}

type HasReviewBySortTypeReply struct {
	Result bool `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *HasReviewBySortTypeReply) Reset()                    { *m = HasReviewBySortTypeReply{} }
func (m *HasReviewBySortTypeReply) String() string            { return proto.CompactTextString(m) }
func (*HasReviewBySortTypeReply) ProtoMessage()               {}
func (*HasReviewBySortTypeReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *HasReviewBySortTypeReply) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type DeleteReviewRequest struct {
	ReviewID int32 `protobuf:"varint,1,opt,name=reviewID" json:"reviewID,omitempty"`
}

func (m *DeleteReviewRequest) Reset()                    { *m = DeleteReviewRequest{} }
func (m *DeleteReviewRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteReviewRequest) ProtoMessage()               {}
func (*DeleteReviewRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteReviewRequest) GetReviewID() int32 {
	if m != nil {
		return m.ReviewID
	}
	return 0
}

type DeleteReviewReply struct {
	Result bool `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *DeleteReviewReply) Reset()                    { *m = DeleteReviewReply{} }
func (m *DeleteReviewReply) String() string            { return proto.CompactTextString(m) }
func (*DeleteReviewReply) ProtoMessage()               {}
func (*DeleteReviewReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeleteReviewReply) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func init() {
	proto.RegisterType((*ReviewIDsBySortTypeRequest)(nil), "saman.ReviewIDsBySortTypeRequest")
	proto.RegisterType((*ReviewIDsBySortTypeReply)(nil), "saman.ReviewIDsBySortTypeReply")
	proto.RegisterType((*HasReviewBySortTypeRequest)(nil), "saman.HasReviewBySortTypeRequest")
	proto.RegisterType((*HasReviewBySortTypeReply)(nil), "saman.HasReviewBySortTypeReply")
	proto.RegisterType((*DeleteReviewRequest)(nil), "saman.DeleteReviewRequest")
	proto.RegisterType((*DeleteReviewReply)(nil), "saman.DeleteReviewReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SamanService service

type SamanServiceClient interface {
	GetReviewIDsBySortType(ctx context.Context, in *ReviewIDsBySortTypeRequest, opts ...grpc.CallOption) (*ReviewIDsBySortTypeReply, error)
	HasReviewBySortType(ctx context.Context, in *HasReviewBySortTypeRequest, opts ...grpc.CallOption) (*HasReviewBySortTypeReply, error)
	DeleteReview(ctx context.Context, in *DeleteReviewRequest, opts ...grpc.CallOption) (*DeleteReviewReply, error)
}

type samanServiceClient struct {
	cc *grpc.ClientConn
}

func NewSamanServiceClient(cc *grpc.ClientConn) SamanServiceClient {
	return &samanServiceClient{cc}
}

func (c *samanServiceClient) GetReviewIDsBySortType(ctx context.Context, in *ReviewIDsBySortTypeRequest, opts ...grpc.CallOption) (*ReviewIDsBySortTypeReply, error) {
	out := new(ReviewIDsBySortTypeReply)
	err := grpc.Invoke(ctx, "/saman.SamanService/GetReviewIDsBySortType", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *samanServiceClient) HasReviewBySortType(ctx context.Context, in *HasReviewBySortTypeRequest, opts ...grpc.CallOption) (*HasReviewBySortTypeReply, error) {
	out := new(HasReviewBySortTypeReply)
	err := grpc.Invoke(ctx, "/saman.SamanService/HasReviewBySortType", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *samanServiceClient) DeleteReview(ctx context.Context, in *DeleteReviewRequest, opts ...grpc.CallOption) (*DeleteReviewReply, error) {
	out := new(DeleteReviewReply)
	err := grpc.Invoke(ctx, "/saman.SamanService/DeleteReview", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SamanService service

type SamanServiceServer interface {
	GetReviewIDsBySortType(context.Context, *ReviewIDsBySortTypeRequest) (*ReviewIDsBySortTypeReply, error)
	HasReviewBySortType(context.Context, *HasReviewBySortTypeRequest) (*HasReviewBySortTypeReply, error)
	DeleteReview(context.Context, *DeleteReviewRequest) (*DeleteReviewReply, error)
}

func RegisterSamanServiceServer(s *grpc.Server, srv SamanServiceServer) {
	s.RegisterService(&_SamanService_serviceDesc, srv)
}

func _SamanService_GetReviewIDsBySortType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewIDsBySortTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SamanServiceServer).GetReviewIDsBySortType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/saman.SamanService/GetReviewIDsBySortType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SamanServiceServer).GetReviewIDsBySortType(ctx, req.(*ReviewIDsBySortTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SamanService_HasReviewBySortType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasReviewBySortTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SamanServiceServer).HasReviewBySortType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/saman.SamanService/HasReviewBySortType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SamanServiceServer).HasReviewBySortType(ctx, req.(*HasReviewBySortTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SamanService_DeleteReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SamanServiceServer).DeleteReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/saman.SamanService/DeleteReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SamanServiceServer).DeleteReview(ctx, req.(*DeleteReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SamanService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "saman.SamanService",
	HandlerType: (*SamanServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReviewIDsBySortType",
			Handler:    _SamanService_GetReviewIDsBySortType_Handler,
		},
		{
			MethodName: "HasReviewBySortType",
			Handler:    _SamanService_HasReviewBySortType_Handler,
		},
		{
			MethodName: "DeleteReview",
			Handler:    _SamanService_DeleteReview_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "saman.proto",
}

func init() { proto.RegisterFile("saman.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x9b, 0xd6, 0x84, 0x74, 0xec, 0xc5, 0xad, 0x84, 0x65, 0x11, 0x8c, 0x7b, 0x0a, 0x08,
	0x05, 0xeb, 0xc5, 0xb3, 0x04, 0xb5, 0x17, 0x0f, 0x1b, 0x6f, 0x82, 0x10, 0x65, 0x0a, 0x81, 0xd4,
	0xac, 0xd9, 0x6d, 0x25, 0xe0, 0xb7, 0xf0, 0x0b, 0x4b, 0x36, 0x7f, 0x54, 0xd8, 0xf4, 0xd4, 0xe3,
	0x6f, 0x67, 0xde, 0x9b, 0xc9, 0x9b, 0xc0, 0xb1, 0x4a, 0x37, 0xe9, 0xfb, 0x42, 0x96, 0x85, 0x2e,
	0x88, 0x6b, 0x80, 0x7f, 0x01, 0x13, 0xb8, 0xcb, 0xf0, 0x73, 0x15, 0xab, 0xdb, 0x2a, 0x29, 0x4a,
	0xfd, 0x54, 0x49, 0x14, 0xf8, 0xb1, 0x45, 0xa5, 0xc9, 0x29, 0xb8, 0xa9, 0x94, 0xab, 0x98, 0x3a,
	0xa1, 0x13, 0xb9, 0xa2, 0x01, 0xc2, 0xc0, 0x57, 0x6d, 0x23, 0x1d, 0x87, 0x4e, 0x34, 0x15, 0x3d,
	0xd7, 0x8a, 0x3c, 0xdb, 0x64, 0x9a, 0x4e, 0x1a, 0x85, 0x01, 0x12, 0x80, 0x57, 0xac, 0xd7, 0x0a,
	0x35, 0x3d, 0x32, 0xcf, 0x2d, 0xf1, 0x1b, 0xa0, 0xd6, 0xe9, 0x32, 0xaf, 0xc8, 0x19, 0x4c, 0xcb,
	0xae, 0x46, 0x9d, 0x70, 0x12, 0xb9, 0xe2, 0xf7, 0x81, 0x3f, 0x02, 0x7b, 0x48, 0x55, 0x23, 0x3e,
	0xc0, 0xde, 0x7c, 0x09, 0xd4, 0xea, 0x57, 0x6f, 0x12, 0x80, 0x57, 0xa2, 0xda, 0xe6, 0xda, 0xd8,
	0xf9, 0xa2, 0x25, 0x7e, 0x05, 0xf3, 0x18, 0x73, 0xd4, 0xd8, 0xc8, 0xba, 0xe1, 0x0c, 0xfc, 0x6e,
	0xcf, 0x76, 0x7e, 0xcf, 0xfc, 0x12, 0x4e, 0xfe, 0x4b, 0xf6, 0xf8, 0x2f, 0xbf, 0xc7, 0x30, 0x4b,
	0xea, 0x2b, 0x25, 0x58, 0xee, 0xb2, 0x37, 0x24, 0x2f, 0x10, 0xdc, 0xa3, 0xb6, 0x24, 0x46, 0x2e,
	0x16, 0xcd, 0x6d, 0x87, 0x6f, 0xc9, 0xce, 0xf7, 0xb5, 0xc8, 0xbc, 0xe2, 0x23, 0xf2, 0x0c, 0x73,
	0x4b, 0x08, 0xbd, 0xf9, 0x70, 0xe0, 0xbd, 0xf9, 0x50, 0x86, 0x7c, 0x44, 0xee, 0x60, 0xf6, 0xf7,
	0xd3, 0x09, 0x6b, 0x25, 0x96, 0x08, 0x19, 0xb5, 0xd6, 0x8c, 0xcf, 0xab, 0x67, 0xfe, 0xdf, 0xeb,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x84, 0xc7, 0x96, 0x94, 0xce, 0x02, 0x00, 0x00,
}
