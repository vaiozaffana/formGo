package main

import (
	controller "formGo/Controller"
	database "formGo/Database"
	model "formGo/Model"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&model.User{})

	r := gin.Default()

	r.GET("/", controller.RenderForm)

	r.POST("/submit", controller.SubmitForm)

	r.Run(":8080")
}
