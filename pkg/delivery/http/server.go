package http_server

import (
	"course-backend/pkg/database"
	"course-backend/pkg/router"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	route    *gin.Engine
	logger   *slog.Logger
	database *database.Database
}

func (self *HttpServer) RouteAll() {
	router.RouteAPI(self.route, self.database)
}

func (self *HttpServer) SetupLogger(logger *slog.Logger) {
	self.logger = logger
}

func (self *HttpServer) Start(addr string) error {
	self.logger.Info("Server started on %s", addr)
	return self.route.Run(addr)
}
