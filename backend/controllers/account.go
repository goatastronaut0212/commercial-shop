package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"commercial-shop.com/access"
	"commercial-shop.com/models"
)

func GetAllAccount(c *gin.Context) {
	data, err := access.FindAllAccount()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get all account value"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetAccount(c *gin.Context) {
	id := c.Param("id")
	data, err := access.FindAccount(&id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get account value"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func CreateAccount(c *gin.Context) {
	data := models.Account{}
	c.ShouldBindJSON(&data)
	fmt.Printf(data.UserName);
	err := access.CreateAccount(&data)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't create account!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "create account successfully!"})
}

func UpdateAccount(c *gin.Context) {
	data := models.Account{}
	c.ShouldBindJSON(&data)
	data.UserName = c.Param("id")

	err := access.UpdateAccount(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't update account!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "update account successfully!"})
}


func DeleteAccount(c *gin.Context) {
	id := c.Param("id")
	err := access.DeleteAccount(&id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't delete account!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "delete account successfully!"})
}