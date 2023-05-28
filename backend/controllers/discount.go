package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"commercial-shop.com/access"
	"commercial-shop.com/models"
)

func GetAllDiscount(c *gin.Context) {
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

	data, err := access.FindAllDiscount(&limit, &page)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get all discount value"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetDiscount(c *gin.Context) {
	id := c.Param("id")
	data, err := access.FindDiscount(&id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get discount value"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func CreateDiscount(c *gin.Context) {
	data := models.Discount{}
	c.ShouldBindJSON(&data)
	err := access.CreateDiscount(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't create discount!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "create discount successfully!"})
}

func UpdateDiscount(c *gin.Context) {
	data := models.Discount{}
	c.ShouldBindJSON(&data)
	data.Id = c.Param("id")

	err := access.UpdateDiscount(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't update discount!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "update discount successfully!"})
}

func DeleteDiscount(c *gin.Context) {
	id := c.Param("id")
	err := access.DeleteDiscount(&id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't delete discount!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "delete discount successfully!"})
}
