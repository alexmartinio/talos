// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: resource/definitions/kubespan/kubespan.proto

package kubespan

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/siderolabs/talos/pkg/machinery/api/common"
	enums "github.com/siderolabs/talos/pkg/machinery/api/resource/definitions/enums"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ConfigSpec describes KubeSpan configuration..
type ConfigSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled                     bool     `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	ClusterId                   string   `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	SharedSecret                string   `protobuf:"bytes,3,opt,name=shared_secret,json=sharedSecret,proto3" json:"shared_secret,omitempty"`
	ForceRouting                bool     `protobuf:"varint,4,opt,name=force_routing,json=forceRouting,proto3" json:"force_routing,omitempty"`
	AdvertiseKubernetesNetworks bool     `protobuf:"varint,5,opt,name=advertise_kubernetes_networks,json=advertiseKubernetesNetworks,proto3" json:"advertise_kubernetes_networks,omitempty"`
	Mtu                         uint32   `protobuf:"varint,6,opt,name=mtu,proto3" json:"mtu,omitempty"`
	EndpointFilters             []string `protobuf:"bytes,7,rep,name=endpoint_filters,json=endpointFilters,proto3" json:"endpoint_filters,omitempty"`
	HarvestExtraEndpoints       bool     `protobuf:"varint,8,opt,name=harvest_extra_endpoints,json=harvestExtraEndpoints,proto3" json:"harvest_extra_endpoints,omitempty"`
}

func (x *ConfigSpec) Reset() {
	*x = ConfigSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_definitions_kubespan_kubespan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigSpec) ProtoMessage() {}

func (x *ConfigSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_kubespan_kubespan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigSpec.ProtoReflect.Descriptor instead.
func (*ConfigSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_kubespan_kubespan_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigSpec) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *ConfigSpec) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *ConfigSpec) GetSharedSecret() string {
	if x != nil {
		return x.SharedSecret
	}
	return ""
}

func (x *ConfigSpec) GetForceRouting() bool {
	if x != nil {
		return x.ForceRouting
	}
	return false
}

func (x *ConfigSpec) GetAdvertiseKubernetesNetworks() bool {
	if x != nil {
		return x.AdvertiseKubernetesNetworks
	}
	return false
}

func (x *ConfigSpec) GetMtu() uint32 {
	if x != nil {
		return x.Mtu
	}
	return 0
}

func (x *ConfigSpec) GetEndpointFilters() []string {
	if x != nil {
		return x.EndpointFilters
	}
	return nil
}

func (x *ConfigSpec) GetHarvestExtraEndpoints() bool {
	if x != nil {
		return x.HarvestExtraEndpoints
	}
	return false
}

// EndpointSpec describes Endpoint state.
type EndpointSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AffiliateId string            `protobuf:"bytes,1,opt,name=affiliate_id,json=affiliateId,proto3" json:"affiliate_id,omitempty"`
	Endpoint    *common.NetIPPort `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *EndpointSpec) Reset() {
	*x = EndpointSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_definitions_kubespan_kubespan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointSpec) ProtoMessage() {}

func (x *EndpointSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_kubespan_kubespan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointSpec.ProtoReflect.Descriptor instead.
func (*EndpointSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_kubespan_kubespan_proto_rawDescGZIP(), []int{1}
}

func (x *EndpointSpec) GetAffiliateId() string {
	if x != nil {
		return x.AffiliateId
	}
	return ""
}

func (x *EndpointSpec) GetEndpoint() *common.NetIPPort {
	if x != nil {
		return x.Endpoint
	}
	return nil
}

// IdentitySpec describes KubeSpan keys and address.
//
// Note: IdentitySpec is persisted on disk in the STATE partition,
// so YAML serialization should be kept backwards compatible.
type IdentitySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address    *common.NetIPPrefix `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Subnet     *common.NetIPPrefix `protobuf:"bytes,2,opt,name=subnet,proto3" json:"subnet,omitempty"`
	PrivateKey string              `protobuf:"bytes,3,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	PublicKey  string              `protobuf:"bytes,4,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *IdentitySpec) Reset() {
	*x = IdentitySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_definitions_kubespan_kubespan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentitySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentitySpec) ProtoMessage() {}

func (x *IdentitySpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_kubespan_kubespan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentitySpec.ProtoReflect.Descriptor instead.
func (*IdentitySpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_kubespan_kubespan_proto_rawDescGZIP(), []int{2}
}

func (x *IdentitySpec) GetAddress() *common.NetIPPrefix {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *IdentitySpec) GetSubnet() *common.NetIPPrefix {
	if x != nil {
		return x.Subnet
	}
	return nil
}

func (x *IdentitySpec) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *IdentitySpec) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

// PeerSpecSpec describes PeerSpec state.
type PeerSpecSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address    *common.NetIP         `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	AllowedIps []*common.NetIPPrefix `protobuf:"bytes,2,rep,name=allowed_ips,json=allowedIps,proto3" json:"allowed_ips,omitempty"`
	Endpoints  []*common.NetIPPort   `protobuf:"bytes,3,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	Label      string                `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *PeerSpecSpec) Reset() {
	*x = PeerSpecSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_definitions_kubespan_kubespan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerSpecSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerSpecSpec) ProtoMessage() {}

func (x *PeerSpecSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_kubespan_kubespan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerSpecSpec.ProtoReflect.Descriptor instead.
func (*PeerSpecSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_kubespan_kubespan_proto_rawDescGZIP(), []int{3}
}

func (x *PeerSpecSpec) GetAddress() *common.NetIP {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *PeerSpecSpec) GetAllowedIps() []*common.NetIPPrefix {
	if x != nil {
		return x.AllowedIps
	}
	return nil
}

func (x *PeerSpecSpec) GetEndpoints() []*common.NetIPPort {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *PeerSpecSpec) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

// PeerStatusSpec describes PeerStatus state.
type PeerStatusSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint           *common.NetIPPort       `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Label              string                  `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	State              enums.KubespanPeerState `protobuf:"varint,3,opt,name=state,proto3,enum=talos.resource.definitions.enums.KubespanPeerState" json:"state,omitempty"`
	ReceiveBytes       int64                   `protobuf:"varint,4,opt,name=receive_bytes,json=receiveBytes,proto3" json:"receive_bytes,omitempty"`
	TransmitBytes      int64                   `protobuf:"varint,5,opt,name=transmit_bytes,json=transmitBytes,proto3" json:"transmit_bytes,omitempty"`
	LastHandshakeTime  *timestamppb.Timestamp  `protobuf:"bytes,6,opt,name=last_handshake_time,json=lastHandshakeTime,proto3" json:"last_handshake_time,omitempty"`
	LastUsedEndpoint   *common.NetIPPort       `protobuf:"bytes,7,opt,name=last_used_endpoint,json=lastUsedEndpoint,proto3" json:"last_used_endpoint,omitempty"`
	LastEndpointChange *timestamppb.Timestamp  `protobuf:"bytes,8,opt,name=last_endpoint_change,json=lastEndpointChange,proto3" json:"last_endpoint_change,omitempty"`
}

func (x *PeerStatusSpec) Reset() {
	*x = PeerStatusSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_definitions_kubespan_kubespan_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerStatusSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerStatusSpec) ProtoMessage() {}

func (x *PeerStatusSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_kubespan_kubespan_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerStatusSpec.ProtoReflect.Descriptor instead.
func (*PeerStatusSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_kubespan_kubespan_proto_rawDescGZIP(), []int{4}
}

func (x *PeerStatusSpec) GetEndpoint() *common.NetIPPort {
	if x != nil {
		return x.Endpoint
	}
	return nil
}

func (x *PeerStatusSpec) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *PeerStatusSpec) GetState() enums.KubespanPeerState {
	if x != nil {
		return x.State
	}
	return enums.KubespanPeerState(0)
}

func (x *PeerStatusSpec) GetReceiveBytes() int64 {
	if x != nil {
		return x.ReceiveBytes
	}
	return 0
}

func (x *PeerStatusSpec) GetTransmitBytes() int64 {
	if x != nil {
		return x.TransmitBytes
	}
	return 0
}

func (x *PeerStatusSpec) GetLastHandshakeTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastHandshakeTime
	}
	return nil
}

func (x *PeerStatusSpec) GetLastUsedEndpoint() *common.NetIPPort {
	if x != nil {
		return x.LastUsedEndpoint
	}
	return nil
}

func (x *PeerStatusSpec) GetLastEndpointChange() *timestamppb.Timestamp {
	if x != nil {
		return x.LastEndpointChange
	}
	return nil
}

var File_resource_definitions_kubespan_kubespan_proto protoreflect.FileDescriptor

var file_resource_definitions_kubespan_kubespan_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x73, 0x70, 0x61, 0x6e, 0x2f,
	0x6b, 0x75, 0x62, 0x65, 0x73, 0x70, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23,
	0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x73,
	0x70, 0x61, 0x6e, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc8, 0x02, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x52, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x42, 0x0a, 0x1d, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65,
	0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1b, 0x61, 0x64, 0x76, 0x65,
	0x72, 0x74, 0x69, 0x73, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x74, 0x75, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6d, 0x74, 0x75, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x36, 0x0a, 0x17, 0x68, 0x61, 0x72, 0x76, 0x65, 0x73, 0x74, 0x5f,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x68, 0x61, 0x72, 0x76, 0x65, 0x73, 0x74, 0x45, 0x78,
	0x74, 0x72, 0x61, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x60, 0x0a, 0x0c,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12,
	0x2d, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x65, 0x74, 0x49, 0x50,
	0x50, 0x6f, 0x72, 0x74, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0xaa,
	0x01, 0x0a, 0x0c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x2d, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x65, 0x74, 0x49, 0x50, 0x50,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2b,
	0x0a, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x65, 0x74, 0x49, 0x50, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x52, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0xb4, 0x01, 0x0a, 0x0c,
	0x50, 0x65, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x53, 0x70, 0x65, 0x63, 0x12, 0x27, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x65, 0x74, 0x49, 0x50, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x34, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x5f, 0x69, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x65, 0x74, 0x49, 0x50, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52,
	0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x49, 0x70, 0x73, 0x12, 0x2f, 0x0a, 0x09, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x65, 0x74, 0x49, 0x50, 0x50, 0x6f, 0x72,
	0x74, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x22, 0xc7, 0x03, 0x0a, 0x0e, 0x50, 0x65, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x2d, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x4e, 0x65, 0x74, 0x49, 0x50, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x49, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x74, 0x61, 0x6c, 0x6f,
	0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x4b, 0x75, 0x62,
	0x65, 0x73, 0x70, 0x61, 0x6e, 0x50, 0x65, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x4a, 0x0a, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68,
	0x61, 0x6b, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x6c, 0x61, 0x73, 0x74,
	0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3f, 0x0a,
	0x12, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x4e, 0x65, 0x74, 0x49, 0x50, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x10, 0x6c, 0x61,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x64, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x4c,
	0x0a, 0x14, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x4d, 0x5a, 0x4b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x72,
	0x6f, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x73, 0x70, 0x61, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_resource_definitions_kubespan_kubespan_proto_rawDescOnce sync.Once
	file_resource_definitions_kubespan_kubespan_proto_rawDescData = file_resource_definitions_kubespan_kubespan_proto_rawDesc
)

func file_resource_definitions_kubespan_kubespan_proto_rawDescGZIP() []byte {
	file_resource_definitions_kubespan_kubespan_proto_rawDescOnce.Do(func() {
		file_resource_definitions_kubespan_kubespan_proto_rawDescData = protoimpl.X.CompressGZIP(file_resource_definitions_kubespan_kubespan_proto_rawDescData)
	})
	return file_resource_definitions_kubespan_kubespan_proto_rawDescData
}

var file_resource_definitions_kubespan_kubespan_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_resource_definitions_kubespan_kubespan_proto_goTypes = []interface{}{
	(*ConfigSpec)(nil),            // 0: talos.resource.definitions.kubespan.ConfigSpec
	(*EndpointSpec)(nil),          // 1: talos.resource.definitions.kubespan.EndpointSpec
	(*IdentitySpec)(nil),          // 2: talos.resource.definitions.kubespan.IdentitySpec
	(*PeerSpecSpec)(nil),          // 3: talos.resource.definitions.kubespan.PeerSpecSpec
	(*PeerStatusSpec)(nil),        // 4: talos.resource.definitions.kubespan.PeerStatusSpec
	(*common.NetIPPort)(nil),      // 5: common.NetIPPort
	(*common.NetIPPrefix)(nil),    // 6: common.NetIPPrefix
	(*common.NetIP)(nil),          // 7: common.NetIP
	(enums.KubespanPeerState)(0),  // 8: talos.resource.definitions.enums.KubespanPeerState
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_resource_definitions_kubespan_kubespan_proto_depIdxs = []int32{
	5,  // 0: talos.resource.definitions.kubespan.EndpointSpec.endpoint:type_name -> common.NetIPPort
	6,  // 1: talos.resource.definitions.kubespan.IdentitySpec.address:type_name -> common.NetIPPrefix
	6,  // 2: talos.resource.definitions.kubespan.IdentitySpec.subnet:type_name -> common.NetIPPrefix
	7,  // 3: talos.resource.definitions.kubespan.PeerSpecSpec.address:type_name -> common.NetIP
	6,  // 4: talos.resource.definitions.kubespan.PeerSpecSpec.allowed_ips:type_name -> common.NetIPPrefix
	5,  // 5: talos.resource.definitions.kubespan.PeerSpecSpec.endpoints:type_name -> common.NetIPPort
	5,  // 6: talos.resource.definitions.kubespan.PeerStatusSpec.endpoint:type_name -> common.NetIPPort
	8,  // 7: talos.resource.definitions.kubespan.PeerStatusSpec.state:type_name -> talos.resource.definitions.enums.KubespanPeerState
	9,  // 8: talos.resource.definitions.kubespan.PeerStatusSpec.last_handshake_time:type_name -> google.protobuf.Timestamp
	5,  // 9: talos.resource.definitions.kubespan.PeerStatusSpec.last_used_endpoint:type_name -> common.NetIPPort
	9,  // 10: talos.resource.definitions.kubespan.PeerStatusSpec.last_endpoint_change:type_name -> google.protobuf.Timestamp
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_resource_definitions_kubespan_kubespan_proto_init() }
func file_resource_definitions_kubespan_kubespan_proto_init() {
	if File_resource_definitions_kubespan_kubespan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resource_definitions_kubespan_kubespan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigSpec); i {
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
		file_resource_definitions_kubespan_kubespan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointSpec); i {
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
		file_resource_definitions_kubespan_kubespan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentitySpec); i {
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
		file_resource_definitions_kubespan_kubespan_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerSpecSpec); i {
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
		file_resource_definitions_kubespan_kubespan_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerStatusSpec); i {
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
			RawDescriptor: file_resource_definitions_kubespan_kubespan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resource_definitions_kubespan_kubespan_proto_goTypes,
		DependencyIndexes: file_resource_definitions_kubespan_kubespan_proto_depIdxs,
		MessageInfos:      file_resource_definitions_kubespan_kubespan_proto_msgTypes,
	}.Build()
	File_resource_definitions_kubespan_kubespan_proto = out.File
	file_resource_definitions_kubespan_kubespan_proto_rawDesc = nil
	file_resource_definitions_kubespan_kubespan_proto_goTypes = nil
	file_resource_definitions_kubespan_kubespan_proto_depIdxs = nil
}
