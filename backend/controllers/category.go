package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"commercial-shop.com/access"
	"commercial-shop.com/models"
)

func GetCategoryAll(c *gin.Context) {
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

	data, err := access.FindCategoryAll(&limit, &page)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get all category value"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetCategoryById(c *gin.Context) {
	id := c.Param("id")
	data, err := access.FindCategoryById(&id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get category value"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func CreateCategory(c *gin.Context) {
	category := models.Category{}
	c.ShouldBindJSON(&category)

	err := access.CreateCategory(&category)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't create category!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "create category successfully!"})
}

func UpdateCategory(c *gin.Context) {
	category := models.Category{}
	c.ShouldBindJSON(&category)
	category.Id = c.Param("id")

	err := access.UpdateCategory(&category)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't update category!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "update category successfully!"})
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	err := access.DeleteCategory(&id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't delete category!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "delete category successfully!"})
}
