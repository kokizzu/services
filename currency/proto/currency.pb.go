// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: proto/currency.proto

package currency

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

type Code struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// e.g USD
	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// e.g United States Dollar
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Code) Reset() {
	*x = Code{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Code) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Code) ProtoMessage() {}

func (x *Code) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Code.ProtoReflect.Descriptor instead.
func (*Code) Descriptor() ([]byte, []int) {
	return file_proto_currency_proto_rawDescGZIP(), []int{0}
}

func (x *Code) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Code) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Codes returns the supported currency codes for the API
type CodesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CodesRequest) Reset() {
	*x = CodesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodesRequest) ProtoMessage() {}

func (x *CodesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodesRequest.ProtoReflect.Descriptor instead.
func (*CodesRequest) Descriptor() ([]byte, []int) {
	return file_proto_currency_proto_rawDescGZIP(), []int{1}
}

type CodesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Codes []*Code `protobuf:"bytes,1,rep,name=codes,proto3" json:"codes,omitempty"`
}

func (x *CodesResponse) Reset() {
	*x = CodesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currency_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodesResponse) ProtoMessage() {}

func (x *CodesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currency_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodesResponse.ProtoReflect.Descriptor instead.
func (*CodesResponse) Descriptor() ([]byte, []int) {
	return file_proto_currency_proto_rawDescGZIP(), []int{2}
}

func (x *CodesResponse) GetCodes() []*Code {
	if x != nil {
		return x.Codes
	}
	return nil
}

// Rates returns the currency rates for a given code e.g USD
type RatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The currency code to get rates for e.g USD
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *RatesRequest) Reset() {
	*x = RatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currency_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatesRequest) ProtoMessage() {}

func (x *RatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currency_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatesRequest.ProtoReflect.Descriptor instead.
func (*RatesRequest) Descriptor() ([]byte, []int) {
	return file_proto_currency_proto_rawDescGZIP(), []int{3}
}

func (x *RatesRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type RatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The code requested e.g USD
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	// The rates for the given code as key-value pairs code:rate
	Rates map[string]float64 `protobuf:"bytes,2,rep,name=rates,proto3" json:"rates,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *RatesResponse) Reset() {
	*x = RatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currency_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatesResponse) ProtoMessage() {}

func (x *RatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currency_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatesResponse.ProtoReflect.Descriptor instead.
func (*RatesResponse) Descriptor() ([]byte, []int) {
	return file_proto_currency_proto_rawDescGZIP(), []int{4}
}

func (x *RatesResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *RatesResponse) GetRates() map[string]float64 {
	if x != nil {
		return x.Rates
	}
	return nil
}

// Convert returns the currency conversion rate between two pairs e.g USD/GBP
type ConvertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// base code to convert from e.g USD
	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	// target code to convert to e.g GBP
	To string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	// optional amount to convert e.g 10.0
	Amount float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *ConvertRequest) Reset() {
	*x = ConvertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currency_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertRequest) ProtoMessage() {}

func (x *ConvertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currency_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertRequest.ProtoReflect.Descriptor instead.
func (*ConvertRequest) Descriptor() ([]byte, []int) {
	return file_proto_currency_proto_rawDescGZIP(), []int{5}
}

func (x *ConvertRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *ConvertRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *ConvertRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type ConvertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the base code e.g USD
	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	// the target code e.g GBP
	To string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	// conversion rate e.g 0.71
	Rate float64 `protobuf:"fixed64,3,opt,name=rate,proto3" json:"rate,omitempty"`
	// converted amount e.g 7.10
	Amount float64 `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *ConvertResponse) Reset() {
	*x = ConvertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currency_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertResponse) ProtoMessage() {}

func (x *ConvertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currency_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertResponse.ProtoReflect.Descriptor instead.
func (*ConvertResponse) Descriptor() ([]byte, []int) {
	return file_proto_currency_proto_rawDescGZIP(), []int{6}
}

func (x *ConvertResponse) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *ConvertResponse) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *ConvertResponse) GetRate() float64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *ConvertResponse) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_proto_currency_proto protoreflect.FileDescriptor

var file_proto_currency_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x22, 0x32, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x0d, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x22, 0x0a, 0x0c, 0x52,
	0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22,
	0x97, 0x01, 0x0a, 0x0d, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x72, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e,
	0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x61,
	0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x72, 0x61, 0x74, 0x65, 0x73, 0x1a,
	0x38, 0x0a, 0x0a, 0x52, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4c, 0x0a, 0x0e, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x61, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x72, 0x61,
	0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xc4, 0x01, 0x0a, 0x08, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x3a, 0x0a, 0x05, 0x43, 0x6f, 0x64, 0x65, 0x73,
	0x12, 0x16, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x05, 0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e,
	0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x40, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x12, 0x18, 0x2e, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_currency_proto_rawDescOnce sync.Once
	file_proto_currency_proto_rawDescData = file_proto_currency_proto_rawDesc
)

func file_proto_currency_proto_rawDescGZIP() []byte {
	file_proto_currency_proto_rawDescOnce.Do(func() {
		file_proto_currency_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_currency_proto_rawDescData)
	})
	return file_proto_currency_proto_rawDescData
}

var file_proto_currency_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_currency_proto_goTypes = []interface{}{
	(*Code)(nil),            // 0: currency.Code
	(*CodesRequest)(nil),    // 1: currency.CodesRequest
	(*CodesResponse)(nil),   // 2: currency.CodesResponse
	(*RatesRequest)(nil),    // 3: currency.RatesRequest
	(*RatesResponse)(nil),   // 4: currency.RatesResponse
	(*ConvertRequest)(nil),  // 5: currency.ConvertRequest
	(*ConvertResponse)(nil), // 6: currency.ConvertResponse
	nil,                     // 7: currency.RatesResponse.RatesEntry
}
var file_proto_currency_proto_depIdxs = []int32{
	0, // 0: currency.CodesResponse.codes:type_name -> currency.Code
	7, // 1: currency.RatesResponse.rates:type_name -> currency.RatesResponse.RatesEntry
	1, // 2: currency.Currency.Codes:input_type -> currency.CodesRequest
	3, // 3: currency.Currency.Rates:input_type -> currency.RatesRequest
	5, // 4: currency.Currency.Convert:input_type -> currency.ConvertRequest
	2, // 5: currency.Currency.Codes:output_type -> currency.CodesResponse
	4, // 6: currency.Currency.Rates:output_type -> currency.RatesResponse
	6, // 7: currency.Currency.Convert:output_type -> currency.ConvertResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_currency_proto_init() }
func file_proto_currency_proto_init() {
	if File_proto_currency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_currency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Code); i {
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
		file_proto_currency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodesRequest); i {
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
		file_proto_currency_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodesResponse); i {
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
		file_proto_currency_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RatesRequest); i {
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
		file_proto_currency_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RatesResponse); i {
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
		file_proto_currency_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertRequest); i {
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
		file_proto_currency_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertResponse); i {
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
			RawDescriptor: file_proto_currency_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_currency_proto_goTypes,
		DependencyIndexes: file_proto_currency_proto_depIdxs,
		MessageInfos:      file_proto_currency_proto_msgTypes,
	}.Build()
	File_proto_currency_proto = out.File
	file_proto_currency_proto_rawDesc = nil
	file_proto_currency_proto_goTypes = nil
	file_proto_currency_proto_depIdxs = nil
}
