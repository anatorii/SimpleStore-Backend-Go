package main

import (
	"fmt"
	"storeapi/internal/di"

	"go.uber.org/fx"
)

func main() {
	(fx.New(di.Module)).Run()

	fmt.Println("Server stopped")
}
