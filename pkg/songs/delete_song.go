package songs

import (
	"github.com/Araks1255/libraryofsongs/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) DeleteSong(c *gin.Context) { // Хэндлер удаления песен
	var song models.Song // Переменная песни
	
	id := c.Param("id") // Параметр айди из ссылки

	if result := h.DB.First(&song, id); result.Error != nil { // Поиск песни по айди из ссылки
		c.AbortWithError(404, result.Error) // Обработка ошибок
		return
	}

	h.DB.Delete(&song) // Удаление найденной песни

	c.Status(200) // Возврат статус-кода
}