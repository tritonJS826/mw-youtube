package main

import (
	"context"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) GetUser(ctx context.Context, req){
	return &UserResponse{Message: "Here's the user"}
}

func main(){
	lis, err :=net.Listen("tcp", ":50051")
	if err != nil{
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	RegisterUserServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil{
		log.Fatalf("failed to serve: %v", err)
	}
}
