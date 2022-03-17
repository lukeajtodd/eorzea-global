// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aetheryte-service-consignment/proto/consignment/consignment.proto

package consignment

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

type Consignment struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers           []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vesselId,proto3" json:"vesselId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Consignment) Reset()         { *m = Consignment{} }
func (m *Consignment) String() string { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()    {}
func (*Consignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee6334c5942fcf75, []int{0}
}

func (m *Consignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consignment.Unmarshal(m, b)
}
func (m *Consignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consignment.Marshal(b, m, deterministic)
}
func (m *Consignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consignment.Merge(m, src)
}
func (m *Consignment) XXX_Size() int {
	return xxx_messageInfo_Consignment.Size(m)
}
func (m *Consignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Consignment.DiscardUnknown(m)
}

var xxx_messageInfo_Consignment proto.InternalMessageInfo

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customerId,proto3" json:"customerId,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee6334c5942fcf75, []int{1}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee6334c5942fcf75, []int{2}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type Response struct {
	Created              bool           `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Consignment          *Consignment   `protobuf:"bytes,2,opt,name=consignment,proto3" json:"consignment,omitempty"`
	Consignments         []*Consignment `protobuf:"bytes,3,rep,name=consignments,proto3" json:"consignments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee6334c5942fcf75, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func (m *Response) GetConsignments() []*Consignment {
	if m != nil {
		return m.Consignments
	}
	return nil
}

func init() {
	proto.RegisterType((*Consignment)(nil), "consignment.Consignment")
	proto.RegisterType((*Container)(nil), "consignment.Container")
	proto.RegisterType((*GetRequest)(nil), "consignment.GetRequest")
	proto.RegisterType((*Response)(nil), "consignment.Response")
}

func init() {
	proto.RegisterFile("aetheryte-service-consignment/proto/consignment/consignment.proto", fileDescriptor_ee6334c5942fcf75)
}

var fileDescriptor_ee6334c5942fcf75 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xcd, 0x4e, 0x3a, 0x31,
	0x10, 0xff, 0x2f, 0x5f, 0x7f, 0x98, 0x25, 0x12, 0x9b, 0x88, 0x0d, 0x07, 0xb3, 0xd9, 0x13, 0x17,
	0x20, 0xc1, 0xc4, 0x83, 0xf1, 0x62, 0x36, 0x91, 0x70, 0x2d, 0x4f, 0x80, 0xbb, 0x93, 0xa5, 0x51,
	0xda, 0xb5, 0x53, 0x30, 0xbe, 0x0c, 0x2f, 0xe0, 0x4b, 0x1a, 0xba, 0x2c, 0x14, 0x0d, 0xb7, 0xfe,
	0x3e, 0xa6, 0xfd, 0xcd, 0x74, 0xe0, 0x79, 0x89, 0x76, 0x85, 0xe6, 0xcb, 0xe2, 0x88, 0xd0, 0x6c,
	0x65, 0x8a, 0xa3, 0x54, 0x2b, 0x92, 0xb9, 0x5a, 0xa3, 0xb2, 0x93, 0xc2, 0x68, 0xab, 0x27, 0x3e,
	0xe3, 0x9d, 0xc7, 0x4e, 0x65, 0xa1, 0x47, 0xc5, 0xdf, 0x01, 0x84, 0xc9, 0x09, 0xb3, 0x2b, 0xa8,
	0xc9, 0x8c, 0x07, 0x51, 0x30, 0xec, 0x88, 0x9a, 0xcc, 0x58, 0x04, 0x61, 0x86, 0x94, 0x1a, 0x59,
	0x58, 0xa9, 0x15, 0xaf, 0x39, 0xc1, 0xa7, 0x58, 0x1f, 0x5a, 0x9f, 0x28, 0xf3, 0x95, 0xe5, 0xf5,
	0x28, 0x18, 0x36, 0xc5, 0x01, 0xb1, 0x07, 0x80, 0x54, 0x2b, 0xbb, 0x94, 0x0a, 0x0d, 0xf1, 0x46,
	0x54, 0x1f, 0x86, 0xd3, 0xfe, 0xd8, 0x8f, 0x93, 0x54, 0xb2, 0xf0, 0x9c, 0x6c, 0x00, 0xed, 0x2d,
	0x12, 0xe1, 0xfb, 0x3c, 0xe3, 0x4d, 0xf7, 0xdc, 0x11, 0xc7, 0x6f, 0xd0, 0x39, 0x16, 0xfd, 0x89,
	0x7a, 0x07, 0x90, 0x6e, 0xc8, 0xea, 0x35, 0x9a, 0x79, 0x76, 0x48, 0xea, 0x31, 0xfb, 0xa0, 0xda,
	0xc8, 0x5c, 0x2a, 0x17, 0xb4, 0x23, 0x0e, 0x68, 0xcf, 0x6f, 0xc8, 0xd5, 0x34, 0x4a, 0xbe, 0x44,
	0x71, 0x17, 0x60, 0x86, 0x56, 0xe0, 0xc7, 0x06, 0xc9, 0xc6, 0xbb, 0x00, 0xda, 0x02, 0xa9, 0xd0,
	0x8a, 0x90, 0x71, 0xf8, 0x9f, 0x1a, 0x5c, 0x5a, 0x2c, 0xdf, 0x6f, 0x8b, 0x0a, 0xb2, 0x47, 0xf0,
	0xc7, 0xeb, 0x52, 0x84, 0x53, 0xfe, 0xbb, 0xed, 0xea, 0x2c, 0x7c, 0x33, 0x7b, 0x82, 0xae, 0x07,
	0x89, 0xd7, 0xdd, 0xcc, 0x2e, 0x17, 0x9f, 0xb9, 0xa7, 0xbb, 0x00, 0x7a, 0x8b, 0x95, 0x2c, 0x0a,
	0xa9, 0xf2, 0x45, 0xb9, 0x1a, 0xec, 0x05, 0xae, 0x13, 0x17, 0xcc, 0xff, 0xe2, 0x8b, 0x17, 0x0e,
	0x6e, 0xce, 0x94, 0xaa, 0xdb, 0xf8, 0x1f, 0x4b, 0xa0, 0x37, 0x43, 0xeb, 0x59, 0x89, 0xdd, 0x9e,
	0x79, 0x4f, 0x83, 0xba, 0x78, 0xc9, 0x6b, 0xcb, 0xad, 0xdf, 0xfd, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x27, 0x7f, 0x15, 0x56, 0xc3, 0x02, 0x00, 0x00,
}
