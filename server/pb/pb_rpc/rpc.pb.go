// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: rpc.proto

package pb_rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	pb_enum "server/pb/pb_enum"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DbAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UID           int64  `protobuf:"varint,1,opt,name=UID,proto3" json:"UID,omitempty"`
	Account       string `protobuf:"bytes,2,opt,name=Account,proto3" json:"Account,omitempty"`
	Password      string `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	CreateTime    int64  `protobuf:"varint,4,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	LastLoginTime int64  `protobuf:"varint,5,opt,name=LastLoginTime,proto3" json:"LastLoginTime,omitempty"`
	Token         string `protobuf:"bytes,6,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *DbAccount) Reset() {
	*x = DbAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DbAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DbAccount) ProtoMessage() {}

func (x *DbAccount) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DbAccount.ProtoReflect.Descriptor instead.
func (*DbAccount) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *DbAccount) GetUID() int64 {
	if x != nil {
		return x.UID
	}
	return 0
}

func (x *DbAccount) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *DbAccount) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *DbAccount) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *DbAccount) GetLastLoginTime() int64 {
	if x != nil {
		return x.LastLoginTime
	}
	return 0
}

func (x *DbAccount) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type DbUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UID      int64       `protobuf:"varint,1,opt,name=UID,proto3" json:"UID,omitempty"`
	NickName string      `protobuf:"bytes,2,opt,name=NickName,proto3" json:"NickName,omitempty"`
	Sex      pb_enum.Sex `protobuf:"varint,3,opt,name=Sex,proto3,enum=pb_enum.Sex" json:"Sex,omitempty"`
	Icon     string      `protobuf:"bytes,4,opt,name=Icon,proto3" json:"Icon,omitempty"`
	Gold     uint32      `protobuf:"varint,5,opt,name=Gold,proto3" json:"Gold,omitempty"`
	Diamond  uint32      `protobuf:"varint,6,opt,name=Diamond,proto3" json:"Diamond,omitempty"`
}

func (x *DbUser) Reset() {
	*x = DbUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DbUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DbUser) ProtoMessage() {}

func (x *DbUser) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DbUser.ProtoReflect.Descriptor instead.
func (*DbUser) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *DbUser) GetUID() int64 {
	if x != nil {
		return x.UID
	}
	return 0
}

func (x *DbUser) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *DbUser) GetSex() pb_enum.Sex {
	if x != nil {
		return x.Sex
	}
	return pb_enum.Sex_Unknow
}

func (x *DbUser) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *DbUser) GetGold() uint32 {
	if x != nil {
		return x.Gold
	}
	return 0
}

func (x *DbUser) GetDiamond() uint32 {
	if x != nil {
		return x.Diamond
	}
	return 0
}

type DbTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomType  pb_enum.RoomType `protobuf:"varint,1,opt,name=RoomType,proto3,enum=pb_enum.RoomType" json:"RoomType,omitempty"`
	TableName string           `protobuf:"bytes,2,opt,name=TableName,proto3" json:"TableName,omitempty"`
	Password  string           `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *DbTable) Reset() {
	*x = DbTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DbTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DbTable) ProtoMessage() {}

func (x *DbTable) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DbTable.ProtoReflect.Descriptor instead.
func (*DbTable) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *DbTable) GetRoomType() pb_enum.RoomType {
	if x != nil {
		return x.RoomType
	}
	return pb_enum.RoomType_Tetris
}

func (x *DbTable) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *DbTable) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_rpc_proto protoreflect.FileDescriptor

var file_rpc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x62, 0x5f,
	0x68, 0x74, 0x74, 0x70, 0x1a, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xaf, 0x01, 0x0a, 0x09, 0x44, 0x62, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x55, 0x49, 0x44,
	0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x4c, 0x61, 0x73, 0x74, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x4c,
	0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x98, 0x01, 0x0a, 0x06, 0x44, 0x62, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x55, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x03, 0x53,
	0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e,
	0x75, 0x6d, 0x2e, 0x53, 0x65, 0x78, 0x52, 0x03, 0x53, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x49,
	0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x47, 0x6f, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x47,
	0x6f, 0x6c, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x44, 0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x22, 0x72, 0x0a,
	0x07, 0x44, 0x62, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x5f,
	0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x52,
	0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x42, 0x12, 0x5a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x70,
	0x62, 0x5f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_proto_rawDescOnce sync.Once
	file_rpc_proto_rawDescData = file_rpc_proto_rawDesc
)

func file_rpc_proto_rawDescGZIP() []byte {
	file_rpc_proto_rawDescOnce.Do(func() {
		file_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_proto_rawDescData)
	})
	return file_rpc_proto_rawDescData
}

var file_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rpc_proto_goTypes = []interface{}{
	(*DbAccount)(nil),     // 0: pb_http.DbAccount
	(*DbUser)(nil),        // 1: pb_http.DbUser
	(*DbTable)(nil),       // 2: pb_http.DbTable
	(pb_enum.Sex)(0),      // 3: pb_enum.Sex
	(pb_enum.RoomType)(0), // 4: pb_enum.RoomType
}
var file_rpc_proto_depIdxs = []int32{
	3, // 0: pb_http.DbUser.Sex:type_name -> pb_enum.Sex
	4, // 1: pb_http.DbTable.RoomType:type_name -> pb_enum.RoomType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_proto_init() }
func file_rpc_proto_init() {
	if File_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DbAccount); i {
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
		file_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DbUser); i {
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
		file_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DbTable); i {
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
			RawDescriptor: file_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_proto_goTypes,
		DependencyIndexes: file_rpc_proto_depIdxs,
		MessageInfos:      file_rpc_proto_msgTypes,
	}.Build()
	File_rpc_proto = out.File
	file_rpc_proto_rawDesc = nil
	file_rpc_proto_goTypes = nil
	file_rpc_proto_depIdxs = nil
}
