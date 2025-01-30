package songs

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h handler) GetAlbumsOfBand(c *gin.Context) { // Получение альбомов группы 
	band := strings.ToLower(c.Param("band")) // Берем название группы из ссылки, приводим к нижнему регистру, записываем в переменную

	var bandID uint // Переменная для хранения айди группы

	if result := h.DB.Raw("SELECT id FROM bands WHERE name = ?", band).Scan(&bandID); result.Error != nil { // Берем айди из таблицы групп где название равно тому, что в ссылке, и сканим в переменную
		log.Println(result.Error) // Обрабатваем ошибки
		c.AbortWithStatusJSON(404, result.Error)
		return
	}

	var albums []string // Срез стрингов для хранения названий найденных альбомов

	if result := h.DB.Raw("SELECT name FROM albums WHERE band_id = ?", bandID).Scan(&albums); result.Error != nil { // Берем все названия из таблицы альбомов, где айди равен найденному ранее. Сканим в срез
		log.Println(result.Error) // Обработка ошибок
		c.AbortWithStatusJSON(404, result.Error)
		return
	}

	response := MakeResponse(albums) // Составляем ответ
	c.JSON(200, response) // Отправляем его
}
