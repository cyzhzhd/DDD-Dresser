package handlers

import (
	"dresser/internal/application/brands"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BrandHandler struct {
	brandRegisterService brands.IBrandRegisterApplicationService
	brandQueryService    brands.IBrandQueryApplicationService
	brandDeleteService   brands.IBrandDeleteApplicationService
}

func NewBrandHandler(brandRegisterService brands.IBrandRegisterApplicationService, brandQueryService brands.IBrandQueryApplicationService, brandDeleteService brands.IBrandDeleteApplicationService) *BrandHandler {
	return &BrandHandler{
		brandRegisterService: brandRegisterService,
		brandQueryService:    brandQueryService,
		brandDeleteService:   brandDeleteService,
	}
}

func (h *BrandHandler) RegisterRoutes(r *gin.Engine) {
	g := r.Group("/api/brands")
	{
		g.GET("", h.GetBrands)
		g.GET("/:id", h.GetBrand)
		g.POST("", h.CreateBrand)
		g.DELETE("/:id", h.DeleteBrand)
	}
}

func (h *BrandHandler) GetBrands(ctx *gin.Context) {
	fmt.Println("h.brandQueryService", h.brandQueryService)
	brands, err := h.brandQueryService.GetAll(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get brands: %s", err.Error())})
		return
	}
	ctx.JSON(http.StatusOK, brands)
}

func (h *BrandHandler) GetBrand(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid brand ID: %s", err.Error())})
		return
	}
	brand, err := h.brandQueryService.Get(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get brand: %s", err.Error())})
		return
	}
	ctx.JSON(http.StatusOK, brand)
}

// CreateBrandRequest represents the data needed to create a new brand
type CreateBrandRequest struct {
	Name       string   `json:"name" binding:"required"`
	Categories []string `json:"categories"`
}

func (h *BrandHandler) CreateBrand(ctx *gin.Context) {
	var req CreateBrandRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request body: %s", err.Error())})
		return
	}

	// The binding:"required" tag ensures the name is not empty
	// No need to check if Categories is empty as it's allowed

	cmd := brands.RegisterBrandCommand{
		Name:       req.Name,
		Categories: req.Categories,
	}

	brand, err := h.brandRegisterService.Register(ctx, cmd)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create brand: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusCreated, brand)
}

func (h *BrandHandler) DeleteBrand(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid brand ID: %s", err.Error())})
		return
	}
	err = h.brandDeleteService.Delete(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to delete brand: %s", err.Error())})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Brand deleted successfully", "id": id})
}
