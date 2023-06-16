package main

import (
	"github.com/RoyJoel/LZStudyingProjectBackend/cmd/web"
	"github.com/RoyJoel/LZStudyingProjectBackendBackend/cmd/grpc"
)

func main() {
	go grpc.Run()
	go web.Run()
	select {}
}
