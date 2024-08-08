package main

import (
	"Task-5/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)

	r.Run("localhost:8080")

}
