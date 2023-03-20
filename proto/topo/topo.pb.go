// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: proto/topo.proto

package topo

import (
	knetopo "github.com/openconfig/kne/proto/topo"
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

type Type int32

const (
	Type_UNKNOWN Type = 0
	Type_HOST    Type = 1
	Type_BGPNODE Type = 2
	Type_SUBTOPO Type = 3
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "UNKNOWN",
		1: "HOST",
		2: "BGPNODE",
		3: "SUBTOPO",
	}
	Type_value = map[string]int32{
		"UNKNOWN": 0,
		"HOST":    1,
		"BGPNODE": 2,
		"SUBTOPO": 3,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_topo_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_proto_topo_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_proto_topo_proto_rawDescGZIP(), []int{0}
}

// Topology message defines what nodes and links will be created
// inside the mesh.
type Topology struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string                        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                                                                                                                       // Name of the topology - will be linked to the cluster name
	Nodes      []*Node                       `protobuf:"bytes,2,rep,name=nodes,proto3" json:"nodes,omitempty"`                                                                                                                     // List of nodes in the topology
	Links      []*knetopo.Link               `protobuf:"bytes,3,rep,name=links,proto3" json:"links,omitempty"`                                                                                                                     // connections between Nodes.
	ExportInts map[string]*InternalInterface `protobuf:"bytes,4,rep,name=export_ints,json=exportInts,proto3" json:"export_ints,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //List of interface to be exported out of a topo, optional.
}

func (x *Topology) Reset() {
	*x = Topology{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_topo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Topology) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topology) ProtoMessage() {}

func (x *Topology) ProtoReflect() protoreflect.Message {
	mi := &file_proto_topo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topology.ProtoReflect.Descriptor instead.
func (*Topology) Descriptor() ([]byte, []int) {
	return file_proto_topo_proto_rawDescGZIP(), []int{0}
}

func (x *Topology) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Topology) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *Topology) GetLinks() []*knetopo.Link {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *Topology) GetExportInts() map[string]*InternalInterface {
	if x != nil {
		return x.ExportInts
	}
	return nil
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string                      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type     Type                        `protobuf:"varint,2,opt,name=type,proto3,enum=topo.Type" json:"type,omitempty"`
	Path     *string                     `protobuf:"bytes,3,opt,name=path,proto3,oneof" json:"path,omitempty"` //Path of subtopo's definition file.
	IpAddr   map[string]string           `protobuf:"bytes,4,rep,name=ip_addr,json=ipAddr,proto3" json:"ip_addr,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Services map[uint32]*knetopo.Service `protobuf:"bytes,6,rep,name=services,proto3" json:"services,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Map of services to enable on the node.
	Config   *Config                     `protobuf:"bytes,7,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_topo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_proto_topo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_proto_topo_proto_rawDescGZIP(), []int{1}
}

func (x *Node) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Node) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_UNKNOWN
}

func (x *Node) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

func (x *Node) GetIpAddr() map[string]string {
	if x != nil {
		return x.IpAddr
	}
	return nil
}

func (x *Node) GetServices() map[uint32]*knetopo.Service {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *Node) GetConfig() *Config {
	if x != nil {
		return x.Config
	}
	return nil
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks            []*knetopo.Task                   `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	ExtraImages      map[string]string                 `protobuf:"bytes,2,rep,name=extra_images,json=extraImages,proto3" json:"extra_images,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //Other containers running on node
	ShareVolumes     []string                          `protobuf:"bytes,3,rep,name=share_volumes,json=shareVolumes,proto3" json:"share_volumes,omitempty"`
	ContainerVolumes map[string]*knetopo.PublicVolumes `protobuf:"bytes,4,rep,name=container_volumes,json=containerVolumes,proto3" json:"container_volumes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ConfigPath       string                            `protobuf:"bytes,6,opt,name=config_path,json=configPath,proto3" json:"config_path,omitempty"`
	ConfigFile       string                            `protobuf:"bytes,7,opt,name=config_file,json=configFile,proto3" json:"config_file,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_topo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_proto_topo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_proto_topo_proto_rawDescGZIP(), []int{2}
}

func (x *Config) GetTasks() []*knetopo.Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

func (x *Config) GetExtraImages() map[string]string {
	if x != nil {
		return x.ExtraImages
	}
	return nil
}

func (x *Config) GetShareVolumes() []string {
	if x != nil {
		return x.ShareVolumes
	}
	return nil
}

func (x *Config) GetContainerVolumes() map[string]*knetopo.PublicVolumes {
	if x != nil {
		return x.ContainerVolumes
	}
	return nil
}

func (x *Config) GetConfigPath() string {
	if x != nil {
		return x.ConfigPath
	}
	return ""
}

func (x *Config) GetConfigFile() string {
	if x != nil {
		return x.ConfigFile
	}
	return ""
}

// If this is a subtopo, interface message define how internal interfaces be exported.
type InternalInterface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node    string `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	NodeInt string `protobuf:"bytes,2,opt,name=node_int,json=nodeInt,proto3" json:"node_int,omitempty"`
}

func (x *InternalInterface) Reset() {
	*x = InternalInterface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_topo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalInterface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalInterface) ProtoMessage() {}

func (x *InternalInterface) ProtoReflect() protoreflect.Message {
	mi := &file_proto_topo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalInterface.ProtoReflect.Descriptor instead.
func (*InternalInterface) Descriptor() ([]byte, []int) {
	return file_proto_topo_proto_rawDescGZIP(), []int{3}
}

func (x *InternalInterface) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *InternalInterface) GetNodeInt() string {
	if x != nil {
		return x.NodeInt
	}
	return ""
}

var File_proto_topo_proto protoreflect.FileDescriptor

var file_proto_topo_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x74, 0x6f, 0x70, 0x6f, 0x1a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6b, 0x6e, 0x65, 0x74, 0x6f, 0x70, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x01,
	0x0a, 0x08, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x74, 0x6f, 0x70, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x12, 0x23, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x6b, 0x6e, 0x65, 0x74, 0x6f, 0x70, 0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05,
	0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x3f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x69, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x6f, 0x70,
	0x6f, 0x2e, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x49, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x49, 0x6e, 0x74, 0x73, 0x1a, 0x56, 0x0a, 0x0f, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x49, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x6f, 0x70,
	0x6f, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf3,
	0x02, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x74, 0x6f, 0x70, 0x6f,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x07, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x2e, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x69,
	0x70, 0x41, 0x64, 0x64, 0x72, 0x12, 0x34, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x6f,
	0x70, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x1a, 0x39, 0x0a, 0x0b, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4d, 0x0a, 0x0d,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x6b, 0x6e, 0x65, 0x74, 0x6f, 0x70, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x22, 0xc4, 0x03, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x23, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x6b, 0x6e, 0x65, 0x74, 0x6f, 0x70, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x12, 0x40, 0x0a, 0x0c, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x6f, 0x70,
	0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x12, 0x4f, 0x0a, 0x11, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x1a, 0x3e,
	0x0a, 0x10, 0x45, 0x78, 0x74, 0x72, 0x61, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x5b,
	0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x6e, 0x65, 0x74, 0x6f,
	0x70, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x42, 0x0a, 0x11, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x74, 0x2a,
	0x37, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x53, 0x54, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x42, 0x47, 0x50, 0x4e, 0x4f, 0x44, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x55, 0x42, 0x54, 0x4f, 0x50, 0x4f, 0x10, 0x03, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x33, 0x72, 0x64, 0x79, 0x2f, 0x62, 0x67, 0x70,
	0x65, 0x6d, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_topo_proto_rawDescOnce sync.Once
	file_proto_topo_proto_rawDescData = file_proto_topo_proto_rawDesc
)

func file_proto_topo_proto_rawDescGZIP() []byte {
	file_proto_topo_proto_rawDescOnce.Do(func() {
		file_proto_topo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_topo_proto_rawDescData)
	})
	return file_proto_topo_proto_rawDescData
}

var file_proto_topo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_topo_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_topo_proto_goTypes = []interface{}{
	(Type)(0),                     // 0: topo.Type
	(*Topology)(nil),              // 1: topo.Topology
	(*Node)(nil),                  // 2: topo.Node
	(*Config)(nil),                // 3: topo.Config
	(*InternalInterface)(nil),     // 4: topo.InternalInterface
	nil,                           // 5: topo.Topology.ExportIntsEntry
	nil,                           // 6: topo.Node.IpAddrEntry
	nil,                           // 7: topo.Node.ServicesEntry
	nil,                           // 8: topo.Config.ExtraImagesEntry
	nil,                           // 9: topo.Config.ContainerVolumesEntry
	(*knetopo.Link)(nil),          // 10: knetopo.Link
	(*knetopo.Task)(nil),          // 11: knetopo.Task
	(*knetopo.Service)(nil),       // 12: knetopo.Service
	(*knetopo.PublicVolumes)(nil), // 13: knetopo.PublicVolumes
}
var file_proto_topo_proto_depIdxs = []int32{
	2,  // 0: topo.Topology.nodes:type_name -> topo.Node
	10, // 1: topo.Topology.links:type_name -> knetopo.Link
	5,  // 2: topo.Topology.export_ints:type_name -> topo.Topology.ExportIntsEntry
	0,  // 3: topo.Node.type:type_name -> topo.Type
	6,  // 4: topo.Node.ip_addr:type_name -> topo.Node.IpAddrEntry
	7,  // 5: topo.Node.services:type_name -> topo.Node.ServicesEntry
	3,  // 6: topo.Node.config:type_name -> topo.Config
	11, // 7: topo.Config.tasks:type_name -> knetopo.Task
	8,  // 8: topo.Config.extra_images:type_name -> topo.Config.ExtraImagesEntry
	9,  // 9: topo.Config.container_volumes:type_name -> topo.Config.ContainerVolumesEntry
	4,  // 10: topo.Topology.ExportIntsEntry.value:type_name -> topo.InternalInterface
	12, // 11: topo.Node.ServicesEntry.value:type_name -> knetopo.Service
	13, // 12: topo.Config.ContainerVolumesEntry.value:type_name -> knetopo.PublicVolumes
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_proto_topo_proto_init() }
func file_proto_topo_proto_init() {
	if File_proto_topo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_topo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Topology); i {
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
		file_proto_topo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_proto_topo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_proto_topo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalInterface); i {
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
	file_proto_topo_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_topo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_topo_proto_goTypes,
		DependencyIndexes: file_proto_topo_proto_depIdxs,
		EnumInfos:         file_proto_topo_proto_enumTypes,
		MessageInfos:      file_proto_topo_proto_msgTypes,
	}.Build()
	File_proto_topo_proto = out.File
	file_proto_topo_proto_rawDesc = nil
	file_proto_topo_proto_goTypes = nil
	file_proto_topo_proto_depIdxs = nil
}
