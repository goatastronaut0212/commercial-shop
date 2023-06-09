package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"commercial-shop.com/access"
	"commercial-shop.com/services"
)

func GetCustomer(c *gin.Context) {
	data := services.CustomerService{Items: []services.Customer{{
		Id: c.Param("id"),
	}}}
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

	data := services.CustomerService{}
	err = data.GetAll(&limit, &page)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get all customer value"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func CreateCustomer(c *gin.Context) {
	// Get input
	input := services.Customer{}
	c.ShouldBindJSON(&input)
	fmt.Println(input)

	// Assign to data and create
	data := services.CustomerService{Items: []services.Customer{input}}
	err := data.Create()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't create customer!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "create customer successfully!"})
}

func UpdateCustomer(c *gin.Context) {
	// Get input
	input := services.Customer{}
	c.ShouldBindJSON(&input)
	fmt.Println(input)

	// Assign to data and create
	data := services.CustomerService{Items: []services.Customer{input}}
	fmt.Println(data)
	err := data.Update()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't update customer!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "update customer successfully!"})
}

func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	err := access.DeleteCustomer(&id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't delete customer!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "delete customer successfully!"})
}
