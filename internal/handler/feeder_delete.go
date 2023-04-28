package handler

import (
	"net/http"
	"rest/internal/feeder"

	"github.com/gin-gonic/gin"
)

func FeederDelete(feed feeder.RepoDeleter) gin.HandlerFunc {
	return func(c *gin.Context) {
		var id string
		c.Bind(&id)

		feed.Delete(id)

		c.Status(http.StatusNoContent)
	}
}
