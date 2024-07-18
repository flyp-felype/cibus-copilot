package server

import (
	"copilot/internal/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg    *config.Config
	router *gin.Engine
}

func New(cfg *config.Config, router *gin.Engine) *Server {
	return &Server{
		cfg:    cfg,
		router: router,
	}
}

func (s *Server) Run() error {
	return s.router.Run(":" + s.cfg.Port)
}
