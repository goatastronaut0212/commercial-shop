package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"commercial-shop.com/access"
	"commercial-shop.com/models"
)

func GetAllBill_Detail(c *gin.Context) {
	data, err := access.FindAllBillDetail()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get all bill detail value"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetBill_Detail(c *gin.Context) {
	id := c.Param("id")
	data, err := access.FindBillDetail(&id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get bill detail value"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func CreateBill_Detail(c *gin.Context) {
	data := models.Bill_Detail{}
	c.ShouldBindJSON(&data)
	err := access.CreateBillDetail(&data)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't create bill detail!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "create bill detail successfully!"})
}

func UpdateBill_Detail(c *gin.Context) {
	data := models.Bill_Detail{}
	c.ShouldBindJSON(&data)
	data.ID = c.Param("id")

	err := access.UpdateBillDetail(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't update bill detail!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "update bill detail successfully!"})
}


func DeleteBill_Detail(c *gin.Context) {
	id := c.Param("id")
	err := access.DeleteBillDetail(&id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't delete bill detail!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "delete bill detail successfully!"})
}