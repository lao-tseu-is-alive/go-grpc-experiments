// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.2
// source: remoteCalculator/protobuf/calc.proto

package calc

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

type CalcRequest_OpType int32

const (
	CalcRequest_Add      CalcRequest_OpType = 0
	CalcRequest_Subtract CalcRequest_OpType = 1
	CalcRequest_Multiply CalcRequest_OpType = 2
	CalcRequest_Divide   CalcRequest_OpType = 3
)

// Enum value maps for CalcRequest_OpType.
var (
	CalcRequest_OpType_name = map[int32]string{
		0: "Add",
		1: "Subtract",
		2: "Multiply",
		3: "Divide",
	}
	CalcRequest_OpType_value = map[string]int32{
		"Add":      0,
		"Subtract": 1,
		"Multiply": 2,
		"Divide":   3,
	}
)

func (x CalcRequest_OpType) Enum() *CalcRequest_OpType {
	p := new(CalcRequest_OpType)
	*p = x
	return p
}

func (x CalcRequest_OpType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CalcRequest_OpType) Descriptor() protoreflect.EnumDescriptor {
	return file_remoteCalculator_protobuf_calc_proto_enumTypes[0].Descriptor()
}

func (CalcRequest_OpType) Type() protoreflect.EnumType {
	return &file_remoteCalculator_protobuf_calc_proto_enumTypes[0]
}

func (x CalcRequest_OpType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CalcRequest_OpType.Descriptor instead.
func (CalcRequest_OpType) EnumDescriptor() ([]byte, []int) {
	return file_remoteCalculator_protobuf_calc_proto_rawDescGZIP(), []int{0, 0}
}

type CalcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num1      float64            `protobuf:"fixed64,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2      float64            `protobuf:"fixed64,2,opt,name=num2,proto3" json:"num2,omitempty"`
	Operation CalcRequest_OpType `protobuf:"varint,3,opt,name=operation,proto3,enum=calc.CalcRequest_OpType" json:"operation,omitempty"`
}

func (x *CalcRequest) Reset() {
	*x = CalcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remoteCalculator_protobuf_calc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcRequest) ProtoMessage() {}

func (x *CalcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remoteCalculator_protobuf_calc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcRequest.ProtoReflect.Descriptor instead.
func (*CalcRequest) Descriptor() ([]byte, []int) {
	return file_remoteCalculator_protobuf_calc_proto_rawDescGZIP(), []int{0}
}

func (x *CalcRequest) GetNum1() float64 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *CalcRequest) GetNum2() float64 {
	if x != nil {
		return x.Num2
	}
	return 0
}

func (x *CalcRequest) GetOperation() CalcRequest_OpType {
	if x != nil {
		return x.Operation
	}
	return CalcRequest_Add
}

type CalcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float64 `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	Error  bool    `protobuf:"varint,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CalcResponse) Reset() {
	*x = CalcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remoteCalculator_protobuf_calc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcResponse) ProtoMessage() {}

func (x *CalcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remoteCalculator_protobuf_calc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcResponse.ProtoReflect.Descriptor instead.
func (*CalcResponse) Descriptor() ([]byte, []int) {
	return file_remoteCalculator_protobuf_calc_proto_rawDescGZIP(), []int{1}
}

func (x *CalcResponse) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *CalcResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

var File_remoteCalculator_protobuf_calc_proto protoreflect.FileDescriptor

var file_remoteCalculator_protobuf_calc_proto_rawDesc = []byte{
	0x0a, 0x24, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x61, 0x6c, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x61, 0x6c, 0x63, 0x22, 0xa8, 0x01, 0x0a,
	0x0b, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x75, 0x6d, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x31,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04,
	0x6e, 0x75, 0x6d, 0x32, 0x12, 0x36, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43,
	0x61, 0x6c, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x39, 0x0a, 0x06,
	0x4f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x53, 0x75, 0x62, 0x74, 0x72, 0x61, 0x63, 0x74, 0x10, 0x01, 0x12, 0x0c, 0x0a,
	0x08, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44,
	0x69, 0x76, 0x69, 0x64, 0x65, 0x10, 0x03, 0x22, 0x3c, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x3e, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x63, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x63, 0x12, 0x11, 0x2e, 0x63,
	0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x61, 0x6f, 0x2d, 0x74, 0x73, 0x65, 0x75, 0x2d, 0x69, 0x73, 0x2d,
	0x61, 0x6c, 0x69, 0x76, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_remoteCalculator_protobuf_calc_proto_rawDescOnce sync.Once
	file_remoteCalculator_protobuf_calc_proto_rawDescData = file_remoteCalculator_protobuf_calc_proto_rawDesc
)

func file_remoteCalculator_protobuf_calc_proto_rawDescGZIP() []byte {
	file_remoteCalculator_protobuf_calc_proto_rawDescOnce.Do(func() {
		file_remoteCalculator_protobuf_calc_proto_rawDescData = protoimpl.X.CompressGZIP(file_remoteCalculator_protobuf_calc_proto_rawDescData)
	})
	return file_remoteCalculator_protobuf_calc_proto_rawDescData
}

var file_remoteCalculator_protobuf_calc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_remoteCalculator_protobuf_calc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_remoteCalculator_protobuf_calc_proto_goTypes = []interface{}{
	(CalcRequest_OpType)(0), // 0: calc.CalcRequest.OpType
	(*CalcRequest)(nil),     // 1: calc.CalcRequest
	(*CalcResponse)(nil),    // 2: calc.CalcResponse
}
var file_remoteCalculator_protobuf_calc_proto_depIdxs = []int32{
	0, // 0: calc.CalcRequest.operation:type_name -> calc.CalcRequest.OpType
	1, // 1: calc.CalcService.Calc:input_type -> calc.CalcRequest
	2, // 2: calc.CalcService.Calc:output_type -> calc.CalcResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_remoteCalculator_protobuf_calc_proto_init() }
func file_remoteCalculator_protobuf_calc_proto_init() {
	if File_remoteCalculator_protobuf_calc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_remoteCalculator_protobuf_calc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalcRequest); i {
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
		file_remoteCalculator_protobuf_calc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalcResponse); i {
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
			RawDescriptor: file_remoteCalculator_protobuf_calc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_remoteCalculator_protobuf_calc_proto_goTypes,
		DependencyIndexes: file_remoteCalculator_protobuf_calc_proto_depIdxs,
		EnumInfos:         file_remoteCalculator_protobuf_calc_proto_enumTypes,
		MessageInfos:      file_remoteCalculator_protobuf_calc_proto_msgTypes,
	}.Build()
	File_remoteCalculator_protobuf_calc_proto = out.File
	file_remoteCalculator_protobuf_calc_proto_rawDesc = nil
	file_remoteCalculator_protobuf_calc_proto_goTypes = nil
	file_remoteCalculator_protobuf_calc_proto_depIdxs = nil
}