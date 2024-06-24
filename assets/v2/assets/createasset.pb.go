// Maintainers, please refer to the style guide here:
//     https://developers.google.com/protocol-buffers/docs/style

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: datatrails-common-api/assets/v2/assets/createasset.proto

package assets

import (
	attribute "github.com/datatrails/go-datatrails-common-api-gen/attribute/v2/attribute"
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

type CreateAssetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list of enabled behaviours
	Behaviours []string `protobuf:"bytes,1,rep,name=behaviours,proto3" json:"behaviours,omitempty"`
	// event attributes for the assets creation
	Attributes map[string]*attribute.Attribute `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// mechanism for evidential proof for Events on this Asset
	ProofMechanism ProofMechanism `protobuf:"varint,4,opt,name=proof_mechanism,json=proofMechanism,proto3,enum=archivist.v2.ProofMechanism" json:"proof_mechanism,omitempty"`
	ChainId        string         `protobuf:"bytes,5,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Public         bool           `protobuf:"varint,6,opt,name=public,proto3" json:"public,omitempty"`
}

func (x *CreateAssetRequest) Reset() {
	*x = CreateAssetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datatrails_common_api_assets_v2_assets_createasset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAssetRequest) ProtoMessage() {}

func (x *CreateAssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datatrails_common_api_assets_v2_assets_createasset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAssetRequest.ProtoReflect.Descriptor instead.
func (*CreateAssetRequest) Descriptor() ([]byte, []int) {
	return file_datatrails_common_api_assets_v2_assets_createasset_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAssetRequest) GetBehaviours() []string {
	if x != nil {
		return x.Behaviours
	}
	return nil
}

func (x *CreateAssetRequest) GetAttributes() map[string]*attribute.Attribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *CreateAssetRequest) GetProofMechanism() ProofMechanism {
	if x != nil {
		return x.ProofMechanism
	}
	return ProofMechanism_PROOF_MECHANISM_UNSPECIFIED
}

func (x *CreateAssetRequest) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *CreateAssetRequest) GetPublic() bool {
	if x != nil {
		return x.Public
	}
	return false
}

var File_datatrails_common_api_assets_v2_assets_createasset_proto protoreflect.FileDescriptor

var file_datatrails_common_api_assets_v2_assets_createasset_proto_rawDesc = []byte{
	0x0a, 0x38, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x2d, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76,
	0x32, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x32, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x2d, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f,
	0x76, 0x32, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3c, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x69, 0x6c,
	0x73, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x09, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5e, 0x0a, 0x0a, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x3e,
	0x92, 0x41, 0x2e, 0x32, 0x29, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x73, 0x20, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x61, 0x73, 0x73, 0x65, 0x74, 0x78, 0x80,
	0x08, 0xfa, 0x42, 0x0a, 0x92, 0x01, 0x07, 0x22, 0x05, 0x72, 0x03, 0x18, 0x80, 0x08, 0x52, 0x0a,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x73, 0x12, 0xed, 0x01, 0x0a, 0x0a, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x30, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x42, 0x9a, 0x01, 0x92, 0x41, 0x2a, 0x32, 0x25, 0x6b, 0x65, 0x79, 0x20, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x20, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x20, 0x6f, 0x66, 0x20, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x20, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x78, 0x80,
	0x08, 0xfa, 0x42, 0x6a, 0x9a, 0x01, 0x67, 0x22, 0x65, 0x72, 0x63, 0x10, 0x01, 0x18, 0x80, 0x08,
	0x32, 0x58, 0x5e, 0x5b, 0x5e, 0x5b, 0x3a, 0x63, 0x6e, 0x74, 0x72, 0x6c, 0x3a, 0x5d, 0x5d, 0x2a,
	0x3f, 0x5b, 0x5e, 0x5b, 0x5b, 0x3a, 0x63, 0x6e, 0x74, 0x72, 0x6c, 0x3a, 0x5d, 0x5d, 0x2b, 0x3f,
	0x5b, 0x5e, 0x5b, 0x3a, 0x63, 0x6e, 0x74, 0x72, 0x6c, 0x3a, 0x5d, 0x5d, 0x24, 0x7c, 0x5e, 0x5b,
	0x5e, 0x5b, 0x3a, 0x63, 0x6e, 0x74, 0x72, 0x6c, 0x3a, 0x5d, 0x5d, 0x24, 0x7c, 0x5e, 0x5b, 0x5e,
	0x5b, 0x3a, 0x63, 0x6e, 0x74, 0x72, 0x6c, 0x3a, 0x5d, 0x5d, 0x2a, 0x3f, 0x5b, 0x5e, 0x5d, 0x5b,
	0x3a, 0x63, 0x6e, 0x74, 0x72, 0x6c, 0x3a, 0x5d, 0x5d, 0x24, 0xba, 0x01, 0x01, 0x2e, 0x52, 0x0a,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0xa7, 0x01, 0x0a, 0x0f, 0x70,
	0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74,
	0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69,
	0x73, 0x6d, 0x42, 0x60, 0x92, 0x41, 0x55, 0x32, 0x4f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x79,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x20, 0x75,
	0x73, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x20, 0x65,
	0x76, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x20, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68,
	0x69, 0x73, 0x20, 0x41, 0x73, 0x73, 0x65, 0x74, 0x9a, 0x02, 0x01, 0x07, 0xfa, 0x42, 0x05, 0x82,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x4d, 0x65, 0x63, 0x68, 0x61,
	0x6e, 0x69, 0x73, 0x6d, 0x12, 0x60, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x45, 0x92, 0x41, 0x3a, 0x32, 0x35, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x20, 0x69, 0x64, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x20, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74,
	0x65, 0x64, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x78, 0x80, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0x80, 0x08, 0x52, 0x07, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0xb7, 0x01, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x42, 0x9e, 0x01, 0x92, 0x41, 0x9a, 0x01, 0x32, 0x97,
	0x01, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x20, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x20, 0x41,
	0x20, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x20, 0x61, 0x73, 0x73, 0x65, 0x74, 0x20, 0x61, 0x6e,
	0x64, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x69, 0x74, 0x73, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x20, 0x61, 0x72, 0x65, 0x20, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x20, 0x74, 0x6f, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x20, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x6f, 0x20, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x20, 0x69, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x20, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x1a, 0x56, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74,
	0x2e, 0x76, 0x32, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x98, 0x02, 0x92, 0x41, 0x94, 0x02, 0x0a,
	0x3d, 0x32, 0x3b, 0x54, 0x68, 0x69, 0x73, 0x20, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62, 0x6f, 0x64, 0x79, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x20,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x41, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x32, 0xd2,
	0x01, 0x7b, 0x20, 0x22, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x73, 0x22, 0x3a,
	0x20, 0x5b, 0x22, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63,
	0x65, 0x22, 0x5d, 0x2c, 0x20, 0x22, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x22, 0x3a, 0x20, 0x7b, 0x20, 0x20, 0x20, 0x20, 0x22, 0x61, 0x72, 0x63, 0x5f, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x47, 0x61, 0x72,
	0x64, 0x65, 0x6e, 0x20, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x2c, 0x20, 0x20, 0x20, 0x20, 0x22,
	0x61, 0x72, 0x63, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x3a, 0x20, 0x22, 0x4d, 0x79, 0x20, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x20, 0x46, 0x65,
	0x6e, 0x63, 0x65, 0x22, 0x2c, 0x20, 0x20, 0x20, 0x20, 0x22, 0x63, 0x6f, 0x6c, 0x6f, 0x75, 0x72,
	0x22, 0x3a, 0x20, 0x22, 0x50, 0x6c, 0x61, 0x69, 0x6e, 0x20, 0x77, 0x6f, 0x6f, 0x64, 0x22, 0x20,
	0x7d, 0x2c, 0x20, 0x22, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e,
	0x69, 0x73, 0x6d, 0x22, 0x3a, 0x20, 0x22, 0x4d, 0x45, 0x52, 0x4b, 0x4c, 0x45, 0x5f, 0x4c, 0x4f,
	0x47, 0x22, 0x2c, 0x22, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0x3a, 0x20, 0x66, 0x61, 0x6c,
	0x73, 0x65, 0x7d, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x52, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x69, 0x74, 0x79, 0x42, 0x4c, 0x5a, 0x4a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x74,
	0x72, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61,
	0x69, 0x6c, 0x73, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67,
	0x65, 0x6e, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x3b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_datatrails_common_api_assets_v2_assets_createasset_proto_rawDescOnce sync.Once
	file_datatrails_common_api_assets_v2_assets_createasset_proto_rawDescData = file_datatrails_common_api_assets_v2_assets_createasset_proto_rawDesc
)

func file_datatrails_common_api_assets_v2_assets_createasset_proto_rawDescGZIP() []byte {
	file_datatrails_common_api_assets_v2_assets_createasset_proto_rawDescOnce.Do(func() {
		file_datatrails_common_api_assets_v2_assets_createasset_proto_rawDescData = protoimpl.X.CompressGZIP(file_datatrails_common_api_assets_v2_assets_createasset_proto_rawDescData)
	})
	return file_datatrails_common_api_assets_v2_assets_createasset_proto_rawDescData
}

var file_datatrails_common_api_assets_v2_assets_createasset_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_datatrails_common_api_assets_v2_assets_createasset_proto_goTypes = []interface{}{
	(*CreateAssetRequest)(nil),  // 0: archivist.v2.CreateAssetRequest
	nil,                         // 1: archivist.v2.CreateAssetRequest.AttributesEntry
	(ProofMechanism)(0),         // 2: archivist.v2.ProofMechanism
	(*attribute.Attribute)(nil), // 3: archivist.v2.Attribute
}
var file_datatrails_common_api_assets_v2_assets_createasset_proto_depIdxs = []int32{
	1, // 0: archivist.v2.CreateAssetRequest.attributes:type_name -> archivist.v2.CreateAssetRequest.AttributesEntry
	2, // 1: archivist.v2.CreateAssetRequest.proof_mechanism:type_name -> archivist.v2.ProofMechanism
	3, // 2: archivist.v2.CreateAssetRequest.AttributesEntry.value:type_name -> archivist.v2.Attribute
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_datatrails_common_api_assets_v2_assets_createasset_proto_init() }
func file_datatrails_common_api_assets_v2_assets_createasset_proto_init() {
	if File_datatrails_common_api_assets_v2_assets_createasset_proto != nil {
		return
	}
	file_datatrails_common_api_assets_v2_assets_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_datatrails_common_api_assets_v2_assets_createasset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAssetRequest); i {
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
			RawDescriptor: file_datatrails_common_api_assets_v2_assets_createasset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datatrails_common_api_assets_v2_assets_createasset_proto_goTypes,
		DependencyIndexes: file_datatrails_common_api_assets_v2_assets_createasset_proto_depIdxs,
		MessageInfos:      file_datatrails_common_api_assets_v2_assets_createasset_proto_msgTypes,
	}.Build()
	File_datatrails_common_api_assets_v2_assets_createasset_proto = out.File
	file_datatrails_common_api_assets_v2_assets_createasset_proto_rawDesc = nil
	file_datatrails_common_api_assets_v2_assets_createasset_proto_goTypes = nil
	file_datatrails_common_api_assets_v2_assets_createasset_proto_depIdxs = nil
}
