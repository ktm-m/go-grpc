package services

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"time"
)

type CalculatorService interface {
	Hello(name string) error
	Fibonacci(n uint32) error
	Average(numbers ...float64) error
	Sum(numbers ...int32) error
}

type calculatorService struct {
	calculatorClient CalculatorClient
}

func NewCalculatorService(calculatorClient CalculatorClient) CalculatorService {
	return calculatorService{calculatorClient: calculatorClient}
}

func (s calculatorService) Hello(name string) error {
	req := HelloRequest{
		Name:        name,
		CreatedDate: timestamppb.Now(),
	}

	res, err := s.calculatorClient.Hello(context.Background(), &req)
	if err != nil {
		return err
	}

	fmt.Printf("Service: Hello\n")
	fmt.Printf("Request: %v\n", req.Name)
	fmt.Printf("Response: %v\n", res.Result)

	return nil
}

func (s calculatorService) Fibonacci(n uint32) error {
	req := FibonacciRequest{
		N: n,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // set timeout
	defer cancel()

	stream, err := s.calculatorClient.Fibonacci(ctx, &req)
	if err != nil {
		return err
	}

	fmt.Printf("Service: Fibonacci\n")
	fmt.Printf("Request: %v\n", req.N)
	for {
		res, err := stream.Recv()
		if err == io.EOF { // if stream end
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("Response: %v\n", res.Result)
	}

	return nil
}

func (s calculatorService) Average(numbers ...float64) error {
	stream, err := s.calculatorClient.Average(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("Service: Average\n")
	for _, number := range numbers {
		req := AverageRequest{
			Number: number,
		}
		stream.Send(&req)
		fmt.Printf("Request: %v\n", req.Number)
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	fmt.Printf("Response: %v\n", res.Result)
	return nil
}

func (s calculatorService) Sum(numbers ...int32) error {
	stream, err := s.calculatorClient.Sum(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("Service: Sum\n")
	go func() {
		for _, number := range numbers {
			req := SumRequest{
				Number: number,
			}
			stream.Send(&req)
			fmt.Printf("Request: %v\n", req.Number)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	done := make(chan bool)
	errs := make(chan error)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				errs <- err
			}
			fmt.Printf("Response: %v\n", res.Result)
		}
		done <- true
	}()

	select {
	case <-done:
		return nil
	case err := <-errs:
		return err
	}
}
