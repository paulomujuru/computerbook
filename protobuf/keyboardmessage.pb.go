// Code generated by protoc-gen-go. DO NOT EDIT.
// source: keyboardmessage.proto

/*
Package __protobuf is a generated protocol buffer package.

It is generated from these files:
	keyboardmessage.proto
	laptopmessage.proto
	memorymessage.proto
	processormessage.proto
	screenmessage.proto
	storagemessage.proto

It has these top-level messages:
	Keyboard
	Laptop
	Memory
	CPU
	GPU
	Screen
	Storage
*/
package __protobuf

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

type Keyboard_Layout int32

const (
	Keyboard_UNKNOWN Keyboard_Layout = 0
	Keyboard_QWERTY  Keyboard_Layout = 1
	Keyboard_QWERTZ  Keyboard_Layout = 2
	Keyboard_AZERTY  Keyboard_Layout = 3
)

var Keyboard_Layout_name = map[int32]string{
	0: "UNKNOWN",
	1: "QWERTY",
	2: "QWERTZ",
	3: "AZERTY",
}
var Keyboard_Layout_value = map[string]int32{
	"UNKNOWN": 0,
	"QWERTY":  1,
	"QWERTZ":  2,
	"AZERTY":  3,
}

func (x Keyboard_Layout) String() string {
	return proto.EnumName(Keyboard_Layout_name, int32(x))
}
func (Keyboard_Layout) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Keyboard struct {
	Layout  Keyboard_Layout `protobuf:"varint,1,opt,name=layout,enum=paulomujuru.computerbook.Keyboard_Layout" json:"layout,omitempty"`
	Backlit bool            `protobuf:"varint,2,opt,name=backlit" json:"backlit,omitempty"`
}

func (m *Keyboard) Reset()                    { *m = Keyboard{} }
func (m *Keyboard) String() string            { return proto.CompactTextString(m) }
func (*Keyboard) ProtoMessage()               {}
func (*Keyboard) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Keyboard) GetLayout() Keyboard_Layout {
	if m != nil {
		return m.Layout
	}
	return Keyboard_UNKNOWN
}

func (m *Keyboard) GetBacklit() bool {
	if m != nil {
		return m.Backlit
	}
	return false
}

func init() {
	proto.RegisterType((*Keyboard)(nil), "paulomujuru.computerbook.Keyboard")
	proto.RegisterEnum("paulomujuru.computerbook.Keyboard_Layout", Keyboard_Layout_name, Keyboard_Layout_value)
}

func init() { proto.RegisterFile("keyboardmessage.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0x4e, 0xad, 0x4c,
	0xca, 0x4f, 0x2c, 0x4a, 0xc9, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0x28, 0x48, 0x2c, 0xcd, 0xc9, 0xcf, 0x2d, 0xcd, 0x2a, 0x2d, 0x2a, 0xd5, 0x4b,
	0xce, 0xcf, 0x2d, 0x28, 0x2d, 0x49, 0x2d, 0x4a, 0xca, 0xcf, 0xcf, 0x56, 0x5a, 0xc4, 0xc8, 0xc5,
	0xe1, 0x0d, 0xd5, 0x23, 0xe4, 0xc8, 0xc5, 0x96, 0x93, 0x58, 0x99, 0x5f, 0x5a, 0x22, 0xc1, 0xa8,
	0xc0, 0xa8, 0xc1, 0x67, 0xa4, 0xa9, 0x87, 0x4b, 0x9f, 0x1e, 0x4c, 0x8f, 0x9e, 0x0f, 0x58, 0x43,
	0x10, 0x54, 0xa3, 0x90, 0x04, 0x17, 0x7b, 0x52, 0x62, 0x72, 0x76, 0x4e, 0x66, 0x89, 0x04, 0x93,
	0x02, 0xa3, 0x06, 0x47, 0x10, 0x8c, 0xab, 0x64, 0xc9, 0xc5, 0x06, 0x51, 0x2b, 0xc4, 0xcd, 0xc5,
	0x1e, 0xea, 0xe7, 0xed, 0xe7, 0x1f, 0xee, 0x27, 0xc0, 0x20, 0xc4, 0xc5, 0xc5, 0x16, 0x18, 0xee,
	0x1a, 0x14, 0x12, 0x29, 0xc0, 0x08, 0x67, 0x47, 0x09, 0x30, 0x81, 0xd8, 0x8e, 0x51, 0x60, 0x71,
	0x66, 0x27, 0x9e, 0x28, 0x2e, 0x3d, 0x6b, 0xb0, 0x4f, 0x92, 0x4a, 0xd3, 0x92, 0xd8, 0xc0, 0x2c,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x34, 0x84, 0xb4, 0x66, 0xec, 0x00, 0x00, 0x00,
}
