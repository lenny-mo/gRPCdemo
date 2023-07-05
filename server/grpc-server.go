package server

import (
	"fmt"
	"log"
	"net"

	"testProtobuf/student/student"

	"google.golang.org/grpc"
)	

func StartServer() {
	// 创建一个监听端口的 TCP 连接
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// 创建 gRPC 服务器
	grpcServer := grpc.NewServer()

	// 将 StudentServiceServer 注册到 gRPC 服务器上
	student.RegisterStudentServiceServer(grpcServer, &student.StudentServer{})

	// 开启一个 goroutine 监听端口并接受客户端请求
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	fmt.Println("Server started. Listening on :50051")
	fmt.Println("Press Ctrl+C to stop the server.")

	// 阻塞主 goroutine，直到接收到中断信号
	// 当接收到中断信号后，停止服务器
	<-make(chan struct{})
	grpcServer.Stop()
}
