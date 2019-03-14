// Code generated by protoc-gen-go. DO NOT EDIT.
// source: consensus/synchro/pb/message.proto

package msgpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type RequireType int32

const (
	RequireType_GETBLOCKHASHES         RequireType = 0
	RequireType_GETBLOCKHASHESBYNUMBER RequireType = 1
)

var RequireType_name = map[int32]string{
	0: "GETBLOCKHASHES",
	1: "GETBLOCKHASHESBYNUMBER",
}

var RequireType_value = map[string]int32{
	"GETBLOCKHASHES":         0,
	"GETBLOCKHASHESBYNUMBER": 1,
}

func (x RequireType) String() string {
	return proto.EnumName(RequireType_name, int32(x))
}

func (RequireType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b8c018fb18032427, []int{0}
}

type BlockInfo struct {
	Number               int64    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Hash                 []byte   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockInfo) Reset()         { *m = BlockInfo{} }
func (m *BlockInfo) String() string { return proto.CompactTextString(m) }
func (*BlockInfo) ProtoMessage()    {}
func (*BlockInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8c018fb18032427, []int{0}
}

func (m *BlockInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockInfo.Unmarshal(m, b)
}
func (m *BlockInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockInfo.Marshal(b, m, deterministic)
}
func (m *BlockInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockInfo.Merge(m, src)
}
func (m *BlockInfo) XXX_Size() int {
	return xxx_messageInfo_BlockInfo.Size(m)
}
func (m *BlockInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BlockInfo proto.InternalMessageInfo

func (m *BlockInfo) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *BlockInfo) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type BlockHashQuery struct {
	ReqType              RequireType `protobuf:"varint,1,opt,name=reqType,proto3,enum=msgpb.RequireType" json:"reqType,omitempty"`
	Start                int64       `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	End                  int64       `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
	Nums                 []int64     `protobuf:"varint,4,rep,packed,name=nums,proto3" json:"nums,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BlockHashQuery) Reset()         { *m = BlockHashQuery{} }
func (m *BlockHashQuery) String() string { return proto.CompactTextString(m) }
func (*BlockHashQuery) ProtoMessage()    {}
func (*BlockHashQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8c018fb18032427, []int{1}
}

func (m *BlockHashQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHashQuery.Unmarshal(m, b)
}
func (m *BlockHashQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHashQuery.Marshal(b, m, deterministic)
}
func (m *BlockHashQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHashQuery.Merge(m, src)
}
func (m *BlockHashQuery) XXX_Size() int {
	return xxx_messageInfo_BlockHashQuery.Size(m)
}
func (m *BlockHashQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHashQuery.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHashQuery proto.InternalMessageInfo

func (m *BlockHashQuery) GetReqType() RequireType {
	if m != nil {
		return m.ReqType
	}
	return RequireType_GETBLOCKHASHES
}

func (m *BlockHashQuery) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *BlockHashQuery) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *BlockHashQuery) GetNums() []int64 {
	if m != nil {
		return m.Nums
	}
	return nil
}

type BlockHashResponse struct {
	BlockInfos           []*BlockInfo `protobuf:"bytes,1,rep,name=blockInfos,proto3" json:"blockInfos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BlockHashResponse) Reset()         { *m = BlockHashResponse{} }
func (m *BlockHashResponse) String() string { return proto.CompactTextString(m) }
func (*BlockHashResponse) ProtoMessage()    {}
func (*BlockHashResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8c018fb18032427, []int{2}
}

func (m *BlockHashResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHashResponse.Unmarshal(m, b)
}
func (m *BlockHashResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHashResponse.Marshal(b, m, deterministic)
}
func (m *BlockHashResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHashResponse.Merge(m, src)
}
func (m *BlockHashResponse) XXX_Size() int {
	return xxx_messageInfo_BlockHashResponse.Size(m)
}
func (m *BlockHashResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHashResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHashResponse proto.InternalMessageInfo

func (m *BlockHashResponse) GetBlockInfos() []*BlockInfo {
	if m != nil {
		return m.BlockInfos
	}
	return nil
}

type SyncHeight struct {
	Height               int64    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Time                 int64    `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncHeight) Reset()         { *m = SyncHeight{} }
func (m *SyncHeight) String() string { return proto.CompactTextString(m) }
func (*SyncHeight) ProtoMessage()    {}
func (*SyncHeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8c018fb18032427, []int{3}
}

func (m *SyncHeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncHeight.Unmarshal(m, b)
}
func (m *SyncHeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncHeight.Marshal(b, m, deterministic)
}
func (m *SyncHeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncHeight.Merge(m, src)
}
func (m *SyncHeight) XXX_Size() int {
	return xxx_messageInfo_SyncHeight.Size(m)
}
func (m *SyncHeight) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncHeight.DiscardUnknown(m)
}

var xxx_messageInfo_SyncHeight proto.InternalMessageInfo

func (m *SyncHeight) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *SyncHeight) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func init() {
	proto.RegisterEnum("msgpb.RequireType", RequireType_name, RequireType_value)
	proto.RegisterType((*BlockInfo)(nil), "msgpb.BlockInfo")
	proto.RegisterType((*BlockHashQuery)(nil), "msgpb.BlockHashQuery")
	proto.RegisterType((*BlockHashResponse)(nil), "msgpb.BlockHashResponse")
	proto.RegisterType((*SyncHeight)(nil), "msgpb.SyncHeight")
}

func init() { proto.RegisterFile("consensus/synchro/pb/message.proto", fileDescriptor_b8c018fb18032427) }

var fileDescriptor_b8c018fb18032427 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xcd, 0x4f, 0x83, 0x40,
	0x10, 0xc5, 0xc5, 0x6d, 0x6b, 0x9c, 0x9a, 0x06, 0x27, 0xa6, 0x21, 0x9e, 0x08, 0x27, 0x62, 0x4c,
	0x6b, 0xea, 0x41, 0x2f, 0x1e, 0xc4, 0x10, 0x31, 0x7e, 0xc5, 0x6d, 0x3d, 0x78, 0x04, 0x1c, 0x4b,
	0xa3, 0x2c, 0x74, 0x07, 0x0e, 0xf8, 0xd7, 0x1b, 0xb6, 0x1f, 0xa9, 0xb7, 0xf7, 0x86, 0x19, 0xde,
	0xfb, 0x65, 0xc1, 0x4b, 0x0b, 0xc5, 0xa4, 0xb8, 0xe6, 0x31, 0x37, 0x2a, 0xcd, 0x74, 0x31, 0x2e,
	0x93, 0x71, 0x4e, 0xcc, 0xf1, 0x9c, 0x46, 0xa5, 0x2e, 0xaa, 0x02, 0xbb, 0x39, 0xcf, 0xcb, 0xc4,
	0xbb, 0x82, 0xc3, 0xe0, 0xa7, 0x48, 0xbf, 0x1f, 0xd4, 0x57, 0x81, 0x43, 0xe8, 0xa9, 0x3a, 0x4f,
	0x48, 0x3b, 0x96, 0x6b, 0xf9, 0x42, 0xae, 0x1d, 0x22, 0x74, 0xb2, 0x98, 0x33, 0x67, 0xdf, 0xb5,
	0xfc, 0x23, 0x69, 0xb4, 0xf7, 0x0b, 0x03, 0x73, 0x18, 0xc5, 0x9c, 0xbd, 0xd5, 0xa4, 0x1b, 0x3c,
	0x87, 0x03, 0x4d, 0xcb, 0x59, 0x53, 0x92, 0x39, 0x1f, 0x4c, 0x70, 0x64, 0x32, 0x46, 0x92, 0x96,
	0xf5, 0x42, 0x53, 0xfb, 0x45, 0x6e, 0x56, 0xf0, 0x04, 0xba, 0x5c, 0xc5, 0xba, 0x32, 0x3f, 0x15,
	0x72, 0x65, 0xd0, 0x06, 0x41, 0xea, 0xd3, 0x11, 0x66, 0xd6, 0xca, 0x36, 0x5b, 0xd5, 0x39, 0x3b,
	0x1d, 0x57, 0xf8, 0x42, 0x1a, 0xed, 0x85, 0x70, 0xbc, 0xcd, 0x96, 0xc4, 0x65, 0x4b, 0x8b, 0x17,
	0x00, 0xc9, 0x86, 0x84, 0x1d, 0xcb, 0x15, 0x7e, 0x7f, 0x62, 0xaf, 0x1b, 0x6c, 0x11, 0xe5, 0xce,
	0x8e, 0x77, 0x0d, 0x30, 0x6d, 0x54, 0x1a, 0xd1, 0x62, 0x9e, 0x55, 0x2d, 0x7c, 0x66, 0xd4, 0x06,
	0x7e, 0xe5, 0xda, 0x02, 0xd5, 0x22, 0xa7, 0x75, 0x4f, 0xa3, 0xcf, 0x6e, 0xa0, 0xbf, 0x03, 0x85,
	0x08, 0x83, 0xfb, 0x70, 0x16, 0x3c, 0xbd, 0xde, 0x3d, 0x46, 0xb7, 0xd3, 0x28, 0x9c, 0xda, 0x7b,
	0x78, 0x0a, 0xc3, 0xff, 0xb3, 0xe0, 0xe3, 0xe5, 0xfd, 0x39, 0x08, 0xa5, 0x6d, 0x25, 0x3d, 0xf3,
	0x04, 0x97, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x56, 0x59, 0xbb, 0xa8, 0x01, 0x00, 0x00,
}
