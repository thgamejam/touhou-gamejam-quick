// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: api/download/v1/download.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetDownloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`           // 文件唯一标识符，SHA256
	Location string `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"` // 服务器组地点
}

func (x *GetDownloadRequest) Reset() {
	*x = GetDownloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_download_v1_download_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDownloadRequest) ProtoMessage() {}

func (x *GetDownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_download_v1_download_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDownloadRequest.ProtoReflect.Descriptor instead.
func (*GetDownloadRequest) Descriptor() ([]byte, []int) {
	return file_api_download_v1_download_proto_rawDescGZIP(), []int{0}
}

func (x *GetDownloadRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GetDownloadRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type GetDownloadReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetDownloadReply) Reset() {
	*x = GetDownloadReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_download_v1_download_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDownloadReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDownloadReply) ProtoMessage() {}

func (x *GetDownloadReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_download_v1_download_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDownloadReply.ProtoReflect.Descriptor instead.
func (*GetDownloadReply) Descriptor() ([]byte, []int) {
	return file_api_download_v1_download_proto_rawDescGZIP(), []int{1}
}

func (x *GetDownloadReply) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_api_download_v1_download_proto protoreflect.FileDescriptor

var file_api_download_v1_download_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x42, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x24, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x32, 0x7d, 0x0a, 0x08, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x71, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x55, 0x52, 0x4c, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x2f, 0x7b, 0x6b, 0x65, 0x79, 0x7d, 0x42, 0x30, 0x0a, 0x0f, 0x61, 0x70, 0x69, 0x2e,
	0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1b, 0x64,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_download_v1_download_proto_rawDescOnce sync.Once
	file_api_download_v1_download_proto_rawDescData = file_api_download_v1_download_proto_rawDesc
)

func file_api_download_v1_download_proto_rawDescGZIP() []byte {
	file_api_download_v1_download_proto_rawDescOnce.Do(func() {
		file_api_download_v1_download_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_download_v1_download_proto_rawDescData)
	})
	return file_api_download_v1_download_proto_rawDescData
}

var file_api_download_v1_download_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_download_v1_download_proto_goTypes = []interface{}{
	(*GetDownloadRequest)(nil), // 0: api.download.v1.GetDownloadRequest
	(*GetDownloadReply)(nil),   // 1: api.download.v1.GetDownloadReply
}
var file_api_download_v1_download_proto_depIdxs = []int32{
	0, // 0: api.download.v1.Download.GetDownloadURL:input_type -> api.download.v1.GetDownloadRequest
	1, // 1: api.download.v1.Download.GetDownloadURL:output_type -> api.download.v1.GetDownloadReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_download_v1_download_proto_init() }
func file_api_download_v1_download_proto_init() {
	if File_api_download_v1_download_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_download_v1_download_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDownloadRequest); i {
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
		file_api_download_v1_download_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDownloadReply); i {
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
			RawDescriptor: file_api_download_v1_download_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_download_v1_download_proto_goTypes,
		DependencyIndexes: file_api_download_v1_download_proto_depIdxs,
		MessageInfos:      file_api_download_v1_download_proto_msgTypes,
	}.Build()
	File_api_download_v1_download_proto = out.File
	file_api_download_v1_download_proto_rawDesc = nil
	file_api_download_v1_download_proto_goTypes = nil
	file_api_download_v1_download_proto_depIdxs = nil
}
