// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: node/v1/node.proto

package nodev1

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

type ComputeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start float64 `protobuf:"fixed64,1,opt,name=start,proto3" json:"start,omitempty"`
	End   float64 `protobuf:"fixed64,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *ComputeRequest) Reset() {
	*x = ComputeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeRequest) ProtoMessage() {}

func (x *ComputeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeRequest.ProtoReflect.Descriptor instead.
func (*ComputeRequest) Descriptor() ([]byte, []int) {
	return file_node_v1_node_proto_rawDescGZIP(), []int{0}
}

func (x *ComputeRequest) GetStart() float64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *ComputeRequest) GetEnd() float64 {
	if x != nil {
		return x.End
	}
	return 0
}

type ComputeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum           float64 `protobuf:"fixed64,1,opt,name=sum,proto3" json:"sum,omitempty"`
	ExecutionTime int64   `protobuf:"varint,2,opt,name=execution_time,json=executionTime,proto3" json:"execution_time,omitempty"`
}

func (x *ComputeResponse) Reset() {
	*x = ComputeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeResponse) ProtoMessage() {}

func (x *ComputeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeResponse.ProtoReflect.Descriptor instead.
func (*ComputeResponse) Descriptor() ([]byte, []int) {
	return file_node_v1_node_proto_rawDescGZIP(), []int{1}
}

func (x *ComputeResponse) GetSum() float64 {
	if x != nil {
		return x.Sum
	}
	return 0
}

func (x *ComputeResponse) GetExecutionTime() int64 {
	if x != nil {
		return x.ExecutionTime
	}
	return 0
}

var File_node_v1_node_proto protoreflect.FileDescriptor

var file_node_v1_node_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x76, 0x31, 0x22, 0x38, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x4a, 0x0a,
	0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x73,
	0x75, 0x6d, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x5c, 0x0a, 0x10, 0x4c, 0x65, 0x69,
	0x62, 0x6e, 0x69, 0x7a, 0x50, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a,
	0x07, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x88, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x4e,
	0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x12, 0x69, 0x64, 0x6c, 0x2f,
	0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x6f, 0x64, 0x65, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x50, 0x4e, 0x58, 0xaa, 0x02, 0x0d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x4e, 0x6f, 0x64,
	0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x4e, 0x6f, 0x64,
	0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3a, 0x4e, 0x6f, 0x64, 0x65, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_v1_node_proto_rawDescOnce sync.Once
	file_node_v1_node_proto_rawDescData = file_node_v1_node_proto_rawDesc
)

func file_node_v1_node_proto_rawDescGZIP() []byte {
	file_node_v1_node_proto_rawDescOnce.Do(func() {
		file_node_v1_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_v1_node_proto_rawDescData)
	})
	return file_node_v1_node_proto_rawDescData
}

var file_node_v1_node_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_node_v1_node_proto_goTypes = []interface{}{
	(*ComputeRequest)(nil),  // 0: proto.node.v1.ComputeRequest
	(*ComputeResponse)(nil), // 1: proto.node.v1.ComputeResponse
}
var file_node_v1_node_proto_depIdxs = []int32{
	0, // 0: proto.node.v1.LeibnizPiService.Compute:input_type -> proto.node.v1.ComputeRequest
	1, // 1: proto.node.v1.LeibnizPiService.Compute:output_type -> proto.node.v1.ComputeResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_node_v1_node_proto_init() }
func file_node_v1_node_proto_init() {
	if File_node_v1_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_v1_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeRequest); i {
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
		file_node_v1_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeResponse); i {
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
			RawDescriptor: file_node_v1_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_node_v1_node_proto_goTypes,
		DependencyIndexes: file_node_v1_node_proto_depIdxs,
		MessageInfos:      file_node_v1_node_proto_msgTypes,
	}.Build()
	File_node_v1_node_proto = out.File
	file_node_v1_node_proto_rawDesc = nil
	file_node_v1_node_proto_goTypes = nil
	file_node_v1_node_proto_depIdxs = nil
}
