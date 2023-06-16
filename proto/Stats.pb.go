package proto

// import (
// 	context "context"
// 	reflect "reflect"
// 	sync "sync"

// 	"github.com/RoyJoel/LZStudyingProjectBackend/package/model"
// 	grpc "google.golang.org/grpc"
// 	codes "google.golang.org/grpc/codes"
// 	status "google.golang.org/grpc/status"
// 	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
// 	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
// )

// const (
// 	// Verify that this generated code is sufficiently up-to-date.
// 	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
// 	// Verify that runtime/protoimpl is sufficiently up-to-date.
// 	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
// )

// // 定义客户端请求的数据格式
// // message 对应  生成的代码中的struct
// type StatsInfoRequest struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	// [修饰符] 类型 字段名 = 标识符
// 	LoginName        string         `protobuf:"bytes,1,opt,name=loginName,proto3" json:"loginName,omitempty"`
// 	Name             string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
// 	Icon             []byte         `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
// 	Sex              string         `protobuf:"varint,4,opt,name=sex,proto3" json:"sex,omitempty"`
// 	Age              int64          `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`
// 	YearsPlayed      int64          `protobuf:"varint,6,opt,name=yearsPlayed,proto3" json:"yearsPlayed,omitempty"`
// 	Height           float32        `protobuf:"float,7,opt,name=height,proto3" json:"height,omitempty"`
// 	Width            float32        `protobuf:"float,8,opt,name=width,proto3" json:"width,omitempty"`
// 	Grip             string         `protobuf:"bytes,9,opt,name=grip,proto3" json:"grip,omitempty"`
// 	Backhand         string         `protobuf:"bytes,10,opt,name=backhand,proto3" json:"backhand,omitempty"`
// 	Points           int64          `protobuf:"bytes,11,opt,name=points,proto3" json:"points,omitempty"`
// 	IsAdult          bool           `protobuf:"bytes,12,opt,name=isAdult,proto3" json:"isAdult,omitempty"`
// 	CareerStats      *model.Stats   `protobuf:"bytes,13,opt,name=careerStats,proto3" json:"careerStats,omitempty"`
// 	Friends          []model.Stats `protobuf:"bytes,14,rep,name=criends,proto3" json:"friends,omitempty""`
// 	GamesPlayed      []model.Game   `protobuf:"bytes,15,rep,name=gameStatsed,proto3" json:"gameStatsed,omitempty"`
// 	GamesToContinued []model.Game   `protobuf:"bytes,16,rep,name=gameToContinued,proto3" json:"gameToContinued,omitempty"`
// }

// func (x *StatsInfoRequest) Reset() {
// 	*x = StatsInfoRequest{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_LZ_proto_msgTypes[0]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *StatsInfoRequest) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*StatsInfoRequest) ProtoMessage() {}

// func (x *StatsInfoRequest) ProtoReflect() protoreflect.Message {
// 	mi := &file_LZ_proto_msgTypes[0]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use StatsInfoRequest.ProtoReflect.Descriptor instead.
// func (*StatsInfoRequest) Descriptor() ([]byte, []int) {
// 	return file_LZ_proto_rawDescGZIP(), []int{0}
// }

// func (x *StatsInfoRequest) GetLoginName() string {
// 	if x != nil {
// 		return x.LoginName
// 	}
// 	return ""
// }

// func (x *StatsInfoRequest) GetName() string {
// 	if x != nil {
// 		return x.Name
// 	}
// 	return ""
// }

// func (x *StatsInfoRequest) GetIcon() []byte {
// 	if x != nil {
// 		return x.Icon
// 	}
// 	return []byte{}
// }

// func (x *StatsInfoRequest) GetSex() string {
// 	if x != nil {
// 		return x.Sex
// 	}
// 	return ""
// }

// func (x *StatsInfoRequest) GetAge() int64 {
// 	if x != nil {
// 		return x.Age
// 	}
// 	return 0
// }

// func (x *StatsInfoRequest) GetYearsPlayed() int64 {
// 	if x != nil {
// 		return x.YearsPlayed
// 	}
// 	return 0
// }

// func (x *StatsInfoRequest) GetHeight() float32 {
// 	if x != nil {
// 		return x.Height
// 	}
// 	return 0
// }

// func (x *StatsInfoRequest) GetWidth() float32 {
// 	if x != nil {
// 		return x.Width
// 	}
// 	return 0
// }

// func (x *StatsInfoRequest) GetGrip() string {
// 	if x != nil {
// 		return x.Grip
// 	}
// 	return ""
// }

// func (x *StatsInfoRequest) GetBackhand() string {
// 	if x != nil {
// 		return x.Backhand
// 	}
// 	return ""
// }

// func (x *StatsInfoRequest) GetPoints() int64 {
// 	if x != nil {
// 		return x.Points
// 	}
// 	return 0
// }

// func (x *StatsInfoRequest) GetIsAdult() bool {
// 	if x != nil {
// 		return x.IsAdult
// 	}
// 	return true
// }

// func (x *StatsInfoRequest) GetCareerStats() *model.Stats {
// 	if x != nil {
// 		return x.CareerStats
// 	}
// 	return &model.Stats{}
// }

// func (x *StatsInfoRequest) GetFriends() []model.Stats {
// 	if x != nil {
// 		return x.Friends
// 	}
// 	return make([]model.Stats, 0)
// }

// func (x *StatsInfoRequest) GetGamesPlayed() []model.Game {
// 	if x != nil {
// 		return x.GamesPlayed
// 	}
// 	return make([]model.Game, 0)
// }

// func (x *StatsInfoRequest) GetGamesToContinued() []model.Game {
// 	if x != nil {
// 		return x.GamesToContinued
// 	}
// 	return make([]model.Game, 0)
// }

// type StatsInfoResponse struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	Code  int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
// 	Msg   string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
// 	Count int64  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
// 	Data  string `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
// }

// func (x *StatsInfoResponse) Reset() {
// 	*x = StatsInfoResponse{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_LZ_proto_msgTypes[1]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *StatsInfoResponse) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*StatsInfoResponse) ProtoMessage() {}

// func (x *StatsInfoResponse) ProtoReflect() protoreflect.Message {
// 	mi := &file_LZ_proto_msgTypes[1]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use StatsInfoResponse.ProtoReflect.Descriptor instead.
// func (*StatsInfoResponse) Descriptor() ([]byte, []int) {
// 	return file_LZ_proto_rawDescGZIP(), []int{1}
// }

// func (x *StatsInfoResponse) GetCode() int64 {
// 	if x != nil {
// 		return x.Code
// 	}
// 	return 0
// }

// func (x *StatsInfoResponse) GetMsg() string {
// 	if x != nil {
// 		return x.Msg
// 	}
// 	return ""
// }

// func (x *StatsInfoResponse) GetCount() int64 {
// 	if x != nil {
// 		return x.Count
// 	}
// 	return 0
// }

// func (x *StatsInfoResponse) GetData() string {
// 	if x != nil {
// 		return x.Data
// 	}
// 	return ""
// }

// var File_LZ_proto protoreflect.FileDescriptor

// var file_LZ_proto_rawDesc = []byte{
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
// 	file_LZ_proto_rawDescOnce sync.Once
// 	file_LZ_proto_rawDescData = file_LZ_proto_rawDesc
// )

// func file_LZ_proto_rawDescGZIP() []byte {
// 	file_LZ_proto_rawDescOnce.Do(func() {
// 		file_LZ_proto_rawDescData = protoimpl.X.CompressGZIP(file_LZ_proto_rawDescData)
// 	})
// 	return file_LZ_proto_rawDescData
// }

// var file_LZ_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
// var file_LZ_proto_goTypes = []interface{}{
// 	(*StatsInfoRequest)(nil),  // 0: proto.StatsInfoRequest
// 	(*StatsInfoResponse)(nil), // 1: proto.StatsInfoResponse
// }
// var file_LZ_proto_depIdxs = []int32{
// 	0, // 0: proto.StatsInfoService.GetStatsInfoById:input_type -> proto.StatsInfoRequest
// 	0, // 1: proto.StatsInfoService.AddNumByKey:input_type -> proto.StatsInfoRequest
// 	0, // 2: proto.StatsInfoService.FindStatsInfoByKey:input_type -> proto.StatsInfoRequest
// 	0, // 3: proto.StatsInfoService.SaveStatsInfo:input_type -> proto.StatsInfoRequest
// 	0, // 4: proto.StatsInfoService.DeleteById:input_type -> proto.StatsInfoRequest
// 	0, // 5: proto.StatsInfoService.FindAllStatss:input_type -> proto.StatsInfoRequest
// 	1, // 6: proto.StatsInfoService.GetStatsInfoById:output_type -> proto.StatsInfoResponse
// 	1, // 7: proto.StatsInfoService.AddNumByKey:output_type -> proto.StatsInfoResponse
// 	1, // 8: proto.StatsInfoService.FindStatsInfoByKey:output_type -> proto.StatsInfoResponse
// 	1, // 9: proto.StatsInfoService.SaveStatsInfo:output_type -> proto.StatsInfoResponse
// 	1, // 10: proto.StatsInfoService.DeleteById:output_type -> proto.StatsInfoResponse
// 	1, // 11: proto.StatsInfoService.FindAllStatss:output_type -> proto.StatsInfoResponse
// 	6, // [6:12] is the sub-list for method output_type
// 	0, // [0:6] is the sub-list for method input_type
// 	0, // [0:0] is the sub-list for extension type_name
// 	0, // [0:0] is the sub-list for extension extendee
// 	0, // [0:0] is the sub-list for field type_name
// }

// func init() { file_LZ_proto_init() }
// func file_LZ_proto_init() {
// 	if File_LZ_proto != nil {
// 		return
// 	}
// 	if !protoimpl.UnsafeEnabled {
// 		file_LZ_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*StatsInfoRequest); i {
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
// 		file_LZ_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*StatsInfoResponse); i {
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
// 			RawDescriptor: file_LZ_proto_rawDesc,
// 			NumEnums:      0,
// 			NumMessages:   2,
// 			NumExtensions: 0,
// 			NumServices:   1,
// 		},
// 		GoTypes:           file_LZ_proto_goTypes,
// 		DependencyIndexes: file_LZ_proto_depIdxs,
// 		MessageInfos:      file_LZ_proto_msgTypes,
// 	}.Build()
// 	File_LZ_proto = out.File
// 	file_LZ_proto_rawDesc = nil
// 	file_LZ_proto_goTypes = nil
// 	file_LZ_proto_depIdxs = nil
// }

// // Reference imports to suppress errors if they are not otherwise used.
// var _ context.Context
// var _ grpc.ClientConnInterface

// // This is a compile-time assertion to ensure that this generated file
// // is compatible with the grpc package it is being compiled against.
// const _ = grpc.SupportPackageIsVersion6

// // StatsInfoServiceClient is the client API for StatsInfoService service.
// //
// // For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
// type StatsInfoServiceClient interface {
// 	//rpc 方式；定义GetUserInfo 远程调用
// 	GetStatsInfoById(ctx context.Context, in *StatsInfoRequest, opts ...grpc.CallOption) (*StatsInfoResponse, error)
// 	AddStats(ctx context.Context, in *StatsInfoRequest, opts ...grpc.CallOption) (*StatsInfoResponse, error)
// 	// FindStatsInfoByKey(ctx context.Context, in *StatsInfoRequest, opts ...grpc.CallOption) (*StatsInfoResponse, error)
// 	SaveStatsInfo(ctx context.Context, in *StatsInfoRequest, opts ...grpc.CallOption) (*StatsInfoResponse, error)
// 	// DeleteById(ctx context.Context, in *StatsInfoRequest, opts ...grpc.CallOption) (*StatsInfoResponse, error)
// 	FindAllStatss(ctx context.Context, in *StatsInfoRequest, opts ...grpc.CallOption) (*StatsInfoResponse, error)
// }

// type StatsInfoServiceClient struct {
// 	cc grpc.ClientConnInterface
// }

// func NewStatsInfoServiceClient(cc grpc.ClientConnInterface) StatsInfoServiceClient {
// 	return &StatsInfoServiceClient{cc}
// }

// func (c *StatsInfoServiceClient) GetStatsInfoById(ctx context.Context, in *StatsInfoRequest, opts ...grpc.CallOption) (*StatsInfoResponse, error) {
// 	out := new(StatsInfoResponse)
// 	err := c.cc.Invoke(ctx, "/proto.StatsInfoService/GetStatsInfoById", in, out, opts...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return out, nil
// }

// func (c *StatsInfoServiceClient) AddStats(ctx context.Context, in *StatsInfoRequest, opts ...grpc.CallOption) (*StatsInfoResponse, error) {
// 	out := new(StatsInfoResponse)
// 	err := c.cc.Invoke(ctx, "/proto.StatsInfoService/AddStatsByKey", in, out, opts...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return out, nil
// }

// // func (c *StatsInfoServiceClient) FindStatsInfoByKey(ctx context.Context, in *StatsInfoRequest, opts ...grpc.CallOption) (*StatsInfoResponse, error) {
// // 	out := new(StatsInfoResponse)
// // 	err := c.cc.Invoke(ctx, "/proto.StatsInfoService/FindStatsInfoByKey", in, out, opts...)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	return out, nil
// // }

// func (c *StatsInfoServiceClient) SaveStatsInfo(ctx context.Context, in *StatsInfoRequest, opts ...grpc.CallOption) (*StatsInfoResponse, error) {
// 	out := new(StatsInfoResponse)
// 	err := c.cc.Invoke(ctx, "/proto.StatsInfoService/SaveStatsInfo", in, out, opts...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return out, nil
// }

// // func (c *StatsInfoServiceClient) DeleteById(ctx context.Context, in *StatsInfoRequest, opts ...grpc.CallOption) (*StatsInfoResponse, error) {
// // 	out := new(StatsInfoResponse)
// // 	err := c.cc.Invoke(ctx, "/proto.StatsInfoService/DeleteById", in, out, opts...)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	return out, nil
// // }

// func (c *StatsInfoServiceClient) FindAllStatss(ctx context.Context, in *StatsInfoRequest, opts ...grpc.CallOption) (*StatsInfoResponse, error) {
// 	out := new(StatsInfoResponse)
// 	err := c.cc.Invoke(ctx, "/proto.StatsInfoService/FindAllStatss", in, out, opts...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return out, nil
// }

// // StatsInfoServiceServer is the server API for StatsInfoService service.
// type StatsInfoServiceServer interface {
// 	//rpc 方式；定义GetUserInfo 远程调用
// 	GetStatsInfoByUid(context.Context, *StatsInfoRequest) (*StatsInfoResponse, error)
// 	AddStatsByKey(context.Context, *StatsInfoRequest) (*StatsInfoResponse, error)
// 	// FindStatsInfoByKey(context.Context, *StatsInfoRequest) (*StatsInfoResponse, error)
// 	SaveStatsInfo(context.Context, *StatsInfoRequest) (*StatsInfoResponse, error)
// 	// DeleteById(context.Context, *StatsInfoRequest) (*StatsInfoResponse, error)
// 	FindAllStatss(context.Context, *StatsInfoRequest) (*StatsInfoResponse, error)
// }

// // UnimplementedStatsInfoServiceServer can be embedded to have forward compatible implementations.
// type UnimplementedStatsInfoServiceServer struct {
// }

// func (*UnimplementedStatsInfoServiceServer) GetStatsInfoByUid(context.Context, *StatsInfoRequest) (*StatsInfoResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method GetStatsInfoById not implemented")
// }
// func (*UnimplementedStatsInfoServiceServer) AddNumByKey(context.Context, *StatsInfoRequest) (*StatsInfoResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method AddNumByKey not implemented")
// }

// //	func (*UnimplementedStatsInfoServiceServer) FindStatsInfoByKey(context.Context, *StatsInfoRequest) (*StatsInfoResponse, error) {
// //		return nil, status.Errorf(codes.Unimplemented, "method FindStatsInfoByKey not implemented")
// //	}
// func (*UnimplementedStatsInfoServiceServer) SaveStatsInfo(context.Context, *StatsInfoRequest) (*StatsInfoResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method SaveStatsInfo not implemented")
// }

// //	func (*UnimplementedStatsInfoServiceServer) DeleteById(context.Context, *StatsInfoRequest) (*StatsInfoResponse, error) {
// //		return nil, status.Errorf(codes.Unimplemented, "method DeleteById not implemented")
// //	}
// func (*UnimplementedStatsInfoServiceServer) FindAllStatss(context.Context, *StatsInfoRequest) (*StatsInfoResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method FindAllStatss not implemented")
// }

// func RegisterStatsInfoServiceServer(s *grpc.Server, srv StatsInfoServiceServer) {
// 	s.RegisterService(&_StatsInfoService_serviceDesc, srv)
// }

// func _StatsInfoService_GetStatsInfoByUid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
// 	in := new(StatsInfoRequest)
// 	if err := dec(in); err != nil {
// 		return nil, err
// 	}
// 	if interceptor == nil {
// 		return srv.(StatsInfoServiceServer).GetStatsInfoByUid(ctx, in)
// 	}
// 	info := &grpc.UnaryServerInfo{
// 		Server:     srv,
// 		FullMethod: "/proto.StatsInfoService/GetStatsInfoById",
// 	}
// 	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
// 		return srv.(StatsInfoServiceServer).GetStatsInfoByUid(ctx, req.(*StatsInfoRequest))
// 	}
// 	return interceptor(ctx, in, info, handler)
// }

// func _StatsInfoService_AddStatsByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
// 	in := new(StatsInfoRequest)
// 	if err := dec(in); err != nil {
// 		return nil, err
// 	}
// 	if interceptor == nil {
// 		return srv.(StatsInfoServiceServer).AddStatsByKey(ctx, in)
// 	}
// 	info := &grpc.UnaryServerInfo{
// 		Server:     srv,
// 		FullMethod: "/proto.StatsInfoService/AddNumByKey",
// 	}
// 	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
// 		return srv.(StatsInfoServiceServer).AddStatsByKey(ctx, req.(*StatsInfoRequest))
// 	}
// 	return interceptor(ctx, in, info, handler)
// }

// // func _StatsInfoService_FindStatsInfoByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
// // 	in := new(StatsInfoRequest)
// // 	if err := dec(in); err != nil {
// // 		return nil, err
// // 	}
// // 	if interceptor == nil {
// // 		return srv.(StatsInfoServiceServer).FindStatsInfoByKey(ctx, in)
// // 	}
// // 	info := &grpc.UnaryServerInfo{
// // 		Server:     srv,
// // 		FullMethod: "/proto.StatsInfoService/FindStatsInfoByKey",
// // 	}
// // 	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
// // 		return srv.(StatsInfoServiceServer).FindStatsInfoByKey(ctx, req.(*StatsInfoRequest))
// // 	}
// // 	return interceptor(ctx, in, info, handler)
// // }

// func _StatsInfoService_SaveStatsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
// 	in := new(StatsInfoRequest)
// 	if err := dec(in); err != nil {
// 		return nil, err
// 	}
// 	if interceptor == nil {
// 		return srv.(StatsInfoServiceServer).SaveStatsInfo(ctx, in)
// 	}
// 	info := &grpc.UnaryServerInfo{
// 		Server:     srv,
// 		FullMethod: "/proto.StatsInfoService/SaveStatsInfo",
// 	}
// 	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
// 		return srv.(StatsInfoServiceServer).SaveStatsInfo(ctx, req.(*StatsInfoRequest))
// 	}
// 	return interceptor(ctx, in, info, handler)
// }

// // func _StatsInfoService_DeleteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
// // 	in := new(StatsInfoRequest)
// // 	if err := dec(in); err != nil {
// // 		return nil, err
// // 	}
// // 	if interceptor == nil {
// // 		return srv.(StatsInfoServiceServer).DeleteById(ctx, in)
// // 	}
// // 	info := &grpc.UnaryServerInfo{
// // 		Server:     srv,
// // 		FullMethod: "/proto.StatsInfoService/DeleteById",
// // 	}
// // 	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
// // 		return srv.(StatsInfoServiceServer).DeleteById(ctx, req.(*StatsInfoRequest))
// // 	}
// // 	return interceptor(ctx, in, info, handler)
// // }

// func _StatsInfoService_FindAllStatss_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
// 	in := new(StatsInfoRequest)
// 	if err := dec(in); err != nil {
// 		return nil, err
// 	}
// 	if interceptor == nil {
// 		return srv.(StatsInfoServiceServer).FindAllStatss(ctx, in)
// 	}
// 	info := &grpc.UnaryServerInfo{
// 		Server:     srv,
// 		FullMethod: "/proto.StatsInfoService/FindAllStatss",
// 	}
// 	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
// 		return srv.(StatsInfoServiceServer).FindAllStatss(ctx, req.(*StatsInfoRequest))
// 	}
// 	return interceptor(ctx, in, info, handler)
// }

// var _StatsInfoService_serviceDesc = grpc.ServiceDesc{
// 	ServiceName: "proto.StatsInfoService",
// 	HandlerType: (*StatsInfoServiceServer)(nil),
// 	Methods: []grpc.MethodDesc{
// 		{
// 			MethodName: "GetStatsInfoById",
// 			Handler:    _StatsInfoService_GetStatsInfoByUid_Handler,
// 		},
// 		{
// 			MethodName: "AddNumByKey",
// 			Handler:    _StatsInfoService_AddStatsByKey_Handler,
// 		},
// 		// {
// 		// 	MethodName: "FindStatsInfoByKey",
// 		// 	Handler:    _StatsInfoService_FindStatsInfoByKey_Handler,
// 		// },
// 		{
// 			MethodName: "SaveStatsInfo",
// 			Handler:    _StatsInfoService_SaveStatsInfo_Handler,
// 		},
// 		// {
// 		// 	MethodName: "DeleteById",
// 		// 	Handler:    _StatsInfoService_DeleteById_Handler,
// 		// },
// 		{
// 			MethodName: "FindAllStatss",
// 			Handler:    _StatsInfoService_FindAllStatss_Handler,
// 		},
// 	},
// 	Streams:  []grpc.StreamDesc{},
// 	Metadata: "LZ.proto",
// }
