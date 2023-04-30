package router

import (
	"rest/internal/handler"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		Name:        "Ping",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: handler.PingGet(),
	},
	Route{
		Name:        "FeederGet",
		Method:      "GET",
		Pattern:     "/feed",
		HandlerFunc: handler.FeederGet(nil),
	},
	Route{
		Name:        "FeederPost",
		Method:      "POST",
		Pattern:     "/feed",
		HandlerFunc: handler.FeederPost(nil),
	},
	Route{
		Name:        "FeederDelete",
		Method:      "DELETE",
		Pattern:     "/feed",
		HandlerFunc: handler.FeederDelete(nil),
	},
	Route{
		Name:        "FeederUpdate",
		Method:      "PATCH",
		Pattern:     "/feed",
		HandlerFunc: handler.FeederUpdate(nil),
	},
}
