package router

import (
	"dream/infrastructure/router/routes"
	"dream/interface/handlers"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, h handlers.AppHandler) *gin.Engine {
	r.Use(gin.Recovery())

	routes.RegisterDreamRoutes(r, h.Dream)

	return r
}
