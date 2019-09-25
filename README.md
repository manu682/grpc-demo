# A Sample gRPC server/client

This repo has a gRPC client and server that is used for showcasing how Golang struct composition can be used for doing a forward compatible RPC definition without breaking the code.

## Case 1

1. Create a proto file and generate the pb file.
2. Create a gRPC server which responds to the RPC calls with an "unimplemented" error response.

## Solution

In the gRPC server registration, use the struct that is auto-generated for this - UnimplementedDemoServiceServer

<pre>
...
pb.RegisterDemoServiceServer(myServer, &pb.UnimplementedDemoServiceServer{})
...
</pre>

## Case 2

1. Create a service with an RPC that is already implemented
2. Extend the service with another RPC

## Solution

Use struct composition on the existing struct to embed the auto-generated struct.
This will ensure that any addition of a new RPC will automatically have an unimplemented response auto-generated for it.

<pre>
...
type demoServiceServer struct { 
    pb.UnimplementedDemoServiceServer 
}  
... 
pb.RegisterDemoServiceServer(myServer, &demoServiceServer{})
...
</pre>

## To generate the pb file from the proto definition

`protoc -I demoservice demoservice.proto --go_out=plugins=grpc:demoservice`

## To start the gRPC server

`go run .\servermain.go`

## To start the gRPC client

`go run .\client\client.go`
