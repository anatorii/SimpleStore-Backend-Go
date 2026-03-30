package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "storeapi/docs"
	"storeapi/internal/api/handlers"
)

type Router struct {
	ClientHandler   *handlers.ClientHandler
	SupplierHandler *handlers.SupplierHandler
	ProductHandler  *handlers.ProductHandler
	ImageHandler    *handlers.ImageHandler
}

func NewRouter(clientHandler *handlers.ClientHandler, supplierHandler *handlers.SupplierHandler, productHandler *handlers.ProductHandler, imageHandler *handlers.ImageHandler) *Router {
	return &Router{
		ClientHandler:   clientHandler,
		SupplierHandler: supplierHandler,
		ProductHandler:  productHandler,
		ImageHandler:    imageHandler,
	}
}

func (r *Router) SetupRoutes() http.Handler {
	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60))
	router.Use(middleware.Compress(5))
	// router.Use(middleware.LoggerMiddleware)

	// CORS
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Health check
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Swagger UI
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("list"),
		httpSwagger.DomID("swagger-ui"),
	))

	// OpenAPI spec JSON endpoint
	router.Get("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "docs/swagger.json")
	})

	// API routes
	router.Route("/api/v1", func(chr chi.Router) {
		// client routes
		chr.Route("/clients", func(ch chi.Router) {
			ch.Get("/", r.ClientHandler.GetAllClients)
			ch.Get("/{name}/{surname}", r.ClientHandler.GetClientByName)
			ch.Post("/", r.ClientHandler.CreateClient)
			ch.Patch("/{id}", r.ClientHandler.UpdateClientAddress)
			ch.Delete("/{id}", r.ClientHandler.DeleteClient)
		})

		// product routes
		chr.Route("/products", func(ch chi.Router) {
			ch.Get("/", r.ProductHandler.GetAllProducts)
			ch.Get("/{id}", r.ProductHandler.GetProductById)
			ch.Post("/", r.ProductHandler.CreateProduct)
			ch.Patch("/{id}/available", r.ProductHandler.UpdateProductAvailable)
			ch.Delete("/{id}", r.ProductHandler.DeleteProduct)
		})

		// supplier routes
		chr.Route("/suppliers", func(ch chi.Router) {
			ch.Get("/", r.SupplierHandler.GetAllSuppliers)
			ch.Get("/{id}", r.SupplierHandler.GetSupplierById)
			ch.Post("/", r.SupplierHandler.CreateSupplier)
			ch.Patch("/{id}", r.SupplierHandler.UpdateSupplierAddress)
			ch.Delete("/{id}", r.SupplierHandler.DeleteSupplier)
		})

		// image routes
		chr.Route("/images", func(ch chi.Router) {
			ch.Get("/{id}", r.ImageHandler.GetImageById)
			ch.Get("/product/{id}", r.ImageHandler.GetProductImageById)
			ch.Post("/", r.ImageHandler.CreateImage)
			ch.Patch("/{id}", r.ImageHandler.UpdateImage)
			ch.Delete("/{id}", r.ImageHandler.DeleteImage)
		})
	})

	return router
}
