package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/arsymedes/hacktiv-assignment-2/database"
	"github.com/arsymedes/hacktiv-assignment-2/models"
)

func GetOrders(ctx *gin.Context) {
	var orders []models.Order

	db := database.GetDB()
	db.Model(&models.Order{}).Preload("Items").Find(&orders)

	ctx.JSON(http.StatusOK, gin.H{"data": orders})
}

func CreateOrder(ctx *gin.Context) {
	var input models.NewOrderModel

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	order := models.Order{
		CustomerName: input.CustomerName,
		Items:        input.Items,
		OrderedAt:    input.OrderedAt,
	}
	db := database.GetDB()
	db.Create(&order)

	ctx.JSON(http.StatusOK, gin.H{"data": order})
}

func PatchOrder(ctx *gin.Context) {
	db := database.GetDB()

	var order models.Order
	if err := db.Where("ID = ?", ctx.Param("id")).First(&order).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	var input models.Order
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&order).Updates(input)

	ctx.JSON(http.StatusOK, gin.H{"data": order})
}

func PutOrder(ctx *gin.Context) {
	db := database.GetDB()

	var order models.Order
	if err := db.Where("id = ?", ctx.Param("id")).First(&order).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	var input models.NewOrderModel
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	db.Model(&order).Updates(input)

	ctx.JSON(http.StatusOK, gin.H{"data": order})
}

func DeleteOrder(ctx *gin.Context) {
	db := database.GetDB()

	var order models.Order
	if err := db.Where("id = ?", ctx.Param("id")).First(&order).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	db.Delete(&order)

	ctx.JSON(http.StatusOK, gin.H{"data": order})
}
