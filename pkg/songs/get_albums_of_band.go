package songs

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h handler) GetAlbumsOfBand(c *gin.Context) { // Получение альбомов группы
	band := strings.ToLower(c.Param("band")) // Берем название группы из ссылки, приводим к нижнему регистру, записываем в переменную

	var albums []string // Срез стрингов для хранения найденных альбомов

	if result := h.DB.Raw("SELECT albums.name FROM albums INNER JOIN bands ON albums.band_id = bands.id WHERE bands.name = ?", band).Scan(&albums); result.Error != nil { // Поиск всех альбомов, айди группы которых соответствует айди искомой группы
		log.Println(result.Error) // Обработка ошибок
		c.AbortWithStatusJSON(404, result.Error)
		return
	}

	response := ConvertToMap(albums) // Преобразуем срез в мапу
	c.JSON(200, response)            // Отправляем его
}
