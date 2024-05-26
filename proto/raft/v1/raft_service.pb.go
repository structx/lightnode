// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: proto/raft/v1/raft_service.proto

package v1

import (
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

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term  int64  `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Index int64  `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Cmd   []byte `protobuf:"bytes,3,opt,name=cmd,proto3" json:"cmd,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raft_v1_raft_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_v1_raft_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_proto_raft_v1_raft_service_proto_rawDescGZIP(), []int{0}
}

func (x *Log) GetTerm() int64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *Log) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Log) GetCmd() []byte {
	if x != nil {
		return x.Cmd
	}
	return nil
}

type AppendEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term         int64  `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	LeaderId     string `protobuf:"bytes,2,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
	PrevLogIndex int64  `protobuf:"varint,3,opt,name=prev_log_index,json=prevLogIndex,proto3" json:"prev_log_index,omitempty"`
	PrevLogTerm  int64  `protobuf:"varint,4,opt,name=prev_log_term,json=prevLogTerm,proto3" json:"prev_log_term,omitempty"`
	Entries      []*Log `protobuf:"bytes,5,rep,name=entries,proto3" json:"entries,omitempty"`
	LeaderCommit int64  `protobuf:"varint,6,opt,name=leader_commit,json=leaderCommit,proto3" json:"leader_commit,omitempty"`
}

func (x *AppendEntryRequest) Reset() {
	*x = AppendEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raft_v1_raft_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntryRequest) ProtoMessage() {}

func (x *AppendEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_v1_raft_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntryRequest.ProtoReflect.Descriptor instead.
func (*AppendEntryRequest) Descriptor() ([]byte, []int) {
	return file_proto_raft_v1_raft_service_proto_rawDescGZIP(), []int{1}
}

func (x *AppendEntryRequest) GetTerm() int64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntryRequest) GetLeaderId() string {
	if x != nil {
		return x.LeaderId
	}
	return ""
}

func (x *AppendEntryRequest) GetPrevLogIndex() int64 {
	if x != nil {
		return x.PrevLogIndex
	}
	return 0
}

func (x *AppendEntryRequest) GetPrevLogTerm() int64 {
	if x != nil {
		return x.PrevLogTerm
	}
	return 0
}

func (x *AppendEntryRequest) GetEntries() []*Log {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *AppendEntryRequest) GetLeaderCommit() int64 {
	if x != nil {
		return x.LeaderCommit
	}
	return 0
}

type AppendEntryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term    int64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Success bool  `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AppendEntryResponse) Reset() {
	*x = AppendEntryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raft_v1_raft_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntryResponse) ProtoMessage() {}

func (x *AppendEntryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_v1_raft_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntryResponse.ProtoReflect.Descriptor instead.
func (*AppendEntryResponse) Descriptor() ([]byte, []int) {
	return file_proto_raft_v1_raft_service_proto_rawDescGZIP(), []int{2}
}

func (x *AppendEntryResponse) GetTerm() int64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntryResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type RequestVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term         int64  `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	CandidateId  string `protobuf:"bytes,2,opt,name=candidate_id,json=candidateId,proto3" json:"candidate_id,omitempty"`
	LastLogIndex int64  `protobuf:"varint,3,opt,name=last_log_index,json=lastLogIndex,proto3" json:"last_log_index,omitempty"`
	LastLogTerm  int64  `protobuf:"varint,4,opt,name=last_log_term,json=lastLogTerm,proto3" json:"last_log_term,omitempty"`
}

func (x *RequestVoteRequest) Reset() {
	*x = RequestVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raft_v1_raft_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteRequest) ProtoMessage() {}

func (x *RequestVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_v1_raft_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteRequest.ProtoReflect.Descriptor instead.
func (*RequestVoteRequest) Descriptor() ([]byte, []int) {
	return file_proto_raft_v1_raft_service_proto_rawDescGZIP(), []int{3}
}

func (x *RequestVoteRequest) GetTerm() int64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteRequest) GetCandidateId() string {
	if x != nil {
		return x.CandidateId
	}
	return ""
}

func (x *RequestVoteRequest) GetLastLogIndex() int64 {
	if x != nil {
		return x.LastLogIndex
	}
	return 0
}

func (x *RequestVoteRequest) GetLastLogTerm() int64 {
	if x != nil {
		return x.LastLogTerm
	}
	return 0
}

type RequestVoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term        int64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	VoteGranted bool  `protobuf:"varint,2,opt,name=vote_granted,json=voteGranted,proto3" json:"vote_granted,omitempty"`
}

func (x *RequestVoteResponse) Reset() {
	*x = RequestVoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raft_v1_raft_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteResponse) ProtoMessage() {}

func (x *RequestVoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_v1_raft_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteResponse.ProtoReflect.Descriptor instead.
func (*RequestVoteResponse) Descriptor() ([]byte, []int) {
	return file_proto_raft_v1_raft_service_proto_rawDescGZIP(), []int{4}
}

func (x *RequestVoteResponse) GetTerm() int64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteResponse) GetVoteGranted() bool {
	if x != nil {
		return x.VoteGranted
	}
	return false
}

// Invoke by leader to send chunks of a snapshot to a follower
// Leaders always send chunks in order
type InstallSnapshotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// leaders term
	Term int64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	// follower can redirect clients
	LeaderId string `protobuf:"bytes,2,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
	// snapshot replaces all entries up through and including this index
	LastIncludedIndex int64 `protobuf:"varint,3,opt,name=last_included_index,json=lastIncludedIndex,proto3" json:"last_included_index,omitempty"`
	// term of lastIncludedIndex
	LastIncludedTerm int64 `protobuf:"varint,4,opt,name=last_included_term,json=lastIncludedTerm,proto3" json:"last_included_term,omitempty"`
	// bytes offset where chunk is positioned in the snapshot file
	Offset []byte `protobuf:"bytes,5,opt,name=offset,proto3" json:"offset,omitempty"`
	// raw bytes of the snapshot chunk, starting at offset
	Data []byte `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *InstallSnapshotRequest) Reset() {
	*x = InstallSnapshotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raft_v1_raft_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallSnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallSnapshotRequest) ProtoMessage() {}

func (x *InstallSnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_v1_raft_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallSnapshotRequest.ProtoReflect.Descriptor instead.
func (*InstallSnapshotRequest) Descriptor() ([]byte, []int) {
	return file_proto_raft_v1_raft_service_proto_rawDescGZIP(), []int{5}
}

func (x *InstallSnapshotRequest) GetTerm() int64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *InstallSnapshotRequest) GetLeaderId() string {
	if x != nil {
		return x.LeaderId
	}
	return ""
}

func (x *InstallSnapshotRequest) GetLastIncludedIndex() int64 {
	if x != nil {
		return x.LastIncludedIndex
	}
	return 0
}

func (x *InstallSnapshotRequest) GetLastIncludedTerm() int64 {
	if x != nil {
		return x.LastIncludedTerm
	}
	return 0
}

func (x *InstallSnapshotRequest) GetOffset() []byte {
	if x != nil {
		return x.Offset
	}
	return nil
}

func (x *InstallSnapshotRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type InstallSnapshotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// current term for leader to update itself
	Term int64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
}

func (x *InstallSnapshotResponse) Reset() {
	*x = InstallSnapshotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raft_v1_raft_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallSnapshotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallSnapshotResponse) ProtoMessage() {}

func (x *InstallSnapshotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_v1_raft_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallSnapshotResponse.ProtoReflect.Descriptor instead.
func (*InstallSnapshotResponse) Descriptor() ([]byte, []int) {
	return file_proto_raft_v1_raft_service_proto_rawDescGZIP(), []int{6}
}

func (x *InstallSnapshotResponse) GetTerm() int64 {
	if x != nil {
		return x.Term
	}
	return 0
}

var File_proto_raft_v1_raft_service_proto protoreflect.FileDescriptor

var file_proto_raft_v1_raft_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x61, 0x66, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x41, 0x0a, 0x03, 0x4c,
	0x6f, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x03,
	0x63, 0x6d, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x22, 0xdc,
	0x01, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x6c,
	0x6f, 0x67, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x22, 0x0a, 0x0d,
	0x70, 0x72, 0x65, 0x76, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d,
	0x12, 0x26, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x52,
	0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x22, 0x43, 0x0a,
	0x13, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x95, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64,
	0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f,
	0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x22, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6c,
	0x6f, 0x67, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c,
	0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x22, 0x4c, 0x0a, 0x13, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x67, 0x72,
	0x61, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x76, 0x6f, 0x74,
	0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x22, 0xd3, 0x01, 0x0a, 0x16, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x54, 0x65,
	0x72, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2d,
	0x0a, 0x17, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x32, 0x81, 0x02,
	0x0a, 0x0d, 0x52, 0x61, 0x66, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x31, 0x12,
	0x4c, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x12, 0x1b, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e,
	0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a,
	0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x72,
	0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x61, 0x66, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0f, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x1f, 0x2e, 0x72,
	0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x78, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x6f, 0x64,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_raft_v1_raft_service_proto_rawDescOnce sync.Once
	file_proto_raft_v1_raft_service_proto_rawDescData = file_proto_raft_v1_raft_service_proto_rawDesc
)

func file_proto_raft_v1_raft_service_proto_rawDescGZIP() []byte {
	file_proto_raft_v1_raft_service_proto_rawDescOnce.Do(func() {
		file_proto_raft_v1_raft_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_raft_v1_raft_service_proto_rawDescData)
	})
	return file_proto_raft_v1_raft_service_proto_rawDescData
}

var file_proto_raft_v1_raft_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_raft_v1_raft_service_proto_goTypes = []interface{}{
	(*Log)(nil),                     // 0: raft.v1.Log
	(*AppendEntryRequest)(nil),      // 1: raft.v1.AppendEntryRequest
	(*AppendEntryResponse)(nil),     // 2: raft.v1.AppendEntryResponse
	(*RequestVoteRequest)(nil),      // 3: raft.v1.RequestVoteRequest
	(*RequestVoteResponse)(nil),     // 4: raft.v1.RequestVoteResponse
	(*InstallSnapshotRequest)(nil),  // 5: raft.v1.InstallSnapshotRequest
	(*InstallSnapshotResponse)(nil), // 6: raft.v1.InstallSnapshotResponse
}
var file_proto_raft_v1_raft_service_proto_depIdxs = []int32{
	0, // 0: raft.v1.AppendEntryRequest.entries:type_name -> raft.v1.Log
	1, // 1: raft.v1.RaftServiceV1.AppendEntries:input_type -> raft.v1.AppendEntryRequest
	3, // 2: raft.v1.RaftServiceV1.RequestVote:input_type -> raft.v1.RequestVoteRequest
	5, // 3: raft.v1.RaftServiceV1.InstallSnapshot:input_type -> raft.v1.InstallSnapshotRequest
	2, // 4: raft.v1.RaftServiceV1.AppendEntries:output_type -> raft.v1.AppendEntryResponse
	4, // 5: raft.v1.RaftServiceV1.RequestVote:output_type -> raft.v1.RequestVoteResponse
	6, // 6: raft.v1.RaftServiceV1.InstallSnapshot:output_type -> raft.v1.InstallSnapshotResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_raft_v1_raft_service_proto_init() }
func file_proto_raft_v1_raft_service_proto_init() {
	if File_proto_raft_v1_raft_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_raft_v1_raft_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_proto_raft_v1_raft_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntryRequest); i {
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
		file_proto_raft_v1_raft_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntryResponse); i {
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
		file_proto_raft_v1_raft_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestVoteRequest); i {
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
		file_proto_raft_v1_raft_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestVoteResponse); i {
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
		file_proto_raft_v1_raft_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallSnapshotRequest); i {
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
		file_proto_raft_v1_raft_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallSnapshotResponse); i {
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
			RawDescriptor: file_proto_raft_v1_raft_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_raft_v1_raft_service_proto_goTypes,
		DependencyIndexes: file_proto_raft_v1_raft_service_proto_depIdxs,
		MessageInfos:      file_proto_raft_v1_raft_service_proto_msgTypes,
	}.Build()
	File_proto_raft_v1_raft_service_proto = out.File
	file_proto_raft_v1_raft_service_proto_rawDesc = nil
	file_proto_raft_v1_raft_service_proto_goTypes = nil
	file_proto_raft_v1_raft_service_proto_depIdxs = nil
}
