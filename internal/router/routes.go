package router

import (
	"rest/internal/handler"

	"github.com/gin-gonic/gin"
)

type Routes []gin.RouteInfo

var routes = Routes{
	gin.RouteInfo{
		Method:      "GET",
		Path:        "/",
		Handler:     "Ping",
		HandlerFunc: handler.PingGet(),
	},
	gin.RouteInfo{
		Method:      "GET",
		Path:        "/feed",
		Handler:     "GetFeed",
		HandlerFunc: handler.FeederGet(nil),
	},
	gin.RouteInfo{
		Method:      "PATCH",
		Path:        "/feed",
		Handler:     "UpdateFeed",
		HandlerFunc: handler.FeederUpdate(nil),
	},
	gin.RouteInfo{
		Method:      "POST",
		Path:        "/feed",
		Handler:     "FeederPost",
		HandlerFunc: handler.FeederPost(nil),
	},
	gin.RouteInfo{
		Method:      "DELETE",
		Path:        "/feed",
		Handler:     "FeederDelete",
		HandlerFunc: handler.FeederDelete(nil),
	},
}
