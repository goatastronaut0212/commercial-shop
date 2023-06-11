package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"commercial-shop.com/models"
	"commercial-shop.com/services"
)

func GetAccount(c *gin.Context) {
	// Create service and assign to data
	data := services.AccountService{Items: []models.Account{{
		Username: c.Param("username"),
	}}}

	// Execute method and send status request to user
	err := data.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get account value"})
		return
	}
	c.JSON(http.StatusOK, data.Items)
}

func GetAllAccount(c *gin.Context) {
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
	data := services.AccountService{}
	err = data.GetAll(&limit, &page)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get all account value"})
		return
	}
	c.JSON(http.StatusOK, data.Items)
}

func CreateAccount(c *gin.Context) {
	// Create service and assign to data
	data := services.AccountService{Items: []models.Account{{}}}
	c.ShouldBindJSON(&data.Items[0])

	// Execute method and send status request to user
	err := data.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't create account!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "create account successfully!"})
}

func UpdateAccount(c *gin.Context) {
	// Create service and assign to data
	data := services.AccountService{Items: []models.Account{{}}}
	c.ShouldBindJSON(&data.Items[0])
	data.Items[0].Username = c.Param("username")

	// Execute method and send status request to user
	err := data.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't update account!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "update account successfully!"})
}

func DeleteAccount(c *gin.Context) {
	// Create service and assign to data
	data := services.AccountService{Items: []models.Account{{
		Username: c.Param("username"),
	}}}

	// Execute method and send status request to user
	err := data.Delete()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't delete account!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "delete account successfully!"})
}
