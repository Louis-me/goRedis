package api

import (
	"context"
	"fmt"
	"net/http"

	"example.com/goRedis/models"
	"github.com/gin-gonic/gin"
)

func UserAdd(c *gin.Context, ctx *context.Context) {
	fmt.Println("add=", c)
	var user models.User
	if c.Bind(&user) == nil { //把客户端格式传过来的数据绑定到结构体user中去
		fmt.Println("data=", user)
		err := user.Add(ctx) // 调用model层的对应方法
		if err != nil {

			c.JSON(http.StatusOK, gin.H{
				"msg":  "新增失败",
				"code": -1,
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "新增成功",
				"code": 1,
			})
		}
	} else {
		c.JSON(400, gin.H{"JSON=== status": "binding JSON error!"})
	}
}

func UserGet(c *gin.Context, ctx *context.Context, key string) {
	// 接受key
	users := models.User{
		Key: key,
	}
	if c.Bind(&users) == nil {
		users, err := users.GetUser(ctx)
		if err != nil {

			c.JSON(http.StatusOK, gin.H{
				"msg":  "获取失败",
				"code": -1,
				"user": users,
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "获取成功",
				"code": 1,
				"user": users,
			})
		}
	} else {
		c.JSON(400, gin.H{"JSON=== status": "binding JSON error!"})
	}
}

func UserAll(c *gin.Context, ctx *context.Context) {
	var user models.User
	if c.Bind(&user) == nil {
		users, err := user.QueryAll(ctx)
		if err != nil {

			c.JSON(http.StatusOK, gin.H{
				"msg":  "获取失败",
				"code": -1,
				"user": users,
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "获取成功",
				"code": 1,
				"user": users,
			})
		}
	} else {
		c.JSON(400, gin.H{"JSON=== status": "binding JSON error!"})
	}
}

func UserDel(c *gin.Context, ctx *context.Context) {
	var user models.User
	if c.Bind(&user) == nil {
		err := user.DeleteUser(ctx)
		fmt.Println("del=", err)
		if !err {

			c.JSON(http.StatusOK, gin.H{
				"msg":  "删除失败1",
				"code": -1,
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "删除成功1",
				"code": 1,
			})
		}
	} else {
		c.JSON(400, gin.H{"JSON=== status": "binding JSON error!"})
	}
}

func UserUpdate(c *gin.Context, ctx *context.Context) {
	var user models.User
	if c.Bind(&user) == nil {
		err := user.UpdateUser(ctx)
		if err != nil {

			c.JSON(http.StatusOK, gin.H{
				"msg":  "修改失败",
				"code": -1,
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "修改成功",
				"code": 1,
			})
		}
	} else {
		c.JSON(400, gin.H{"JSON=== status": "binding JSON error!"})
	}
}
