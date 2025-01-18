package songs

import (
	"github.com/Araks1255/libraryofsongs/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetSong (c *gin.Context) { // Хэндлер получения одной песни по айди
	id := c.Param("id") // Получение айди из ссылки

	var song models.Song // Переменная для песни

	if result := h.DB.First(&song, id); result.Error != nil { // Поиск песни по айди
		c.AbortWithError(404, result.Error) // Обработка ошибок
		return
	}

	c.JSON(200, &song) // Возврат статус-кода и структуры песни в виде JSONа
}