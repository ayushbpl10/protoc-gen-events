// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/proto/eventspush/push.proto

package eventspush // import "pb/eventspush"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MyEvents struct {
	Push                 bool     `protobuf:"varint,1,opt,name=push,proto3" json:"push,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MyEvents) Reset()         { *m = MyEvents{} }
func (m *MyEvents) String() string { return proto.CompactTextString(m) }
func (*MyEvents) ProtoMessage()    {}
func (*MyEvents) Descriptor() ([]byte, []int) {
	return fileDescriptor_push_669dbfebbcb7c53f, []int{0}
}
func (m *MyEvents) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyEvents.Unmarshal(m, b)
}
func (m *MyEvents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyEvents.Marshal(b, m, deterministic)
}
func (dst *MyEvents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyEvents.Merge(dst, src)
}
func (m *MyEvents) XXX_Size() int {
	return xxx_messageInfo_MyEvents.Size(m)
}
func (m *MyEvents) XXX_DiscardUnknown() {
	xxx_messageInfo_MyEvents.DiscardUnknown(m)
}

var xxx_messageInfo_MyEvents proto.InternalMessageInfo

func (m *MyEvents) GetPush() bool {
	if m != nil {
		return m.Push
	}
	return false
}

var E_Event = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*MyEvents)(nil),
	Field:         9966,
	Name:          "eventspush.event",
	Tag:           "bytes,9966,opt,name=event",
	Filename:      "example/proto/eventspush/push.proto",
}

func init() {
	proto.RegisterType((*MyEvents)(nil), "eventspush.MyEvents")
	proto.RegisterExtension(E_Event)
}

func init() {
	proto.RegisterFile("example/proto/eventspush/push.proto", fileDescriptor_push_669dbfebbcb7c53f)
}

var fileDescriptor_push_669dbfebbcb7c53f = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0x29,
	0x2e, 0x28, 0x2d, 0xce, 0xd0, 0x07, 0x11, 0x7a, 0x60, 0x51, 0x21, 0x2e, 0x84, 0xb0, 0x94, 0x42,
	0x7a, 0x7e, 0x7e, 0x3a, 0x4c, 0x7d, 0x52, 0x69, 0x9a, 0x7e, 0x4a, 0x6a, 0x71, 0x72, 0x51, 0x66,
	0x41, 0x49, 0x7e, 0x11, 0x44, 0xb5, 0x92, 0x1c, 0x17, 0x87, 0x6f, 0xa5, 0x2b, 0x58, 0x87, 0x90,
	0x10, 0x17, 0x0b, 0x48, 0x97, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x47, 0x10, 0x98, 0x6d, 0xe5, 0xcd,
	0xc5, 0x0a, 0x36, 0x4f, 0x48, 0x4e, 0x0f, 0x62, 0x96, 0x1e, 0xcc, 0x2c, 0x3d, 0xdf, 0xd4, 0x92,
	0x8c, 0xfc, 0x14, 0xff, 0x82, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x89, 0x77, 0xbe, 0x0a, 0x8c, 0x1a,
	0xdc, 0x46, 0x22, 0x7a, 0x08, 0xeb, 0xf5, 0x60, 0x26, 0x07, 0x41, 0xcc, 0x70, 0xe2, 0x8f, 0xe2,
	0x2d, 0x48, 0x42, 0x72, 0x76, 0x12, 0x1b, 0xd8, 0x30, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xc2, 0x1d, 0x60, 0xb3, 0xd9, 0x00, 0x00, 0x00,
}