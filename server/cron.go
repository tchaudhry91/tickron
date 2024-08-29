package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CronServer struct {
	listenAddr string
	srv        *echo.Echo
}

func (s *CronServer) Start() {
	s.srv.Start(s.listenAddr)
}

func NewCronServer(listenAddr string) *CronServer {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	s := &CronServer{
		listenAddr: listenAddr,
		srv:        e,
	}
	s.addRoutes()
	return s
}
