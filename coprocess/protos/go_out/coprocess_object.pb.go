// Code generated by protoc-gen-go.
// source: coprocess_object.proto
// DO NOT EDIT!

package coprocess

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Object struct {
	HookType HookType           `protobuf:"varint,1,opt,name=hook_type,json=hookType,enum=coprocess.HookType" json:"hook_type,omitempty"`
	Request  *MiniRequestObject `protobuf:"bytes,2,opt,name=request" json:"request,omitempty"`
	Session  *SessionState      `protobuf:"bytes,3,opt,name=session" json:"session,omitempty"`
	Metadata map[string]string  `protobuf:"bytes,4,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Spec     map[string]string  `protobuf:"bytes,5,rep,name=spec" json:"spec,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Object) Reset()                    { *m = Object{} }
func (m *Object) String() string            { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()               {}
func (*Object) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Object) GetRequest() *MiniRequestObject {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Object) GetSession() *SessionState {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *Object) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Object) GetSpec() map[string]string {
	if m != nil {
		return m.Spec
	}
	return nil
}

func init() {
	proto.RegisterType((*Object)(nil), "coprocess.Object")
}

func init() { proto.RegisterFile("coprocess_object.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x91, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xe9, 0xba, 0x5f, 0x7d, 0xa2, 0x48, 0x14, 0x2d, 0x55, 0x71, 0xe8, 0xc5, 0x53, 0xd5,
	0x09, 0x2a, 0xee, 0x2c, 0x78, 0x19, 0x42, 0xea, 0xbd, 0x64, 0xf1, 0xc1, 0xea, 0x6c, 0x53, 0x9b,
	0x4c, 0xe8, 0x9f, 0xe0, 0x7f, 0x6d, 0x96, 0x64, 0x5d, 0x87, 0x27, 0x2f, 0x25, 0x79, 0xdf, 0xcf,
	0xe7, 0xbd, 0xd7, 0x16, 0x8e, 0xb8, 0x28, 0x2b, 0xc1, 0x51, 0xca, 0x54, 0xcc, 0x3e, 0x90, 0xab,
	0x58, 0x5f, 0x95, 0x20, 0x41, 0x53, 0x8f, 0x2e, 0x37, 0x48, 0x9e, 0x15, 0x59, 0x5a, 0xe1, 0xd7,
	0x12, 0xa5, 0xda, 0xe2, 0xa3, 0xb3, 0x0d, 0x24, 0xf5, 0x23, 0x13, 0x45, 0x2a, 0x15, 0x53, 0xe8,
	0xe2, 0xd6, 0x18, 0x2e, 0xf2, 0x5c, 0x14, 0xb6, 0x7e, 0xf1, 0xe3, 0x43, 0xff, 0xd5, 0xf4, 0x21,
	0x37, 0x10, 0xcc, 0x85, 0x58, 0xa4, 0xaa, 0x2e, 0x31, 0xf4, 0x46, 0xde, 0xd5, 0xde, 0xf8, 0x20,
	0x6e, 0xb4, 0xf8, 0x45, 0x67, 0x6f, 0x3a, 0xa2, 0xc3, 0xb9, 0x3b, 0x91, 0x7b, 0x18, 0xb8, 0x5d,
	0xc2, 0x8e, 0xe6, 0x77, 0xc6, 0xa7, 0x2d, 0x7e, 0xaa, 0x57, 0xa5, 0x36, 0xb5, 0x03, 0xe8, 0x1a,
	0x26, 0xb7, 0x30, 0x70, 0x3b, 0x86, 0xbe, 0xf1, 0x8e, 0x5b, 0x5e, 0x62, 0x93, 0x64, 0xb5, 0x3c,
	0x5d, 0x73, 0x64, 0x02, 0xc3, 0x1c, 0x15, 0x7b, 0x67, 0x8a, 0x85, 0xdd, 0x91, 0xaf, 0x9d, 0xf3,
	0x96, 0x63, 0x07, 0xc4, 0x53, 0x47, 0x3c, 0x17, 0xaa, 0xaa, 0x69, 0x23, 0x90, 0x6b, 0xe8, 0xca,
	0x12, 0x79, 0xd8, 0x33, 0xe2, 0xc9, 0x5f, 0x31, 0xd1, 0xa9, 0x95, 0x0c, 0x18, 0x4d, 0x60, 0x77,
	0xab, 0x17, 0xd9, 0x07, 0x7f, 0x81, 0xb5, 0xf9, 0x2a, 0x01, 0x5d, 0x1d, 0xc9, 0x21, 0xf4, 0xbe,
	0xd9, 0xe7, 0x12, 0xcd, 0x9b, 0x07, 0xd4, 0x5e, 0x9e, 0x3a, 0x8f, 0x5e, 0xf4, 0x00, 0x41, 0xd3,
	0xef, 0x3f, 0xe2, 0xac, 0x6f, 0x7e, 0xc9, 0xdd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x86,
	0x7e, 0x67, 0x13, 0x02, 0x00, 0x00,
}
