package main

import (
	authService "github.com/ant-joshua/demo-grpc/auth/service"
	protos "github.com/ant-joshua/demo-grpc/invoicer"
	"github.com/ant-joshua/demo-grpc/order/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to serve: %s", err)
		//panic(err)

	}

	protectedGroup := e.Group("")

	registerAuthService := authService.NewJwtService()

	protectedGroup.Use(registerAuthService.JwtMiddleware)
	//protectedGroup.Use(handler.JwtVerify)

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	// create client
	invoiceClient := protos.NewInvoicerClient(conn)

	orderController := handler.NewOrderController(invoiceClient)
	orderController.Route(e, protectedGroup)

	e.Logger.Fatal(e.Start(":8001"))
}
