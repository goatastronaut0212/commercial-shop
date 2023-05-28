package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"commercial-shop.com/access"
	"commercial-shop.com/models"
)

func GetAllCustomer(c *gin.Context) {
	data, err := access.FindCustomerAll()

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

// func UpdateCategory(c *gin.Context) {
// 	data := models.Category{}
// 	c.ShouldBindJSON(&data)
// 	data.Id = c.Param("id")

// 	err := access.UpdateCategory(&data)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "can't update category!"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": "update category successfully!"})
// }

// func DeleteCategory(c *gin.Context) {
// 	id := c.Param("id")
// 	err := access.DeleteCategory(&id)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "can't delete category!"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": "delete category successfully!"})
// }
