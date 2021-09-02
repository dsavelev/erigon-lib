// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: txpool/txpool.proto

package txpool

import (
	types "github.com/ledgerwatch/erigon-lib/gointerfaces/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ImportResult int32

const (
	ImportResult_SUCCESS        ImportResult = 0
	ImportResult_ALREADY_EXISTS ImportResult = 1
	ImportResult_FEE_TOO_LOW    ImportResult = 2
	ImportResult_STALE          ImportResult = 3
	ImportResult_INVALID        ImportResult = 4
	ImportResult_INTERNAL_ERROR ImportResult = 5
)

// Enum value maps for ImportResult.
var (
	ImportResult_name = map[int32]string{
		0: "SUCCESS",
		1: "ALREADY_EXISTS",
		2: "FEE_TOO_LOW",
		3: "STALE",
		4: "INVALID",
		5: "INTERNAL_ERROR",
	}
	ImportResult_value = map[string]int32{
		"SUCCESS":        0,
		"ALREADY_EXISTS": 1,
		"FEE_TOO_LOW":    2,
		"STALE":          3,
		"INVALID":        4,
		"INTERNAL_ERROR": 5,
	}
)

func (x ImportResult) Enum() *ImportResult {
	p := new(ImportResult)
	*p = x
	return p
}

func (x ImportResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImportResult) Descriptor() protoreflect.EnumDescriptor {
	return file_txpool_txpool_proto_enumTypes[0].Descriptor()
}

func (ImportResult) Type() protoreflect.EnumType {
	return &file_txpool_txpool_proto_enumTypes[0]
}

func (x ImportResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImportResult.Descriptor instead.
func (ImportResult) EnumDescriptor() ([]byte, []int) {
	return file_txpool_txpool_proto_rawDescGZIP(), []int{0}
}

type AllReply_Type int32

const (
	AllReply_PENDING  AllReply_Type = 0 // All currently processable transactions
	AllReply_QUEUED   AllReply_Type = 1 // Queued but non-processable transactions
	AllReply_BASE_FEE AllReply_Type = 2 // BaseFee not enough baseFee non-processable transactions
)

// Enum value maps for AllReply_Type.
var (
	AllReply_Type_name = map[int32]string{
		0: "PENDING",
		1: "QUEUED",
		2: "BASE_FEE",
	}
	AllReply_Type_value = map[string]int32{
		"PENDING":  0,
		"QUEUED":   1,
		"BASE_FEE": 2,
	}
)

func (x AllReply_Type) Enum() *AllReply_Type {
	p := new(AllReply_Type)
	*p = x
	return p
}

func (x AllReply_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AllReply_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_txpool_txpool_proto_enumTypes[1].Descriptor()
}

func (AllReply_Type) Type() protoreflect.EnumType {
	return &file_txpool_txpool_proto_enumTypes[1]
}

func (x AllReply_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AllReply_Type.Descriptor instead.
func (AllReply_Type) EnumDescriptor() ([]byte, []int) {
	return file_txpool_txpool_proto_rawDescGZIP(), []int{8, 0}
}

type TxHashes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hashes []*types.H256 `protobuf:"bytes,1,rep,name=hashes,proto3" json:"hashes,omitempty"`
}

func (x *TxHashes) Reset() {
	*x = TxHashes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txpool_txpool_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxHashes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxHashes) ProtoMessage() {}

func (x *TxHashes) ProtoReflect() protoreflect.Message {
	mi := &file_txpool_txpool_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxHashes.ProtoReflect.Descriptor instead.
func (*TxHashes) Descriptor() ([]byte, []int) {
	return file_txpool_txpool_proto_rawDescGZIP(), []int{0}
}

func (x *TxHashes) GetHashes() []*types.H256 {
	if x != nil {
		return x.Hashes
	}
	return nil
}

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RlpTxs [][]byte `protobuf:"bytes,1,rep,name=rlpTxs,proto3" json:"rlpTxs,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txpool_txpool_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_txpool_txpool_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_txpool_txpool_proto_rawDescGZIP(), []int{1}
}

func (x *AddRequest) GetRlpTxs() [][]byte {
	if x != nil {
		return x.RlpTxs
	}
	return nil
}

type AddReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Imported []ImportResult `protobuf:"varint,1,rep,packed,name=imported,proto3,enum=txpool.ImportResult" json:"imported,omitempty"`
	Errors   []string       `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *AddReply) Reset() {
	*x = AddReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txpool_txpool_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReply) ProtoMessage() {}

func (x *AddReply) ProtoReflect() protoreflect.Message {
	mi := &file_txpool_txpool_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReply.ProtoReflect.Descriptor instead.
func (*AddReply) Descriptor() ([]byte, []int) {
	return file_txpool_txpool_proto_rawDescGZIP(), []int{2}
}

func (x *AddReply) GetImported() []ImportResult {
	if x != nil {
		return x.Imported
	}
	return nil
}

func (x *AddReply) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

type TransactionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hashes []*types.H256 `protobuf:"bytes,1,rep,name=hashes,proto3" json:"hashes,omitempty"`
}

func (x *TransactionsRequest) Reset() {
	*x = TransactionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txpool_txpool_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionsRequest) ProtoMessage() {}

func (x *TransactionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_txpool_txpool_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionsRequest.ProtoReflect.Descriptor instead.
func (*TransactionsRequest) Descriptor() ([]byte, []int) {
	return file_txpool_txpool_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionsRequest) GetHashes() []*types.H256 {
	if x != nil {
		return x.Hashes
	}
	return nil
}

type TransactionsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RlpTxs [][]byte `protobuf:"bytes,1,rep,name=rlpTxs,proto3" json:"rlpTxs,omitempty"`
}

func (x *TransactionsReply) Reset() {
	*x = TransactionsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txpool_txpool_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionsReply) ProtoMessage() {}

func (x *TransactionsReply) ProtoReflect() protoreflect.Message {
	mi := &file_txpool_txpool_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionsReply.ProtoReflect.Descriptor instead.
func (*TransactionsReply) Descriptor() ([]byte, []int) {
	return file_txpool_txpool_proto_rawDescGZIP(), []int{4}
}

func (x *TransactionsReply) GetRlpTxs() [][]byte {
	if x != nil {
		return x.RlpTxs
	}
	return nil
}

type OnAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OnAddRequest) Reset() {
	*x = OnAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txpool_txpool_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnAddRequest) ProtoMessage() {}

func (x *OnAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_txpool_txpool_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnAddRequest.ProtoReflect.Descriptor instead.
func (*OnAddRequest) Descriptor() ([]byte, []int) {
	return file_txpool_txpool_proto_rawDescGZIP(), []int{5}
}

type OnAddReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RplTxs [][]byte `protobuf:"bytes,1,rep,name=rplTxs,proto3" json:"rplTxs,omitempty"`
}

func (x *OnAddReply) Reset() {
	*x = OnAddReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txpool_txpool_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnAddReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnAddReply) ProtoMessage() {}

func (x *OnAddReply) ProtoReflect() protoreflect.Message {
	mi := &file_txpool_txpool_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnAddReply.ProtoReflect.Descriptor instead.
func (*OnAddReply) Descriptor() ([]byte, []int) {
	return file_txpool_txpool_proto_rawDescGZIP(), []int{6}
}

func (x *OnAddReply) GetRplTxs() [][]byte {
	if x != nil {
		return x.RplTxs
	}
	return nil
}

type AllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AllRequest) Reset() {
	*x = AllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txpool_txpool_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllRequest) ProtoMessage() {}

func (x *AllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_txpool_txpool_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllRequest.ProtoReflect.Descriptor instead.
func (*AllRequest) Descriptor() ([]byte, []int) {
	return file_txpool_txpool_proto_rawDescGZIP(), []int{7}
}

type AllReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txs []*AllReply_Tx `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (x *AllReply) Reset() {
	*x = AllReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txpool_txpool_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllReply) ProtoMessage() {}

func (x *AllReply) ProtoReflect() protoreflect.Message {
	mi := &file_txpool_txpool_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllReply.ProtoReflect.Descriptor instead.
func (*AllReply) Descriptor() ([]byte, []int) {
	return file_txpool_txpool_proto_rawDescGZIP(), []int{8}
}

func (x *AllReply) GetTxs() []*AllReply_Tx {
	if x != nil {
		return x.Txs
	}
	return nil
}

type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txpool_txpool_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_txpool_txpool_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_txpool_txpool_proto_rawDescGZIP(), []int{9}
}

type StatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PendingCount uint32 `protobuf:"varint,1,opt,name=pendingCount,proto3" json:"pendingCount,omitempty"`
	QueuedCount  uint32 `protobuf:"varint,2,opt,name=queuedCount,proto3" json:"queuedCount,omitempty"`
}

func (x *StatusReply) Reset() {
	*x = StatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txpool_txpool_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusReply) ProtoMessage() {}

func (x *StatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_txpool_txpool_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusReply.ProtoReflect.Descriptor instead.
func (*StatusReply) Descriptor() ([]byte, []int) {
	return file_txpool_txpool_proto_rawDescGZIP(), []int{10}
}

func (x *StatusReply) GetPendingCount() uint32 {
	if x != nil {
		return x.PendingCount
	}
	return 0
}

func (x *StatusReply) GetQueuedCount() uint32 {
	if x != nil {
		return x.QueuedCount
	}
	return 0
}

type AllReply_Tx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   AllReply_Type `protobuf:"varint,1,opt,name=type,proto3,enum=txpool.AllReply_Type" json:"type,omitempty"`
	Sender []byte        `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	RlpTx  []byte        `protobuf:"bytes,3,opt,name=rlpTx,proto3" json:"rlpTx,omitempty"`
}

func (x *AllReply_Tx) Reset() {
	*x = AllReply_Tx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txpool_txpool_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllReply_Tx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllReply_Tx) ProtoMessage() {}

func (x *AllReply_Tx) ProtoReflect() protoreflect.Message {
	mi := &file_txpool_txpool_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllReply_Tx.ProtoReflect.Descriptor instead.
func (*AllReply_Tx) Descriptor() ([]byte, []int) {
	return file_txpool_txpool_proto_rawDescGZIP(), []int{8, 0}
}

func (x *AllReply_Tx) GetType() AllReply_Type {
	if x != nil {
		return x.Type
	}
	return AllReply_PENDING
}

func (x *AllReply_Tx) GetSender() []byte {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *AllReply_Tx) GetRlpTx() []byte {
	if x != nil {
		return x.RlpTx
	}
	return nil
}

var File_txpool_txpool_proto protoreflect.FileDescriptor

var file_txpool_txpool_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a,
	0x08, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x06, 0x68, 0x61, 0x73,
	0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x48, 0x32, 0x35, 0x36, 0x52, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x22, 0x24,
	0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x6c, 0x70, 0x54, 0x78, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x6c,
	0x70, 0x54, 0x78, 0x73, 0x22, 0x54, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x30, 0x0a, 0x08, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x49, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x08, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x3a, 0x0a, 0x13, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x23, 0x0a, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x48, 0x32, 0x35, 0x36, 0x52, 0x06,
	0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x22, 0x2b, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x6c, 0x70, 0x54, 0x78, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x6c, 0x70,
	0x54, 0x78, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x4f, 0x6e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x24, 0x0a, 0x0a, 0x4f, 0x6e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x70, 0x6c, 0x54, 0x78, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0c, 0x52, 0x06, 0x72, 0x70, 0x6c, 0x54, 0x78, 0x73, 0x22, 0x0c, 0x0a, 0x0a, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xbf, 0x01, 0x0a, 0x08, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x03, 0x74, 0x78, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x2e, 0x54, 0x78, 0x52, 0x03, 0x74, 0x78, 0x73, 0x1a, 0x5d, 0x0a, 0x02, 0x54,
	0x78, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6c, 0x70, 0x54, 0x78, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x72, 0x6c, 0x70, 0x54, 0x78, 0x22, 0x2d, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x51, 0x55, 0x45, 0x55, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x42,
	0x41, 0x53, 0x45, 0x5f, 0x46, 0x45, 0x45, 0x10, 0x02, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x53, 0x0a, 0x0b, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0c, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x2a,
	0x6c, 0x0a, 0x0c, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e,
	0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x45, 0x45, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x4f, 0x57, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x54, 0x41, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x54,
	0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x32, 0x80, 0x03,
	0x0a, 0x06, 0x54, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x12, 0x36, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x31, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x12,
	0x10, 0x2e, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x65,
	0x73, 0x1a, 0x10, 0x2e, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x54, 0x78, 0x48, 0x61, 0x73,
	0x68, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x12, 0x2e, 0x74, 0x78, 0x70,
	0x6f, 0x6f, 0x6c, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x46, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x1b, 0x2e, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x12,
	0x12, 0x2e, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x4f, 0x6e, 0x41, 0x64, 0x64, 0x12, 0x14,
	0x2e, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x4f, 0x6e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x4f, 0x6e,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x30, 0x01, 0x12, 0x34, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x15, 0x2e, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x74, 0x78,
	0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x3b, 0x74, 0x78, 0x70,
	0x6f, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_txpool_txpool_proto_rawDescOnce sync.Once
	file_txpool_txpool_proto_rawDescData = file_txpool_txpool_proto_rawDesc
)

func file_txpool_txpool_proto_rawDescGZIP() []byte {
	file_txpool_txpool_proto_rawDescOnce.Do(func() {
		file_txpool_txpool_proto_rawDescData = protoimpl.X.CompressGZIP(file_txpool_txpool_proto_rawDescData)
	})
	return file_txpool_txpool_proto_rawDescData
}

var file_txpool_txpool_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_txpool_txpool_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_txpool_txpool_proto_goTypes = []interface{}{
	(ImportResult)(0),           // 0: txpool.ImportResult
	(AllReply_Type)(0),          // 1: txpool.AllReply.Type
	(*TxHashes)(nil),            // 2: txpool.TxHashes
	(*AddRequest)(nil),          // 3: txpool.AddRequest
	(*AddReply)(nil),            // 4: txpool.AddReply
	(*TransactionsRequest)(nil), // 5: txpool.TransactionsRequest
	(*TransactionsReply)(nil),   // 6: txpool.TransactionsReply
	(*OnAddRequest)(nil),        // 7: txpool.OnAddRequest
	(*OnAddReply)(nil),          // 8: txpool.OnAddReply
	(*AllRequest)(nil),          // 9: txpool.AllRequest
	(*AllReply)(nil),            // 10: txpool.AllReply
	(*StatusRequest)(nil),       // 11: txpool.StatusRequest
	(*StatusReply)(nil),         // 12: txpool.StatusReply
	(*AllReply_Tx)(nil),         // 13: txpool.AllReply.Tx
	(*types.H256)(nil),          // 14: types.H256
	(*emptypb.Empty)(nil),       // 15: google.protobuf.Empty
	(*types.VersionReply)(nil),  // 16: types.VersionReply
}
var file_txpool_txpool_proto_depIdxs = []int32{
	14, // 0: txpool.TxHashes.hashes:type_name -> types.H256
	0,  // 1: txpool.AddReply.imported:type_name -> txpool.ImportResult
	14, // 2: txpool.TransactionsRequest.hashes:type_name -> types.H256
	13, // 3: txpool.AllReply.txs:type_name -> txpool.AllReply.Tx
	1,  // 4: txpool.AllReply.Tx.type:type_name -> txpool.AllReply.Type
	15, // 5: txpool.Txpool.Version:input_type -> google.protobuf.Empty
	2,  // 6: txpool.Txpool.FindUnknown:input_type -> txpool.TxHashes
	3,  // 7: txpool.Txpool.Add:input_type -> txpool.AddRequest
	5,  // 8: txpool.Txpool.Transactions:input_type -> txpool.TransactionsRequest
	9,  // 9: txpool.Txpool.All:input_type -> txpool.AllRequest
	7,  // 10: txpool.Txpool.OnAdd:input_type -> txpool.OnAddRequest
	11, // 11: txpool.Txpool.Status:input_type -> txpool.StatusRequest
	16, // 12: txpool.Txpool.Version:output_type -> types.VersionReply
	2,  // 13: txpool.Txpool.FindUnknown:output_type -> txpool.TxHashes
	4,  // 14: txpool.Txpool.Add:output_type -> txpool.AddReply
	6,  // 15: txpool.Txpool.Transactions:output_type -> txpool.TransactionsReply
	10, // 16: txpool.Txpool.All:output_type -> txpool.AllReply
	8,  // 17: txpool.Txpool.OnAdd:output_type -> txpool.OnAddReply
	12, // 18: txpool.Txpool.Status:output_type -> txpool.StatusReply
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_txpool_txpool_proto_init() }
func file_txpool_txpool_proto_init() {
	if File_txpool_txpool_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_txpool_txpool_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxHashes); i {
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
		file_txpool_txpool_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_txpool_txpool_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReply); i {
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
		file_txpool_txpool_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionsRequest); i {
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
		file_txpool_txpool_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionsReply); i {
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
		file_txpool_txpool_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnAddRequest); i {
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
		file_txpool_txpool_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnAddReply); i {
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
		file_txpool_txpool_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllRequest); i {
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
		file_txpool_txpool_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllReply); i {
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
		file_txpool_txpool_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusRequest); i {
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
		file_txpool_txpool_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusReply); i {
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
		file_txpool_txpool_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllReply_Tx); i {
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
			RawDescriptor: file_txpool_txpool_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_txpool_txpool_proto_goTypes,
		DependencyIndexes: file_txpool_txpool_proto_depIdxs,
		EnumInfos:         file_txpool_txpool_proto_enumTypes,
		MessageInfos:      file_txpool_txpool_proto_msgTypes,
	}.Build()
	File_txpool_txpool_proto = out.File
	file_txpool_txpool_proto_rawDesc = nil
	file_txpool_txpool_proto_goTypes = nil
	file_txpool_txpool_proto_depIdxs = nil
}
