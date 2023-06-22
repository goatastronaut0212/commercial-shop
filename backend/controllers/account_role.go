package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"commercial-shop.com/models"
	"commercial-shop.com/services"
)

func GetAccountRole(c *gin.Context) {
	// Create service and assign to data
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "bad request input"})
	}
	data := services.AccountRoleService{Items: []models.AccountRole{{
		Id: id,
	}}}

	// Execute method and send status request to user
	err = data.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get account role value"})
		return
	}
	c.JSON(http.StatusOK, data.Items)
}

func GetAllAccountRole(c *gin.Context) {
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
	data := services.AccountRoleService{}
	err = data.GetAll(&limit, &page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get all account role value"})
		return
	}
	c.JSON(http.StatusOK, data.Items)
}

func CreateAccountRole(c *gin.Context) {
	// Create service and assign to data
	data := services.AccountRoleService{Items: []models.AccountRole{{}}}
	c.ShouldBindJSON(&data.Items[0])

	// Execute method and send status request to user
	err := data.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't create account role!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "create account role successfully!"})
}

func UpdateAccountRole(c *gin.Context) {
	// Create service and assign to data
	data := services.AccountRoleService{Items: []models.AccountRole{{}}}
	c.ShouldBindJSON(&data.Items[0])
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "bad request input"})
	}
	data.Items[0].Id = id

	// Execute method and send status request to user
	err = data.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't update account role!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "update account role successfully!"})
}

func DeleteAccountRole(c *gin.Context) {
	// Create service and assign to data
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "bad request input"})
	}
	data := services.AccountRoleService{Items: []models.AccountRole{{
		Id: id,
	}}}

	// Execute method and send status request to user
	err = data.Delete()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't delete account role!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "delete account role successfully!"})
}
