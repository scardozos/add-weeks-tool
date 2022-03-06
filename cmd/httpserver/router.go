package httpserver

import (
	"net/http"
	"path"

	"github.com/julienschmidt/httprouter"
	"github.com/scardozos/add-weeks-tool/cmd/dbclient"
)

type WeekHandlerRouterContext struct {
	Rtr      *httprouter.Router
	DbClient *dbclient.LocalClient
}
type WeekHandlerRouter struct {
	*WeekHandlerRouterContext
}

func (w *WeekHandlerRouterContext) InitRoutes() {
	router := httprouter.New()
	var placeholder httprouter.Handle

	// UI
	router.GET("/", HtmlTemplate)

	// Actual HTTP API
	router.POST("/date", w.InsertDates)
	router.DELETE("/date", placeholder)
	router.GET("/date", w.GetDates)
	router.ServeFiles("/src/*filepath", http.Dir(path.Join("public", "src")))
	w.Rtr = router
}
