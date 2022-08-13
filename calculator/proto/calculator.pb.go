// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.1
// source: calculator.proto

package proto

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

type PrimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *PrimeRequest) Reset() {
	*x = PrimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeRequest) ProtoMessage() {}

func (x *PrimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeRequest.ProtoReflect.Descriptor instead.
func (*PrimeRequest) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *PrimeRequest) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type PrimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PrimeResponse) Reset() {
	*x = PrimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeResponse) ProtoMessage() {}

func (x *PrimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeResponse.ProtoReflect.Descriptor instead.
func (*PrimeResponse) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *PrimeResponse) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

type Inp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *Inp) Reset() {
	*x = Inp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inp) ProtoMessage() {}

func (x *Inp) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inp.ProtoReflect.Descriptor instead.
func (*Inp) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{2}
}

func (x *Inp) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type Outp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res float64 `protobuf:"fixed64,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *Outp) Reset() {
	*x = Outp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outp) ProtoMessage() {}

func (x *Outp) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outp.ProtoReflect.Descriptor instead.
func (*Outp) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{3}
}

func (x *Outp) GetRes() float64 {
	if x != nil {
		return x.Res
	}
	return 0
}

type Number struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inp int64 `protobuf:"varint,1,opt,name=inp,proto3" json:"inp,omitempty"`
}

func (x *Number) Reset() {
	*x = Number{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Number) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Number) ProtoMessage() {}

func (x *Number) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Number.ProtoReflect.Descriptor instead.
func (*Number) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{4}
}

func (x *Number) GetInp() int64 {
	if x != nil {
		return x.Inp
	}
	return 0
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Max int64 `protobuf:"varint,1,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{5}
}

func (x *Result) GetMax() int64 {
	if x != nil {
		return x.Max
	}
	return 0
}

type SqrtRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *SqrtRequest) Reset() {
	*x = SqrtRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SqrtRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SqrtRequest) ProtoMessage() {}

func (x *SqrtRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SqrtRequest.ProtoReflect.Descriptor instead.
func (*SqrtRequest) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{6}
}

func (x *SqrtRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type SqrtResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float64 `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SqrtResponse) Reset() {
	*x = SqrtResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SqrtResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SqrtResponse) ProtoMessage() {}

func (x *SqrtResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SqrtResponse.ProtoReflect.Descriptor instead.
func (*SqrtResponse) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{7}
}

func (x *SqrtResponse) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calculator_proto protoreflect.FileDescriptor

var file_calculator_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x20,
	0x0a, 0x0c, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d,
	0x22, 0x27, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x17, 0x0a, 0x03, 0x49, 0x6e, 0x70,
	0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e,
	0x75, 0x6d, 0x22, 0x18, 0x0a, 0x04, 0x4f, 0x75, 0x74, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x72, 0x65, 0x73, 0x22, 0x1a, 0x0a, 0x06,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6e, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x69, 0x6e, 0x70, 0x22, 0x1a, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x6d, 0x61, 0x78, 0x22, 0x25, 0x0a, 0x0b, 0x53, 0x71, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x26, 0x0a, 0x0c, 0x53,
	0x71, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x32, 0xf1, 0x01, 0x0a, 0x09, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x46, 0x69, 0x6e,
	0x64, 0x12, 0x3f, 0x0a, 0x06, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x63, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x30, 0x01, 0x12, 0x2e, 0x0a, 0x07, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x0f, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x49, 0x6e, 0x70, 0x1a, 0x10,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x4f, 0x75, 0x74, 0x70,
	0x28, 0x01, 0x12, 0x38, 0x0a, 0x0a, 0x4d, 0x61, 0x78, 0x54, 0x69, 0x6c, 0x6c, 0x4e, 0x6f, 0x77,
	0x12, 0x12, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x1a, 0x12, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x28, 0x01, 0x30, 0x01, 0x12, 0x39, 0x0a, 0x04,
	0x53, 0x71, 0x72, 0x74, 0x12, 0x17, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x53, 0x71, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x71, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x73, 0x68, 0x35, 0x34, 0x36, 0x31, 0x2f, 0x47,
	0x6f, 0x5f, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calculator_proto_rawDescOnce sync.Once
	file_calculator_proto_rawDescData = file_calculator_proto_rawDesc
)

func file_calculator_proto_rawDescGZIP() []byte {
	file_calculator_proto_rawDescOnce.Do(func() {
		file_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculator_proto_rawDescData)
	})
	return file_calculator_proto_rawDescData
}

var file_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_calculator_proto_goTypes = []interface{}{
	(*PrimeRequest)(nil),  // 0: calculator.PrimeRequest
	(*PrimeResponse)(nil), // 1: calculator.PrimeResponse
	(*Inp)(nil),           // 2: calculator.Inp
	(*Outp)(nil),          // 3: calculator.Outp
	(*Number)(nil),        // 4: calculator.Number
	(*Result)(nil),        // 5: calculator.Result
	(*SqrtRequest)(nil),   // 6: calculator.SqrtRequest
	(*SqrtResponse)(nil),  // 7: calculator.SqrtResponse
}
var file_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.PrimeFind.Primes:input_type -> calculator.PrimeRequest
	2, // 1: calculator.PrimeFind.Average:input_type -> calculator.Inp
	4, // 2: calculator.PrimeFind.MaxTillNow:input_type -> calculator.Number
	6, // 3: calculator.PrimeFind.Sqrt:input_type -> calculator.SqrtRequest
	1, // 4: calculator.PrimeFind.Primes:output_type -> calculator.PrimeResponse
	3, // 5: calculator.PrimeFind.Average:output_type -> calculator.Outp
	5, // 6: calculator.PrimeFind.MaxTillNow:output_type -> calculator.Result
	7, // 7: calculator.PrimeFind.Sqrt:output_type -> calculator.SqrtResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calculator_proto_init() }
func file_calculator_proto_init() {
	if File_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeRequest); i {
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
		file_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeResponse); i {
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
		file_calculator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inp); i {
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
		file_calculator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Outp); i {
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
		file_calculator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Number); i {
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
		file_calculator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_calculator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SqrtRequest); i {
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
		file_calculator_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SqrtResponse); i {
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
			RawDescriptor: file_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_proto_goTypes,
		DependencyIndexes: file_calculator_proto_depIdxs,
		MessageInfos:      file_calculator_proto_msgTypes,
	}.Build()
	File_calculator_proto = out.File
	file_calculator_proto_rawDesc = nil
	file_calculator_proto_goTypes = nil
	file_calculator_proto_depIdxs = nil
}
