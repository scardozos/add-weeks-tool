package httpserver

import (
	"net/http"
	"path"

	"github.com/julienschmidt/httprouter"
)

type WeekHandlerRouterContext struct {
	Rtr *httprouter.Router
}
type WeekHandlerRouter struct {
	*WeekHandlerRouterContext
}

func InitRoutes() WeekHandlerRouter {
	router := httprouter.New()
	var placeholder httprouter.Handle
	router.GET("/", placeholder)
	router.GET("/add", placeholder)
	router.GET("/remove", placeholder)
	router.GET("/list", placeholder)
	router.ServeFiles("/src/*filepath", http.Dir(path.Join("public", "src")))
	return WeekHandlerRouter{
		&WeekHandlerRouterContext{Rtr: router},
	}
}
