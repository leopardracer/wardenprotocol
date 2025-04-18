// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: warden/warden/v1beta3/signature.proto

package v1beta3

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// SignRequestStatus indicates the status of a signature request.
//
// The possible state transitions are:
//   - PENDING -> FULFILLED
//   - PENDING -> REJECTED
type SignRequestStatus int32

const (
	// The request is missing the status field.
	SignRequestStatus_SIGN_REQUEST_STATUS_UNSPECIFIED SignRequestStatus = 0
	// The request is waiting to be fulfilled. This is the initial state of a
	// request.
	SignRequestStatus_SIGN_REQUEST_STATUS_PENDING SignRequestStatus = 1
	// The request was fulfilled. This is a final state for a request.
	SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED SignRequestStatus = 2
	// The request was rejected. This is a final state for a request.
	SignRequestStatus_SIGN_REQUEST_STATUS_REJECTED SignRequestStatus = 3
)

var SignRequestStatus_name = map[int32]string{
	0: "SIGN_REQUEST_STATUS_UNSPECIFIED",
	1: "SIGN_REQUEST_STATUS_PENDING",
	2: "SIGN_REQUEST_STATUS_FULFILLED",
	3: "SIGN_REQUEST_STATUS_REJECTED",
}

var SignRequestStatus_value = map[string]int32{
	"SIGN_REQUEST_STATUS_UNSPECIFIED": 0,
	"SIGN_REQUEST_STATUS_PENDING":     1,
	"SIGN_REQUEST_STATUS_FULFILLED":   2,
	"SIGN_REQUEST_STATUS_REJECTED":    3,
}

func (x SignRequestStatus) String() string {
	return proto.EnumName(SignRequestStatus_name, int32(x))
}

func (SignRequestStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_950b6dc74b0989bd, []int{0}
}

// BroadcastType specifies how the transaction should be broadcasted.
type BroadcastType int32

const (
	// The signature should be broadcasted manually by the caller.
	BroadcastType_BROADCAST_TYPE_DISABLED BroadcastType = 0
	// The signature should be automatically broadcasted by an offchain relayer.
	BroadcastType_BROADCAST_TYPE_AUTOMATIC BroadcastType = 1
)

var BroadcastType_name = map[int32]string{
	0: "BROADCAST_TYPE_DISABLED",
	1: "BROADCAST_TYPE_AUTOMATIC",
}

var BroadcastType_value = map[string]int32{
	"BROADCAST_TYPE_DISABLED":  0,
	"BROADCAST_TYPE_AUTOMATIC": 1,
}

func (x BroadcastType) String() string {
	return proto.EnumName(BroadcastType_name, int32(x))
}

func (BroadcastType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_950b6dc74b0989bd, []int{1}
}

// SignRequest is the request from a user (creator) to a Keychain to sign a
// message (data_for_signing).
//
// Once that the Keychain has received the request, it will either fulfill it
// or reject it. The result of the request will be stored in the result field.
type SignRequest struct {
	// Unique id of the request.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Address of the creator of the request.
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	// Key ID of the key to be used for signing.
	KeyId uint64 `protobuf:"varint,3,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// Data to be signed.
	DataForSigning []byte `protobuf:"bytes,4,opt,name=data_for_signing,json=dataForSigning,proto3" json:"data_for_signing,omitempty"`
	// Status of the request.
	Status SignRequestStatus `protobuf:"varint,5,opt,name=status,proto3,enum=warden.warden.v1beta3.SignRequestStatus" json:"status,omitempty"`
	// Result of the request, depending on the status:
	//   If pending, this field is empty.
	//   If approved, this field contains the signed data.
	//   If rejected, this field contains the reason.
	//
	// Types that are valid to be assigned to Result:
	//	*SignRequest_SignedData
	//	*SignRequest_RejectReason
	Result        isSignRequest_Result `protobuf_oneof:"result"`
	EncryptionKey []byte               `protobuf:"bytes,8,opt,name=encryption_key,json=encryptionKey,proto3" json:"encryption_key,omitempty"`
	// Amount of fees deducted during new sign request
	DeductedKeychainFees github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,9,rep,name=deducted_keychain_fees,json=deductedKeychainFees,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"deducted_keychain_fees"`
	// Broadcast type of the sign request, indicating how the transaction should be broadcasted.
	BroadcastType BroadcastType `protobuf:"varint,10,opt,name=broadcast_type,json=broadcastType,proto3,enum=warden.warden.v1beta3.BroadcastType" json:"broadcast_type,omitempty"`
}

func (m *SignRequest) Reset()         { *m = SignRequest{} }
func (m *SignRequest) String() string { return proto.CompactTextString(m) }
func (*SignRequest) ProtoMessage()    {}
func (*SignRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_950b6dc74b0989bd, []int{0}
}
func (m *SignRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignRequest.Merge(m, src)
}
func (m *SignRequest) XXX_Size() int {
	return m.Size()
}
func (m *SignRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignRequest proto.InternalMessageInfo

type isSignRequest_Result interface {
	isSignRequest_Result()
	MarshalTo([]byte) (int, error)
	Size() int
}

type SignRequest_SignedData struct {
	SignedData []byte `protobuf:"bytes,6,opt,name=signed_data,json=signedData,proto3,oneof" json:"signed_data,omitempty"`
}
type SignRequest_RejectReason struct {
	RejectReason string `protobuf:"bytes,7,opt,name=reject_reason,json=rejectReason,proto3,oneof" json:"reject_reason,omitempty"`
}

func (*SignRequest_SignedData) isSignRequest_Result()   {}
func (*SignRequest_RejectReason) isSignRequest_Result() {}

func (m *SignRequest) GetResult() isSignRequest_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *SignRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SignRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *SignRequest) GetKeyId() uint64 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func (m *SignRequest) GetDataForSigning() []byte {
	if m != nil {
		return m.DataForSigning
	}
	return nil
}

func (m *SignRequest) GetStatus() SignRequestStatus {
	if m != nil {
		return m.Status
	}
	return SignRequestStatus_SIGN_REQUEST_STATUS_UNSPECIFIED
}

func (m *SignRequest) GetSignedData() []byte {
	if x, ok := m.GetResult().(*SignRequest_SignedData); ok {
		return x.SignedData
	}
	return nil
}

func (m *SignRequest) GetRejectReason() string {
	if x, ok := m.GetResult().(*SignRequest_RejectReason); ok {
		return x.RejectReason
	}
	return ""
}

func (m *SignRequest) GetEncryptionKey() []byte {
	if m != nil {
		return m.EncryptionKey
	}
	return nil
}

func (m *SignRequest) GetDeductedKeychainFees() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.DeductedKeychainFees
	}
	return nil
}

func (m *SignRequest) GetBroadcastType() BroadcastType {
	if m != nil {
		return m.BroadcastType
	}
	return BroadcastType_BROADCAST_TYPE_DISABLED
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SignRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SignRequest_SignedData)(nil),
		(*SignRequest_RejectReason)(nil),
	}
}

func init() {
	proto.RegisterEnum("warden.warden.v1beta3.SignRequestStatus", SignRequestStatus_name, SignRequestStatus_value)
	proto.RegisterEnum("warden.warden.v1beta3.BroadcastType", BroadcastType_name, BroadcastType_value)
	proto.RegisterType((*SignRequest)(nil), "warden.warden.v1beta3.SignRequest")
}

func init() {
	proto.RegisterFile("warden/warden/v1beta3/signature.proto", fileDescriptor_950b6dc74b0989bd)
}

var fileDescriptor_950b6dc74b0989bd = []byte{
	// 673 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x4d, 0x4f, 0xdb, 0x4a,
	0x14, 0x8d, 0x03, 0x04, 0x18, 0x48, 0x5e, 0x18, 0xc1, 0x7b, 0x7e, 0x40, 0x93, 0xf4, 0x03, 0xc9,
	0x8a, 0x84, 0x2d, 0x60, 0xd3, 0x65, 0xe3, 0xc4, 0x01, 0x43, 0x1a, 0xa8, 0xed, 0x2c, 0xda, 0xcd,
	0x68, 0x62, 0x0f, 0xc1, 0x0d, 0x78, 0x52, 0xcf, 0xa4, 0xad, 0xff, 0x45, 0x77, 0x95, 0xba, 0xe8,
	0xba, 0xea, 0x0a, 0xf5, 0x17, 0x74, 0xc9, 0x92, 0x65, 0x57, 0x6d, 0x05, 0x0b, 0xfe, 0x46, 0xe5,
	0xb1, 0xd3, 0xa6, 0x34, 0xdd, 0x64, 0xe6, 0xdc, 0x7b, 0xe6, 0xde, 0x7b, 0xee, 0x51, 0x0c, 0x36,
	0x5e, 0xe1, 0xd0, 0x23, 0x81, 0x96, 0x1e, 0x2f, 0xb7, 0xba, 0x84, 0xe3, 0x1d, 0x8d, 0xf9, 0xbd,
	0x00, 0xf3, 0x61, 0x48, 0xd4, 0x41, 0x48, 0x39, 0x85, 0x2b, 0x49, 0x5e, 0x4d, 0x8f, 0x94, 0xb6,
	0xba, 0x84, 0xcf, 0xfc, 0x80, 0x6a, 0xe2, 0x37, 0x61, 0xae, 0x96, 0x5c, 0xca, 0xce, 0x28, 0xd3,
	0xba, 0x98, 0x91, 0xb4, 0xdc, 0x96, 0xe6, 0x52, 0x3f, 0x48, 0xf3, 0xcb, 0x3d, 0xda, 0xa3, 0xe2,
	0xaa, 0xc5, 0xb7, 0x24, 0x7a, 0xef, 0xf3, 0x34, 0x58, 0xb0, 0xfd, 0x5e, 0x60, 0x91, 0x17, 0x43,
	0xc2, 0x38, 0x2c, 0x80, 0xac, 0xef, 0xc9, 0x52, 0x45, 0x52, 0xa6, 0xad, 0xac, 0xef, 0x41, 0x19,
	0xcc, 0xba, 0x21, 0xc1, 0x9c, 0x86, 0x72, 0xb6, 0x22, 0x29, 0xf3, 0xd6, 0x08, 0xc2, 0x15, 0x90,
	0xeb, 0x93, 0x08, 0xf9, 0x9e, 0x3c, 0x25, 0xd8, 0x33, 0x7d, 0x12, 0x99, 0x1e, 0x54, 0x40, 0xd1,
	0xc3, 0x1c, 0xa3, 0x63, 0x1a, 0xa2, 0x58, 0x8c, 0x1f, 0xf4, 0xe4, 0xe9, 0x8a, 0xa4, 0x2c, 0x5a,
	0x85, 0x38, 0xde, 0xa4, 0xa1, 0x9d, 0x44, 0xe1, 0x23, 0x90, 0x63, 0x1c, 0xf3, 0x21, 0x93, 0x67,
	0x2a, 0x92, 0x52, 0xd8, 0x56, 0xd4, 0x89, 0x5a, 0xd5, 0xb1, 0xf1, 0x6c, 0xc1, 0xb7, 0xd2, 0x77,
	0x70, 0x1b, 0x2c, 0xc4, 0x2d, 0x88, 0x87, 0xe2, 0xd2, 0x72, 0x2e, 0x6e, 0xa3, 0xff, 0xf3, 0xe9,
	0xe6, 0xbc, 0x0a, 0x6c, 0x11, 0x6f, 0x60, 0x8e, 0xf7, 0x32, 0x16, 0x60, 0x3f, 0x11, 0x7c, 0x08,
	0xf2, 0x21, 0x79, 0x4e, 0x5c, 0x8e, 0x42, 0x82, 0x19, 0x0d, 0xe4, 0xd9, 0x58, 0x96, 0xbe, 0x14,
	0xbf, 0x5a, 0xb4, 0x44, 0xc6, 0x12, 0x89, 0xbd, 0x8c, 0xb5, 0x18, 0x8e, 0x61, 0xb8, 0x01, 0x0a,
	0x24, 0x70, 0xc3, 0x68, 0xc0, 0x7d, 0x1a, 0xa0, 0x3e, 0x89, 0xe4, 0x39, 0xa1, 0x2b, 0xff, 0x2b,
	0x7a, 0x40, 0x22, 0xf8, 0x56, 0x02, 0xff, 0x7a, 0xc4, 0x1b, 0xba, 0x9c, 0x78, 0x31, 0xcb, 0x3d,
	0xc1, 0x7e, 0x80, 0x8e, 0x09, 0x61, 0xf2, 0x7c, 0x65, 0x4a, 0x59, 0xd8, 0xfe, 0x5f, 0x4d, 0x9c,
	0x52, 0x63, 0xa7, 0x52, 0x95, 0x5b, 0x6a, 0x9d, 0xfa, 0x81, 0xde, 0xbc, 0xf8, 0x5a, 0xce, 0x7c,
	0xfc, 0x56, 0x56, 0x7a, 0x3e, 0x3f, 0x19, 0x76, 0x55, 0x97, 0x9e, 0x69, 0xa9, 0xad, 0xc9, 0xb1,
	0xc9, 0xbc, 0xbe, 0xc6, 0xa3, 0x01, 0x61, 0xe2, 0x01, 0x7b, 0x17, 0x4f, 0x7d, 0x4a, 0x7a, 0xd8,
	0x8d, 0x50, 0xec, 0x35, 0xfb, 0x70, 0x73, 0x5e, 0x95, 0xac, 0xe5, 0xd1, 0x00, 0x07, 0x69, 0xff,
	0x26, 0x21, 0x0c, 0x1e, 0x80, 0x42, 0x37, 0xa4, 0xd8, 0x73, 0x31, 0xe3, 0x28, 0x2e, 0x22, 0x03,
	0xb1, 0xf8, 0x07, 0x7f, 0x59, 0xbc, 0x3e, 0x22, 0x3b, 0xd1, 0x80, 0x58, 0xf9, 0xee, 0x38, 0xd4,
	0xe7, 0x40, 0x2e, 0x24, 0x6c, 0x78, 0xca, 0xab, 0xef, 0x25, 0xb0, 0xf4, 0x87, 0x47, 0xf0, 0x3e,
	0x28, 0xdb, 0xe6, 0x6e, 0x1b, 0x59, 0xc6, 0x93, 0x8e, 0x61, 0x3b, 0xc8, 0x76, 0x6a, 0x4e, 0xc7,
	0x46, 0x9d, 0xb6, 0x7d, 0x64, 0xd4, 0xcd, 0xa6, 0x69, 0x34, 0x8a, 0x19, 0x58, 0x06, 0x6b, 0x93,
	0x48, 0x47, 0x46, 0xbb, 0x61, 0xb6, 0x77, 0x8b, 0x12, 0xbc, 0x0b, 0xee, 0x4c, 0x22, 0x34, 0x3b,
	0xad, 0xa6, 0xd9, 0x6a, 0x19, 0x8d, 0x62, 0x16, 0x56, 0xc0, 0xfa, 0x24, 0x8a, 0x65, 0xec, 0x1b,
	0x75, 0xc7, 0x68, 0x14, 0xa7, 0xaa, 0xfb, 0x20, 0xff, 0x9b, 0x14, 0xb8, 0x06, 0xfe, 0xd3, 0xad,
	0xc3, 0x5a, 0xa3, 0x5e, 0xb3, 0x1d, 0xe4, 0x3c, 0x3d, 0x32, 0x50, 0xc3, 0xb4, 0x6b, 0x7a, 0x4b,
	0xcc, 0xb4, 0x0e, 0xe4, 0x5b, 0xc9, 0x5a, 0xc7, 0x39, 0x7c, 0x5c, 0x73, 0xcc, 0x7a, 0x51, 0xd2,
	0xf1, 0xc5, 0x55, 0x49, 0xba, 0xbc, 0x2a, 0x49, 0xdf, 0xaf, 0x4a, 0xd2, 0x9b, 0xeb, 0x52, 0xe6,
	0xf2, 0xba, 0x94, 0xf9, 0x72, 0x5d, 0xca, 0x3c, 0xdb, 0x1d, 0xf3, 0x2c, 0x59, 0xe4, 0xa6, 0xf8,
	0x8b, 0xb9, 0xf4, 0x34, 0xc5, 0xb7, 0xa0, 0xf6, 0x7a, 0x74, 0x11, 0x86, 0x8e, 0x3e, 0x01, 0xdd,
	0x9c, 0xe0, 0xed, 0xfc, 0x08, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x41, 0xc6, 0x2d, 0x22, 0x04, 0x00,
	0x00,
}

func (m *SignRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BroadcastType != 0 {
		i = encodeVarintSignature(dAtA, i, uint64(m.BroadcastType))
		i--
		dAtA[i] = 0x50
	}
	if len(m.DeductedKeychainFees) > 0 {
		for iNdEx := len(m.DeductedKeychainFees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DeductedKeychainFees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSignature(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.EncryptionKey) > 0 {
		i -= len(m.EncryptionKey)
		copy(dAtA[i:], m.EncryptionKey)
		i = encodeVarintSignature(dAtA, i, uint64(len(m.EncryptionKey)))
		i--
		dAtA[i] = 0x42
	}
	if m.Result != nil {
		{
			size := m.Result.Size()
			i -= size
			if _, err := m.Result.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.Status != 0 {
		i = encodeVarintSignature(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if len(m.DataForSigning) > 0 {
		i -= len(m.DataForSigning)
		copy(dAtA[i:], m.DataForSigning)
		i = encodeVarintSignature(dAtA, i, uint64(len(m.DataForSigning)))
		i--
		dAtA[i] = 0x22
	}
	if m.KeyId != 0 {
		i = encodeVarintSignature(dAtA, i, uint64(m.KeyId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSignature(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintSignature(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SignRequest_SignedData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignRequest_SignedData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SignedData != nil {
		i -= len(m.SignedData)
		copy(dAtA[i:], m.SignedData)
		i = encodeVarintSignature(dAtA, i, uint64(len(m.SignedData)))
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func (m *SignRequest_RejectReason) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignRequest_RejectReason) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.RejectReason)
	copy(dAtA[i:], m.RejectReason)
	i = encodeVarintSignature(dAtA, i, uint64(len(m.RejectReason)))
	i--
	dAtA[i] = 0x3a
	return len(dAtA) - i, nil
}
func encodeVarintSignature(dAtA []byte, offset int, v uint64) int {
	offset -= sovSignature(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SignRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovSignature(uint64(m.Id))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSignature(uint64(l))
	}
	if m.KeyId != 0 {
		n += 1 + sovSignature(uint64(m.KeyId))
	}
	l = len(m.DataForSigning)
	if l > 0 {
		n += 1 + l + sovSignature(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovSignature(uint64(m.Status))
	}
	if m.Result != nil {
		n += m.Result.Size()
	}
	l = len(m.EncryptionKey)
	if l > 0 {
		n += 1 + l + sovSignature(uint64(l))
	}
	if len(m.DeductedKeychainFees) > 0 {
		for _, e := range m.DeductedKeychainFees {
			l = e.Size()
			n += 1 + l + sovSignature(uint64(l))
		}
	}
	if m.BroadcastType != 0 {
		n += 1 + sovSignature(uint64(m.BroadcastType))
	}
	return n
}

func (m *SignRequest_SignedData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SignedData != nil {
		l = len(m.SignedData)
		n += 1 + l + sovSignature(uint64(l))
	}
	return n
}
func (m *SignRequest_RejectReason) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RejectReason)
	n += 1 + l + sovSignature(uint64(l))
	return n
}

func sovSignature(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSignature(x uint64) (n int) {
	return sovSignature(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SignRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSignature
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SignRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSignature
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSignature
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyId", wireType)
			}
			m.KeyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataForSigning", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSignature
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSignature
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataForSigning = append(m.DataForSigning[:0], dAtA[iNdEx:postIndex]...)
			if m.DataForSigning == nil {
				m.DataForSigning = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= SignRequestStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSignature
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSignature
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.Result = &SignRequest_SignedData{v}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RejectReason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSignature
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSignature
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = &SignRequest_RejectReason{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptionKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSignature
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSignature
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncryptionKey = append(m.EncryptionKey[:0], dAtA[iNdEx:postIndex]...)
			if m.EncryptionKey == nil {
				m.EncryptionKey = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeductedKeychainFees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSignature
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSignature
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeductedKeychainFees = append(m.DeductedKeychainFees, types.Coin{})
			if err := m.DeductedKeychainFees[len(m.DeductedKeychainFees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BroadcastType", wireType)
			}
			m.BroadcastType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BroadcastType |= BroadcastType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSignature(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSignature
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSignature(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSignature
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSignature
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSignature
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSignature
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSignature        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSignature          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSignature = fmt.Errorf("proto: unexpected end of group")
)
