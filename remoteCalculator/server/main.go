package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/lao-tseu-is-alive/go-grpc-experiments/remoteCalculator/calc"
	"google.golang.org/grpc"
	"log"
	"net"
)

// calcServer is used to implement calc Server
type calcServer struct {
	pb.UnimplementedCalcServiceServer
}

func (s *calcServer) Calc(ctx context.Context, in *pb.CalcRequest) (*pb.CalcResponse, error) {
	log.Printf("Received: num1: %v , num2: %v ,  operation : %v", in.GetNum1(), in.GetNum2(), in.GetOperation())

	var res float64
	isError := false
	switch in.Operation {
	case pb.CalcRequest_Add:
		res = in.GetNum1() + in.GetNum2()
	case pb.CalcRequest_Subtract:
		res = in.GetNum1() - in.GetNum2()
	case pb.CalcRequest_Multiply:
		res = in.GetNum1() * in.GetNum2()
	case pb.CalcRequest_Divide:
		if in.GetNum2() == 0 {
			isError = true
			res = 0
		} else {
			res = in.GetNum1() / in.GetNum2()
		}
	default:
		res = in.GetNum1() + in.GetNum2()
	}

	return &pb.CalcResponse{
		Result: res,
		Error:  isError,
	}, nil
}

func main() {
	port := flag.Int("port", 50051, "The server port")
	flag.Parse()
	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//var opts []grpc.ServerOption
	//grpcServer := grpc.NewServer(opts...)
	grpcServer := grpc.NewServer()
	pb.RegisterCalcServiceServer(grpcServer, &calcServer{})
	log.Printf("gRPC Calc Server listening at %v", l.Addr())
	grpcServer.Serve(l)
}
