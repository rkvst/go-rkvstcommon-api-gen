// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: datatrails-common-api/caps/v1/caps/caps.proto

//
// Caps messages
//

package caps

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type GetCapsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCapsRequest) Reset() {
	*x = GetCapsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datatrails_common_api_caps_v1_caps_caps_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCapsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCapsRequest) ProtoMessage() {}

func (x *GetCapsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datatrails_common_api_caps_v1_caps_caps_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCapsRequest.ProtoReflect.Descriptor instead.
func (*GetCapsRequest) Descriptor() ([]byte, []int) {
	return file_datatrails_common_api_caps_v1_caps_caps_proto_rawDescGZIP(), []int{0}
}

type Cap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceType      string `protobuf:"bytes,1,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	ResourceRemaining int64  `protobuf:"varint,2,opt,name=resource_remaining,json=resourceRemaining,proto3" json:"resource_remaining,omitempty"`
}

func (x *Cap) Reset() {
	*x = Cap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datatrails_common_api_caps_v1_caps_caps_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cap) ProtoMessage() {}

func (x *Cap) ProtoReflect() protoreflect.Message {
	mi := &file_datatrails_common_api_caps_v1_caps_caps_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cap.ProtoReflect.Descriptor instead.
func (*Cap) Descriptor() ([]byte, []int) {
	return file_datatrails_common_api_caps_v1_caps_caps_proto_rawDescGZIP(), []int{1}
}

func (x *Cap) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *Cap) GetResourceRemaining() int64 {
	if x != nil {
		return x.ResourceRemaining
	}
	return 0
}

type Caps struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Caps []*Cap `protobuf:"bytes,1,rep,name=caps,proto3" json:"caps,omitempty"`
}

func (x *Caps) Reset() {
	*x = Caps{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datatrails_common_api_caps_v1_caps_caps_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Caps) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Caps) ProtoMessage() {}

func (x *Caps) ProtoReflect() protoreflect.Message {
	mi := &file_datatrails_common_api_caps_v1_caps_caps_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Caps.ProtoReflect.Descriptor instead.
func (*Caps) Descriptor() ([]byte, []int) {
	return file_datatrails_common_api_caps_v1_caps_caps_proto_rawDescGZIP(), []int{2}
}

func (x *Caps) GetCaps() []*Cap {
	if x != nil {
		return x.Caps
	}
	return nil
}

var File_datatrails_common_api_caps_v1_caps_caps_proto protoreflect.FileDescriptor

var file_datatrails_common_api_caps_v1_caps_caps_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x2d, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x61, 0x70, 0x73, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x70,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x3a, 0x3e, 0x92, 0x41, 0x3b, 0x0a, 0x39, 0x32,
	0x37, 0x47, 0x45, 0x54, 0x20, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x66, 0x6f, 0x72,
	0x20, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x20, 0x63, 0x61, 0x70, 0x70, 0x65,
	0x64, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x20, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x22, 0xd0, 0x01, 0x0a, 0x03, 0x43, 0x61, 0x70,
	0x12, 0x5f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3a, 0x92, 0x41, 0x2d, 0x32, 0x2b, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x79, 0x69, 0x6e, 0x67,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x61, 0x70, 0x70, 0x65, 0x64, 0x20, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x20, 0x74, 0x79, 0x70, 0x65, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01,
	0x18, 0x80, 0x08, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x68, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x65,
	0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x39, 0x92,
	0x41, 0x26, 0x32, 0x24, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x20, 0x6f, 0x66, 0x20, 0x63, 0x61,
	0x70, 0x70, 0x65, 0x64, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x20, 0x72,
	0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0xfa, 0x42, 0x0d, 0x22, 0x0b, 0x28, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x52, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0x55, 0x0a, 0x04, 0x43,
	0x61, 0x70, 0x73, 0x12, 0x4d, 0x0a, 0x04, 0x63, 0x61, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x61, 0x70, 0x42, 0x26, 0x92, 0x41, 0x23, 0x32, 0x21, 0x4c, 0x69, 0x73, 0x74, 0x20,
	0x6f, 0x66, 0x20, 0x72, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x6e, 0x74, 0x20, 0x63, 0x61, 0x70, 0x70,
	0x65, 0x64, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x04, 0x63, 0x61,
	0x70, 0x73, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x64,
	0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x61, 0x70, 0x73, 0x3b, 0x63, 0x61, 0x70, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_datatrails_common_api_caps_v1_caps_caps_proto_rawDescOnce sync.Once
	file_datatrails_common_api_caps_v1_caps_caps_proto_rawDescData = file_datatrails_common_api_caps_v1_caps_caps_proto_rawDesc
)

func file_datatrails_common_api_caps_v1_caps_caps_proto_rawDescGZIP() []byte {
	file_datatrails_common_api_caps_v1_caps_caps_proto_rawDescOnce.Do(func() {
		file_datatrails_common_api_caps_v1_caps_caps_proto_rawDescData = protoimpl.X.CompressGZIP(file_datatrails_common_api_caps_v1_caps_caps_proto_rawDescData)
	})
	return file_datatrails_common_api_caps_v1_caps_caps_proto_rawDescData
}

var file_datatrails_common_api_caps_v1_caps_caps_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_datatrails_common_api_caps_v1_caps_caps_proto_goTypes = []interface{}{
	(*GetCapsRequest)(nil), // 0: archivist.v1.GetCapsRequest
	(*Cap)(nil),            // 1: archivist.v1.Cap
	(*Caps)(nil),           // 2: archivist.v1.Caps
}
var file_datatrails_common_api_caps_v1_caps_caps_proto_depIdxs = []int32{
	1, // 0: archivist.v1.Caps.caps:type_name -> archivist.v1.Cap
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_datatrails_common_api_caps_v1_caps_caps_proto_init() }
func file_datatrails_common_api_caps_v1_caps_caps_proto_init() {
	if File_datatrails_common_api_caps_v1_caps_caps_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datatrails_common_api_caps_v1_caps_caps_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCapsRequest); i {
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
		file_datatrails_common_api_caps_v1_caps_caps_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cap); i {
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
		file_datatrails_common_api_caps_v1_caps_caps_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Caps); i {
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
			RawDescriptor: file_datatrails_common_api_caps_v1_caps_caps_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datatrails_common_api_caps_v1_caps_caps_proto_goTypes,
		DependencyIndexes: file_datatrails_common_api_caps_v1_caps_caps_proto_depIdxs,
		MessageInfos:      file_datatrails_common_api_caps_v1_caps_caps_proto_msgTypes,
	}.Build()
	File_datatrails_common_api_caps_v1_caps_caps_proto = out.File
	file_datatrails_common_api_caps_v1_caps_caps_proto_rawDesc = nil
	file_datatrails_common_api_caps_v1_caps_caps_proto_goTypes = nil
	file_datatrails_common_api_caps_v1_caps_caps_proto_depIdxs = nil
}
