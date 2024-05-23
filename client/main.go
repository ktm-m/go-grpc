package main

import (
	"client/services"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
)

func main() {
	var cc *grpc.ClientConn
	var err error
	var creds credentials.TransportCredentials

	host := flag.String("host", "localhost:50051", "gRPC server host")
	tls := flag.Bool("tls", false, "use a secure TLS connection")
	flag.Parse()

	if *tls {
		certFile := "../tls/ca.crt"
		creds, err = credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		creds = insecure.NewCredentials() // use because of grpc require transport security
	}

	cc, err = grpc.NewClient(*host, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	calculatorClient := services.NewCalculatorClient(cc)
	calculatorService := services.NewCalculatorService(calculatorClient)

	// err = calculatorService.Hello("Oat")
	// err = calculatorService.Fibonacci(3)
	// err = calculatorService.Average(1, 2, 3, 4, 5)
	err = calculatorService.Sum(1, 2, 3, 4, 5)
	if err != nil {
		if grpcErr, ok := status.FromError(err); ok {
			// if err from gRPC ok will be true
			log.Printf("[%v] %v", grpcErr.Code(), grpcErr.Message())
		} else {
			log.Fatal(err)
		}
	}
}
