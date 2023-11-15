package main

import (
	"github.com/ant-joshua/demo-grpc/auth/handler"
	"github.com/ant-joshua/demo-grpc/auth/service"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	jwtService := service.NewJwtService()

	authController := handler.NewAuthController(jwtService)
	authController.Route(e)

	e.Logger.Fatal(e.Start(":8002"))
}
