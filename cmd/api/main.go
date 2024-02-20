// Entry point for all the golang project.
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World ")
}

func main() {
	fmt.Println("Entry point for application !!!")
	router := gin.Default()
	router.GET("/", hello)

	router.Run("localhost:8080")
}
