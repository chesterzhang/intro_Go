package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"time"
)




func main() {
	r := gin.Default() //返回一个 gin 的 engine
	logger,err:=zap.NewProduction() //zap logger

	if err!=nil{
		panic(err)
	}

	//middleware, 对所有请求 进行 log
	r.Use(func(c *gin.Context) {
		s:=time.Now()
		c.Next()//log 以后继续执行
		//path, response code,log latency
		logger.Info("incoming request", zap.String("path",c.Request.URL.Path),zap.Int("status", c.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)))

	})

	r.Use(func(c *gin.Context) {
		c.Set("requestId", rand.Int())
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		h:=gin.H{
			"message": "pong"}

		if rid, exists :=c.Get("requestId"); exists{
			h["requestId"]=rid
		}
		c.JSON(200, h)
		})


	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")

	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}