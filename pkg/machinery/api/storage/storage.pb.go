// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: storage/storage.proto

package storage

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	common "github.com/siderolabs/talos/pkg/machinery/api/common"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Disk_DiskType int32

const (
	Disk_UNKNOWN Disk_DiskType = 0
	Disk_SSD     Disk_DiskType = 1
	Disk_HDD     Disk_DiskType = 2
	Disk_NVME    Disk_DiskType = 3
	Disk_SD      Disk_DiskType = 4
)

// Enum value maps for Disk_DiskType.
var (
	Disk_DiskType_name = map[int32]string{
		0: "UNKNOWN",
		1: "SSD",
		2: "HDD",
		3: "NVME",
		4: "SD",
	}
	Disk_DiskType_value = map[string]int32{
		"UNKNOWN": 0,
		"SSD":     1,
		"HDD":     2,
		"NVME":    3,
		"SD":      4,
	}
)

func (x Disk_DiskType) Enum() *Disk_DiskType {
	p := new(Disk_DiskType)
	*p = x
	return p
}

func (x Disk_DiskType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Disk_DiskType) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_storage_proto_enumTypes[0].Descriptor()
}

func (Disk_DiskType) Type() protoreflect.EnumType {
	return &file_storage_storage_proto_enumTypes[0]
}

func (x Disk_DiskType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Disk_DiskType.Descriptor instead.
func (Disk_DiskType) EnumDescriptor() ([]byte, []int) {
	return file_storage_storage_proto_rawDescGZIP(), []int{0, 0}
}

// Disk represents a disk.
type Disk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Size indicates the disk size in bytes.
	Size uint64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// Model idicates the disk model.
	Model string `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	// DeviceName indicates the disk name (e.g. `sda`).
	DeviceName string `protobuf:"bytes,3,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	// Name as in `/sys/block/<dev>/device/name`.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Serial as in `/sys/block/<dev>/device/serial`.
	Serial string `protobuf:"bytes,5,opt,name=serial,proto3" json:"serial,omitempty"`
	// Modalias as in `/sys/block/<dev>/device/modalias`.
	Modalias string `protobuf:"bytes,6,opt,name=modalias,proto3" json:"modalias,omitempty"`
	// Uuid as in `/sys/block/<dev>/device/uuid`.
	Uuid string `protobuf:"bytes,7,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Wwid as in `/sys/block/<dev>/device/wwid`.
	Wwid string `protobuf:"bytes,8,opt,name=wwid,proto3" json:"wwid,omitempty"`
	// Type is a type of the disk: nvme, ssd, hdd, sd card.
	Type Disk_DiskType `protobuf:"varint,9,opt,name=type,proto3,enum=storage.Disk_DiskType" json:"type,omitempty"`
	// BusPath is the bus path of the disk.
	BusPath string `protobuf:"bytes,10,opt,name=bus_path,json=busPath,proto3" json:"bus_path,omitempty"`
	// SystemDisk indicates that the disk is used as Talos system disk.
	SystemDisk bool `protobuf:"varint,11,opt,name=system_disk,json=systemDisk,proto3" json:"system_disk,omitempty"`
	// Subsystem is the symlink path in the `/sys/block/<dev>/subsystem`.
	Subsystem string `protobuf:"bytes,12,opt,name=subsystem,proto3" json:"subsystem,omitempty"`
	// Readonly specifies if the disk is read only.
	Readonly bool `protobuf:"varint,13,opt,name=readonly,proto3" json:"readonly,omitempty"`
}

func (x *Disk) Reset() {
	*x = Disk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Disk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Disk) ProtoMessage() {}

func (x *Disk) ProtoReflect() protoreflect.Message {
	mi := &file_storage_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Disk.ProtoReflect.Descriptor instead.
func (*Disk) Descriptor() ([]byte, []int) {
	return file_storage_storage_proto_rawDescGZIP(), []int{0}
}

func (x *Disk) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Disk) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Disk) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *Disk) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Disk) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *Disk) GetModalias() string {
	if x != nil {
		return x.Modalias
	}
	return ""
}

func (x *Disk) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Disk) GetWwid() string {
	if x != nil {
		return x.Wwid
	}
	return ""
}

func (x *Disk) GetType() Disk_DiskType {
	if x != nil {
		return x.Type
	}
	return Disk_UNKNOWN
}

func (x *Disk) GetBusPath() string {
	if x != nil {
		return x.BusPath
	}
	return ""
}

func (x *Disk) GetSystemDisk() bool {
	if x != nil {
		return x.SystemDisk
	}
	return false
}

func (x *Disk) GetSubsystem() string {
	if x != nil {
		return x.Subsystem
	}
	return ""
}

func (x *Disk) GetReadonly() bool {
	if x != nil {
		return x.Readonly
	}
	return false
}

// DisksResponse represents the response of the `Disks` RPC.
type Disks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *common.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Disks    []*Disk          `protobuf:"bytes,2,rep,name=disks,proto3" json:"disks,omitempty"`
}

func (x *Disks) Reset() {
	*x = Disks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Disks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Disks) ProtoMessage() {}

func (x *Disks) ProtoReflect() protoreflect.Message {
	mi := &file_storage_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Disks.ProtoReflect.Descriptor instead.
func (*Disks) Descriptor() ([]byte, []int) {
	return file_storage_storage_proto_rawDescGZIP(), []int{1}
}

func (x *Disks) GetMetadata() *common.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Disks) GetDisks() []*Disk {
	if x != nil {
		return x.Disks
	}
	return nil
}

type DisksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*Disks `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *DisksResponse) Reset() {
	*x = DisksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisksResponse) ProtoMessage() {}

func (x *DisksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisksResponse.ProtoReflect.Descriptor instead.
func (*DisksResponse) Descriptor() ([]byte, []int) {
	return file_storage_storage_proto_rawDescGZIP(), []int{2}
}

func (x *DisksResponse) GetMessages() []*Disks {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_storage_storage_proto protoreflect.FileDescriptor

var file_storage_storage_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa0, 0x03, 0x0a, 0x04, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x77, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x77, 0x77, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44,
	0x69, 0x73, 0x6b, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x73, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x73, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x64, 0x69, 0x73, 0x6b, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x61, 0x64, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x72, 0x65, 0x61, 0x64, 0x6f, 0x6e, 0x6c, 0x79, 0x22, 0x3b, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x6b,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x53, 0x44, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x48, 0x44,
	0x44, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x56, 0x4d, 0x45, 0x10, 0x03, 0x12, 0x06, 0x0a,
	0x02, 0x53, 0x44, 0x10, 0x04, 0x22, 0x5a, 0x0a, 0x05, 0x44, 0x69, 0x73, 0x6b, 0x73, 0x12, 0x2c,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05,
	0x64, 0x69, 0x73, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x05, 0x64, 0x69, 0x73, 0x6b,
	0x73, 0x22, 0x3b, 0x0a, 0x0d, 0x44, 0x69, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44,
	0x69, 0x73, 0x6b, 0x73, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x32, 0x49,
	0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x37, 0x0a, 0x05, 0x44, 0x69, 0x73, 0x6b, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x16, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x6b,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x61,
	0x62, 0x73, 0x2f, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_storage_proto_rawDescOnce sync.Once
	file_storage_storage_proto_rawDescData = file_storage_storage_proto_rawDesc
)

func file_storage_storage_proto_rawDescGZIP() []byte {
	file_storage_storage_proto_rawDescOnce.Do(func() {
		file_storage_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_storage_proto_rawDescData)
	})
	return file_storage_storage_proto_rawDescData
}

var file_storage_storage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_storage_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_storage_storage_proto_goTypes = []interface{}{
	(Disk_DiskType)(0),      // 0: storage.Disk.DiskType
	(*Disk)(nil),            // 1: storage.Disk
	(*Disks)(nil),           // 2: storage.Disks
	(*DisksResponse)(nil),   // 3: storage.DisksResponse
	(*common.Metadata)(nil), // 4: common.Metadata
	(*emptypb.Empty)(nil),   // 5: google.protobuf.Empty
}
var file_storage_storage_proto_depIdxs = []int32{
	0, // 0: storage.Disk.type:type_name -> storage.Disk.DiskType
	4, // 1: storage.Disks.metadata:type_name -> common.Metadata
	1, // 2: storage.Disks.disks:type_name -> storage.Disk
	2, // 3: storage.DisksResponse.messages:type_name -> storage.Disks
	5, // 4: storage.StorageService.Disks:input_type -> google.protobuf.Empty
	3, // 5: storage.StorageService.Disks:output_type -> storage.DisksResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_storage_storage_proto_init() }
func file_storage_storage_proto_init() {
	if File_storage_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storage_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Disk); i {
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
		file_storage_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Disks); i {
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
		file_storage_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisksResponse); i {
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
			RawDescriptor: file_storage_storage_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_storage_storage_proto_goTypes,
		DependencyIndexes: file_storage_storage_proto_depIdxs,
		EnumInfos:         file_storage_storage_proto_enumTypes,
		MessageInfos:      file_storage_storage_proto_msgTypes,
	}.Build()
	File_storage_storage_proto = out.File
	file_storage_storage_proto_rawDesc = nil
	file_storage_storage_proto_goTypes = nil
	file_storage_storage_proto_depIdxs = nil
}
