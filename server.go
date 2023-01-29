package main

import (
	"awesomeProject/controller"
	"awesomeProject/repository"
	"awesomeProject/util"
	"github.com/gin-gonic/gin"
	"os"
)

// 通过gin搭建web框架
func main() {
	// 初始化数据索引
	if err := Init(); err != nil {
		os.Exit(-1)
	}
	// 初始化引擎配置
	r := gin.Default()

	r.Use(gin.Logger())

	// 构建路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := controller.QueryPageInfo(topicId)
		c.JSON(200, data)
	})

	r.POST("/community/post/do", func(c *gin.Context) {
		uid, _ := c.GetPostForm("uid")
		topicId, _ := c.GetPostForm("topic_id")
		content, _ := c.GetPostForm("content")
		data := controller.PublishPost(uid, topicId, content)
		c.JSON(200, data)
	})

	// 启动服务
	err := r.Run()
	if err != nil {
		return
	}
}

func Init() error {
	if err := repository.Init(); err != nil {
		return err
	}
	if err := util.InitLogger(); err != nil {
		return err
	}
	return nil
}
