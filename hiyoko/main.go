package main

import (
	"net/http"
  "os"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hiyoko")
	})

	return r
}

func main() {
	r := setupRouter()
	
	port := os.Getenv("PORT")
	if len(port) == 0 {
			port = "8080"
	}
	r.Run(":" + port)
}
