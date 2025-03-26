package handlers

import (
	"net/http"
	"strconv"

	"service/internal/domain/entities"
	"service/internal/service"

	"github.com/gin-gonic/gin"
)

type SaleHandler struct {
	service service.SaleService
}

func NewSaleHandler(service service.SaleService) *SaleHandler {
	return &SaleHandler{service: service}
}

// CreateSale crea una nueva venta
func (h *SaleHandler) CreateSale(ctx *gin.Context) {
	var sale entities.Sale
	if err := ctx.ShouldBindJSON(&sale); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	if err := h.service.CreateSale(&sale); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Venta creada exitosamente",
		"data":    sale,
	})
}

// GetSaleByID obtiene una venta por su ID
func (h *SaleHandler) GetSaleByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	sale, err := h.service.GetSaleByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Venta no encontrada"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": sale})
}

// GetAllSales obtiene todas las ventas
func (h *SaleHandler) GetAllSales(ctx *gin.Context) {
	sales, err := h.service.GetAllSales()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": sales})
}

// UpdateSale actualiza una venta existente
func (h *SaleHandler) UpdateSale(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var sale entities.Sale
	if err := ctx.ShouldBindJSON(&sale); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	sale.ID = uint(id)
	if err := h.service.UpdateSale(&sale); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Venta actualizada exitosamente",
		"data":    sale,
	})
}

// DeleteSale elimina una venta por su ID
func (h *SaleHandler) DeleteSale(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.service.DeleteSale(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Venta eliminada exitosamente"})
}
