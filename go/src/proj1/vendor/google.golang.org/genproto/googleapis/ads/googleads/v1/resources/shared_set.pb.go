// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/shared_set.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// SharedSets are used for sharing criterion exclusions across multiple
// campaigns.
type SharedSet struct {
	// The resource name of the shared set.
	// Shared set resource names have the form:
	//
	// `customers/{customer_id}/sharedSets/{shared_set_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of this shared set. Read only.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The type of this shared set: each shared set holds only a single kind
	// of resource. Required. Immutable.
	Type enums.SharedSetTypeEnum_SharedSetType `protobuf:"varint,3,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.SharedSetTypeEnum_SharedSetType" json:"type,omitempty"`
	// The name of this shared set. Required.
	// Shared Sets must have names that are unique among active shared sets of
	// the same type.
	// The length of this string should be between 1 and 255 UTF-8 bytes,
	// inclusive.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The status of this shared set. Read only.
	Status enums.SharedSetStatusEnum_SharedSetStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.SharedSetStatusEnum_SharedSetStatus" json:"status,omitempty"`
	// The number of shared criteria within this shared set. Read only.
	MemberCount *wrappers.Int64Value `protobuf:"bytes,6,opt,name=member_count,json=memberCount,proto3" json:"member_count,omitempty"`
	// The number of campaigns associated with this shared set. Read only.
	ReferenceCount       *wrappers.Int64Value `protobuf:"bytes,7,opt,name=reference_count,json=referenceCount,proto3" json:"reference_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SharedSet) Reset()         { *m = SharedSet{} }
func (m *SharedSet) String() string { return proto.CompactTextString(m) }
func (*SharedSet) ProtoMessage()    {}
func (*SharedSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_shared_set_2ded0d66d5306c68, []int{0}
}
func (m *SharedSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharedSet.Unmarshal(m, b)
}
func (m *SharedSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharedSet.Marshal(b, m, deterministic)
}
func (dst *SharedSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharedSet.Merge(dst, src)
}
func (m *SharedSet) XXX_Size() int {
	return xxx_messageInfo_SharedSet.Size(m)
}
func (m *SharedSet) XXX_DiscardUnknown() {
	xxx_messageInfo_SharedSet.DiscardUnknown(m)
}

var xxx_messageInfo_SharedSet proto.InternalMessageInfo

func (m *SharedSet) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *SharedSet) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SharedSet) GetType() enums.SharedSetTypeEnum_SharedSetType {
	if m != nil {
		return m.Type
	}
	return enums.SharedSetTypeEnum_UNSPECIFIED
}

func (m *SharedSet) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *SharedSet) GetStatus() enums.SharedSetStatusEnum_SharedSetStatus {
	if m != nil {
		return m.Status
	}
	return enums.SharedSetStatusEnum_UNSPECIFIED
}

func (m *SharedSet) GetMemberCount() *wrappers.Int64Value {
	if m != nil {
		return m.MemberCount
	}
	return nil
}

func (m *SharedSet) GetReferenceCount() *wrappers.Int64Value {
	if m != nil {
		return m.ReferenceCount
	}
	return nil
}

func init() {
	proto.RegisterType((*SharedSet)(nil), "google.ads.googleads.v1.resources.SharedSet")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/shared_set.proto", fileDescriptor_shared_set_2ded0d66d5306c68)
}

var fileDescriptor_shared_set_2ded0d66d5306c68 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x6b, 0xdb, 0x30,
	0x14, 0xc7, 0xb1, 0x93, 0x65, 0x54, 0xed, 0x32, 0xf0, 0xc9, 0x74, 0x65, 0xa4, 0x1b, 0x85, 0xc0,
	0x40, 0x9e, 0xd3, 0x6d, 0x07, 0x0d, 0x0a, 0xce, 0x36, 0xca, 0x76, 0x18, 0xc5, 0x19, 0x39, 0x94,
	0x40, 0x50, 0xa2, 0x57, 0xcf, 0x10, 0x4b, 0x46, 0x92, 0x3b, 0xfa, 0x75, 0x76, 0xdc, 0xf7, 0xd8,
	0x65, 0x1f, 0x65, 0x5f, 0x61, 0x97, 0x61, 0xc9, 0x12, 0x2b, 0xa3, 0x4d, 0x6e, 0xcf, 0x4f, 0xff,
	0xff, 0x4f, 0xff, 0xf7, 0x90, 0xd1, 0xa4, 0x10, 0xa2, 0xd8, 0x40, 0x42, 0x99, 0x4a, 0x6c, 0xd9,
	0x56, 0xd7, 0x69, 0x22, 0x41, 0x89, 0x46, 0xae, 0x41, 0x25, 0xea, 0x2b, 0x95, 0xc0, 0x96, 0x0a,
	0x34, 0xae, 0xa5, 0xd0, 0x22, 0x3a, 0xb6, 0x42, 0x4c, 0x99, 0xc2, 0xde, 0x83, 0xaf, 0x53, 0xec,
	0x3d, 0x87, 0xaf, 0xef, 0xc2, 0x02, 0x6f, 0xaa, 0x7f, 0x91, 0x4b, 0xa5, 0xa9, 0x6e, 0x94, 0x25,
	0x1f, 0x9e, 0xee, 0x6c, 0xd3, 0x37, 0x35, 0x74, 0xa6, 0xa7, 0x9d, 0xc9, 0x7c, 0xad, 0x9a, 0xab,
	0xe4, 0x9b, 0xa4, 0x75, 0x0d, 0xd2, 0x41, 0x8f, 0x1c, 0xb4, 0x2e, 0x13, 0xca, 0xb9, 0xd0, 0x54,
	0x97, 0x82, 0x77, 0xa7, 0xcf, 0x7e, 0xf6, 0xd0, 0xde, 0xcc, 0x70, 0x67, 0xa0, 0xa3, 0xe7, 0xe8,
	0x91, 0x1b, 0x62, 0xc9, 0x69, 0x05, 0x71, 0x30, 0x0a, 0xc6, 0x7b, 0xf9, 0x81, 0x6b, 0x7e, 0xa6,
	0x15, 0x44, 0x2f, 0x50, 0x58, 0xb2, 0x38, 0x1c, 0x05, 0xe3, 0xfd, 0xc9, 0x93, 0x6e, 0x03, 0xd8,
	0xdd, 0x8e, 0x3f, 0x72, 0xfd, 0xe6, 0xd5, 0x9c, 0x6e, 0x1a, 0xc8, 0xc3, 0x92, 0x45, 0x39, 0xea,
	0xb7, 0x59, 0xe3, 0xde, 0x28, 0x18, 0x0f, 0x27, 0x67, 0xf8, 0xae, 0xdd, 0x99, 0x09, 0xb1, 0x4f,
	0xf2, 0xe5, 0xa6, 0x86, 0x0f, 0xbc, 0xa9, 0x6e, 0x77, 0x72, 0xc3, 0x8a, 0x5e, 0xa2, 0xbe, 0x09,
	0xd7, 0x37, 0x11, 0x8e, 0xfe, 0x8b, 0x30, 0xd3, 0xb2, 0xe4, 0x85, 0xcd, 0x60, 0x94, 0xd1, 0x25,
	0x1a, 0xd8, 0x45, 0xc7, 0x0f, 0x4c, 0x8e, 0xe9, 0xae, 0x39, 0x66, 0xc6, 0x75, 0x3b, 0x89, 0xed,
	0xe5, 0x1d, 0x31, 0x3a, 0x43, 0x07, 0x15, 0x54, 0x2b, 0x90, 0xcb, 0xb5, 0x68, 0xb8, 0x8e, 0x07,
	0xdb, 0x17, 0xb3, 0x6f, 0x0d, 0xef, 0x5a, 0x7d, 0xf4, 0x1e, 0x3d, 0x96, 0x70, 0x05, 0x12, 0xf8,
	0x1a, 0x3a, 0xc4, 0xc3, 0xed, 0x88, 0xa1, 0xf7, 0x18, 0xca, 0xf4, 0x4f, 0x80, 0x4e, 0xd6, 0xa2,
	0xc2, 0x5b, 0xdf, 0xe6, 0x74, 0xe8, 0x07, 0xb9, 0x68, 0xb9, 0x17, 0xc1, 0xe5, 0xa7, 0xce, 0x54,
	0x88, 0x0d, 0xe5, 0x05, 0x16, 0xb2, 0x48, 0x0a, 0xe0, 0xe6, 0x56, 0xf7, 0x0c, 0xeb, 0x52, 0xdd,
	0xf3, 0x8f, 0xbc, 0xf5, 0xd5, 0xf7, 0xb0, 0x77, 0x9e, 0x65, 0x3f, 0xc2, 0xe3, 0x73, 0x8b, 0xcc,
	0x98, 0xc2, 0xb6, 0x6c, 0xab, 0x79, 0x8a, 0x73, 0xa7, 0xfc, 0xe5, 0x34, 0x8b, 0x8c, 0xa9, 0x85,
	0xd7, 0x2c, 0xe6, 0xe9, 0xc2, 0x6b, 0x7e, 0x87, 0x27, 0xf6, 0x80, 0x90, 0x8c, 0x29, 0x42, 0xbc,
	0x8a, 0x90, 0x79, 0x4a, 0x88, 0xd7, 0xad, 0x06, 0x26, 0xec, 0xe9, 0xdf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xe0, 0xe4, 0x1a, 0x66, 0xcf, 0x03, 0x00, 0x00,
}
