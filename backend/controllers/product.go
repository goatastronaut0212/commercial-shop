package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"commercial-shop.com/access"
	"commercial-shop.com/models"
)

func GetProductAll(c *gin.Context) {
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

	data, err := access.FindProductAll(&limit, &page)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get all product value"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetProductById(c *gin.Context) {
	id := c.Param("id")
	data, err := access.FindProductById(&id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get product value"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func CreateProduct(c *gin.Context) {
	data := models.Product{}
	c.ShouldBindJSON(&data)
	err := access.CreateProduct(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't create product!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "create product successfully!"})
}

func UpdateProduct(c *gin.Context) {
	data := models.Product{}
	c.ShouldBindJSON(&data)
	data.Id = c.Param("id")

	err := access.UpdateProduct(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't update product!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "update product successfully!"})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	err := access.DeleteProduct(&id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't delete product!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "delete product successfully!"})
}
