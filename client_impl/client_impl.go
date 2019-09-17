package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"google.golang.org/grpc"

	pb "grpc-demo/demoservice"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:5000", "The server address in the format of host:port")
)

func main() {
	fmt.Println("In.... client...")
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		fmt.Println("fail to dial: ", err)
	}
	defer conn.Close()

	client := pb.NewDemoServiceClient(conn)
	getDemoServiceResult(client, &pb.Input{RequestId: 32455})
}

func getDemoServiceResult(client pb.DemoServiceClient, input *pb.Input) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	output, err := client.GetData(ctx, input)
	if err != nil {
		fmt.Println("Error : ", client, err)
	}
	fmt.Println(output)
}
