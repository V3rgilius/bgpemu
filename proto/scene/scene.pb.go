// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: proto/scene.proto

package scene

import (
	_ "github.com/p3rdy/bgpemu/proto/gobgp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Step_Type int32

const (
	Step_UNKNOWN Step_Type = 0
	Step_CMDS    Step_Type = 1
	Step_WAIT    Step_Type = 2
	Step_FT      Step_Type = 3
)

// Enum value maps for Step_Type.
var (
	Step_Type_name = map[int32]string{
		0: "UNKNOWN",
		1: "CMDS",
		2: "WAIT",
		3: "FT",
	}
	Step_Type_value = map[string]int32{
		"UNKNOWN": 0,
		"CMDS":    1,
		"WAIT":    2,
		"FT":      3,
	}
)

func (x Step_Type) Enum() *Step_Type {
	p := new(Step_Type)
	*p = x
	return p
}

func (x Step_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Step_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_scene_proto_enumTypes[0].Descriptor()
}

func (Step_Type) Type() protoreflect.EnumType {
	return &file_proto_scene_proto_enumTypes[0]
}

func (x Step_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Step_Type.Descriptor instead.
func (Step_Type) EnumDescriptor() ([]byte, []int) {
	return file_proto_scene_proto_rawDescGZIP(), []int{2, 0}
}

type Scene struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopoName  string      `protobuf:"bytes,1,opt,name=topo_name,json=topoName,proto3" json:"topo_name,omitempty"`
	Behaviors []*Behavior `protobuf:"bytes,2,rep,name=behaviors,proto3" json:"behaviors,omitempty"`
}

func (x *Scene) Reset() {
	*x = Scene{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scene_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scene) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scene) ProtoMessage() {}

func (x *Scene) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scene_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scene.ProtoReflect.Descriptor instead.
func (*Scene) Descriptor() ([]byte, []int) {
	return file_proto_scene_proto_rawDescGZIP(), []int{0}
}

func (x *Scene) GetTopoName() string {
	if x != nil {
		return x.TopoName
	}
	return ""
}

func (x *Scene) GetBehaviors() []*Behavior {
	if x != nil {
		return x.Behaviors
	}
	return nil
}

type Behavior struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DeviceName string  `protobuf:"bytes,2,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	IsAsync    bool    `protobuf:"varint,3,opt,name=isAsync,proto3" json:"isAsync,omitempty"`
	Steps      []*Step `protobuf:"bytes,4,rep,name=steps,proto3" json:"steps,omitempty"`
}

func (x *Behavior) Reset() {
	*x = Behavior{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scene_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Behavior) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Behavior) ProtoMessage() {}

func (x *Behavior) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scene_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Behavior.ProtoReflect.Descriptor instead.
func (*Behavior) Descriptor() ([]byte, []int) {
	return file_proto_scene_proto_rawDescGZIP(), []int{1}
}

func (x *Behavior) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Behavior) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *Behavior) GetIsAsync() bool {
	if x != nil {
		return x.IsAsync
	}
	return false
}

func (x *Behavior) GetSteps() []*Step {
	if x != nil {
		return x.Steps
	}
	return nil
}

type Step struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Step_Type  `protobuf:"varint,1,opt,name=type,proto3,enum=scene.Step_Type" json:"type,omitempty"`
	Name string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Body *anypb.Any `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Step) Reset() {
	*x = Step{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scene_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Step) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Step) ProtoMessage() {}

func (x *Step) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scene_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Step.ProtoReflect.Descriptor instead.
func (*Step) Descriptor() ([]byte, []int) {
	return file_proto_scene_proto_rawDescGZIP(), []int{2}
}

func (x *Step) GetType() Step_Type {
	if x != nil {
		return x.Type
	}
	return Step_UNKNOWN
}

func (x *Step) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Step) GetBody() *anypb.Any {
	if x != nil {
		return x.Body
	}
	return nil
}

type Commands struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmds []string `protobuf:"bytes,1,rep,name=cmds,proto3" json:"cmds,omitempty"`
}

func (x *Commands) Reset() {
	*x = Commands{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scene_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commands) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commands) ProtoMessage() {}

func (x *Commands) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scene_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commands.ProtoReflect.Descriptor instead.
func (*Commands) Descriptor() ([]byte, []int) {
	return file_proto_scene_proto_rawDescGZIP(), []int{3}
}

func (x *Commands) GetCmds() []string {
	if x != nil {
		return x.Cmds
	}
	return nil
}

type Wait struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time      uint32 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Timestamp uint32 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Wait) Reset() {
	*x = Wait{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scene_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wait) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wait) ProtoMessage() {}

func (x *Wait) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scene_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wait.ProtoReflect.Descriptor instead.
func (*Wait) Descriptor() ([]byte, []int) {
	return file_proto_scene_proto_rawDescGZIP(), []int{4}
}

func (x *Wait) GetTime() uint32 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Wait) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type FileTrans struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Src string `protobuf:"bytes,1,opt,name=src,proto3" json:"src,omitempty"`
	Des string `protobuf:"bytes,2,opt,name=des,proto3" json:"des,omitempty"`
}

func (x *FileTrans) Reset() {
	*x = FileTrans{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scene_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileTrans) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileTrans) ProtoMessage() {}

func (x *FileTrans) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scene_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileTrans.ProtoReflect.Descriptor instead.
func (*FileTrans) Descriptor() ([]byte, []int) {
	return file_proto_scene_proto_rawDescGZIP(), []int{5}
}

func (x *FileTrans) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

func (x *FileTrans) GetDes() string {
	if x != nil {
		return x.Des
	}
	return ""
}

var File_proto_scene_proto protoreflect.FileDescriptor

var file_proto_scene_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x62, 0x67, 0x70, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x62, 0x67, 0x70, 0x2f, 0x67, 0x6f, 0x62, 0x67, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x05, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x70, 0x6f, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x70, 0x6f, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x2d, 0x0a, 0x09, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x2e, 0x42, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x52, 0x09, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x73,
	0x22, 0x7c, 0x0a, 0x08, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x12, 0x21, 0x0a, 0x05, 0x73,
	0x74, 0x65, 0x70, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x73, 0x63, 0x65,
	0x6e, 0x65, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x22, 0x9b,
	0x01, 0x0a, 0x04, 0x53, 0x74, 0x65, 0x70, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x2e, 0x53, 0x74,
	0x65, 0x70, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x28, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x2f, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x43, 0x4d, 0x44, 0x53, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x41,
	0x49, 0x54, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x46, 0x54, 0x10, 0x03, 0x22, 0x1e, 0x0a, 0x08,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6d, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6d, 0x64, 0x73, 0x22, 0x38, 0x0a, 0x04,
	0x57, 0x61, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x2f, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x73, 0x72, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x64, 0x65, 0x73, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_scene_proto_rawDescOnce sync.Once
	file_proto_scene_proto_rawDescData = file_proto_scene_proto_rawDesc
)

func file_proto_scene_proto_rawDescGZIP() []byte {
	file_proto_scene_proto_rawDescOnce.Do(func() {
		file_proto_scene_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_scene_proto_rawDescData)
	})
	return file_proto_scene_proto_rawDescData
}

var file_proto_scene_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_scene_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_scene_proto_goTypes = []interface{}{
	(Step_Type)(0),    // 0: scene.Step.Type
	(*Scene)(nil),     // 1: scene.Scene
	(*Behavior)(nil),  // 2: scene.Behavior
	(*Step)(nil),      // 3: scene.Step
	(*Commands)(nil),  // 4: scene.Commands
	(*Wait)(nil),      // 5: scene.Wait
	(*FileTrans)(nil), // 6: scene.FileTrans
	(*anypb.Any)(nil), // 7: google.protobuf.Any
}
var file_proto_scene_proto_depIdxs = []int32{
	2, // 0: scene.Scene.behaviors:type_name -> scene.Behavior
	3, // 1: scene.Behavior.steps:type_name -> scene.Step
	0, // 2: scene.Step.type:type_name -> scene.Step.Type
	7, // 3: scene.Step.body:type_name -> google.protobuf.Any
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_scene_proto_init() }
func file_proto_scene_proto_init() {
	if File_proto_scene_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_scene_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scene); i {
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
		file_proto_scene_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Behavior); i {
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
		file_proto_scene_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Step); i {
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
		file_proto_scene_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commands); i {
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
		file_proto_scene_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wait); i {
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
		file_proto_scene_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileTrans); i {
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
			RawDescriptor: file_proto_scene_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_scene_proto_goTypes,
		DependencyIndexes: file_proto_scene_proto_depIdxs,
		EnumInfos:         file_proto_scene_proto_enumTypes,
		MessageInfos:      file_proto_scene_proto_msgTypes,
	}.Build()
	File_proto_scene_proto = out.File
	file_proto_scene_proto_rawDesc = nil
	file_proto_scene_proto_goTypes = nil
	file_proto_scene_proto_depIdxs = nil
}
