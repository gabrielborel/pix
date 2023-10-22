// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: pixkey.proto

package protofiles

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

type PixKeyRegistration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind      string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Key       string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	AccountId string `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *PixKeyRegistration) Reset() {
	*x = PixKeyRegistration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pixkey_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PixKeyRegistration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PixKeyRegistration) ProtoMessage() {}

func (x *PixKeyRegistration) ProtoReflect() protoreflect.Message {
	mi := &file_pixkey_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PixKeyRegistration.ProtoReflect.Descriptor instead.
func (*PixKeyRegistration) Descriptor() ([]byte, []int) {
	return file_pixkey_proto_rawDescGZIP(), []int{0}
}

func (x *PixKeyRegistration) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *PixKeyRegistration) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PixKeyRegistration) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type PixKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Key  string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *PixKey) Reset() {
	*x = PixKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pixkey_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PixKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PixKey) ProtoMessage() {}

func (x *PixKey) ProtoReflect() protoreflect.Message {
	mi := &file_pixkey_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PixKey.ProtoReflect.Descriptor instead.
func (*PixKey) Descriptor() ([]byte, []int) {
	return file_pixkey_proto_rawDescGZIP(), []int{1}
}

func (x *PixKey) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *PixKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId     string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	AccountNumber string `protobuf:"bytes,2,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	BankId        string `protobuf:"bytes,3,opt,name=bank_id,json=bankId,proto3" json:"bank_id,omitempty"`
	BankName      string `protobuf:"bytes,4,opt,name=bank_name,json=bankName,proto3" json:"bank_name,omitempty"`
	OwnerName     string `protobuf:"bytes,5,opt,name=owner_name,json=ownerName,proto3" json:"owner_name,omitempty"`
	CreatedAt     string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pixkey_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_pixkey_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_pixkey_proto_rawDescGZIP(), []int{2}
}

func (x *Account) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Account) GetAccountNumber() string {
	if x != nil {
		return x.AccountNumber
	}
	return ""
}

func (x *Account) GetBankId() string {
	if x != nil {
		return x.BankId
	}
	return ""
}

func (x *Account) GetBankName() string {
	if x != nil {
		return x.BankName
	}
	return ""
}

func (x *Account) GetOwnerName() string {
	if x != nil {
		return x.OwnerName
	}
	return ""
}

func (x *Account) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type PixKeyInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Kind      string   `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Key       string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Account   *Account `protobuf:"bytes,4,opt,name=account,proto3" json:"account,omitempty"`
	CreatedAt string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *PixKeyInfo) Reset() {
	*x = PixKeyInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pixkey_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PixKeyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PixKeyInfo) ProtoMessage() {}

func (x *PixKeyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pixkey_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PixKeyInfo.ProtoReflect.Descriptor instead.
func (*PixKeyInfo) Descriptor() ([]byte, []int) {
	return file_pixkey_proto_rawDescGZIP(), []int{3}
}

func (x *PixKeyInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PixKeyInfo) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *PixKeyInfo) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PixKeyInfo) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *PixKeyInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type PixKeyCreatedResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *PixKeyCreatedResult) Reset() {
	*x = PixKeyCreatedResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pixkey_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PixKeyCreatedResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PixKeyCreatedResult) ProtoMessage() {}

func (x *PixKeyCreatedResult) ProtoReflect() protoreflect.Message {
	mi := &file_pixkey_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PixKeyCreatedResult.ProtoReflect.Descriptor instead.
func (*PixKeyCreatedResult) Descriptor() ([]byte, []int) {
	return file_pixkey_proto_rawDescGZIP(), []int{4}
}

func (x *PixKeyCreatedResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PixKeyCreatedResult) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PixKeyCreatedResult) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pixkey_proto protoreflect.FileDescriptor

var file_pixkey_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x69, 0x78, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x61, 0x62, 0x72, 0x69,
	0x65, 0x6c, 0x62, 0x6f, 0x72, 0x65, 0x6c, 0x2e, 0x70, 0x69, 0x78, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x70, 0x69, 0x78, 0x22, 0x59, 0x0a, 0x12, 0x50, 0x69, 0x78, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x2e,
	0x0a, 0x06, 0x50, 0x69, 0x78, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0xc3,
	0x01, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x17, 0x0a, 0x07, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x62, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x6e,
	0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61,
	0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0xa9, 0x01, 0x0a, 0x0a, 0x50, 0x69, 0x78, 0x4b, 0x65, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x46, 0x0a, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x61, 0x62, 0x72, 0x69, 0x65, 0x6c, 0x62,
	0x6f, 0x72, 0x65, 0x6c, 0x2e, 0x70, 0x69, 0x78, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x70, 0x69, 0x78,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x53, 0x0a, 0x13, 0x50, 0x69, 0x78, 0x4b, 0x65, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x82, 0x02, 0x0a, 0x0a, 0x50, 0x69, 0x78, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x50, 0x69, 0x78, 0x4b, 0x65, 0x79, 0x12, 0x37, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x61, 0x62, 0x72, 0x69, 0x65, 0x6c, 0x62, 0x6f, 0x72, 0x65,
	0x6c, 0x2e, 0x70, 0x69, 0x78, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x70, 0x69, 0x78, 0x2e, 0x50, 0x69,
	0x78, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x38, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x61,
	0x62, 0x72, 0x69, 0x65, 0x6c, 0x62, 0x6f, 0x72, 0x65, 0x6c, 0x2e, 0x70, 0x69, 0x78, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x70, 0x69, 0x78, 0x2e, 0x50, 0x69, 0x78, 0x4b, 0x65, 0x79, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x6c, 0x0a, 0x0a,
	0x46, 0x69, 0x6e, 0x64, 0x50, 0x69, 0x78, 0x4b, 0x65, 0x79, 0x12, 0x2b, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x61, 0x62, 0x72, 0x69, 0x65, 0x6c, 0x62,
	0x6f, 0x72, 0x65, 0x6c, 0x2e, 0x70, 0x69, 0x78, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x70, 0x69, 0x78,
	0x2e, 0x50, 0x69, 0x78, 0x4b, 0x65, 0x79, 0x1a, 0x2f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x61, 0x62, 0x72, 0x69, 0x65, 0x6c, 0x62, 0x6f, 0x72, 0x65,
	0x6c, 0x2e, 0x70, 0x69, 0x78, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x70, 0x69, 0x78, 0x2e, 0x50, 0x69,
	0x78, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x62, 0x72, 0x69, 0x65, 0x6c,
	0x62, 0x6f, 0x72, 0x65, 0x6c, 0x2e, 0x70, 0x69, 0x78, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x70, 0x69,
	0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pixkey_proto_rawDescOnce sync.Once
	file_pixkey_proto_rawDescData = file_pixkey_proto_rawDesc
)

func file_pixkey_proto_rawDescGZIP() []byte {
	file_pixkey_proto_rawDescOnce.Do(func() {
		file_pixkey_proto_rawDescData = protoimpl.X.CompressGZIP(file_pixkey_proto_rawDescData)
	})
	return file_pixkey_proto_rawDescData
}

var file_pixkey_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pixkey_proto_goTypes = []interface{}{
	(*PixKeyRegistration)(nil),  // 0: github.com.gabrielborel.pix.codepix.PixKeyRegistration
	(*PixKey)(nil),              // 1: github.com.gabrielborel.pix.codepix.PixKey
	(*Account)(nil),             // 2: github.com.gabrielborel.pix.codepix.Account
	(*PixKeyInfo)(nil),          // 3: github.com.gabrielborel.pix.codepix.PixKeyInfo
	(*PixKeyCreatedResult)(nil), // 4: github.com.gabrielborel.pix.codepix.PixKeyCreatedResult
}
var file_pixkey_proto_depIdxs = []int32{
	2, // 0: github.com.gabrielborel.pix.codepix.PixKeyInfo.account:type_name -> github.com.gabrielborel.pix.codepix.Account
	0, // 1: github.com.gabrielborel.pix.codepix.PixService.RegisterPixKey:input_type -> github.com.gabrielborel.pix.codepix.PixKeyRegistration
	1, // 2: github.com.gabrielborel.pix.codepix.PixService.FindPixKey:input_type -> github.com.gabrielborel.pix.codepix.PixKey
	4, // 3: github.com.gabrielborel.pix.codepix.PixService.RegisterPixKey:output_type -> github.com.gabrielborel.pix.codepix.PixKeyCreatedResult
	3, // 4: github.com.gabrielborel.pix.codepix.PixService.FindPixKey:output_type -> github.com.gabrielborel.pix.codepix.PixKeyInfo
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pixkey_proto_init() }
func file_pixkey_proto_init() {
	if File_pixkey_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pixkey_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PixKeyRegistration); i {
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
		file_pixkey_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PixKey); i {
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
		file_pixkey_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_pixkey_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PixKeyInfo); i {
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
		file_pixkey_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PixKeyCreatedResult); i {
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
			RawDescriptor: file_pixkey_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pixkey_proto_goTypes,
		DependencyIndexes: file_pixkey_proto_depIdxs,
		MessageInfos:      file_pixkey_proto_msgTypes,
	}.Build()
	File_pixkey_proto = out.File
	file_pixkey_proto_rawDesc = nil
	file_pixkey_proto_goTypes = nil
	file_pixkey_proto_depIdxs = nil
}