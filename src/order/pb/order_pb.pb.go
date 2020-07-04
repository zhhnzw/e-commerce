// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/order_pb.proto

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

type OrderRequest struct {
	GoodsUuid            string   `protobuf:"bytes,1,opt,name=goodsUuid,proto3" json:"goodsUuid,omitempty"`
	GoodsTypeId          int64    `protobuf:"varint,2,opt,name=goodsTypeId,proto3" json:"goodsTypeId,omitempty"`
	PrimaryType          string   `protobuf:"bytes,3,opt,name=primaryType,proto3" json:"primaryType,omitempty"`
	SecondaryType        string   `protobuf:"bytes,4,opt,name=secondaryType,proto3" json:"secondaryType,omitempty"`
	Img                  string   `protobuf:"bytes,5,opt,name=img,proto3" json:"img,omitempty"`
	Title                string   `protobuf:"bytes,6,opt,name=title,proto3" json:"title,omitempty"`
	Subtitle             string   `protobuf:"bytes,7,opt,name=subtitle,proto3" json:"subtitle,omitempty"`
	Price                int64    `protobuf:"varint,8,opt,name=price,proto3" json:"price,omitempty"`
	OrderStatus          string   `protobuf:"bytes,9,opt,name=orderStatus,proto3" json:"orderStatus,omitempty"`
	PageIndex            int64    `protobuf:"varint,10,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	PageSize             int64    `protobuf:"varint,11,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderRequest) Reset()         { *m = OrderRequest{} }
func (m *OrderRequest) String() string { return proto.CompactTextString(m) }
func (*OrderRequest) ProtoMessage()    {}
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_55e302acbb9449ef, []int{0}
}

func (m *OrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderRequest.Unmarshal(m, b)
}
func (m *OrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderRequest.Marshal(b, m, deterministic)
}
func (m *OrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderRequest.Merge(m, src)
}
func (m *OrderRequest) XXX_Size() int {
	return xxx_messageInfo_OrderRequest.Size(m)
}
func (m *OrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OrderRequest proto.InternalMessageInfo

func (m *OrderRequest) GetGoodsUuid() string {
	if m != nil {
		return m.GoodsUuid
	}
	return ""
}

func (m *OrderRequest) GetGoodsTypeId() int64 {
	if m != nil {
		return m.GoodsTypeId
	}
	return 0
}

func (m *OrderRequest) GetPrimaryType() string {
	if m != nil {
		return m.PrimaryType
	}
	return ""
}

func (m *OrderRequest) GetSecondaryType() string {
	if m != nil {
		return m.SecondaryType
	}
	return ""
}

func (m *OrderRequest) GetImg() string {
	if m != nil {
		return m.Img
	}
	return ""
}

func (m *OrderRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *OrderRequest) GetSubtitle() string {
	if m != nil {
		return m.Subtitle
	}
	return ""
}

func (m *OrderRequest) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *OrderRequest) GetOrderStatus() string {
	if m != nil {
		return m.OrderStatus
	}
	return ""
}

func (m *OrderRequest) GetPageIndex() int64 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *OrderRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type OrderReplyItem struct {
	GoodsUuid            string   `protobuf:"bytes,1,opt,name=goodsUuid,proto3" json:"goodsUuid,omitempty"`
	GoodsTypeId          int64    `protobuf:"varint,2,opt,name=goodsTypeId,proto3" json:"goodsTypeId,omitempty"`
	PrimaryType          string   `protobuf:"bytes,3,opt,name=primaryType,proto3" json:"primaryType,omitempty"`
	SecondaryType        string   `protobuf:"bytes,4,opt,name=secondaryType,proto3" json:"secondaryType,omitempty"`
	Img                  string   `protobuf:"bytes,5,opt,name=img,proto3" json:"img,omitempty"`
	Title                string   `protobuf:"bytes,6,opt,name=title,proto3" json:"title,omitempty"`
	Subtitle             string   `protobuf:"bytes,7,opt,name=subtitle,proto3" json:"subtitle,omitempty"`
	Price                int64    `protobuf:"varint,8,opt,name=price,proto3" json:"price,omitempty"`
	OrderStatus          string   `protobuf:"bytes,9,opt,name=orderStatus,proto3" json:"orderStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderReplyItem) Reset()         { *m = OrderReplyItem{} }
func (m *OrderReplyItem) String() string { return proto.CompactTextString(m) }
func (*OrderReplyItem) ProtoMessage()    {}
func (*OrderReplyItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_55e302acbb9449ef, []int{1}
}

func (m *OrderReplyItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderReplyItem.Unmarshal(m, b)
}
func (m *OrderReplyItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderReplyItem.Marshal(b, m, deterministic)
}
func (m *OrderReplyItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderReplyItem.Merge(m, src)
}
func (m *OrderReplyItem) XXX_Size() int {
	return xxx_messageInfo_OrderReplyItem.Size(m)
}
func (m *OrderReplyItem) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderReplyItem.DiscardUnknown(m)
}

var xxx_messageInfo_OrderReplyItem proto.InternalMessageInfo

func (m *OrderReplyItem) GetGoodsUuid() string {
	if m != nil {
		return m.GoodsUuid
	}
	return ""
}

func (m *OrderReplyItem) GetGoodsTypeId() int64 {
	if m != nil {
		return m.GoodsTypeId
	}
	return 0
}

func (m *OrderReplyItem) GetPrimaryType() string {
	if m != nil {
		return m.PrimaryType
	}
	return ""
}

func (m *OrderReplyItem) GetSecondaryType() string {
	if m != nil {
		return m.SecondaryType
	}
	return ""
}

func (m *OrderReplyItem) GetImg() string {
	if m != nil {
		return m.Img
	}
	return ""
}

func (m *OrderReplyItem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *OrderReplyItem) GetSubtitle() string {
	if m != nil {
		return m.Subtitle
	}
	return ""
}

func (m *OrderReplyItem) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *OrderReplyItem) GetOrderStatus() string {
	if m != nil {
		return m.OrderStatus
	}
	return ""
}

type OrderReply struct {
	Data                 []*OrderReplyItem `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *OrderReply) Reset()         { *m = OrderReply{} }
func (m *OrderReply) String() string { return proto.CompactTextString(m) }
func (*OrderReply) ProtoMessage()    {}
func (*OrderReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_55e302acbb9449ef, []int{2}
}

func (m *OrderReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderReply.Unmarshal(m, b)
}
func (m *OrderReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderReply.Marshal(b, m, deterministic)
}
func (m *OrderReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderReply.Merge(m, src)
}
func (m *OrderReply) XXX_Size() int {
	return xxx_messageInfo_OrderReply.Size(m)
}
func (m *OrderReply) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderReply.DiscardUnknown(m)
}

var xxx_messageInfo_OrderReply proto.InternalMessageInfo

func (m *OrderReply) GetData() []*OrderReplyItem {
	if m != nil {
		return m.Data
	}
	return nil
}

type OrderStatisticItem struct {
	GoodsType            string   `protobuf:"bytes,1,opt,name=goodsType,proto3" json:"goodsType,omitempty"`
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderStatisticItem) Reset()         { *m = OrderStatisticItem{} }
func (m *OrderStatisticItem) String() string { return proto.CompactTextString(m) }
func (*OrderStatisticItem) ProtoMessage()    {}
func (*OrderStatisticItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_55e302acbb9449ef, []int{3}
}

func (m *OrderStatisticItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderStatisticItem.Unmarshal(m, b)
}
func (m *OrderStatisticItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderStatisticItem.Marshal(b, m, deterministic)
}
func (m *OrderStatisticItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderStatisticItem.Merge(m, src)
}
func (m *OrderStatisticItem) XXX_Size() int {
	return xxx_messageInfo_OrderStatisticItem.Size(m)
}
func (m *OrderStatisticItem) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderStatisticItem.DiscardUnknown(m)
}

var xxx_messageInfo_OrderStatisticItem proto.InternalMessageInfo

func (m *OrderStatisticItem) GetGoodsType() string {
	if m != nil {
		return m.GoodsType
	}
	return ""
}

func (m *OrderStatisticItem) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type OrderStatisticReply struct {
	Data                 []*OrderStatisticItem `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OrderStatisticReply) Reset()         { *m = OrderStatisticReply{} }
func (m *OrderStatisticReply) String() string { return proto.CompactTextString(m) }
func (*OrderStatisticReply) ProtoMessage()    {}
func (*OrderStatisticReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_55e302acbb9449ef, []int{4}
}

func (m *OrderStatisticReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderStatisticReply.Unmarshal(m, b)
}
func (m *OrderStatisticReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderStatisticReply.Marshal(b, m, deterministic)
}
func (m *OrderStatisticReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderStatisticReply.Merge(m, src)
}
func (m *OrderStatisticReply) XXX_Size() int {
	return xxx_messageInfo_OrderStatisticReply.Size(m)
}
func (m *OrderStatisticReply) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderStatisticReply.DiscardUnknown(m)
}

var xxx_messageInfo_OrderStatisticReply proto.InternalMessageInfo

func (m *OrderStatisticReply) GetData() []*OrderStatisticItem {
	if m != nil {
		return m.Data
	}
	return nil
}

type OrderCommonReply struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderCommonReply) Reset()         { *m = OrderCommonReply{} }
func (m *OrderCommonReply) String() string { return proto.CompactTextString(m) }
func (*OrderCommonReply) ProtoMessage()    {}
func (*OrderCommonReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_55e302acbb9449ef, []int{5}
}

func (m *OrderCommonReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderCommonReply.Unmarshal(m, b)
}
func (m *OrderCommonReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderCommonReply.Marshal(b, m, deterministic)
}
func (m *OrderCommonReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderCommonReply.Merge(m, src)
}
func (m *OrderCommonReply) XXX_Size() int {
	return xxx_messageInfo_OrderCommonReply.Size(m)
}
func (m *OrderCommonReply) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderCommonReply.DiscardUnknown(m)
}

var xxx_messageInfo_OrderCommonReply proto.InternalMessageInfo

func (m *OrderCommonReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*OrderRequest)(nil), "pb.orderRequest")
	proto.RegisterType((*OrderReplyItem)(nil), "pb.orderReplyItem")
	proto.RegisterType((*OrderReply)(nil), "pb.orderReply")
	proto.RegisterType((*OrderStatisticItem)(nil), "pb.orderStatisticItem")
	proto.RegisterType((*OrderStatisticReply)(nil), "pb.orderStatisticReply")
	proto.RegisterType((*OrderCommonReply)(nil), "pb.orderCommonReply")
}

func init() { proto.RegisterFile("pb/order_pb.proto", fileDescriptor_55e302acbb9449ef) }

var fileDescriptor_55e302acbb9449ef = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x54, 0xcd, 0xca, 0xd3, 0x40,
	0x14, 0xfd, 0x92, 0x7c, 0xf9, 0x6c, 0x6f, 0x6a, 0x49, 0xc7, 0xa2, 0x43, 0x71, 0x11, 0x42, 0x91,
	0xe2, 0xa2, 0x42, 0x15, 0xdc, 0x2a, 0x5d, 0x68, 0x41, 0x10, 0x52, 0x5d, 0x4b, 0x7e, 0x86, 0x30,
	0xd0, 0x64, 0xc6, 0xcc, 0x04, 0xac, 0xcf, 0xe1, 0xb3, 0xf8, 0x14, 0x3e, 0x94, 0xcc, 0x9d, 0x34,
	0x4d, 0xb0, 0x8f, 0xe0, 0x2e, 0xe7, 0x9c, 0x7b, 0xee, 0x9c, 0x7b, 0xa0, 0x85, 0x85, 0xcc, 0x5e,
	0x89, 0xa6, 0x60, 0xcd, 0x37, 0x99, 0x6d, 0x65, 0x23, 0xb4, 0x20, 0xae, 0xcc, 0xe2, 0x3f, 0x2e,
	0xcc, 0x90, 0x4e, 0xd8, 0xf7, 0x96, 0x29, 0x4d, 0x9e, 0xc3, 0xb4, 0x14, 0xa2, 0x50, 0x5f, 0x5b,
	0x5e, 0x50, 0x27, 0x72, 0x36, 0xd3, 0xe4, 0x4a, 0x90, 0x08, 0x02, 0x04, 0x5f, 0xce, 0x92, 0x1d,
	0x0a, 0xea, 0x46, 0xce, 0xc6, 0x4b, 0x86, 0x94, 0x99, 0x90, 0x0d, 0xaf, 0xd2, 0xe6, 0x6c, 0x08,
	0xea, 0xe1, 0x86, 0x21, 0x45, 0xd6, 0xf0, 0x58, 0xb1, 0x5c, 0xd4, 0xc5, 0x65, 0xe6, 0x1e, 0x67,
	0xc6, 0x24, 0x09, 0xc1, 0xe3, 0x55, 0x49, 0x7d, 0xd4, 0xcc, 0x27, 0x59, 0x82, 0xaf, 0xb9, 0x3e,
	0x31, 0xfa, 0x80, 0x9c, 0x05, 0x64, 0x05, 0x13, 0xd5, 0x66, 0x56, 0x78, 0x84, 0x42, 0x8f, 0x8d,
	0x43, 0x36, 0x3c, 0x67, 0x74, 0x82, 0x39, 0x2d, 0x30, 0x09, 0xf1, 0xe2, 0xa3, 0x4e, 0x75, 0xab,
	0xe8, 0xd4, 0x26, 0x1c, 0x50, 0xa6, 0x03, 0x99, 0x96, 0xec, 0x50, 0x17, 0xec, 0x07, 0x05, 0xf4,
	0x5e, 0x09, 0xf3, 0xa2, 0x01, 0x47, 0xfe, 0x93, 0xd1, 0x00, 0xc5, 0x1e, 0xc7, 0xbf, 0x5c, 0x98,
	0x77, 0x75, 0xca, 0xd3, 0xf9, 0xa0, 0x59, 0xf5, 0xbf, 0xd0, 0x56, 0xc5, 0x6f, 0x00, 0xae, 0xad,
	0x90, 0x17, 0x70, 0x5f, 0xa4, 0x3a, 0xa5, 0x4e, 0xe4, 0x6d, 0x82, 0x1d, 0xd9, 0xca, 0x6c, 0x3b,
	0xee, 0x2c, 0x41, 0x3d, 0xfe, 0x08, 0xa4, 0x5f, 0xc2, 0x95, 0xe6, 0xf9, 0xa8, 0x4f, 0xbc, 0x74,
	0xd8, 0x27, 0x5e, 0xb9, 0x04, 0x3f, 0x17, 0x6d, 0xad, 0xbb, 0x26, 0x2d, 0x88, 0xdf, 0xc3, 0x93,
	0xf1, 0x26, 0x1b, 0xe4, 0xe5, 0x28, 0xc8, 0xd3, 0x3e, 0xc8, 0xe8, 0xc1, 0x2e, 0xcc, 0x1a, 0x42,
	0xd4, 0xf6, 0xa2, 0xaa, 0x44, 0x6d, 0xfd, 0x21, 0x78, 0x95, 0x2a, 0xbb, 0x10, 0xe6, 0x73, 0xf7,
	0xdb, 0x01, 0xff, 0xb3, 0x19, 0x23, 0x6f, 0x21, 0xd8, 0x37, 0x2c, 0xd5, 0xcc, 0xc2, 0x70, 0x70,
	0x25, 0xfe, 0xd0, 0x56, 0xcb, 0x9e, 0x19, 0xac, 0x8c, 0xef, 0xc8, 0x0e, 0x66, 0x1f, 0x98, 0x46,
	0xd7, 0x27, 0xae, 0xf4, 0x0d, 0xe7, 0x7c, 0xdc, 0x58, 0x7c, 0x47, 0xde, 0xc1, 0xe2, 0xe2, 0xe9,
	0xb3, 0xdf, 0x30, 0x3e, 0xfb, 0xf7, 0xc2, 0x6e, 0x43, 0xf6, 0x80, 0x7f, 0x09, 0xaf, 0xff, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x26, 0x74, 0x14, 0x6e, 0x27, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderClient interface {
	CreateOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderCommonReply, error)
	GetOrderList(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderReply, error)
	GetOrderStatistic(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderStatisticReply, error)
}

type orderClient struct {
	cc *grpc.ClientConn
}

func NewOrderClient(cc *grpc.ClientConn) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) CreateOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderCommonReply, error) {
	out := new(OrderCommonReply)
	err := c.cc.Invoke(ctx, "/pb.Order/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) GetOrderList(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderReply, error) {
	out := new(OrderReply)
	err := c.cc.Invoke(ctx, "/pb.Order/GetOrderList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) GetOrderStatistic(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderStatisticReply, error) {
	out := new(OrderStatisticReply)
	err := c.cc.Invoke(ctx, "/pb.Order/GetOrderStatistic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
type OrderServer interface {
	CreateOrder(context.Context, *OrderRequest) (*OrderCommonReply, error)
	GetOrderList(context.Context, *OrderRequest) (*OrderReply, error)
	GetOrderStatistic(context.Context, *OrderRequest) (*OrderStatisticReply, error)
}

func RegisterOrderServer(s *grpc.Server, srv OrderServer) {
	s.RegisterService(&_Order_serviceDesc, srv)
}

func _Order_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Order/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateOrder(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_GetOrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetOrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Order/GetOrderList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetOrderList(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_GetOrderStatistic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetOrderStatistic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Order/GetOrderStatistic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetOrderStatistic(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Order_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _Order_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrderList",
			Handler:    _Order_GetOrderList_Handler,
		},
		{
			MethodName: "GetOrderStatistic",
			Handler:    _Order_GetOrderStatistic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/order_pb.proto",
}
