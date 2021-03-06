// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helpers.proto

package types // import "github.com/inn4science/exonum-go/types"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Hash struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hash) Reset()         { *m = Hash{} }
func (m *Hash) String() string { return proto.CompactTextString(m) }
func (*Hash) ProtoMessage()    {}
func (*Hash) Descriptor() ([]byte, []int) {
	return fileDescriptor_helpers_aedb05ef7e2aea53, []int{0}
}
func (m *Hash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hash.Unmarshal(m, b)
}
func (m *Hash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hash.Marshal(b, m, deterministic)
}
func (dst *Hash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hash.Merge(dst, src)
}
func (m *Hash) XXX_Size() int {
	return xxx_messageInfo_Hash.Size(m)
}
func (m *Hash) XXX_DiscardUnknown() {
	xxx_messageInfo_Hash.DiscardUnknown(m)
}

var xxx_messageInfo_Hash proto.InternalMessageInfo

func (m *Hash) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type PublicKey struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublicKey) Reset()         { *m = PublicKey{} }
func (m *PublicKey) String() string { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()    {}
func (*PublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_helpers_aedb05ef7e2aea53, []int{1}
}
func (m *PublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicKey.Unmarshal(m, b)
}
func (m *PublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicKey.Marshal(b, m, deterministic)
}
func (dst *PublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicKey.Merge(dst, src)
}
func (m *PublicKey) XXX_Size() int {
	return xxx_messageInfo_PublicKey.Size(m)
}
func (m *PublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_PublicKey proto.InternalMessageInfo

func (m *PublicKey) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type SecretKey struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecretKey) Reset()         { *m = SecretKey{} }
func (m *SecretKey) String() string { return proto.CompactTextString(m) }
func (*SecretKey) ProtoMessage()    {}
func (*SecretKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_helpers_aedb05ef7e2aea53, []int{2}
}
func (m *SecretKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecretKey.Unmarshal(m, b)
}
func (m *SecretKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecretKey.Marshal(b, m, deterministic)
}
func (dst *SecretKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecretKey.Merge(dst, src)
}
func (m *SecretKey) XXX_Size() int {
	return xxx_messageInfo_SecretKey.Size(m)
}
func (m *SecretKey) XXX_DiscardUnknown() {
	xxx_messageInfo_SecretKey.DiscardUnknown(m)
}

var xxx_messageInfo_SecretKey proto.InternalMessageInfo

func (m *SecretKey) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type BitVec struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Len                  uint64   `protobuf:"varint,2,opt,name=len,proto3" json:"len,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BitVec) Reset()         { *m = BitVec{} }
func (m *BitVec) String() string { return proto.CompactTextString(m) }
func (*BitVec) ProtoMessage()    {}
func (*BitVec) Descriptor() ([]byte, []int) {
	return fileDescriptor_helpers_aedb05ef7e2aea53, []int{3}
}
func (m *BitVec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BitVec.Unmarshal(m, b)
}
func (m *BitVec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BitVec.Marshal(b, m, deterministic)
}
func (dst *BitVec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BitVec.Merge(dst, src)
}
func (m *BitVec) XXX_Size() int {
	return xxx_messageInfo_BitVec.Size(m)
}
func (m *BitVec) XXX_DiscardUnknown() {
	xxx_messageInfo_BitVec.DiscardUnknown(m)
}

var xxx_messageInfo_BitVec proto.InternalMessageInfo

func (m *BitVec) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *BitVec) GetLen() uint64 {
	if m != nil {
		return m.Len
	}
	return 0
}

func init() {
	proto.RegisterType((*Hash)(nil), "exonum.Hash")
	proto.RegisterType((*PublicKey)(nil), "exonum.PublicKey")
	proto.RegisterType((*SecretKey)(nil), "exonum.SecretKey")
	proto.RegisterType((*BitVec)(nil), "exonum.BitVec")
}

func init() { proto.RegisterFile("helpers.proto", fileDescriptor_helpers_aedb05ef7e2aea53) }

var fileDescriptor_helpers_aedb05ef7e2aea53 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x48, 0xcd, 0x29,
	0x48, 0x2d, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xad, 0xc8, 0xcf, 0x2b,
	0xcd, 0x55, 0x92, 0xe2, 0x62, 0xf1, 0x48, 0x2c, 0xce, 0x10, 0x12, 0xe2, 0x62, 0x49, 0x49, 0x2c,
	0x49, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x02, 0xb3, 0x95, 0xe4, 0xb9, 0x38, 0x03, 0x4a,
	0x93, 0x72, 0x32, 0x93, 0xbd, 0x53, 0x2b, 0x71, 0x29, 0x08, 0x4e, 0x4d, 0x2e, 0x4a, 0x2d, 0xc1,
	0xa5, 0x40, 0x8f, 0x8b, 0xcd, 0x29, 0xb3, 0x24, 0x2c, 0x35, 0x19, 0x9b, 0xac, 0x90, 0x00, 0x17,
	0x73, 0x4e, 0x6a, 0x9e, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0x88, 0xe9, 0xa4, 0x11, 0xa5,
	0x96, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x99, 0x97, 0x67, 0x52,
	0x9c, 0x9c, 0x99, 0x9a, 0x97, 0x9c, 0xaa, 0x0f, 0x71, 0xae, 0x6e, 0x7a, 0xbe, 0x7e, 0x49, 0x65,
	0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x1b, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x27,
	0xb9, 0xa8, 0xd7, 0x00, 0x00, 0x00,
}
