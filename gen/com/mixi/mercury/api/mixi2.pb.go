// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: com/mixi/mercury/api/mixi2.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreatePostRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Text          string                 `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePostRequest) Reset() {
	*x = CreatePostRequest{}
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostRequest) ProtoMessage() {}

func (x *CreatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostRequest.ProtoReflect.Descriptor instead.
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return file_com_mixi_mercury_api_mixi2_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePostRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type ImageInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LargeUrl      string                 `protobuf:"bytes,1,opt,name=large_url,json=largeUrl,proto3" json:"large_url,omitempty"`
	LargeMimeType string                 `protobuf:"bytes,2,opt,name=large_mime_type,json=largeMimeType,proto3" json:"large_mime_type,omitempty"`
	LargeWidth    int32                  `protobuf:"varint,3,opt,name=large_width,json=largeWidth,proto3" json:"large_width,omitempty"`
	LargeHeight   int32                  `protobuf:"varint,4,opt,name=large_height,json=largeHeight,proto3" json:"large_height,omitempty"`
	SmallUrl      string                 `protobuf:"bytes,5,opt,name=small_url,json=smallUrl,proto3" json:"small_url,omitempty"`
	SmallMimeType string                 `protobuf:"bytes,6,opt,name=small_mime_type,json=smallMimeType,proto3" json:"small_mime_type,omitempty"`
	SmallWidth    int32                  `protobuf:"varint,7,opt,name=small_width,json=smallWidth,proto3" json:"small_width,omitempty"`
	SmallHeight   int32                  `protobuf:"varint,8,opt,name=small_height,json=smallHeight,proto3" json:"small_height,omitempty"`
	Unknown9      string                 `protobuf:"bytes,9,opt,name=unknown9,proto3" json:"unknown9,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImageInfo) Reset() {
	*x = ImageInfo{}
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageInfo) ProtoMessage() {}

func (x *ImageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageInfo.ProtoReflect.Descriptor instead.
func (*ImageInfo) Descriptor() ([]byte, []int) {
	return file_com_mixi_mercury_api_mixi2_proto_rawDescGZIP(), []int{1}
}

func (x *ImageInfo) GetLargeUrl() string {
	if x != nil {
		return x.LargeUrl
	}
	return ""
}

func (x *ImageInfo) GetLargeMimeType() string {
	if x != nil {
		return x.LargeMimeType
	}
	return ""
}

func (x *ImageInfo) GetLargeWidth() int32 {
	if x != nil {
		return x.LargeWidth
	}
	return 0
}

func (x *ImageInfo) GetLargeHeight() int32 {
	if x != nil {
		return x.LargeHeight
	}
	return 0
}

func (x *ImageInfo) GetSmallUrl() string {
	if x != nil {
		return x.SmallUrl
	}
	return ""
}

func (x *ImageInfo) GetSmallMimeType() string {
	if x != nil {
		return x.SmallMimeType
	}
	return ""
}

func (x *ImageInfo) GetSmallWidth() int32 {
	if x != nil {
		return x.SmallWidth
	}
	return 0
}

func (x *ImageInfo) GetSmallHeight() int32 {
	if x != nil {
		return x.SmallHeight
	}
	return 0
}

func (x *ImageInfo) GetUnknown9() string {
	if x != nil {
		return x.Unknown9
	}
	return ""
}

type MediaItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MediaId       string                 `protobuf:"bytes,1,opt,name=media_id,json=mediaId,proto3" json:"media_id,omitempty"`
	Unknown3      int32                  `protobuf:"varint,3,opt,name=unknown3,proto3" json:"unknown3,omitempty"`
	Unknown4      int32                  `protobuf:"varint,4,opt,name=unknown4,proto3" json:"unknown4,omitempty"`
	Image         *ImageInfo             `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MediaItem) Reset() {
	*x = MediaItem{}
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MediaItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaItem) ProtoMessage() {}

func (x *MediaItem) ProtoReflect() protoreflect.Message {
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaItem.ProtoReflect.Descriptor instead.
func (*MediaItem) Descriptor() ([]byte, []int) {
	return file_com_mixi_mercury_api_mixi2_proto_rawDescGZIP(), []int{2}
}

func (x *MediaItem) GetMediaId() string {
	if x != nil {
		return x.MediaId
	}
	return ""
}

func (x *MediaItem) GetUnknown3() int32 {
	if x != nil {
		return x.Unknown3
	}
	return 0
}

func (x *MediaItem) GetUnknown4() int32 {
	if x != nil {
		return x.Unknown4
	}
	return 0
}

func (x *MediaItem) GetImage() *ImageInfo {
	if x != nil {
		return x.Image
	}
	return nil
}

type Timestamp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Timestamp     int64                  `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Timestamp) Reset() {
	*x = Timestamp{}
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Timestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timestamp) ProtoMessage() {}

func (x *Timestamp) ProtoReflect() protoreflect.Message {
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timestamp.ProtoReflect.Descriptor instead.
func (*Timestamp) Descriptor() ([]byte, []int) {
	return file_com_mixi_mercury_api_mixi2_proto_rawDescGZIP(), []int{3}
}

func (x *Timestamp) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Post struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PostId        string                 `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	TimestampNs   string                 `protobuf:"bytes,2,opt,name=timestamp_ns,json=timestampNs,proto3" json:"timestamp_ns,omitempty"`
	CreatedAt     *Timestamp             `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Unknown4Id    string                 `protobuf:"bytes,4,opt,name=unknown4_id,json=unknown4Id,proto3" json:"unknown4_id,omitempty"`
	Media         *MediaItem             `protobuf:"bytes,7,opt,name=media,proto3" json:"media,omitempty"`
	Text          string                 `protobuf:"bytes,12,opt,name=text,proto3" json:"text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Post) Reset() {
	*x = Post{}
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_com_mixi_mercury_api_mixi2_proto_rawDescGZIP(), []int{4}
}

func (x *Post) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *Post) GetTimestampNs() string {
	if x != nil {
		return x.TimestampNs
	}
	return ""
}

func (x *Post) GetCreatedAt() *Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Post) GetUnknown4Id() string {
	if x != nil {
		return x.Unknown4Id
	}
	return ""
}

func (x *Post) GetMedia() *MediaItem {
	if x != nil {
		return x.Media
	}
	return nil
}

func (x *Post) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type CreatePostResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Post          *Post                  `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePostResponse) Reset() {
	*x = CreatePostResponse{}
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostResponse) ProtoMessage() {}

func (x *CreatePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostResponse.ProtoReflect.Descriptor instead.
func (*CreatePostResponse) Descriptor() ([]byte, []int) {
	return file_com_mixi_mercury_api_mixi2_proto_rawDescGZIP(), []int{5}
}

func (x *CreatePostResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type GetPostsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Text          string                 `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPostsRequest) Reset() {
	*x = GetPostsRequest{}
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostsRequest) ProtoMessage() {}

func (x *GetPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostsRequest.ProtoReflect.Descriptor instead.
func (*GetPostsRequest) Descriptor() ([]byte, []int) {
	return file_com_mixi_mercury_api_mixi2_proto_rawDescGZIP(), []int{6}
}

func (x *GetPostsRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type GetPostsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Post          *Post                  `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPostsResponse) Reset() {
	*x = GetPostsResponse{}
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostsResponse) ProtoMessage() {}

func (x *GetPostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostsResponse.ProtoReflect.Descriptor instead.
func (*GetPostsResponse) Descriptor() ([]byte, []int) {
	return file_com_mixi_mercury_api_mixi2_proto_rawDescGZIP(), []int{7}
}

func (x *GetPostsResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type SwitchPersonaRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PersonaId     string                 `protobuf:"bytes,1,opt,name=persona_id,json=personaId,proto3" json:"persona_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SwitchPersonaRequest) Reset() {
	*x = SwitchPersonaRequest{}
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SwitchPersonaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwitchPersonaRequest) ProtoMessage() {}

func (x *SwitchPersonaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwitchPersonaRequest.ProtoReflect.Descriptor instead.
func (*SwitchPersonaRequest) Descriptor() ([]byte, []int) {
	return file_com_mixi_mercury_api_mixi2_proto_rawDescGZIP(), []int{8}
}

func (x *SwitchPersonaRequest) GetPersonaId() string {
	if x != nil {
		return x.PersonaId
	}
	return ""
}

type SwitchPersonaResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SwitchPersonaResponse) Reset() {
	*x = SwitchPersonaResponse{}
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SwitchPersonaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwitchPersonaResponse) ProtoMessage() {}

func (x *SwitchPersonaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwitchPersonaResponse.ProtoReflect.Descriptor instead.
func (*SwitchPersonaResponse) Descriptor() ([]byte, []int) {
	return file_com_mixi_mercury_api_mixi2_proto_rawDescGZIP(), []int{9}
}

type GetPersonasRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PersonaId     string                 `protobuf:"bytes,1,opt,name=persona_id,json=personaId,proto3" json:"persona_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPersonasRequest) Reset() {
	*x = GetPersonasRequest{}
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPersonasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonasRequest) ProtoMessage() {}

func (x *GetPersonasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonasRequest.ProtoReflect.Descriptor instead.
func (*GetPersonasRequest) Descriptor() ([]byte, []int) {
	return file_com_mixi_mercury_api_mixi2_proto_rawDescGZIP(), []int{10}
}

func (x *GetPersonasRequest) GetPersonaId() string {
	if x != nil {
		return x.PersonaId
	}
	return ""
}

type GetPersonasResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Persona       *Persona               `protobuf:"bytes,1,opt,name=persona,proto3" json:"persona,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPersonasResponse) Reset() {
	*x = GetPersonasResponse{}
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPersonasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonasResponse) ProtoMessage() {}

func (x *GetPersonasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonasResponse.ProtoReflect.Descriptor instead.
func (*GetPersonasResponse) Descriptor() ([]byte, []int) {
	return file_com_mixi_mercury_api_mixi2_proto_rawDescGZIP(), []int{11}
}

func (x *GetPersonasResponse) GetPersona() *Persona {
	if x != nil {
		return x.Persona
	}
	return nil
}

type Persona struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PersonaId     string                 `protobuf:"bytes,1,opt,name=persona_id,json=personaId,proto3" json:"persona_id,omitempty"`
	AccountName   string                 `protobuf:"bytes,2,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	DisplayName   string                 `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	IconUrl       string                 `protobuf:"bytes,4,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url,omitempty"`
	Bio           string                 `protobuf:"bytes,8,opt,name=bio,proto3" json:"bio,omitempty"`
	ProfileUrl    string                 `protobuf:"bytes,9,opt,name=profile_url,json=profileUrl,proto3" json:"profile_url,omitempty"`
	CreatedAt     *Timestamp             `protobuf:"bytes,17,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Persona) Reset() {
	*x = Persona{}
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Persona) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Persona) ProtoMessage() {}

func (x *Persona) ProtoReflect() protoreflect.Message {
	mi := &file_com_mixi_mercury_api_mixi2_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Persona.ProtoReflect.Descriptor instead.
func (*Persona) Descriptor() ([]byte, []int) {
	return file_com_mixi_mercury_api_mixi2_proto_rawDescGZIP(), []int{12}
}

func (x *Persona) GetPersonaId() string {
	if x != nil {
		return x.PersonaId
	}
	return ""
}

func (x *Persona) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *Persona) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Persona) GetIconUrl() string {
	if x != nil {
		return x.IconUrl
	}
	return ""
}

func (x *Persona) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *Persona) GetProfileUrl() string {
	if x != nil {
		return x.ProfileUrl
	}
	return ""
}

func (x *Persona) GetCreatedAt() *Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_com_mixi_mercury_api_mixi2_proto protoreflect.FileDescriptor

const file_com_mixi_mercury_api_mixi2_proto_rawDesc = "" +
	"\n" +
	" com/mixi/mercury/api/mixi2.proto\x12\x14com.mixi.mercury.api\"'\n" +
	"\x11CreatePostRequest\x12\x12\n" +
	"\x04text\x18\x01 \x01(\tR\x04text\"\xb9\x02\n" +
	"\tImageInfo\x12\x1b\n" +
	"\tlarge_url\x18\x01 \x01(\tR\blargeUrl\x12&\n" +
	"\x0flarge_mime_type\x18\x02 \x01(\tR\rlargeMimeType\x12\x1f\n" +
	"\vlarge_width\x18\x03 \x01(\x05R\n" +
	"largeWidth\x12!\n" +
	"\flarge_height\x18\x04 \x01(\x05R\vlargeHeight\x12\x1b\n" +
	"\tsmall_url\x18\x05 \x01(\tR\bsmallUrl\x12&\n" +
	"\x0fsmall_mime_type\x18\x06 \x01(\tR\rsmallMimeType\x12\x1f\n" +
	"\vsmall_width\x18\a \x01(\x05R\n" +
	"smallWidth\x12!\n" +
	"\fsmall_height\x18\b \x01(\x05R\vsmallHeight\x12\x1a\n" +
	"\bunknown9\x18\t \x01(\tR\bunknown9\"\x95\x01\n" +
	"\tMediaItem\x12\x19\n" +
	"\bmedia_id\x18\x01 \x01(\tR\amediaId\x12\x1a\n" +
	"\bunknown3\x18\x03 \x01(\x05R\bunknown3\x12\x1a\n" +
	"\bunknown4\x18\x04 \x01(\x05R\bunknown4\x125\n" +
	"\x05image\x18\x06 \x01(\v2\x1f.com.mixi.mercury.api.ImageInfoR\x05image\")\n" +
	"\tTimestamp\x12\x1c\n" +
	"\ttimestamp\x18\x01 \x01(\x03R\ttimestamp\"\xee\x01\n" +
	"\x04Post\x12\x17\n" +
	"\apost_id\x18\x01 \x01(\tR\x06postId\x12!\n" +
	"\ftimestamp_ns\x18\x02 \x01(\tR\vtimestampNs\x12>\n" +
	"\n" +
	"created_at\x18\x03 \x01(\v2\x1f.com.mixi.mercury.api.TimestampR\tcreatedAt\x12\x1f\n" +
	"\vunknown4_id\x18\x04 \x01(\tR\n" +
	"unknown4Id\x125\n" +
	"\x05media\x18\a \x01(\v2\x1f.com.mixi.mercury.api.MediaItemR\x05media\x12\x12\n" +
	"\x04text\x18\f \x01(\tR\x04text\"D\n" +
	"\x12CreatePostResponse\x12.\n" +
	"\x04post\x18\x01 \x01(\v2\x1a.com.mixi.mercury.api.PostR\x04post\"%\n" +
	"\x0fGetPostsRequest\x12\x12\n" +
	"\x04text\x18\x01 \x01(\tR\x04text\"B\n" +
	"\x10GetPostsResponse\x12.\n" +
	"\x04post\x18\x01 \x01(\v2\x1a.com.mixi.mercury.api.PostR\x04post\"5\n" +
	"\x14SwitchPersonaRequest\x12\x1d\n" +
	"\n" +
	"persona_id\x18\x01 \x01(\tR\tpersonaId\"\x17\n" +
	"\x15SwitchPersonaResponse\"3\n" +
	"\x12GetPersonasRequest\x12\x1d\n" +
	"\n" +
	"persona_id\x18\x01 \x01(\tR\tpersonaId\"N\n" +
	"\x13GetPersonasResponse\x127\n" +
	"\apersona\x18\x01 \x01(\v2\x1d.com.mixi.mercury.api.PersonaR\apersona\"\xfc\x01\n" +
	"\aPersona\x12\x1d\n" +
	"\n" +
	"persona_id\x18\x01 \x01(\tR\tpersonaId\x12!\n" +
	"\faccount_name\x18\x02 \x01(\tR\vaccountName\x12!\n" +
	"\fdisplay_name\x18\x03 \x01(\tR\vdisplayName\x12\x19\n" +
	"\bicon_url\x18\x04 \x01(\tR\aiconUrl\x12\x10\n" +
	"\x03bio\x18\b \x01(\tR\x03bio\x12\x1f\n" +
	"\vprofile_url\x18\t \x01(\tR\n" +
	"profileUrl\x12>\n" +
	"\n" +
	"created_at\x18\x11 \x01(\v2\x1f.com.mixi.mercury.api.TimestampR\tcreatedAt2\xa2\x03\n" +
	"\x0eMercuryService\x12a\n" +
	"\n" +
	"CreatePost\x12'.com.mixi.mercury.api.CreatePostRequest\x1a(.com.mixi.mercury.api.CreatePostResponse\"\x00\x12[\n" +
	"\bGetPosts\x12%.com.mixi.mercury.api.GetPostsRequest\x1a&.com.mixi.mercury.api.GetPostsResponse\"\x00\x12j\n" +
	"\rSwitchPersona\x12*.com.mixi.mercury.api.SwitchPersonaRequest\x1a+.com.mixi.mercury.api.SwitchPersonaResponse\"\x00\x12d\n" +
	"\vGetPersonas\x12(.com.mixi.mercury.api.GetPersonasRequest\x1a).com.mixi.mercury.api.GetPersonasResponse\"\x00B9Z7github.com/matsuu/go-mixi2/gen/com/mixi/mercury/api;apib\x06proto3"

var (
	file_com_mixi_mercury_api_mixi2_proto_rawDescOnce sync.Once
	file_com_mixi_mercury_api_mixi2_proto_rawDescData []byte
)

func file_com_mixi_mercury_api_mixi2_proto_rawDescGZIP() []byte {
	file_com_mixi_mercury_api_mixi2_proto_rawDescOnce.Do(func() {
		file_com_mixi_mercury_api_mixi2_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_mixi_mercury_api_mixi2_proto_rawDesc), len(file_com_mixi_mercury_api_mixi2_proto_rawDesc)))
	})
	return file_com_mixi_mercury_api_mixi2_proto_rawDescData
}

var file_com_mixi_mercury_api_mixi2_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_com_mixi_mercury_api_mixi2_proto_goTypes = []any{
	(*CreatePostRequest)(nil),     // 0: com.mixi.mercury.api.CreatePostRequest
	(*ImageInfo)(nil),             // 1: com.mixi.mercury.api.ImageInfo
	(*MediaItem)(nil),             // 2: com.mixi.mercury.api.MediaItem
	(*Timestamp)(nil),             // 3: com.mixi.mercury.api.Timestamp
	(*Post)(nil),                  // 4: com.mixi.mercury.api.Post
	(*CreatePostResponse)(nil),    // 5: com.mixi.mercury.api.CreatePostResponse
	(*GetPostsRequest)(nil),       // 6: com.mixi.mercury.api.GetPostsRequest
	(*GetPostsResponse)(nil),      // 7: com.mixi.mercury.api.GetPostsResponse
	(*SwitchPersonaRequest)(nil),  // 8: com.mixi.mercury.api.SwitchPersonaRequest
	(*SwitchPersonaResponse)(nil), // 9: com.mixi.mercury.api.SwitchPersonaResponse
	(*GetPersonasRequest)(nil),    // 10: com.mixi.mercury.api.GetPersonasRequest
	(*GetPersonasResponse)(nil),   // 11: com.mixi.mercury.api.GetPersonasResponse
	(*Persona)(nil),               // 12: com.mixi.mercury.api.Persona
}
var file_com_mixi_mercury_api_mixi2_proto_depIdxs = []int32{
	1,  // 0: com.mixi.mercury.api.MediaItem.image:type_name -> com.mixi.mercury.api.ImageInfo
	3,  // 1: com.mixi.mercury.api.Post.created_at:type_name -> com.mixi.mercury.api.Timestamp
	2,  // 2: com.mixi.mercury.api.Post.media:type_name -> com.mixi.mercury.api.MediaItem
	4,  // 3: com.mixi.mercury.api.CreatePostResponse.post:type_name -> com.mixi.mercury.api.Post
	4,  // 4: com.mixi.mercury.api.GetPostsResponse.post:type_name -> com.mixi.mercury.api.Post
	12, // 5: com.mixi.mercury.api.GetPersonasResponse.persona:type_name -> com.mixi.mercury.api.Persona
	3,  // 6: com.mixi.mercury.api.Persona.created_at:type_name -> com.mixi.mercury.api.Timestamp
	0,  // 7: com.mixi.mercury.api.MercuryService.CreatePost:input_type -> com.mixi.mercury.api.CreatePostRequest
	6,  // 8: com.mixi.mercury.api.MercuryService.GetPosts:input_type -> com.mixi.mercury.api.GetPostsRequest
	8,  // 9: com.mixi.mercury.api.MercuryService.SwitchPersona:input_type -> com.mixi.mercury.api.SwitchPersonaRequest
	10, // 10: com.mixi.mercury.api.MercuryService.GetPersonas:input_type -> com.mixi.mercury.api.GetPersonasRequest
	5,  // 11: com.mixi.mercury.api.MercuryService.CreatePost:output_type -> com.mixi.mercury.api.CreatePostResponse
	7,  // 12: com.mixi.mercury.api.MercuryService.GetPosts:output_type -> com.mixi.mercury.api.GetPostsResponse
	9,  // 13: com.mixi.mercury.api.MercuryService.SwitchPersona:output_type -> com.mixi.mercury.api.SwitchPersonaResponse
	11, // 14: com.mixi.mercury.api.MercuryService.GetPersonas:output_type -> com.mixi.mercury.api.GetPersonasResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_com_mixi_mercury_api_mixi2_proto_init() }
func file_com_mixi_mercury_api_mixi2_proto_init() {
	if File_com_mixi_mercury_api_mixi2_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_mixi_mercury_api_mixi2_proto_rawDesc), len(file_com_mixi_mercury_api_mixi2_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_mixi_mercury_api_mixi2_proto_goTypes,
		DependencyIndexes: file_com_mixi_mercury_api_mixi2_proto_depIdxs,
		MessageInfos:      file_com_mixi_mercury_api_mixi2_proto_msgTypes,
	}.Build()
	File_com_mixi_mercury_api_mixi2_proto = out.File
	file_com_mixi_mercury_api_mixi2_proto_goTypes = nil
	file_com_mixi_mercury_api_mixi2_proto_depIdxs = nil
}
