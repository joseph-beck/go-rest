package handler

import (
	"net/http"
	"rest/internal/feeder"

	"github.com/gin-gonic/gin"
)

func FeederGet(feed feeder.RepoGetter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
