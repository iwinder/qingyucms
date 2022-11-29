// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: api/qycms_bff/web/v1/qy_web_menus.proto

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

type MenusWebInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Url      string                  `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Blanked  int32                   `protobuf:"varint,4,opt,name=blanked,proto3" json:"blanked,omitempty"`
	Children []*MenusWebInfoResponse `protobuf:"bytes,7,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *MenusWebInfoResponse) Reset() {
	*x = MenusWebInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenusWebInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenusWebInfoResponse) ProtoMessage() {}

func (x *MenusWebInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenusWebInfoResponse.ProtoReflect.Descriptor instead.
func (*MenusWebInfoResponse) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_web_v1_qy_web_menus_proto_rawDescGZIP(), []int{0}
}

func (x *MenusWebInfoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MenusWebInfoResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *MenusWebInfoResponse) GetBlanked() int32 {
	if x != nil {
		return x.Blanked
	}
	return 0
}

func (x *MenusWebInfoResponse) GetChildren() []*MenusWebInfoResponse {
	if x != nil {
		return x.Children
	}
	return nil
}

type CreateQyWebMenusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateQyWebMenusRequest) Reset() {
	*x = CreateQyWebMenusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQyWebMenusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQyWebMenusRequest) ProtoMessage() {}

func (x *CreateQyWebMenusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQyWebMenusRequest.ProtoReflect.Descriptor instead.
func (*CreateQyWebMenusRequest) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_web_v1_qy_web_menus_proto_rawDescGZIP(), []int{1}
}

type CreateQyWebMenusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateQyWebMenusReply) Reset() {
	*x = CreateQyWebMenusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQyWebMenusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQyWebMenusReply) ProtoMessage() {}

func (x *CreateQyWebMenusReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQyWebMenusReply.ProtoReflect.Descriptor instead.
func (*CreateQyWebMenusReply) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_web_v1_qy_web_menus_proto_rawDescGZIP(), []int{2}
}

type UpdateQyWebMenusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateQyWebMenusRequest) Reset() {
	*x = UpdateQyWebMenusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQyWebMenusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQyWebMenusRequest) ProtoMessage() {}

func (x *UpdateQyWebMenusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQyWebMenusRequest.ProtoReflect.Descriptor instead.
func (*UpdateQyWebMenusRequest) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_web_v1_qy_web_menus_proto_rawDescGZIP(), []int{3}
}

type UpdateQyWebMenusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateQyWebMenusReply) Reset() {
	*x = UpdateQyWebMenusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQyWebMenusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQyWebMenusReply) ProtoMessage() {}

func (x *UpdateQyWebMenusReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQyWebMenusReply.ProtoReflect.Descriptor instead.
func (*UpdateQyWebMenusReply) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_web_v1_qy_web_menus_proto_rawDescGZIP(), []int{4}
}

type DeleteQyWebMenusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteQyWebMenusRequest) Reset() {
	*x = DeleteQyWebMenusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteQyWebMenusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteQyWebMenusRequest) ProtoMessage() {}

func (x *DeleteQyWebMenusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteQyWebMenusRequest.ProtoReflect.Descriptor instead.
func (*DeleteQyWebMenusRequest) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_web_v1_qy_web_menus_proto_rawDescGZIP(), []int{5}
}

type DeleteQyWebMenusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteQyWebMenusReply) Reset() {
	*x = DeleteQyWebMenusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteQyWebMenusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteQyWebMenusReply) ProtoMessage() {}

func (x *DeleteQyWebMenusReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteQyWebMenusReply.ProtoReflect.Descriptor instead.
func (*DeleteQyWebMenusReply) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_web_v1_qy_web_menus_proto_rawDescGZIP(), []int{6}
}

type GetQyWebMenusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetQyWebMenusRequest) Reset() {
	*x = GetQyWebMenusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQyWebMenusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQyWebMenusRequest) ProtoMessage() {}

func (x *GetQyWebMenusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQyWebMenusRequest.ProtoReflect.Descriptor instead.
func (*GetQyWebMenusRequest) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_web_v1_qy_web_menus_proto_rawDescGZIP(), []int{7}
}

type GetQyWebMenusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetQyWebMenusReply) Reset() {
	*x = GetQyWebMenusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQyWebMenusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQyWebMenusReply) ProtoMessage() {}

func (x *GetQyWebMenusReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQyWebMenusReply.ProtoReflect.Descriptor instead.
func (*GetQyWebMenusReply) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_web_v1_qy_web_menus_proto_rawDescGZIP(), []int{8}
}

type ListQyWebMenusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListQyWebMenusRequest) Reset() {
	*x = ListQyWebMenusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQyWebMenusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQyWebMenusRequest) ProtoMessage() {}

func (x *ListQyWebMenusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQyWebMenusRequest.ProtoReflect.Descriptor instead.
func (*ListQyWebMenusRequest) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_web_v1_qy_web_menus_proto_rawDescGZIP(), []int{9}
}

type ListQyWebMenusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListQyWebMenusReply) Reset() {
	*x = ListQyWebMenusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQyWebMenusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQyWebMenusReply) ProtoMessage() {}

func (x *ListQyWebMenusReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQyWebMenusReply.ProtoReflect.Descriptor instead.
func (*ListQyWebMenusReply) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_web_v1_qy_web_menus_proto_rawDescGZIP(), []int{10}
}

var File_api_qycms_bff_web_v1_qy_web_menus_proto protoreflect.FileDescriptor

var file_api_qycms_bff_web_v1_qy_web_menus_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2f,
	0x77, 0x65, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x79, 0x5f, 0x77, 0x65, 0x62, 0x5f, 0x6d, 0x65,
	0x6e, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x71,
	0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01,
	0x0a, 0x14, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x57, 0x65, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x62, 0x6c, 0x61, 0x6e, 0x6b, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62,
	0x6c, 0x61, 0x6e, 0x6b, 0x65, 0x64, 0x12, 0x46, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72,
	0x65, 0x6e, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71,
	0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x6e, 0x75, 0x73, 0x57, 0x65, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0x19,
	0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x79, 0x57, 0x65, 0x62, 0x4d, 0x65, 0x6e,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x51, 0x79, 0x57, 0x65, 0x62, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x19, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x79, 0x57, 0x65,
	0x62, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x17, 0x0a,
	0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x79, 0x57, 0x65, 0x62, 0x4d, 0x65, 0x6e, 0x75,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x51, 0x79, 0x57, 0x65, 0x62, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x79, 0x57, 0x65, 0x62,
	0x4d, 0x65, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x51, 0x79, 0x57, 0x65, 0x62, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x51, 0x79, 0x57, 0x65, 0x62, 0x4d, 0x65,
	0x6e, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74,
	0x51, 0x79, 0x57, 0x65, 0x62, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x79, 0x57, 0x65, 0x62, 0x4d, 0x65,
	0x6e, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xe6, 0x05, 0x0a, 0x0a, 0x51, 0x79, 0x57,
	0x65, 0x62, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x12, 0x6e, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x51, 0x79, 0x57, 0x65, 0x62, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x12, 0x2d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x79, 0x57, 0x65, 0x62, 0x4d, 0x65,
	0x6e, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x79, 0x57, 0x65, 0x62, 0x4d, 0x65, 0x6e,
	0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x6e, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x51, 0x79, 0x57, 0x65, 0x62, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x12, 0x2d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x79, 0x57, 0x65, 0x62, 0x4d, 0x65,
	0x6e, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x79, 0x57, 0x65, 0x62, 0x4d, 0x65, 0x6e,
	0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x6e, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x51, 0x79, 0x57, 0x65, 0x62, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x12, 0x2d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x79, 0x57, 0x65, 0x62, 0x4d, 0x65,
	0x6e, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x79, 0x57, 0x65, 0x62, 0x4d, 0x65, 0x6e,
	0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x8d, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x51,
	0x79, 0x57, 0x65, 0x62, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x12,
	0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e,
	0x77, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x79, 0x57, 0x65, 0x62, 0x4d,
	0x65, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x79, 0x57, 0x65, 0x62, 0x4d, 0x65, 0x6e, 0x75, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x73,
	0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x8d, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x51,
	0x79, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x12,
	0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e,
	0x77, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x79, 0x57, 0x65, 0x62, 0x4d,
	0x65, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x79, 0x57, 0x65, 0x62, 0x4d, 0x65, 0x6e, 0x75, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x73,
	0x2f, 0x66, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x12, 0x68, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x51,
	0x79, 0x57, 0x65, 0x62, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x79, 0x57, 0x65, 0x62, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x79, 0x63,
	0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x51, 0x79, 0x57, 0x65, 0x62, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x42, 0x4e, 0x0a, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62,
	0x66, 0x66, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x77, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2f,
	0x71, 0x69, 0x6e, 0x67, 0x79, 0x75, 0x63, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x79,
	0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2f, 0x77, 0x65, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_qycms_bff_web_v1_qy_web_menus_proto_rawDescOnce sync.Once
	file_api_qycms_bff_web_v1_qy_web_menus_proto_rawDescData = file_api_qycms_bff_web_v1_qy_web_menus_proto_rawDesc
)

func file_api_qycms_bff_web_v1_qy_web_menus_proto_rawDescGZIP() []byte {
	file_api_qycms_bff_web_v1_qy_web_menus_proto_rawDescOnce.Do(func() {
		file_api_qycms_bff_web_v1_qy_web_menus_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_qycms_bff_web_v1_qy_web_menus_proto_rawDescData)
	})
	return file_api_qycms_bff_web_v1_qy_web_menus_proto_rawDescData
}

var file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_qycms_bff_web_v1_qy_web_menus_proto_goTypes = []interface{}{
	(*MenusWebInfoResponse)(nil),    // 0: api.qycms_bff.web.v1.MenusWebInfoResponse
	(*CreateQyWebMenusRequest)(nil), // 1: api.qycms_bff.web.v1.CreateQyWebMenusRequest
	(*CreateQyWebMenusReply)(nil),   // 2: api.qycms_bff.web.v1.CreateQyWebMenusReply
	(*UpdateQyWebMenusRequest)(nil), // 3: api.qycms_bff.web.v1.UpdateQyWebMenusRequest
	(*UpdateQyWebMenusReply)(nil),   // 4: api.qycms_bff.web.v1.UpdateQyWebMenusReply
	(*DeleteQyWebMenusRequest)(nil), // 5: api.qycms_bff.web.v1.DeleteQyWebMenusRequest
	(*DeleteQyWebMenusReply)(nil),   // 6: api.qycms_bff.web.v1.DeleteQyWebMenusReply
	(*GetQyWebMenusRequest)(nil),    // 7: api.qycms_bff.web.v1.GetQyWebMenusRequest
	(*GetQyWebMenusReply)(nil),      // 8: api.qycms_bff.web.v1.GetQyWebMenusReply
	(*ListQyWebMenusRequest)(nil),   // 9: api.qycms_bff.web.v1.ListQyWebMenusRequest
	(*ListQyWebMenusReply)(nil),     // 10: api.qycms_bff.web.v1.ListQyWebMenusReply
}
var file_api_qycms_bff_web_v1_qy_web_menus_proto_depIdxs = []int32{
	0,  // 0: api.qycms_bff.web.v1.MenusWebInfoResponse.children:type_name -> api.qycms_bff.web.v1.MenusWebInfoResponse
	1,  // 1: api.qycms_bff.web.v1.QyWebMenus.CreateQyWebMenus:input_type -> api.qycms_bff.web.v1.CreateQyWebMenusRequest
	3,  // 2: api.qycms_bff.web.v1.QyWebMenus.UpdateQyWebMenus:input_type -> api.qycms_bff.web.v1.UpdateQyWebMenusRequest
	5,  // 3: api.qycms_bff.web.v1.QyWebMenus.DeleteQyWebMenus:input_type -> api.qycms_bff.web.v1.DeleteQyWebMenusRequest
	7,  // 4: api.qycms_bff.web.v1.QyWebMenus.GetQyWebHeaderMenus:input_type -> api.qycms_bff.web.v1.GetQyWebMenusRequest
	7,  // 5: api.qycms_bff.web.v1.QyWebMenus.GetQyWebFooterMenus:input_type -> api.qycms_bff.web.v1.GetQyWebMenusRequest
	9,  // 6: api.qycms_bff.web.v1.QyWebMenus.ListQyWebMenus:input_type -> api.qycms_bff.web.v1.ListQyWebMenusRequest
	2,  // 7: api.qycms_bff.web.v1.QyWebMenus.CreateQyWebMenus:output_type -> api.qycms_bff.web.v1.CreateQyWebMenusReply
	4,  // 8: api.qycms_bff.web.v1.QyWebMenus.UpdateQyWebMenus:output_type -> api.qycms_bff.web.v1.UpdateQyWebMenusReply
	6,  // 9: api.qycms_bff.web.v1.QyWebMenus.DeleteQyWebMenus:output_type -> api.qycms_bff.web.v1.DeleteQyWebMenusReply
	8,  // 10: api.qycms_bff.web.v1.QyWebMenus.GetQyWebHeaderMenus:output_type -> api.qycms_bff.web.v1.GetQyWebMenusReply
	8,  // 11: api.qycms_bff.web.v1.QyWebMenus.GetQyWebFooterMenus:output_type -> api.qycms_bff.web.v1.GetQyWebMenusReply
	10, // 12: api.qycms_bff.web.v1.QyWebMenus.ListQyWebMenus:output_type -> api.qycms_bff.web.v1.ListQyWebMenusReply
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_qycms_bff_web_v1_qy_web_menus_proto_init() }
func file_api_qycms_bff_web_v1_qy_web_menus_proto_init() {
	if File_api_qycms_bff_web_v1_qy_web_menus_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenusWebInfoResponse); i {
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
		file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQyWebMenusRequest); i {
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
		file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQyWebMenusReply); i {
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
		file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQyWebMenusRequest); i {
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
		file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQyWebMenusReply); i {
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
		file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteQyWebMenusRequest); i {
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
		file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteQyWebMenusReply); i {
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
		file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQyWebMenusRequest); i {
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
		file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQyWebMenusReply); i {
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
		file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQyWebMenusRequest); i {
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
		file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQyWebMenusReply); i {
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
			RawDescriptor: file_api_qycms_bff_web_v1_qy_web_menus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_qycms_bff_web_v1_qy_web_menus_proto_goTypes,
		DependencyIndexes: file_api_qycms_bff_web_v1_qy_web_menus_proto_depIdxs,
		MessageInfos:      file_api_qycms_bff_web_v1_qy_web_menus_proto_msgTypes,
	}.Build()
	File_api_qycms_bff_web_v1_qy_web_menus_proto = out.File
	file_api_qycms_bff_web_v1_qy_web_menus_proto_rawDesc = nil
	file_api_qycms_bff_web_v1_qy_web_menus_proto_goTypes = nil
	file_api_qycms_bff_web_v1_qy_web_menus_proto_depIdxs = nil
}
