package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"commercial-shop.com/access"
	"commercial-shop.com/models"
)

func GetAllCustomer(c *gin.Context) {
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

	data, err := access.FindAllCustomer(&limit, &page)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get all customer value"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetCustomer(c *gin.Context) {
	id := c.Param("id")
	data, err := access.FindCustomer(&id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get category value"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func CreateCustomer(c *gin.Context) {
	data := models.Customer{}
	c.ShouldBindJSON(&data)

	err := access.CreateCustomer(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't create category!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "create customer successfully!"})
}

func UpdateCustomer(c *gin.Context) {
	data := models.Customer{}
	c.ShouldBindJSON(&data)
	data.Id = c.Param("id")

	err := access.UpdateCustomer(&data)

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
