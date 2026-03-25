package main

import (
	"fmt"
	"storeapi/internal/di"

	"go.uber.org/fx"
)

// @title ShopAPI
// @version 1.0
// @description RESTful API for ShopAPI application
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /api/v1
// @schemes http https

func main() {
	(fx.New(di.Module)).Run()
	fmt.Println("Server stopped")
}
