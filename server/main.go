package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/nari-z/grpc-sample/server/protocol"
)

func main() {
	fmt.Println("Hello gRPC.")

	// launch server
	listenPort, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	service := &SampleService{}
	sample.RegisterSampleServiceServer(server, service)
	// Register reflection service on gRPC server.
	reflection.Register(server)
	server.Serve(listenPort)
}

// implement gRPC service
type SampleService struct {
}

func (s *SampleService) RequestSampleMethod(ctx context.Context, req *sample.SampleRequest) (*sample.SampleResponse, error) {
	if (req == nil) {
		return nil, errors.New("SampleRequest is nil")
	}
	fmt.Println("Receive -> " + req.GetRequestMessage())
	return &sample.SampleResponse{ Id: req.GetId(), ResponseMessage: "This is ResponseMessage." }, nil
}
