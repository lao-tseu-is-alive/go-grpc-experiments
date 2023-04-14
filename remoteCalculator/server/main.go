package main

import (
	"context"
	"crypto/rand"
	"flag"
	"fmt"
	pb "github.com/lao-tseu-is-alive/go-grpc-experiments/remoteCalculator/calc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
)

const (
	VERSION          = "0.1.1"
	APP              = "go-grpc-calcServer"
	baseSslPath      = "openssl_certificates/"
	defaultPort      = 50051
	defaultServerIp  = "localhost"
	useTLS           = true
	serverCrt        = "localhost_server.crt"
	serverPem        = "localhost_server.pem"
	defaultTokenSize = 32
)

// calcServer is used to implement calc gRPC Server
type calcServer struct {
	pb.UnimplementedCalcServiceServer
	l     *log.Logger
	token string
}

// Calc handles the gRPC CalcRequest by doing basic math operation on num1 and num2, it then returns the result in CalcResponse
func (s *calcServer) Calc(ctx context.Context, in *pb.CalcRequest) (*pb.CalcResponse, error) {
	s.l.Printf("TRACE: entering Calc(num1: %v , num2: %v ,  operation : %v)", in.GetNum1(), in.GetNum2(), in.GetOperation())

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

func TokenGenerator(l int) string {
	b := make([]byte, l)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// LogInterceptor allows to intercept any request done to the gRPC server
func LogInterceptor(l *log.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		l.Printf("TRACE: entering LogInterceptor() received request:\n%+v\n", req)

		headers, ok := metadata.FromIncomingContext(ctx)
		if ok {
			l.Printf("# LogInterceptor() Received headers: \n%+v\n", headers)
		}
		return handler(ctx, req)
	}
}

func CheckHeaderInterceptor(l *log.Logger, token string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		l.Printf("TRACE: entering CheckHeaderInterceptor() ctx:\n %+v\n", ctx)
		headers, ok := metadata.FromIncomingContext(ctx)

		if !ok {
			l.Printf("🔥🔥 Error: reading metadata.FromIncomingContext(ctx:\n %+v)\n", ctx)
			return nil, status.Error(codes.Internal, "Error while reading the context")
		}

		if len(headers.Get("authorization")) == 0 {
			l.Println("🤔🤔 WARN: in CheckHeaderInterceptor did not find authorization header !")
			// return nil, status.Error(codes.Unauthenticated, "Expected authorization header")
			//for now just accept this unauthorized call :-)
			return handler(ctx, req)
		}

		//TODO : also add check for valid token when authorization header was found

		return handler(ctx, req)
	}
}

func main() {
	port := flag.Int("port", defaultPort, "The server port")
	serverIP := flag.String("server-ip", defaultServerIp, "The server Ip address")
	flag.Parse()
	l := log.New(os.Stdout, fmt.Sprintf("%s ", APP), log.Ldate|log.Ltime|log.Lshortfile)
	listenAddr := fmt.Sprintf("%s:%d", *serverIP, *port)
	l.Printf("INFO: 'Starting %s version:%s server on %s'", APP, VERSION, listenAddr)

	rootToken := TokenGenerator(defaultTokenSize)

	netListener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		l.Fatalf("💥💥 error doing net.Listen(%s) : %v", listenAddr, err)
	}
	var opts []grpc.ServerOption
	if useTLS {
		certFile := baseSslPath + serverCrt
		l.Printf("INFO: 'TLS enabled, certFile: %s'", certFile)
		keyFile := baseSslPath + serverPem
		l.Printf("INFO: 'TLS enabled  keyFile: %s'", keyFile)
		sslCredentials, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			l.Printf("INFO: for TLS to work you should have your certificates in directory : %s", baseSslPath)
			l.Printf("INFO: you can generate self-signed certificates using : scripts/make_ssl_certificates.sh")
			l.Fatalf("💥💥 error loading certificates: %v", err)
		}
		opts = append(opts, grpc.Creds(sslCredentials))
	}

	opts = append(opts, grpc.ChainUnaryInterceptor(LogInterceptor(l), CheckHeaderInterceptor(l, rootToken)))

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCalcServiceServer(grpcServer, &calcServer{
		UnimplementedCalcServiceServer: pb.UnimplementedCalcServiceServer{},
		l:                              l,
		token:                          rootToken,
	})
	l.Printf("INFO: 'will start server %s'", listenAddr)
	err = grpcServer.Serve(netListener)
	if err != nil {
		l.Fatalf("💥💥 error doing grpcServer.Serve : %v", err)
	}
}
