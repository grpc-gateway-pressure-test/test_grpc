package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "test_grpc/gw/pb"
)

const (
	port = ":51001"
)

type server struct {
}

func (s *server) Echo(c context.Context, v *pb.StringMessage) (*pb.StringMessage, error) {
	result := &pb.StringMessage{Value: v.Value}
	return result, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, &server{})
	s.Serve(lis)

}
