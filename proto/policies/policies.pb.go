// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: proto/policies.proto

package policies

import (
	gobgp "github.com/p3rdy/bgpemu/proto/gobgp"
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

type PolicyDeployments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopoName          string              `protobuf:"bytes,1,opt,name=topo_name,json=topoName,proto3" json:"topo_name,omitempty"`
	PolicyDeployments []*PolicyDeployment `protobuf:"bytes,2,rep,name=policy_deployments,json=policyDeployments,proto3" json:"policy_deployments,omitempty"`
}

func (x *PolicyDeployments) Reset() {
	*x = PolicyDeployments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_policies_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyDeployments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyDeployments) ProtoMessage() {}

func (x *PolicyDeployments) ProtoReflect() protoreflect.Message {
	mi := &file_proto_policies_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyDeployments.ProtoReflect.Descriptor instead.
func (*PolicyDeployments) Descriptor() ([]byte, []int) {
	return file_proto_policies_proto_rawDescGZIP(), []int{0}
}

func (x *PolicyDeployments) GetTopoName() string {
	if x != nil {
		return x.TopoName
	}
	return ""
}

func (x *PolicyDeployments) GetPolicyDeployments() []*PolicyDeployment {
	if x != nil {
		return x.PolicyDeployments
	}
	return nil
}

type PolicyDeployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouterName  string                  `protobuf:"bytes,1,opt,name=router_name,json=routerName,proto3" json:"router_name,omitempty"`
	DefinedSets []*gobgp.DefinedSet     `protobuf:"bytes,2,rep,name=defined_sets,json=definedSets,proto3" json:"defined_sets,omitempty"`
	Policies    []*gobgp.Policy         `protobuf:"bytes,3,rep,name=policies,proto3" json:"policies,omitempty"`
	Assignments *gobgp.PolicyAssignment `protobuf:"bytes,4,opt,name=assignments,proto3" json:"assignments,omitempty"`
	PeerGroups  []*gobgp.PeerGroup      `protobuf:"bytes,5,rep,name=peer_groups,json=peerGroups,proto3" json:"peer_groups,omitempty"`
}

func (x *PolicyDeployment) Reset() {
	*x = PolicyDeployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_policies_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyDeployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyDeployment) ProtoMessage() {}

func (x *PolicyDeployment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_policies_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyDeployment.ProtoReflect.Descriptor instead.
func (*PolicyDeployment) Descriptor() ([]byte, []int) {
	return file_proto_policies_proto_rawDescGZIP(), []int{1}
}

func (x *PolicyDeployment) GetRouterName() string {
	if x != nil {
		return x.RouterName
	}
	return ""
}

func (x *PolicyDeployment) GetDefinedSets() []*gobgp.DefinedSet {
	if x != nil {
		return x.DefinedSets
	}
	return nil
}

func (x *PolicyDeployment) GetPolicies() []*gobgp.Policy {
	if x != nil {
		return x.Policies
	}
	return nil
}

func (x *PolicyDeployment) GetAssignments() *gobgp.PolicyAssignment {
	if x != nil {
		return x.Assignments
	}
	return nil
}

func (x *PolicyDeployment) GetPeerGroups() []*gobgp.PeerGroup {
	if x != nil {
		return x.PeerGroups
	}
	return nil
}

var File_proto_policies_proto protoreflect.FileDescriptor

var file_proto_policies_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x62, 0x67, 0x70, 0x2f, 0x67, 0x6f,
	0x62, 0x67, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x11, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x6f, 0x70, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x6f, 0x70, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x49, 0x0a, 0x12, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x11, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x82, 0x02, 0x0a, 0x10, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x0c,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x65, 0x64, 0x53, 0x65, 0x74, 0x52, 0x0b, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x53, 0x65,
	0x74, 0x73, 0x12, 0x29, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x39, 0x0a,
	0x0b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x61, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x31, 0x0a, 0x0b, 0x70, 0x65, 0x65, 0x72,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x61, 0x70, 0x69, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x0a, 0x70, 0x65, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x42, 0x10, 0x5a, 0x0e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_policies_proto_rawDescOnce sync.Once
	file_proto_policies_proto_rawDescData = file_proto_policies_proto_rawDesc
)

func file_proto_policies_proto_rawDescGZIP() []byte {
	file_proto_policies_proto_rawDescOnce.Do(func() {
		file_proto_policies_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_policies_proto_rawDescData)
	})
	return file_proto_policies_proto_rawDescData
}

var file_proto_policies_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_policies_proto_goTypes = []interface{}{
	(*PolicyDeployments)(nil),      // 0: policies.PolicyDeployments
	(*PolicyDeployment)(nil),       // 1: policies.PolicyDeployment
	(*gobgp.DefinedSet)(nil),       // 2: apipb.DefinedSet
	(*gobgp.Policy)(nil),           // 3: apipb.Policy
	(*gobgp.PolicyAssignment)(nil), // 4: apipb.PolicyAssignment
	(*gobgp.PeerGroup)(nil),        // 5: apipb.PeerGroup
}
var file_proto_policies_proto_depIdxs = []int32{
	1, // 0: policies.PolicyDeployments.policy_deployments:type_name -> policies.PolicyDeployment
	2, // 1: policies.PolicyDeployment.defined_sets:type_name -> apipb.DefinedSet
	3, // 2: policies.PolicyDeployment.policies:type_name -> apipb.Policy
	4, // 3: policies.PolicyDeployment.assignments:type_name -> apipb.PolicyAssignment
	5, // 4: policies.PolicyDeployment.peer_groups:type_name -> apipb.PeerGroup
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_policies_proto_init() }
func file_proto_policies_proto_init() {
	if File_proto_policies_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_policies_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyDeployments); i {
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
		file_proto_policies_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyDeployment); i {
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
			RawDescriptor: file_proto_policies_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_policies_proto_goTypes,
		DependencyIndexes: file_proto_policies_proto_depIdxs,
		MessageInfos:      file_proto_policies_proto_msgTypes,
	}.Build()
	File_proto_policies_proto = out.File
	file_proto_policies_proto_rawDesc = nil
	file_proto_policies_proto_goTypes = nil
	file_proto_policies_proto_depIdxs = nil
}
