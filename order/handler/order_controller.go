package handler

import (
	"fmt"
	protos "github.com/ant-joshua/demo-grpc/invoicer"
	"github.com/ant-joshua/demo-grpc/order/data"
	"github.com/labstack/echo/v4"
	"time"
)

type OrderController struct {
	invoice protos.InvoicerClient
}

func NewOrderController(invoice protos.InvoicerClient) *OrderController {
	return &OrderController{
		invoice: invoice,
	}
}

func (controller *OrderController) Route(e *echo.Echo) {
	e.GET("/orders", controller.GetOrderList)
	e.POST("/orders", controller.CreateOrder)
	e.GET("/orders/invoice", controller.GetOrderInvoice)
}

func (controller *OrderController) GetOrderList(ctx echo.Context) error {

	return ctx.JSON(200, data.OrderListData)
}

func (controller *OrderController) CreateOrder(ctx echo.Context) error {

	createOrderRequest := new(data.CreateOrderRequest)

	err := ctx.Bind(createOrderRequest)

	if err != nil {
		return ctx.JSON(400, err)
	}
	orderList := data.OrderListData

	var orderDetailList []data.OrderDetail

	var orderID int = len(orderList) + 1

	total := 0

	for _, detailRequest := range createOrderRequest.Details {

		product := data.OrderProduct{
			ProductID: detailRequest.Product.ProductID,
			Name:      detailRequest.Product.Name,
			Price:     detailRequest.Product.Price,
		}

		subTotal := product.Price * detailRequest.Qty

		orderDetail := data.OrderDetail{
			OrderDetailID: len(orderDetailList) + 1,
			OrderID:       orderID,
			Product:       product,
			Qty:           detailRequest.Qty,
			SubTotal:      subTotal,
		}

		orderDetailList = append(orderDetailList, orderDetail)

		total += subTotal
	}

	createdAt := time.Now()

	order := data.Order{
		OrderID:     orderID,
		OrderNumber: createOrderRequest.OrderNumber,
		Details:     orderDetailList,
		Total:       total,
		CreatedAt:   &createdAt,
	}

	orderList = append(orderList, &order)

	customer := data.Customer{
		Name:  "Antonius Joshua",
		Email: "antoniusjoshua47@gmail.com",
	}
	// create invoice
	amount := int64(order.Total)
	// create invoice request
	invoiceRequest := &protos.CreateRequest{
		Amount: &protos.Amount{
			Amount:   amount,
			Currency: "IDR",
		},
		From:      "Vultra",
		To:        fmt.Sprintf("%s - %s", customer.Name, customer.Email),
		VATNumber: "",
	}

	createdInvoice, err := controller.invoice.Create(ctx.Request().Context(), invoiceRequest)
	if err != nil {
		return ctx.JSON(500, "Failed to create invoice")
	}

	return ctx.JSON(200, map[string]interface{}{
		"data":    orderList,
		"invoice": createdInvoice,
	})
}

func (controller *OrderController) GetOrderInvoice(ctx echo.Context) error {

	return nil
}
