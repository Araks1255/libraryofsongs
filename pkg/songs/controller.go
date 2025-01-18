package songs

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := handler{
		DB: db,
	}

	routes := r.Group("/songs")

	routes.GET("/:id", h.GetSong)
	routes.GET("/", h.GetSongs)
	routes.POST("/", h.AddSong)
	routes.PUT("/:id", h.UpdateSong)
	routes.DELETE("/:id", h.DeleteSong)
}
