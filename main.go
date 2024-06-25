package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello world")

	router := gin.New()
	router.Run(":8080")
}
