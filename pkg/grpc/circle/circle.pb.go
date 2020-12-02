// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: internal/manager/circle/circle.proto

package circle

import (
	proto "github.com/golang/protobuf/proto"
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

type ResourceHealth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  *string `protobuf:"bytes,1,opt,name=status,proto3,oneof" json:"status,omitempty"`
	Message *string `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
}

func (x *ResourceHealth) Reset() {
	*x = ResourceHealth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_manager_circle_circle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceHealth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceHealth) ProtoMessage() {}

func (x *ResourceHealth) ProtoReflect() protoreflect.Message {
	mi := &file_internal_manager_circle_circle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceHealth.ProtoReflect.Descriptor instead.
func (*ResourceHealth) Descriptor() ([]byte, []int) {
	return file_internal_manager_circle_circle_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceHealth) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *ResourceHealth) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type ResourceStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind              string          `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	Name              string          `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	CreationTimestamp string          `protobuf:"bytes,7,opt,name=creationTimestamp,proto3" json:"creationTimestamp,omitempty"`
	Health            *ResourceHealth `protobuf:"bytes,6,opt,name=health,proto3,oneof" json:"health,omitempty"`
}

func (x *ResourceStatus) Reset() {
	*x = ResourceStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_manager_circle_circle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceStatus) ProtoMessage() {}

func (x *ResourceStatus) ProtoReflect() protoreflect.Message {
	mi := &file_internal_manager_circle_circle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceStatus.ProtoReflect.Descriptor instead.
func (*ResourceStatus) Descriptor() ([]byte, []int) {
	return file_internal_manager_circle_circle_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceStatus) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ResourceStatus) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceStatus) GetCreationTimestamp() string {
	if x != nil {
		return x.CreationTimestamp
	}
	return ""
}

func (x *ResourceStatus) GetHealth() *ResourceHealth {
	if x != nil {
		return x.Health
	}
	return nil
}

type ResourceParent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Kind       string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Controller bool   `protobuf:"varint,3,opt,name=controller,proto3" json:"controller,omitempty"`
}

func (x *ResourceParent) Reset() {
	*x = ResourceParent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_manager_circle_circle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceParent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceParent) ProtoMessage() {}

func (x *ResourceParent) ProtoReflect() protoreflect.Message {
	mi := &file_internal_manager_circle_circle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceParent.ProtoReflect.Descriptor instead.
func (*ResourceParent) Descriptor() ([]byte, []int) {
	return file_internal_manager_circle_circle_proto_rawDescGZIP(), []int{2}
}

func (x *ResourceParent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceParent) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ResourceParent) GetController() bool {
	if x != nil {
		return x.Controller
	}
	return false
}

type ResourceNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref     *ResourceStatus   `protobuf:"bytes,1,opt,name=ref,proto3,oneof" json:"ref,omitempty"`
	Parents []*ResourceParent `protobuf:"bytes,2,rep,name=parents,proto3" json:"parents,omitempty"`
}

func (x *ResourceNode) Reset() {
	*x = ResourceNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_manager_circle_circle_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceNode) ProtoMessage() {}

func (x *ResourceNode) ProtoReflect() protoreflect.Message {
	mi := &file_internal_manager_circle_circle_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceNode.ProtoReflect.Descriptor instead.
func (*ResourceNode) Descriptor() ([]byte, []int) {
	return file_internal_manager_circle_circle_proto_rawDescGZIP(), []int{3}
}

func (x *ResourceNode) GetRef() *ResourceStatus {
	if x != nil {
		return x.Ref
	}
	return nil
}

func (x *ResourceNode) GetParents() []*ResourceParent {
	if x != nil {
		return x.Parents
	}
	return nil
}

type ProjectNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Resources []*ResourceNode `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *ProjectNode) Reset() {
	*x = ProjectNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_manager_circle_circle_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectNode) ProtoMessage() {}

func (x *ProjectNode) ProtoReflect() protoreflect.Message {
	mi := &file_internal_manager_circle_circle_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectNode.ProtoReflect.Descriptor instead.
func (*ProjectNode) Descriptor() ([]byte, []int) {
	return file_internal_manager_circle_circle_proto_rawDescGZIP(), []int{4}
}

func (x *ProjectNode) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProjectNode) GetResources() []*ResourceNode {
	if x != nil {
		return x.Resources
	}
	return nil
}

type CircleTreeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*ProjectNode `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *CircleTreeResponse) Reset() {
	*x = CircleTreeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_manager_circle_circle_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CircleTreeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CircleTreeResponse) ProtoMessage() {}

func (x *CircleTreeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_manager_circle_circle_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CircleTreeResponse.ProtoReflect.Descriptor instead.
func (*CircleTreeResponse) Descriptor() ([]byte, []int) {
	return file_internal_manager_circle_circle_proto_rawDescGZIP(), []int{5}
}

func (x *CircleTreeResponse) GetNodes() []*ProjectNode {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type Circle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *Circle) Reset() {
	*x = Circle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_manager_circle_circle_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Circle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Circle) ProtoMessage() {}

func (x *Circle) ProtoReflect() protoreflect.Message {
	mi := &file_internal_manager_circle_circle_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Circle.ProtoReflect.Descriptor instead.
func (*Circle) Descriptor() ([]byte, []int) {
	return file_internal_manager_circle_circle_proto_rawDescGZIP(), []int{6}
}

func (x *Circle) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Circle) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

var File_internal_manager_circle_circle_proto protoreflect.FileDescriptor

var file_internal_manager_circle_circle_proto_rawDesc = []byte{
	0x0a, 0x24, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2f, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x2f, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x22, 0x63,
	0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0xa6, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c,
	0x0a, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x33, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63,
	0x69, 0x72, 0x63, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x48, 0x00, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x88, 0x01,
	0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x22, 0x58, 0x0a, 0x0e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x22, 0x77, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x2d, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x03, 0x72,
	0x65, 0x66, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x07,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x72, 0x65, 0x66, 0x22,
	0x55, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x3f, 0x0a, 0x12, 0x43, 0x69, 0x72, 0x63, 0x6c, 0x65,
	0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x69,
	0x72, 0x63, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x3a, 0x0a, 0x06, 0x43, 0x69, 0x72, 0x63, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x32, 0x4b, 0x0a, 0x0d, 0x43, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x43, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x54, 0x72,
	0x65, 0x65, 0x12, 0x0e, 0x2e, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x2e, 0x43, 0x69, 0x72, 0x63,
	0x6c, 0x65, 0x1a, 0x1a, 0x2e, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x2e, 0x43, 0x69, 0x72, 0x63,
	0x6c, 0x65, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x09, 0x5a, 0x07, 0x2f, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_internal_manager_circle_circle_proto_rawDescOnce sync.Once
	file_internal_manager_circle_circle_proto_rawDescData = file_internal_manager_circle_circle_proto_rawDesc
)

func file_internal_manager_circle_circle_proto_rawDescGZIP() []byte {
	file_internal_manager_circle_circle_proto_rawDescOnce.Do(func() {
		file_internal_manager_circle_circle_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_manager_circle_circle_proto_rawDescData)
	})
	return file_internal_manager_circle_circle_proto_rawDescData
}

var file_internal_manager_circle_circle_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_internal_manager_circle_circle_proto_goTypes = []interface{}{
	(*ResourceHealth)(nil),     // 0: circle.ResourceHealth
	(*ResourceStatus)(nil),     // 1: circle.ResourceStatus
	(*ResourceParent)(nil),     // 2: circle.ResourceParent
	(*ResourceNode)(nil),       // 3: circle.ResourceNode
	(*ProjectNode)(nil),        // 4: circle.ProjectNode
	(*CircleTreeResponse)(nil), // 5: circle.CircleTreeResponse
	(*Circle)(nil),             // 6: circle.Circle
}
var file_internal_manager_circle_circle_proto_depIdxs = []int32{
	0, // 0: circle.ResourceStatus.health:type_name -> circle.ResourceHealth
	1, // 1: circle.ResourceNode.ref:type_name -> circle.ResourceStatus
	2, // 2: circle.ResourceNode.parents:type_name -> circle.ResourceParent
	3, // 3: circle.ProjectNode.resources:type_name -> circle.ResourceNode
	4, // 4: circle.CircleTreeResponse.nodes:type_name -> circle.ProjectNode
	6, // 5: circle.CircleService.CircleTree:input_type -> circle.Circle
	5, // 6: circle.CircleService.CircleTree:output_type -> circle.CircleTreeResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_internal_manager_circle_circle_proto_init() }
func file_internal_manager_circle_circle_proto_init() {
	if File_internal_manager_circle_circle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_manager_circle_circle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceHealth); i {
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
		file_internal_manager_circle_circle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceStatus); i {
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
		file_internal_manager_circle_circle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceParent); i {
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
		file_internal_manager_circle_circle_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceNode); i {
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
		file_internal_manager_circle_circle_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectNode); i {
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
		file_internal_manager_circle_circle_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CircleTreeResponse); i {
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
		file_internal_manager_circle_circle_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Circle); i {
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
	file_internal_manager_circle_circle_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_internal_manager_circle_circle_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_internal_manager_circle_circle_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_manager_circle_circle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_manager_circle_circle_proto_goTypes,
		DependencyIndexes: file_internal_manager_circle_circle_proto_depIdxs,
		MessageInfos:      file_internal_manager_circle_circle_proto_msgTypes,
	}.Build()
	File_internal_manager_circle_circle_proto = out.File
	file_internal_manager_circle_circle_proto_rawDesc = nil
	file_internal_manager_circle_circle_proto_goTypes = nil
	file_internal_manager_circle_circle_proto_depIdxs = nil
}
