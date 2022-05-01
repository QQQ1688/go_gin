package main

import (
	"go-api/controller"

	"github.com/gin-gonic/gin"
)

// sets up an association in which getDatas handles requests to the /iplogs endpoint path.
func main() {
	// Initialize a Gin router using Default.
	router := gin.Default()
	router.GET("/mysql", controller.GetDatas)
	router.GET("/mysql/:ip", controller.GetDataByIP)

	// Use the Run function to attach the router to an http.Server and start the server.
	router.Run("localhost:8080")
}
