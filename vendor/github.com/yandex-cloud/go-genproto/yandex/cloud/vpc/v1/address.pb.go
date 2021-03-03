// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: yandex/cloud/vpc/v1/address.proto

package vpc

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// An Address resource. For more information, see [Address](/docs/vpc/concepts/address).
type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the address. Generated at creation time.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the address belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the address.
	// The name is unique within the folder.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the address.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Resource labels as `key:value` pairs.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// External ipv4 address specification.
	//
	// Types that are assignable to Address:
	//	*Address_ExternalIpv4Address
	Address isAddress_Address `protobuf_oneof:"address"`
	// Specifies if address is reserved or not.
	Reserved bool `protobuf:"varint,15,opt,name=reserved,proto3" json:"reserved,omitempty"`
	// Specifies if address is used or not.
	Used bool `protobuf:"varint,16,opt,name=used,proto3" json:"used,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_vpc_v1_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_vpc_v1_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_vpc_v1_address_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Address) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *Address) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Address) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Address) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Address) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (m *Address) GetAddress() isAddress_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (x *Address) GetExternalIpv4Address() *ExternalIpv4Address {
	if x, ok := x.GetAddress().(*Address_ExternalIpv4Address); ok {
		return x.ExternalIpv4Address
	}
	return nil
}

func (x *Address) GetReserved() bool {
	if x != nil {
		return x.Reserved
	}
	return false
}

func (x *Address) GetUsed() bool {
	if x != nil {
		return x.Used
	}
	return false
}

type isAddress_Address interface {
	isAddress_Address()
}

type Address_ExternalIpv4Address struct {
	ExternalIpv4Address *ExternalIpv4Address `protobuf:"bytes,7,opt,name=external_ipv4_address,json=externalIpv4Address,proto3,oneof"`
}

func (*Address_ExternalIpv4Address) isAddress_Address() {}

type ExternalIpv4Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value of address.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Availability zone from which the address will be allocated.
	ZoneId string `protobuf:"bytes,2,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	// Parameters of the allocated address, for example DDoS Protection.
	Requirements *AddressRequirements `protobuf:"bytes,3,opt,name=requirements,proto3" json:"requirements,omitempty"`
}

func (x *ExternalIpv4Address) Reset() {
	*x = ExternalIpv4Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_vpc_v1_address_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalIpv4Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalIpv4Address) ProtoMessage() {}

func (x *ExternalIpv4Address) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_vpc_v1_address_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalIpv4Address.ProtoReflect.Descriptor instead.
func (*ExternalIpv4Address) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_vpc_v1_address_proto_rawDescGZIP(), []int{1}
}

func (x *ExternalIpv4Address) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ExternalIpv4Address) GetZoneId() string {
	if x != nil {
		return x.ZoneId
	}
	return ""
}

func (x *ExternalIpv4Address) GetRequirements() *AddressRequirements {
	if x != nil {
		return x.Requirements
	}
	return nil
}

type AddressRequirements struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DDoS protection provider ID.
	DdosProtectionProvider string `protobuf:"bytes,1,opt,name=ddos_protection_provider,json=ddosProtectionProvider,proto3" json:"ddos_protection_provider,omitempty"`
	// Capability to send SMTP traffic.
	OutgoingSmtpCapability string `protobuf:"bytes,2,opt,name=outgoing_smtp_capability,json=outgoingSmtpCapability,proto3" json:"outgoing_smtp_capability,omitempty"`
}

func (x *AddressRequirements) Reset() {
	*x = AddressRequirements{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_vpc_v1_address_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressRequirements) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressRequirements) ProtoMessage() {}

func (x *AddressRequirements) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_vpc_v1_address_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressRequirements.ProtoReflect.Descriptor instead.
func (*AddressRequirements) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_vpc_v1_address_proto_rawDescGZIP(), []int{2}
}

func (x *AddressRequirements) GetDdosProtectionProvider() string {
	if x != nil {
		return x.DdosProtectionProvider
	}
	return ""
}

func (x *AddressRequirements) GetOutgoingSmtpCapability() string {
	if x != nil {
		return x.OutgoingSmtpCapability
	}
	return ""
}

var File_yandex_cloud_vpc_v1_address_proto protoreflect.FileDescriptor

var file_yandex_cloud_vpc_v1_address_proto_rawDesc = []byte{
	0x0a, 0x21, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x03, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x5e, 0x0a, 0x15, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x49, 0x70, 0x76, 0x34, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52,
	0x13, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x70, 0x76, 0x34, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x64, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x0f, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x04, 0xc0, 0xc1, 0x31, 0x01,
	0x22, 0x96, 0x01, 0x0a, 0x13, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x70, 0x76,
	0x34, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x7a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x4c, 0x0a, 0x0c, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x0c, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x13, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x38, 0x0a, 0x18, 0x64, 0x64, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x16, 0x64, 0x64, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x18, 0x6f,
	0x75, 0x74, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x6d, 0x74, 0x70, 0x5f, 0x63, 0x61, 0x70,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x6f,
	0x75, 0x74, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x53, 0x6d, 0x74, 0x70, 0x43, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x42, 0x56, 0x0a, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31,
	0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_vpc_v1_address_proto_rawDescOnce sync.Once
	file_yandex_cloud_vpc_v1_address_proto_rawDescData = file_yandex_cloud_vpc_v1_address_proto_rawDesc
)

func file_yandex_cloud_vpc_v1_address_proto_rawDescGZIP() []byte {
	file_yandex_cloud_vpc_v1_address_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_vpc_v1_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_vpc_v1_address_proto_rawDescData)
	})
	return file_yandex_cloud_vpc_v1_address_proto_rawDescData
}

var file_yandex_cloud_vpc_v1_address_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_yandex_cloud_vpc_v1_address_proto_goTypes = []interface{}{
	(*Address)(nil),             // 0: yandex.cloud.vpc.v1.Address
	(*ExternalIpv4Address)(nil), // 1: yandex.cloud.vpc.v1.ExternalIpv4Address
	(*AddressRequirements)(nil), // 2: yandex.cloud.vpc.v1.AddressRequirements
	nil,                         // 3: yandex.cloud.vpc.v1.Address.LabelsEntry
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_yandex_cloud_vpc_v1_address_proto_depIdxs = []int32{
	4, // 0: yandex.cloud.vpc.v1.Address.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: yandex.cloud.vpc.v1.Address.labels:type_name -> yandex.cloud.vpc.v1.Address.LabelsEntry
	1, // 2: yandex.cloud.vpc.v1.Address.external_ipv4_address:type_name -> yandex.cloud.vpc.v1.ExternalIpv4Address
	2, // 3: yandex.cloud.vpc.v1.ExternalIpv4Address.requirements:type_name -> yandex.cloud.vpc.v1.AddressRequirements
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_yandex_cloud_vpc_v1_address_proto_init() }
func file_yandex_cloud_vpc_v1_address_proto_init() {
	if File_yandex_cloud_vpc_v1_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_vpc_v1_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_yandex_cloud_vpc_v1_address_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalIpv4Address); i {
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
		file_yandex_cloud_vpc_v1_address_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressRequirements); i {
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
	file_yandex_cloud_vpc_v1_address_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Address_ExternalIpv4Address)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_vpc_v1_address_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_vpc_v1_address_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_vpc_v1_address_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_vpc_v1_address_proto_msgTypes,
	}.Build()
	File_yandex_cloud_vpc_v1_address_proto = out.File
	file_yandex_cloud_vpc_v1_address_proto_rawDesc = nil
	file_yandex_cloud_vpc_v1_address_proto_goTypes = nil
	file_yandex_cloud_vpc_v1_address_proto_depIdxs = nil
}
