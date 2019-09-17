package main

import (
	"flag"
	"fmt"
	"net"
	context "context"

	"google.golang.org/grpc"

	pb "grpc-demo/demoservice"
)

var (
	port = ":5000"
)

// Note the struct composition of the auto-generated "pb.UnimplementedDemoServiceServer"
type demoServiceServer struct {
	pb.UnimplementedDemoServiceServer
}

// Getdata returns the data
func (s *demoServiceServer) GetData(ctx context.Context, input *pb.Input) (*pb.Output, error) {
	// No feature was found, return an unnamed feature
	return &pb.Output{ResponseId: 10}, nil
}

func main() {
	flag.Parse()
	fmt.Println("In... server...")
	myServer := grpc.NewServer()

	//This can be used when the implementation has not yet started.
	//This will return the method not implemented for all RPCs
	//pb.RegisterDemoServiceServer(myServer, &pb.UnimplementedDemoServiceServer{})
	
	pb.RegisterDemoServiceServer(myServer, &demoServiceServer{})
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	fmt.Println("Running the server...")
	err1 := myServer.Serve(listen)
	if err1 != nil {
		fmt.Println("Starting the server has failed : ", err)
	}
}
