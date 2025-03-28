package router

import (
	"service/internal/delivery/handlers"
	"service/internal/repository"
	"service/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// Repositorios
	productRepo := repository.NewProductRepository(db)
	supplierRepo := repository.NewSupplierRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	purchaseRepo := repository.NewPurchaseRepository(db)
	purchaseDetailRepo := repository.NewPurchaseDetailRepository(db)
	saleRepo := repository.NewSaleRepository(db)
	saleDetailRepo := repository.NewSaleDetailRepository(db)

	// Servicios
	productSvc := service.NewProductService(productRepo)
	supplierSvc := service.NewSupplierService(supplierRepo)
	categorySvc := service.NewCategoryService(categoryRepo)
	purchaseSvc := service.NewPurchaseService(purchaseRepo, purchaseDetailRepo)
	purchaseDetailSvc := service.NewPurchaseDetailService(purchaseDetailRepo)
	saleSvc := service.NewSaleService(saleRepo, saleDetailRepo, productRepo)
	saleDetailSvc := service.NewSaleDetailService(saleDetailRepo)

	// Handlers
	productHandler := handlers.NewProductHandler(productSvc)
	supplierHandler := handlers.NewSupplierHandler(supplierSvc)
	categoryHandler := handlers.NewCategoryHandler(categorySvc)
	purchaseHandler := handlers.NewPurchaseHandler(purchaseSvc)
	purchaseDetailHandler := handlers.NewPurchaseDetailHandler(purchaseDetailSvc)
	saleHandler := handlers.NewSaleHandler(saleSvc)
	saleDetailHandler := handlers.NewSaleDetailHandler(saleDetailSvc)

	// Grupo de rutas para Product
	productGroup := router.Group("/products")
	{
		productGroup.POST("/", productHandler.CreateProduct)
		productGroup.GET("/:id", productHandler.GetProductByID)
		productGroup.GET("/", productHandler.GetAllProducts)
		productGroup.PUT("/:id", productHandler.UpdateProduct)
		productGroup.DELETE("/:id", productHandler.DeleteProduct)
	}

	// Grupo de rutas para Supplier
	supplierGroup := router.Group("/suppliers")
	{
		supplierGroup.POST("/", supplierHandler.CreateSupplier)
		supplierGroup.GET("/:id", supplierHandler.GetSupplierByID)
		supplierGroup.GET("/", supplierHandler.GetAllSuppliers)
		supplierGroup.PUT("/:id", supplierHandler.UpdateSupplier)
		supplierGroup.DELETE("/:id", supplierHandler.DeleteSupplier)
	}

	// Grupo de rutas para Category
	categoryGroup := router.Group("/categories")
	{
		categoryGroup.POST("/", categoryHandler.CreateCategory)
		categoryGroup.GET("/:id", categoryHandler.GetCategoryByID)
		categoryGroup.GET("/", categoryHandler.GetAllCategories)
		categoryGroup.PUT("/:id", categoryHandler.UpdateCategory)
		categoryGroup.DELETE("/:id", categoryHandler.DeleteCategory)
	}

	// Grupo de rutas para Purchase
	purchaseGroup := router.Group("/purchases")
	{
		purchaseGroup.POST("/", purchaseHandler.CreatePurchase)
		purchaseGroup.GET("/:id", purchaseHandler.GetPurchaseByID)
		purchaseGroup.GET("/", purchaseHandler.GetAllPurchases)
		purchaseGroup.PUT("/:id", purchaseHandler.UpdatePurchase)
		purchaseGroup.DELETE("/:id", purchaseHandler.DeletePurchase)
	}

	// Grupo de rutas para PurchaseDetail
	purchaseDetailGroup := router.Group("/purchase-details")
	{
		purchaseDetailGroup.GET("/:id", purchaseDetailHandler.GetPurchaseDetailByID)
		purchaseDetailGroup.GET("/purchases/:purchase_id", purchaseDetailHandler.GetAllByPurchaseID)
		purchaseDetailGroup.POST("/", purchaseDetailHandler.CreatePurchaseDetail)
		purchaseDetailGroup.PUT("/:id", purchaseDetailHandler.UpdatePurchaseDetail)
		purchaseDetailGroup.DELETE("/:id", purchaseDetailHandler.DeletePurchaseDetail)
	}

	// Grupo de rutas para Sale
	saleGroup := router.Group("/sales")
	{
		saleGroup.POST("/", saleHandler.CreateSale)
		saleGroup.GET("/:id", saleHandler.GetSaleByID)
		saleGroup.GET("/", saleHandler.GetAllSales)
		saleGroup.PUT("/:id", saleHandler.UpdateSale)
		saleGroup.DELETE("/:id", saleHandler.DeleteSale)
	}

	// Grupo de rutas para SaleDetail
	saleDetailGroup := router.Group("/sale-details")
	{
		saleDetailGroup.GET("/:id", saleDetailHandler.GetSaleDetailByID)
		saleDetailGroup.GET("/sales/:sale_id", saleDetailHandler.GetAllBySaleID)
		saleDetailGroup.POST("/", saleDetailHandler.CreateSaleDetail)
		saleDetailGroup.PUT("/:id", saleDetailHandler.UpdateSaleDetail)
		saleDetailGroup.DELETE("/:id", saleDetailHandler.DeleteSaleDetail)
	}
}
