package main

import (
	protos "github.com/ant-joshua/demo-grpc/invoicer"
	"google.golang.org/grpc"
	"net/http"
)

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))

	conn, err := grpc.Dial("localhost:9000")

	if err != nil {
		panic(err)
		//log.Fatalf("Failed to serve: %s", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	// create client
	invoiceClient := protos.NewInvoicerClient(conn)

}
