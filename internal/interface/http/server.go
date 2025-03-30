package http

import (
	"context"
	"dresser/internal/application/brands"
	"dresser/internal/application/productcollection"
	"dresser/internal/application/products"
	"dresser/internal/interface/http/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	BrandRegisterService brands.IBrandRegisterApplicationService
	BrandQueryService    brands.IBrandQueryApplicationService
	BrandDeleteService   brands.IBrandDeleteApplicationService

	ProductRegisterService products.IProductRegisterApplicationService
	ProductQueryService    products.IProductQueryApplicationService
	ProductDeleteService   products.IProductDeleteApplicationService
	ProductUpdateService   products.IProductUpdateApplicationService

	CategoriesQueryService productcollection.ICategoriesQueryApplicationService

	server *http.Server
	router *gin.Engine
}

type ServerConfig struct {
	Port                 string
	BrandRegisterService brands.IBrandRegisterApplicationService
	BrandQueryService    brands.IBrandQueryApplicationService
	BrandDeleteService   brands.IBrandDeleteApplicationService

	ProductRegisterService products.IProductRegisterApplicationService
	ProductQueryService    products.IProductQueryApplicationService
	ProductDeleteService   products.IProductDeleteApplicationService
	ProductUpdateService   products.IProductUpdateApplicationService

	CategoriesQueryService productcollection.ICategoriesQueryApplicationService
}

func NewHTTPServer(config ServerConfig) *Server {
	router := gin.Default()

	return &Server{
		BrandRegisterService: config.BrandRegisterService,
		BrandQueryService:    config.BrandQueryService,
		BrandDeleteService:   config.BrandDeleteService,

		ProductRegisterService: config.ProductRegisterService,
		ProductQueryService:    config.ProductQueryService,
		ProductDeleteService:   config.ProductDeleteService,
		ProductUpdateService:   config.ProductUpdateService,

		CategoriesQueryService: config.CategoriesQueryService,

		router: router,
		server: &http.Server{
			Addr:    ":" + config.Port,
			Handler: router,
		},
	}
}

func (s *Server) Start() error {
	brandHandler := handlers.NewBrandHandler(s.BrandRegisterService, s.BrandQueryService, s.BrandDeleteService)
	brandHandler.RegisterRoutes(s.router)

	productHandler := handlers.NewProductHandler(s.ProductRegisterService, s.ProductQueryService, s.ProductDeleteService, s.ProductUpdateService)
	productHandler.RegisterRoutes(s.router)

	categoriesHandler := handlers.NewCategoriesHandler(s.CategoriesQueryService)
	categoriesHandler.RegisterRoutes(s.router)

	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
