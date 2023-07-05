package student

import (
	"context"
)

// 实现服务端的代码
type StudentServer struct{}

func (s *StudentServer) GetStudent(ctx context.Context, in *StudentRequest) (*StudentResponse, error) {
	return &StudentResponse{
		Student: &Student{
			Name:    in.Name,
			Age:     20,
			Address: "Beijing",
		},
	}, nil
}

func (s *StudentServer) mustEmbedUnimplementedStudentServiceServer() {
}
