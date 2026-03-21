package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"storeapi/internal/api"
	"storeapi/internal/api/handlers"
	"storeapi/internal/config"
	"storeapi/internal/data/postgres"
	"syscall"
	"time"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		fmt.Println("load config error: ", err.Error())
	}
	dbConf := postgres.Config{
		Host:     conf.DBHost,
		Port:     conf.DBPort,
		DBName:   conf.DBName,
		User:     conf.DBUser,
		Password: conf.DBPassword,
		SSLMode:  conf.DBSSLMode,
	}
	db, err := postgres.NewPostgres(dbConf)
	if err != nil {
		fmt.Println("postgres connection error: ", err.Error())
	}
	defer db.Close()

	router := api.NewRouter(handlers.NewClientHandler(), handlers.NewSupplierHandler(), handlers.NewProductHandler(), handlers.NewImageHandler())

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", conf.ServerPort),
		Handler:      router.SetupRoutes(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		fmt.Printf("Server starting on port %s", conf.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Server failed: ", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Server forced to shutdown: ", err.Error())
	}

	fmt.Println("Server stopped")
}
