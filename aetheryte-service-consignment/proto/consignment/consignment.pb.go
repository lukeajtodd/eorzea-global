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
	Contains             []*Container `protobuf:"bytes,4,rep,name=contains,proto3" json:"contains,omitempty"`
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

func (m *Consignment) GetContains() []*Container {
	if m != nil {
		return m.Contains
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
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xcb, 0x6e, 0xea, 0x30,
	0x10, 0xbd, 0xe1, 0x75, 0xc3, 0x04, 0x5d, 0x74, 0x2d, 0x95, 0x5a, 0x2c, 0xaa, 0x28, 0x2b, 0x36,
	0x80, 0x94, 0xee, 0xaa, 0x6e, 0xaa, 0x48, 0x45, 0x6c, 0xcd, 0x17, 0xd0, 0x64, 0x14, 0xac, 0x16,
	0x3b, 0xf5, 0x18, 0xaa, 0xfe, 0x0c, 0xdb, 0xfe, 0x66, 0x85, 0xc3, 0xc3, 0xb4, 0x62, 0xe7, 0xf3,
	0x18, 0xfb, 0xcc, 0x78, 0xe0, 0x69, 0x89, 0x76, 0x85, 0xe6, 0xd3, 0xe2, 0x98, 0xd0, 0x6c, 0x65,
	0x8e, 0xe3, 0x5c, 0x2b, 0x92, 0xa5, 0x5a, 0xa3, 0xb2, 0xd3, 0xca, 0x68, 0xab, 0xa7, 0x3e, 0xe3,
	0x9d, 0x27, 0x4e, 0x65, 0x91, 0x47, 0x25, 0x5f, 0x01, 0x44, 0xd9, 0x19, 0xb3, 0x7f, 0xd0, 0x90,
	0x05, 0x0f, 0xe2, 0x60, 0xd4, 0x15, 0x0d, 0x59, 0xb0, 0x18, 0xa2, 0x02, 0x29, 0x37, 0xb2, 0xb2,
	0x52, 0x2b, 0xde, 0x70, 0x82, 0x4f, 0xb1, 0x01, 0x74, 0x3e, 0x50, 0x96, 0x2b, 0xcb, 0x9b, 0x71,
	0x30, 0x6a, 0x8b, 0x03, 0x62, 0x29, 0x84, 0xb9, 0x56, 0x76, 0x29, 0x15, 0xf1, 0x56, 0xdc, 0x1c,
	0x45, 0xe9, 0x60, 0xe2, 0x87, 0xc9, 0x6a, 0x11, 0x8d, 0x38, 0xf9, 0xd8, 0x10, 0xc2, 0x2d, 0x12,
	0xe1, 0xdb, 0xbc, 0xe0, 0x6d, 0xf7, 0xd4, 0x09, 0x27, 0xaf, 0xd0, 0x3d, 0x95, 0xfc, 0x8a, 0x79,
	0x07, 0x90, 0x6f, 0xc8, 0xea, 0x35, 0x9a, 0x79, 0x71, 0x48, 0xe9, 0x31, 0xfb, 0x90, 0xda, 0xc8,
	0x52, 0x2a, 0x17, 0xb2, 0x2b, 0x0e, 0x68, 0xcf, 0x6f, 0xc8, 0xd5, 0xb4, 0x6a, 0xbe, 0x46, 0x49,
	0x0f, 0x60, 0x86, 0x56, 0xe0, 0xfb, 0x06, 0xc9, 0x26, 0xbb, 0x00, 0x42, 0x81, 0x54, 0x69, 0x45,
	0xc8, 0x38, 0xfc, 0xcd, 0x0d, 0x2e, 0x2d, 0xd6, 0xef, 0x87, 0xe2, 0x08, 0xd9, 0x03, 0xf8, 0xa3,
	0x75, 0x29, 0xa2, 0x94, 0xff, 0x6c, 0xfa, 0x78, 0x16, 0xbe, 0x99, 0x3d, 0x42, 0xcf, 0x83, 0xc4,
	0x9b, 0x6e, 0x62, 0xd7, 0x8b, 0x2f, 0xdc, 0xe9, 0x2e, 0x80, 0xfe, 0x62, 0x25, 0xab, 0x4a, 0xaa,
	0x72, 0x51, 0xaf, 0x05, 0x7b, 0x86, 0xff, 0x99, 0x0b, 0xe6, 0x7f, 0xef, 0xd5, 0x0b, 0x87, 0x37,
	0x17, 0xca, 0xb1, 0xdb, 0xe4, 0x0f, 0xcb, 0xa0, 0x3f, 0x43, 0xeb, 0x59, 0x89, 0xdd, 0x5e, 0x78,
	0xcf, 0x83, 0xba, 0x7a, 0xc9, 0x4b, 0xc7, 0xad, 0xde, 0xfd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x05, 0x89, 0xf9, 0x26, 0xbf, 0x02, 0x00, 0x00,
}
