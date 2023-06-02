package app

import (
	"log"
	"rest/internal/feeder"
	"rest/internal/handler"
	"rest/pkg/firestore"

	"github.com/gin-gonic/gin"
)

const conf = "configs/service-acc.json"

func Run() {
	feed := feeder.NewRepo(conf)
	r := gin.Default()

	r.GET("/ping", handler.PingGet())

	r.GET("/feed", handler.FeederGet(feed))
	r.POST("/feed", handler.FeederPost(feed))
	r.PATCH("/feed", handler.FeederUpdate(feed))
	r.DELETE("/feed", handler.FeederDelete(feed))

	r.Run()
}

func Fb() {
	s := firestore.NewStore(conf, "user")
	err := s.Close()

	if err != nil {
		log.Fatalln(err)
	}
}
