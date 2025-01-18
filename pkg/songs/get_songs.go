package songs

import (
	"github.com/Araks1255/libraryofsongs/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetSongs(c *gin.Context) { // Хэндлер получения всех песен
	var songs []models.Song // Переменная с массивом объектов структуры песни

	if result := h.DB.Find(&songs); result.Error != nil { // Поиск по бд и запись всего в переменную
		c.AbortWithError(404, result.Error) // Обработка ошибок
		return
	}

	c.JSON(200, &songs) // Возврат статус-кода и массива JSONов с песнями
}
