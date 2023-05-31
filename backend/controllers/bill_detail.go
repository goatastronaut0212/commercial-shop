package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"commercial-shop.com/access"
	"commercial-shop.com/models"
)

func GetAllBillDetail(c *gin.Context) {
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

	data, err := access.FindAllBillDetail(&limit, &page)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get all bill detail value"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetBillDetail(c *gin.Context) {
	id := c.Param("id")
	data, err := access.FindBillDetail(&id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get bill detail value"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func CreateBillDetail(c *gin.Context) {
	data := models.BillDetail{}
	c.ShouldBindJSON(&data)
	err := access.CreateBillDetail(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't create bill detail!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "create bill detail successfully!"})
}

func UpdateBillDetail(c *gin.Context) {
	data := models.BillDetail{}
	c.ShouldBindJSON(&data)
	data.Id = c.Param("id")

	err := access.UpdateBillDetail(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't update bill detail!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "update bill detail successfully!"})
}

func DeleteBillDetail(c *gin.Context) {
	id := c.Param("id")
	err := access.DeleteBillDetail(&id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't delete bill detail!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "delete bill detail successfully!"})
}
