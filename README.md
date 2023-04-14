# go-grpc-experiments
some code to experiment with gRPC in Go

### How to try the gRPC calc example

go the base directory of the project, and run from your shell :

     go run remoteCalculator/server/main.go 

open another terminal and run the client 

    go run remoteCalculator/client/main.go -num1 20 -num2 10 -operation Add

### How to generate the calc protoc golang files

    scripts/protobuf_generate.sh
    

### More info

+ [gRPC Go basic tutorial](https://grpc.io/docs/languages/go/basics/)
+ [Protocol Buffers doc](https://protobuf.dev/)
+ [Protocol Buffer Compiler Installation](https://grpc.io/docs/protoc-installation/)
+ [Evans: universal gRPC client](https://github.com/ktr0731/evans)