package app

import (
	"rest/internal/feeder"
	"rest/pkg/handler"

	"github.com/gin-gonic/gin"
)

func Run() {
	feed := feeder.NewRepo()
	r := gin.Default()

	r.GET("/ping", handler.PingGet())

	r.GET("/feed", handler.FeederGet(feed))
	r.POST("/feed", handler.FeederPost(feed))

	r.Run()
}
