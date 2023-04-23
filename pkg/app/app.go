package app

import (
	"rest/internal/feeder"
	"rest/pkg/firebase/firestore"
	"rest/pkg/handler"

	"github.com/gin-gonic/gin"
)

const conf = "conf/service-acc.json"

func Run() {
	feed := feeder.NewRepo()
	r := gin.Default()

	r.GET("/ping", handler.PingGet())

	r.GET("/feed", handler.FeederGet(feed))
	r.POST("/feed", handler.FeederPost(feed))

	r.Run()
}

func Fb() {
	c := firestore.Conn(conf)
	firestore.Close(c)
}
