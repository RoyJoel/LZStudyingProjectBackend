package main

import (
	"github.com/RoyJoel/LZStudyingProject/cmd/grpc"
	"github.com/RoyJoel/LZStudyingProject/cmd/web"
)

func main() {
	go grpc.Run()
	go web.Run()
	select {}
}
