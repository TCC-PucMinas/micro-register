// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.2
// source: users.proto

package communicate

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

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Street       string `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	State        string `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Number       string `protobuf:"bytes,3,opt,name=number,proto3" json:"number,omitempty"`
	Country      string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Neighborhood string `protobuf:"bytes,5,opt,name=neighborhood,proto3" json:"neighborhood,omitempty"`
	Complement   string `protobuf:"bytes,6,opt,name=complement,proto3" json:"complement,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetNeighborhood() string {
	if x != nil {
		return x.Neighborhood
	}
	return ""
}

func (x *Address) GetComplement() string {
	if x != nil {
		return x.Complement
	}
	return ""
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string     `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string     `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Phone     string     `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	CpfCpnj   string     `protobuf:"bytes,4,opt,name=cpf_cpnj,json=cpfCpnj,proto3" json:"cpf_cpnj,omitempty"`
	Busines   string     `protobuf:"bytes,5,opt,name=busines,proto3" json:"busines,omitempty"`
	Password  string     `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	Email     string     `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	Address   []*Address `protobuf:"bytes,8,rep,name=address,proto3" json:"address,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateUserRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreateUserRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateUserRequest) GetCpfCpnj() string {
	if x != nil {
		return x.CpfCpnj
	}
	return ""
}

func (x *CreateUserRequest) GetBusines() string {
	if x != nil {
		return x.Busines
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetAddress() []*Address {
	if x != nil {
		return x.Address
	}
	return nil
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Created bool `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserResponse) GetCreated() bool {
	if x != nil {
		return x.Created
	}
	return false
}

type ValidateUserExistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email   string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	CpfCnpj string `protobuf:"bytes,2,opt,name=cpf_cnpj,json=cpfCnpj,proto3" json:"cpf_cnpj,omitempty"`
}

func (x *ValidateUserExistRequest) Reset() {
	*x = ValidateUserExistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateUserExistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateUserExistRequest) ProtoMessage() {}

func (x *ValidateUserExistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateUserExistRequest.ProtoReflect.Descriptor instead.
func (*ValidateUserExistRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{3}
}

func (x *ValidateUserExistRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ValidateUserExistRequest) GetCpfCnpj() string {
	if x != nil {
		return x.CpfCnpj
	}
	return ""
}

type ValidateUserExistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *ValidateUserExistResponse) Reset() {
	*x = ValidateUserExistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateUserExistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateUserExistResponse) ProtoMessage() {}

func (x *ValidateUserExistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateUserExistResponse.ProtoReflect.Descriptor instead.
func (*ValidateUserExistResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{4}
}

func (x *ValidateUserExistResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

type ValidateCodeActiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CodeActive string `protobuf:"bytes,1,opt,name=code_active,json=codeActive,proto3" json:"code_active,omitempty"`
}

func (x *ValidateCodeActiveRequest) Reset() {
	*x = ValidateCodeActiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateCodeActiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateCodeActiveRequest) ProtoMessage() {}

func (x *ValidateCodeActiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateCodeActiveRequest.ProtoReflect.Descriptor instead.
func (*ValidateCodeActiveRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{5}
}

func (x *ValidateCodeActiveRequest) GetCodeActive() string {
	if x != nil {
		return x.CodeActive
	}
	return ""
}

type ValidateCodeActiveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *ValidateCodeActiveResponse) Reset() {
	*x = ValidateCodeActiveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateCodeActiveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateCodeActiveResponse) ProtoMessage() {}

func (x *ValidateCodeActiveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateCodeActiveResponse.ProtoReflect.Descriptor instead.
func (*ValidateCodeActiveResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{6}
}

func (x *ValidateCodeActiveResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

type ActiveCodeUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CodeActive string `protobuf:"bytes,1,opt,name=code_active,json=codeActive,proto3" json:"code_active,omitempty"`
}

func (x *ActiveCodeUserRequest) Reset() {
	*x = ActiveCodeUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveCodeUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveCodeUserRequest) ProtoMessage() {}

func (x *ActiveCodeUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveCodeUserRequest.ProtoReflect.Descriptor instead.
func (*ActiveCodeUserRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{7}
}

func (x *ActiveCodeUserRequest) GetCodeActive() string {
	if x != nil {
		return x.CodeActive
	}
	return ""
}

type ActiveCodeUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Actived bool `protobuf:"varint,1,opt,name=actived,proto3" json:"actived,omitempty"`
}

func (x *ActiveCodeUserResponse) Reset() {
	*x = ActiveCodeUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveCodeUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveCodeUserResponse) ProtoMessage() {}

func (x *ActiveCodeUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveCodeUserResponse.ProtoReflect.Descriptor instead.
func (*ActiveCodeUserResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{8}
}

func (x *ActiveCodeUserResponse) GetActived() bool {
	if x != nil {
		return x.Actived
	}
	return false
}

var File_users_proto protoreflect.FileDescriptor

var file_users_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72,
	0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x65, 0x69,
	0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xf0, 0x01,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x70, 0x66, 0x5f, 0x63, 0x70, 0x6e,
	0x6a, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x70, 0x66, 0x43, 0x70, 0x6e, 0x6a,
	0x12, 0x18, 0x0a, 0x07, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x22, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x2e, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x22, 0x4b, 0x0a, 0x18, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x70, 0x66, 0x5f, 0x63, 0x6e, 0x70, 0x6a, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x70, 0x66, 0x43, 0x6e, 0x70, 0x6a, 0x22, 0x31, 0x0a,
	0x19, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x22, 0x3c, 0x0a, 0x19, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x32,
	0x0a, 0x1a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x22, 0x38, 0x0a, 0x15, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x6f, 0x64, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x32, 0x0a, 0x16,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x64,
	0x32, 0xb7, 0x02, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a,
	0x11, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x12, 0x19, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x12, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x12, 0x1a, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x17,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x16, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_users_proto_rawDescOnce sync.Once
	file_users_proto_rawDescData = file_users_proto_rawDesc
)

func file_users_proto_rawDescGZIP() []byte {
	file_users_proto_rawDescOnce.Do(func() {
		file_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_users_proto_rawDescData)
	})
	return file_users_proto_rawDescData
}

var file_users_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_users_proto_goTypes = []interface{}{
	(*Address)(nil),                    // 0: Address
	(*CreateUserRequest)(nil),          // 1: CreateUserRequest
	(*CreateUserResponse)(nil),         // 2: CreateUserResponse
	(*ValidateUserExistRequest)(nil),   // 3: ValidateUserExistRequest
	(*ValidateUserExistResponse)(nil),  // 4: ValidateUserExistResponse
	(*ValidateCodeActiveRequest)(nil),  // 5: ValidateCodeActiveRequest
	(*ValidateCodeActiveResponse)(nil), // 6: ValidateCodeActiveResponse
	(*ActiveCodeUserRequest)(nil),      // 7: ActiveCodeUserRequest
	(*ActiveCodeUserResponse)(nil),     // 8: ActiveCodeUserResponse
}
var file_users_proto_depIdxs = []int32{
	0, // 0: CreateUserRequest.address:type_name -> Address
	1, // 1: UserCommunicate.CreateUser:input_type -> CreateUserRequest
	3, // 2: UserCommunicate.ValidateUserExist:input_type -> ValidateUserExistRequest
	5, // 3: UserCommunicate.ValidateCodeActive:input_type -> ValidateCodeActiveRequest
	7, // 4: UserCommunicate.ActivatedUserCodeActive:input_type -> ActiveCodeUserRequest
	2, // 5: UserCommunicate.CreateUser:output_type -> CreateUserResponse
	4, // 6: UserCommunicate.ValidateUserExist:output_type -> ValidateUserExistResponse
	6, // 7: UserCommunicate.ValidateCodeActive:output_type -> ValidateCodeActiveResponse
	8, // 8: UserCommunicate.ActivatedUserCodeActive:output_type -> ActiveCodeUserResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_users_proto_init() }
func file_users_proto_init() {
	if File_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_users_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
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
		file_users_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateUserExistRequest); i {
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
		file_users_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateUserExistResponse); i {
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
		file_users_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateCodeActiveRequest); i {
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
		file_users_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateCodeActiveResponse); i {
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
		file_users_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveCodeUserRequest); i {
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
		file_users_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveCodeUserResponse); i {
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
			RawDescriptor: file_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_users_proto_goTypes,
		DependencyIndexes: file_users_proto_depIdxs,
		MessageInfos:      file_users_proto_msgTypes,
	}.Build()
	File_users_proto = out.File
	file_users_proto_rawDesc = nil
	file_users_proto_goTypes = nil
	file_users_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserCommunicateClient is the client API for UserCommunicate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserCommunicateClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	ValidateUserExist(ctx context.Context, in *ValidateUserExistRequest, opts ...grpc.CallOption) (*ValidateUserExistResponse, error)
	ValidateCodeActive(ctx context.Context, in *ValidateCodeActiveRequest, opts ...grpc.CallOption) (*ValidateCodeActiveResponse, error)
	ActivatedUserCodeActive(ctx context.Context, in *ActiveCodeUserRequest, opts ...grpc.CallOption) (*ActiveCodeUserResponse, error)
}

type userCommunicateClient struct {
	cc grpc.ClientConnInterface
}

func NewUserCommunicateClient(cc grpc.ClientConnInterface) UserCommunicateClient {
	return &userCommunicateClient{cc}
}

func (c *userCommunicateClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/UserCommunicate/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCommunicateClient) ValidateUserExist(ctx context.Context, in *ValidateUserExistRequest, opts ...grpc.CallOption) (*ValidateUserExistResponse, error) {
	out := new(ValidateUserExistResponse)
	err := c.cc.Invoke(ctx, "/UserCommunicate/ValidateUserExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCommunicateClient) ValidateCodeActive(ctx context.Context, in *ValidateCodeActiveRequest, opts ...grpc.CallOption) (*ValidateCodeActiveResponse, error) {
	out := new(ValidateCodeActiveResponse)
	err := c.cc.Invoke(ctx, "/UserCommunicate/ValidateCodeActive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCommunicateClient) ActivatedUserCodeActive(ctx context.Context, in *ActiveCodeUserRequest, opts ...grpc.CallOption) (*ActiveCodeUserResponse, error) {
	out := new(ActiveCodeUserResponse)
	err := c.cc.Invoke(ctx, "/UserCommunicate/ActivatedUserCodeActive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserCommunicateServer is the server API for UserCommunicate service.
type UserCommunicateServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	ValidateUserExist(context.Context, *ValidateUserExistRequest) (*ValidateUserExistResponse, error)
	ValidateCodeActive(context.Context, *ValidateCodeActiveRequest) (*ValidateCodeActiveResponse, error)
	ActivatedUserCodeActive(context.Context, *ActiveCodeUserRequest) (*ActiveCodeUserResponse, error)
}

// UnimplementedUserCommunicateServer can be embedded to have forward compatible implementations.
type UnimplementedUserCommunicateServer struct {
}

func (*UnimplementedUserCommunicateServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserCommunicateServer) ValidateUserExist(context.Context, *ValidateUserExistRequest) (*ValidateUserExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateUserExist not implemented")
}
func (*UnimplementedUserCommunicateServer) ValidateCodeActive(context.Context, *ValidateCodeActiveRequest) (*ValidateCodeActiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateCodeActive not implemented")
}
func (*UnimplementedUserCommunicateServer) ActivatedUserCodeActive(context.Context, *ActiveCodeUserRequest) (*ActiveCodeUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivatedUserCodeActive not implemented")
}

func RegisterUserCommunicateServer(s *grpc.Server, srv UserCommunicateServer) {
	s.RegisterService(&_UserCommunicate_serviceDesc, srv)
}

func _UserCommunicate_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCommunicateServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserCommunicate/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCommunicateServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCommunicate_ValidateUserExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateUserExistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCommunicateServer).ValidateUserExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserCommunicate/ValidateUserExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCommunicateServer).ValidateUserExist(ctx, req.(*ValidateUserExistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCommunicate_ValidateCodeActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateCodeActiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCommunicateServer).ValidateCodeActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserCommunicate/ValidateCodeActive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCommunicateServer).ValidateCodeActive(ctx, req.(*ValidateCodeActiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCommunicate_ActivatedUserCodeActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActiveCodeUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCommunicateServer).ActivatedUserCodeActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserCommunicate/ActivatedUserCodeActive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCommunicateServer).ActivatedUserCodeActive(ctx, req.(*ActiveCodeUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserCommunicate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "UserCommunicate",
	HandlerType: (*UserCommunicateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserCommunicate_CreateUser_Handler,
		},
		{
			MethodName: "ValidateUserExist",
			Handler:    _UserCommunicate_ValidateUserExist_Handler,
		},
		{
			MethodName: "ValidateCodeActive",
			Handler:    _UserCommunicate_ValidateCodeActive_Handler,
		},
		{
			MethodName: "ActivatedUserCodeActive",
			Handler:    _UserCommunicate_ActivatedUserCodeActive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}
