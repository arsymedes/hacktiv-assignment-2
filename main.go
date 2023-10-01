package main

import (
	"github.com/arsymedes/hacktiv-assignment-2/controllers"
	"github.com/arsymedes/hacktiv-assignment-2/database"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.StartDB()

	router.GET("/order", controllers.GetOrders)
	router.POST("/order", controllers.CreateOrder)
	router.PATCH("/order/:id", controllers.PatchOrder)
	router.PUT("/order/:id", controllers.PutOrder)
	router.DELETE("/order/:id", controllers.DeleteOrder)

	router.Run()
}
