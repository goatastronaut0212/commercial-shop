package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"commercial-shop.com/models"
	"commercial-shop.com/services"
)

func CreateCustomer(c *gin.Context) {
	// Create service and assign to data
	data := services.CustomerService{Items: []models.Customer{{}}}
	c.ShouldBindJSON(&data.Items[0])

	// Execute method and send status request to user
	err := data.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't create customer!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "create customer successfully!"})
}

func DeleteCustomer(c *gin.Context) {
	// Create service and assign to data
	data := services.CustomerService{Items: []models.Customer{{
		Id: c.Param("id"),
	}}}

	// Execute method and send status request to user
	err := data.Delete()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't delete customer!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "delete customer successfully!"})
}

func GetCustomer(c *gin.Context) {
	// Create service and assign to data
	data := services.CustomerService{Items: []models.Customer{{
		Id: c.Param("id"),
	}}}

	// Execute method and send status request to user
	err := data.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get customer value"})
		return
	}
	c.JSON(http.StatusOK, data.Items)
}

func GetAllCustomer(c *gin.Context) {
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
	data := services.CustomerService{}
	err = data.GetAll(&limit, &page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get all customer value"})
		return
	}
	c.JSON(http.StatusOK, data.Items)
}

func UpdateCustomer(c *gin.Context) {
	// Create service and assign to data
	data := services.CustomerService{Items: []models.Customer{{}}}
	c.ShouldBindJSON(&data.Items[0])
	data.Items[0].Id = c.Param("id")

	// Execute method and send status request to user
	err := data.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't update customer!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "update customer successfully!"})
}
