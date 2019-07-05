package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

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
	result := &pb.StringMessage{Value: g_content}
	return result, nil
}

var g_content string

func ReadContent(level string) string {
	path := ""
	if level == "1" {
		path = "./1B.txt"
	} else if level == "2" {
		path = "./1KB.txt"
	} else if level == "3" {
		path = "./1MB.txt"
	}
	fmt.Println(path)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("error : %s", err)
		return ""
	}
	return string(bytes)
}

func main() {
	level := os.Args[1]
	g_content = ReadContent(level)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, &server{})
	s.Serve(lis)

}
