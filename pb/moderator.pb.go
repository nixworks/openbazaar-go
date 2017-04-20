// Code generated by protoc-gen-go.
// source: moderator.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Moderator_Fee_FeeType int32

const (
	Moderator_Fee_FIXED                 Moderator_Fee_FeeType = 0
	Moderator_Fee_PERCENTAGE            Moderator_Fee_FeeType = 1
	Moderator_Fee_FIXED_PLUS_PERCENTAGE Moderator_Fee_FeeType = 2
)

var Moderator_Fee_FeeType_name = map[int32]string{
	0: "FIXED",
	1: "PERCENTAGE",
	2: "FIXED_PLUS_PERCENTAGE",
}
var Moderator_Fee_FeeType_value = map[string]int32{
	"FIXED":                 0,
	"PERCENTAGE":            1,
	"FIXED_PLUS_PERCENTAGE": 2,
}

func (x Moderator_Fee_FeeType) String() string {
	return proto.EnumName(Moderator_Fee_FeeType_name, int32(x))
}
func (Moderator_Fee_FeeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 0, 0} }

type Moderator struct {
	Description        string         `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	TermsAndConditions string         `protobuf:"bytes,2,opt,name=termsAndConditions" json:"termsAndConditions,omitempty"`
	Languages          []string       `protobuf:"bytes,3,rep,name=languages" json:"languages,omitempty"`
	AcceptedCurrency   string         `protobuf:"bytes,4,opt,name=acceptedCurrency" json:"acceptedCurrency,omitempty"`
	Fee                *Moderator_Fee `protobuf:"bytes,5,opt,name=fee" json:"fee,omitempty"`
}

func (m *Moderator) Reset()                    { *m = Moderator{} }
func (m *Moderator) String() string            { return proto.CompactTextString(m) }
func (*Moderator) ProtoMessage()               {}
func (*Moderator) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Moderator) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Moderator) GetTermsAndConditions() string {
	if m != nil {
		return m.TermsAndConditions
	}
	return ""
}

func (m *Moderator) GetLanguages() []string {
	if m != nil {
		return m.Languages
	}
	return nil
}

func (m *Moderator) GetAcceptedCurrency() string {
	if m != nil {
		return m.AcceptedCurrency
	}
	return ""
}

func (m *Moderator) GetFee() *Moderator_Fee {
	if m != nil {
		return m.Fee
	}
	return nil
}

type Moderator_Fee struct {
	FixedFee   *Moderator_Price      `protobuf:"bytes,1,opt,name=fixedFee" json:"fixedFee,omitempty"`
	Percentage float32               `protobuf:"fixed32,2,opt,name=percentage" json:"percentage,omitempty"`
	FeeType    Moderator_Fee_FeeType `protobuf:"varint,3,opt,name=feeType,enum=Moderator_Fee_FeeType" json:"feeType,omitempty"`
}

func (m *Moderator_Fee) Reset()                    { *m = Moderator_Fee{} }
func (m *Moderator_Fee) String() string            { return proto.CompactTextString(m) }
func (*Moderator_Fee) ProtoMessage()               {}
func (*Moderator_Fee) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 0} }

func (m *Moderator_Fee) GetFixedFee() *Moderator_Price {
	if m != nil {
		return m.FixedFee
	}
	return nil
}

func (m *Moderator_Fee) GetPercentage() float32 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

func (m *Moderator_Fee) GetFeeType() Moderator_Fee_FeeType {
	if m != nil {
		return m.FeeType
	}
	return Moderator_Fee_FIXED
}

type Moderator_Price struct {
	CurrencyCode string `protobuf:"bytes,1,opt,name=currencyCode" json:"currencyCode,omitempty"`
	Amount       uint64 `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *Moderator_Price) Reset()                    { *m = Moderator_Price{} }
func (m *Moderator_Price) String() string            { return proto.CompactTextString(m) }
func (*Moderator_Price) ProtoMessage()               {}
func (*Moderator_Price) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 1} }

func (m *Moderator_Price) GetCurrencyCode() string {
	if m != nil {
		return m.CurrencyCode
	}
	return ""
}

func (m *Moderator_Price) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type DisputeUpdate struct {
	OrderId            string      `protobuf:"bytes,1,opt,name=orderId" json:"orderId,omitempty"`
	PayoutAddress      string      `protobuf:"bytes,2,opt,name=payoutAddress" json:"payoutAddress,omitempty"`
	Outpoints          []*Outpoint `protobuf:"bytes,3,rep,name=outpoints" json:"outpoints,omitempty"`
	SerializedContract []byte      `protobuf:"bytes,4,opt,name=serializedContract,proto3" json:"serializedContract,omitempty"`
}

func (m *DisputeUpdate) Reset()                    { *m = DisputeUpdate{} }
func (m *DisputeUpdate) String() string            { return proto.CompactTextString(m) }
func (*DisputeUpdate) ProtoMessage()               {}
func (*DisputeUpdate) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *DisputeUpdate) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *DisputeUpdate) GetPayoutAddress() string {
	if m != nil {
		return m.PayoutAddress
	}
	return ""
}

func (m *DisputeUpdate) GetOutpoints() []*Outpoint {
	if m != nil {
		return m.Outpoints
	}
	return nil
}

func (m *DisputeUpdate) GetSerializedContract() []byte {
	if m != nil {
		return m.SerializedContract
	}
	return nil
}

func init() {
	proto.RegisterType((*Moderator)(nil), "Moderator")
	proto.RegisterType((*Moderator_Fee)(nil), "Moderator.Fee")
	proto.RegisterType((*Moderator_Price)(nil), "Moderator.Price")
	proto.RegisterType((*DisputeUpdate)(nil), "DisputeUpdate")
	proto.RegisterEnum("Moderator_Fee_FeeType", Moderator_Fee_FeeType_name, Moderator_Fee_FeeType_value)
}

func init() { proto.RegisterFile("moderator.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0xdf, 0x6a, 0xdb, 0x30,
	0x14, 0xc6, 0xe7, 0x38, 0x6d, 0xe6, 0x93, 0x36, 0x0d, 0x07, 0x56, 0xbc, 0x30, 0x86, 0x09, 0x83,
	0xe5, 0x62, 0x98, 0x91, 0x3d, 0xc0, 0xc8, 0xd2, 0x64, 0x14, 0xf6, 0x27, 0x68, 0x2d, 0x8c, 0xdd,
	0x14, 0xd5, 0x3a, 0x09, 0x82, 0x46, 0x12, 0x92, 0x0c, 0xcb, 0x9e, 0x68, 0xb0, 0xf7, 0xd8, 0x73,
	0x15, 0xab, 0x76, 0xeb, 0x40, 0x2e, 0x74, 0x71, 0x7e, 0xdf, 0x87, 0xf4, 0xa1, 0xef, 0xc0, 0xd9,
	0x56, 0x0b, 0xb2, 0xdc, 0x6b, 0x9b, 0x1b, 0xab, 0xbd, 0x1e, 0x9d, 0x15, 0x5a, 0x79, 0xcb, 0x0b,
	0xef, 0x1e, 0xc0, 0xf8, 0x5f, 0x0c, 0xc9, 0xd7, 0xc6, 0x84, 0x19, 0xf4, 0x05, 0xb9, 0xc2, 0x4a,
	0xe3, 0xa5, 0x56, 0x69, 0x94, 0x45, 0x93, 0x84, 0xb5, 0x11, 0xe6, 0x80, 0x9e, 0xec, 0xd6, 0xcd,
	0x94, 0x98, 0x6b, 0x25, 0x64, 0x05, 0x5d, 0xda, 0x09, 0xc6, 0x03, 0x0a, 0xbe, 0x82, 0xe4, 0x8e,
	0xab, 0x4d, 0xc9, 0x37, 0xe4, 0xd2, 0x38, 0x8b, 0x27, 0x09, 0x7b, 0x02, 0x98, 0x41, 0xbc, 0x26,
	0x4a, 0xbb, 0x59, 0x34, 0xe9, 0x4f, 0x07, 0xf9, 0x63, 0x90, 0x7c, 0x49, 0xc4, 0x2a, 0x69, 0xf4,
	0x3f, 0x82, 0x78, 0x49, 0x84, 0xef, 0xe0, 0xf9, 0x5a, 0xfe, 0x26, 0xb1, 0x24, 0x0a, 0xb1, 0xfa,
	0xd3, 0x61, 0xcb, 0xbe, 0xb2, 0xb2, 0x20, 0xf6, 0xe8, 0xc0, 0xd7, 0x00, 0x86, 0x6c, 0x41, 0xca,
	0xf3, 0x0d, 0x85, 0x74, 0x1d, 0xd6, 0x22, 0xf8, 0x1e, 0x7a, 0x6b, 0xa2, 0xab, 0x9d, 0xa1, 0x34,
	0xce, 0xa2, 0xc9, 0x60, 0x7a, 0xbe, 0xff, 0x76, 0x75, 0x2a, 0x95, 0x35, 0xb6, 0xf1, 0x47, 0xe8,
	0xd5, 0x0c, 0x13, 0x38, 0x5a, 0x5e, 0xfe, 0x5c, 0x5c, 0x0c, 0x9f, 0xe1, 0x00, 0x60, 0xb5, 0x60,
	0xf3, 0xc5, 0xb7, 0xab, 0xd9, 0xe7, 0xc5, 0x30, 0xc2, 0x97, 0xf0, 0x22, 0x48, 0x37, 0xab, 0x2f,
	0xd7, 0x3f, 0x6e, 0x5a, 0x52, 0x67, 0x34, 0x87, 0xa3, 0x90, 0x12, 0xc7, 0x70, 0x52, 0x94, 0xd6,
	0x92, 0x2a, 0x76, 0x73, 0x2d, 0xa8, 0xfe, 0xe4, 0x3d, 0x86, 0xe7, 0x70, 0xcc, 0xb7, 0xba, 0x54,
	0x3e, 0x64, 0xef, 0xb2, 0x7a, 0x1a, 0xff, 0x8d, 0xe0, 0xf4, 0x42, 0x3a, 0x53, 0x7a, 0xba, 0x36,
	0x82, 0x7b, 0xc2, 0x14, 0x7a, 0xda, 0x0a, 0xb2, 0x97, 0xa2, 0xbe, 0xa8, 0x19, 0xf1, 0x0d, 0x9c,
	0x1a, 0xbe, 0xd3, 0xa5, 0x9f, 0x09, 0x61, 0xc9, 0x35, 0x25, 0xed, 0x43, 0x7c, 0x0b, 0x89, 0x2e,
	0xbd, 0xd1, 0x52, 0xf9, 0x87, 0x7e, 0xfa, 0xd3, 0x24, 0xff, 0x5e, 0x13, 0xf6, 0xa4, 0x55, 0xc5,
	0x3b, 0xb2, 0x92, 0xdf, 0xc9, 0x3f, 0x54, 0x15, 0x1c, 0xb6, 0x28, 0x34, 0x77, 0xc2, 0x0e, 0x28,
	0x9f, 0xba, 0xbf, 0x3a, 0xe6, 0xf6, 0xf6, 0x38, 0x6c, 0xd9, 0x87, 0xfb, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xd6, 0x38, 0x96, 0x16, 0x89, 0x02, 0x00, 0x00,
}
