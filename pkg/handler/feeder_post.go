package handler

import (
	"net/http"
	"rest/internal/feeder"

	"github.com/gin-gonic/gin"
)

type feederPost struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func FeederPost(feed feeder.RepoAdder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := feederPost{}
		c.Bind(&requestBody)

		item := feeder.Item{
			Name: requestBody.Name,
			Data: requestBody.Data,
		}
		feed.Add(item)

		c.Status(http.StatusNoContent)
	}
}
