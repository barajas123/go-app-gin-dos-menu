package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/menu", "./static/")

	r.Run(":80")
}
