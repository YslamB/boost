// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: boost_server.proto

package im_repo_boost

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

type BoostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PostId        string `protobuf:"bytes,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	DurationHours int32  `protobuf:"varint,3,opt,name=duration_hours,json=durationHours,proto3" json:"duration_hours,omitempty"`
	ViewTarget    int32  `protobuf:"varint,4,opt,name=view_target,json=viewTarget,proto3" json:"view_target,omitempty"`
}

func (x *BoostRequest) Reset() {
	*x = BoostRequest{}
	mi := &file_boost_server_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BoostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoostRequest) ProtoMessage() {}

func (x *BoostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boost_server_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoostRequest.ProtoReflect.Descriptor instead.
func (*BoostRequest) Descriptor() ([]byte, []int) {
	return file_boost_server_proto_rawDescGZIP(), []int{0}
}

func (x *BoostRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BoostRequest) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *BoostRequest) GetDurationHours() int32 {
	if x != nil {
		return x.DurationHours
	}
	return 0
}

func (x *BoostRequest) GetViewTarget() int32 {
	if x != nil {
		return x.ViewTarget
	}
	return 0
}

type BoostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *BoostResponse) Reset() {
	*x = BoostResponse{}
	mi := &file_boost_server_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BoostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoostResponse) ProtoMessage() {}

func (x *BoostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_boost_server_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoostResponse.ProtoReflect.Descriptor instead.
func (*BoostResponse) Descriptor() ([]byte, []int) {
	return file_boost_server_proto_rawDescGZIP(), []int{1}
}

func (x *BoostResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BoostResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type BoostStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId string `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
}

func (x *BoostStatsRequest) Reset() {
	*x = BoostStatsRequest{}
	mi := &file_boost_server_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BoostStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoostStatsRequest) ProtoMessage() {}

func (x *BoostStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boost_server_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoostStatsRequest.ProtoReflect.Descriptor instead.
func (*BoostStatsRequest) Descriptor() ([]byte, []int) {
	return file_boost_server_proto_rawDescGZIP(), []int{2}
}

func (x *BoostStatsRequest) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

type BoostStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Views         int32 `protobuf:"varint,1,opt,name=views,proto3" json:"views,omitempty"`
	ProfileClicks int32 `protobuf:"varint,2,opt,name=profile_clicks,json=profileClicks,proto3" json:"profile_clicks,omitempty"`
}

func (x *BoostStatsResponse) Reset() {
	*x = BoostStatsResponse{}
	mi := &file_boost_server_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BoostStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoostStatsResponse) ProtoMessage() {}

func (x *BoostStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_boost_server_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoostStatsResponse.ProtoReflect.Descriptor instead.
func (*BoostStatsResponse) Descriptor() ([]byte, []int) {
	return file_boost_server_proto_rawDescGZIP(), []int{3}
}

func (x *BoostStatsResponse) GetViews() int32 {
	if x != nil {
		return x.Views
	}
	return 0
}

func (x *BoostStatsResponse) GetProfileClicks() int32 {
	if x != nil {
		return x.ProfileClicks
	}
	return 0
}

type BoostedPostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*BoostedPost `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *BoostedPostsResponse) Reset() {
	*x = BoostedPostsResponse{}
	mi := &file_boost_server_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BoostedPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoostedPostsResponse) ProtoMessage() {}

func (x *BoostedPostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_boost_server_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoostedPostsResponse.ProtoReflect.Descriptor instead.
func (*BoostedPostsResponse) Descriptor() ([]byte, []int) {
	return file_boost_server_proto_rawDescGZIP(), []int{4}
}

func (x *BoostedPostsResponse) GetPosts() []*BoostedPost {
	if x != nil {
		return x.Posts
	}
	return nil
}

type BoostedPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId        string `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	UserId        string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Views         int32  `protobuf:"varint,3,opt,name=views,proto3" json:"views,omitempty"`
	ProfileClicks int32  `protobuf:"varint,4,opt,name=profile_clicks,json=profileClicks,proto3" json:"profile_clicks,omitempty"`
}

func (x *BoostedPost) Reset() {
	*x = BoostedPost{}
	mi := &file_boost_server_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BoostedPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoostedPost) ProtoMessage() {}

func (x *BoostedPost) ProtoReflect() protoreflect.Message {
	mi := &file_boost_server_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoostedPost.ProtoReflect.Descriptor instead.
func (*BoostedPost) Descriptor() ([]byte, []int) {
	return file_boost_server_proto_rawDescGZIP(), []int{5}
}

func (x *BoostedPost) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *BoostedPost) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BoostedPost) GetViews() int32 {
	if x != nil {
		return x.Views
	}
	return 0
}

func (x *BoostedPost) GetProfileClicks() int32 {
	if x != nil {
		return x.ProfileClicks
	}
	return 0
}

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	mi := &file_boost_server_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boost_server_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_boost_server_proto_rawDescGZIP(), []int{6}
}

var File_boost_server_proto protoreflect.FileDescriptor

var file_boost_server_proto_rawDesc = []byte{
	0x0a, 0x12, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x22, 0x88, 0x01, 0x0a, 0x0c,
	0x42, 0x6f, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x25,
	0x0a, 0x0e, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x76, 0x69, 0x65, 0x77,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x43, 0x0a, 0x0d, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x2c, 0x0a, 0x11, 0x42,
	0x6f, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x12, 0x42, 0x6f, 0x6f,
	0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x22, 0x40, 0x0a, 0x14,
	0x42, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x42, 0x6f, 0x6f, 0x73,
	0x74, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x7c,
	0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x22, 0x0e, 0x0a, 0x0c,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xd1, 0x01, 0x0a,
	0x0c, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a,
	0x09, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x62, 0x6f, 0x6f,
	0x73, 0x74, 0x2e, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x73,
	0x74, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x13, 0x2e, 0x62, 0x6f, 0x6f, 0x73, 0x74,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x62, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x50, 0x6f, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x62, 0x6f,
	0x6f, 0x73, 0x74, 0x2e, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x42, 0x6f,
	0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x15, 0x5a, 0x13, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x2f, 0x69, 0x6d, 0x5f, 0x72, 0x65, 0x70,
	0x6f, 0x5f, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_boost_server_proto_rawDescOnce sync.Once
	file_boost_server_proto_rawDescData = file_boost_server_proto_rawDesc
)

func file_boost_server_proto_rawDescGZIP() []byte {
	file_boost_server_proto_rawDescOnce.Do(func() {
		file_boost_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_boost_server_proto_rawDescData)
	})
	return file_boost_server_proto_rawDescData
}

var file_boost_server_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_boost_server_proto_goTypes = []any{
	(*BoostRequest)(nil),         // 0: boost.BoostRequest
	(*BoostResponse)(nil),        // 1: boost.BoostResponse
	(*BoostStatsRequest)(nil),    // 2: boost.BoostStatsRequest
	(*BoostStatsResponse)(nil),   // 3: boost.BoostStatsResponse
	(*BoostedPostsResponse)(nil), // 4: boost.BoostedPostsResponse
	(*BoostedPost)(nil),          // 5: boost.BoostedPost
	(*EmptyRequest)(nil),         // 6: boost.EmptyRequest
}
var file_boost_server_proto_depIdxs = []int32{
	5, // 0: boost.BoostedPostsResponse.posts:type_name -> boost.BoostedPost
	0, // 1: boost.BoostService.BoostPost:input_type -> boost.BoostRequest
	6, // 2: boost.BoostService.GetBoostedPosts:input_type -> boost.EmptyRequest
	2, // 3: boost.BoostService.GetBoostStats:input_type -> boost.BoostStatsRequest
	1, // 4: boost.BoostService.BoostPost:output_type -> boost.BoostResponse
	4, // 5: boost.BoostService.GetBoostedPosts:output_type -> boost.BoostedPostsResponse
	3, // 6: boost.BoostService.GetBoostStats:output_type -> boost.BoostStatsResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_boost_server_proto_init() }
func file_boost_server_proto_init() {
	if File_boost_server_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_boost_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_boost_server_proto_goTypes,
		DependencyIndexes: file_boost_server_proto_depIdxs,
		MessageInfos:      file_boost_server_proto_msgTypes,
	}.Build()
	File_boost_server_proto = out.File
	file_boost_server_proto_rawDesc = nil
	file_boost_server_proto_goTypes = nil
	file_boost_server_proto_depIdxs = nil
}
