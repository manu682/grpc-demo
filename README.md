# A Sample gRPC server/client

This repo has a gRPC client and server that is used for showcasing how Golang struct composition/type enbedding can be used for doing a forward compatible RPC definition without breaking the code.

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

## Commands

### To generate the pb file from the proto definition

`protoc -I demoservice demoservice.proto --go_out=plugins=grpc:demoservice`

### To start the gRPC server

`go run .\servermain.go`

Output:
In... server...
Running the server...

### To start the gRPC client that calls the implemented method

`go run .\client_impl\client_impl.go`

Output:
In.... client...
responseId:10

### To start the gRPC client that calls the un-implemented method

`go run .\client_nonimpl\client_nonimpl.go`

Output:
In.... client...
Error : &{0xc000047c00} rpc error: code = Unimplemented desc = method GetData1 not implemented
<nil>
