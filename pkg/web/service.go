package web

import (
	"context"
	"net/http"
	"ship-manage-local/pkg/config"
	"ship-manage-local/pkg/web/middleware"
	"ship-manage-local/pkg/web/router"

	"github.com/gin-gonic/gin"
)

type Service struct {
	httpSvr *http.Server
	engine  *gin.Engine
}

func NewService() *Service {
	g := gin.New()
	g.Use(
		gin.Recovery(),
		middleware.Logger(),
		middleware.Cors(),
	)

	svr := &Service{
		engine: g,
	}

	svr.httpSvr = &http.Server{
		Addr:    config.Get().Server.Addr,
		Handler: g,
	}

	return svr
}

func (srv *Service) Start(ctx context.Context) error {

	// 注册路由
	router.RegisterRouter(srv.engine)

	return srv.httpSvr.ListenAndServe()
}

func (srv *Service) Stop(ctx context.Context) error {

	return srv.httpSvr.Shutdown(ctx)
}
