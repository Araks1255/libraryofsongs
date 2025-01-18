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

	routes.GET("/:id", h.GetSong) // Обработка нужных запросов соответсвующими методами
	routes.GET("/", h.GetSongs)
	routes.POST("/", h.AddSong)
	routes.PUT("/:id", h.UpdateSong)
	routes.DELETE("/:id", h.DeleteSong)
}
