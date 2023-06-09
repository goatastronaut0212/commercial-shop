package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"commercial-shop.com/services"
)

func GetProduct(c *gin.Context) {
	data := services.ProductService{Items: []services.Product{{Id: c.Param("id")}}}
	err := data.Get()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get product value"})
		return
	}
	c.JSON(http.StatusOK, data)
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

	data := services.ProductService{}
	err = data.GetAll(&limit, &page)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get all product value"})
		return
	}
	c.JSON(http.StatusOK, data.Items)
}

func CreateProduct(c *gin.Context) {
	// Get input
	input := services.Product{}
	c.ShouldBindJSON(&input)
	fmt.Println(input)

	// Assign to data and create
	data := services.ProductService{Items: []services.Product{input}}
	fmt.Println(data)
	err := data.Create()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't create product!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "create product successfully!"})
}

func UpdateProduct(c *gin.Context) {
	// Get input
	input := services.Product{}
	c.ShouldBindJSON(&input)
	input.Id = c.Param("id")

	// Assign to data and update
	data := services.ProductService{Items: []services.Product{input}}
	err := data.Update()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't update product!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "update product successfully!"})
}

func DeleteProduct(c *gin.Context) {
	data := services.ProductService{Items: []services.Product{{
		Id: c.Param("id"),
	}}}
	err := data.Delete()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't delete product!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "delete product successfully!"})
}
