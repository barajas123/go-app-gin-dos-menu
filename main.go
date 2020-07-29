package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	r.GET("/menu",func(c *gin.Context){
		c.File("./static/lunch.jpg")
		c.File("./static/pg1.jpg")
		c.File("./static/pg2.jpg")
		c.File("./static/pg3.jpg")
		c.File("./static/pg4.jpg")
		c.File("./static/pg5.jpg")
		c.File("./static/pg6.jpg")
		c.File("./static/pg7.jpg")
	})

	r.Run(":8080")
}
