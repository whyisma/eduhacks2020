// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Send2ClientReq struct {
	SystemID             string   `protobuf:"bytes,1,opt,name=systemID,proto3" json:"systemID,omitempty"`
	MessageID            string   `protobuf:"bytes,2,opt,name=messageID,proto3" json:"messageID,omitempty"`
	SendUserID           string   `protobuf:"bytes,3,opt,name=sendUserID,proto3" json:"sendUserID,omitempty"`
	ClientID             string   `protobuf:"bytes,4,opt,name=clientID,proto3" json:"clientID,omitempty"`
	Code                 int32    `protobuf:"varint,5,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	Data                 string   `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Send2ClientReq) Reset()         { *m = Send2ClientReq{} }
func (m *Send2ClientReq) String() string { return proto.CompactTextString(m) }
func (*Send2ClientReq) ProtoMessage()    {}
func (*Send2ClientReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{0}
}

func (m *Send2ClientReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Send2ClientReq.Unmarshal(m, b)
}
func (m *Send2ClientReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Send2ClientReq.Marshal(b, m, deterministic)
}
func (m *Send2ClientReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Send2ClientReq.Merge(m, src)
}
func (m *Send2ClientReq) XXX_Size() int {
	return xxx_messageInfo_Send2ClientReq.Size(m)
}
func (m *Send2ClientReq) XXX_DiscardUnknown() {
	xxx_messageInfo_Send2ClientReq.DiscardUnknown(m)
}

var xxx_messageInfo_Send2ClientReq proto.InternalMessageInfo

func (m *Send2ClientReq) GetSystemID() string {
	if m != nil {
		return m.SystemID
	}
	return ""
}

func (m *Send2ClientReq) GetMessageID() string {
	if m != nil {
		return m.MessageID
	}
	return ""
}

func (m *Send2ClientReq) GetSendUserID() string {
	if m != nil {
		return m.SendUserID
	}
	return ""
}

func (m *Send2ClientReq) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *Send2ClientReq) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Send2ClientReq) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Send2ClientReq) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type CloseClientReq struct {
	SystemID             string   `protobuf:"bytes,1,opt,name=systemID,proto3" json:"systemID,omitempty"`
	ClientID             string   `protobuf:"bytes,2,opt,name=clientID,proto3" json:"clientID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseClientReq) Reset()         { *m = CloseClientReq{} }
func (m *CloseClientReq) String() string { return proto.CompactTextString(m) }
func (*CloseClientReq) ProtoMessage()    {}
func (*CloseClientReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{1}
}

func (m *CloseClientReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseClientReq.Unmarshal(m, b)
}
func (m *CloseClientReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseClientReq.Marshal(b, m, deterministic)
}
func (m *CloseClientReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseClientReq.Merge(m, src)
}
func (m *CloseClientReq) XXX_Size() int {
	return xxx_messageInfo_CloseClientReq.Size(m)
}
func (m *CloseClientReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseClientReq.DiscardUnknown(m)
}

var xxx_messageInfo_CloseClientReq proto.InternalMessageInfo

func (m *CloseClientReq) GetSystemID() string {
	if m != nil {
		return m.SystemID
	}
	return ""
}

func (m *CloseClientReq) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

type BindGroupReq struct {
	SystemID             string   `protobuf:"bytes,1,opt,name=systemID,proto3" json:"systemID,omitempty"`
	GroupName            string   `protobuf:"bytes,2,opt,name=groupName,proto3" json:"groupName,omitempty"`
	ClientID             string   `protobuf:"bytes,3,opt,name=clientID,proto3" json:"clientID,omitempty"`
	UserID               string   `protobuf:"bytes,4,opt,name=userID,proto3" json:"userID,omitempty"`
	Extend               string   `protobuf:"bytes,5,opt,name=extend,proto3" json:"extend,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BindGroupReq) Reset()         { *m = BindGroupReq{} }
func (m *BindGroupReq) String() string { return proto.CompactTextString(m) }
func (*BindGroupReq) ProtoMessage()    {}
func (*BindGroupReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{2}
}

func (m *BindGroupReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindGroupReq.Unmarshal(m, b)
}
func (m *BindGroupReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindGroupReq.Marshal(b, m, deterministic)
}
func (m *BindGroupReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindGroupReq.Merge(m, src)
}
func (m *BindGroupReq) XXX_Size() int {
	return xxx_messageInfo_BindGroupReq.Size(m)
}
func (m *BindGroupReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BindGroupReq.DiscardUnknown(m)
}

var xxx_messageInfo_BindGroupReq proto.InternalMessageInfo

func (m *BindGroupReq) GetSystemID() string {
	if m != nil {
		return m.SystemID
	}
	return ""
}

func (m *BindGroupReq) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *BindGroupReq) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *BindGroupReq) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *BindGroupReq) GetExtend() string {
	if m != nil {
		return m.Extend
	}
	return ""
}

type Send2GroupReq struct {
	SystemID             string   `protobuf:"bytes,1,opt,name=systemID,proto3" json:"systemID,omitempty"`
	MessageID            string   `protobuf:"bytes,2,opt,name=messageID,proto3" json:"messageID,omitempty"`
	SendUserID           string   `protobuf:"bytes,3,opt,name=sendUserID,proto3" json:"sendUserID,omitempty"`
	GroupName            string   `protobuf:"bytes,4,opt,name=groupName,proto3" json:"groupName,omitempty"`
	Code                 int32    `protobuf:"varint,5,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	Data                 string   `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Send2GroupReq) Reset()         { *m = Send2GroupReq{} }
func (m *Send2GroupReq) String() string { return proto.CompactTextString(m) }
func (*Send2GroupReq) ProtoMessage()    {}
func (*Send2GroupReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{3}
}

func (m *Send2GroupReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Send2GroupReq.Unmarshal(m, b)
}
func (m *Send2GroupReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Send2GroupReq.Marshal(b, m, deterministic)
}
func (m *Send2GroupReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Send2GroupReq.Merge(m, src)
}
func (m *Send2GroupReq) XXX_Size() int {
	return xxx_messageInfo_Send2GroupReq.Size(m)
}
func (m *Send2GroupReq) XXX_DiscardUnknown() {
	xxx_messageInfo_Send2GroupReq.DiscardUnknown(m)
}

var xxx_messageInfo_Send2GroupReq proto.InternalMessageInfo

func (m *Send2GroupReq) GetSystemID() string {
	if m != nil {
		return m.SystemID
	}
	return ""
}

func (m *Send2GroupReq) GetMessageID() string {
	if m != nil {
		return m.MessageID
	}
	return ""
}

func (m *Send2GroupReq) GetSendUserID() string {
	if m != nil {
		return m.SendUserID
	}
	return ""
}

func (m *Send2GroupReq) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *Send2GroupReq) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Send2GroupReq) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Send2GroupReq) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type Send2SystemReq struct {
	SystemID             string   `protobuf:"bytes,1,opt,name=systemID,proto3" json:"systemID,omitempty"`
	MessageID            string   `protobuf:"bytes,2,opt,name=messageID,proto3" json:"messageID,omitempty"`
	SendUserID           string   `protobuf:"bytes,3,opt,name=sendUserID,proto3" json:"sendUserID,omitempty"`
	Code                 int32    `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	Data                 string   `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Send2SystemReq) Reset()         { *m = Send2SystemReq{} }
func (m *Send2SystemReq) String() string { return proto.CompactTextString(m) }
func (*Send2SystemReq) ProtoMessage()    {}
func (*Send2SystemReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{4}
}

func (m *Send2SystemReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Send2SystemReq.Unmarshal(m, b)
}
func (m *Send2SystemReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Send2SystemReq.Marshal(b, m, deterministic)
}
func (m *Send2SystemReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Send2SystemReq.Merge(m, src)
}
func (m *Send2SystemReq) XXX_Size() int {
	return xxx_messageInfo_Send2SystemReq.Size(m)
}
func (m *Send2SystemReq) XXX_DiscardUnknown() {
	xxx_messageInfo_Send2SystemReq.DiscardUnknown(m)
}

var xxx_messageInfo_Send2SystemReq proto.InternalMessageInfo

func (m *Send2SystemReq) GetSystemID() string {
	if m != nil {
		return m.SystemID
	}
	return ""
}

func (m *Send2SystemReq) GetMessageID() string {
	if m != nil {
		return m.MessageID
	}
	return ""
}

func (m *Send2SystemReq) GetSendUserID() string {
	if m != nil {
		return m.SendUserID
	}
	return ""
}

func (m *Send2SystemReq) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Send2SystemReq) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Send2SystemReq) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type GetGroupClientsReq struct {
	SystemID             string   `protobuf:"bytes,1,opt,name=systemID,proto3" json:"systemID,omitempty"`
	GroupName            string   `protobuf:"bytes,2,opt,name=groupName,proto3" json:"groupName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupClientsReq) Reset()         { *m = GetGroupClientsReq{} }
func (m *GetGroupClientsReq) String() string { return proto.CompactTextString(m) }
func (*GetGroupClientsReq) ProtoMessage()    {}
func (*GetGroupClientsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{5}
}

func (m *GetGroupClientsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupClientsReq.Unmarshal(m, b)
}
func (m *GetGroupClientsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupClientsReq.Marshal(b, m, deterministic)
}
func (m *GetGroupClientsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupClientsReq.Merge(m, src)
}
func (m *GetGroupClientsReq) XXX_Size() int {
	return xxx_messageInfo_GetGroupClientsReq.Size(m)
}
func (m *GetGroupClientsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupClientsReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupClientsReq proto.InternalMessageInfo

func (m *GetGroupClientsReq) GetSystemID() string {
	if m != nil {
		return m.SystemID
	}
	return ""
}

func (m *GetGroupClientsReq) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

type Send2ClientReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Send2ClientReply) Reset()         { *m = Send2ClientReply{} }
func (m *Send2ClientReply) String() string { return proto.CompactTextString(m) }
func (*Send2ClientReply) ProtoMessage()    {}
func (*Send2ClientReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{6}
}

func (m *Send2ClientReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Send2ClientReply.Unmarshal(m, b)
}
func (m *Send2ClientReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Send2ClientReply.Marshal(b, m, deterministic)
}
func (m *Send2ClientReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Send2ClientReply.Merge(m, src)
}
func (m *Send2ClientReply) XXX_Size() int {
	return xxx_messageInfo_Send2ClientReply.Size(m)
}
func (m *Send2ClientReply) XXX_DiscardUnknown() {
	xxx_messageInfo_Send2ClientReply.DiscardUnknown(m)
}

var xxx_messageInfo_Send2ClientReply proto.InternalMessageInfo

type CloseClientReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseClientReply) Reset()         { *m = CloseClientReply{} }
func (m *CloseClientReply) String() string { return proto.CompactTextString(m) }
func (*CloseClientReply) ProtoMessage()    {}
func (*CloseClientReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{7}
}

func (m *CloseClientReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseClientReply.Unmarshal(m, b)
}
func (m *CloseClientReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseClientReply.Marshal(b, m, deterministic)
}
func (m *CloseClientReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseClientReply.Merge(m, src)
}
func (m *CloseClientReply) XXX_Size() int {
	return xxx_messageInfo_CloseClientReply.Size(m)
}
func (m *CloseClientReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseClientReply.DiscardUnknown(m)
}

var xxx_messageInfo_CloseClientReply proto.InternalMessageInfo

type BindGroupReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BindGroupReply) Reset()         { *m = BindGroupReply{} }
func (m *BindGroupReply) String() string { return proto.CompactTextString(m) }
func (*BindGroupReply) ProtoMessage()    {}
func (*BindGroupReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{8}
}

func (m *BindGroupReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindGroupReply.Unmarshal(m, b)
}
func (m *BindGroupReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindGroupReply.Marshal(b, m, deterministic)
}
func (m *BindGroupReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindGroupReply.Merge(m, src)
}
func (m *BindGroupReply) XXX_Size() int {
	return xxx_messageInfo_BindGroupReply.Size(m)
}
func (m *BindGroupReply) XXX_DiscardUnknown() {
	xxx_messageInfo_BindGroupReply.DiscardUnknown(m)
}

var xxx_messageInfo_BindGroupReply proto.InternalMessageInfo

type Send2GroupReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Send2GroupReply) Reset()         { *m = Send2GroupReply{} }
func (m *Send2GroupReply) String() string { return proto.CompactTextString(m) }
func (*Send2GroupReply) ProtoMessage()    {}
func (*Send2GroupReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{9}
}

func (m *Send2GroupReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Send2GroupReply.Unmarshal(m, b)
}
func (m *Send2GroupReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Send2GroupReply.Marshal(b, m, deterministic)
}
func (m *Send2GroupReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Send2GroupReply.Merge(m, src)
}
func (m *Send2GroupReply) XXX_Size() int {
	return xxx_messageInfo_Send2GroupReply.Size(m)
}
func (m *Send2GroupReply) XXX_DiscardUnknown() {
	xxx_messageInfo_Send2GroupReply.DiscardUnknown(m)
}

var xxx_messageInfo_Send2GroupReply proto.InternalMessageInfo

type Send2SystemReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Send2SystemReply) Reset()         { *m = Send2SystemReply{} }
func (m *Send2SystemReply) String() string { return proto.CompactTextString(m) }
func (*Send2SystemReply) ProtoMessage()    {}
func (*Send2SystemReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{10}
}

func (m *Send2SystemReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Send2SystemReply.Unmarshal(m, b)
}
func (m *Send2SystemReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Send2SystemReply.Marshal(b, m, deterministic)
}
func (m *Send2SystemReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Send2SystemReply.Merge(m, src)
}
func (m *Send2SystemReply) XXX_Size() int {
	return xxx_messageInfo_Send2SystemReply.Size(m)
}
func (m *Send2SystemReply) XXX_DiscardUnknown() {
	xxx_messageInfo_Send2SystemReply.DiscardUnknown(m)
}

var xxx_messageInfo_Send2SystemReply proto.InternalMessageInfo

type GetGroupClientsReply struct {
	List                 []string `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupClientsReply) Reset()         { *m = GetGroupClientsReply{} }
func (m *GetGroupClientsReply) String() string { return proto.CompactTextString(m) }
func (*GetGroupClientsReply) ProtoMessage()    {}
func (*GetGroupClientsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{11}
}

func (m *GetGroupClientsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupClientsReply.Unmarshal(m, b)
}
func (m *GetGroupClientsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupClientsReply.Marshal(b, m, deterministic)
}
func (m *GetGroupClientsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupClientsReply.Merge(m, src)
}
func (m *GetGroupClientsReply) XXX_Size() int {
	return xxx_messageInfo_GetGroupClientsReply.Size(m)
}
func (m *GetGroupClientsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupClientsReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupClientsReply proto.InternalMessageInfo

func (m *GetGroupClientsReply) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*Send2ClientReq)(nil), "Send2ClientReq")
	proto.RegisterType((*CloseClientReq)(nil), "CloseClientReq")
	proto.RegisterType((*BindGroupReq)(nil), "BindGroupReq")
	proto.RegisterType((*Send2GroupReq)(nil), "Send2GroupReq")
	proto.RegisterType((*Send2SystemReq)(nil), "Send2SystemReq")
	proto.RegisterType((*GetGroupClientsReq)(nil), "GetGroupClientsReq")
	proto.RegisterType((*Send2ClientReply)(nil), "Send2ClientReply")
	proto.RegisterType((*CloseClientReply)(nil), "CloseClientReply")
	proto.RegisterType((*BindGroupReply)(nil), "BindGroupReply")
	proto.RegisterType((*Send2GroupReply)(nil), "Send2GroupReply")
	proto.RegisterType((*Send2SystemReply)(nil), "Send2SystemReply")
	proto.RegisterType((*GetGroupClientsReply)(nil), "GetGroupClientsReply")
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_bedfbfc9b54e5600) }

var fileDescriptor_bedfbfc9b54e5600 = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcb, 0x8a, 0xdb, 0x30,
	0x14, 0xb5, 0xf3, 0x9a, 0xe6, 0x4e, 0x13, 0x27, 0xb7, 0x0f, 0x84, 0x29, 0x25, 0x78, 0x35, 0x14,
	0xaa, 0x96, 0x99, 0x0f, 0x28, 0x4c, 0x02, 0xd3, 0x6c, 0x66, 0x91, 0xd0, 0x4d, 0x77, 0x19, 0xfb,
	0x12, 0x02, 0xb6, 0xe5, 0x5a, 0x9e, 0xd2, 0xfc, 0x47, 0x3f, 0xa3, 0x1f, 0xd2, 0xee, 0xfa, 0x49,
	0x45, 0x72, 0xec, 0x91, 0xec, 0x2c, 0x02, 0xa1, 0x3b, 0xdd, 0xa3, 0xd7, 0x39, 0x47, 0xe7, 0x0a,
	0x60, 0x9b, 0x67, 0x21, 0xcf, 0x72, 0x51, 0x88, 0xe0, 0xb7, 0x0b, 0xe3, 0x35, 0xa5, 0xd1, 0xf5,
	0x3c, 0xde, 0x51, 0x5a, 0xac, 0xe8, 0x1b, 0xfa, 0xf0, 0x4c, 0xee, 0x65, 0x41, 0xc9, 0x72, 0xc1,
	0xdc, 0x99, 0x7b, 0x35, 0x5c, 0xd5, 0x35, 0xbe, 0x81, 0x61, 0x42, 0x52, 0x6e, 0xb6, 0xb4, 0x5c,
	0xb0, 0x8e, 0x9e, 0x7c, 0x02, 0xf0, 0x2d, 0x80, 0xa4, 0x34, 0xfa, 0x22, 0x29, 0x5f, 0x2e, 0x58,
	0x57, 0x4f, 0x1b, 0x88, 0x3a, 0x39, 0xd4, 0xd7, 0x2c, 0x17, 0xac, 0x57, 0x9e, 0x5c, 0xd5, 0x88,
	0xd0, 0x0b, 0x45, 0x44, 0xac, 0x3f, 0x73, 0xaf, 0xfa, 0x2b, 0x3d, 0x46, 0x06, 0x17, 0x87, 0xc3,
	0xd9, 0x40, 0x2f, 0xaf, 0x4a, 0xb5, 0x3a, 0xda, 0x14, 0x1b, 0x76, 0xa1, 0x61, 0x3d, 0x0e, 0x3e,
	0xc3, 0x78, 0x1e, 0x0b, 0x49, 0xa7, 0x29, 0x31, 0xb9, 0x74, 0x6c, 0x2e, 0xc1, 0x4f, 0x17, 0x9e,
	0xdf, 0xee, 0xd2, 0xe8, 0x2e, 0x17, 0x8f, 0xd9, 0x09, 0x96, 0x6c, 0xd5, 0xba, 0xfb, 0x4d, 0x42,
	0x95, 0x25, 0x35, 0x60, 0x5d, 0xd3, 0x6d, 0x48, 0x7e, 0x0d, 0x83, 0xc7, 0xd2, 0xaa, 0xd2, 0x8c,
	0x43, 0xa5, 0x70, 0xfa, 0x51, 0x50, 0x1a, 0x69, 0x33, 0x86, 0xab, 0x43, 0x15, 0xfc, 0x71, 0x61,
	0xa4, 0xdf, 0xea, 0x54, 0x5e, 0x67, 0x3c, 0x95, 0xa5, 0xaa, 0xd7, 0x54, 0x75, 0xfe, 0x63, 0xfd,
	0xaa, 0x72, 0xb7, 0xd6, 0x7c, 0xff, 0xaf, 0x98, 0x8a, 0x6e, 0xef, 0x38, 0xdd, 0xfe, 0x71, 0xba,
	0x03, 0x83, 0xee, 0x3d, 0xe0, 0x1d, 0x15, 0xda, 0xf7, 0x32, 0x5e, 0xf2, 0xac, 0x58, 0x04, 0x08,
	0x13, 0xab, 0xeb, 0xb2, 0x78, 0xaf, 0x30, 0x2b, 0xbf, 0x0a, 0x9b, 0xc0, 0xd8, 0x08, 0xa2, 0x42,
	0xa6, 0xe0, 0x99, 0x19, 0x38, 0x6c, 0xb4, 0xac, 0x54, 0xd8, 0x3b, 0x78, 0xd9, 0x22, 0x9c, 0xc5,
	0x7b, 0x25, 0x2e, 0xde, 0xc9, 0x82, 0xb9, 0xb3, 0xae, 0x12, 0xa7, 0xc6, 0xd7, 0x7f, 0x3b, 0x30,
	0x9a, 0x8b, 0x24, 0x11, 0xe9, 0x9a, 0xf2, 0xef, 0xbb, 0x90, 0xf0, 0x06, 0x2e, 0x0d, 0x7a, 0xe8,
	0x71, 0xfb, 0x8b, 0xf0, 0xa7, 0xbc, 0xc5, 0xde, 0x51, 0x9b, 0x0c, 0xfe, 0xe8, 0x71, 0xbb, 0x1b,
	0xfd, 0x29, 0x6f, 0xc9, 0x73, 0xf0, 0x3d, 0x0c, 0x6b, 0x81, 0x38, 0xe2, 0x66, 0xd7, 0xf9, 0x1e,
	0x6f, 0x68, 0x77, 0xf0, 0x23, 0xc0, 0x93, 0x7a, 0x1c, 0x73, 0xab, 0x1d, 0xfc, 0x09, 0x6f, 0x5a,
	0xe3, 0xd4, 0x52, 0x4a, 0x73, 0x2a, 0x29, 0x75, 0xea, 0x2a, 0x29, 0xa6, 0x77, 0x0e, 0x7e, 0x02,
	0xaf, 0xe1, 0x1e, 0xbe, 0xe0, 0xed, 0x00, 0xf8, 0xaf, 0xf8, 0x31, 0x93, 0x03, 0xe7, 0x76, 0xf4,
	0xf5, 0x52, 0xff, 0xaf, 0xa1, 0x88, 0x3f, 0x64, 0x0f, 0x0f, 0x03, 0x5d, 0xdc, 0xfc, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0xdd, 0xeb, 0x5a, 0x9e, 0x7a, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommonServiceClient is the client API for CommonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommonServiceClient interface {
	Send2Client(ctx context.Context, in *Send2ClientReq, opts ...grpc.CallOption) (*Send2ClientReply, error)
	CloseClient(ctx context.Context, in *CloseClientReq, opts ...grpc.CallOption) (*CloseClientReply, error)
	BindGroup(ctx context.Context, in *BindGroupReq, opts ...grpc.CallOption) (*BindGroupReply, error)
	Send2Group(ctx context.Context, in *Send2GroupReq, opts ...grpc.CallOption) (*Send2GroupReply, error)
	Send2System(ctx context.Context, in *Send2SystemReq, opts ...grpc.CallOption) (*Send2SystemReply, error)
	GetGroupClients(ctx context.Context, in *GetGroupClientsReq, opts ...grpc.CallOption) (*GetGroupClientsReply, error)
}

type commonServiceClient struct {
	cc *grpc.ClientConn
}

func NewCommonServiceClient(cc *grpc.ClientConn) CommonServiceClient {
	return &commonServiceClient{cc}
}

func (c *commonServiceClient) Send2Client(ctx context.Context, in *Send2ClientReq, opts ...grpc.CallOption) (*Send2ClientReply, error) {
	out := new(Send2ClientReply)
	err := c.cc.Invoke(ctx, "/CommonService/Send2Client", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonServiceClient) CloseClient(ctx context.Context, in *CloseClientReq, opts ...grpc.CallOption) (*CloseClientReply, error) {
	out := new(CloseClientReply)
	err := c.cc.Invoke(ctx, "/CommonService/CloseClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonServiceClient) BindGroup(ctx context.Context, in *BindGroupReq, opts ...grpc.CallOption) (*BindGroupReply, error) {
	out := new(BindGroupReply)
	err := c.cc.Invoke(ctx, "/CommonService/BindGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonServiceClient) Send2Group(ctx context.Context, in *Send2GroupReq, opts ...grpc.CallOption) (*Send2GroupReply, error) {
	out := new(Send2GroupReply)
	err := c.cc.Invoke(ctx, "/CommonService/Send2Group", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonServiceClient) Send2System(ctx context.Context, in *Send2SystemReq, opts ...grpc.CallOption) (*Send2SystemReply, error) {
	out := new(Send2SystemReply)
	err := c.cc.Invoke(ctx, "/CommonService/Send2System", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonServiceClient) GetGroupClients(ctx context.Context, in *GetGroupClientsReq, opts ...grpc.CallOption) (*GetGroupClientsReply, error) {
	out := new(GetGroupClientsReply)
	err := c.cc.Invoke(ctx, "/CommonService/GetGroupClients", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommonServiceServer is the server API for CommonService service.
type CommonServiceServer interface {
	Send2Client(context.Context, *Send2ClientReq) (*Send2ClientReply, error)
	CloseClient(context.Context, *CloseClientReq) (*CloseClientReply, error)
	BindGroup(context.Context, *BindGroupReq) (*BindGroupReply, error)
	Send2Group(context.Context, *Send2GroupReq) (*Send2GroupReply, error)
	Send2System(context.Context, *Send2SystemReq) (*Send2SystemReply, error)
	GetGroupClients(context.Context, *GetGroupClientsReq) (*GetGroupClientsReply, error)
}

// UnimplementedCommonServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCommonServiceServer struct {
}

func (*UnimplementedCommonServiceServer) Send2Client(ctx context.Context, req *Send2ClientReq) (*Send2ClientReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send2Client not implemented")
}
func (*UnimplementedCommonServiceServer) CloseClient(ctx context.Context, req *CloseClientReq) (*CloseClientReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseClient not implemented")
}
func (*UnimplementedCommonServiceServer) BindGroup(ctx context.Context, req *BindGroupReq) (*BindGroupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindGroup not implemented")
}
func (*UnimplementedCommonServiceServer) Send2Group(ctx context.Context, req *Send2GroupReq) (*Send2GroupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send2Group not implemented")
}
func (*UnimplementedCommonServiceServer) Send2System(ctx context.Context, req *Send2SystemReq) (*Send2SystemReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send2System not implemented")
}
func (*UnimplementedCommonServiceServer) GetGroupClients(ctx context.Context, req *GetGroupClientsReq) (*GetGroupClientsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupClients not implemented")
}

func RegisterCommonServiceServer(s *grpc.Server, srv CommonServiceServer) {
	s.RegisterService(&_CommonService_serviceDesc, srv)
}

func _CommonService_Send2Client_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Send2ClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).Send2Client(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CommonService/Send2Client",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).Send2Client(ctx, req.(*Send2ClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommonService_CloseClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).CloseClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CommonService/CloseClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).CloseClient(ctx, req.(*CloseClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommonService_BindGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).BindGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CommonService/BindGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).BindGroup(ctx, req.(*BindGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommonService_Send2Group_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Send2GroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).Send2Group(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CommonService/Send2Group",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).Send2Group(ctx, req.(*Send2GroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommonService_Send2System_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Send2SystemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).Send2System(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CommonService/Send2System",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).Send2System(ctx, req.(*Send2SystemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommonService_GetGroupClients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupClientsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).GetGroupClients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CommonService/GetGroupClients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).GetGroupClients(ctx, req.(*GetGroupClientsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommonService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CommonService",
	HandlerType: (*CommonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send2Client",
			Handler:    _CommonService_Send2Client_Handler,
		},
		{
			MethodName: "CloseClient",
			Handler:    _CommonService_CloseClient_Handler,
		},
		{
			MethodName: "BindGroup",
			Handler:    _CommonService_BindGroup_Handler,
		},
		{
			MethodName: "Send2Group",
			Handler:    _CommonService_Send2Group_Handler,
		},
		{
			MethodName: "Send2System",
			Handler:    _CommonService_Send2System_Handler,
		},
		{
			MethodName: "GetGroupClients",
			Handler:    _CommonService_GetGroupClients_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}