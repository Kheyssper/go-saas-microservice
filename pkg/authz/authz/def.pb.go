// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: authz/authz/def.proto

package authz

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PermissionAllowSide int32

const (
	PermissionAllowSide_BOTH        PermissionAllowSide = 0
	PermissionAllowSide_HOST_ONLY   PermissionAllowSide = 1
	PermissionAllowSide_TENANT_ONLY PermissionAllowSide = 2
)

// Enum value maps for PermissionAllowSide.
var (
	PermissionAllowSide_name = map[int32]string{
		0: "BOTH",
		1: "HOST_ONLY",
		2: "TENANT_ONLY",
	}
	PermissionAllowSide_value = map[string]int32{
		"BOTH":        0,
		"HOST_ONLY":   1,
		"TENANT_ONLY": 2,
	}
)

func (x PermissionAllowSide) Enum() *PermissionAllowSide {
	p := new(PermissionAllowSide)
	*p = x
	return p
}

func (x PermissionAllowSide) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PermissionAllowSide) Descriptor() protoreflect.EnumDescriptor {
	return file_authz_authz_def_proto_enumTypes[0].Descriptor()
}

func (PermissionAllowSide) Type() protoreflect.EnumType {
	return &file_authz_authz_def_proto_enumTypes[0]
}

func (x PermissionAllowSide) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PermissionAllowSide.Descriptor instead.
func (PermissionAllowSide) EnumDescriptor() ([]byte, []int) {
	return file_authz_authz_def_proto_rawDescGZIP(), []int{0}
}

// PermissionDefGroup group multiple permission definition
type PermissionDefGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Side     PermissionAllowSide `protobuf:"varint,2,opt,name=side,proto3,enum=authz.PermissionAllowSide" json:"side,omitempty"`
	Def      []*PermissionDef    `protobuf:"bytes,3,rep,name=def,proto3" json:"def,omitempty"`
	Extra    *structpb.Struct    `protobuf:"bytes,4,opt,name=extra,proto3" json:"extra,omitempty"`
	Internal bool                `protobuf:"varint,5,opt,name=internal,proto3" json:"internal,omitempty"`
	Priority int32               `protobuf:"varint,6,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *PermissionDefGroup) Reset() {
	*x = PermissionDefGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_authz_def_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionDefGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionDefGroup) ProtoMessage() {}

func (x *PermissionDefGroup) ProtoReflect() protoreflect.Message {
	mi := &file_authz_authz_def_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionDefGroup.ProtoReflect.Descriptor instead.
func (*PermissionDefGroup) Descriptor() ([]byte, []int) {
	return file_authz_authz_def_proto_rawDescGZIP(), []int{0}
}

func (x *PermissionDefGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PermissionDefGroup) GetSide() PermissionAllowSide {
	if x != nil {
		return x.Side
	}
	return PermissionAllowSide_BOTH
}

func (x *PermissionDefGroup) GetDef() []*PermissionDef {
	if x != nil {
		return x.Def
	}
	return nil
}

func (x *PermissionDefGroup) GetExtra() *structpb.Struct {
	if x != nil {
		return x.Extra
	}
	return nil
}

func (x *PermissionDefGroup) GetInternal() bool {
	if x != nil {
		return x.Internal
	}
	return false
}

func (x *PermissionDefGroup) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

type PermissionDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name user friendly name
	Name      string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Side      PermissionAllowSide `protobuf:"varint,2,opt,name=side,proto3,enum=authz.PermissionAllowSide" json:"side,omitempty"`
	Namespace string              `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Action    string              `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`
	Extra     *structpb.Struct    `protobuf:"bytes,5,opt,name=extra,proto3" json:"extra,omitempty"`
	// internal will not be displayed by ui
	Internal bool  `protobuf:"varint,6,opt,name=internal,proto3" json:"internal,omitempty"`
	Priority int32 `protobuf:"varint,7,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *PermissionDef) Reset() {
	*x = PermissionDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_authz_def_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionDef) ProtoMessage() {}

func (x *PermissionDef) ProtoReflect() protoreflect.Message {
	mi := &file_authz_authz_def_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionDef.ProtoReflect.Descriptor instead.
func (*PermissionDef) Descriptor() ([]byte, []int) {
	return file_authz_authz_def_proto_rawDescGZIP(), []int{1}
}

func (x *PermissionDef) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PermissionDef) GetSide() PermissionAllowSide {
	if x != nil {
		return x.Side
	}
	return PermissionAllowSide_BOTH
}

func (x *PermissionDef) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *PermissionDef) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *PermissionDef) GetExtra() *structpb.Struct {
	if x != nil {
		return x.Extra
	}
	return nil
}

func (x *PermissionDef) GetInternal() bool {
	if x != nil {
		return x.Internal
	}
	return false
}

func (x *PermissionDef) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

type PermissionConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups []*PermissionDefGroup `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *PermissionConf) Reset() {
	*x = PermissionConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_authz_def_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionConf) ProtoMessage() {}

func (x *PermissionConf) ProtoReflect() protoreflect.Message {
	mi := &file_authz_authz_def_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionConf.ProtoReflect.Descriptor instead.
func (*PermissionConf) Descriptor() ([]byte, []int) {
	return file_authz_authz_def_proto_rawDescGZIP(), []int{2}
}

func (x *PermissionConf) GetGroups() []*PermissionDefGroup {
	if x != nil {
		return x.Groups
	}
	return nil
}

var File_authz_authz_def_proto protoreflect.FileDescriptor

var file_authz_authz_def_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x64, 0x65,
	0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x01, 0x0a,
	0x12, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x66, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x73, 0x69, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x69, 0x64,
	0x65, 0x52, 0x04, 0x73, 0x69, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x03, 0x64, 0x65, 0x66, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x66, 0x52, 0x03, 0x64, 0x65, 0x66, 0x12,
	0x2d, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0xf0, 0x01, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x04,
	0x73, 0x69, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x6c,
	0x6f, 0x77, 0x53, 0x69, 0x64, 0x65, 0x52, 0x04, 0x73, 0x69, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x43, 0x0a, 0x0e, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x31, 0x0a, 0x06, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x66, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2a, 0x3f,
	0x0a, 0x13, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x6c, 0x6f,
	0x77, 0x53, 0x69, 0x64, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4f, 0x54, 0x48, 0x10, 0x00, 0x12,
	0x0d, 0x0a, 0x09, 0x48, 0x4f, 0x53, 0x54, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x01, 0x12, 0x0f,
	0x0a, 0x0b, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x02, 0x42,
	0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authz_authz_def_proto_rawDescOnce sync.Once
	file_authz_authz_def_proto_rawDescData = file_authz_authz_def_proto_rawDesc
)

func file_authz_authz_def_proto_rawDescGZIP() []byte {
	file_authz_authz_def_proto_rawDescOnce.Do(func() {
		file_authz_authz_def_proto_rawDescData = protoimpl.X.CompressGZIP(file_authz_authz_def_proto_rawDescData)
	})
	return file_authz_authz_def_proto_rawDescData
}

var file_authz_authz_def_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_authz_authz_def_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_authz_authz_def_proto_goTypes = []interface{}{
	(PermissionAllowSide)(0),   // 0: authz.PermissionAllowSide
	(*PermissionDefGroup)(nil), // 1: authz.PermissionDefGroup
	(*PermissionDef)(nil),      // 2: authz.PermissionDef
	(*PermissionConf)(nil),     // 3: authz.PermissionConf
	(*structpb.Struct)(nil),    // 4: google.protobuf.Struct
}
var file_authz_authz_def_proto_depIdxs = []int32{
	0, // 0: authz.PermissionDefGroup.side:type_name -> authz.PermissionAllowSide
	2, // 1: authz.PermissionDefGroup.def:type_name -> authz.PermissionDef
	4, // 2: authz.PermissionDefGroup.extra:type_name -> google.protobuf.Struct
	0, // 3: authz.PermissionDef.side:type_name -> authz.PermissionAllowSide
	4, // 4: authz.PermissionDef.extra:type_name -> google.protobuf.Struct
	1, // 5: authz.PermissionConf.groups:type_name -> authz.PermissionDefGroup
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_authz_authz_def_proto_init() }
func file_authz_authz_def_proto_init() {
	if File_authz_authz_def_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authz_authz_def_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionDefGroup); i {
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
		file_authz_authz_def_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionDef); i {
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
		file_authz_authz_def_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionConf); i {
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
			RawDescriptor: file_authz_authz_def_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_authz_authz_def_proto_goTypes,
		DependencyIndexes: file_authz_authz_def_proto_depIdxs,
		EnumInfos:         file_authz_authz_def_proto_enumTypes,
		MessageInfos:      file_authz_authz_def_proto_msgTypes,
	}.Build()
	File_authz_authz_def_proto = out.File
	file_authz_authz_def_proto_rawDesc = nil
	file_authz_authz_def_proto_goTypes = nil
	file_authz_authz_def_proto_depIdxs = nil
}
