package main

import (
	"context"

	"example.com/goRedis/api"
	"example.com/goRedis/db"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	db.InitRedis(ctx)
	defer db.Close()
	r := gin.Default()
	r.POST("/userAdd", func(c *gin.Context) {
		api.UserAdd(c, &ctx)
	})
	r.GET("/userAll", func(c *gin.Context) {
		api.UserAll(c, &ctx)
	})
	r.GET("/userGet/:key", func(c *gin.Context) {
		key := c.Param("key")

		api.UserGet(c, &ctx, key)
	})
	r.POST("/userUpdate", func(c *gin.Context) {
		api.UserUpdate(c, &ctx)
	})
	r.POST("/userDel", func(c *gin.Context) {
		api.UserDel(c, &ctx)
	})
	r.Run(":8000")
}
