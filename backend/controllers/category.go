package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"commercial-shop.com/services"
)

func CreateCategory(c *gin.Context) {
	// Get input
	input := services.Category{}
	c.ShouldBindJSON(&input)

	// Assign to data and create
	data := services.CategoryService{Items: []services.Category{input}}
	err := data.Create()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't create category!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "create category successfully!"})
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

	data := services.CategoryService{}
	err = data.GetAll(&limit, &page)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get all category value"})
		return
	}
	c.JSON(http.StatusOK, data.Items)
}

func GetCategory(c *gin.Context) {
	data := services.CategoryService{Items: []services.Category{{
		Id: c.Param("id"),
	}}}
	err := data.Get()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get category value"})
		return
	}
	c.JSON(http.StatusOK, data.Items[0])
}

func UpdateCategory(c *gin.Context) {
	// Get input
	input := services.Category{}
	c.ShouldBindJSON(&input)
	input.Id = c.Param("id")

	// Assign to data and update
	data := services.CategoryService{Items: []services.Category{input}}
	err := data.Update()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't update category!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "update category successfully!"})
}

func DeleteCategory(c *gin.Context) {
	data := services.CategoryService{Items: []services.Category{{
		Id: c.Param("id"),
	}}}
	err := data.Delete()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't delete category!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "delete category successfully!"})
}
