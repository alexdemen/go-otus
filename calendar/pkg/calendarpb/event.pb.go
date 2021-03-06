// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package calendarpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Event struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	StartDate            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	Duration             *duration.Duration   `protobuf:"bytes,5,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Event) GetStartDate() *timestamp.Timestamp {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *Event) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "Event")
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_2d17a9d3f0ddf27e) }

var fileDescriptor_2d17a9d3f0ddf27e = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8e, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xe5, 0xb4, 0x45, 0xf4, 0x82, 0x18, 0x3c, 0x99, 0x0e, 0x60, 0x31, 0x65, 0x72, 0x25,
	0x10, 0x03, 0x2b, 0x2a, 0x2f, 0x60, 0x31, 0xb1, 0xa0, 0x4b, 0x7d, 0x54, 0x96, 0x1a, 0xdb, 0x72,
	0xae, 0xbc, 0x1e, 0xaf, 0x86, 0xe2, 0x24, 0x08, 0xd1, 0xcd, 0xbe, 0xff, 0xfb, 0x7f, 0x7d, 0x50,
	0xd3, 0x17, 0x05, 0x36, 0x29, 0x47, 0x8e, 0x9b, 0xbb, 0x43, 0x8c, 0x87, 0x23, 0x6d, 0xcb, 0xaf,
	0x3d, 0x7d, 0x6e, 0xd9, 0x77, 0xd4, 0x33, 0x76, 0x69, 0x02, 0x6e, 0xff, 0x03, 0xee, 0x94, 0x91,
	0x7d, 0x0c, 0x63, 0x7e, 0xff, 0x2d, 0x60, 0xf5, 0x3a, 0x0c, 0xca, 0x6b, 0xa8, 0xbc, 0x53, 0x42,
	0x8b, 0x66, 0x61, 0x2b, 0xef, 0xa4, 0x84, 0x65, 0xc0, 0x8e, 0x54, 0xa5, 0x45, 0xb3, 0xb6, 0xe5,
	0x2d, 0x35, 0xd4, 0x8e, 0xfa, 0x7d, 0xf6, 0x69, 0x98, 0x50, 0x8b, 0x12, 0xfd, 0x3d, 0xc9, 0x67,
	0x80, 0x9e, 0x31, 0xf3, 0x87, 0x43, 0x26, 0xb5, 0xd4, 0xa2, 0xa9, 0x1f, 0x36, 0x66, 0x94, 0x30,
	0xb3, 0x84, 0x79, 0x9b, 0x2d, 0xed, 0xba, 0xd0, 0x3b, 0x64, 0x92, 0x4f, 0x70, 0x39, 0xcb, 0xa9,
	0x55, 0x29, 0xde, 0x9c, 0x15, 0x77, 0x13, 0x60, 0x7f, 0xd1, 0x97, 0xab, 0x77, 0xd8, 0xe3, 0x91,
	0x82, 0xc3, 0x9c, 0xda, 0xf6, 0xa2, 0xa0, 0x8f, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x72, 0x02,
	0x87, 0x3e, 0x26, 0x01, 0x00, 0x00,
}
