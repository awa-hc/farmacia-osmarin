package handlers

import (
	"net/http"
	"strconv"

	"service/internal/domain/entities"
	"service/internal/service"

	"github.com/gin-gonic/gin"
)

type SupplierHandler struct {
	service service.SupplierService
}

func NewSupplierHandler(service service.SupplierService) *SupplierHandler {
	return &SupplierHandler{service: service}
}

func (h *SupplierHandler) CreateSupplier(ctx *gin.Context) {
	var supplier entities.Supplier
	if err := ctx.ShouldBindJSON(&supplier); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateSupplier(&supplier); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Proveedor creado exitosamente", "data": supplier})
}

func (h *SupplierHandler) GetSupplierByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	supplier, err := h.service.GetSupplierByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Proveedor no encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": supplier})
}

func (h *SupplierHandler) GetAllSuppliers(ctx *gin.Context) {
	suppliers, err := h.service.GetAllSuppliers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": suppliers})
}

func (h *SupplierHandler) UpdateSupplier(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var supplier entities.Supplier
	if err := ctx.ShouldBindJSON(&supplier); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	supplier.ID = uint(id)
	if err := h.service.UpdateSupplier(&supplier); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Proveedor actualizado exitosamente", "data": supplier})
}

func (h *SupplierHandler) DeleteSupplier(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.service.DeleteSupplier(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Proveedor eliminado exitosamente"})
}
