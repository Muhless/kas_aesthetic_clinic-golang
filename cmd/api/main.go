package main

import (
	"kas_aesthetic_clinic/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	router := gin.Default()

	router.Run("localhost:8080")
}
