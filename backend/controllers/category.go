package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"commercial-shop.com/models"
	"commercial-shop.com/services"
)

func CreateCategory(c *gin.Context) {
	// Create service and assign to data
	data := services.CategoryService{Items: []models.Category{{}}}
	c.ShouldBindJSON(&data.Items[0])

	// Execute method and send status request to user
	err := data.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't create category!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "create category successfully!"})
}

func DeleteCategory(c *gin.Context) {
	// Create service and assign to data
	data := services.CategoryService{Items: []models.Category{{
		Id: c.Param("id"),
	}}}

	// Execute method and send status request to user
	err := data.Delete()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't delete category!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "delete category successfully!"})
}

func GetCategory(c *gin.Context) {
	// Create service and assign to data
	data := services.CategoryService{Items: []models.Category{{
		Id: c.Param("id"),
	}}}

	// Execute method and send status request to user
	err := data.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get category value"})
		return
	}
	c.JSON(http.StatusOK, data.Items)
}

func GetAllCategory(c *gin.Context) {
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
	data := services.CategoryService{}
	err = data.GetAll(&limit, &page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get all category value"})
		return
	}
	c.JSON(http.StatusOK, data.Items)
}

func UpdateCategory(c *gin.Context) {
	// Create service and assign to data
	data := services.CategoryService{Items: []models.Category{{}}}
	c.ShouldBindJSON(&data.Items[0])
	data.Items[0].Id = c.Param("id")

	// Execute method and send status request to user
	err := data.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't update category!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "update category successfully!"})
}
