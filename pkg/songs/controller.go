package songs

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct { // Структура хэндлера
	DB *gorm.DB // С бдшкой
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) { // Функция регистрации маршрутов
	h := handler{ // Создание объекта структуры хэндлера с бдшкой из аргумента
		DB: db,
	}

	routes := r.Group("/songs") // Роутер группа на субдомене /songs для роутера из аргументов

	// Create
	routes.POST("/", h.CreateSong)
	// Get
	routes.GET("/:album", h.GetSongsFromAlbum)
	routes.GET("/albums/:band", h.GetAlbumsOfBand)
	routes.GET("/bands/:genre", h.GetBandsByGenre)
	routes.GET("/genres", h.GetGenres)
	routes.GET("/file/:genre/:band/:album/:song", h.GetSongFile)
	// Find
	routes.GET("/song/:desiredSong", h.FindSong)
	routes.GET("/album/:desiredAlbum", h.FindAlbum)
	routes.GET("/band/:desiredBand", h.FindBand)
}
