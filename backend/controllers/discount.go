package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"commercial-shop.com/models"
	"commercial-shop.com/services"
)

func CreateDiscount(c *gin.Context) {
	// Create service and assign to data
	data := services.DiscountService{Items: []models.Discount{{}}}
	c.ShouldBindJSON(&data.Items[0])

	// Execute method and send status request to user
	err := data.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't create discount!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "create discount successfully!"})
}

func DeleteDiscount(c *gin.Context) {
	// Create service and assign to data
	data := services.DiscountService{Items: []models.Discount{{
		Id: c.Param("id"),
	}}}

	// Execute method and send status request to user
	err := data.Delete()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't delete discount!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "delete discount successfully!"})
}

func GetDiscount(c *gin.Context) {
	// Create service and assign to data
	data := services.DiscountService{Items: []models.Discount{{
		Id: c.Param("id"),
	}}}

	// Execute method and send status request to user
	err := data.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get discount value"})
		return
	}
	c.JSON(http.StatusOK, data.Items)
}

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

	// Create service and assign to data
	// Then execute method and send status request to user
	data := services.DiscountService{}
	err = data.GetAll(&limit, &page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get all discount value"})
		return
	}
	c.JSON(http.StatusOK, data.Items)
}

func UpdateDiscount(c *gin.Context) {
	// Create service and assign to data
	data := services.DiscountService{Items: []models.Discount{{}}}
	c.ShouldBindJSON(&data.Items[0])
	data.Items[0].Id = c.Param("id")

	// Execute method and send status request to user
	err := data.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't update discount!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "update discount successfully!"})
}
