package songs

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h handler) GetBandsByGenre(c *gin.Context) { // Получение групп по жанру
	genre := strings.ToLower(c.Param("genre")) // Получение жанра из ссылки (и приведение к нижнему регистру)

	var bands []string // Срез стрингов для хранения найденных групп

	if result := h.DB.Raw("SELECT bands.name FROM bands INNER JOIN genres ON bands.genre_id = genres.id WHERE genres.name = ?", genre).Scan(&bands); result.Error != nil { // Запись в этот срез всех групп, айди жанра которых совпадает с айди искомого жанра
		log.Println(result.Error) // Обработка ошибок
		c.AbortWithStatusJSON(404, result.Error)
		return
	}

	response := ConvertToMap(bands) // Преобразуем срез в мапу
	c.JSON(200, response)           // Отправляем её
}
