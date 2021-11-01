// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: hostGroupCommand.proto

package nowin

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type HostGroupCommandCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostGroupName    string `protobuf:"bytes,1,opt,name=hostGroupName,proto3" json:"hostGroupName,omitempty"`
	HomePageURL      string `protobuf:"bytes,2,opt,name=homePageURL,proto3" json:"homePageURL,omitempty"`
	Address          string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	IsEnterprise     string `protobuf:"bytes,4,opt,name=is_enterprise,json=isEnterprise,proto3" json:"is_enterprise,omitempty"`
	TelephoneNumber1 string `protobuf:"bytes,5,opt,name=telephoneNumber1,proto3" json:"telephoneNumber1,omitempty"`
	TelephoneNumber2 string `protobuf:"bytes,6,opt,name=telephoneNumber2,proto3" json:"telephoneNumber2,omitempty"`
	TelephoneNumber3 string `protobuf:"bytes,7,opt,name=telephoneNumber3,proto3" json:"telephoneNumber3,omitempty"`
	Email            string `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *HostGroupCommandCreateRequest) Reset() {
	*x = HostGroupCommandCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostGroupCommand_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostGroupCommandCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostGroupCommandCreateRequest) ProtoMessage() {}

func (x *HostGroupCommandCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hostGroupCommand_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostGroupCommandCreateRequest.ProtoReflect.Descriptor instead.
func (*HostGroupCommandCreateRequest) Descriptor() ([]byte, []int) {
	return file_hostGroupCommand_proto_rawDescGZIP(), []int{0}
}

func (x *HostGroupCommandCreateRequest) GetHostGroupName() string {
	if x != nil {
		return x.HostGroupName
	}
	return ""
}

func (x *HostGroupCommandCreateRequest) GetHomePageURL() string {
	if x != nil {
		return x.HomePageURL
	}
	return ""
}

func (x *HostGroupCommandCreateRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *HostGroupCommandCreateRequest) GetIsEnterprise() string {
	if x != nil {
		return x.IsEnterprise
	}
	return ""
}

func (x *HostGroupCommandCreateRequest) GetTelephoneNumber1() string {
	if x != nil {
		return x.TelephoneNumber1
	}
	return ""
}

func (x *HostGroupCommandCreateRequest) GetTelephoneNumber2() string {
	if x != nil {
		return x.TelephoneNumber2
	}
	return ""
}

func (x *HostGroupCommandCreateRequest) GetTelephoneNumber3() string {
	if x != nil {
		return x.TelephoneNumber3
	}
	return ""
}

func (x *HostGroupCommandCreateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type HostGroupCommandCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostGroupUUID string `protobuf:"bytes,1,opt,name=hostGroupUUID,proto3" json:"hostGroupUUID,omitempty"`
}

func (x *HostGroupCommandCreateResponse) Reset() {
	*x = HostGroupCommandCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostGroupCommand_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostGroupCommandCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostGroupCommandCreateResponse) ProtoMessage() {}

func (x *HostGroupCommandCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hostGroupCommand_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostGroupCommandCreateResponse.ProtoReflect.Descriptor instead.
func (*HostGroupCommandCreateResponse) Descriptor() ([]byte, []int) {
	return file_hostGroupCommand_proto_rawDescGZIP(), []int{1}
}

func (x *HostGroupCommandCreateResponse) GetHostGroupUUID() string {
	if x != nil {
		return x.HostGroupUUID
	}
	return ""
}

type HostGroupCommandUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostGroupUUID    string `protobuf:"bytes,1,opt,name=hostGroupUUID,proto3" json:"hostGroupUUID,omitempty"`
	HostGroupName    string `protobuf:"bytes,2,opt,name=hostGroupName,proto3" json:"hostGroupName,omitempty"`
	HomePageURL      string `protobuf:"bytes,3,opt,name=homePageURL,proto3" json:"homePageURL,omitempty"`
	Address          string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	IsEnterprise     string `protobuf:"bytes,5,opt,name=is_enterprise,json=isEnterprise,proto3" json:"is_enterprise,omitempty"`
	TelephoneNumber1 string `protobuf:"bytes,6,opt,name=telephoneNumber1,proto3" json:"telephoneNumber1,omitempty"`
	TelephoneNumber2 string `protobuf:"bytes,7,opt,name=telephoneNumber2,proto3" json:"telephoneNumber2,omitempty"`
	TelephoneNumber3 string `protobuf:"bytes,8,opt,name=telephoneNumber3,proto3" json:"telephoneNumber3,omitempty"`
	Email            string `protobuf:"bytes,9,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *HostGroupCommandUpdateRequest) Reset() {
	*x = HostGroupCommandUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostGroupCommand_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostGroupCommandUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostGroupCommandUpdateRequest) ProtoMessage() {}

func (x *HostGroupCommandUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hostGroupCommand_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostGroupCommandUpdateRequest.ProtoReflect.Descriptor instead.
func (*HostGroupCommandUpdateRequest) Descriptor() ([]byte, []int) {
	return file_hostGroupCommand_proto_rawDescGZIP(), []int{2}
}

func (x *HostGroupCommandUpdateRequest) GetHostGroupUUID() string {
	if x != nil {
		return x.HostGroupUUID
	}
	return ""
}

func (x *HostGroupCommandUpdateRequest) GetHostGroupName() string {
	if x != nil {
		return x.HostGroupName
	}
	return ""
}

func (x *HostGroupCommandUpdateRequest) GetHomePageURL() string {
	if x != nil {
		return x.HomePageURL
	}
	return ""
}

func (x *HostGroupCommandUpdateRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *HostGroupCommandUpdateRequest) GetIsEnterprise() string {
	if x != nil {
		return x.IsEnterprise
	}
	return ""
}

func (x *HostGroupCommandUpdateRequest) GetTelephoneNumber1() string {
	if x != nil {
		return x.TelephoneNumber1
	}
	return ""
}

func (x *HostGroupCommandUpdateRequest) GetTelephoneNumber2() string {
	if x != nil {
		return x.TelephoneNumber2
	}
	return ""
}

func (x *HostGroupCommandUpdateRequest) GetTelephoneNumber3() string {
	if x != nil {
		return x.TelephoneNumber3
	}
	return ""
}

func (x *HostGroupCommandUpdateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type HostGroupCommandUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HostGroupCommandUpdateResponse) Reset() {
	*x = HostGroupCommandUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostGroupCommand_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostGroupCommandUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostGroupCommandUpdateResponse) ProtoMessage() {}

func (x *HostGroupCommandUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hostGroupCommand_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostGroupCommandUpdateResponse.ProtoReflect.Descriptor instead.
func (*HostGroupCommandUpdateResponse) Descriptor() ([]byte, []int) {
	return file_hostGroupCommand_proto_rawDescGZIP(), []int{3}
}

type HostGroupCommandDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostGroupUUID string `protobuf:"bytes,1,opt,name=hostGroupUUID,proto3" json:"hostGroupUUID,omitempty"`
}

func (x *HostGroupCommandDeleteRequest) Reset() {
	*x = HostGroupCommandDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostGroupCommand_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostGroupCommandDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostGroupCommandDeleteRequest) ProtoMessage() {}

func (x *HostGroupCommandDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hostGroupCommand_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostGroupCommandDeleteRequest.ProtoReflect.Descriptor instead.
func (*HostGroupCommandDeleteRequest) Descriptor() ([]byte, []int) {
	return file_hostGroupCommand_proto_rawDescGZIP(), []int{4}
}

func (x *HostGroupCommandDeleteRequest) GetHostGroupUUID() string {
	if x != nil {
		return x.HostGroupUUID
	}
	return ""
}

type HostGroupCommandDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HostGroupCommandDeleteResponse) Reset() {
	*x = HostGroupCommandDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostGroupCommand_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostGroupCommandDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostGroupCommandDeleteResponse) ProtoMessage() {}

func (x *HostGroupCommandDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hostGroupCommand_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostGroupCommandDeleteResponse.ProtoReflect.Descriptor instead.
func (*HostGroupCommandDeleteResponse) Descriptor() ([]byte, []int) {
	return file_hostGroupCommand_proto_rawDescGZIP(), []int{5}
}

var File_hostGroupCommand_proto protoreflect.FileDescriptor

var file_hostGroupCommand_proto_rawDesc = []byte{
	0x0a, 0x16, 0x68, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6e, 0x6f, 0x77, 0x69, 0x6e, 0x22,
	0xc0, 0x02, 0x0a, 0x1d, 0x48, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x24, 0x0a, 0x0d, 0x68, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x68, 0x6f, 0x73, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x68, 0x6f, 0x6d, 0x65, 0x50,
	0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x6f,
	0x6d, 0x65, 0x50, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70,
	0x72, 0x69, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x73, 0x45, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x31, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x31, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32,
	0x12, 0x2a, 0x0a, 0x10, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x33, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x33, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x46, 0x0a, 0x1e, 0x48, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x68, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x68, 0x6f, 0x73,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x55, 0x49, 0x44, 0x22, 0xe6, 0x02, 0x0a, 0x1d, 0x48,
	0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d,
	0x68, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x68, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x55,
	0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x68, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x68, 0x6f, 0x73, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x68, 0x6f, 0x6d, 0x65,
	0x50, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68,
	0x6f, 0x6d, 0x65, 0x50, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x70, 0x72, 0x69, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x73, 0x45,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x31, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x31, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x32, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x33, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x33, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x22, 0x20, 0x0a, 0x1e, 0x48, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x45, 0x0a, 0x1d, 0x48, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x68, 0x6f, 0x73, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x68,
	0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x55, 0x49, 0x44, 0x22, 0x20, 0x0a, 0x1e,
	0x48, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa4,
	0x02, 0x0a, 0x17, 0x48, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x6e, 0x6f, 0x77, 0x69, 0x6e, 0x2e, 0x48, 0x6f, 0x73,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6e, 0x6f, 0x77,
	0x69, 0x6e, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x24, 0x2e,
	0x6e, 0x6f, 0x77, 0x69, 0x6e, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6e, 0x6f, 0x77, 0x69, 0x6e, 0x2e, 0x48, 0x6f, 0x73, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x6e, 0x6f, 0x77, 0x69, 0x6e, 0x2e, 0x48,
	0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6e,
	0x6f, 0x77, 0x69, 0x6e, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x66, 0x74, 0x69, 0x61, 0x2d, 0x69, 0x6e, 0x63, 0x2f, 0x6e,
	0x6f, 0x77, 0x69, 0x6e, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3b,
	0x6e, 0x6f, 0x77, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hostGroupCommand_proto_rawDescOnce sync.Once
	file_hostGroupCommand_proto_rawDescData = file_hostGroupCommand_proto_rawDesc
)

func file_hostGroupCommand_proto_rawDescGZIP() []byte {
	file_hostGroupCommand_proto_rawDescOnce.Do(func() {
		file_hostGroupCommand_proto_rawDescData = protoimpl.X.CompressGZIP(file_hostGroupCommand_proto_rawDescData)
	})
	return file_hostGroupCommand_proto_rawDescData
}

var file_hostGroupCommand_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_hostGroupCommand_proto_goTypes = []interface{}{
	(*HostGroupCommandCreateRequest)(nil),  // 0: nowin.HostGroupCommandCreateRequest
	(*HostGroupCommandCreateResponse)(nil), // 1: nowin.HostGroupCommandCreateResponse
	(*HostGroupCommandUpdateRequest)(nil),  // 2: nowin.HostGroupCommandUpdateRequest
	(*HostGroupCommandUpdateResponse)(nil), // 3: nowin.HostGroupCommandUpdateResponse
	(*HostGroupCommandDeleteRequest)(nil),  // 4: nowin.HostGroupCommandDeleteRequest
	(*HostGroupCommandDeleteResponse)(nil), // 5: nowin.HostGroupCommandDeleteResponse
}
var file_hostGroupCommand_proto_depIdxs = []int32{
	0, // 0: nowin.HostGroupCommandService.Create:input_type -> nowin.HostGroupCommandCreateRequest
	2, // 1: nowin.HostGroupCommandService.Update:input_type -> nowin.HostGroupCommandUpdateRequest
	4, // 2: nowin.HostGroupCommandService.Delete:input_type -> nowin.HostGroupCommandDeleteRequest
	1, // 3: nowin.HostGroupCommandService.Create:output_type -> nowin.HostGroupCommandCreateResponse
	3, // 4: nowin.HostGroupCommandService.Update:output_type -> nowin.HostGroupCommandUpdateResponse
	5, // 5: nowin.HostGroupCommandService.Delete:output_type -> nowin.HostGroupCommandDeleteResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hostGroupCommand_proto_init() }
func file_hostGroupCommand_proto_init() {
	if File_hostGroupCommand_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hostGroupCommand_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostGroupCommandCreateRequest); i {
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
		file_hostGroupCommand_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostGroupCommandCreateResponse); i {
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
		file_hostGroupCommand_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostGroupCommandUpdateRequest); i {
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
		file_hostGroupCommand_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostGroupCommandUpdateResponse); i {
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
		file_hostGroupCommand_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostGroupCommandDeleteRequest); i {
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
		file_hostGroupCommand_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostGroupCommandDeleteResponse); i {
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
			RawDescriptor: file_hostGroupCommand_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hostGroupCommand_proto_goTypes,
		DependencyIndexes: file_hostGroupCommand_proto_depIdxs,
		MessageInfos:      file_hostGroupCommand_proto_msgTypes,
	}.Build()
	File_hostGroupCommand_proto = out.File
	file_hostGroupCommand_proto_rawDesc = nil
	file_hostGroupCommand_proto_goTypes = nil
	file_hostGroupCommand_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HostGroupCommandServiceClient is the client API for HostGroupCommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HostGroupCommandServiceClient interface {
	Create(ctx context.Context, in *HostGroupCommandCreateRequest, opts ...grpc.CallOption) (*HostGroupCommandCreateResponse, error)
	Update(ctx context.Context, in *HostGroupCommandUpdateRequest, opts ...grpc.CallOption) (*HostGroupCommandUpdateResponse, error)
	Delete(ctx context.Context, in *HostGroupCommandDeleteRequest, opts ...grpc.CallOption) (*HostGroupCommandDeleteResponse, error)
}

type hostGroupCommandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHostGroupCommandServiceClient(cc grpc.ClientConnInterface) HostGroupCommandServiceClient {
	return &hostGroupCommandServiceClient{cc}
}

func (c *hostGroupCommandServiceClient) Create(ctx context.Context, in *HostGroupCommandCreateRequest, opts ...grpc.CallOption) (*HostGroupCommandCreateResponse, error) {
	out := new(HostGroupCommandCreateResponse)
	err := c.cc.Invoke(ctx, "/nowin.HostGroupCommandService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostGroupCommandServiceClient) Update(ctx context.Context, in *HostGroupCommandUpdateRequest, opts ...grpc.CallOption) (*HostGroupCommandUpdateResponse, error) {
	out := new(HostGroupCommandUpdateResponse)
	err := c.cc.Invoke(ctx, "/nowin.HostGroupCommandService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostGroupCommandServiceClient) Delete(ctx context.Context, in *HostGroupCommandDeleteRequest, opts ...grpc.CallOption) (*HostGroupCommandDeleteResponse, error) {
	out := new(HostGroupCommandDeleteResponse)
	err := c.cc.Invoke(ctx, "/nowin.HostGroupCommandService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostGroupCommandServiceServer is the server API for HostGroupCommandService service.
type HostGroupCommandServiceServer interface {
	Create(context.Context, *HostGroupCommandCreateRequest) (*HostGroupCommandCreateResponse, error)
	Update(context.Context, *HostGroupCommandUpdateRequest) (*HostGroupCommandUpdateResponse, error)
	Delete(context.Context, *HostGroupCommandDeleteRequest) (*HostGroupCommandDeleteResponse, error)
}

// UnimplementedHostGroupCommandServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHostGroupCommandServiceServer struct {
}

func (*UnimplementedHostGroupCommandServiceServer) Create(context.Context, *HostGroupCommandCreateRequest) (*HostGroupCommandCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedHostGroupCommandServiceServer) Update(context.Context, *HostGroupCommandUpdateRequest) (*HostGroupCommandUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedHostGroupCommandServiceServer) Delete(context.Context, *HostGroupCommandDeleteRequest) (*HostGroupCommandDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterHostGroupCommandServiceServer(s *grpc.Server, srv HostGroupCommandServiceServer) {
	s.RegisterService(&_HostGroupCommandService_serviceDesc, srv)
}

func _HostGroupCommandService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostGroupCommandCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostGroupCommandServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nowin.HostGroupCommandService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostGroupCommandServiceServer).Create(ctx, req.(*HostGroupCommandCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostGroupCommandService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostGroupCommandUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostGroupCommandServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nowin.HostGroupCommandService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostGroupCommandServiceServer).Update(ctx, req.(*HostGroupCommandUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostGroupCommandService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostGroupCommandDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostGroupCommandServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nowin.HostGroupCommandService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostGroupCommandServiceServer).Delete(ctx, req.(*HostGroupCommandDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HostGroupCommandService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nowin.HostGroupCommandService",
	HandlerType: (*HostGroupCommandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _HostGroupCommandService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _HostGroupCommandService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _HostGroupCommandService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hostGroupCommand.proto",
}
