// Code generated by protoc-gen-go. DO NOT EDIT.
// source: screenmessage.proto

package __protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Screen_Panel int32

const (
	Screen_UNKNOWN Screen_Panel = 0
	Screen_IPS     Screen_Panel = 1
	Screen_OLED    Screen_Panel = 2
)

var Screen_Panel_name = map[int32]string{
	0: "UNKNOWN",
	1: "IPS",
	2: "OLED",
}
var Screen_Panel_value = map[string]int32{
	"UNKNOWN": 0,
	"IPS":     1,
	"OLED":    2,
}

func (x Screen_Panel) String() string {
	return proto.EnumName(Screen_Panel_name, int32(x))
}
func (Screen_Panel) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

type Screen struct {
	SizeInch   float32            `protobuf:"fixed32,1,opt,name=size_inch,json=sizeInch" json:"size_inch,omitempty"`
	Resolution *Screen_Resolution `protobuf:"bytes,2,opt,name=resolution" json:"resolution,omitempty"`
	Panel      Screen_Panel       `protobuf:"varint,3,opt,name=panel,enum=paulomujuru.computerbook.Screen_Panel" json:"panel,omitempty"`
	Multitouch bool               `protobuf:"varint,4,opt,name=multitouch" json:"multitouch,omitempty"`
}

func (m *Screen) Reset()                    { *m = Screen{} }
func (m *Screen) String() string            { return proto.CompactTextString(m) }
func (*Screen) ProtoMessage()               {}
func (*Screen) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Screen) GetSizeInch() float32 {
	if m != nil {
		return m.SizeInch
	}
	return 0
}

func (m *Screen) GetResolution() *Screen_Resolution {
	if m != nil {
		return m.Resolution
	}
	return nil
}

func (m *Screen) GetPanel() Screen_Panel {
	if m != nil {
		return m.Panel
	}
	return Screen_UNKNOWN
}

func (m *Screen) GetMultitouch() bool {
	if m != nil {
		return m.Multitouch
	}
	return false
}

type Screen_Resolution struct {
	Width  uint32 `protobuf:"varint,1,opt,name=width" json:"width,omitempty"`
	Height uint32 `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
}

func (m *Screen_Resolution) Reset()                    { *m = Screen_Resolution{} }
func (m *Screen_Resolution) String() string            { return proto.CompactTextString(m) }
func (*Screen_Resolution) ProtoMessage()               {}
func (*Screen_Resolution) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

func (m *Screen_Resolution) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Screen_Resolution) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterType((*Screen)(nil), "paulomujuru.computerbook.Screen")
	proto.RegisterType((*Screen_Resolution)(nil), "paulomujuru.computerbook.Screen.Resolution")
	proto.RegisterEnum("paulomujuru.computerbook.Screen_Panel", Screen_Panel_name, Screen_Panel_value)
}

func init() { proto.RegisterFile("screenmessage.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x18, 0xc5, 0x4d, 0xb7, 0x76, 0xf5, 0x9b, 0x93, 0xf2, 0x29, 0x52, 0x14, 0xa4, 0xec, 0xa0, 0x05,
	0x21, 0x87, 0x79, 0x53, 0x4f, 0xa2, 0x87, 0x31, 0xe9, 0x46, 0x86, 0x08, 0x5e, 0xa4, 0xad, 0x71,
	0x8d, 0xb6, 0x4d, 0x69, 0x12, 0x04, 0xff, 0x2f, 0xff, 0x3f, 0x59, 0x2a, 0xba, 0x8b, 0x78, 0xcb,
	0x7b, 0xe4, 0xf7, 0x1e, 0xef, 0x83, 0x3d, 0x95, 0xb7, 0x9c, 0xd7, 0x15, 0x57, 0x2a, 0x5d, 0x71,
	0xda, 0xb4, 0x52, 0x4b, 0x0c, 0x9b, 0xd4, 0x94, 0xb2, 0x32, 0xaf, 0xa6, 0x35, 0x34, 0x97, 0x55,
	0x63, 0x34, 0x6f, 0x33, 0x29, 0xdf, 0xc6, 0x9f, 0x0e, 0x78, 0x4b, 0x4b, 0xe0, 0x11, 0x6c, 0x2b,
	0xf1, 0xc1, 0x9f, 0x44, 0x9d, 0x17, 0x21, 0x89, 0x48, 0xec, 0x30, 0x7f, 0x6d, 0x4c, 0xeb, 0xbc,
	0xc0, 0x19, 0x40, 0xcb, 0x95, 0x2c, 0x8d, 0x16, 0xb2, 0x0e, 0x9d, 0x88, 0xc4, 0xc3, 0xc9, 0x19,
	0xfd, 0x2b, 0x96, 0x76, 0x91, 0x94, 0xfd, 0x20, 0x6c, 0x03, 0xc7, 0x2b, 0x70, 0x9b, 0xb4, 0xe6,
	0x65, 0xd8, 0x8b, 0x48, 0xbc, 0x3b, 0x39, 0xf9, 0x37, 0x67, 0xb1, 0xfe, 0xcd, 0x3a, 0x08, 0x8f,
	0x01, 0x2a, 0x53, 0x6a, 0xa1, 0xa5, 0xc9, 0x8b, 0xb0, 0x1f, 0x91, 0xd8, 0x67, 0x1b, 0xce, 0xe1,
	0x05, 0xc0, 0x6f, 0x2f, 0xee, 0x83, 0xfb, 0x2e, 0x9e, 0x75, 0xb7, 0x68, 0xc4, 0x3a, 0x81, 0x07,
	0xe0, 0x15, 0x5c, 0xac, 0x0a, 0x6d, 0xa7, 0x8c, 0xd8, 0xb7, 0x1a, 0x9f, 0x82, 0x6b, 0xbb, 0x70,
	0x08, 0x83, 0xfb, 0x64, 0x96, 0xcc, 0x1f, 0x92, 0x60, 0x0b, 0x07, 0xd0, 0x9b, 0x2e, 0x96, 0x01,
	0x41, 0x1f, 0xfa, 0xf3, 0xbb, 0xdb, 0x9b, 0xc0, 0xb9, 0xde, 0x79, 0x04, 0x7a, 0x69, 0x8f, 0x9b,
	0x99, 0x97, 0xcc, 0xb3, 0xaf, 0xf3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x82, 0x57, 0x81, 0xb4,
	0x7d, 0x01, 0x00, 0x00,
}