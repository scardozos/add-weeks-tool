package httpserver

import (
	"net/http"

	"github.com/scardozos/add-weeks-tool/cmd/dbclient"
)

type HttpServer struct {
	router         WeekHandlerRouter
	grpcServerAddr string
	httpServerAddr string
}

func NewHttpServer(grpcServerAddress string, httpServerAddress string) *HttpServer {
	return &HttpServer{
		grpcServerAddr: grpcServerAddress,
		httpServerAddr: httpServerAddress,
		router: WeekHandlerRouter{
			&WeekHandlerRouterContext{},
		},
	}
}
func (s *HttpServer) Serve() {
	s.router.InitRoutes()

	s.router.DbClient = dbclient.NewLocalClient(s.grpcServerAddr, false)
	http.ListenAndServe(s.httpServerAddr, s.router.Rtr)
}
