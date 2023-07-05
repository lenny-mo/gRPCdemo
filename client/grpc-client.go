package client

import (
	"context"
	"fmt"
	"log"

	"testProtobuf/student/student"

	"google.golang.org/grpc"
)

func StartClient() {
	// 1. 创建一个 `studentServiceClient` 实例，并将 gRPC 客户端连接传递给它
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := student.NewStudentServiceClient(conn)

	// 2. 使用 `client` 变量调用 `GetStudent` 方法
	req := &student.StudentRequest{
		Name: "Alice",
	}

	resp, err := client.GetStudent(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call GetStudent: %v", err)
	}

	fmt.Println("client gets result from server: ", resp.Student)
}
