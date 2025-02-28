// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: agoric/swingset/swingset.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// CoreEvalProposal is a gov Content type for evaluating code in the SwingSet
// core.
// See `agoric-sdk/packages/vats/src/core/eval.js`.
type CoreEvalProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Although evals are sequential, they may run concurrently, since they each
	// can return a Promise.
	Evals []CoreEval `protobuf:"bytes,3,rep,name=evals,proto3" json:"evals"`
}

func (m *CoreEvalProposal) Reset()         { *m = CoreEvalProposal{} }
func (m *CoreEvalProposal) String() string { return proto.CompactTextString(m) }
func (*CoreEvalProposal) ProtoMessage()    {}
func (*CoreEvalProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff9c341e0de15f8b, []int{0}
}
func (m *CoreEvalProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CoreEvalProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CoreEvalProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CoreEvalProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoreEvalProposal.Merge(m, src)
}
func (m *CoreEvalProposal) XXX_Size() int {
	return m.Size()
}
func (m *CoreEvalProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CoreEvalProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CoreEvalProposal proto.InternalMessageInfo

// CoreEval defines an individual SwingSet core evaluation, for use in
// CoreEvalProposal.
type CoreEval struct {
	// Grant these JSON-stringified core bootstrap permits to the jsCode, as the
	// `powers` endowment.
	JsonPermits string `protobuf:"bytes,1,opt,name=json_permits,json=jsonPermits,proto3" json:"json_permits,omitempty" yaml:"json_permits"`
	// Evaluate this JavaScript code in a Compartment endowed with `powers` as
	// well as some powerless helpers.
	JsCode string `protobuf:"bytes,2,opt,name=js_code,json=jsCode,proto3" json:"js_code,omitempty" yaml:"js_code"`
}

func (m *CoreEval) Reset()         { *m = CoreEval{} }
func (m *CoreEval) String() string { return proto.CompactTextString(m) }
func (*CoreEval) ProtoMessage()    {}
func (*CoreEval) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff9c341e0de15f8b, []int{1}
}
func (m *CoreEval) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CoreEval) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CoreEval.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CoreEval) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoreEval.Merge(m, src)
}
func (m *CoreEval) XXX_Size() int {
	return m.Size()
}
func (m *CoreEval) XXX_DiscardUnknown() {
	xxx_messageInfo_CoreEval.DiscardUnknown(m)
}

var xxx_messageInfo_CoreEval proto.InternalMessageInfo

func (m *CoreEval) GetJsonPermits() string {
	if m != nil {
		return m.JsonPermits
	}
	return ""
}

func (m *CoreEval) GetJsCode() string {
	if m != nil {
		return m.JsCode
	}
	return ""
}

// Params are the swingset configuration/governance parameters.
type Params struct {
	// Map from unit name to a value in SwingSet "beans".
	// Must not be negative.
	//
	// These values are used by SwingSet to normalize named per-resource charges
	// (maybe rent) in a single Nat usage unit, the "bean".
	//
	// There is no required order to this list of entries, but all the chain
	// nodes must all serialize and deserialize the existing order without
	// permuting it.
	BeansPerUnit []StringBeans `protobuf:"bytes,1,rep,name=beans_per_unit,json=beansPerUnit,proto3" json:"beans_per_unit"`
	// The price in Coins per the unit named "fee".  This value is used by
	// cosmic-swingset JS code to decide how many tokens to charge.
	//
	// cost = beans_used * fee_unit_price / beans_per_unit["fee"]
	FeeUnitPrice github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=fee_unit_price,json=feeUnitPrice,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fee_unit_price"`
	// The SwingSet bootstrap vat configuration file.  Not usefully modifiable
	// via governance as it is only referenced by the chain's initial
	// construction.
	BootstrapVatConfig string `protobuf:"bytes,3,opt,name=bootstrap_vat_config,json=bootstrapVatConfig,proto3" json:"bootstrap_vat_config,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff9c341e0de15f8b, []int{2}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetBeansPerUnit() []StringBeans {
	if m != nil {
		return m.BeansPerUnit
	}
	return nil
}

func (m *Params) GetFeeUnitPrice() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.FeeUnitPrice
	}
	return nil
}

func (m *Params) GetBootstrapVatConfig() string {
	if m != nil {
		return m.BootstrapVatConfig
	}
	return ""
}

// Map element of a string key to a Nat bean count.
type StringBeans struct {
	// What the beans are for.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The actual bean value.
	Beans github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,2,opt,name=beans,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"beans"`
}

func (m *StringBeans) Reset()         { *m = StringBeans{} }
func (m *StringBeans) String() string { return proto.CompactTextString(m) }
func (*StringBeans) ProtoMessage()    {}
func (*StringBeans) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff9c341e0de15f8b, []int{3}
}
func (m *StringBeans) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StringBeans) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StringBeans.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StringBeans) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringBeans.Merge(m, src)
}
func (m *StringBeans) XXX_Size() int {
	return m.Size()
}
func (m *StringBeans) XXX_DiscardUnknown() {
	xxx_messageInfo_StringBeans.DiscardUnknown(m)
}

var xxx_messageInfo_StringBeans proto.InternalMessageInfo

func (m *StringBeans) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func init() {
	proto.RegisterType((*CoreEvalProposal)(nil), "agoric.swingset.CoreEvalProposal")
	proto.RegisterType((*CoreEval)(nil), "agoric.swingset.CoreEval")
	proto.RegisterType((*Params)(nil), "agoric.swingset.Params")
	proto.RegisterType((*StringBeans)(nil), "agoric.swingset.StringBeans")
}

func init() { proto.RegisterFile("agoric/swingset/swingset.proto", fileDescriptor_ff9c341e0de15f8b) }

var fileDescriptor_ff9c341e0de15f8b = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xbf, 0x8f, 0xd3, 0x3e,
	0x14, 0x4f, 0xae, 0x3f, 0xbe, 0xf7, 0x75, 0xab, 0x72, 0x32, 0x95, 0x28, 0x27, 0x94, 0x54, 0x59,
	0xa8, 0x84, 0x48, 0xee, 0x40, 0x2c, 0x65, 0x22, 0xd5, 0x49, 0x8c, 0x55, 0x50, 0x19, 0x58, 0x22,
	0x27, 0x75, 0x83, 0x7b, 0xa9, 0x1d, 0x6c, 0x5f, 0xa1, 0x23, 0x13, 0x8c, 0x8c, 0x8c, 0x9d, 0xf9,
	0x4b, 0x6e, 0xbc, 0x11, 0x31, 0x14, 0xd4, 0x2e, 0xcc, 0xf7, 0x17, 0x20, 0xdb, 0xc9, 0x51, 0xc1,
	0xc2, 0x94, 0x67, 0x7f, 0xde, 0x7b, 0x9f, 0xcf, 0x7b, 0x1f, 0x07, 0x38, 0x28, 0x63, 0x9c, 0xa4,
	0x81, 0x78, 0x4b, 0x68, 0x26, 0xb0, 0xbc, 0x09, 0xfc, 0x82, 0x33, 0xc9, 0xe0, 0x2d, 0x83, 0xfb,
	0xd5, 0xf5, 0x71, 0x37, 0x63, 0x19, 0xd3, 0x58, 0xa0, 0x22, 0x93, 0x76, 0xec, 0xa4, 0x4c, 0x2c,
	0x98, 0x08, 0x12, 0x24, 0x70, 0xb0, 0x3c, 0x4d, 0xb0, 0x44, 0xa7, 0x41, 0xca, 0x08, 0x35, 0xb8,
	0xf7, 0xc1, 0x06, 0x47, 0x23, 0xc6, 0xf1, 0xd9, 0x12, 0xe5, 0x63, 0xce, 0x0a, 0x26, 0x50, 0x0e,
	0xbb, 0xa0, 0x21, 0x89, 0xcc, 0x71, 0xcf, 0xee, 0xdb, 0x83, 0xff, 0x23, 0x73, 0x80, 0x7d, 0xd0,
	0x9a, 0x62, 0x91, 0x72, 0x52, 0x48, 0xc2, 0x68, 0xef, 0x40, 0x63, 0xfb, 0x57, 0xf0, 0x09, 0x68,
	0xe0, 0x25, 0xca, 0x45, 0xaf, 0xd6, 0xaf, 0x0d, 0x5a, 0x8f, 0xee, 0xfa, 0x7f, 0x68, 0xf4, 0x2b,
	0xa6, 0xb0, 0x7e, 0xb9, 0x71, 0xad, 0xc8, 0x64, 0x0f, 0xeb, 0x1f, 0xd7, 0xae, 0xe5, 0x09, 0x70,
	0x58, 0xc1, 0x70, 0x08, 0xda, 0x73, 0xc1, 0x68, 0x5c, 0x60, 0xbe, 0x20, 0x52, 0x18, 0x1d, 0xe1,
	0x9d, 0xeb, 0x8d, 0x7b, 0x7b, 0x85, 0x16, 0xf9, 0xd0, 0xdb, 0x47, 0xbd, 0xa8, 0xa5, 0x8e, 0x63,
	0x73, 0x82, 0x0f, 0xc0, 0x7f, 0x73, 0x11, 0xa7, 0x6c, 0x8a, 0x8d, 0xc4, 0x10, 0x5e, 0x6f, 0xdc,
	0x4e, 0x55, 0xa6, 0x01, 0x2f, 0x6a, 0xce, 0xc5, 0x48, 0x05, 0xef, 0x0f, 0x40, 0x73, 0x8c, 0x38,
	0x5a, 0x08, 0xf8, 0x1c, 0x74, 0x12, 0x8c, 0xa8, 0x50, 0x6d, 0xe3, 0x0b, 0x4a, 0x64, 0xcf, 0xd6,
	0x53, 0xdc, 0xfb, 0x6b, 0x8a, 0x17, 0x92, 0x13, 0x9a, 0x85, 0x2a, 0xb9, 0x1c, 0xa4, 0xad, 0x2b,
	0xc7, 0x98, 0x4f, 0x28, 0x91, 0xf0, 0x0d, 0xe8, 0xcc, 0x30, 0xd6, 0x3d, 0xe2, 0x82, 0x93, 0x54,
	0x09, 0x31, 0xfb, 0x30, 0x66, 0xf8, 0xca, 0x0c, 0xbf, 0x34, 0xc3, 0x1f, 0x31, 0x42, 0xc3, 0x13,
	0xd5, 0xe6, 0xcb, 0x77, 0x77, 0x90, 0x11, 0xf9, 0xfa, 0x22, 0xf1, 0x53, 0xb6, 0x08, 0x4a, 0xe7,
	0xcc, 0xe7, 0xa1, 0x98, 0x9e, 0x07, 0x72, 0x55, 0x60, 0xa1, 0x0b, 0x44, 0xd4, 0x9e, 0x61, 0xac,
	0xd8, 0xc6, 0x8a, 0x00, 0x9e, 0x80, 0x6e, 0xc2, 0x98, 0x14, 0x92, 0xa3, 0x22, 0x5e, 0x22, 0x19,
	0xa7, 0x8c, 0xce, 0x48, 0xd6, 0xab, 0x69, 0x93, 0xe0, 0x0d, 0xf6, 0x12, 0xc9, 0x91, 0x46, 0x86,
	0x87, 0x9f, 0xd7, 0xae, 0xf5, 0x73, 0xed, 0xda, 0x5e, 0x0e, 0x5a, 0x7b, 0x13, 0xc1, 0x23, 0x50,
	0x3b, 0xc7, 0xab, 0xd2, 0x7a, 0x15, 0xc2, 0x33, 0xd0, 0xd0, 0xf3, 0x95, 0xfb, 0x0c, 0x94, 0xd6,
	0x6f, 0x1b, 0xf7, 0xfe, 0x3f, 0x68, 0x9d, 0x10, 0x2a, 0x23, 0x53, 0x3d, 0xac, 0x2b, 0xb6, 0x70,
	0x72, 0xb9, 0x75, 0xec, 0xab, 0xad, 0x63, 0xff, 0xd8, 0x3a, 0xf6, 0xa7, 0x9d, 0x63, 0x5d, 0xed,
	0x1c, 0xeb, 0xeb, 0xce, 0xb1, 0x5e, 0x3d, 0xdd, 0xeb, 0xf7, 0xcc, 0x3c, 0x7e, 0xb3, 0x79, 0xdd,
	0x2f, 0x63, 0x39, 0xa2, 0x59, 0x45, 0xf4, 0xee, 0xf7, 0x7f, 0xa1, 0x89, 0x92, 0xa6, 0x7e, 0xce,
	0x8f, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x6e, 0x33, 0x65, 0x37, 0x03, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.BeansPerUnit) != len(that1.BeansPerUnit) {
		return false
	}
	for i := range this.BeansPerUnit {
		if !this.BeansPerUnit[i].Equal(&that1.BeansPerUnit[i]) {
			return false
		}
	}
	if len(this.FeeUnitPrice) != len(that1.FeeUnitPrice) {
		return false
	}
	for i := range this.FeeUnitPrice {
		if !this.FeeUnitPrice[i].Equal(&that1.FeeUnitPrice[i]) {
			return false
		}
	}
	if this.BootstrapVatConfig != that1.BootstrapVatConfig {
		return false
	}
	return true
}
func (this *StringBeans) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StringBeans)
	if !ok {
		that2, ok := that.(StringBeans)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if !this.Beans.Equal(that1.Beans) {
		return false
	}
	return true
}
func (m *CoreEvalProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoreEvalProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CoreEvalProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Evals) > 0 {
		for iNdEx := len(m.Evals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Evals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSwingset(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintSwingset(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintSwingset(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CoreEval) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoreEval) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CoreEval) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.JsCode) > 0 {
		i -= len(m.JsCode)
		copy(dAtA[i:], m.JsCode)
		i = encodeVarintSwingset(dAtA, i, uint64(len(m.JsCode)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.JsonPermits) > 0 {
		i -= len(m.JsonPermits)
		copy(dAtA[i:], m.JsonPermits)
		i = encodeVarintSwingset(dAtA, i, uint64(len(m.JsonPermits)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BootstrapVatConfig) > 0 {
		i -= len(m.BootstrapVatConfig)
		copy(dAtA[i:], m.BootstrapVatConfig)
		i = encodeVarintSwingset(dAtA, i, uint64(len(m.BootstrapVatConfig)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.FeeUnitPrice) > 0 {
		for iNdEx := len(m.FeeUnitPrice) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeUnitPrice[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSwingset(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.BeansPerUnit) > 0 {
		for iNdEx := len(m.BeansPerUnit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BeansPerUnit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSwingset(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *StringBeans) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StringBeans) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StringBeans) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Beans.Size()
		i -= size
		if _, err := m.Beans.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSwingset(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintSwingset(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSwingset(dAtA []byte, offset int, v uint64) int {
	offset -= sovSwingset(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CoreEvalProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovSwingset(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovSwingset(uint64(l))
	}
	if len(m.Evals) > 0 {
		for _, e := range m.Evals {
			l = e.Size()
			n += 1 + l + sovSwingset(uint64(l))
		}
	}
	return n
}

func (m *CoreEval) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.JsonPermits)
	if l > 0 {
		n += 1 + l + sovSwingset(uint64(l))
	}
	l = len(m.JsCode)
	if l > 0 {
		n += 1 + l + sovSwingset(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.BeansPerUnit) > 0 {
		for _, e := range m.BeansPerUnit {
			l = e.Size()
			n += 1 + l + sovSwingset(uint64(l))
		}
	}
	if len(m.FeeUnitPrice) > 0 {
		for _, e := range m.FeeUnitPrice {
			l = e.Size()
			n += 1 + l + sovSwingset(uint64(l))
		}
	}
	l = len(m.BootstrapVatConfig)
	if l > 0 {
		n += 1 + l + sovSwingset(uint64(l))
	}
	return n
}

func (m *StringBeans) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovSwingset(uint64(l))
	}
	l = m.Beans.Size()
	n += 1 + l + sovSwingset(uint64(l))
	return n
}

func sovSwingset(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSwingset(x uint64) (n int) {
	return sovSwingset(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CoreEvalProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSwingset
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
			return fmt.Errorf("proto: CoreEvalProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoreEvalProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwingset
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
				return ErrInvalidLengthSwingset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwingset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwingset
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
				return ErrInvalidLengthSwingset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwingset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Evals", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwingset
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
				return ErrInvalidLengthSwingset
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSwingset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Evals = append(m.Evals, CoreEval{})
			if err := m.Evals[len(m.Evals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSwingset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSwingset
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
func (m *CoreEval) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSwingset
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
			return fmt.Errorf("proto: CoreEval: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoreEval: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JsonPermits", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwingset
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
				return ErrInvalidLengthSwingset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwingset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JsonPermits = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JsCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwingset
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
				return ErrInvalidLengthSwingset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwingset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JsCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSwingset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSwingset
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSwingset
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeansPerUnit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwingset
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
				return ErrInvalidLengthSwingset
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSwingset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BeansPerUnit = append(m.BeansPerUnit, StringBeans{})
			if err := m.BeansPerUnit[len(m.BeansPerUnit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeUnitPrice", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwingset
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
				return ErrInvalidLengthSwingset
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSwingset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeUnitPrice = append(m.FeeUnitPrice, types.Coin{})
			if err := m.FeeUnitPrice[len(m.FeeUnitPrice)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BootstrapVatConfig", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwingset
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
				return ErrInvalidLengthSwingset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwingset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BootstrapVatConfig = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSwingset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSwingset
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
func (m *StringBeans) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSwingset
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
			return fmt.Errorf("proto: StringBeans: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StringBeans: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwingset
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
				return ErrInvalidLengthSwingset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwingset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Beans", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwingset
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
				return ErrInvalidLengthSwingset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwingset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Beans.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSwingset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSwingset
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
func skipSwingset(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSwingset
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
					return 0, ErrIntOverflowSwingset
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
					return 0, ErrIntOverflowSwingset
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
				return 0, ErrInvalidLengthSwingset
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSwingset
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSwingset
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSwingset        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSwingset          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSwingset = fmt.Errorf("proto: unexpected end of group")
)
