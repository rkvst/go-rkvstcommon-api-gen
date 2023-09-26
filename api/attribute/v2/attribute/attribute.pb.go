// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: attribute/v2/attribute/attribute.proto

//
// Structured attributes for attributes in assets and locations
//

package attribute

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type DictAttr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value map[string]string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DictAttr) Reset() {
	*x = DictAttr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attribute_v2_attribute_attribute_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DictAttr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DictAttr) ProtoMessage() {}

func (x *DictAttr) ProtoReflect() protoreflect.Message {
	mi := &file_attribute_v2_attribute_attribute_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DictAttr.ProtoReflect.Descriptor instead.
func (*DictAttr) Descriptor() ([]byte, []int) {
	return file_attribute_v2_attribute_attribute_proto_rawDescGZIP(), []int{0}
}

func (x *DictAttr) GetValue() map[string]string {
	if x != nil {
		return x.Value
	}
	return nil
}

type ListAttr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []*DictAttr `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *ListAttr) Reset() {
	*x = ListAttr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attribute_v2_attribute_attribute_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAttr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAttr) ProtoMessage() {}

func (x *ListAttr) ProtoReflect() protoreflect.Message {
	mi := &file_attribute_v2_attribute_attribute_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAttr.ProtoReflect.Descriptor instead.
func (*ListAttr) Descriptor() ([]byte, []int) {
	return file_attribute_v2_attribute_attribute_proto_rawDescGZIP(), []int{1}
}

func (x *ListAttr) GetValue() []*DictAttr {
	if x != nil {
		return x.Value
	}
	return nil
}

type Attribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*Attribute_StrVal
	//	*Attribute_DictVal
	//	*Attribute_ListVal
	Value isAttribute_Value `protobuf_oneof:"value"`
}

func (x *Attribute) Reset() {
	*x = Attribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attribute_v2_attribute_attribute_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute) ProtoMessage() {}

func (x *Attribute) ProtoReflect() protoreflect.Message {
	mi := &file_attribute_v2_attribute_attribute_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attribute.ProtoReflect.Descriptor instead.
func (*Attribute) Descriptor() ([]byte, []int) {
	return file_attribute_v2_attribute_attribute_proto_rawDescGZIP(), []int{2}
}

func (m *Attribute) GetValue() isAttribute_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Attribute) GetStrVal() string {
	if x, ok := x.GetValue().(*Attribute_StrVal); ok {
		return x.StrVal
	}
	return ""
}

func (x *Attribute) GetDictVal() *DictAttr {
	if x, ok := x.GetValue().(*Attribute_DictVal); ok {
		return x.DictVal
	}
	return nil
}

func (x *Attribute) GetListVal() *ListAttr {
	if x, ok := x.GetValue().(*Attribute_ListVal); ok {
		return x.ListVal
	}
	return nil
}

type isAttribute_Value interface {
	isAttribute_Value()
}

type Attribute_StrVal struct {
	StrVal string `protobuf:"bytes,1,opt,name=str_val,json=strVal,proto3,oneof"`
}

type Attribute_DictVal struct {
	DictVal *DictAttr `protobuf:"bytes,2,opt,name=dict_val,json=dictVal,proto3,oneof"`
}

type Attribute_ListVal struct {
	ListVal *ListAttr `protobuf:"bytes,3,opt,name=list_val,json=listVal,proto3,oneof"`
}

func (*Attribute_StrVal) isAttribute_Value() {}

func (*Attribute_DictVal) isAttribute_Value() {}

func (*Attribute_ListVal) isAttribute_Value() {}

var File_attribute_v2_attribute_attribute_proto protoreflect.FileDescriptor

var file_attribute_v2_attribute_attribute_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x97, 0x01, 0x0a, 0x08, 0x44, 0x69, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x12, 0x51, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x41,
	0x74, 0x74, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x18,
	0xfa, 0x42, 0x15, 0x9a, 0x01, 0x12, 0x22, 0x09, 0x72, 0x07, 0x18, 0x80, 0x08, 0xba, 0x01, 0x01,
	0x2e, 0x2a, 0x05, 0x72, 0x03, 0x18, 0x80, 0x20, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a,
	0x38, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x38, 0x0a, 0x08, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x74, 0x74, 0x72, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74,
	0x2e, 0x76, 0x32, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0xa3, 0x01, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x12, 0x23, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0x80, 0x20, 0x48, 0x00, 0x52, 0x06,
	0x73, 0x74, 0x72, 0x56, 0x61, 0x6c, 0x12, 0x33, 0x0a, 0x08, 0x64, 0x69, 0x63, 0x74, 0x5f, 0x76,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72,
	0x48, 0x00, 0x52, 0x07, 0x64, 0x69, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x33, 0x0a, 0x08, 0x6c,
	0x69, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x74, 0x74, 0x72, 0x48, 0x00, 0x52, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c,
	0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74,
	0x68, 0x69, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6b, 0x76, 0x73, 0x74, 0x2f, 0x67, 0x6f,
	0x2d, 0x72, 0x6b, 0x76, 0x73, 0x74, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69,
	0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x3b,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_attribute_v2_attribute_attribute_proto_rawDescOnce sync.Once
	file_attribute_v2_attribute_attribute_proto_rawDescData = file_attribute_v2_attribute_attribute_proto_rawDesc
)

func file_attribute_v2_attribute_attribute_proto_rawDescGZIP() []byte {
	file_attribute_v2_attribute_attribute_proto_rawDescOnce.Do(func() {
		file_attribute_v2_attribute_attribute_proto_rawDescData = protoimpl.X.CompressGZIP(file_attribute_v2_attribute_attribute_proto_rawDescData)
	})
	return file_attribute_v2_attribute_attribute_proto_rawDescData
}

var file_attribute_v2_attribute_attribute_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_attribute_v2_attribute_attribute_proto_goTypes = []interface{}{
	(*DictAttr)(nil),  // 0: archivist.v2.DictAttr
	(*ListAttr)(nil),  // 1: archivist.v2.ListAttr
	(*Attribute)(nil), // 2: archivist.v2.Attribute
	nil,               // 3: archivist.v2.DictAttr.ValueEntry
}
var file_attribute_v2_attribute_attribute_proto_depIdxs = []int32{
	3, // 0: archivist.v2.DictAttr.value:type_name -> archivist.v2.DictAttr.ValueEntry
	0, // 1: archivist.v2.ListAttr.value:type_name -> archivist.v2.DictAttr
	0, // 2: archivist.v2.Attribute.dict_val:type_name -> archivist.v2.DictAttr
	1, // 3: archivist.v2.Attribute.list_val:type_name -> archivist.v2.ListAttr
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_attribute_v2_attribute_attribute_proto_init() }
func file_attribute_v2_attribute_attribute_proto_init() {
	if File_attribute_v2_attribute_attribute_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_attribute_v2_attribute_attribute_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DictAttr); i {
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
		file_attribute_v2_attribute_attribute_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAttr); i {
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
		file_attribute_v2_attribute_attribute_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attribute); i {
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
	file_attribute_v2_attribute_attribute_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Attribute_StrVal)(nil),
		(*Attribute_DictVal)(nil),
		(*Attribute_ListVal)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_attribute_v2_attribute_attribute_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_attribute_v2_attribute_attribute_proto_goTypes,
		DependencyIndexes: file_attribute_v2_attribute_attribute_proto_depIdxs,
		MessageInfos:      file_attribute_v2_attribute_attribute_proto_msgTypes,
	}.Build()
	File_attribute_v2_attribute_attribute_proto = out.File
	file_attribute_v2_attribute_attribute_proto_rawDesc = nil
	file_attribute_v2_attribute_attribute_proto_goTypes = nil
	file_attribute_v2_attribute_attribute_proto_depIdxs = nil
}
