// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: assets/v2/assets/enums.proto

package assets

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

type ConfirmationStatus int32

const (
	ConfirmationStatus_CONFIRMATION_STATUS_UNSPECIFIED ConfirmationStatus = 0
	ConfirmationStatus_PENDING                         ConfirmationStatus = 1 // not yet committed
	ConfirmationStatus_CONFIRMED                       ConfirmationStatus = 2 // committed
	ConfirmationStatus_FAILED                          ConfirmationStatus = 3 // permanent failure
)

// Enum value maps for ConfirmationStatus.
var (
	ConfirmationStatus_name = map[int32]string{
		0: "CONFIRMATION_STATUS_UNSPECIFIED",
		1: "PENDING",
		2: "CONFIRMED",
		3: "FAILED",
	}
	ConfirmationStatus_value = map[string]int32{
		"CONFIRMATION_STATUS_UNSPECIFIED": 0,
		"PENDING":                         1,
		"CONFIRMED":                       2,
		"FAILED":                          3,
	}
)

func (x ConfirmationStatus) Enum() *ConfirmationStatus {
	p := new(ConfirmationStatus)
	*p = x
	return p
}

func (x ConfirmationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConfirmationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_assets_v2_assets_enums_proto_enumTypes[0].Descriptor()
}

func (ConfirmationStatus) Type() protoreflect.EnumType {
	return &file_assets_v2_assets_enums_proto_enumTypes[0]
}

func (x ConfirmationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConfirmationStatus.Descriptor instead.
func (ConfirmationStatus) EnumDescriptor() ([]byte, []int) {
	return file_assets_v2_assets_enums_proto_rawDescGZIP(), []int{0}
}

type TrackedStatus int32

const (
	TrackedStatus_TRACKED_STATUS_UNSPECIFIED TrackedStatus = 0
	TrackedStatus_TRACKED                    TrackedStatus = 1
	TrackedStatus_NOT_TRACKED                TrackedStatus = 2
	TrackedStatus_ANY                        TrackedStatus = 3
)

// Enum value maps for TrackedStatus.
var (
	TrackedStatus_name = map[int32]string{
		0: "TRACKED_STATUS_UNSPECIFIED",
		1: "TRACKED",
		2: "NOT_TRACKED",
		3: "ANY",
	}
	TrackedStatus_value = map[string]int32{
		"TRACKED_STATUS_UNSPECIFIED": 0,
		"TRACKED":                    1,
		"NOT_TRACKED":                2,
		"ANY":                        3,
	}
)

func (x TrackedStatus) Enum() *TrackedStatus {
	p := new(TrackedStatus)
	*p = x
	return p
}

func (x TrackedStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TrackedStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_assets_v2_assets_enums_proto_enumTypes[1].Descriptor()
}

func (TrackedStatus) Type() protoreflect.EnumType {
	return &file_assets_v2_assets_enums_proto_enumTypes[1]
}

func (x TrackedStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TrackedStatus.Descriptor instead.
func (TrackedStatus) EnumDescriptor() ([]byte, []int) {
	return file_assets_v2_assets_enums_proto_rawDescGZIP(), []int{1}
}

// Specifies how the asset data will be stored. This is set once on creation
// and does not change.
type StorageIntegrity int32

const (
	StorageIntegrity_STORAGE_INTEGRITY_UNSPECIFIED StorageIntegrity = 0
	StorageIntegrity_LEDGER                        StorageIntegrity = 1
	StorageIntegrity_TENANT_STORAGE                StorageIntegrity = 2
)

// Enum value maps for StorageIntegrity.
var (
	StorageIntegrity_name = map[int32]string{
		0: "STORAGE_INTEGRITY_UNSPECIFIED",
		1: "LEDGER",
		2: "TENANT_STORAGE",
	}
	StorageIntegrity_value = map[string]int32{
		"STORAGE_INTEGRITY_UNSPECIFIED": 0,
		"LEDGER":                        1,
		"TENANT_STORAGE":                2,
	}
)

func (x StorageIntegrity) Enum() *StorageIntegrity {
	p := new(StorageIntegrity)
	*p = x
	return p
}

func (x StorageIntegrity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StorageIntegrity) Descriptor() protoreflect.EnumDescriptor {
	return file_assets_v2_assets_enums_proto_enumTypes[2].Descriptor()
}

func (StorageIntegrity) Type() protoreflect.EnumType {
	return &file_assets_v2_assets_enums_proto_enumTypes[2]
}

func (x StorageIntegrity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StorageIntegrity.Descriptor instead.
func (StorageIntegrity) EnumDescriptor() ([]byte, []int) {
	return file_assets_v2_assets_enums_proto_rawDescGZIP(), []int{2}
}

type ProofMechanism int32

const (
	ProofMechanism_PROOF_MECHANISM_UNSPECIFIED ProofMechanism = 0
	ProofMechanism_KHIPU                       ProofMechanism = 1
	ProofMechanism_SIMPLE_HASH                 ProofMechanism = 2
	ProofMechanism_TRIE_HASH                   ProofMechanism = 3
	ProofMechanism_TRIE_HASH_MERKLE            ProofMechanism = 4
	ProofMechanism_TRIE_HASH_VERKLE            ProofMechanism = 5
)

// Enum value maps for ProofMechanism.
var (
	ProofMechanism_name = map[int32]string{
		0: "PROOF_MECHANISM_UNSPECIFIED",
		1: "KHIPU",
		2: "SIMPLE_HASH",
		3: "TRIE_HASH",
		4: "TRIE_HASH_MERKLE",
		5: "TRIE_HASH_VERKLE",
	}
	ProofMechanism_value = map[string]int32{
		"PROOF_MECHANISM_UNSPECIFIED": 0,
		"KHIPU":                       1,
		"SIMPLE_HASH":                 2,
		"TRIE_HASH":                   3,
		"TRIE_HASH_MERKLE":            4,
		"TRIE_HASH_VERKLE":            5,
	}
)

func (x ProofMechanism) Enum() *ProofMechanism {
	p := new(ProofMechanism)
	*p = x
	return p
}

func (x ProofMechanism) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProofMechanism) Descriptor() protoreflect.EnumDescriptor {
	return file_assets_v2_assets_enums_proto_enumTypes[3].Descriptor()
}

func (ProofMechanism) Type() protoreflect.EnumType {
	return &file_assets_v2_assets_enums_proto_enumTypes[3]
}

func (x ProofMechanism) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProofMechanism.Descriptor instead.
func (ProofMechanism) EnumDescriptor() ([]byte, []int) {
	return file_assets_v2_assets_enums_proto_rawDescGZIP(), []int{3}
}

var File_assets_v2_assets_enums_proto protoreflect.FileDescriptor

var file_assets_v2_assets_enums_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x2a, 0x61, 0x0a, 0x12,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x23, 0x0a, 0x1f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x45,
	0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x2a,
	0x56, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x45, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0f, 0x0a,
	0x0b, 0x4e, 0x4f, 0x54, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x02, 0x12, 0x07,
	0x0a, 0x03, 0x41, 0x4e, 0x59, 0x10, 0x03, 0x2a, 0x55, 0x0a, 0x10, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x69, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x1d, 0x53,
	0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x47, 0x52, 0x49, 0x54, 0x59,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x4c, 0x45, 0x44, 0x47, 0x45, 0x52, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x45,
	0x4e, 0x41, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x10, 0x02, 0x2a, 0x88,
	0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73,
	0x6d, 0x12, 0x1f, 0x0a, 0x1b, 0x50, 0x52, 0x4f, 0x4f, 0x46, 0x5f, 0x4d, 0x45, 0x43, 0x48, 0x41,
	0x4e, 0x49, 0x53, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4b, 0x48, 0x49, 0x50, 0x55, 0x10, 0x01, 0x12, 0x0f, 0x0a,
	0x0b, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x48, 0x41, 0x53, 0x48, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x54, 0x52, 0x49, 0x45, 0x5f, 0x48, 0x41, 0x53, 0x48, 0x10, 0x03, 0x12, 0x14, 0x0a,
	0x10, 0x54, 0x52, 0x49, 0x45, 0x5f, 0x48, 0x41, 0x53, 0x48, 0x5f, 0x4d, 0x45, 0x52, 0x4b, 0x4c,
	0x45, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x52, 0x49, 0x45, 0x5f, 0x48, 0x41, 0x53, 0x48,
	0x5f, 0x56, 0x45, 0x52, 0x4b, 0x4c, 0x45, 0x10, 0x05, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74,
	0x68, 0x69, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6b, 0x76, 0x73, 0x74, 0x2f, 0x67, 0x6f,
	0x2d, 0x72, 0x6b, 0x76, 0x73, 0x74, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69,
	0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f,
	0x76, 0x32, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x3b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_assets_v2_assets_enums_proto_rawDescOnce sync.Once
	file_assets_v2_assets_enums_proto_rawDescData = file_assets_v2_assets_enums_proto_rawDesc
)

func file_assets_v2_assets_enums_proto_rawDescGZIP() []byte {
	file_assets_v2_assets_enums_proto_rawDescOnce.Do(func() {
		file_assets_v2_assets_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_assets_v2_assets_enums_proto_rawDescData)
	})
	return file_assets_v2_assets_enums_proto_rawDescData
}

var file_assets_v2_assets_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_assets_v2_assets_enums_proto_goTypes = []interface{}{
	(ConfirmationStatus)(0), // 0: archivist.v2.ConfirmationStatus
	(TrackedStatus)(0),      // 1: archivist.v2.TrackedStatus
	(StorageIntegrity)(0),   // 2: archivist.v2.StorageIntegrity
	(ProofMechanism)(0),     // 3: archivist.v2.ProofMechanism
}
var file_assets_v2_assets_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_assets_v2_assets_enums_proto_init() }
func file_assets_v2_assets_enums_proto_init() {
	if File_assets_v2_assets_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_assets_v2_assets_enums_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_assets_v2_assets_enums_proto_goTypes,
		DependencyIndexes: file_assets_v2_assets_enums_proto_depIdxs,
		EnumInfos:         file_assets_v2_assets_enums_proto_enumTypes,
	}.Build()
	File_assets_v2_assets_enums_proto = out.File
	file_assets_v2_assets_enums_proto_rawDesc = nil
	file_assets_v2_assets_enums_proto_goTypes = nil
	file_assets_v2_assets_enums_proto_depIdxs = nil
}
