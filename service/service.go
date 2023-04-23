package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/grpc-template/protos"
	"google.golang.org/grpc"
)

type UserServer struct {
	pb.UnimplementedUserManagementServer
}

func (s *UserServer) CreateNewUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("Received: %v", req.GetName())
	var userId int32 = int32(rand.Intn(1000))
	
	return &pb.UserResponse{Name: req.GetName(), Age: req.GetAge(), Id: userId}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, &UserServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}