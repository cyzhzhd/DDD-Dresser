package handlers

import (
	"dresser/internal/application/productcollection"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoriesHandler struct {
	categoriesQueryService productcollection.ICategoriesQueryApplicationService
}

func NewCategoriesHandler(categoriesQueryService productcollection.ICategoriesQueryApplicationService) *CategoriesHandler {
	return &CategoriesHandler{
		categoriesQueryService: categoriesQueryService,
	}
}

func (h *CategoriesHandler) RegisterRoutes(r *gin.Engine) {
	g := r.Group("/api/categories")
	{
		g.GET("/lowest", h.GetLowestProducts)
		g.GET("/lowest/brand", h.GetLowestProductsByBrand)
		g.GET("/lohi/:category", h.GetLowestAndHighestProductsByCategories)
	}
}

// GetLowestProducts returns the lowest price product for each category
func (h *CategoriesHandler) GetLowestProducts(ctx *gin.Context) {
	lowestProducts, err := h.categoriesQueryService.GetLowestProducts(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get lowest products: %s", err.Error())})
		return
	}
	ctx.JSON(http.StatusOK, lowestProducts)
}

func (h *CategoriesHandler) GetLowestProductsByBrand(ctx *gin.Context) {
	lowestProductsByBrand, err := h.categoriesQueryService.GetLowestProductsByBrand(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get lowest products by brand: %s", err.Error())})
		return
	}
	ctx.JSON(http.StatusOK, lowestProductsByBrand)
}

func (h *CategoriesHandler) GetLowestAndHighestProductsByCategories(ctx *gin.Context) {
	category := ctx.Param("category")
	lohiProductsByCategory, err := h.categoriesQueryService.GetLowestAndHighestProductsByCategories(ctx, category)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get lowest and highest products by category: %s", err.Error())})
		return
	}
	ctx.JSON(http.StatusOK, lohiProductsByCategory)
}
