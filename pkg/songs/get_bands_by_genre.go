package songs

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h handler) GetBandsByGenre(c *gin.Context) { // Получение групп по жанру
	genre := strings.ToLower(c.Param("genre")) // Получение жанра из ссылки (и приведение к нижнему регистру)

	var genreID uint // Переменная для айди найденного жанра

	if result := h.DB.Raw("SELECT id FROM genres WHERE name = ?", genre).Scan(&genreID); result.Error != nil { // Берем айди из таблицы genres, где название как в ссылке
		log.Println(result.Error) // Обработка ошибок
		c.AbortWithStatusJSON(404, result.Error)
		return
	}

	var bands []string // Срез стрингов для названий альбомов

	if result := h.DB.Raw("SELECT name FROM bands WHERE genre_id = ?", genreID).Scan(&bands); result.Error != nil { // Берем все названия из таблицы групп, где айди жанра как тот, что нашли ранее
		log.Println(result.Error) // Обработка ошибок
		c.AbortWithStatusJSON(404, result.Error)
		return
	}

	response := MakeResponse(bands) // Создаём ответ 
	c.JSON(200, response) // Отправляем его
}
