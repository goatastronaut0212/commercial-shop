package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"commercial-shop.com/models"
	"commercial-shop.com/services"
)

func CreateProduct(c *gin.Context) {
	// Create service and assign to data
	data := services.ProductService{Items: []models.Product{{}}}
	c.ShouldBindJSON(&data.Items[0])

	// Execute method and send status request to user
	err := data.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't create product!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "create product successfully!"})
}

func DeleteProduct(c *gin.Context) {
	// Create service and assign to data
	data := services.ProductService{Items: []models.Product{{
		Id: c.Param("id"),
	}}}

	// Execute method and send status request to user
	err := data.Delete()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't delete product!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "delete product successfully!"})
}

func GetProduct(c *gin.Context) {
	// Create service and assign to data
	data := services.ProductService{Items: []models.Product{{
		Id: c.Param("id"),
	}}}

	// Execute method and send status request to user
	err := data.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get product value"})
		return
	}
	c.JSON(http.StatusOK, data.Items)
}

func GetAllProduct(c *gin.Context) {
	// Get limit
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}
	if limit <= 0 {
		limit = 10
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
	data := services.ProductService{}
	err = data.GetAll(&limit, &page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get all product value"})
		return
	}
	c.JSON(http.StatusOK, data.Items)
}

func UpdateProduct(c *gin.Context) {
	// Create service and assign to data
	data := services.ProductService{Items: []models.Product{{}}}
	c.ShouldBindJSON(&data.Items[0])
	data.Items[0].Id = c.Param("id")

	// Check options
	category_option, name_option := true, true

	if data.Items[0].IdCategory == "" {
		category_option = false
	}
	if data.Items[0].Name == "" {
		name_option = false
	}

	// Execute method and send status request to user
	err := data.Update(&category_option, &name_option)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't update product!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "update product successfully!"})
}
