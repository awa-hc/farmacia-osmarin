package handlers

import (
	"net/http"
	"strconv"

	"service/internal/domain/entities"
	"service/internal/service"

	"github.com/gin-gonic/gin"
)

type SaleDetailHandler struct {
	service service.SaleDetailService
}

func NewSaleDetailHandler(service service.SaleDetailService) *SaleDetailHandler {
	return &SaleDetailHandler{service: service}
}

// GetSaleDetailByID obtiene un detalle de venta por su ID
func (h *SaleDetailHandler) GetSaleDetailByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	detail, err := h.service.GetSaleDetailByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Detalle de venta no encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": detail})
}

// GetAllBySaleID obtiene todos los detalles asociados a una venta específica
func (h *SaleDetailHandler) GetAllBySaleID(ctx *gin.Context) {
	saleID, err := strconv.ParseUint(ctx.Param("sale_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de venta inválido"})
		return
	}

	details, err := h.service.GetAllBySaleID(uint(saleID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": details})
}

// CreateSaleDetail crea un nuevo detalle de venta
func (h *SaleDetailHandler) CreateSaleDetail(ctx *gin.Context) {
	var detail entities.SaleDetail
	if err := ctx.ShouldBindJSON(&detail); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	if err := h.service.CreateSaleDetail(&detail); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Detalle de venta creado exitosamente",
		"data":    detail,
	})
}

// UpdateSaleDetail actualiza un detalle de venta existente
func (h *SaleDetailHandler) UpdateSaleDetail(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var detail entities.SaleDetail
	if err := ctx.ShouldBindJSON(&detail); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	detail.ID = uint(id)
	if err := h.service.UpdateSaleDetail(&detail); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Detalle de venta actualizado exitosamente",
		"data":    detail,
	})
}

// DeleteSaleDetail elimina un detalle de venta por su ID
func (h *SaleDetailHandler) DeleteSaleDetail(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.service.DeleteSaleDetail(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Detalle de venta eliminado exitosamente"})
}
