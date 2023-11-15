package main

import (
	protos "github.com/ant-joshua/demo-grpc/invoicer"
	"github.com/ant-joshua/demo-grpc/product/models"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

type ProductController struct {
	invoice protos.InvoicerClient
}

func NewProductController(invoice protos.InvoicerClient) *ProductController {
	return &ProductController{invoice: invoice}
}

func (p *ProductController) CreateInvoice(c echo.Context) error {
	//var createProductRequest *data.CreateProductRequest

	createProductRequest := new(models.CreateProductRequest)

	err := c.Bind(createProductRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	var amount int64 = 0

	amount = int64(createProductRequest.Price)

	// create invoice request
	invoiceRequest := &protos.CreateRequest{
		Amount: &protos.Amount{
			Amount:   amount,
			Currency: "joshua",
		},
		From:      "Ant Joshua",
		To:        "Joshua",
		VATNumber: "",
	}

	// call invoice service
	invoiceResponse, err := p.invoice.Create(c.Request().Context(), invoiceRequest)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	response := map[string]interface{}{
		"invoice": invoiceResponse,
		"product": createProductRequest,
	}

	return c.JSON(http.StatusOK, response)
}

func (p *ProductController) Route(e *echo.Echo) {
	e.POST("/product/invoice", p.CreateInvoice)
}

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

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	productController := NewProductController(invoiceClient)
	productController.Route(e)

	e.Logger.Fatal(e.Start(":8000"))

}
