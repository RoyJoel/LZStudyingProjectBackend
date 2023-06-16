package main

import (
	// "context"
	// "fmt"
	"log"

	"github.com/RoyJoel/LZStudyingProject/proto"
	"google.golang.org/grpc"
)

type PlayerInfoServiceClientImpl struct {
	proto.PlayerInfoServiceClient
}

func NewPlayerInfoServiceClientImpl() *PlayerInfoServiceClientImpl {
	conn, err := grpc.Dial(":6666", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("正在监听服务端 : %v\n", err)
	}
	return &PlayerInfoServiceClientImpl{proto.NewPlayerInfoServiceClient(conn)}
}

func playerActive() {
	//1 配置grpc服务端的端口作为客户端的监听
	conn, err := grpc.Dial(":6666", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("正在监听服务端 : %v\n", err)
	}
	defer conn.Close()
	//2 实例化 UserInfoService 服务的客户端
	// client := proto.NewPlayerInfoServiceClient(conn)
	//3 调用grpc服务
	// req := new(proto.PlayerInfoRequest)
	// req.LoginName = "abc"
	// resp, err := client.FindAllPlayers(context.Background(), req)
	// if err != nil {
	// 	log.Fatalf("请求错误 : %v\n", err)
	// }
	// fmt.Printf("响应内容 : %v\n", resp)
}
