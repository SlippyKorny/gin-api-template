package main

import (
	"fmt"

	"github.com/SlippyKorny/gin-api-template/internal/cat"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World!")
	r := gin.Default()
	r.GET("/ping", cat.NewCatController().GetCats)
	r.Run()
}
