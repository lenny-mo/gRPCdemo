package main

import (
	"context"
	"fmt"
	"sync"
	"testProtobuf/client"
	"testProtobuf/server"

	"time"
)

func main() {

	wg := sync.WaitGroup{}

	// 创建一个带有取消功能的上下文
	ctx, cancel := context.WithCancel(context.Background())

	// 启动 gRPC 服务器
	go server.StartServer()

	// 启动 gRPC 客户端
	wg.Add(1)
	go worker(ctx, &wg, client.StartClient)

	// 等待服务器和客户端执行结束
	time.Sleep(5 * time.Second)
	fmt.Println("Stopping server and client...") 
	cancel()
	wg.Wait()
}

func worker(ctx context.Context, wg *sync.WaitGroup, f func()) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(1 * time.Second)
			f()
		}
	}

}
