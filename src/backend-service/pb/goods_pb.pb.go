// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/goods_pb.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GoodsRequest struct {
	GoodsUuid            string   `protobuf:"bytes,1,opt,name=GoodsUuid,proto3" json:"GoodsUuid,omitempty"`
	GoodsFrom            string   `protobuf:"bytes,2,opt,name=GoodsFrom,proto3" json:"GoodsFrom,omitempty"`
	GoodsTypeId          int64    `protobuf:"varint,3,opt,name=GoodsTypeId,proto3" json:"GoodsTypeId,omitempty"`
	PrimaryType          string   `protobuf:"bytes,4,opt,name=PrimaryType,proto3" json:"PrimaryType,omitempty"`
	SecondaryType        string   `protobuf:"bytes,5,opt,name=SecondaryType,proto3" json:"SecondaryType,omitempty"`
	Img                  string   `protobuf:"bytes,6,opt,name=Img,proto3" json:"Img,omitempty"`
	Imgs                 string   `protobuf:"bytes,7,opt,name=Imgs,proto3" json:"Imgs,omitempty"`
	IsValid              bool     `protobuf:"varint,8,opt,name=IsValid,proto3" json:"IsValid,omitempty"`
	Title                string   `protobuf:"bytes,9,opt,name=Title,proto3" json:"Title,omitempty"`
	Subtitle             string   `protobuf:"bytes,10,opt,name=Subtitle,proto3" json:"Subtitle,omitempty"`
	Price                int64    `protobuf:"varint,11,opt,name=Price,proto3" json:"Price,omitempty"`
	PublishDate          string   `protobuf:"bytes,12,opt,name=PublishDate,proto3" json:"PublishDate,omitempty"`
	PageIndex            int64    `protobuf:"varint,13,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	PageSize             int64    `protobuf:"varint,14,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoodsRequest) Reset()         { *m = GoodsRequest{} }
func (m *GoodsRequest) String() string { return proto.CompactTextString(m) }
func (*GoodsRequest) ProtoMessage()    {}
func (*GoodsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_088a779269859b27, []int{0}
}

func (m *GoodsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoodsRequest.Unmarshal(m, b)
}
func (m *GoodsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoodsRequest.Marshal(b, m, deterministic)
}
func (m *GoodsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodsRequest.Merge(m, src)
}
func (m *GoodsRequest) XXX_Size() int {
	return xxx_messageInfo_GoodsRequest.Size(m)
}
func (m *GoodsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GoodsRequest proto.InternalMessageInfo

func (m *GoodsRequest) GetGoodsUuid() string {
	if m != nil {
		return m.GoodsUuid
	}
	return ""
}

func (m *GoodsRequest) GetGoodsFrom() string {
	if m != nil {
		return m.GoodsFrom
	}
	return ""
}

func (m *GoodsRequest) GetGoodsTypeId() int64 {
	if m != nil {
		return m.GoodsTypeId
	}
	return 0
}

func (m *GoodsRequest) GetPrimaryType() string {
	if m != nil {
		return m.PrimaryType
	}
	return ""
}

func (m *GoodsRequest) GetSecondaryType() string {
	if m != nil {
		return m.SecondaryType
	}
	return ""
}

func (m *GoodsRequest) GetImg() string {
	if m != nil {
		return m.Img
	}
	return ""
}

func (m *GoodsRequest) GetImgs() string {
	if m != nil {
		return m.Imgs
	}
	return ""
}

func (m *GoodsRequest) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

func (m *GoodsRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *GoodsRequest) GetSubtitle() string {
	if m != nil {
		return m.Subtitle
	}
	return ""
}

func (m *GoodsRequest) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *GoodsRequest) GetPublishDate() string {
	if m != nil {
		return m.PublishDate
	}
	return ""
}

func (m *GoodsRequest) GetPageIndex() int64 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *GoodsRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type GoodsReplyItem struct {
	GoodsUuid            string   `protobuf:"bytes,1,opt,name=GoodsUuid,proto3" json:"GoodsUuid,omitempty"`
	GoodsFrom            string   `protobuf:"bytes,2,opt,name=GoodsFrom,proto3" json:"GoodsFrom,omitempty"`
	GoodsTypeId          int64    `protobuf:"varint,3,opt,name=GoodsTypeId,proto3" json:"GoodsTypeId,omitempty"`
	PrimaryType          string   `protobuf:"bytes,4,opt,name=PrimaryType,proto3" json:"PrimaryType,omitempty"`
	SecondaryType        string   `protobuf:"bytes,5,opt,name=SecondaryType,proto3" json:"SecondaryType,omitempty"`
	Img                  string   `protobuf:"bytes,6,opt,name=Img,proto3" json:"Img,omitempty"`
	Imgs                 string   `protobuf:"bytes,7,opt,name=Imgs,proto3" json:"Imgs,omitempty"`
	IsValid              bool     `protobuf:"varint,8,opt,name=IsValid,proto3" json:"IsValid,omitempty"`
	Title                string   `protobuf:"bytes,9,opt,name=Title,proto3" json:"Title,omitempty"`
	Subtitle             string   `protobuf:"bytes,10,opt,name=Subtitle,proto3" json:"Subtitle,omitempty"`
	Price                int64    `protobuf:"varint,11,opt,name=Price,proto3" json:"Price,omitempty"`
	PublishDate          string   `protobuf:"bytes,12,opt,name=PublishDate,proto3" json:"PublishDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoodsReplyItem) Reset()         { *m = GoodsReplyItem{} }
func (m *GoodsReplyItem) String() string { return proto.CompactTextString(m) }
func (*GoodsReplyItem) ProtoMessage()    {}
func (*GoodsReplyItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_088a779269859b27, []int{1}
}

func (m *GoodsReplyItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoodsReplyItem.Unmarshal(m, b)
}
func (m *GoodsReplyItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoodsReplyItem.Marshal(b, m, deterministic)
}
func (m *GoodsReplyItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodsReplyItem.Merge(m, src)
}
func (m *GoodsReplyItem) XXX_Size() int {
	return xxx_messageInfo_GoodsReplyItem.Size(m)
}
func (m *GoodsReplyItem) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodsReplyItem.DiscardUnknown(m)
}

var xxx_messageInfo_GoodsReplyItem proto.InternalMessageInfo

func (m *GoodsReplyItem) GetGoodsUuid() string {
	if m != nil {
		return m.GoodsUuid
	}
	return ""
}

func (m *GoodsReplyItem) GetGoodsFrom() string {
	if m != nil {
		return m.GoodsFrom
	}
	return ""
}

func (m *GoodsReplyItem) GetGoodsTypeId() int64 {
	if m != nil {
		return m.GoodsTypeId
	}
	return 0
}

func (m *GoodsReplyItem) GetPrimaryType() string {
	if m != nil {
		return m.PrimaryType
	}
	return ""
}

func (m *GoodsReplyItem) GetSecondaryType() string {
	if m != nil {
		return m.SecondaryType
	}
	return ""
}

func (m *GoodsReplyItem) GetImg() string {
	if m != nil {
		return m.Img
	}
	return ""
}

func (m *GoodsReplyItem) GetImgs() string {
	if m != nil {
		return m.Imgs
	}
	return ""
}

func (m *GoodsReplyItem) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

func (m *GoodsReplyItem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *GoodsReplyItem) GetSubtitle() string {
	if m != nil {
		return m.Subtitle
	}
	return ""
}

func (m *GoodsReplyItem) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *GoodsReplyItem) GetPublishDate() string {
	if m != nil {
		return m.PublishDate
	}
	return ""
}

type GoodsReply struct {
	Data                 []*GoodsReplyItem `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GoodsReply) Reset()         { *m = GoodsReply{} }
func (m *GoodsReply) String() string { return proto.CompactTextString(m) }
func (*GoodsReply) ProtoMessage()    {}
func (*GoodsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_088a779269859b27, []int{2}
}

func (m *GoodsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoodsReply.Unmarshal(m, b)
}
func (m *GoodsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoodsReply.Marshal(b, m, deterministic)
}
func (m *GoodsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodsReply.Merge(m, src)
}
func (m *GoodsReply) XXX_Size() int {
	return xxx_messageInfo_GoodsReply.Size(m)
}
func (m *GoodsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodsReply.DiscardUnknown(m)
}

var xxx_messageInfo_GoodsReply proto.InternalMessageInfo

func (m *GoodsReply) GetData() []*GoodsReplyItem {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GoodsRequest)(nil), "pb.GoodsRequest")
	proto.RegisterType((*GoodsReplyItem)(nil), "pb.GoodsReplyItem")
	proto.RegisterType((*GoodsReply)(nil), "pb.GoodsReply")
}

func init() { proto.RegisterFile("pb/goods_pb.proto", fileDescriptor_088a779269859b27) }

var fileDescriptor_088a779269859b27 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x93, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0xc6, 0x2d, 0xcb, 0x7f, 0xc7, 0xb6, 0xea, 0x0e, 0x3d, 0x2c, 0xa6, 0x07, 0x21, 0x4a, 0xd1,
	0xc9, 0x05, 0xb7, 0x85, 0x3e, 0x80, 0xa9, 0x2b, 0xe8, 0xc1, 0xc8, 0x6e, 0xaf, 0x41, 0xb2, 0x16,
	0x65, 0x41, 0xb2, 0x36, 0xd2, 0x0a, 0xe2, 0x3c, 0x49, 0x1e, 0x22, 0x0f, 0x95, 0x47, 0x09, 0x3b,
	0xb6, 0x2c, 0x19, 0x7c, 0xc8, 0x21, 0xc7, 0xdc, 0xf6, 0xfb, 0x7d, 0xf3, 0xed, 0x32, 0x33, 0x2c,
	0x7c, 0x94, 0xe1, 0xb7, 0x38, 0xcb, 0xa2, 0xe2, 0x46, 0x86, 0x73, 0x99, 0x67, 0x2a, 0xc3, 0xb6,
	0x0c, 0x9d, 0x47, 0x13, 0xc6, 0x2b, 0x8d, 0x7d, 0x7e, 0x57, 0xf2, 0x42, 0xe1, 0x67, 0x18, 0x92,
	0xfe, 0x57, 0x8a, 0x88, 0x19, 0xb6, 0xe1, 0x0e, 0xfd, 0x1a, 0x9c, 0xdd, 0xdf, 0x79, 0x96, 0xb2,
	0x76, 0xc3, 0xd5, 0x00, 0x6d, 0x18, 0x91, 0xd8, 0x1e, 0x24, 0xf7, 0x22, 0x66, 0xda, 0x86, 0x6b,
	0xfa, 0x4d, 0xa4, 0x2b, 0xd6, 0xb9, 0x48, 0x83, 0xfc, 0xa0, 0x01, 0xeb, 0xd0, 0x0d, 0x4d, 0x84,
	0x5f, 0x60, 0xb2, 0xe1, 0xbb, 0x6c, 0x1f, 0x55, 0x35, 0x5d, 0xaa, 0xb9, 0x84, 0x38, 0x05, 0xd3,
	0x4b, 0x63, 0xd6, 0x23, 0x4f, 0x1f, 0x11, 0xa1, 0xe3, 0xa5, 0x71, 0xc1, 0xfa, 0x84, 0xe8, 0x8c,
	0x0c, 0xfa, 0x5e, 0xf1, 0x3f, 0x48, 0x44, 0xc4, 0x06, 0xb6, 0xe1, 0x0e, 0xfc, 0x4a, 0xe2, 0x27,
	0xe8, 0x6e, 0x85, 0x4a, 0x38, 0x1b, 0x52, 0xf9, 0x51, 0xe0, 0x0c, 0x06, 0x9b, 0x32, 0x54, 0x64,
	0x00, 0x19, 0x67, 0xad, 0x13, 0xeb, 0x5c, 0xec, 0x38, 0x1b, 0x51, 0x57, 0x47, 0x41, 0xfd, 0x94,
	0x61, 0x22, 0x8a, 0xdb, 0x65, 0xa0, 0x38, 0x1b, 0x9f, 0xfa, 0xa9, 0x91, 0x9e, 0x98, 0x0c, 0x62,
	0xee, 0xed, 0x23, 0x7e, 0xcf, 0x26, 0x94, 0xad, 0x81, 0x7e, 0x51, 0x8b, 0x8d, 0x78, 0xe0, 0xcc,
	0x22, 0xf3, 0xac, 0x9d, 0xe7, 0x36, 0x58, 0xa7, 0xd5, 0xc8, 0xe4, 0xe0, 0x29, 0x9e, 0xbe, 0x2f,
	0xe7, 0x6d, 0x97, 0xe3, 0xfc, 0x00, 0xa8, 0x27, 0x8c, 0x5f, 0xa1, 0x13, 0x05, 0x2a, 0x60, 0x86,
	0x6d, 0xba, 0xa3, 0x05, 0xce, 0x65, 0x38, 0xbf, 0x9c, 0xbf, 0x4f, 0xfe, 0xe2, 0xc9, 0x80, 0x2e,
	0x19, 0xb8, 0x80, 0xf1, 0x8a, 0x2b, 0x3a, 0xff, 0x15, 0x85, 0xc2, 0x69, 0x23, 0x43, 0xdf, 0x69,
	0x66, 0x5d, 0xde, 0xe2, 0xb4, 0xf0, 0x17, 0x58, 0x55, 0x66, 0xc9, 0x55, 0x20, 0x92, 0x2b, 0xa9,
	0x2b, 0x6f, 0x3b, 0x2d, 0xfc, 0x09, 0x1f, 0xaa, 0xe4, 0x9f, 0x4c, 0xbd, 0xf6, 0xc1, 0xb0, 0x47,
	0xbf, 0xfd, 0xfb, 0x4b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x99, 0xdd, 0xd0, 0xcf, 0x02, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GoodsClient is the client API for Goods service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GoodsClient interface {
	GetGoodsList(ctx context.Context, in *GoodsRequest, opts ...grpc.CallOption) (*GoodsReply, error)
	GetGoodsDetail(ctx context.Context, in *GoodsRequest, opts ...grpc.CallOption) (*GoodsReplyItem, error)
	GetGoodsHotList(ctx context.Context, in *GoodsRequest, opts ...grpc.CallOption) (*GoodsReply, error)
}

type goodsClient struct {
	cc *grpc.ClientConn
}

func NewGoodsClient(cc *grpc.ClientConn) GoodsClient {
	return &goodsClient{cc}
}

func (c *goodsClient) GetGoodsList(ctx context.Context, in *GoodsRequest, opts ...grpc.CallOption) (*GoodsReply, error) {
	out := new(GoodsReply)
	err := c.cc.Invoke(ctx, "/pb.Goods/GetGoodsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GetGoodsDetail(ctx context.Context, in *GoodsRequest, opts ...grpc.CallOption) (*GoodsReplyItem, error) {
	out := new(GoodsReplyItem)
	err := c.cc.Invoke(ctx, "/pb.Goods/GetGoodsDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GetGoodsHotList(ctx context.Context, in *GoodsRequest, opts ...grpc.CallOption) (*GoodsReply, error) {
	out := new(GoodsReply)
	err := c.cc.Invoke(ctx, "/pb.Goods/GetGoodsHotList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodsServer is the server API for Goods service.
type GoodsServer interface {
	GetGoodsList(context.Context, *GoodsRequest) (*GoodsReply, error)
	GetGoodsDetail(context.Context, *GoodsRequest) (*GoodsReplyItem, error)
	GetGoodsHotList(context.Context, *GoodsRequest) (*GoodsReply, error)
}

func RegisterGoodsServer(s *grpc.Server, srv GoodsServer) {
	s.RegisterService(&_Goods_serviceDesc, srv)
}

func _Goods_GetGoodsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetGoodsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Goods/GetGoodsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetGoodsList(ctx, req.(*GoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GetGoodsDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetGoodsDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Goods/GetGoodsDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetGoodsDetail(ctx, req.(*GoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GetGoodsHotList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetGoodsHotList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Goods/GetGoodsHotList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetGoodsHotList(ctx, req.(*GoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Goods_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Goods",
	HandlerType: (*GoodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGoodsList",
			Handler:    _Goods_GetGoodsList_Handler,
		},
		{
			MethodName: "GetGoodsDetail",
			Handler:    _Goods_GetGoodsDetail_Handler,
		},
		{
			MethodName: "GetGoodsHotList",
			Handler:    _Goods_GetGoodsHotList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/goods_pb.proto",
}