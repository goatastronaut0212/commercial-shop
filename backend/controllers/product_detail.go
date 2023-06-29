package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"commercial-shop.com/models"
	"commercial-shop.com/services"
)

func CreateProductDetail(c *gin.Context) {
	// Create service and assign to data
	data := services.ProductDetailService{Items: []models.ProductDetail{{}}}
	c.ShouldBindJSON(&data.Items[0])

	// Execute method and send status request to user
	err := data.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't create product detail!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "create product detail successfully!"})
}

func DeleteProductDetail(c *gin.Context) {
	// Create service and assign to data
	data := services.ProductDetailService{Items: []models.ProductDetail{{
		Id: c.Param("id"),
	}}}

	// Execute method and send status request to user
	err := data.Delete()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't delete product detail!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "delete product detail successfully!"})
}

func GetProductDetail(c *gin.Context) {
	// Create service and assign to data
	data := services.ProductDetailService{Items: []models.ProductDetail{{
		Id: c.Param("id"),
	}}}

	// Execute method and send status request to user
	err := data.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get product detail value"})
		return
	}
	c.JSON(http.StatusOK, data.Items)
}

func GetAllProductDetail(c *gin.Context) {
	// Get limit
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 20
	}
	if limit <= 0 {
		limit = 20
	}

	// Get page
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	if page <= 0 {
		page = 1
	}

	// Create service and assign to data
	// Then execute method and send status request to user
	data := services.ProductDetailService{}
	err = data.GetAll(&limit, &page)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get all product detail value"})
		return
	}
	c.JSON(http.StatusOK, data.Items)
}

func UpdateProductDetail(c *gin.Context) {
	// Create service and assign to data
	data := services.ProductDetailService{Items: []models.ProductDetail{{}}}
	c.ShouldBindJSON(&data.Items[0])
	data.Items[0].Id = c.Param("id")

	// Execute method and send status request to user
	err := data.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't update product detail!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "update product detail successfully!"})
}
