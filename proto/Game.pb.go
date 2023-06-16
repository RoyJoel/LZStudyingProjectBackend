package proto

import (
	context "context"
	// reflect "reflect"
	// sync "sync"

	"github.com/RoyJoel/LZStudyingProject/package/utils"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	// protoreflect "google.golang.org/protobuf/reflect/protoreflect"
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
type GameInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// [修饰符] 类型 字段名 = 标识符
	Date                float64          `protobuf:"varint,1,opt,name=date,proto3" json:"date,omitempty"`
	Place               string           `protobuf:"bytes,2,opt,name=place,proto3" json:"place,omitempty"`
	Surface             string           `protobuf:"bytes,3,opt,name=surface,proto3" json:"surface,omitempty"`
	SetNum              int              `protobuf:"varint,4,opt,name=setNum,proto3" json:"setNum,omitempty"`
	GameNum             int              `protobuf:"varint,5,opt,name=gameNum,proto3" json:"gameNum,omitempty"`
	IsGoldenGoal        bool             `protobuf:"varint,6,opt,name=isGoldenGoal,proto3" json:"isGoldenGoal,omitempty"`
	IsPlayer1Serving    bool             `protobuf:"varint,7,opt,name=isPlayer1Serving,proto3" json:"isPlayer1Serving,omitempty"`
	IsCompleted         bool             `protobuf:"float,8,opt,name=isCompleted,proto3" json:"isCompleted,omitempty"`
	Player1LoginName    string           `protobuf:"float,9,opt,name=player1LoginName,proto3" json:"player1LoginName,omitempty"`
	Player1StatsId      int              `protobuf:"varint,10,opt,name=player1StatsId,proto3" json:"player1StatsId,omitempty"`
	Player2LoginName    string           `protobuf:"bytes,11,opt,name=player2LoginName,proto3" json:"player2LoginName,omitempty"`
	Player2StatsId      int              `protobuf:"varint,12,opt,name=player2StatsId ,proto3" json:"player2StatsId,omitempty"`
	IsPlayer1FirstServe bool             `protobuf:"varint,13,opt,name=isPlayer1FirstServe ,proto3" json:"isPlayer1FirstServe,omitempty"`
	IsPlayer2FirstServe bool             `protobuf:"varint,14,opt,name=isPlayer2FirstServe ,proto3" json:"isPlayer2FirstServe,omitempty"`
	IsPlayer1Left       bool             `protobuf:"varint,15,opt,name=isPlayer1Left ,proto3" json:"isPlayer1Left,omitempty"`
	IsChangePosition    bool             `protobuf:"varint,16,opt,name=isChangePosition ,proto3" json:"isChangePosition,omitempty"`
	Result              utils.IntMatrix3 `protobuf:"bytes,17,opt,name=result,proto3" json:"result,omitempty"`
}

// func (x *GameInfoRequest) Reset() {
// 	*x = GameInfoRequest{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_Game_proto_msgTypes[0]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *GameInfoRequest) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*GameInfoRequest) ProtoMessage() {}

// func (x *GameInfoRequest) ProtoReflect() protoreflect.Message {
// 	mi := &file_Game_proto_msgTypes[0]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// Deprecated: Use GameInfoRequest.ProtoReflect.Descriptor instead.
// func (*GameInfoRequest) Descriptor() ([]byte, []int) {
// 	return file_Game_proto_rawDescGZIP(), []int{0}
// }

func (x *GameInfoRequest) GetDate() float64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *GameInfoRequest) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

func (x *GameInfoRequest) GetSurface() string {
	if x != nil {
		return x.Surface
	}
	return ""
}

func (x *GameInfoRequest) GetSetNum() int {
	if x != nil {
		return x.SetNum
	}
	return 0
}

func (x *GameInfoRequest) GetGameNum() int {
	if x != nil {
		return x.GameNum
	}
	return 0
}

func (x *GameInfoRequest) GetIsGoldenGoal() bool {
	if x != nil {
		return x.IsGoldenGoal
	}
	return false
}

func (x *GameInfoRequest) GetIsPlayer1Servingl() bool {
	if x != nil {
		return x.IsPlayer1Serving
	}
	return false
}

func (x *GameInfoRequest) GetIsCompleted() bool {
	if x != nil {
		return x.IsCompleted
	}
	return false
}

func (x *GameInfoRequest) GetPlayer1LoginName() string {
	if x != nil {
		return x.Player1LoginName
	}
	return ""
}

func (x *GameInfoRequest) GetPlayer1StatsId() int {
	if x != nil {
		return x.Player1StatsId
	}
	return 0
}

func (x *GameInfoRequest) GetPlayer2LoginName() string {
	if x != nil {
		return x.Player2LoginName
	}
	return ""
}

func (x *GameInfoRequest) GetPlayer2StatsId() int {
	if x != nil {
		return x.Player2StatsId
	}
	return 0
}

func (x *GameInfoRequest) GetIsPlayer1Left() bool {
	if x != nil {
		return x.IsPlayer1Left
	}
	return true
}

func (x *GameInfoRequest) GetIsChangePosition() bool {
	if x != nil {
		return x.IsChangePosition
	}
	return true
}

func (x *GameInfoRequest) GetResult() utils.IntMatrix3 {
	if x != nil {
		return x.Result
	}
	return make([][][]int64, 0)
}

type GameInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg   string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Count int64  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Data  string `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

// func (x *GameInfoResponse) Reset() {
// 	*x = GameInfoResponse{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_Game_proto_msgTypes[1]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *GameInfoResponse) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*GameInfoResponse) ProtoMessage() {}

// func (x *GameInfoResponse) ProtoReflect() protoreflect.Message {
// 	mi := &file_Game_proto_msgTypes[1]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use GameInfoResponse.ProtoReflect.Descriptor instead.
// func (*GameInfoResponse) Descriptor() ([]byte, []int) {
// 	return file_Game_proto_rawDescGZIP(), []int{1}
// }

func (x *GameInfoResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GameInfoResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GameInfoResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GameInfoResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

// var File_Game_proto protoreflect.FileDescriptor

// var file_Game_proto_rawDesc = []byte{
// 	0x0a, 0x0f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
// 	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x0b, 0x49, 0x6e, 0x66, 0x6f,
// 	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
// 	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
// 	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69,
// 	0x6e, 0x66, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69,
// 	0x6e, 0x66, 0x6f, 0x4b, 0x65, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6e,
// 	0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x6e, 0x66, 0x6f, 0x4e, 0x75,
// 	0x6d, 0x22, 0x5e, 0x0a, 0x0c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
// 	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
// 	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
// 	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
// 	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a,
// 	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
// 	0x61, 0x32, 0xef, 0x02, 0x0a, 0x0e, 0x4e, 0x75, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72,
// 	0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x49, 0x6e,
// 	0x66, 0x6f, 0x42, 0x79, 0x49, 0x64, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49,
// 	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f,
// 	0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
// 	0x00, 0x12, 0x38, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x4e, 0x75, 0x6d, 0x42, 0x79, 0x4b, 0x65, 0x79,
// 	0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
// 	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x66,
// 	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x10, 0x46,
// 	0x69, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x12,
// 	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
// 	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f,
// 	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0b, 0x53, 0x61,
// 	0x76, 0x65, 0x4e, 0x75, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
// 	0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
// 	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
// 	0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79,
// 	0x49, 0x64, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
// 	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49,
// 	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a,
// 	0x07, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
// 	0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70,
// 	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
// 	0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
// 	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
// }

// var (
// 	file_Game_proto_rawDescOnce sync.Once
// 	file_Game_proto_rawDescData = file_Game_proto_rawDesc
// )

// func file_Game_proto_rawDescGZIP() []byte {
// 	file_Game_proto_rawDescOnce.Do(func() {
// 		file_Game_proto_rawDescData = protoimpl.X.CompressGZIP(file_Game_proto_rawDescData)
// 	})
// 	return file_Game_proto_rawDescData
// }

// var file_Game_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
// var file_Game_proto_goTypes = []interface{}{
// 	(*GameInfoRequest)(nil),  // 0: proto.GameInfoRequest
// 	(*GameInfoResponse)(nil), // 1: proto.GameInfoResponse
// }
// var file_Game_proto_depIdxs = []int32{
// 	0, // 0: proto.GameInfoService.GetGameInfoById:input_type -> proto.GameInfoRequest
// 	0, // 1: proto.GameInfoService.AddNumByKey:input_type -> proto.GameInfoRequest
// 	0, // 2: proto.GameInfoService.FindGameInfoByKey:input_type -> proto.GameInfoRequest
// 	0, // 3: proto.GameInfoService.SaveGameInfo:input_type -> proto.GameInfoRequest
// 	0, // 4: proto.GameInfoService.DeleteById:input_type -> proto.GameInfoRequest
// 	0, // 5: proto.GameInfoService.FindAllGames:input_type -> proto.GameInfoRequest
// 	1, // 6: proto.GameInfoService.GetGameInfoById:output_type -> proto.GameInfoResponse
// 	1, // 7: proto.GameInfoService.AddNumByKey:output_type -> proto.GameInfoResponse
// 	1, // 8: proto.GameInfoService.FindGameInfoByKey:output_type -> proto.GameInfoResponse
// 	1, // 9: proto.GameInfoService.SaveGameInfo:output_type -> proto.GameInfoResponse
// 	1, // 10: proto.GameInfoService.DeleteById:output_type -> proto.GameInfoResponse
// 	1, // 11: proto.GameInfoService.FindAllGames:output_type -> proto.GameInfoResponse
// 	6, // [6:12] is the sub-list for method output_type
// 	0, // [0:6] is the sub-list for method input_type
// 	0, // [0:0] is the sub-list for extension type_name
// 	0, // [0:0] is the sub-list for extension extendee
// 	0, // [0:0] is the sub-list for field type_name
// }

// func init() { file_Game_proto_init() }
// func file_Game_proto_init() {
// 	if File_Game_proto != nil {
// 		return
// 	}
// 	if !protoimpl.UnsafeEnabled {
// 		file_Game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*GameInfoRequest); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_Game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*GameInfoResponse); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			default:
// 				return nil
// 			}
// 		}
// 	}
// 	type x struct{}
// 	out := protoimpl.TypeBuilder{
// 		File: protoimpl.DescBuilder{
// 			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
// 			RawDescriptor: file_Game_proto_rawDesc,
// 			NumEnums:      0,
// 			NumMessages:   2,
// 			NumExtensions: 0,
// 			NumServices:   1,
// 		},
// 		GoTypes:           file_Game_proto_goTypes,
// 		DependencyIndexes: file_Game_proto_depIdxs,
// 		MessageInfos:      file_Game_proto_msgTypes,
// 	}.Build()
// 	File_Game_proto = out.File
// 	file_Game_proto_rawDesc = nil
// 	file_Game_proto_goTypes = nil
// 	file_Game_proto_depIdxs = nil
// }

// // Reference imports to suppress errors if they are not otherwise used.
// var _ context.Context
// var _ grpc.ClientConnInterface

// // This is a compile-time assertion to ensure that this generated file
// // is compatible with the grpc package it is being compiled against.
// const _ = grpc.SupportPackageIsVersion6

// GameInfoServiceClient is the client API for GameInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameInfoServiceClient interface {
	//rpc 方式；定义GetUserInfo 远程调用
	SearchGame(ctx context.Context, in *GameInfoRequest, opts ...grpc.CallOption) (*GameInfoResponse, error)
	AddGame(ctx context.Context, in *GameInfoRequest, opts ...grpc.CallOption) (*GameInfoResponse, error)
	UpdateGame(ctx context.Context, in *GameInfoRequest, opts ...grpc.CallOption) (*GameInfoResponse, error)
	// FindGameInfoByKey(ctx context.Context, in *GameInfoRequest, opts ...grpc.CallOption) (*GameInfoResponse, error)
	// DeleteById(ctx context.Context, in *GameInfoRequest, opts ...grpc.CallOption) (*GameInfoResponse, error)
	// FindAllGames(ctx context.Context, in *GameInfoRequest, opts ...grpc.CallOption) (*GameInfoResponse, error)
}

type gameInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGameInfoServiceClient(cc grpc.ClientConnInterface) GameInfoServiceClient {
	return &gameInfoServiceClient{cc}
}

func (c *gameInfoServiceClient) SearchGame(ctx context.Context, in *GameInfoRequest, opts ...grpc.CallOption) (*GameInfoResponse, error) {
	out := new(GameInfoResponse)
	err := c.cc.Invoke(ctx, "/proto.GameInfoService/SearchGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameInfoServiceClient) AddGame(ctx context.Context, in *GameInfoRequest, opts ...grpc.CallOption) (*GameInfoResponse, error) {
	out := new(GameInfoResponse)
	err := c.cc.Invoke(ctx, "/proto.GameInfoService/AddGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameInfoServiceClient) UpdateGame(ctx context.Context, in *GameInfoRequest, opts ...grpc.CallOption) (*GameInfoResponse, error) {
	out := new(GameInfoResponse)
	err := c.cc.Invoke(ctx, "/proto.GameInfoService/UpdateGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// func (c *GameInfoServiceClient) FindGameInfoByKey(ctx context.Context, in *GameInfoRequest, opts ...grpc.CallOption) (*GameInfoResponse, error) {
// 	out := new(GameInfoResponse)
// 	err := c.cc.Invoke(ctx, "/proto.GameInfoService/FindGameInfoByKey", in, out, opts...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return out, nil
// }

// func (c *GameInfoServiceClient) DeleteById(ctx context.Context, in *GameInfoRequest, opts ...grpc.CallOption) (*GameInfoResponse, error) {
// 	out := new(GameInfoResponse)
// 	err := c.cc.Invoke(ctx, "/proto.GameInfoService/DeleteById", in, out, opts...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return out, nil
// }

// func (c *GameInfoServiceClient) FindAllGames(ctx context.Context, in *GameInfoRequest, opts ...grpc.CallOption) (*GameInfoResponse, error) {
// 	out := new(GameInfoResponse)
// 	err := c.cc.Invoke(ctx, "/proto.GameInfoService/FindAllGames", in, out, opts...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return out, nil
// }

// GameInfoServiceServer is the server API for GameInfoService service.
type GameInfoServiceServer interface {
	//rpc 方式；定义GetUserInfo 远程调用
	SearchGame(context.Context, *GameInfoRequest) (*GameInfoResponse, error)
	AddGame(context.Context, *GameInfoRequest) (*GameInfoResponse, error)
	UpdateGame(context.Context, *GameInfoRequest) (*GameInfoResponse, error)
	// FindGameInfoByKey(context.Context, *GameInfoRequest) (*GameInfoResponse, error)
	// DeleteById(context.Context, *GameInfoRequest) (*GameInfoResponse, error)
	// FindAllGames(context.Context, *GameInfoRequest) (*GameInfoResponse, error)
}

// UnimplementedGameInfoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGameInfoServiceServer struct {
}

func (*UnimplementedGameInfoServiceServer) SearchGame(context.Context, *GameInfoRequest) (*GameInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchGame not implemented")
}
func (*UnimplementedGameInfoServiceServer) AddGame(context.Context, *GameInfoRequest) (*GameInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGame not implemented")
}

func (*UnimplementedGameInfoServiceServer) UpdateGame(context.Context, *GameInfoRequest) (*GameInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGame not implemented")
}

//	func (*UnimplementedGameInfoServiceServer) FindGameInfoByKey(context.Context, *GameInfoRequest) (*GameInfoResponse, error) {
//		return nil, status.Errorf(codes.Unimplemented, "method FindGameInfoByKey not implemented")
//	}

//	func (*UnimplementedGameInfoServiceServer) DeleteById(context.Context, *GameInfoRequest) (*GameInfoResponse, error) {
//		return nil, status.Errorf(codes.Unimplemented, "method DeleteById not implemented")
//	}
// func (*UnimplementedGameInfoServiceServer) FindAllGames(context.Context, *GameInfoRequest) (*GameInfoResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method FindAllGames not implemented")
// }

func RegisterGameInfoServiceServer(s *grpc.Server, srv GameInfoServiceServer) {
	s.RegisterService(&_GameInfoService_serviceDesc, srv)
}

func _GameInfoService_SearchGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameInfoServiceServer).SearchGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GameInfoService/SearchGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameInfoServiceServer).SearchGame(ctx, req.(*GameInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameInfoService_AddGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameInfoServiceServer).AddGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GameInfoService/AddGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameInfoServiceServer).AddGame(ctx, req.(*GameInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameInfoService_UpdateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameInfoServiceServer).UpdateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GameInfoService/UpdateGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameInfoServiceServer).UpdateGame(ctx, req.(*GameInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// func _GameInfoService_FindGameInfoByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
// 	in := new(GameInfoRequest)
// 	if err := dec(in); err != nil {
// 		return nil, err
// 	}
// 	if interceptor == nil {
// 		return srv.(GameInfoServiceServer).FindGameInfoByKey(ctx, in)
// 	}
// 	info := &grpc.UnaryServerInfo{
// 		Server:     srv,
// 		FullMethod: "/proto.GameInfoService/FindGameInfoByKey",
// 	}
// 	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
// 		return srv.(GameInfoServiceServer).FindGameInfoByKey(ctx, req.(*GameInfoRequest))
// 	}
// 	return interceptor(ctx, in, info, handler)
// }

// func _GameInfoService_DeleteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
// 	in := new(GameInfoRequest)
// 	if err := dec(in); err != nil {
// 		return nil, err
// 	}
// 	if interceptor == nil {
// 		return srv.(GameInfoServiceServer).DeleteById(ctx, in)
// 	}
// 	info := &grpc.UnaryServerInfo{
// 		Server:     srv,
// 		FullMethod: "/proto.GameInfoService/DeleteById",
// 	}
// 	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
// 		return srv.(GameInfoServiceServer).DeleteById(ctx, req.(*GameInfoRequest))
// 	}
// 	return interceptor(ctx, in, info, handler)
// }

// func _GameInfoService_FindAllGames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
// 	in := new(GameInfoRequest)
// 	if err := dec(in); err != nil {
// 		return nil, err
// 	}
// 	if interceptor == nil {
// 		return srv.(GameInfoServiceServer).FindAllGames(ctx, in)
// 	}
// 	info := &grpc.UnaryServerInfo{
// 		Server:     srv,
// 		FullMethod: "/proto.GameInfoService/FindAllGames",
// 	}
// 	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
// 		return srv.(GameInfoServiceServer).FindAllGames(ctx, req.(*GameInfoRequest))
// 	}
// 	return interceptor(ctx, in, info, handler)
// }

var _GameInfoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GameInfoService",
	HandlerType: (*GameInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddGame",
			Handler:    _GameInfoService_AddGame_Handler,
		},
		{
			MethodName: "UpdateGame",
			Handler:    _GameInfoService_UpdateGame_Handler,
		},
		// {
		// 	MethodName: "FindGameInfoByKey",
		// 	Handler:    _GameInfoService_FindGameInfoByKey_Handler,
		// },
		{
			MethodName: "SearchGame",
			Handler:    _GameInfoService_SearchGame_Handler,
		},
		// {
		// 	MethodName: "DeleteById",
		// 	Handler:    _GameInfoService_DeleteById_Handler,
		// },
		// {
		// 	MethodName: "FindAllGames",
		// 	Handler:    _GameInfoService_FindAllGames_Handler,
		// },
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Game.proto",
}
