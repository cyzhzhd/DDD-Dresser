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
