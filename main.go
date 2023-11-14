package main

import (
	"context"
	"fmt"
	"github.com/ant-joshua/demo-grpc/invoicer"
	"google.golang.org/grpc"
	"log"
	"net"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s *myInvoicerServer) Create(ctx context.Context, request *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {

	fmt.Println("Create invoicer", request)

	return &invoicer.CreateResponse{
		Pdf:  []byte(request.From),
		Docx: []byte("test"),
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)

	}

	grpcServer := grpc.NewServer()

	invoicerService := &myInvoicerServer{}
	invoicer.RegisterInvoicerServer(grpcServer, invoicerService)

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to serve grpc server: %s", err)
	}
}
