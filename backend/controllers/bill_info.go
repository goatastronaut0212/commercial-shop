package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"commercial-shop.com/access"
	"commercial-shop.com/models"
)

func GetAllBill_Info(c *gin.Context) {
	data, err := access.FindAllBillInfo()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get all bill info value"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetBill_Info(c *gin.Context) {
	id := c.Param("id")
	data, err := access.FindBillInfo(&id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get bill info value"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func CreateBillInfo(c *gin.Context) {
	data := models.Bill_Info{}
	c.ShouldBindJSON(&data)
	err := access.CreateBillInfo(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't create bill info!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "create bill info successfully!"})
}

func UpdateBillInfo(c *gin.Context) {
	data := models.Bill_Info{}
	c.ShouldBindJSON(&data)
	data.Id = c.Param("id")

	err := access.UpdateBillInfo(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't update bill info!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "update bill info successfully!"})
}

func DeleteBillInfo(c *gin.Context) {
	id := c.Param("id")
	err := access.DeleteBillInfo(&id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't delete bill info!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "delete bill info successfully!"})
}
