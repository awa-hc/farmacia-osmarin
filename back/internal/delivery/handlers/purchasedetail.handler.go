package handlers

import (
	"net/http"
	"strconv"

	"service/internal/domain/entities"
	"service/internal/service"

	"github.com/gin-gonic/gin"
)

type PurchaseDetailHandler struct {
	service service.PurchaseDetailService
}

func NewPurchaseDetailHandler(service service.PurchaseDetailService) *PurchaseDetailHandler {
	return &PurchaseDetailHandler{service: service}
}

// GetPurchaseDetailByID obtiene un detalle de compra por su ID
func (h *PurchaseDetailHandler) GetPurchaseDetailByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	detail, err := h.service.GetPurchaseDetailByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Detalle de compra no encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": detail})
}

// GetAllByPurchaseID obtiene todos los detalles asociados a una compra específica
func (h *PurchaseDetailHandler) GetAllByPurchaseID(ctx *gin.Context) {
	purchaseID, err := strconv.ParseUint(ctx.Param("purchase_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de compra inválido"})
		return
	}

	details, err := h.service.GetAllByPurchaseID(uint(purchaseID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": details})
}

// CreatePurchaseDetail crea un nuevo detalle de compra
func (h *PurchaseDetailHandler) CreatePurchaseDetail(ctx *gin.Context) {
	var detail entities.PurchaseDetail
	if err := ctx.ShouldBindJSON(&detail); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	if err := h.service.CreatePurchaseDetail(&detail); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Detalle de compra creado exitosamente",
		"data":    detail,
	})
}

// UpdatePurchaseDetail actualiza un detalle de compra existente
func (h *PurchaseDetailHandler) UpdatePurchaseDetail(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var detail entities.PurchaseDetail
	if err := ctx.ShouldBindJSON(&detail); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	detail.ID = uint(id)
	if err := h.service.UpdatePurchaseDetail(&detail); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Detalle de compra actualizado exitosamente",
		"data":    detail,
	})
}

// DeletePurchaseDetail elimina un detalle de compra por su ID
func (h *PurchaseDetailHandler) DeletePurchaseDetail(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.service.DeletePurchaseDetail(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Detalle de compra eliminado exitosamente"})
}
