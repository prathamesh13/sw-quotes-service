package main

import (
	qgen "github.com/prathamesh13/starwarsquotesgo"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/quote", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": getQuote(),
		})
	})

	r.Run(":3000") // listen and serve on 0.0.0.0:8080

}

func getQuote() string {
	return qgen.GetRandomQuote()
}
