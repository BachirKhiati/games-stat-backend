package main

import "github.com/gin-gonic/gin"

func handleGetGunsList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello earth",
	})
}
func NameHandler(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"message": "Hello " + name,
	})
}

func main() {
	router := gin.Default()
	router.GET("/guns", handleGetGunsList)
	router.GET("/:name", NameHandler)
	router.Run(":5000")
}
