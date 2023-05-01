// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: user/api/auth/v1/error_reason.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ErrorReason int32

const (
	ErrorReason_INVALID_CREDENTIALS       ErrorReason = 0
	ErrorReason_INVALID_OPERATION         ErrorReason = 1
	ErrorReason_USER_LOCKED               ErrorReason = 2
	ErrorReason_EMAIL_NOT_CONFIRMED       ErrorReason = 3
	ErrorReason_PHONE_NOT_CONFIRMED       ErrorReason = 4
	ErrorReason_EMAIL_RECOVER_FAILED      ErrorReason = 5
	ErrorReason_EMAIL_CONFIRM_FAILED      ErrorReason = 6
	ErrorReason_PHONE_RECOVER_FAILED      ErrorReason = 7
	ErrorReason_PHONE_CONFIRM_FAILED      ErrorReason = 8
	ErrorReason_TWO_STEP_FAILED           ErrorReason = 9
	ErrorReason_CONFIRM_PASSWORD_MISMATCH ErrorReason = 10
	ErrorReason_REMEMBER_TOKEN_NOT_FOUND  ErrorReason = 11
	ErrorReason_REMEMBER_TOKEN_USED       ErrorReason = 14
	ErrorReason_USER_DELETED              ErrorReason = 12
	ErrorReason_REFRESH_TOKEN_INVALID     ErrorReason = 13
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0:  "INVALID_CREDENTIALS",
		1:  "INVALID_OPERATION",
		2:  "USER_LOCKED",
		3:  "EMAIL_NOT_CONFIRMED",
		4:  "PHONE_NOT_CONFIRMED",
		5:  "EMAIL_RECOVER_FAILED",
		6:  "EMAIL_CONFIRM_FAILED",
		7:  "PHONE_RECOVER_FAILED",
		8:  "PHONE_CONFIRM_FAILED",
		9:  "TWO_STEP_FAILED",
		10: "CONFIRM_PASSWORD_MISMATCH",
		11: "REMEMBER_TOKEN_NOT_FOUND",
		14: "REMEMBER_TOKEN_USED",
		12: "USER_DELETED",
		13: "REFRESH_TOKEN_INVALID",
	}
	ErrorReason_value = map[string]int32{
		"INVALID_CREDENTIALS":       0,
		"INVALID_OPERATION":         1,
		"USER_LOCKED":               2,
		"EMAIL_NOT_CONFIRMED":       3,
		"PHONE_NOT_CONFIRMED":       4,
		"EMAIL_RECOVER_FAILED":      5,
		"EMAIL_CONFIRM_FAILED":      6,
		"PHONE_RECOVER_FAILED":      7,
		"PHONE_CONFIRM_FAILED":      8,
		"TWO_STEP_FAILED":           9,
		"CONFIRM_PASSWORD_MISMATCH": 10,
		"REMEMBER_TOKEN_NOT_FOUND":  11,
		"REMEMBER_TOKEN_USED":       14,
		"USER_DELETED":              12,
		"REFRESH_TOKEN_INVALID":     13,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_user_api_auth_v1_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_user_api_auth_v1_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_user_api_auth_v1_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_user_api_auth_v1_error_reason_proto protoreflect.FileDescriptor

var file_user_api_auth_v1_error_reason_proto_rawDesc = []byte{
	0x0a, 0x23, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xe0, 0x03, 0x0a,
	0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x13,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x52, 0x45, 0x44, 0x45, 0x4e, 0x54, 0x49,
	0x41, 0x4c, 0x53, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1b, 0x0a, 0x11, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x15, 0x0a, 0x0b, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x93, 0x03, 0x12,
	0x1d, 0x0a, 0x13, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x4f, 0x4e,
	0x46, 0x49, 0x52, 0x4d, 0x45, 0x44, 0x10, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x93, 0x03, 0x12, 0x1d,
	0x0a, 0x13, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x46,
	0x49, 0x52, 0x4d, 0x45, 0x44, 0x10, 0x04, 0x1a, 0x04, 0xa8, 0x45, 0x93, 0x03, 0x12, 0x1e, 0x0a,
	0x14, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x1a, 0x04, 0xa8, 0x45, 0x93, 0x03, 0x12, 0x1e, 0x0a,
	0x14, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x06, 0x1a, 0x04, 0xa8, 0x45, 0x93, 0x03, 0x12, 0x1e, 0x0a,
	0x14, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x07, 0x1a, 0x04, 0xa8, 0x45, 0x93, 0x03, 0x12, 0x1e, 0x0a,
	0x14, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x08, 0x1a, 0x04, 0xa8, 0x45, 0x93, 0x03, 0x12, 0x19, 0x0a,
	0x0f, 0x54, 0x57, 0x4f, 0x5f, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x10, 0x09, 0x1a, 0x04, 0xa8, 0x45, 0x93, 0x03, 0x12, 0x23, 0x0a, 0x19, 0x43, 0x4f, 0x4e, 0x46,
	0x49, 0x52, 0x4d, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x4d, 0x49, 0x53,
	0x4d, 0x41, 0x54, 0x43, 0x48, 0x10, 0x0a, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x22, 0x0a,
	0x18, 0x52, 0x45, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x0b, 0x1a, 0x04, 0xa8, 0x45, 0x93,
	0x03, 0x12, 0x1d, 0x0a, 0x13, 0x52, 0x45, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x54, 0x4f,
	0x4b, 0x45, 0x4e, 0x5f, 0x55, 0x53, 0x45, 0x44, 0x10, 0x0e, 0x1a, 0x04, 0xa8, 0x45, 0x93, 0x03,
	0x12, 0x16, 0x0a, 0x0c, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44,
	0x10, 0x0c, 0x1a, 0x04, 0xa8, 0x45, 0x93, 0x03, 0x12, 0x1f, 0x0a, 0x15, 0x52, 0x45, 0x46, 0x52,
	0x45, 0x53, 0x48, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x10, 0x0d, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42,
	0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x6b, 0x69, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_api_auth_v1_error_reason_proto_rawDescOnce sync.Once
	file_user_api_auth_v1_error_reason_proto_rawDescData = file_user_api_auth_v1_error_reason_proto_rawDesc
)

func file_user_api_auth_v1_error_reason_proto_rawDescGZIP() []byte {
	file_user_api_auth_v1_error_reason_proto_rawDescOnce.Do(func() {
		file_user_api_auth_v1_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_api_auth_v1_error_reason_proto_rawDescData)
	})
	return file_user_api_auth_v1_error_reason_proto_rawDescData
}

var file_user_api_auth_v1_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_api_auth_v1_error_reason_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: user.api.auth.v1.ErrorReason
}
var file_user_api_auth_v1_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_api_auth_v1_error_reason_proto_init() }
func file_user_api_auth_v1_error_reason_proto_init() {
	if File_user_api_auth_v1_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_api_auth_v1_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_api_auth_v1_error_reason_proto_goTypes,
		DependencyIndexes: file_user_api_auth_v1_error_reason_proto_depIdxs,
		EnumInfos:         file_user_api_auth_v1_error_reason_proto_enumTypes,
	}.Build()
	File_user_api_auth_v1_error_reason_proto = out.File
	file_user_api_auth_v1_error_reason_proto_rawDesc = nil
	file_user_api_auth_v1_error_reason_proto_goTypes = nil
	file_user_api_auth_v1_error_reason_proto_depIdxs = nil
}
