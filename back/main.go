package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Modelo de Producto
type Producto struct {
	ID           uint    `gorm:"primaryKey" json:"id"`
	Producto     string  `gorm:"not null" json:"producto"`
	Presentacion string  `json:"presentacion"`
	Stock        int     `json:"stock"`
	Precio       float64 `json:"precio"`
	Vencimiento  string  `json:"vencimiento"`
}

// Modelo de Venta
type Venta struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	ProductoID uint      `json:"producto_id"`
	Cantidad   int       `json:"cantidad"`
	Total      float64   `json:"total"`
	Fecha      time.Time `json:"fecha"`
}

var db *gorm.DB

func initDB() {
	dsn := os.Getenv("DATABASE_URL")
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error al conectar con la base de datos")
	}
	db.AutoMigrate(&Producto{}, &Venta{})
}

func getProductos(c *gin.Context) {
	var productos []Producto
	db.Find(&productos)
	c.JSON(http.StatusOK, productos)
}

func addProducto(c *gin.Context) {
	var producto Producto
	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&producto)
	c.JSON(http.StatusCreated, producto)
}

func venderProducto(c *gin.Context) {
	var venta Venta
	if err := c.ShouldBindJSON(&venta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var producto Producto
	db.First(&producto, venta.ProductoID)
	if producto.ID == 0 || producto.Stock < venta.Cantidad {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Stock insuficiente"})
		return
	}

	venta.Total = float64(venta.Cantidad) * producto.Precio
	venta.Fecha = time.Now()
	producto.Stock -= venta.Cantidad
	db.Save(&producto)
	db.Create(&venta)
	c.JSON(http.StatusOK, venta)
}

func getVentas(c *gin.Context) {
	var ventas []Venta
	db.Find(&ventas)
	c.JSON(http.StatusOK, ventas)
}

func main() {
	initDB()
	r := gin.Default()
	r.GET("/productos", getProductos)
	r.POST("/productos", addProducto)
	r.POST("/vender", venderProducto)
	r.GET("/ventas", getVentas)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
