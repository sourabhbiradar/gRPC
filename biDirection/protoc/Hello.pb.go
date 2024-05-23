// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: Hello.proto

package __

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

type CReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reqt string `protobuf:"bytes,1,opt,name=Reqt,proto3" json:"Reqt,omitempty"`
}

func (x *CReq) Reset() {
	*x = CReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CReq) ProtoMessage() {}

func (x *CReq) ProtoReflect() protoreflect.Message {
	mi := &file_Hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CReq.ProtoReflect.Descriptor instead.
func (*CReq) Descriptor() ([]byte, []int) {
	return file_Hello_proto_rawDescGZIP(), []int{0}
}

func (x *CReq) GetReqt() string {
	if x != nil {
		return x.Reqt
	}
	return ""
}

type SResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resp string `protobuf:"bytes,1,opt,name=Resp,proto3" json:"Resp,omitempty"`
}

func (x *SResp) Reset() {
	*x = SResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SResp) ProtoMessage() {}

func (x *SResp) ProtoReflect() protoreflect.Message {
	mi := &file_Hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SResp.ProtoReflect.Descriptor instead.
func (*SResp) Descriptor() ([]byte, []int) {
	return file_Hello_proto_rawDescGZIP(), []int{1}
}

func (x *SResp) GetResp() string {
	if x != nil {
		return x.Resp
	}
	return ""
}

var File_Hello_proto protoreflect.FileDescriptor

var file_Hello_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a,
	0x04, 0x43, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x65, 0x71, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x65, 0x71, 0x74, 0x22, 0x1b, 0x0a, 0x05, 0x53, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x52, 0x65, 0x73, 0x70, 0x32, 0x29, 0x0a, 0x05, 0x42, 0x69, 0x44, 0x69, 0x72, 0x12,
	0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x05,
	0x2e, 0x43, 0x52, 0x65, 0x71, 0x1a, 0x06, 0x2e, 0x53, 0x52, 0x65, 0x73, 0x70, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Hello_proto_rawDescOnce sync.Once
	file_Hello_proto_rawDescData = file_Hello_proto_rawDesc
)

func file_Hello_proto_rawDescGZIP() []byte {
	file_Hello_proto_rawDescOnce.Do(func() {
		file_Hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_Hello_proto_rawDescData)
	})
	return file_Hello_proto_rawDescData
}

var file_Hello_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Hello_proto_goTypes = []interface{}{
	(*CReq)(nil),  // 0: CReq
	(*SResp)(nil), // 1: SResp
}
var file_Hello_proto_depIdxs = []int32{
	0, // 0: BiDir.ServerReply:input_type -> CReq
	1, // 1: BiDir.ServerReply:output_type -> SResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Hello_proto_init() }
func file_Hello_proto_init() {
	if File_Hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CReq); i {
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
		file_Hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SResp); i {
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
			RawDescriptor: file_Hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Hello_proto_goTypes,
		DependencyIndexes: file_Hello_proto_depIdxs,
		MessageInfos:      file_Hello_proto_msgTypes,
	}.Build()
	File_Hello_proto = out.File
	file_Hello_proto_rawDesc = nil
	file_Hello_proto_goTypes = nil
	file_Hello_proto_depIdxs = nil
}