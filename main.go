package main

import (
	"web/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/tambah", controllers.Add)
	router.POST("/kurang", controllers.Substract)
	router.POST("/kali", controllers.Times)
	router.POST("/bagi", controllers.Divide)
	router.POST("/faktorial", controllers.Factorial)

	router.Run()
}
