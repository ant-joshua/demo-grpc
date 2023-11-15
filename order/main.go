package main

import (
	protos "github.com/ant-joshua/demo-grpc/invoicer"
	"github.com/ant-joshua/demo-grpc/order/handler"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	e := echo.New()

	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to serve: %s", err)
		//panic(err)

	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	// create client
	invoiceClient := protos.NewInvoicerClient(conn)

	orderController := handler.NewOrderController(invoiceClient)
	orderController.Route(e)

	e.Logger.Fatal(e.Start(":8001"))
}
