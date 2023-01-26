package routes

import (
	"dream/interface/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterDreamRoutes(r *gin.Engine, h handlers.DreamHandler) *gin.Engine {
	r.GET("dream/:id")
	r.POST("dream", h.Add)
	r.PUT("dream")
	r.DELETE("dream/:id")

	return r
}
