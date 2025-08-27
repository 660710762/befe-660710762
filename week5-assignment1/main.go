// main.go
package main

import (
    "net/http"
	"strings"
    "github.com/gin-gonic/gin"
)

// CosmeticProduct struct
type CosmeticProduct struct {
    ID           string  `json:"id"`
    Name         string  `json:"name"`
    Brand        string  `json:"brand"`
    Category     string  `json:"category"` // เช่น "lipstick", "cushion", "skincare"
    Price        float64 `json:"price"`
    InStock      bool    `json:"in_stock"`
}

// In-memory database (ในโปรเจคจริงใช้ database)
var products = []CosmeticProduct{
    {ID: "1", Name: "Matte Lipstick", Brand: "4U2", Category: "lipstick", Price: 299.00, InStock: true},
    {ID: "2", Name: "Flawless Cushion", Brand: "CHY", Category: "cushion", Price: 429.00, InStock: false},
}

func getProducts(c *gin.Context) {
	categoryQuery := c.Query("category")

	if categoryQuery != "" {
		filter := []CosmeticProduct{}
		for _, product := range products {
			if strings.ToLower(product.Category) == strings.ToLower(categoryQuery) {
				filter = append(filter, product)
			}
		}
		c.JSON(http.StatusOK, filter)
		return
	}
	c.JSON(http.StatusOK, products)
}

func main() {
	r := gin.Default()

	// สำหรับตรวจสอบว่าเซิร์ฟเวอร์ยังทำงานอยู่หรือไม่
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "healthy"})
	})

	api := r.Group("/api/v1")
	{
		// ดึงข้อมูลสินค้าเครื่องสำอางทั้งหมด หรือ filter ตาม category
		api.GET("/products", getProducts)
	}
	r.Run(":8080")
}