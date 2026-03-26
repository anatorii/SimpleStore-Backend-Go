package di

import (
	"context"
	"fmt"
	"net/http"
	"storeapi/internal/api"
	"storeapi/internal/api/handlers"
	"storeapi/internal/config"
	"storeapi/internal/data/postgres"
	"storeapi/internal/data/repository"
	"storeapi/internal/service"
	"time"

	"go.uber.org/fx"
)

var Module = fx.Module("shop",
	fx.Provide(
		NewHttp,
		api.NewRouter,
		handlers.NewClientHandler,
		handlers.NewSupplierHandler,
		handlers.NewProductHandler,
		handlers.NewImageHandler,
		service.NewClientService,
		service.NewSupplierService,
		service.NewProductService,
		service.NewImageService,
		repository.NewClientRepo,
		repository.NewImageRepo,
		repository.NewProductRepo,
		repository.NewSupplierRepo,
		config.NewConfig,
		postgres.NewPostgres,
		postgres.NewPostgresConfig,
	),
	fx.Invoke(
		RegisterHooks,
	),
)

func NewHttp(router *api.Router, conf *config.Config) *http.Server {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", conf.ServerPort),
		Handler:      router.SetupRoutes(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	return server
}

func RegisterHooks(lc fx.Lifecycle, server *http.Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				fmt.Printf("Server starting on port %s\n", server.Addr)
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					fmt.Println("Server failed: ", err.Error())
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Shutting down server...")
			return server.Shutdown(ctx)
		},
	})
}
