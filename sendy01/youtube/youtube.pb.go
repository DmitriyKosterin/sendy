// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: youtube.proto

package youtube

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

type PlaylistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaylistId string `protobuf:"bytes,1,opt,name=playlist_id,json=playlistId,proto3" json:"playlist_id,omitempty"`
}

func (x *PlaylistRequest) Reset() {
	*x = PlaylistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_youtube_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaylistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaylistRequest) ProtoMessage() {}

func (x *PlaylistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_youtube_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaylistRequest.ProtoReflect.Descriptor instead.
func (*PlaylistRequest) Descriptor() ([]byte, []int) {
	return file_youtube_proto_rawDescGZIP(), []int{0}
}

func (x *PlaylistRequest) GetPlaylistId() string {
	if x != nil {
		return x.PlaylistId
	}
	return ""
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId string `protobuf:"bytes,1,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_youtube_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_youtube_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_youtube_proto_rawDescGZIP(), []int{1}
}

func (x *Video) GetVideoId() string {
	if x != nil {
		return x.VideoId
	}
	return ""
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type PlaylistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Videos []*Video `protobuf:"bytes,1,rep,name=videos,proto3" json:"videos,omitempty"`
}

func (x *PlaylistResponse) Reset() {
	*x = PlaylistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_youtube_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaylistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaylistResponse) ProtoMessage() {}

func (x *PlaylistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_youtube_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaylistResponse.ProtoReflect.Descriptor instead.
func (*PlaylistResponse) Descriptor() ([]byte, []int) {
	return file_youtube_proto_rawDescGZIP(), []int{2}
}

func (x *PlaylistResponse) GetVideos() []*Video {
	if x != nil {
		return x.Videos
	}
	return nil
}

var File_youtube_proto protoreflect.FileDescriptor

var file_youtube_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x32, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73,
	0x74, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x19, 0x0a, 0x08,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x32, 0x0a,
	0x10, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x06, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x73, 0x32, 0x3d, 0x0a, 0x07, 0x59, 0x6f, 0x75, 0x54, 0x75, 0x62, 0x65, 0x12, 0x32, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x2a, 0x5a, 0x28, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x2f, 0x64, 0x61, 0x6e, 0x63, 0x72, 0x69,
	0x6d, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x79, 0x2f, 0x73, 0x65, 0x6e,
	0x64, 0x79, 0x30, 0x31, 0x2f, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_youtube_proto_rawDescOnce sync.Once
	file_youtube_proto_rawDescData = file_youtube_proto_rawDesc
)

func file_youtube_proto_rawDescGZIP() []byte {
	file_youtube_proto_rawDescOnce.Do(func() {
		file_youtube_proto_rawDescData = protoimpl.X.CompressGZIP(file_youtube_proto_rawDescData)
	})
	return file_youtube_proto_rawDescData
}

var file_youtube_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_youtube_proto_goTypes = []interface{}{
	(*PlaylistRequest)(nil),  // 0: PlaylistRequest
	(*Video)(nil),            // 1: Video
	(*PlaylistResponse)(nil), // 2: PlaylistResponse
}
var file_youtube_proto_depIdxs = []int32{
	1, // 0: PlaylistResponse.videos:type_name -> Video
	0, // 1: YouTube.GetPlaylist:input_type -> PlaylistRequest
	2, // 2: YouTube.GetPlaylist:output_type -> PlaylistResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_youtube_proto_init() }
func file_youtube_proto_init() {
	if File_youtube_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_youtube_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaylistRequest); i {
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
		file_youtube_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_youtube_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaylistResponse); i {
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
			RawDescriptor: file_youtube_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_youtube_proto_goTypes,
		DependencyIndexes: file_youtube_proto_depIdxs,
		MessageInfos:      file_youtube_proto_msgTypes,
	}.Build()
	File_youtube_proto = out.File
	file_youtube_proto_rawDesc = nil
	file_youtube_proto_goTypes = nil
	file_youtube_proto_depIdxs = nil
}
