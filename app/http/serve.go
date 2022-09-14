package http

import (
	"net/http"
	"time"

	"gosky/app/http/routers"
	"gosky/infra/conf"
	"gosky/infra/console"
	"gosky/infra/helpers"
)

func NewServe() *http.Server {
	router := routers.NewRouter()
	addr := conf.GetString("app.http_host") + ":" + conf.GetString("app.http_port")
	s := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	console.Success(time.Now().Format(helpers.CSTLayout) + " http listening on " + addr)
	return s
}
