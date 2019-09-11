package main

import (
	. "./apis"
	db "./db"
	myjwt "./middleware/jwt"
	. "./utils"
	"github.com/gin-gonic/gin"
)

// // MiddleWare for router
// func MiddleWare() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Set("whatTheFuck", "guyFucks")
// 		c.Next()
// 	}
// }

func main() {
	defer db.Con.Close()

	router := gin.Default()
	router.POST("/login", HandleLogin)

	router.Use(myjwt.Auth())
	{
		router.GET("/users", HandleUserGetAll)
		router.GET("/users/:uid", HandleUserGetOne)
		router.POST("/users", HandleUserCreate)
		router.PUT("/users/:uid", HandleUserUpdate)
		router.DELETE("/users/:uid", HandleUserDelete)
	}

	var c Config
	c.GetConfig()

	router.Run(c.Dev.Port)
	// router.Run(":8080")
}
