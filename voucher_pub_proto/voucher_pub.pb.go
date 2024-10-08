// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: voucher_pub.proto

package voucher_api_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateVoucherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Percentage int64  `protobuf:"varint,2,opt,name=percentage,proto3" json:"percentage,omitempty"`
	ExpiresAt  string `protobuf:"bytes,3,opt,name=expiresAt,proto3" json:"expiresAt,omitempty"`
}

func (x *CreateVoucherRequest) Reset() {
	*x = CreateVoucherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voucher_pub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVoucherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVoucherRequest) ProtoMessage() {}

func (x *CreateVoucherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_voucher_pub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVoucherRequest.ProtoReflect.Descriptor instead.
func (*CreateVoucherRequest) Descriptor() ([]byte, []int) {
	return file_voucher_pub_proto_rawDescGZIP(), []int{0}
}

func (x *CreateVoucherRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateVoucherRequest) GetPercentage() int64 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

func (x *CreateVoucherRequest) GetExpiresAt() string {
	if x != nil {
		return x.ExpiresAt
	}
	return ""
}

type CreateVoucherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid       string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Code       string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Percentage int64  `protobuf:"varint,3,opt,name=percentage,proto3" json:"percentage,omitempty"`
	CreatedAt  string `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	ExpiresAt  string `protobuf:"bytes,5,opt,name=expiresAt,proto3" json:"expiresAt,omitempty"`
}

func (x *CreateVoucherResponse) Reset() {
	*x = CreateVoucherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voucher_pub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVoucherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVoucherResponse) ProtoMessage() {}

func (x *CreateVoucherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_voucher_pub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVoucherResponse.ProtoReflect.Descriptor instead.
func (*CreateVoucherResponse) Descriptor() ([]byte, []int) {
	return file_voucher_pub_proto_rawDescGZIP(), []int{1}
}

func (x *CreateVoucherResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *CreateVoucherResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateVoucherResponse) GetPercentage() int64 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

func (x *CreateVoucherResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CreateVoucherResponse) GetExpiresAt() string {
	if x != nil {
		return x.ExpiresAt
	}
	return ""
}

type GetVoucherByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetVoucherByIDRequest) Reset() {
	*x = GetVoucherByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voucher_pub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVoucherByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVoucherByIDRequest) ProtoMessage() {}

func (x *GetVoucherByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_voucher_pub_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVoucherByIDRequest.ProtoReflect.Descriptor instead.
func (*GetVoucherByIDRequest) Descriptor() ([]byte, []int) {
	return file_voucher_pub_proto_rawDescGZIP(), []int{2}
}

func (x *GetVoucherByIDRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GetVoucherByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid       string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Code       string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Percentage int64  `protobuf:"varint,3,opt,name=percentage,proto3" json:"percentage,omitempty"`
	CreatedAt  string `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	ExpiresAt  string `protobuf:"bytes,5,opt,name=ExpiresAt,proto3" json:"ExpiresAt,omitempty"`
}

func (x *GetVoucherByIDResponse) Reset() {
	*x = GetVoucherByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voucher_pub_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVoucherByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVoucherByIDResponse) ProtoMessage() {}

func (x *GetVoucherByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_voucher_pub_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVoucherByIDResponse.ProtoReflect.Descriptor instead.
func (*GetVoucherByIDResponse) Descriptor() ([]byte, []int) {
	return file_voucher_pub_proto_rawDescGZIP(), []int{3}
}

func (x *GetVoucherByIDResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GetVoucherByIDResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetVoucherByIDResponse) GetPercentage() int64 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

func (x *GetVoucherByIDResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GetVoucherByIDResponse) GetExpiresAt() string {
	if x != nil {
		return x.ExpiresAt
	}
	return ""
}

type UpdateVoucherByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid       string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Code       string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Percentage int64  `protobuf:"varint,3,opt,name=percentage,proto3" json:"percentage,omitempty"`
	CreatedAt  string `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	ExpiresAt  string `protobuf:"bytes,5,opt,name=expiresAt,proto3" json:"expiresAt,omitempty"`
}

func (x *UpdateVoucherByIDRequest) Reset() {
	*x = UpdateVoucherByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voucher_pub_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateVoucherByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVoucherByIDRequest) ProtoMessage() {}

func (x *UpdateVoucherByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_voucher_pub_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVoucherByIDRequest.ProtoReflect.Descriptor instead.
func (*UpdateVoucherByIDRequest) Descriptor() ([]byte, []int) {
	return file_voucher_pub_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateVoucherByIDRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *UpdateVoucherByIDRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateVoucherByIDRequest) GetPercentage() int64 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

func (x *UpdateVoucherByIDRequest) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UpdateVoucherByIDRequest) GetExpiresAt() string {
	if x != nil {
		return x.ExpiresAt
	}
	return ""
}

type UpdateVoucherByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid       string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Code       string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Percentage int64  `protobuf:"varint,3,opt,name=percentage,proto3" json:"percentage,omitempty"`
	CreatedAt  string `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	ExpiresAt  string `protobuf:"bytes,5,opt,name=expiresAt,proto3" json:"expiresAt,omitempty"`
}

func (x *UpdateVoucherByIDResponse) Reset() {
	*x = UpdateVoucherByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voucher_pub_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateVoucherByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVoucherByIDResponse) ProtoMessage() {}

func (x *UpdateVoucherByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_voucher_pub_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVoucherByIDResponse.ProtoReflect.Descriptor instead.
func (*UpdateVoucherByIDResponse) Descriptor() ([]byte, []int) {
	return file_voucher_pub_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateVoucherByIDResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *UpdateVoucherByIDResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateVoucherByIDResponse) GetPercentage() int64 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

func (x *UpdateVoucherByIDResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UpdateVoucherByIDResponse) GetExpiresAt() string {
	if x != nil {
		return x.ExpiresAt
	}
	return ""
}

var File_voucher_pub_proto protoreflect.FileDescriptor

var file_voucher_pub_proto_rawDesc = []byte{
	0x0a, 0x11, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x75,
	0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x22, 0x9b, 0x01,
	0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x22, 0x2b, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x9c, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x22, 0x9e, 0x01, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x22, 0x9f, 0x01, 0x0a, 0x19, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x32, 0xd8, 0x01, 0x0a, 0x07, 0x56,
	0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x75,
	0x63, 0x68, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x6f,
	0x75, 0x63, 0x68, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12, 0x19,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2a, 0x5a, 0x28, 0x68, 0x66, 0x2d, 0x76, 0x6f, 0x75, 0x63,
	0x68, 0x65, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_voucher_pub_proto_rawDescOnce sync.Once
	file_voucher_pub_proto_rawDescData = file_voucher_pub_proto_rawDesc
)

func file_voucher_pub_proto_rawDescGZIP() []byte {
	file_voucher_pub_proto_rawDescOnce.Do(func() {
		file_voucher_pub_proto_rawDescData = protoimpl.X.CompressGZIP(file_voucher_pub_proto_rawDescData)
	})
	return file_voucher_pub_proto_rawDescData
}

var file_voucher_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_voucher_pub_proto_goTypes = []any{
	(*CreateVoucherRequest)(nil),      // 0: CreateVoucherRequest
	(*CreateVoucherResponse)(nil),     // 1: CreateVoucherResponse
	(*GetVoucherByIDRequest)(nil),     // 2: GetVoucherByIDRequest
	(*GetVoucherByIDResponse)(nil),    // 3: GetVoucherByIDResponse
	(*UpdateVoucherByIDRequest)(nil),  // 4: UpdateVoucherByIDRequest
	(*UpdateVoucherByIDResponse)(nil), // 5: UpdateVoucherByIDResponse
}
var file_voucher_pub_proto_depIdxs = []int32{
	0, // 0: Voucher.CreateVoucher:input_type -> CreateVoucherRequest
	2, // 1: Voucher.GetVoucherByID:input_type -> GetVoucherByIDRequest
	4, // 2: Voucher.UpdateVoucherByID:input_type -> UpdateVoucherByIDRequest
	1, // 3: Voucher.CreateVoucher:output_type -> CreateVoucherResponse
	3, // 4: Voucher.GetVoucherByID:output_type -> GetVoucherByIDResponse
	5, // 5: Voucher.UpdateVoucherByID:output_type -> UpdateVoucherByIDResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_voucher_pub_proto_init() }
func file_voucher_pub_proto_init() {
	if File_voucher_pub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_voucher_pub_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateVoucherRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_voucher_pub_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateVoucherResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_voucher_pub_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetVoucherByIDRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_voucher_pub_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetVoucherByIDResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_voucher_pub_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateVoucherByIDRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_voucher_pub_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateVoucherByIDResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_voucher_pub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_voucher_pub_proto_goTypes,
		DependencyIndexes: file_voucher_pub_proto_depIdxs,
		MessageInfos:      file_voucher_pub_proto_msgTypes,
	}.Build()
	File_voucher_pub_proto = out.File
	file_voucher_pub_proto_rawDesc = nil
	file_voucher_pub_proto_goTypes = nil
	file_voucher_pub_proto_depIdxs = nil
}
