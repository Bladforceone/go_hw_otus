package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Bladforceone/go_hw_otus/hw15_go_sql/internal/db"
	"github.com/Bladforceone/go_hw_otus/hw15_go_sql/internal/repository"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Repo *repository.Repository
}

func NewProductHandler(rep *repository.Repository) *ProductHandler {
	return &ProductHandler{Repo: rep}
}

func (h *ProductHandler) ProductCreate(c *gin.Context) {
	input := db.Product{}
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.Repo.ProductCreate(context.Background(), input.Name, input.Price, input.Stock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid product ID"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32) // 10 - система счисления, 32 - битность
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := h.Repo.ProductGet(context.Background(), int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}
