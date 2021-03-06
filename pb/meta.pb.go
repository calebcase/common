// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: meta.proto

package pb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// SerializableMeta is the object metadata that will be stored serialized
type SerializableMeta struct {
	UserDefined          map[string]string `protobuf:"bytes,2,rep,name=user_defined,json=userDefined,proto3" json:"user_defined,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ContentType          string            `protobuf:"bytes,1,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	ContentLength        int64             `protobuf:"varint,3,opt,name=content_length,json=contentLength,proto3" json:"content_length,omitempty"`
	FileCreated          *time.Time        `protobuf:"bytes,4,opt,name=file_created,json=fileCreated,proto3,stdtime" json:"file_created,omitempty"`
	FileModified         *time.Time        `protobuf:"bytes,5,opt,name=file_modified,json=fileModified,proto3,stdtime" json:"file_modified,omitempty"`
	FilePermissions      uint32            `protobuf:"varint,6,opt,name=file_permissions,json=filePermissions,proto3" json:"file_permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SerializableMeta) Reset()         { *m = SerializableMeta{} }
func (m *SerializableMeta) String() string { return proto.CompactTextString(m) }
func (*SerializableMeta) ProtoMessage()    {}
func (*SerializableMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5ea8fe65782bcc, []int{0}
}
func (m *SerializableMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SerializableMeta.Unmarshal(m, b)
}
func (m *SerializableMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SerializableMeta.Marshal(b, m, deterministic)
}
func (m *SerializableMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SerializableMeta.Merge(m, src)
}
func (m *SerializableMeta) XXX_Size() int {
	return xxx_messageInfo_SerializableMeta.Size(m)
}
func (m *SerializableMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_SerializableMeta.DiscardUnknown(m)
}

var xxx_messageInfo_SerializableMeta proto.InternalMessageInfo

func (m *SerializableMeta) GetUserDefined() map[string]string {
	if m != nil {
		return m.UserDefined
	}
	return nil
}

func (m *SerializableMeta) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *SerializableMeta) GetContentLength() int64 {
	if m != nil {
		return m.ContentLength
	}
	return 0
}

func (m *SerializableMeta) GetFileCreated() *time.Time {
	if m != nil {
		return m.FileCreated
	}
	return nil
}

func (m *SerializableMeta) GetFileModified() *time.Time {
	if m != nil {
		return m.FileModified
	}
	return nil
}

func (m *SerializableMeta) GetFilePermissions() uint32 {
	if m != nil {
		return m.FilePermissions
	}
	return 0
}

func init() {
	proto.RegisterType((*SerializableMeta)(nil), "objects.SerializableMeta")
	proto.RegisterMapType((map[string]string)(nil), "objects.SerializableMeta.UserDefinedEntry")
}

func init() { proto.RegisterFile("meta.proto", fileDescriptor_3b5ea8fe65782bcc) }

var fileDescriptor_3b5ea8fe65782bcc = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0x69, 0xbb, 0x4d, 0x4c, 0x37, 0x2d, 0xc1, 0x43, 0xe8, 0x65, 0x55, 0x10, 0xaa, 0x87,
	0x0c, 0xe6, 0x45, 0x3c, 0x78, 0xd8, 0xdc, 0xcd, 0x81, 0xd4, 0x79, 0xf1, 0x32, 0xd2, 0xf5, 0x6b,
	0x8d, 0xb6, 0x4d, 0x49, 0x52, 0xa1, 0xfe, 0x0a, 0x7f, 0x83, 0x7f, 0xcc, 0xbf, 0x22, 0x4b, 0x3b,
	0x06, 0x3b, 0x79, 0xcb, 0xf7, 0xf0, 0xbe, 0xef, 0x97, 0xf7, 0x43, 0xa8, 0x00, 0xcd, 0x68, 0x25,
	0x85, 0x16, 0xf8, 0x48, 0xc4, 0xef, 0xb0, 0xd1, 0xca, 0x47, 0x99, 0xc8, 0x44, 0x0b, 0xfd, 0x71,
	0x26, 0x44, 0x96, 0xc3, 0xc4, 0x4c, 0x71, 0x9d, 0x4e, 0x34, 0x2f, 0x40, 0x69, 0x56, 0x54, 0xad,
	0xe0, 0xe2, 0xc7, 0x41, 0xde, 0x33, 0x48, 0xce, 0x72, 0xfe, 0xc5, 0xe2, 0x1c, 0x96, 0xa0, 0x19,
	0x5e, 0xa2, 0x61, 0xad, 0x40, 0xae, 0x13, 0x48, 0x79, 0x09, 0x09, 0xb1, 0x03, 0x27, 0x74, 0xa7,
	0xd7, 0xb4, 0xdb, 0x40, 0x0f, 0x0d, 0xf4, 0x45, 0x81, 0x7c, 0x68, 0xc5, 0x8b, 0x52, 0xcb, 0x26,
	0x72, 0xeb, 0x3d, 0xc1, 0xe7, 0x68, 0xb8, 0x11, 0xa5, 0x86, 0x52, 0xaf, 0x75, 0x53, 0x01, 0xb1,
	0x02, 0x2b, 0x3c, 0x8e, 0xdc, 0x8e, 0xad, 0x9a, 0x0a, 0xf0, 0x25, 0x3a, 0xd9, 0x49, 0x72, 0x28,
	0x33, 0xfd, 0x46, 0x9c, 0xc0, 0x0a, 0x9d, 0x68, 0xd4, 0xd1, 0x47, 0x03, 0xf1, 0x1c, 0x0d, 0x53,
	0x9e, 0xc3, 0x7a, 0x23, 0x81, 0x69, 0x48, 0x48, 0x2f, 0xb0, 0x42, 0x77, 0xea, 0xd3, 0xb6, 0x25,
	0xdd, 0xb5, 0xa4, 0xab, 0x5d, 0xcb, 0x59, 0xef, 0xfb, 0x77, 0x6c, 0x45, 0xee, 0xd6, 0x35, 0x6f,
	0x4d, 0x78, 0x81, 0x46, 0x26, 0xa4, 0x10, 0x09, 0x4f, 0x39, 0x24, 0xa4, 0xff, 0xcf, 0x14, 0xb3,
	0x7b, 0xd9, 0xb9, 0xf0, 0x15, 0xf2, 0x4c, 0x4c, 0x05, 0xb2, 0xe0, 0x4a, 0x71, 0x51, 0x2a, 0x32,
	0x08, 0xac, 0x70, 0x14, 0x9d, 0x6e, 0xf9, 0xd3, 0x1e, 0xfb, 0xf7, 0xc8, 0x3b, 0xbc, 0x10, 0xf6,
	0x90, 0xf3, 0x01, 0x4d, 0x77, 0x8b, 0xed, 0x13, 0x9f, 0xa1, 0xfe, 0x27, 0xcb, 0x6b, 0x20, 0xb6,
	0x61, 0xed, 0x70, 0x67, 0xdf, 0x5a, 0xb3, 0xde, 0xab, 0x5d, 0xc5, 0xf1, 0xc0, 0x7c, 0xec, 0xe6,
	0x2f, 0x00, 0x00, 0xff, 0xff, 0x86, 0x86, 0xfa, 0xf2, 0xf5, 0x01, 0x00, 0x00,
}
