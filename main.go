package main

import (
	"tugas_restApi/models"

	"github.com/gin-gonic/gin"

	"tugas_restApi/controller/usercontroller"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/users", usercontroller.Index)
	r.GET("/api/users/:id", usercontroller.Show)
	r.POST("/api/users", usercontroller.Create)
	r.PUT("/api/users/:id", usercontroller.Update)
	r.DELETE("/api/users", usercontroller.Delete)

	r.Run()
}