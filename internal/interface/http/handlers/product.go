package handlers

import (
	"dresser/internal/application/products"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productRegisterService products.IProductRegisterApplicationService
	productQueryService    products.IProductQueryApplicationService
	productDeleteService   products.IProductDeleteApplicationService
	productUpdateService   products.IProductUpdateApplicationService
}

func NewProductHandler(
	productRegisterService products.IProductRegisterApplicationService,
	productQueryService products.IProductQueryApplicationService,
	productDeleteService products.IProductDeleteApplicationService,
	productUpdateService products.IProductUpdateApplicationService,
) *ProductHandler {
	return &ProductHandler{
		productRegisterService: productRegisterService,
		productQueryService:    productQueryService,
		productDeleteService:   productDeleteService,
		productUpdateService:   productUpdateService,
	}
}

func (h *ProductHandler) RegisterRoutes(r *gin.Engine) {
	g := r.Group("/api/products")
	{
		g.GET("", h.GetProducts)
		g.GET("/:id", h.GetProduct)
		g.GET("/brand/:brandId", h.GetProductsByBrand)
		g.GET("/category/:category", h.GetProductsByCategory)
		g.POST("", h.CreateProduct)
		g.PUT("/:id", h.UpdateProduct)
		g.DELETE("/:id", h.DeleteProduct)
	}
}

func (h *ProductHandler) GetProducts(ctx *gin.Context) {
	products, err := h.productQueryService.GetAll(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get products: %s", err.Error())})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProduct(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid product ID: %s", err.Error())})
		return
	}
	product, err := h.productQueryService.Get(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get product: %s", err.Error())})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (h *ProductHandler) GetProductsByBrand(ctx *gin.Context) {
	brandIdParam := ctx.Param("brandId")
	brandId, err := strconv.Atoi(brandIdParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid brand ID: %s", err.Error())})
		return
	}
	products, err := h.productQueryService.GetByBrand(ctx, brandId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get products by brand: %s", err.Error())})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductsByCategory(ctx *gin.Context) {
	category := ctx.Param("category")
	products, err := h.productQueryService.GetByCategory(ctx, category)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get products by category: %s", err.Error())})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

// CreateProductRequest represents the data needed to create a new product
type CreateProductRequest struct {
	Name     string `json:"name" binding:"required"`
	BrandID  int    `json:"brand_id" binding:"required"`
	Category string `json:"category" binding:"required"`
	Amount   int    `json:"amount" binding:"required"`
	Currency string `json:"currency" binding:"required"`
	Size     string `json:"size" binding:"required"`
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var req CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request body: %s", err.Error())})
		return
	}
	fmt.Println("req", req.Name, req.BrandID, req.Category, req.Amount, req.Currency, req.Size)

	cmd := products.RegisterProductCommand{
		Name:     req.Name,
		BrandID:  req.BrandID,
		Category: req.Category,
		Amount:   req.Amount,
		Currency: req.Currency,
		Size:     req.Size,
	}

	product, err := h.productRegisterService.Register(ctx, cmd)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create product: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

// UpdateProductRequest represents the data needed to update a product
type UpdateProductRequest struct {
	Name     string `json:"name" binding:"required"`
	BrandID  int    `json:"brand_id" binding:"required"`
	Category string `json:"category" binding:"required"`
	Amount   int    `json:"amount" binding:"required"`
	Currency string `json:"currency" binding:"required"`
	Size     string `json:"size" binding:"required"`
}

func (h *ProductHandler) UpdateProduct(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid product ID: %s", err.Error())})
		return
	}

	var req UpdateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request body: %s", err.Error())})
		return
	}

	cmd := products.UpdateProductCommand{
		ID:       id,
		Name:     req.Name,
		BrandID:  req.BrandID,
		Category: req.Category,
		Amount:   req.Amount,
		Currency: req.Currency,
		Size:     req.Size,
	}

	product, err := h.productUpdateService.Update(ctx, cmd)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update product: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid product ID: %s", err.Error())})
		return
	}
	err = h.productDeleteService.Delete(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to delete product: %s", err.Error())})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully", "id": id})
}
