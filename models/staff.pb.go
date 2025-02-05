// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.7
// source: staff.proto

package models

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

type PollStatus int32

const (
	PollStatus_InvalidPollStatus PollStatus = 0
	PollStatus_Valid             PollStatus = 1
	PollStatus_Empty             PollStatus = 2
	PollStatus_Terminated        PollStatus = 100
)

// Enum value maps for PollStatus.
var (
	PollStatus_name = map[int32]string{
		0:   "InvalidPollStatus",
		1:   "Valid",
		2:   "Empty",
		100: "Terminated",
	}
	PollStatus_value = map[string]int32{
		"InvalidPollStatus": 0,
		"Valid":             1,
		"Empty":             2,
		"Terminated":        100,
	}
)

func (x PollStatus) Enum() *PollStatus {
	p := new(PollStatus)
	*p = x
	return p
}

func (x PollStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PollStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_staff_proto_enumTypes[0].Descriptor()
}

func (PollStatus) Type() protoreflect.EnumType {
	return &file_staff_proto_enumTypes[0]
}

func (x PollStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PollStatus.Descriptor instead.
func (PollStatus) EnumDescriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{0}
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaffId string `protobuf:"bytes,1,opt,name=staff_id,json=staffId,proto3" json:"staff_id,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetStaffId() string {
	if x != nil {
		return x.StaffId
	}
	return ""
}

type RegisterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterReply) Reset() {
	*x = RegisterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReply) ProtoMessage() {}

func (x *RegisterReply) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReply.ProtoReflect.Descriptor instead.
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{1}
}

type ElectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId    string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	StaffId   string `protobuf:"bytes,2,opt,name=staff_id,json=staffId,proto3" json:"staff_id,omitempty"`
	StaffAddr string `protobuf:"bytes,3,opt,name=staff_addr,json=staffAddr,proto3" json:"staff_addr,omitempty"`
}

func (x *ElectRequest) Reset() {
	*x = ElectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElectRequest) ProtoMessage() {}

func (x *ElectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElectRequest.ProtoReflect.Descriptor instead.
func (*ElectRequest) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{2}
}

func (x *ElectRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *ElectRequest) GetStaffId() string {
	if x != nil {
		return x.StaffId
	}
	return ""
}

func (x *ElectRequest) GetStaffAddr() string {
	if x != nil {
		return x.StaffAddr
	}
	return ""
}

type ElectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeaderId   string `protobuf:"bytes,1,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
	LeaderAddr string `protobuf:"bytes,2,opt,name=leader_addr,json=leaderAddr,proto3" json:"leader_addr,omitempty"`
}

func (x *ElectReply) Reset() {
	*x = ElectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElectReply) ProtoMessage() {}

func (x *ElectReply) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElectReply.ProtoReflect.Descriptor instead.
func (*ElectReply) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{3}
}

func (x *ElectReply) GetLeaderId() string {
	if x != nil {
		return x.LeaderId
	}
	return ""
}

func (x *ElectReply) GetLeaderAddr() string {
	if x != nil {
		return x.LeaderAddr
	}
	return ""
}

type PollRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaffId string `protobuf:"bytes,1,opt,name=staff_id,json=staffId,proto3" json:"staff_id,omitempty"`
}

func (x *PollRequest) Reset() {
	*x = PollRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PollRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PollRequest) ProtoMessage() {}

func (x *PollRequest) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PollRequest.ProtoReflect.Descriptor instead.
func (*PollRequest) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{4}
}

func (x *PollRequest) GetStaffId() string {
	if x != nil {
		return x.StaffId
	}
	return ""
}

type PollReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status PollStatus `protobuf:"varint,1,opt,name=status,proto3,enum=staff.PollStatus" json:"status,omitempty"`
	Task   *Task      `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *PollReply) Reset() {
	*x = PollReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PollReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PollReply) ProtoMessage() {}

func (x *PollReply) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PollReply.ProtoReflect.Descriptor instead.
func (*PollReply) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{5}
}

func (x *PollReply) GetStatus() PollStatus {
	if x != nil {
		return x.Status
	}
	return PollStatus_InvalidPollStatus
}

func (x *PollReply) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type FinishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *FinishRequest) Reset() {
	*x = FinishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishRequest) ProtoMessage() {}

func (x *FinishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishRequest.ProtoReflect.Descriptor instead.
func (*FinishRequest) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{6}
}

func (x *FinishRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

type FinishReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FinishReply) Reset() {
	*x = FinishReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishReply) ProtoMessage() {}

func (x *FinishReply) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishReply.ProtoReflect.Descriptor instead.
func (*FinishReply) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{7}
}

var File_staff_proto protoreflect.FileDescriptor

var file_staff_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x1a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x2c, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x22, 0x0f,
	0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x61, 0x0a, 0x0c, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x66, 0x66, 0x41, 0x64,
	0x64, 0x72, 0x22, 0x4a, 0x0a, 0x0a, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x22, 0x28,
	0x0a, 0x0b, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x09, 0x50, 0x6f, 0x6c, 0x6c,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x50, 0x6f,
	0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1e, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b,
	0x22, 0x28, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2a, 0x49, 0x0a, 0x0a, 0x50, 0x6f, 0x6c,
	0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x6e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x50, 0x6f, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x64, 0x10, 0x64, 0x32, 0xde, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x66, 0x66, 0x12, 0x3a,
	0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x73, 0x74, 0x61,
	0x66, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x05, 0x45, 0x6c,
	0x65, 0x63, 0x74, 0x12, 0x13, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x45, 0x6c, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66,
	0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x30, 0x0a,
	0x04, 0x50, 0x6f, 0x6c, 0x6c, 0x12, 0x12, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x50, 0x6f,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x2e, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x34, 0x0a, 0x06, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x14, 0x2e, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2f, 0x64, 0x6d, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_staff_proto_rawDescOnce sync.Once
	file_staff_proto_rawDescData = file_staff_proto_rawDesc
)

func file_staff_proto_rawDescGZIP() []byte {
	file_staff_proto_rawDescOnce.Do(func() {
		file_staff_proto_rawDescData = protoimpl.X.CompressGZIP(file_staff_proto_rawDescData)
	})
	return file_staff_proto_rawDescData
}

var file_staff_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_staff_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_staff_proto_goTypes = []interface{}{
	(PollStatus)(0),         // 0: staff.PollStatus
	(*RegisterRequest)(nil), // 1: staff.RegisterRequest
	(*RegisterReply)(nil),   // 2: staff.RegisterReply
	(*ElectRequest)(nil),    // 3: staff.ElectRequest
	(*ElectReply)(nil),      // 4: staff.ElectReply
	(*PollRequest)(nil),     // 5: staff.PollRequest
	(*PollReply)(nil),       // 6: staff.PollReply
	(*FinishRequest)(nil),   // 7: staff.FinishRequest
	(*FinishReply)(nil),     // 8: staff.FinishReply
	(*Task)(nil),            // 9: task.Task
}
var file_staff_proto_depIdxs = []int32{
	0, // 0: staff.PollReply.status:type_name -> staff.PollStatus
	9, // 1: staff.PollReply.task:type_name -> task.Task
	1, // 2: staff.Staff.Register:input_type -> staff.RegisterRequest
	3, // 3: staff.Staff.Elect:input_type -> staff.ElectRequest
	5, // 4: staff.Staff.Poll:input_type -> staff.PollRequest
	7, // 5: staff.Staff.Finish:input_type -> staff.FinishRequest
	2, // 6: staff.Staff.Register:output_type -> staff.RegisterReply
	4, // 7: staff.Staff.Elect:output_type -> staff.ElectReply
	6, // 8: staff.Staff.Poll:output_type -> staff.PollReply
	8, // 9: staff.Staff.Finish:output_type -> staff.FinishReply
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_staff_proto_init() }
func file_staff_proto_init() {
	if File_staff_proto != nil {
		return
	}
	file_task_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_staff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_staff_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReply); i {
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
		file_staff_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElectRequest); i {
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
		file_staff_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElectReply); i {
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
		file_staff_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PollRequest); i {
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
		file_staff_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PollReply); i {
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
		file_staff_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishRequest); i {
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
		file_staff_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishReply); i {
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
			RawDescriptor: file_staff_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_staff_proto_goTypes,
		DependencyIndexes: file_staff_proto_depIdxs,
		EnumInfos:         file_staff_proto_enumTypes,
		MessageInfos:      file_staff_proto_msgTypes,
	}.Build()
	File_staff_proto = out.File
	file_staff_proto_rawDesc = nil
	file_staff_proto_goTypes = nil
	file_staff_proto_depIdxs = nil
}
