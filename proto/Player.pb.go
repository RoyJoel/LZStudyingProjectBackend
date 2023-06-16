package proto

import (
	context "context"
	reflect "reflect"
	sync "sync"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 定义客户端请求的数据格式
// message 对应  生成的代码中的struct
type PlayerInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// [修饰符] 类型 字段名 = 标识符
	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	LoginName string `protobuf:"bytes,2,opt,name=loginName,proto3" json:"loginName,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Icon      string `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
	Sex       string `protobuf:"bytes,5,opt,name=sex,proto3" json:"sex,omitempty"`
	Age       int64  `protobuf:"varint,6,opt,name=age,proto3" json:"age,omitempty"`
	Points    int64  `protobuf:"bytes,12,opt,name=points,proto3" json:"points,omitempty"`
}

func (x *PlayerInfoRequest) Reset() {
	*x = PlayerInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInfoRequest) ProtoMessage() {}

func (x *PlayerInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Player_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInfoRequest.ProtoReflect.Descriptor instead.
func (*PlayerInfoRequest) Descriptor() ([]byte, []int) {
	return file_Player_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerInfoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PlayerInfoRequest) GetLoginName() string {
	if x != nil {
		return x.LoginName
	}
	return ""
}

func (x *PlayerInfoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlayerInfoRequest) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *PlayerInfoRequest) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *PlayerInfoRequest) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *PlayerInfoRequest) GetPoints() int64 {
	if x != nil {
		return x.Points
	}
	return 0
}

type PlayerInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg   string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Count int64  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Data  string `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PlayerInfoResponse) Reset() {
	*x = PlayerInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Player_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInfoResponse) ProtoMessage() {}

func (x *PlayerInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Player_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInfoResponse.ProtoReflect.Descriptor instead.
func (*PlayerInfoResponse) Descriptor() ([]byte, []int) {
	return file_Player_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerInfoResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PlayerInfoResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *PlayerInfoResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *PlayerInfoResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_Player_proto protoreflect.FileDescriptor

var file_Player_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x0b, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x6e, 0x66, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69,
	0x6e, 0x66, 0x6f, 0x4b, 0x65, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x6e, 0x66, 0x6f, 0x4e, 0x75,
	0x6d, 0x22, 0x5e, 0x0a, 0x0c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x32, 0xef, 0x02, 0x0a, 0x0e, 0x4e, 0x75, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x49, 0x6e,
	0x66, 0x6f, 0x42, 0x79, 0x49, 0x64, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x38, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x4e, 0x75, 0x6d, 0x42, 0x79, 0x4b, 0x65, 0x79,
	0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x10, 0x46,
	0x69, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x12,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0b, 0x53, 0x61,
	0x76, 0x65, 0x4e, 0x75, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a,
	0x07, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Player_proto_rawDescOnce sync.Once
	file_Player_proto_rawDescData = file_Player_proto_rawDesc
)

func file_Player_proto_rawDescGZIP() []byte {
	file_Player_proto_rawDescOnce.Do(func() {
		file_Player_proto_rawDescData = protoimpl.X.CompressGZIP(file_Player_proto_rawDescData)
	})
	return file_Player_proto_rawDescData
}

var file_Player_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Player_proto_goTypes = []interface{}{
	(*PlayerInfoRequest)(nil),  // 0: proto.PlayerInfoRequest
	(*PlayerInfoResponse)(nil), // 1: proto.PlayerInfoResponse
}
var file_Player_proto_depIdxs = []int32{
	0, // 0: proto.PlayerInfoService.GetPlayerInfoById:input_type -> proto.PlayerInfoRequest
	0, // 1: proto.PlayerInfoService.AddNumByKey:input_type -> proto.PlayerInfoRequest
	0, // 2: proto.PlayerInfoService.FindPlayerInfoByKey:input_type -> proto.PlayerInfoRequest
	0, // 3: proto.PlayerInfoService.SavePlayerInfo:input_type -> proto.PlayerInfoRequest
	0, // 4: proto.PlayerInfoService.DeleteById:input_type -> proto.PlayerInfoRequest
	0, // 5: proto.PlayerInfoService.FindAllPlayers:input_type -> proto.PlayerInfoRequest
	1, // 6: proto.PlayerInfoService.GetPlayerInfoById:output_type -> proto.PlayerInfoResponse
	1, // 7: proto.PlayerInfoService.AddNumByKey:output_type -> proto.PlayerInfoResponse
	1, // 8: proto.PlayerInfoService.FindPlayerInfoByKey:output_type -> proto.PlayerInfoResponse
	1, // 9: proto.PlayerInfoService.SavePlayerInfo:output_type -> proto.PlayerInfoResponse
	1, // 10: proto.PlayerInfoService.DeleteById:output_type -> proto.PlayerInfoResponse
	1, // 11: proto.PlayerInfoService.FindAllPlayers:output_type -> proto.PlayerInfoResponse
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Player_proto_init() }
func file_Player_proto_init() {
	if File_Player_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerInfoRequest); i {
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
		file_Player_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerInfoResponse); i {
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
			RawDescriptor: file_Player_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Player_proto_goTypes,
		DependencyIndexes: file_Player_proto_depIdxs,
		MessageInfos:      file_Player_proto_msgTypes,
	}.Build()
	File_Player_proto = out.File
	file_Player_proto_rawDesc = nil
	file_Player_proto_goTypes = nil
	file_Player_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PlayerInfoServiceClient is the client API for PlayerInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlayerInfoServiceClient interface {
	//rpc 方式；定义GetUserInfo 远程调用
	AddPlayer(ctx context.Context, in *PlayerInfoRequest, opts ...grpc.CallOption) (*PlayerInfoResponse, error)
	UpdatePlayer(ctx context.Context, in *PlayerInfoRequest, opts ...grpc.CallOption) (*PlayerInfoResponse, error)
	GetPlayerInfo(ctx context.Context, in *PlayerInfoRequest, opts ...grpc.CallOption) (*PlayerInfoResponse, error)
	SearchPlayer(ctx context.Context, in *PlayerInfoRequest, opts ...grpc.CallOption) (*PlayerInfoResponse, error)
	AddFriend(ctx context.Context, in *PlayerInfoRequest, opts ...grpc.CallOption) (*PlayerInfoResponse, error)
	DeleteFriend(ctx context.Context, in *PlayerInfoRequest, opts ...grpc.CallOption) (*PlayerInfoResponse, error)
}

type playerInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlayerInfoServiceClient(cc grpc.ClientConnInterface) PlayerInfoServiceClient {
	return &playerInfoServiceClient{cc}
}

func (c *playerInfoServiceClient) AddPlayer(ctx context.Context, in *PlayerInfoRequest, opts ...grpc.CallOption) (*PlayerInfoResponse, error) {
	out := new(PlayerInfoResponse)
	err := c.cc.Invoke(ctx, "/proto.PlayerInfoService/AddPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerInfoServiceClient) UpdatePlayer(ctx context.Context, in *PlayerInfoRequest, opts ...grpc.CallOption) (*PlayerInfoResponse, error) {
	out := new(PlayerInfoResponse)
	err := c.cc.Invoke(ctx, "/proto.PlayerInfoService/UpdatePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerInfoServiceClient) GetPlayerInfo(ctx context.Context, in *PlayerInfoRequest, opts ...grpc.CallOption) (*PlayerInfoResponse, error) {
	out := new(PlayerInfoResponse)
	err := c.cc.Invoke(ctx, "/proto.PlayerInfoService/GetPlayerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerInfoServiceClient) SearchPlayer(ctx context.Context, in *PlayerInfoRequest, opts ...grpc.CallOption) (*PlayerInfoResponse, error) {
	out := new(PlayerInfoResponse)
	err := c.cc.Invoke(ctx, "/proto.PlayerInfoService/SearchPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerInfoServiceClient) AddFriend(ctx context.Context, in *PlayerInfoRequest, opts ...grpc.CallOption) (*PlayerInfoResponse, error) {
	out := new(PlayerInfoResponse)
	err := c.cc.Invoke(ctx, "/proto.PlayerInfoService/AddFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerInfoServiceClient) DeleteFriend(ctx context.Context, in *PlayerInfoRequest, opts ...grpc.CallOption) (*PlayerInfoResponse, error) {
	out := new(PlayerInfoResponse)
	err := c.cc.Invoke(ctx, "/proto.PlayerInfoService/DeleteFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// func (c *PlayerInfoServiceClient) FindPlayerInfoByKey(ctx context.Context, in *PlayerInfoRequest, opts ...grpc.CallOption) (*PlayerInfoResponse, error) {
// 	out := new(PlayerInfoResponse)
// 	err := c.cc.Invoke(ctx, "/proto.PlayerInfoService/FindPlayerInfoByKey", in, out, opts...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return out, nil
// }

// func (c *PlayerInfoServiceClient) DeleteById(ctx context.Context, in *PlayerInfoRequest, opts ...grpc.CallOption) (*PlayerInfoResponse, error) {
// 	out := new(PlayerInfoResponse)
// 	err := c.cc.Invoke(ctx, "/proto.PlayerInfoService/DeleteById", in, out, opts...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return out, nil
// }

// func (c *playerInfoServiceClient) FindAllPlayers(ctx context.Context, in *PlayerInfoRequest, opts ...grpc.CallOption) (*PlayerInfoResponse, error) {
// 	out := new(PlayerInfoResponse)
// 	err := c.cc.Invoke(ctx, "/proto.PlayerInfoService/FindAllPlayers", in, out, opts...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return out, nil
// }

// PlayerInfoServiceServer is the server API for PlayerInfoService service.
type PlayerInfoServiceServer interface {
	//rpc 方式；定义GetUserInfo 远程调用
	AddPlayer(context.Context, *PlayerInfoRequest) (*PlayerInfoResponse, error)
	UpdatePlayer(context.Context, *PlayerInfoRequest) (*PlayerInfoResponse, error)
	GetPlayerInfo(context.Context, *PlayerInfoRequest) (*PlayerInfoResponse, error)
	SearchPlayer(context.Context, *PlayerInfoRequest) (*PlayerInfoResponse, error)
	AddFriend(context.Context, *PlayerInfoRequest) (*PlayerInfoResponse, error)
	DeleteFriend(context.Context, *PlayerInfoRequest) (*PlayerInfoResponse, error)
}

// UnimplementedPlayerInfoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPlayerInfoServiceServer struct {
}

func (*UnimplementedPlayerInfoServiceServer) GetPlayerInfo(context.Context, *PlayerInfoRequest) (*PlayerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerInfo not implemented")
}
func (*UnimplementedPlayerInfoServiceServer) AddPlayer(context.Context, *PlayerInfoRequest) (*PlayerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPlayer not implemented")
}

//	func (*UnimplementedPlayerInfoServiceServer) FindPlayerInfoByKey(context.Context, *PlayerInfoRequest) (*PlayerInfoResponse, error) {
//		return nil, status.Errorf(codes.Unimplemented, "method FindPlayerInfoByKey not implemented")
//	}
func (*UnimplementedPlayerInfoServiceServer) UpdatePlayer(context.Context, *PlayerInfoRequest) (*PlayerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlayer not implemented")
}

//	func (*UnimplementedPlayerInfoServiceServer) DeleteById(context.Context, *PlayerInfoRequest) (*PlayerInfoResponse, error) {
//		return nil, status.Errorf(codes.Unimplemented, "method DeleteById not implemented")
//	}
func (*UnimplementedPlayerInfoServiceServer) SearchPlayer(context.Context, *PlayerInfoRequest) (*PlayerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPlayer not implemented")
}

func (*UnimplementedPlayerInfoServiceServer) AddFriend(context.Context, *PlayerInfoRequest) (*PlayerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}

func (*UnimplementedPlayerInfoServiceServer) DeleteFriend(context.Context, *PlayerInfoRequest) (*PlayerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFriend not implemented")
}

func RegisterPlayerInfoServiceServer(s *grpc.Server, srv PlayerInfoServiceServer) {
	s.RegisterService(&_PlayerInfoService_serviceDesc, srv)
}

func _PlayerInfoService_GetPlayerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerInfoServiceServer).GetPlayerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PlayerInfoService/GetPlayerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerInfoServiceServer).GetPlayerInfo(ctx, req.(*PlayerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerInfoService_AddPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerInfoServiceServer).AddPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PlayerInfoService/AddPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerInfoServiceServer).AddPlayer(ctx, req.(*PlayerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerInfoService_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerInfoServiceServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PlayerInfoService/AddFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerInfoServiceServer).AddFriend(ctx, req.(*PlayerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerInfoService_UpdatePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerInfoServiceServer).UpdatePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PlayerInfoService/UpdatePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerInfoServiceServer).UpdatePlayer(ctx, req.(*PlayerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerInfoService_DeleteFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerInfoServiceServer).DeleteFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PlayerInfoService/DeleteFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerInfoServiceServer).DeleteFriend(ctx, req.(*PlayerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerInfoService_SearchPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerInfoServiceServer).SearchPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PlayerInfoService/SearchPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerInfoServiceServer).SearchPlayer(ctx, req.(*PlayerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlayerInfoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PlayerInfoService",
	HandlerType: (*PlayerInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlayerInfo",
			Handler:    _PlayerInfoService_GetPlayerInfo_Handler,
		},
		{
			MethodName: "AddFriend",
			Handler:    _PlayerInfoService_AddFriend_Handler,
		},
		{
			MethodName: "AddPlayer",
			Handler:    _PlayerInfoService_AddPlayer_Handler,
		},
		{
			MethodName: "UpdatePlayer",
			Handler:    _PlayerInfoService_UpdatePlayer_Handler,
		},
		{
			MethodName: "DeleteFriend",
			Handler:    _PlayerInfoService_DeleteFriend_Handler,
		},
		{
			MethodName: "SearchPlayer",
			Handler:    _PlayerInfoService_SearchPlayer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Player.proto",
}
