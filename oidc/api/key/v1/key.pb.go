// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: oidc/api/key/v1/key.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/go-kratos/kratos/v2/errors"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeleteJsonWebKeySetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Set string `protobuf:"bytes,1,opt,name=set,proto3" json:"set,omitempty"`
}

func (x *DeleteJsonWebKeySetRequest) Reset() {
	*x = DeleteJsonWebKeySetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidc_api_key_v1_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJsonWebKeySetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJsonWebKeySetRequest) ProtoMessage() {}

func (x *DeleteJsonWebKeySetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oidc_api_key_v1_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJsonWebKeySetRequest.ProtoReflect.Descriptor instead.
func (*DeleteJsonWebKeySetRequest) Descriptor() ([]byte, []int) {
	return file_oidc_api_key_v1_key_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteJsonWebKeySetRequest) GetSet() string {
	if x != nil {
		return x.Set
	}
	return ""
}

type GetJsonWebKeySetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Set string `protobuf:"bytes,1,opt,name=set,proto3" json:"set,omitempty"`
}

func (x *GetJsonWebKeySetRequest) Reset() {
	*x = GetJsonWebKeySetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidc_api_key_v1_key_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJsonWebKeySetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJsonWebKeySetRequest) ProtoMessage() {}

func (x *GetJsonWebKeySetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oidc_api_key_v1_key_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJsonWebKeySetRequest.ProtoReflect.Descriptor instead.
func (*GetJsonWebKeySetRequest) Descriptor() ([]byte, []int) {
	return file_oidc_api_key_v1_key_proto_rawDescGZIP(), []int{1}
}

func (x *GetJsonWebKeySetRequest) GetSet() string {
	if x != nil {
		return x.Set
	}
	return ""
}

type UpdateJsonWebKeySetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Set  string         `protobuf:"bytes,1,opt,name=set,proto3" json:"set,omitempty"`
	Keys *JsonWebKeySet `protobuf:"bytes,2,opt,name=keys,proto3" json:"keys,omitempty"`
}

func (x *UpdateJsonWebKeySetRequest) Reset() {
	*x = UpdateJsonWebKeySetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidc_api_key_v1_key_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJsonWebKeySetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJsonWebKeySetRequest) ProtoMessage() {}

func (x *UpdateJsonWebKeySetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oidc_api_key_v1_key_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJsonWebKeySetRequest.ProtoReflect.Descriptor instead.
func (*UpdateJsonWebKeySetRequest) Descriptor() ([]byte, []int) {
	return file_oidc_api_key_v1_key_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateJsonWebKeySetRequest) GetSet() string {
	if x != nil {
		return x.Set
	}
	return ""
}

func (x *UpdateJsonWebKeySetRequest) GetKeys() *JsonWebKeySet {
	if x != nil {
		return x.Keys
	}
	return nil
}

type CreateJsonWebKeySetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Set  string                         `protobuf:"bytes,1,opt,name=set,proto3" json:"set,omitempty"`
	Keys *JsonWebKeySetGeneratorRequest `protobuf:"bytes,2,opt,name=keys,proto3" json:"keys,omitempty"`
}

func (x *CreateJsonWebKeySetRequest) Reset() {
	*x = CreateJsonWebKeySetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidc_api_key_v1_key_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJsonWebKeySetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJsonWebKeySetRequest) ProtoMessage() {}

func (x *CreateJsonWebKeySetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oidc_api_key_v1_key_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJsonWebKeySetRequest.ProtoReflect.Descriptor instead.
func (*CreateJsonWebKeySetRequest) Descriptor() ([]byte, []int) {
	return file_oidc_api_key_v1_key_proto_rawDescGZIP(), []int{3}
}

func (x *CreateJsonWebKeySetRequest) GetSet() string {
	if x != nil {
		return x.Set
	}
	return ""
}

func (x *CreateJsonWebKeySetRequest) GetKeys() *JsonWebKeySetGeneratorRequest {
	if x != nil {
		return x.Keys
	}
	return nil
}

type JsonWebKeySetGeneratorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alg string `protobuf:"bytes,1,opt,name=alg,proto3" json:"alg,omitempty"`
	Kid string `protobuf:"bytes,8,opt,name=kid,proto3" json:"kid,omitempty"`
	Use string `protobuf:"bytes,14,opt,name=use,proto3" json:"use,omitempty"`
}

func (x *JsonWebKeySetGeneratorRequest) Reset() {
	*x = JsonWebKeySetGeneratorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidc_api_key_v1_key_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonWebKeySetGeneratorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonWebKeySetGeneratorRequest) ProtoMessage() {}

func (x *JsonWebKeySetGeneratorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oidc_api_key_v1_key_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonWebKeySetGeneratorRequest.ProtoReflect.Descriptor instead.
func (*JsonWebKeySetGeneratorRequest) Descriptor() ([]byte, []int) {
	return file_oidc_api_key_v1_key_proto_rawDescGZIP(), []int{4}
}

func (x *JsonWebKeySetGeneratorRequest) GetAlg() string {
	if x != nil {
		return x.Alg
	}
	return ""
}

func (x *JsonWebKeySetGeneratorRequest) GetKid() string {
	if x != nil {
		return x.Kid
	}
	return ""
}

func (x *JsonWebKeySetGeneratorRequest) GetUse() string {
	if x != nil {
		return x.Use
	}
	return ""
}

type DeleteJsonWebKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Set string `protobuf:"bytes,1,opt,name=set,proto3" json:"set,omitempty"`
	Kid string `protobuf:"bytes,2,opt,name=kid,proto3" json:"kid,omitempty"`
}

func (x *DeleteJsonWebKeyRequest) Reset() {
	*x = DeleteJsonWebKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidc_api_key_v1_key_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJsonWebKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJsonWebKeyRequest) ProtoMessage() {}

func (x *DeleteJsonWebKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oidc_api_key_v1_key_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJsonWebKeyRequest.ProtoReflect.Descriptor instead.
func (*DeleteJsonWebKeyRequest) Descriptor() ([]byte, []int) {
	return file_oidc_api_key_v1_key_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteJsonWebKeyRequest) GetSet() string {
	if x != nil {
		return x.Set
	}
	return ""
}

func (x *DeleteJsonWebKeyRequest) GetKid() string {
	if x != nil {
		return x.Kid
	}
	return ""
}

type GetJsonWebKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Set string `protobuf:"bytes,1,opt,name=set,proto3" json:"set,omitempty"`
	Kid string `protobuf:"bytes,2,opt,name=kid,proto3" json:"kid,omitempty"`
}

func (x *GetJsonWebKeyRequest) Reset() {
	*x = GetJsonWebKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidc_api_key_v1_key_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJsonWebKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJsonWebKeyRequest) ProtoMessage() {}

func (x *GetJsonWebKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oidc_api_key_v1_key_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJsonWebKeyRequest.ProtoReflect.Descriptor instead.
func (*GetJsonWebKeyRequest) Descriptor() ([]byte, []int) {
	return file_oidc_api_key_v1_key_proto_rawDescGZIP(), []int{6}
}

func (x *GetJsonWebKeyRequest) GetSet() string {
	if x != nil {
		return x.Set
	}
	return ""
}

func (x *GetJsonWebKeyRequest) GetKid() string {
	if x != nil {
		return x.Kid
	}
	return ""
}

type UpdateJsonWebKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Set string      `protobuf:"bytes,1,opt,name=set,proto3" json:"set,omitempty"`
	Kid string      `protobuf:"bytes,2,opt,name=kid,proto3" json:"kid,omitempty"`
	Key *JsonWebKey `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *UpdateJsonWebKeyRequest) Reset() {
	*x = UpdateJsonWebKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidc_api_key_v1_key_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJsonWebKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJsonWebKeyRequest) ProtoMessage() {}

func (x *UpdateJsonWebKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oidc_api_key_v1_key_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJsonWebKeyRequest.ProtoReflect.Descriptor instead.
func (*UpdateJsonWebKeyRequest) Descriptor() ([]byte, []int) {
	return file_oidc_api_key_v1_key_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateJsonWebKeyRequest) GetSet() string {
	if x != nil {
		return x.Set
	}
	return ""
}

func (x *UpdateJsonWebKeyRequest) GetKid() string {
	if x != nil {
		return x.Kid
	}
	return ""
}

func (x *UpdateJsonWebKeyRequest) GetKey() *JsonWebKey {
	if x != nil {
		return x.Key
	}
	return nil
}

type JsonWebKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alg string   `protobuf:"bytes,1,opt,name=alg,proto3" json:"alg,omitempty"`
	Crv *string  `protobuf:"bytes,2,opt,name=crv,proto3,oneof" json:"crv,omitempty"`
	D   *string  `protobuf:"bytes,3,opt,name=d,proto3,oneof" json:"d,omitempty"`
	Dp  *string  `protobuf:"bytes,4,opt,name=dp,proto3,oneof" json:"dp,omitempty"`
	Dq  *string  `protobuf:"bytes,5,opt,name=dq,proto3,oneof" json:"dq,omitempty"`
	E   *string  `protobuf:"bytes,6,opt,name=e,proto3,oneof" json:"e,omitempty"`
	K   *string  `protobuf:"bytes,7,opt,name=k,proto3,oneof" json:"k,omitempty"`
	Kid string   `protobuf:"bytes,8,opt,name=kid,proto3" json:"kid,omitempty"`
	Kty string   `protobuf:"bytes,9,opt,name=kty,proto3" json:"kty,omitempty"`
	N   *string  `protobuf:"bytes,10,opt,name=n,proto3,oneof" json:"n,omitempty"`
	P   *string  `protobuf:"bytes,11,opt,name=p,proto3,oneof" json:"p,omitempty"`
	Q   *string  `protobuf:"bytes,12,opt,name=q,proto3,oneof" json:"q,omitempty"`
	Qi  *string  `protobuf:"bytes,13,opt,name=qi,proto3,oneof" json:"qi,omitempty"`
	Use string   `protobuf:"bytes,14,opt,name=use,proto3" json:"use,omitempty"`
	X   *string  `protobuf:"bytes,15,opt,name=x,proto3,oneof" json:"x,omitempty"`
	X5C []string `protobuf:"bytes,16,rep,name=x5c,proto3" json:"x5c,omitempty"`
	Y   *string  `protobuf:"bytes,17,opt,name=y,proto3,oneof" json:"y,omitempty"`
}

func (x *JsonWebKey) Reset() {
	*x = JsonWebKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidc_api_key_v1_key_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonWebKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonWebKey) ProtoMessage() {}

func (x *JsonWebKey) ProtoReflect() protoreflect.Message {
	mi := &file_oidc_api_key_v1_key_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonWebKey.ProtoReflect.Descriptor instead.
func (*JsonWebKey) Descriptor() ([]byte, []int) {
	return file_oidc_api_key_v1_key_proto_rawDescGZIP(), []int{8}
}

func (x *JsonWebKey) GetAlg() string {
	if x != nil {
		return x.Alg
	}
	return ""
}

func (x *JsonWebKey) GetCrv() string {
	if x != nil && x.Crv != nil {
		return *x.Crv
	}
	return ""
}

func (x *JsonWebKey) GetD() string {
	if x != nil && x.D != nil {
		return *x.D
	}
	return ""
}

func (x *JsonWebKey) GetDp() string {
	if x != nil && x.Dp != nil {
		return *x.Dp
	}
	return ""
}

func (x *JsonWebKey) GetDq() string {
	if x != nil && x.Dq != nil {
		return *x.Dq
	}
	return ""
}

func (x *JsonWebKey) GetE() string {
	if x != nil && x.E != nil {
		return *x.E
	}
	return ""
}

func (x *JsonWebKey) GetK() string {
	if x != nil && x.K != nil {
		return *x.K
	}
	return ""
}

func (x *JsonWebKey) GetKid() string {
	if x != nil {
		return x.Kid
	}
	return ""
}

func (x *JsonWebKey) GetKty() string {
	if x != nil {
		return x.Kty
	}
	return ""
}

func (x *JsonWebKey) GetN() string {
	if x != nil && x.N != nil {
		return *x.N
	}
	return ""
}

func (x *JsonWebKey) GetP() string {
	if x != nil && x.P != nil {
		return *x.P
	}
	return ""
}

func (x *JsonWebKey) GetQ() string {
	if x != nil && x.Q != nil {
		return *x.Q
	}
	return ""
}

func (x *JsonWebKey) GetQi() string {
	if x != nil && x.Qi != nil {
		return *x.Qi
	}
	return ""
}

func (x *JsonWebKey) GetUse() string {
	if x != nil {
		return x.Use
	}
	return ""
}

func (x *JsonWebKey) GetX() string {
	if x != nil && x.X != nil {
		return *x.X
	}
	return ""
}

func (x *JsonWebKey) GetX5C() []string {
	if x != nil {
		return x.X5C
	}
	return nil
}

func (x *JsonWebKey) GetY() string {
	if x != nil && x.Y != nil {
		return *x.Y
	}
	return ""
}

type JsonWebKeySet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*JsonWebKey `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *JsonWebKeySet) Reset() {
	*x = JsonWebKeySet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidc_api_key_v1_key_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonWebKeySet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonWebKeySet) ProtoMessage() {}

func (x *JsonWebKeySet) ProtoReflect() protoreflect.Message {
	mi := &file_oidc_api_key_v1_key_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonWebKeySet.ProtoReflect.Descriptor instead.
func (*JsonWebKeySet) Descriptor() ([]byte, []int) {
	return file_oidc_api_key_v1_key_proto_rawDescGZIP(), []int{9}
}

func (x *JsonWebKeySet) GetKeys() []*JsonWebKey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_oidc_api_key_v1_key_proto protoreflect.FileDescriptor

var file_oidc_api_key_v1_key_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x79, 0x2f, 0x76,
	0x31, 0x2f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6f, 0x69, 0x64,
	0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x79, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x73, 0x6f,
	0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b,
	0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x73, 0x65, 0x74,
	0x22, 0x38, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65,
	0x79, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x03, 0x73,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x73, 0x65, 0x74, 0x22, 0x6c, 0x0a, 0x1a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x53, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x03, 0x73, 0x65, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6b, 0x65, 0x79, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x53,
	0x65, 0x74, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x7c, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0b, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x03, 0x73, 0x65, 0x74, 0x12, 0x3f, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b,
	0x65, 0x79, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x74,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x7c, 0x0a, 0x1d, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65,
	0x62, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x03, 0x61, 0x6c, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x03, 0x61, 0x6c, 0x67, 0x12, 0x1d, 0x0a, 0x03, 0x6b, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0b, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x03, 0x6b, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x03, 0x75, 0x73, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0b, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x03, 0x75, 0x73, 0x65, 0x22, 0x57, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x73,
	0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xe2, 0x41,
	0x01, 0x02, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x73, 0x65, 0x74, 0x12, 0x1d,
	0x0a, 0x03, 0x6b, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xe2, 0x41, 0x01,
	0x02, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x6b, 0x69, 0x64, 0x22, 0x54, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0b, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x03, 0x73, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x03, 0x6b, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0b, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03,
	0x6b, 0x69, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x73,
	0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xe2, 0x41,
	0x01, 0x02, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x73, 0x65, 0x74, 0x12, 0x1d,
	0x0a, 0x03, 0x6b, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xe2, 0x41, 0x01,
	0x02, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x6b, 0x69, 0x64, 0x12, 0x2a, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6f, 0x69, 0x64,
	0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x79, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65,
	0x62, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0xd5, 0x03, 0x0a, 0x0a, 0x4a, 0x73,
	0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x03, 0x61, 0x6c, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x03, 0x61, 0x6c, 0x67, 0x12, 0x15, 0x0a, 0x03, 0x63, 0x72, 0x76, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x63, 0x72, 0x76, 0x88, 0x01, 0x01, 0x12, 0x11,
	0x0a, 0x01, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x01, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x13, 0x0a, 0x02, 0x64, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52,
	0x02, 0x64, 0x70, 0x88, 0x01, 0x01, 0x12, 0x13, 0x0a, 0x02, 0x64, 0x71, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x03, 0x52, 0x02, 0x64, 0x71, 0x88, 0x01, 0x01, 0x12, 0x11, 0x0a, 0x01, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x01, 0x65, 0x88, 0x01, 0x01, 0x12, 0x11,
	0x0a, 0x01, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x01, 0x6b, 0x88, 0x01,
	0x01, 0x12, 0x1d, 0x0a, 0x03, 0x6b, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b,
	0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x6b, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x03, 0x6b, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xe2,
	0x41, 0x01, 0x02, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x6b, 0x74, 0x79, 0x12,
	0x11, 0x0a, 0x01, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x01, 0x6e, 0x88,
	0x01, 0x01, 0x12, 0x11, 0x0a, 0x01, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52,
	0x01, 0x70, 0x88, 0x01, 0x01, 0x12, 0x11, 0x0a, 0x01, 0x71, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x08, 0x52, 0x01, 0x71, 0x88, 0x01, 0x01, 0x12, 0x13, 0x0a, 0x02, 0x71, 0x69, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x02, 0x71, 0x69, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a,
	0x03, 0x75, 0x73, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xe2, 0x41, 0x01, 0x02,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x75, 0x73, 0x65, 0x12, 0x11, 0x0a, 0x01,
	0x78, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0a, 0x52, 0x01, 0x78, 0x88, 0x01, 0x01, 0x12,
	0x10, 0x0a, 0x03, 0x78, 0x35, 0x63, 0x18, 0x10, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x78, 0x35,
	0x63, 0x12, 0x11, 0x0a, 0x01, 0x79, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0b, 0x52, 0x01,
	0x79, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x63, 0x72, 0x76, 0x42, 0x04, 0x0a, 0x02,
	0x5f, 0x64, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x64, 0x70, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x64, 0x71,
	0x42, 0x04, 0x0a, 0x02, 0x5f, 0x65, 0x42, 0x04, 0x0a, 0x02, 0x5f, 0x6b, 0x42, 0x04, 0x0a, 0x02,
	0x5f, 0x6e, 0x42, 0x04, 0x0a, 0x02, 0x5f, 0x70, 0x42, 0x04, 0x0a, 0x02, 0x5f, 0x71, 0x42, 0x05,
	0x0a, 0x03, 0x5f, 0x71, 0x69, 0x42, 0x04, 0x0a, 0x02, 0x5f, 0x78, 0x42, 0x04, 0x0a, 0x02, 0x5f,
	0x79, 0x22, 0x3d, 0x0a, 0x0d, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x53,
	0x65, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x79, 0x2e,
	0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73,
	0x32, 0xd9, 0x06, 0x0a, 0x0a, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x74, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62,
	0x4b, 0x65, 0x79, 0x53, 0x65, 0x74, 0x12, 0x28, 0x2e, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6b, 0x65, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x73, 0x6f, 0x6e,
	0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x2a, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x2f,
	0x7b, 0x73, 0x65, 0x74, 0x7d, 0x12, 0x73, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4a, 0x73, 0x6f, 0x6e,
	0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x74, 0x12, 0x25, 0x2e, 0x6f, 0x69, 0x64, 0x63,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x73, 0x6f, 0x6e,
	0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x79, 0x2e,
	0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x74, 0x22, 0x1b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x69, 0x64, 0x63, 0x2f,
	0x6b, 0x65, 0x79, 0x73, 0x2f, 0x7b, 0x73, 0x65, 0x74, 0x7d, 0x12, 0x7c, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x53, 0x65,
	0x74, 0x12, 0x28, 0x2e, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x79,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65,
	0x79, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6f, 0x69,
	0x64, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x79, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x57,
	0x65, 0x62, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x74, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18,
	0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x6b, 0x65,
	0x79, 0x73, 0x2f, 0x7b, 0x73, 0x65, 0x74, 0x7d, 0x12, 0x7c, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x74, 0x12,
	0x28, 0x2e, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x79, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x53,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6f, 0x69, 0x64, 0x63,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x79, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62,
	0x4b, 0x65, 0x79, 0x53, 0x65, 0x74, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01,
	0x2a, 0x1a, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x6b, 0x65, 0x79, 0x73,
	0x2f, 0x7b, 0x73, 0x65, 0x74, 0x7d, 0x12, 0x74, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x25, 0x2e, 0x6f, 0x69, 0x64,
	0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1b, 0x2a, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x6b, 0x65, 0x79, 0x73,
	0x2f, 0x7b, 0x73, 0x65, 0x74, 0x7d, 0x2f, 0x7b, 0x6b, 0x69, 0x64, 0x7d, 0x12, 0x73, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x22, 0x2e,
	0x6f, 0x69, 0x64, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x79, 0x2e, 0x47, 0x65, 0x74,
	0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x79,
	0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x74, 0x22, 0x21,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x69, 0x64, 0x63,
	0x2f, 0x6b, 0x65, 0x79, 0x73, 0x2f, 0x7b, 0x73, 0x65, 0x74, 0x7d, 0x2f, 0x7b, 0x6b, 0x69, 0x64,
	0x7d, 0x12, 0x79, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x57,
	0x65, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x25, 0x2e, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6b, 0x65, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x57,
	0x65, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6f,
	0x69, 0x64, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x79, 0x2e, 0x4a, 0x73, 0x6f, 0x6e,
	0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01,
	0x2a, 0x1a, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x6b, 0x65, 0x79, 0x73,
	0x2f, 0x7b, 0x73, 0x65, 0x74, 0x7d, 0x2f, 0x7b, 0x6b, 0x69, 0x64, 0x7d, 0x42, 0x2b, 0x5a, 0x29,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x61,
	0x61, 0x73, 0x2f, 0x6b, 0x69, 0x74, 0x2f, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6b, 0x65, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_oidc_api_key_v1_key_proto_rawDescOnce sync.Once
	file_oidc_api_key_v1_key_proto_rawDescData = file_oidc_api_key_v1_key_proto_rawDesc
)

func file_oidc_api_key_v1_key_proto_rawDescGZIP() []byte {
	file_oidc_api_key_v1_key_proto_rawDescOnce.Do(func() {
		file_oidc_api_key_v1_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_oidc_api_key_v1_key_proto_rawDescData)
	})
	return file_oidc_api_key_v1_key_proto_rawDescData
}

var file_oidc_api_key_v1_key_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_oidc_api_key_v1_key_proto_goTypes = []interface{}{
	(*DeleteJsonWebKeySetRequest)(nil),    // 0: oidc.api.key.DeleteJsonWebKeySetRequest
	(*GetJsonWebKeySetRequest)(nil),       // 1: oidc.api.key.GetJsonWebKeySetRequest
	(*UpdateJsonWebKeySetRequest)(nil),    // 2: oidc.api.key.UpdateJsonWebKeySetRequest
	(*CreateJsonWebKeySetRequest)(nil),    // 3: oidc.api.key.CreateJsonWebKeySetRequest
	(*JsonWebKeySetGeneratorRequest)(nil), // 4: oidc.api.key.JsonWebKeySetGeneratorRequest
	(*DeleteJsonWebKeyRequest)(nil),       // 5: oidc.api.key.DeleteJsonWebKeyRequest
	(*GetJsonWebKeyRequest)(nil),          // 6: oidc.api.key.GetJsonWebKeyRequest
	(*UpdateJsonWebKeyRequest)(nil),       // 7: oidc.api.key.UpdateJsonWebKeyRequest
	(*JsonWebKey)(nil),                    // 8: oidc.api.key.JsonWebKey
	(*JsonWebKeySet)(nil),                 // 9: oidc.api.key.JsonWebKeySet
	(*emptypb.Empty)(nil),                 // 10: google.protobuf.Empty
}
var file_oidc_api_key_v1_key_proto_depIdxs = []int32{
	9,  // 0: oidc.api.key.UpdateJsonWebKeySetRequest.keys:type_name -> oidc.api.key.JsonWebKeySet
	4,  // 1: oidc.api.key.CreateJsonWebKeySetRequest.keys:type_name -> oidc.api.key.JsonWebKeySetGeneratorRequest
	8,  // 2: oidc.api.key.UpdateJsonWebKeyRequest.key:type_name -> oidc.api.key.JsonWebKey
	8,  // 3: oidc.api.key.JsonWebKeySet.keys:type_name -> oidc.api.key.JsonWebKey
	0,  // 4: oidc.api.key.KeyService.DeleteJsonWebKeySet:input_type -> oidc.api.key.DeleteJsonWebKeySetRequest
	1,  // 5: oidc.api.key.KeyService.GetJsonWebKeySet:input_type -> oidc.api.key.GetJsonWebKeySetRequest
	3,  // 6: oidc.api.key.KeyService.CreateJsonWebKeySet:input_type -> oidc.api.key.CreateJsonWebKeySetRequest
	2,  // 7: oidc.api.key.KeyService.UpdateJsonWebKeySet:input_type -> oidc.api.key.UpdateJsonWebKeySetRequest
	5,  // 8: oidc.api.key.KeyService.DeleteJsonWebKey:input_type -> oidc.api.key.DeleteJsonWebKeyRequest
	6,  // 9: oidc.api.key.KeyService.GetJsonWebKey:input_type -> oidc.api.key.GetJsonWebKeyRequest
	7,  // 10: oidc.api.key.KeyService.UpdateJsonWebKey:input_type -> oidc.api.key.UpdateJsonWebKeyRequest
	10, // 11: oidc.api.key.KeyService.DeleteJsonWebKeySet:output_type -> google.protobuf.Empty
	9,  // 12: oidc.api.key.KeyService.GetJsonWebKeySet:output_type -> oidc.api.key.JsonWebKeySet
	9,  // 13: oidc.api.key.KeyService.CreateJsonWebKeySet:output_type -> oidc.api.key.JsonWebKeySet
	9,  // 14: oidc.api.key.KeyService.UpdateJsonWebKeySet:output_type -> oidc.api.key.JsonWebKeySet
	10, // 15: oidc.api.key.KeyService.DeleteJsonWebKey:output_type -> google.protobuf.Empty
	9,  // 16: oidc.api.key.KeyService.GetJsonWebKey:output_type -> oidc.api.key.JsonWebKeySet
	8,  // 17: oidc.api.key.KeyService.UpdateJsonWebKey:output_type -> oidc.api.key.JsonWebKey
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_oidc_api_key_v1_key_proto_init() }
func file_oidc_api_key_v1_key_proto_init() {
	if File_oidc_api_key_v1_key_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oidc_api_key_v1_key_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteJsonWebKeySetRequest); i {
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
		file_oidc_api_key_v1_key_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJsonWebKeySetRequest); i {
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
		file_oidc_api_key_v1_key_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateJsonWebKeySetRequest); i {
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
		file_oidc_api_key_v1_key_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJsonWebKeySetRequest); i {
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
		file_oidc_api_key_v1_key_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsonWebKeySetGeneratorRequest); i {
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
		file_oidc_api_key_v1_key_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteJsonWebKeyRequest); i {
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
		file_oidc_api_key_v1_key_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJsonWebKeyRequest); i {
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
		file_oidc_api_key_v1_key_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateJsonWebKeyRequest); i {
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
		file_oidc_api_key_v1_key_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsonWebKey); i {
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
		file_oidc_api_key_v1_key_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsonWebKeySet); i {
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
	file_oidc_api_key_v1_key_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_oidc_api_key_v1_key_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_oidc_api_key_v1_key_proto_goTypes,
		DependencyIndexes: file_oidc_api_key_v1_key_proto_depIdxs,
		MessageInfos:      file_oidc_api_key_v1_key_proto_msgTypes,
	}.Build()
	File_oidc_api_key_v1_key_proto = out.File
	file_oidc_api_key_v1_key_proto_rawDesc = nil
	file_oidc_api_key_v1_key_proto_goTypes = nil
	file_oidc_api_key_v1_key_proto_depIdxs = nil
}
