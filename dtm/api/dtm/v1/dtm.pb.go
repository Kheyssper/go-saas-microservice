// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: dtm/api/dtm/v1/dtm.proto

package v1

import (
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

type QueryPreparedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *QueryPreparedRequest) Reset() {
	*x = QueryPreparedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dtm_api_dtm_v1_dtm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPreparedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPreparedRequest) ProtoMessage() {}

func (x *QueryPreparedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dtm_api_dtm_v1_dtm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPreparedRequest.ProtoReflect.Descriptor instead.
func (*QueryPreparedRequest) Descriptor() ([]byte, []int) {
	return file_dtm_api_dtm_v1_dtm_proto_rawDescGZIP(), []int{0}
}

func (x *QueryPreparedRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

var File_dtm_api_dtm_v1_dtm_proto protoreflect.FileDescriptor

var file_dtm_api_dtm_v1_dtm_proto_rawDesc = []byte{
	0x0a, 0x18, 0x64, 0x74, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x74, 0x6d, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x74, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x64, 0x74, 0x6d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x64, 0x74, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x14, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x32, 0x85, 0x01, 0x0a,
	0x0a, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x77, 0x0a, 0x0d, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x64, 0x12, 0x24, 0x2e, 0x64,
	0x74, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x74, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x22, 0x12, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x7d, 0x2f, 0x64, 0x74, 0x6d, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2d, 0x70, 0x72, 0x65, 0x70,
	0x61, 0x72, 0x65, 0x64, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x6b, 0x69, 0x74, 0x2f, 0x64,
	0x74, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x74, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dtm_api_dtm_v1_dtm_proto_rawDescOnce sync.Once
	file_dtm_api_dtm_v1_dtm_proto_rawDescData = file_dtm_api_dtm_v1_dtm_proto_rawDesc
)

func file_dtm_api_dtm_v1_dtm_proto_rawDescGZIP() []byte {
	file_dtm_api_dtm_v1_dtm_proto_rawDescOnce.Do(func() {
		file_dtm_api_dtm_v1_dtm_proto_rawDescData = protoimpl.X.CompressGZIP(file_dtm_api_dtm_v1_dtm_proto_rawDescData)
	})
	return file_dtm_api_dtm_v1_dtm_proto_rawDescData
}

var file_dtm_api_dtm_v1_dtm_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_dtm_api_dtm_v1_dtm_proto_goTypes = []interface{}{
	(*QueryPreparedRequest)(nil), // 0: dtm.api.dtm.v1.QueryPreparedRequest
	(*emptypb.Empty)(nil),        // 1: google.protobuf.Empty
}
var file_dtm_api_dtm_v1_dtm_proto_depIdxs = []int32{
	0, // 0: dtm.api.dtm.v1.MsgService.QueryPrepared:input_type -> dtm.api.dtm.v1.QueryPreparedRequest
	1, // 1: dtm.api.dtm.v1.MsgService.QueryPrepared:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dtm_api_dtm_v1_dtm_proto_init() }
func file_dtm_api_dtm_v1_dtm_proto_init() {
	if File_dtm_api_dtm_v1_dtm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dtm_api_dtm_v1_dtm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPreparedRequest); i {
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
			RawDescriptor: file_dtm_api_dtm_v1_dtm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dtm_api_dtm_v1_dtm_proto_goTypes,
		DependencyIndexes: file_dtm_api_dtm_v1_dtm_proto_depIdxs,
		MessageInfos:      file_dtm_api_dtm_v1_dtm_proto_msgTypes,
	}.Build()
	File_dtm_api_dtm_v1_dtm_proto = out.File
	file_dtm_api_dtm_v1_dtm_proto_rawDesc = nil
	file_dtm_api_dtm_v1_dtm_proto_goTypes = nil
	file_dtm_api_dtm_v1_dtm_proto_depIdxs = nil
}
