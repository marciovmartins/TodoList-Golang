package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/todos", func(c *gin.Context) {
		items := []gin.H{
			{"id": 1, "title": "Create README"},
			{"id": 2, "title": "Create First Gin Gonic resource"},
			{"id": 3, "title": "Create"},
		}
		c.JSON(http.StatusOK, gin.H{
			"items": items,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
