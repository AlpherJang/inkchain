// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/chaincode_shim.proto

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ChaincodeMessage_Type int32

const (
	ChaincodeMessage_UNDEFINED           ChaincodeMessage_Type = 0
	ChaincodeMessage_REGISTER            ChaincodeMessage_Type = 1
	ChaincodeMessage_REGISTERED          ChaincodeMessage_Type = 2
	ChaincodeMessage_INIT                ChaincodeMessage_Type = 3
	ChaincodeMessage_READY               ChaincodeMessage_Type = 4
	ChaincodeMessage_TRANSACTION         ChaincodeMessage_Type = 5
	ChaincodeMessage_COMPLETED           ChaincodeMessage_Type = 6
	ChaincodeMessage_ERROR               ChaincodeMessage_Type = 7
	ChaincodeMessage_GET_STATE           ChaincodeMessage_Type = 8
	ChaincodeMessage_PUT_STATE           ChaincodeMessage_Type = 9
	ChaincodeMessage_DEL_STATE           ChaincodeMessage_Type = 10
	ChaincodeMessage_INVOKE_CHAINCODE    ChaincodeMessage_Type = 11
	ChaincodeMessage_RESPONSE            ChaincodeMessage_Type = 13
	ChaincodeMessage_GET_STATE_BY_RANGE  ChaincodeMessage_Type = 14
	ChaincodeMessage_GET_QUERY_RESULT    ChaincodeMessage_Type = 15
	ChaincodeMessage_QUERY_STATE_NEXT    ChaincodeMessage_Type = 16
	ChaincodeMessage_QUERY_STATE_CLOSE   ChaincodeMessage_Type = 17
	ChaincodeMessage_KEEPALIVE           ChaincodeMessage_Type = 18
	ChaincodeMessage_GET_HISTORY_FOR_KEY ChaincodeMessage_Type = 19
	ChaincodeMessage_TRANSFER            ChaincodeMessage_Type = 820
	ChaincodeMessage_ISSUE_TOKEN         ChaincodeMessage_Type = 821
	ChaincodeMessage_GET_ACCOUNT         ChaincodeMessage_Type = 823
	ChaincodeMessage_CROSS_TRANSFER      ChaincodeMessage_Type = 824
	ChaincodeMessage_GET_FEE             ChaincodeMessage_Type = 825
)

var ChaincodeMessage_Type_name = map[int32]string{
	0:   "UNDEFINED",
	1:   "REGISTER",
	2:   "REGISTERED",
	3:   "INIT",
	4:   "READY",
	5:   "TRANSACTION",
	6:   "COMPLETED",
	7:   "ERROR",
	8:   "GET_STATE",
	9:   "PUT_STATE",
	10:  "DEL_STATE",
	11:  "INVOKE_CHAINCODE",
	13:  "RESPONSE",
	14:  "GET_STATE_BY_RANGE",
	15:  "GET_QUERY_RESULT",
	16:  "QUERY_STATE_NEXT",
	17:  "QUERY_STATE_CLOSE",
	18:  "KEEPALIVE",
	19:  "GET_HISTORY_FOR_KEY",
	820: "TRANSFER",
	821: "ISSUE_TOKEN",
	823: "GET_ACCOUNT",
	824: "CROSS_TRANSFER",
	825: "GET_FEE",
}
var ChaincodeMessage_Type_value = map[string]int32{
	"UNDEFINED":           0,
	"REGISTER":            1,
	"REGISTERED":          2,
	"INIT":                3,
	"READY":               4,
	"TRANSACTION":         5,
	"COMPLETED":           6,
	"ERROR":               7,
	"GET_STATE":           8,
	"PUT_STATE":           9,
	"DEL_STATE":           10,
	"INVOKE_CHAINCODE":    11,
	"RESPONSE":            13,
	"GET_STATE_BY_RANGE":  14,
	"GET_QUERY_RESULT":    15,
	"QUERY_STATE_NEXT":    16,
	"QUERY_STATE_CLOSE":   17,
	"KEEPALIVE":           18,
	"GET_HISTORY_FOR_KEY": 19,
	"TRANSFER":            820,
	"ISSUE_TOKEN":         821,
	"GET_ACCOUNT":         823,
	"CROSS_TRANSFER":      824,
	"GET_FEE":             825,
}

func (x ChaincodeMessage_Type) String() string {
	return proto.EnumName(ChaincodeMessage_Type_name, int32(x))
}
func (ChaincodeMessage_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

type ChaincodeMessage struct {
	Type      ChaincodeMessage_Type       `protobuf:"varint,1,opt,name=type,enum=protos.ChaincodeMessage_Type" json:"type,omitempty"`
	Timestamp *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Payload   []byte                      `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Txid      string                      `protobuf:"bytes,4,opt,name=txid" json:"txid,omitempty"`
	Proposal  *SignedProposal             `protobuf:"bytes,5,opt,name=proposal" json:"proposal,omitempty"`
	// event emmited by chaincode. Used only with Init or Invoke.
	// This event is then stored (currently)
	// with Block.NonHashData.TransactionResult
	ChaincodeEvent *ChaincodeEvent `protobuf:"bytes,6,opt,name=chaincode_event,json=chaincodeEvent" json:"chaincode_event,omitempty"`
}

func (m *ChaincodeMessage) Reset()                    { *m = ChaincodeMessage{} }
func (m *ChaincodeMessage) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeMessage) ProtoMessage()               {}
func (*ChaincodeMessage) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *ChaincodeMessage) GetType() ChaincodeMessage_Type {
	if m != nil {
		return m.Type
	}
	return ChaincodeMessage_UNDEFINED
}

func (m *ChaincodeMessage) GetTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ChaincodeMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ChaincodeMessage) GetTxid() string {
	if m != nil {
		return m.Txid
	}
	return ""
}

func (m *ChaincodeMessage) GetProposal() *SignedProposal {
	if m != nil {
		return m.Proposal
	}
	return nil
}

func (m *ChaincodeMessage) GetChaincodeEvent() *ChaincodeEvent {
	if m != nil {
		return m.ChaincodeEvent
	}
	return nil
}

type PutStateInfo struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *PutStateInfo) Reset()                    { *m = PutStateInfo{} }
func (m *PutStateInfo) String() string            { return proto.CompactTextString(m) }
func (*PutStateInfo) ProtoMessage()               {}
func (*PutStateInfo) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *PutStateInfo) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PutStateInfo) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type GetStateByRange struct {
	StartKey string `protobuf:"bytes,1,opt,name=startKey" json:"startKey,omitempty"`
	EndKey   string `protobuf:"bytes,2,opt,name=endKey" json:"endKey,omitempty"`
}

func (m *GetStateByRange) Reset()                    { *m = GetStateByRange{} }
func (m *GetStateByRange) String() string            { return proto.CompactTextString(m) }
func (*GetStateByRange) ProtoMessage()               {}
func (*GetStateByRange) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *GetStateByRange) GetStartKey() string {
	if m != nil {
		return m.StartKey
	}
	return ""
}

func (m *GetStateByRange) GetEndKey() string {
	if m != nil {
		return m.EndKey
	}
	return ""
}

type GetQueryResult struct {
	Query string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
}

func (m *GetQueryResult) Reset()                    { *m = GetQueryResult{} }
func (m *GetQueryResult) String() string            { return proto.CompactTextString(m) }
func (*GetQueryResult) ProtoMessage()               {}
func (*GetQueryResult) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *GetQueryResult) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type GetHistoryForKey struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *GetHistoryForKey) Reset()                    { *m = GetHistoryForKey{} }
func (m *GetHistoryForKey) String() string            { return proto.CompactTextString(m) }
func (*GetHistoryForKey) ProtoMessage()               {}
func (*GetHistoryForKey) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *GetHistoryForKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type QueryStateNext struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *QueryStateNext) Reset()                    { *m = QueryStateNext{} }
func (m *QueryStateNext) String() string            { return proto.CompactTextString(m) }
func (*QueryStateNext) ProtoMessage()               {}
func (*QueryStateNext) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *QueryStateNext) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type QueryStateClose struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *QueryStateClose) Reset()                    { *m = QueryStateClose{} }
func (m *QueryStateClose) String() string            { return proto.CompactTextString(m) }
func (*QueryStateClose) ProtoMessage()               {}
func (*QueryStateClose) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *QueryStateClose) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type QueryResultBytes struct {
	ResultBytes []byte `protobuf:"bytes,1,opt,name=resultBytes,proto3" json:"resultBytes,omitempty"`
}

func (m *QueryResultBytes) Reset()                    { *m = QueryResultBytes{} }
func (m *QueryResultBytes) String() string            { return proto.CompactTextString(m) }
func (*QueryResultBytes) ProtoMessage()               {}
func (*QueryResultBytes) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *QueryResultBytes) GetResultBytes() []byte {
	if m != nil {
		return m.ResultBytes
	}
	return nil
}

type QueryResponse struct {
	Results []*QueryResultBytes `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
	HasMore bool                `protobuf:"varint,2,opt,name=has_more,json=hasMore" json:"has_more,omitempty"`
	Id      string              `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
}

func (m *QueryResponse) Reset()                    { *m = QueryResponse{} }
func (m *QueryResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()               {}
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *QueryResponse) GetResults() []*QueryResultBytes {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *QueryResponse) GetHasMore() bool {
	if m != nil {
		return m.HasMore
	}
	return false
}

func (m *QueryResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CrossTransfer struct {
	To     []byte `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Amount []byte `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *CrossTransfer) Reset()                    { *m = CrossTransfer{} }
func (m *CrossTransfer) String() string            { return proto.CompactTextString(m) }
func (*CrossTransfer) ProtoMessage()               {}
func (*CrossTransfer) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

func (m *CrossTransfer) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *CrossTransfer) GetAmount() []byte {
	if m != nil {
		return m.Amount
	}
	return nil
}

type CrossTransferInfo struct {
	TranSet      []*CrossTransfer `protobuf:"bytes,1,rep,name=tranSet" json:"tranSet,omitempty"`
	PubTxId      []byte           `protobuf:"bytes,2,opt,name=pubTxId,proto3" json:"pubTxId,omitempty"`
	FromPlatForm []byte           `protobuf:"bytes,3,opt,name=fromPlatForm,proto3" json:"fromPlatForm,omitempty"`
	FromAccount  []byte           `protobuf:"bytes,4,opt,name=from_account,json=fromAccount,proto3" json:"from_account,omitempty"`
	BalanceType  []byte           `protobuf:"bytes,5,opt,name=balance_type,json=balanceType,proto3" json:"balance_type,omitempty"`
}

func (m *CrossTransferInfo) Reset()                    { *m = CrossTransferInfo{} }
func (m *CrossTransferInfo) String() string            { return proto.CompactTextString(m) }
func (*CrossTransferInfo) ProtoMessage()               {}
func (*CrossTransferInfo) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{10} }

func (m *CrossTransferInfo) GetTranSet() []*CrossTransfer {
	if m != nil {
		return m.TranSet
	}
	return nil
}

func (m *CrossTransferInfo) GetPubTxId() []byte {
	if m != nil {
		return m.PubTxId
	}
	return nil
}

func (m *CrossTransferInfo) GetFromPlatForm() []byte {
	if m != nil {
		return m.FromPlatForm
	}
	return nil
}

func (m *CrossTransferInfo) GetFromAccount() []byte {
	if m != nil {
		return m.FromAccount
	}
	return nil
}

func (m *CrossTransferInfo) GetBalanceType() []byte {
	if m != nil {
		return m.BalanceType
	}
	return nil
}

type Transfer struct {
	To          []byte `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	BalanceType []byte `protobuf:"bytes,2,opt,name=balanceType,proto3" json:"balanceType,omitempty"`
	Amount      []byte `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *Transfer) Reset()                    { *m = Transfer{} }
func (m *Transfer) String() string            { return proto.CompactTextString(m) }
func (*Transfer) ProtoMessage()               {}
func (*Transfer) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{11} }

func (m *Transfer) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *Transfer) GetBalanceType() []byte {
	if m != nil {
		return m.BalanceType
	}
	return nil
}

func (m *Transfer) GetAmount() []byte {
	if m != nil {
		return m.Amount
	}
	return nil
}

type TransferInfo struct {
	TranSet []*Transfer `protobuf:"bytes,1,rep,name=tranSet" json:"tranSet,omitempty"`
}

func (m *TransferInfo) Reset()                    { *m = TransferInfo{} }
func (m *TransferInfo) String() string            { return proto.CompactTextString(m) }
func (*TransferInfo) ProtoMessage()               {}
func (*TransferInfo) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{12} }

func (m *TransferInfo) GetTranSet() []*Transfer {
	if m != nil {
		return m.TranSet
	}
	return nil
}

type IssueTokenInfo struct {
	Address     []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	BalanceType []byte `protobuf:"bytes,2,opt,name=balanceType,proto3" json:"balanceType,omitempty"`
	Amount      []byte `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *IssueTokenInfo) Reset()                    { *m = IssueTokenInfo{} }
func (m *IssueTokenInfo) String() string            { return proto.CompactTextString(m) }
func (*IssueTokenInfo) ProtoMessage()               {}
func (*IssueTokenInfo) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{13} }

func (m *IssueTokenInfo) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *IssueTokenInfo) GetBalanceType() []byte {
	if m != nil {
		return m.BalanceType
	}
	return nil
}

func (m *IssueTokenInfo) GetAmount() []byte {
	if m != nil {
		return m.Amount
	}
	return nil
}

func init() {
	proto.RegisterType((*ChaincodeMessage)(nil), "protos.ChaincodeMessage")
	proto.RegisterType((*PutStateInfo)(nil), "protos.PutStateInfo")
	proto.RegisterType((*GetStateByRange)(nil), "protos.GetStateByRange")
	proto.RegisterType((*GetQueryResult)(nil), "protos.GetQueryResult")
	proto.RegisterType((*GetHistoryForKey)(nil), "protos.GetHistoryForKey")
	proto.RegisterType((*QueryStateNext)(nil), "protos.QueryStateNext")
	proto.RegisterType((*QueryStateClose)(nil), "protos.QueryStateClose")
	proto.RegisterType((*QueryResultBytes)(nil), "protos.QueryResultBytes")
	proto.RegisterType((*QueryResponse)(nil), "protos.QueryResponse")
	proto.RegisterType((*CrossTransfer)(nil), "protos.CrossTransfer")
	proto.RegisterType((*CrossTransferInfo)(nil), "protos.CrossTransferInfo")
	proto.RegisterType((*Transfer)(nil), "protos.Transfer")
	proto.RegisterType((*TransferInfo)(nil), "protos.TransferInfo")
	proto.RegisterType((*IssueTokenInfo)(nil), "protos.IssueTokenInfo")
	proto.RegisterEnum("protos.ChaincodeMessage_Type", ChaincodeMessage_Type_name, ChaincodeMessage_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ChaincodeSupport service

type ChaincodeSupportClient interface {
	Register(ctx context.Context, opts ...grpc.CallOption) (ChaincodeSupport_RegisterClient, error)
}

type chaincodeSupportClient struct {
	cc *grpc.ClientConn
}

func NewChaincodeSupportClient(cc *grpc.ClientConn) ChaincodeSupportClient {
	return &chaincodeSupportClient{cc}
}

func (c *chaincodeSupportClient) Register(ctx context.Context, opts ...grpc.CallOption) (ChaincodeSupport_RegisterClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ChaincodeSupport_serviceDesc.Streams[0], c.cc, "/protos.ChaincodeSupport/Register", opts...)
	if err != nil {
		return nil, err
	}
	x := &chaincodeSupportRegisterClient{stream}
	return x, nil
}

type ChaincodeSupport_RegisterClient interface {
	Send(*ChaincodeMessage) error
	Recv() (*ChaincodeMessage, error)
	grpc.ClientStream
}

type chaincodeSupportRegisterClient struct {
	grpc.ClientStream
}

func (x *chaincodeSupportRegisterClient) Send(m *ChaincodeMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chaincodeSupportRegisterClient) Recv() (*ChaincodeMessage, error) {
	m := new(ChaincodeMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ChaincodeSupport service

type ChaincodeSupportServer interface {
	Register(ChaincodeSupport_RegisterServer) error
}

func RegisterChaincodeSupportServer(s *grpc.Server, srv ChaincodeSupportServer) {
	s.RegisterService(&_ChaincodeSupport_serviceDesc, srv)
}

func _ChaincodeSupport_Register_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChaincodeSupportServer).Register(&chaincodeSupportRegisterServer{stream})
}

type ChaincodeSupport_RegisterServer interface {
	Send(*ChaincodeMessage) error
	Recv() (*ChaincodeMessage, error)
	grpc.ServerStream
}

type chaincodeSupportRegisterServer struct {
	grpc.ServerStream
}

func (x *chaincodeSupportRegisterServer) Send(m *ChaincodeMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chaincodeSupportRegisterServer) Recv() (*ChaincodeMessage, error) {
	m := new(ChaincodeMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ChaincodeSupport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ChaincodeSupport",
	HandlerType: (*ChaincodeSupportServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Register",
			Handler:       _ChaincodeSupport_Register_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "peer/chaincode_shim.proto",
}

func init() { proto.RegisterFile("peer/chaincode_shim.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 1012 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x5f, 0x6f, 0xe2, 0xc6,
	0x17, 0x5d, 0xfe, 0xc3, 0x85, 0x90, 0xd9, 0xc9, 0x6e, 0x7e, 0x2c, 0xd2, 0x4f, 0x65, 0xad, 0xaa,
	0x42, 0xfb, 0x00, 0x5d, 0x5a, 0xb5, 0x55, 0x5f, 0x2a, 0x02, 0x03, 0xb1, 0x92, 0xd8, 0xec, 0xd8,
	0xac, 0x36, 0x7d, 0xb1, 0x0c, 0x9e, 0x10, 0x2b, 0xe0, 0xa1, 0x9e, 0x61, 0x15, 0x3e, 0x4c, 0xbf,
	0x45, 0xab, 0xb6, 0xaf, 0x7d, 0xea, 0xc7, 0xaa, 0xc6, 0x7f, 0x08, 0x24, 0xca, 0x4b, 0x9f, 0xe0,
	0x9c, 0x7b, 0xce, 0x9d, 0x33, 0xd7, 0xf6, 0x0c, 0xbc, 0x59, 0x33, 0x16, 0x76, 0xe7, 0xb7, 0xae,
	0x1f, 0xcc, 0xb9, 0xc7, 0x1c, 0x71, 0xeb, 0xaf, 0x3a, 0xeb, 0x90, 0x4b, 0x8e, 0x8b, 0xd1, 0x8f,
	0x68, 0x36, 0x1f, 0x49, 0xd8, 0x67, 0x16, 0xc8, 0x58, 0xd3, 0x3c, 0x89, 0x6a, 0xeb, 0x90, 0xaf,
	0xb9, 0x70, 0x97, 0x09, 0xf9, 0xc5, 0x82, 0xf3, 0xc5, 0x92, 0x75, 0x23, 0x34, 0xdb, 0xdc, 0x74,
	0xa5, 0xbf, 0x62, 0x42, 0xba, 0xab, 0x75, 0x2c, 0xd0, 0xfe, 0x2e, 0x00, 0x1a, 0xa4, 0xfd, 0xae,
	0x98, 0x10, 0xee, 0x82, 0xe1, 0xf7, 0x90, 0x97, 0xdb, 0x35, 0x6b, 0x64, 0x5a, 0x99, 0x76, 0xbd,
	0xf7, 0xff, 0x58, 0x2a, 0x3a, 0x8f, 0x75, 0x1d, 0x7b, 0xbb, 0x66, 0x34, 0x92, 0xe2, 0x1f, 0xa0,
	0xb2, 0x6b, 0xdd, 0xc8, 0xb6, 0x32, 0xed, 0x6a, 0xaf, 0xd9, 0x89, 0x17, 0xef, 0xa4, 0x8b, 0x77,
	0xec, 0x54, 0x41, 0x1f, 0xc4, 0xb8, 0x01, 0xa5, 0xb5, 0xbb, 0x5d, 0x72, 0xd7, 0x6b, 0xe4, 0x5a,
	0x99, 0x76, 0x8d, 0xa6, 0x10, 0x63, 0xc8, 0xcb, 0x7b, 0xdf, 0x6b, 0xe4, 0x5b, 0x99, 0x76, 0x85,
	0x46, 0xff, 0x71, 0x0f, 0xca, 0xe9, 0x16, 0x1b, 0x85, 0x68, 0x99, 0xd3, 0x34, 0x9e, 0xe5, 0x2f,
	0x02, 0xe6, 0x4d, 0x92, 0x2a, 0xdd, 0xe9, 0xf0, 0x4f, 0x70, 0xfc, 0x68, 0x64, 0x8d, 0xe2, 0xa1,
	0x75, 0xb7, 0x33, 0xa2, 0xaa, 0xb4, 0x3e, 0x3f, 0xc0, 0xda, 0xaf, 0x39, 0xc8, 0xab, 0xbd, 0xe2,
	0x23, 0xa8, 0x4c, 0x8d, 0x21, 0x19, 0xe9, 0x06, 0x19, 0xa2, 0x17, 0xb8, 0x06, 0x65, 0x4a, 0xc6,
	0xba, 0x65, 0x13, 0x8a, 0x32, 0xb8, 0x0e, 0x90, 0x22, 0x32, 0x44, 0x59, 0x5c, 0x86, 0xbc, 0x6e,
	0xe8, 0x36, 0xca, 0xe1, 0x0a, 0x14, 0x28, 0xe9, 0x0f, 0xaf, 0x51, 0x1e, 0x1f, 0x43, 0xd5, 0xa6,
	0x7d, 0xc3, 0xea, 0x0f, 0x6c, 0xdd, 0x34, 0x50, 0x41, 0xb5, 0x1c, 0x98, 0x57, 0x93, 0x4b, 0x62,
	0x93, 0x21, 0x2a, 0x2a, 0x29, 0xa1, 0xd4, 0xa4, 0xa8, 0xa4, 0x2a, 0x63, 0x62, 0x3b, 0x96, 0xdd,
	0xb7, 0x09, 0x2a, 0x2b, 0x38, 0x99, 0xa6, 0xb0, 0xa2, 0xe0, 0x90, 0x5c, 0x26, 0x10, 0xf0, 0x2b,
	0x40, 0xba, 0xf1, 0xd1, 0xbc, 0x20, 0xce, 0xe0, 0xbc, 0xaf, 0x1b, 0x03, 0x73, 0x48, 0x50, 0x35,
	0x0e, 0x68, 0x4d, 0x4c, 0xc3, 0x22, 0xe8, 0x08, 0x9f, 0x02, 0xde, 0x35, 0x74, 0xce, 0xae, 0x1d,
	0xda, 0x37, 0xc6, 0x04, 0xd5, 0x95, 0x57, 0xf1, 0x1f, 0xa6, 0x84, 0x5e, 0x3b, 0x94, 0x58, 0xd3,
	0x4b, 0x1b, 0x1d, 0x2b, 0x36, 0x66, 0x62, 0xbd, 0x41, 0x3e, 0xd9, 0x08, 0xe1, 0xd7, 0xf0, 0x72,
	0x9f, 0x1d, 0x5c, 0x9a, 0x16, 0x41, 0x2f, 0x55, 0x9a, 0x0b, 0x42, 0x26, 0xfd, 0x4b, 0xfd, 0x23,
	0x41, 0x18, 0xff, 0x0f, 0x4e, 0x54, 0xc7, 0x73, 0xdd, 0xb2, 0x4d, 0x7a, 0xed, 0x8c, 0x4c, 0xea,
	0x5c, 0x90, 0x6b, 0x74, 0x82, 0x8f, 0xa0, 0x1c, 0x6d, 0x7f, 0x44, 0x28, 0xfa, 0xad, 0x88, 0x11,
	0x54, 0x75, 0xcb, 0x9a, 0x12, 0xc7, 0x36, 0x2f, 0x88, 0x81, 0x7e, 0x8f, 0x18, 0xe5, 0xec, 0x0f,
	0x06, 0xe6, 0xd4, 0xb0, 0xd1, 0x1f, 0x45, 0x7c, 0x02, 0xf5, 0x01, 0x35, 0x2d, 0xcb, 0xd9, 0x19,
	0xff, 0x2c, 0xe2, 0x1a, 0x94, 0x94, 0x6c, 0x44, 0x08, 0xfa, 0xab, 0xa8, 0x7d, 0x07, 0xb5, 0xc9,
	0x46, 0x5a, 0xd2, 0x95, 0x4c, 0x0f, 0x6e, 0x38, 0x46, 0x90, 0xbb, 0x63, 0xdb, 0xe8, 0xf5, 0xad,
	0x50, 0xf5, 0x17, 0xbf, 0x82, 0xc2, 0x67, 0x77, 0xb9, 0x61, 0xd1, 0xab, 0x59, 0xa3, 0x31, 0xd0,
	0x08, 0x1c, 0x8f, 0x59, 0xec, 0x3b, 0xdb, 0x52, 0x37, 0x58, 0x30, 0xdc, 0x84, 0xb2, 0x90, 0x6e,
	0x28, 0x2f, 0x76, 0xfe, 0x1d, 0xc6, 0xa7, 0x50, 0x64, 0x81, 0xa7, 0x2a, 0xd9, 0xa8, 0x92, 0x20,
	0xed, 0x2b, 0xa8, 0x8f, 0x99, 0xfc, 0xb0, 0x61, 0xe1, 0x96, 0x32, 0xb1, 0x59, 0x4a, 0xb5, 0xdc,
	0x2f, 0x0a, 0x26, 0x2d, 0x62, 0xa0, 0x7d, 0x09, 0x68, 0xcc, 0xe4, 0xb9, 0x2f, 0x24, 0x0f, 0xb7,
	0x23, 0x1e, 0xaa, 0x9e, 0x4f, 0xa2, 0x6a, 0x2d, 0xa8, 0x47, 0xad, 0xa2, 0x58, 0x06, 0xbb, 0x97,
	0xb8, 0x0e, 0x59, 0xdf, 0x4b, 0x24, 0x59, 0xdf, 0xd3, 0xde, 0xc2, 0xf1, 0x83, 0x62, 0xb0, 0xe4,
	0x82, 0x3d, 0x91, 0x7c, 0x0b, 0x68, 0x2f, 0xcf, 0xd9, 0x56, 0x32, 0x81, 0x5b, 0x50, 0x0d, 0x1f,
	0x60, 0x24, 0xae, 0xd1, 0x7d, 0x4a, 0x0b, 0xe0, 0x28, 0x75, 0xad, 0x79, 0x20, 0x18, 0xee, 0x41,
	0x29, 0xae, 0x2b, 0x79, 0xae, 0x5d, 0xed, 0x35, 0xd2, 0x2f, 0xe6, 0x71, 0x77, 0x9a, 0x0a, 0xf1,
	0x1b, 0x28, 0xdf, 0xba, 0xc2, 0x59, 0xf1, 0x30, 0x9e, 0x76, 0x99, 0x96, 0x6e, 0x5d, 0x71, 0xc5,
	0xc3, 0x34, 0x65, 0x6e, 0x97, 0xf2, 0x7b, 0x38, 0x1a, 0x84, 0x5c, 0x08, 0x3b, 0x74, 0x03, 0x71,
	0xc3, 0x42, 0x25, 0x90, 0x3c, 0x49, 0x96, 0x95, 0x5c, 0x4d, 0xdc, 0x5d, 0xf1, 0x4d, 0x20, 0x93,
	0xe7, 0x96, 0x20, 0xed, 0x9f, 0x0c, 0xbc, 0x3c, 0x70, 0x46, 0x8f, 0xbd, 0x0b, 0x25, 0x19, 0xba,
	0x81, 0xc5, 0x64, 0x92, 0xf6, 0xf5, 0xee, 0xfb, 0xde, 0xd7, 0xd2, 0x54, 0x15, 0x1d, 0x3d, 0x9b,
	0x99, 0x7d, 0xaf, 0x7b, 0x49, 0xff, 0x14, 0x62, 0x0d, 0x6a, 0x37, 0x21, 0x5f, 0x4d, 0x96, 0xae,
	0x1c, 0xf1, 0x70, 0x95, 0x9c, 0x4c, 0x07, 0x1c, 0x7e, 0x1b, 0x6b, 0x1c, 0x77, 0x3e, 0x8f, 0x22,
	0xe6, 0xe3, 0x81, 0x2a, 0xae, 0x1f, 0x53, 0x4a, 0x32, 0x73, 0x97, 0x6e, 0x30, 0x67, 0x4e, 0x74,
	0xa0, 0x16, 0x62, 0x49, 0xc2, 0xa9, 0x23, 0x45, 0xb3, 0xa1, 0xfc, 0xec, 0xf6, 0x5b, 0xb0, 0x2f,
	0x4d, 0x32, 0xee, 0x53, 0x7b, 0x03, 0xca, 0x1d, 0x0c, 0xe8, 0x47, 0xa8, 0x1d, 0x8c, 0xe6, 0xdd,
	0xe3, 0xd1, 0xa0, 0x74, 0x34, 0x4f, 0xa6, 0xa2, 0x79, 0x50, 0xd7, 0x85, 0xd8, 0x30, 0x9b, 0xdf,
	0xb1, 0x20, 0x72, 0x37, 0xa0, 0xe4, 0x7a, 0x5e, 0xc8, 0x44, 0xfa, 0xd6, 0xa4, 0xf0, 0xbf, 0x27,
	0xec, 0x7d, 0xda, 0xbb, 0x77, 0xac, 0xcd, 0x7a, 0xcd, 0x43, 0x89, 0x87, 0x50, 0xa6, 0x6c, 0xe1,
	0x0b, 0xc9, 0x42, 0xdc, 0x78, 0xee, 0xd6, 0x69, 0x3e, 0x5b, 0xd1, 0x5e, 0xb4, 0x33, 0x5f, 0x67,
	0xce, 0x1c, 0x78, 0xc7, 0xc3, 0x45, 0xc7, 0x0f, 0xee, 0x96, 0xee, 0x4c, 0xdc, 0xf0, 0x4d, 0xe0,
	0xb9, 0xd2, 0xe7, 0x81, 0x62, 0xa2, 0x83, 0x3d, 0xf5, 0xab, 0x0b, 0xf3, 0xe7, 0xf7, 0x0b, 0x5f,
	0xde, 0x6e, 0x66, 0x9d, 0x39, 0x5f, 0x75, 0x9f, 0x58, 0xba, 0xa9, 0x25, 0xbe, 0x40, 0x45, 0x57,
	0x59, 0x66, 0xf1, 0x6d, 0xfc, 0xcd, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x69, 0x2c, 0x14, 0xa9,
	0xb1, 0x07, 0x00, 0x00,
}
