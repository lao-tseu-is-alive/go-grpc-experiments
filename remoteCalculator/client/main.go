package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/lao-tseu-is-alive/go-grpc-experiments/remoteCalculator/calc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const defaultPort = 50051

// main is client to call the gRPC calc server, just try this from your shell :  go run remoteCalculator/client/main.go -num1 2 -num2 2 -operation M
func main() {
	port := flag.Int("port", defaultPort, "the address to connect to")
	num1 := flag.Float64("num1", 0, "first operand number")
	num2 := flag.Float64("num2", 0, "second operand number")
	op := flag.String("operation", "Add", "Operation to apply to the numbers [Add|Subtract|Multiply|Divide]")
	flag.Parse()
	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	//	127.0.0.1:50051
	log.Printf("Will try to connect to : %s", addr)
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalcServiceClient(conn)

	var calcOp pb.CalcRequest_OpType
	switch *op {
	case "Add", "A":
		calcOp = pb.CalcRequest_Add
	case "Subtract", "S":
		calcOp = pb.CalcRequest_Subtract
	case "Multiply", "M":
		calcOp = pb.CalcRequest_Multiply
	case "Divide", "D":
		calcOp = pb.CalcRequest_Divide
	default:
		calcOp = pb.CalcRequest_Add
	}

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	log.Printf("Sending num1: %.3f , num2: %.3f and operation :%s", *num1, *num2, calcOp.String())
	r, err := c.Calc(ctx, &pb.CalcRequest{
		Num1:      *num1,
		Num2:      *num2,
		Operation: calcOp,
	})
	if err != nil {
		log.Fatalf("Unable to call gRPC Calc() error: %v", err)
	}
	log.Printf("Server sended response  result: %v error:%v", r.GetResult(), r.GetError())
}
