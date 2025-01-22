package songs

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h handler) GetbandsByGenre(c *gin.Context) { // Хэндлер получения всех групп по жанру (корявый, но я всё равно пол проекта перелопатить собираюсь)
	genre := c.Param("genre") // Получаем жанр из ссылки

	var bands []string // Массив групп (стрингов просто)

	if result := h.DB.Raw("SELECT band FROM songs WHERE genre = ?", genre).Scan(&bands); result.Error != nil { // Поиск всех групп, где жанр равен тому, что в ссылке
		c.AbortWithError(400, result.Error) // Обработка ошибок
		log.Println(result.Error)
		return
	}

	response := MakeResponse(bands) // Создание ответа (функцией из другого файла, её бы по хорошему отдельно вынести, но я всё равно сейчас всё переделывать буду)

	c.JSON(200, response) // Отправляем ответ
}
