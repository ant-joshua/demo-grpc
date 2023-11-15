package handler

import (
	"github.com/ant-joshua/demo-grpc/auth/data"
	"github.com/ant-joshua/demo-grpc/auth/service"
	"github.com/labstack/echo/v4"
	"log"
	"strings"
)

type AuthController struct {
	jwtService *service.JwtService
}

func NewAuthController(jwtService *service.JwtService) *AuthController {
	return &AuthController{
		jwtService: jwtService,
	}
}

func (a *AuthController) Route(e *echo.Echo) {
	e.POST("/login", a.Login)
	e.GET("/verify", a.ManualVerify)
}

func (a *AuthController) Login(ctx echo.Context) error {

	request := new(data.LoginRequest)

	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(400, err)
	}

	if request.Username != "antoniusjoshua47@gmail.com" || request.Password != "joshua" {
		return ctx.JSON(401, map[string]interface{}{
			"message": "invalid username or password",
		})
	}

	user := data.User{
		UserID:   1,
		Email:    request.Username,
		Name:     "Joshua",
		Password: "",
		Admin:    true,
	}

	// Check username and password from database here.

	// Generate encoded token and send it as response.
	t, err := a.jwtService.GenerateToken(user.Name, user.Email, user.Admin)

	if err != nil {
		return err
	}

	return ctx.JSON(200, map[string]interface{}{
		"token": t,
	})
}

func (a *AuthController) ManualVerify(ctx echo.Context) error {

	header := ctx.Request().Header.Get("Authorization")

	splitString := strings.Split(header, " ")

	if len(splitString) != 2 {
		return ctx.JSON(401, map[string]interface{}{
			"message": "invalid token",
		})
	}

	tokenString := splitString[1]

	log.Println(tokenString)

	jwtService := service.NewJwtService()

	_, err := jwtService.VerifyToken(tokenString)

	if err != nil {
		return ctx.JSON(401, err)
	}

	return ctx.JSON(200, map[string]interface{}{
		"message": "valid",
	})
}
