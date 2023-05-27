package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"commercial-shop.com/access"
	"commercial-shop.com/models"
)

func GetProductImageAll(c *gin.Context) {
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

	data, err := access.FindProductImageAll(&limit, &page)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get all product image value"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetProductImageById(c *gin.Context) {
	id := c.Param("id")
	data, err := access.FindProductImageById(&id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get product image value"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func CreateProductImage(c *gin.Context) {
	data := models.ProductImage{}
	c.ShouldBindJSON(&data)
	err := access.CreateProductImage(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't create product image!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "create product image successfully!"})
}

func UpdateProductImage(c *gin.Context) {
	data := models.ProductImage{}
	c.ShouldBindJSON(&data)
	data.Id = c.Param("id")

	err := access.UpdateProductImage(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't update product image!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "update product image successfully!"})
}

func DeleteProductImage(c *gin.Context) {
	id := c.Param("id")
	err := access.DeleteProductImage(&id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't delete product image!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "delete product image successfully!"})
}
