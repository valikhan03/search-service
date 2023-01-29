package api

import (
	"search-service/api/v1/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, handler *handlers.Handler) {
	r.GET("/auctions/:id", handler.GetAuction)
	r.GET("/products/:id", handler.GetProduct)

	r.POST("/search", handler.Search)

	r.GET("/suggestions", handler.Suggestions)

	r.GET("/health", handler.HealthCheck)
}
