// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.5.0
// source: service.proto

package messages

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x0f, 0x61, 0x65, 0x72, 0x6f, 0x73,
	0x70, 0x69, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd1, 0x06, 0x0a,
	0x18, 0x41, 0x65, 0x72, 0x6f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x41, 0x50, 0x49, 0x12, 0x66, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x28, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x25, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x60, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x26, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x66, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x28, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x28, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x64, 0x0a, 0x13, 0x49, 0x73, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x49, 0x73, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x49, 0x73, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x73, 0x0a, 0x18, 0x4d, 0x61, 0x70, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x4d, 0x61, 0x70, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2a, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a,
	0x12, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4d,
	0x61, 0x70, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x32, 0xf9, 0x04, 0x0a, 0x17, 0x41, 0x65, 0x72, 0x6f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x65, 0x72, 0x6f, 0x73, 0x70, 0x69, 0x6b, 0x65, 0x41, 0x50, 0x49, 0x12, 0x64, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x27, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x65, 0x72, 0x6f, 0x73, 0x70, 0x69, 0x6b, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x65, 0x72, 0x6f, 0x73, 0x70, 0x69, 0x6b,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x5b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x24, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x65, 0x72, 0x6f, 0x73, 0x70, 0x69, 0x6b, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x65, 0x72, 0x6f, 0x73, 0x70, 0x69, 0x6b, 0x65, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x69, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x30,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x65, 0x72,
	0x6f, 0x73, 0x70, 0x69, 0x6b, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x42, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x65, 0x72, 0x6f, 0x73, 0x70, 0x69, 0x6b, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x28, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x65,
	0x72, 0x6f, 0x73, 0x70, 0x69, 0x6b, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x65, 0x72, 0x6f, 0x73, 0x70, 0x69, 0x6b, 0x65, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x64, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x27, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x65, 0x72, 0x6f, 0x73, 0x70, 0x69, 0x6b, 0x65, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x65, 0x72, 0x6f,
	0x73, 0x70, 0x69, 0x6b, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x65, 0x72, 0x6f, 0x73, 0x70, 0x69,
	0x6b, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x65, 0x72, 0x6f, 0x73, 0x70, 0x69, 0x6b, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x65, 0x72, 0x6f, 0x73,
	0x70, 0x69, 0x6b, 0x65, 0x2f, 0x61, 0x65, 0x72, 0x6f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x63, 0x61, 0x70, 0x69, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_proto_goTypes = []interface{}{
	(*CreateKubernetesClusterRequest)(nil),         // 0: messages.CreateKubernetesClusterRequest
	(*GetKubernetesClusterRequest)(nil),            // 1: messages.GetKubernetesClusterRequest
	(*GetKubernetesClustersRequest)(nil),           // 2: messages.GetKubernetesClustersRequest
	(*UpdateKubernetesClusterRequest)(nil),         // 3: messages.UpdateKubernetesClusterRequest
	(*DeleteKubernetesClusterRequest)(nil),         // 4: messages.DeleteKubernetesClusterRequest
	(*IsKubernetesClusterRequest)(nil),             // 5: messages.IsKubernetesClusterRequest
	(*MapRegionToSharedClusterRequest)(nil),        // 6: messages.MapRegionToSharedClusterRequest
	(*MapRegionToClusterRequest)(nil),              // 7: messages.MapRegionToClusterRequest
	(*CreateAerospikeClusterRequest)(nil),          // 8: messages.CreateAerospikeClusterRequest
	(*GetAerospikeClusterRequest)(nil),             // 9: messages.GetAerospikeClusterRequest
	(*GetAerospikeClustersByNamespaceRequest)(nil), // 10: messages.GetAerospikeClustersByNamespaceRequest
	(*GetAllAerospikeClustersRequest)(nil),         // 11: messages.GetAllAerospikeClustersRequest
	(*UpdateAerospikeClusterRequest)(nil),          // 12: messages.UpdateAerospikeClusterRequest
	(*DeleteAerospikeClusterRequest)(nil),          // 13: messages.DeleteAerospikeClusterRequest
	(*CreateKubernetesClusterResponse)(nil),        // 14: messages.CreateKubernetesClusterResponse
	(*GetKubernetesClusterResponse)(nil),           // 15: messages.GetKubernetesClusterResponse
	(*GetKubernetesClustersResponse)(nil),          // 16: messages.GetKubernetesClustersResponse
	(*UpdateKubernetesClusterResponse)(nil),        // 17: messages.UpdateKubernetesClusterResponse
	(*DeleteKubernetesClusterResponse)(nil),        // 18: messages.DeleteKubernetesClusterResponse
	(*IsKubernetesClusterResponse)(nil),            // 19: messages.IsKubernetesClusterResponse
	(*MapRegionToSharedClusterResponse)(nil),       // 20: messages.MapRegionToSharedClusterResponse
	(*MapRegionToClusterResponse)(nil),             // 21: messages.MapRegionToClusterResponse
	(*CreateAerospikeClusterResponse)(nil),         // 22: messages.CreateAerospikeClusterResponse
	(*GetAerospikeClusterResponse)(nil),            // 23: messages.GetAerospikeClusterResponse
	(*GetAerospikeClustersResponse)(nil),           // 24: messages.GetAerospikeClustersResponse
	(*UpdateAerospikeClusterResponse)(nil),         // 25: messages.UpdateAerospikeClusterResponse
	(*DeleteAerospikeClusterResponse)(nil),         // 26: messages.DeleteAerospikeClusterResponse
}
var file_service_proto_depIdxs = []int32{
	0,  // 0: messages.AerostationKubernetesAPI.CreateCluster:input_type -> messages.CreateKubernetesClusterRequest
	1,  // 1: messages.AerostationKubernetesAPI.GetCluster:input_type -> messages.GetKubernetesClusterRequest
	2,  // 2: messages.AerostationKubernetesAPI.GetClusters:input_type -> messages.GetKubernetesClustersRequest
	3,  // 3: messages.AerostationKubernetesAPI.UpdateCluster:input_type -> messages.UpdateKubernetesClusterRequest
	4,  // 4: messages.AerostationKubernetesAPI.DeleteCluster:input_type -> messages.DeleteKubernetesClusterRequest
	5,  // 5: messages.AerostationKubernetesAPI.IsKubernetesCluster:input_type -> messages.IsKubernetesClusterRequest
	6,  // 6: messages.AerostationKubernetesAPI.MapRegionToSharedCluster:input_type -> messages.MapRegionToSharedClusterRequest
	7,  // 7: messages.AerostationKubernetesAPI.MapRegionToCluster:input_type -> messages.MapRegionToClusterRequest
	8,  // 8: messages.AerostationAerospikeAPI.CreateCluster:input_type -> messages.CreateAerospikeClusterRequest
	9,  // 9: messages.AerostationAerospikeAPI.GetCluster:input_type -> messages.GetAerospikeClusterRequest
	10, // 10: messages.AerostationAerospikeAPI.GetClusters:input_type -> messages.GetAerospikeClustersByNamespaceRequest
	11, // 11: messages.AerostationAerospikeAPI.GetAllClusters:input_type -> messages.GetAllAerospikeClustersRequest
	12, // 12: messages.AerostationAerospikeAPI.UpdateCluster:input_type -> messages.UpdateAerospikeClusterRequest
	13, // 13: messages.AerostationAerospikeAPI.DeleteCluster:input_type -> messages.DeleteAerospikeClusterRequest
	14, // 14: messages.AerostationKubernetesAPI.CreateCluster:output_type -> messages.CreateKubernetesClusterResponse
	15, // 15: messages.AerostationKubernetesAPI.GetCluster:output_type -> messages.GetKubernetesClusterResponse
	16, // 16: messages.AerostationKubernetesAPI.GetClusters:output_type -> messages.GetKubernetesClustersResponse
	17, // 17: messages.AerostationKubernetesAPI.UpdateCluster:output_type -> messages.UpdateKubernetesClusterResponse
	18, // 18: messages.AerostationKubernetesAPI.DeleteCluster:output_type -> messages.DeleteKubernetesClusterResponse
	19, // 19: messages.AerostationKubernetesAPI.IsKubernetesCluster:output_type -> messages.IsKubernetesClusterResponse
	20, // 20: messages.AerostationKubernetesAPI.MapRegionToSharedCluster:output_type -> messages.MapRegionToSharedClusterResponse
	21, // 21: messages.AerostationKubernetesAPI.MapRegionToCluster:output_type -> messages.MapRegionToClusterResponse
	22, // 22: messages.AerostationAerospikeAPI.CreateCluster:output_type -> messages.CreateAerospikeClusterResponse
	23, // 23: messages.AerostationAerospikeAPI.GetCluster:output_type -> messages.GetAerospikeClusterResponse
	24, // 24: messages.AerostationAerospikeAPI.GetClusters:output_type -> messages.GetAerospikeClustersResponse
	24, // 25: messages.AerostationAerospikeAPI.GetAllClusters:output_type -> messages.GetAerospikeClustersResponse
	25, // 26: messages.AerostationAerospikeAPI.UpdateCluster:output_type -> messages.UpdateAerospikeClusterResponse
	26, // 27: messages.AerostationAerospikeAPI.DeleteCluster:output_type -> messages.DeleteAerospikeClusterResponse
	14, // [14:28] is the sub-list for method output_type
	0,  // [0:14] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_aerospike_proto_init()
	file_kubernetes_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
