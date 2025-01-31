package songs

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h handler) GetGenres(c *gin.Context) { // Хэндлер получения всех жанров
	var genres []string // Срез стрингов для найденных жанров

	if result := h.DB.Raw("SELECT name FROM genres").Scan(&genres); result.Error != nil { // Получаем все названия из таблицы жанров, сканируем в срез
		log.Println(result.Error) // Обработка ошибок
		c.AbortWithStatusJSON(404, result.Error)
		return
	}

	response := ConvertToMap(genres) // Создаём ответ с помощью самописной функции (из другого файла)
	c.JSON(200, response)            // Отправляем его
}
