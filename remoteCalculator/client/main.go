package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/lao-tseu-is-alive/go-grpc-experiments/remoteCalculator/calc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"os"
	"time"
)

const (
	VERSION              = "0.1.1"
	APP                  = "go-grpc-calcServer"
	baseSslPath          = "openssl_certificates/"
	defaultPort          = 50051
	defaultServerIp      = "localhost"
	useTLS               = true
	certificateAuthority = "ca.crt"
	defaultTimeoutSecond = 3
	defaultAuthTOKEN     = "14d3d4de9273fb1b5cc33e6bae2225f115885d028ddfb82555172188f9dee91c"
)

// main is client to call the gRPC calc server, just try this from your shell :  go run remoteCalculator/client/main.go -num1 2 -num2 2 -operation M
func main() {
	port := flag.Int("port", defaultPort, "the address to connect to")
	serverIP := flag.String("server-ip", defaultServerIp, "The server Ip address")
	num1 := flag.Float64("num1", 0, "first operand number")
	num2 := flag.Float64("num2", 0, "second operand number")
	op := flag.String("operation", "Add", "Operation to apply to the numbers [Add|Subtract|Multiply|Divide]")
	flag.Parse()
	l := log.New(os.Stdout, fmt.Sprintf("%s ", APP), log.Ldate|log.Ltime|log.Lshortfile)

	listenAddr := fmt.Sprintf("%s:%d", *serverIP, *port)
	l.Printf("INFO: 'Starting %s version:%s. Will try to connect to %s'", APP, VERSION, listenAddr)
	var opts []grpc.DialOption
	if useTLS {
		certAuth := baseSslPath + certificateAuthority
		l.Printf("INFO: 'TLS enabled CA certAuth: %s'", certAuth)
		sslCredentials, err := credentials.NewClientTLSFromFile(certAuth, "")
		if err != nil {
			l.Printf("INFO: for TLS to work you should have your certificates in directory : %s", baseSslPath)
			l.Printf("INFO: you can generate self-signed certificates using : scripts/make_ssl_certificates.sh")
			l.Fatalf("💥💥 error loading CA certificate: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(sslCredentials))
	} else {
		l.Println("INFO: TLS is disabled !")
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	// Set up a connection to the server allowing insecure (no ssl).
	conn, err := grpc.Dial(listenAddr, opts...)
	if err != nil {
		l.Fatalf("💥💥 ERROR: unable to connect doing grpc.Dial(%s) : %v", listenAddr, err)
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
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeoutSecond*time.Second)
	defer cancel()
	// Anything linked to this variable will transmit request headers.
	md := metadata.New(map[string]string{"x-request-id": "req-123"})
	ctx = metadata.NewOutgoingContext(ctx, md)

	// Anything linked to this variable will fetch response headers.
	var header metadata.MD
	l.Printf("INFO: will send gRPC request Calc(num1: %.3f , num2: %.3f, Operation :%s)", *num1, *num2, calcOp.String())
	r, err := c.Calc(ctx, &pb.CalcRequest{
		Num1:      *num1,
		Num2:      *num2,
		Operation: calcOp,
	},
		grpc.Header(&header))
	if err != nil {
		l.Fatalf("💥💥 ERROR: failed call gRPC Calc() : %v", err)
	}
	l.Printf("INFO: server response is result: %v  with error:%v", r.GetResult(), r.GetError())
}
