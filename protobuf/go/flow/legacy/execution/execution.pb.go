// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: flow/legacy/execution/execution.proto

package execution

import (
	entities "github.com/onflow/flow/protobuf/go/flow/legacy/entities"
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

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_legacy_execution_execution_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flow_legacy_execution_execution_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_flow_legacy_execution_execution_proto_rawDescGZIP(), []int{0}
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_legacy_execution_execution_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flow_legacy_execution_execution_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_flow_legacy_execution_execution_proto_rawDescGZIP(), []int{1}
}

type GetAccountAtBlockIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId []byte `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Address []byte `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetAccountAtBlockIDRequest) Reset() {
	*x = GetAccountAtBlockIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_legacy_execution_execution_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountAtBlockIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountAtBlockIDRequest) ProtoMessage() {}

func (x *GetAccountAtBlockIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flow_legacy_execution_execution_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountAtBlockIDRequest.ProtoReflect.Descriptor instead.
func (*GetAccountAtBlockIDRequest) Descriptor() ([]byte, []int) {
	return file_flow_legacy_execution_execution_proto_rawDescGZIP(), []int{2}
}

func (x *GetAccountAtBlockIDRequest) GetBlockId() []byte {
	if x != nil {
		return x.BlockId
	}
	return nil
}

func (x *GetAccountAtBlockIDRequest) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

type GetAccountAtBlockIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *entities.Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *GetAccountAtBlockIDResponse) Reset() {
	*x = GetAccountAtBlockIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_legacy_execution_execution_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountAtBlockIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountAtBlockIDResponse) ProtoMessage() {}

func (x *GetAccountAtBlockIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flow_legacy_execution_execution_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountAtBlockIDResponse.ProtoReflect.Descriptor instead.
func (*GetAccountAtBlockIDResponse) Descriptor() ([]byte, []int) {
	return file_flow_legacy_execution_execution_proto_rawDescGZIP(), []int{3}
}

func (x *GetAccountAtBlockIDResponse) GetAccount() *entities.Account {
	if x != nil {
		return x.Account
	}
	return nil
}

type ExecuteScriptAtBlockIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId   []byte   `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Script    []byte   `protobuf:"bytes,2,opt,name=script,proto3" json:"script,omitempty"`
	Arguments [][]byte `protobuf:"bytes,3,rep,name=arguments,proto3" json:"arguments,omitempty"`
}

func (x *ExecuteScriptAtBlockIDRequest) Reset() {
	*x = ExecuteScriptAtBlockIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_legacy_execution_execution_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteScriptAtBlockIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteScriptAtBlockIDRequest) ProtoMessage() {}

func (x *ExecuteScriptAtBlockIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flow_legacy_execution_execution_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteScriptAtBlockIDRequest.ProtoReflect.Descriptor instead.
func (*ExecuteScriptAtBlockIDRequest) Descriptor() ([]byte, []int) {
	return file_flow_legacy_execution_execution_proto_rawDescGZIP(), []int{4}
}

func (x *ExecuteScriptAtBlockIDRequest) GetBlockId() []byte {
	if x != nil {
		return x.BlockId
	}
	return nil
}

func (x *ExecuteScriptAtBlockIDRequest) GetScript() []byte {
	if x != nil {
		return x.Script
	}
	return nil
}

func (x *ExecuteScriptAtBlockIDRequest) GetArguments() [][]byte {
	if x != nil {
		return x.Arguments
	}
	return nil
}

type ExecuteScriptAtBlockIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ExecuteScriptAtBlockIDResponse) Reset() {
	*x = ExecuteScriptAtBlockIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_legacy_execution_execution_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteScriptAtBlockIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteScriptAtBlockIDResponse) ProtoMessage() {}

func (x *ExecuteScriptAtBlockIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flow_legacy_execution_execution_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteScriptAtBlockIDResponse.ProtoReflect.Descriptor instead.
func (*ExecuteScriptAtBlockIDResponse) Descriptor() ([]byte, []int) {
	return file_flow_legacy_execution_execution_proto_rawDescGZIP(), []int{5}
}

func (x *ExecuteScriptAtBlockIDResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type GetEventsForBlockIDsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*GetEventsForBlockIDsResponse_Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *GetEventsForBlockIDsResponse) Reset() {
	*x = GetEventsForBlockIDsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_legacy_execution_execution_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsForBlockIDsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsForBlockIDsResponse) ProtoMessage() {}

func (x *GetEventsForBlockIDsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flow_legacy_execution_execution_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsForBlockIDsResponse.ProtoReflect.Descriptor instead.
func (*GetEventsForBlockIDsResponse) Descriptor() ([]byte, []int) {
	return file_flow_legacy_execution_execution_proto_rawDescGZIP(), []int{6}
}

func (x *GetEventsForBlockIDsResponse) GetResults() []*GetEventsForBlockIDsResponse_Result {
	if x != nil {
		return x.Results
	}
	return nil
}

type GetEventsForBlockIDsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	BlockIds [][]byte `protobuf:"bytes,2,rep,name=block_ids,json=blockIds,proto3" json:"block_ids,omitempty"`
}

func (x *GetEventsForBlockIDsRequest) Reset() {
	*x = GetEventsForBlockIDsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_legacy_execution_execution_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsForBlockIDsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsForBlockIDsRequest) ProtoMessage() {}

func (x *GetEventsForBlockIDsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flow_legacy_execution_execution_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsForBlockIDsRequest.ProtoReflect.Descriptor instead.
func (*GetEventsForBlockIDsRequest) Descriptor() ([]byte, []int) {
	return file_flow_legacy_execution_execution_proto_rawDescGZIP(), []int{7}
}

func (x *GetEventsForBlockIDsRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetEventsForBlockIDsRequest) GetBlockIds() [][]byte {
	if x != nil {
		return x.BlockIds
	}
	return nil
}

type GetTransactionResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId       []byte `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	TransactionId []byte `protobuf:"bytes,2,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
}

func (x *GetTransactionResultRequest) Reset() {
	*x = GetTransactionResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_legacy_execution_execution_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionResultRequest) ProtoMessage() {}

func (x *GetTransactionResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flow_legacy_execution_execution_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionResultRequest.ProtoReflect.Descriptor instead.
func (*GetTransactionResultRequest) Descriptor() ([]byte, []int) {
	return file_flow_legacy_execution_execution_proto_rawDescGZIP(), []int{8}
}

func (x *GetTransactionResultRequest) GetBlockId() []byte {
	if x != nil {
		return x.BlockId
	}
	return nil
}

func (x *GetTransactionResultRequest) GetTransactionId() []byte {
	if x != nil {
		return x.TransactionId
	}
	return nil
}

type GetTransactionResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode   uint32            `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	ErrorMessage string            `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	Events       []*entities.Event `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *GetTransactionResultResponse) Reset() {
	*x = GetTransactionResultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_legacy_execution_execution_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionResultResponse) ProtoMessage() {}

func (x *GetTransactionResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flow_legacy_execution_execution_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionResultResponse.ProtoReflect.Descriptor instead.
func (*GetTransactionResultResponse) Descriptor() ([]byte, []int) {
	return file_flow_legacy_execution_execution_proto_rawDescGZIP(), []int{9}
}

func (x *GetTransactionResultResponse) GetStatusCode() uint32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetTransactionResultResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *GetTransactionResultResponse) GetEvents() []*entities.Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type GetEventsForBlockIDsResponse_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId     []byte            `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	BlockHeight uint64            `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Events      []*entities.Event `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *GetEventsForBlockIDsResponse_Result) Reset() {
	*x = GetEventsForBlockIDsResponse_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_legacy_execution_execution_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsForBlockIDsResponse_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsForBlockIDsResponse_Result) ProtoMessage() {}

func (x *GetEventsForBlockIDsResponse_Result) ProtoReflect() protoreflect.Message {
	mi := &file_flow_legacy_execution_execution_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsForBlockIDsResponse_Result.ProtoReflect.Descriptor instead.
func (*GetEventsForBlockIDsResponse_Result) Descriptor() ([]byte, []int) {
	return file_flow_legacy_execution_execution_proto_rawDescGZIP(), []int{6, 0}
}

func (x *GetEventsForBlockIDsResponse_Result) GetBlockId() []byte {
	if x != nil {
		return x.BlockId
	}
	return nil
}

func (x *GetEventsForBlockIDsResponse_Result) GetBlockHeight() uint64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *GetEventsForBlockIDsResponse_Result) GetEvents() []*entities.Event {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_flow_legacy_execution_execution_proto protoreflect.FileDescriptor

var file_flow_legacy_execution_execution_proto_rawDesc = []byte{
	0x0a, 0x25, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x22, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6c, 0x65, 0x67,
	0x61, 0x63, 0x79, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x4a, 0x0a, 0x1b, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x70, 0x0a, 0x1d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x41, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72,
	0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x61,
	0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x36, 0x0a, 0x1e, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x41, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0xd9, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x6f,
	0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x48, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47,
	0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x6f, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x49, 0x44, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x6f, 0x0a, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x4e, 0x0a, 0x1b,
	0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x6f, 0x72, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x73, 0x22, 0x5f, 0x0a, 0x1b,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x8d, 0x01,
	0x0a, 0x1c, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xee, 0x03,
	0x0a, 0x0c, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x50, 0x49, 0x12, 0x37,
	0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x12, 0x25,
	0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6d, 0x0a,
	0x16, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x41, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x12, 0x28, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x41, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x41, 0x74, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x6f, 0x72, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x49, 0x44, 0x73, 0x12, 0x26, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x6f, 0x72, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x46, 0x6f, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x26, 0x2e,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x60,
	0x0a, 0x24, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x6e, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2e, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6f, 0x6e, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flow_legacy_execution_execution_proto_rawDescOnce sync.Once
	file_flow_legacy_execution_execution_proto_rawDescData = file_flow_legacy_execution_execution_proto_rawDesc
)

func file_flow_legacy_execution_execution_proto_rawDescGZIP() []byte {
	file_flow_legacy_execution_execution_proto_rawDescOnce.Do(func() {
		file_flow_legacy_execution_execution_proto_rawDescData = protoimpl.X.CompressGZIP(file_flow_legacy_execution_execution_proto_rawDescData)
	})
	return file_flow_legacy_execution_execution_proto_rawDescData
}

var file_flow_legacy_execution_execution_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_flow_legacy_execution_execution_proto_goTypes = []interface{}{
	(*PingRequest)(nil),                         // 0: execution.PingRequest
	(*PingResponse)(nil),                        // 1: execution.PingResponse
	(*GetAccountAtBlockIDRequest)(nil),          // 2: execution.GetAccountAtBlockIDRequest
	(*GetAccountAtBlockIDResponse)(nil),         // 3: execution.GetAccountAtBlockIDResponse
	(*ExecuteScriptAtBlockIDRequest)(nil),       // 4: execution.ExecuteScriptAtBlockIDRequest
	(*ExecuteScriptAtBlockIDResponse)(nil),      // 5: execution.ExecuteScriptAtBlockIDResponse
	(*GetEventsForBlockIDsResponse)(nil),        // 6: execution.GetEventsForBlockIDsResponse
	(*GetEventsForBlockIDsRequest)(nil),         // 7: execution.GetEventsForBlockIDsRequest
	(*GetTransactionResultRequest)(nil),         // 8: execution.GetTransactionResultRequest
	(*GetTransactionResultResponse)(nil),        // 9: execution.GetTransactionResultResponse
	(*GetEventsForBlockIDsResponse_Result)(nil), // 10: execution.GetEventsForBlockIDsResponse.Result
	(*entities.Account)(nil),                    // 11: entities.Account
	(*entities.Event)(nil),                      // 12: entities.Event
}
var file_flow_legacy_execution_execution_proto_depIdxs = []int32{
	11, // 0: execution.GetAccountAtBlockIDResponse.account:type_name -> entities.Account
	10, // 1: execution.GetEventsForBlockIDsResponse.results:type_name -> execution.GetEventsForBlockIDsResponse.Result
	12, // 2: execution.GetTransactionResultResponse.events:type_name -> entities.Event
	12, // 3: execution.GetEventsForBlockIDsResponse.Result.events:type_name -> entities.Event
	0,  // 4: execution.ExecutionAPI.Ping:input_type -> execution.PingRequest
	2,  // 5: execution.ExecutionAPI.GetAccountAtBlockID:input_type -> execution.GetAccountAtBlockIDRequest
	4,  // 6: execution.ExecutionAPI.ExecuteScriptAtBlockID:input_type -> execution.ExecuteScriptAtBlockIDRequest
	7,  // 7: execution.ExecutionAPI.GetEventsForBlockIDs:input_type -> execution.GetEventsForBlockIDsRequest
	8,  // 8: execution.ExecutionAPI.GetTransactionResult:input_type -> execution.GetTransactionResultRequest
	1,  // 9: execution.ExecutionAPI.Ping:output_type -> execution.PingResponse
	3,  // 10: execution.ExecutionAPI.GetAccountAtBlockID:output_type -> execution.GetAccountAtBlockIDResponse
	5,  // 11: execution.ExecutionAPI.ExecuteScriptAtBlockID:output_type -> execution.ExecuteScriptAtBlockIDResponse
	6,  // 12: execution.ExecutionAPI.GetEventsForBlockIDs:output_type -> execution.GetEventsForBlockIDsResponse
	9,  // 13: execution.ExecutionAPI.GetTransactionResult:output_type -> execution.GetTransactionResultResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_flow_legacy_execution_execution_proto_init() }
func file_flow_legacy_execution_execution_proto_init() {
	if File_flow_legacy_execution_execution_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flow_legacy_execution_execution_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_flow_legacy_execution_execution_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_flow_legacy_execution_execution_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountAtBlockIDRequest); i {
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
		file_flow_legacy_execution_execution_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountAtBlockIDResponse); i {
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
		file_flow_legacy_execution_execution_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteScriptAtBlockIDRequest); i {
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
		file_flow_legacy_execution_execution_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteScriptAtBlockIDResponse); i {
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
		file_flow_legacy_execution_execution_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsForBlockIDsResponse); i {
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
		file_flow_legacy_execution_execution_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsForBlockIDsRequest); i {
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
		file_flow_legacy_execution_execution_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionResultRequest); i {
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
		file_flow_legacy_execution_execution_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionResultResponse); i {
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
		file_flow_legacy_execution_execution_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsForBlockIDsResponse_Result); i {
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
			RawDescriptor: file_flow_legacy_execution_execution_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_flow_legacy_execution_execution_proto_goTypes,
		DependencyIndexes: file_flow_legacy_execution_execution_proto_depIdxs,
		MessageInfos:      file_flow_legacy_execution_execution_proto_msgTypes,
	}.Build()
	File_flow_legacy_execution_execution_proto = out.File
	file_flow_legacy_execution_execution_proto_rawDesc = nil
	file_flow_legacy_execution_execution_proto_goTypes = nil
	file_flow_legacy_execution_execution_proto_depIdxs = nil
}
