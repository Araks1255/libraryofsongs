package songs

import (
	"github.com/Araks1255/libraryofsongs/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddSongRequestBody struct { // Структура JSON тела POST запроса
	Name          string `json:"name"`
	Band          string `json:"band"`
	Album         string `json:"album"`
	Genre         string `json:"genre"`
	YearOfRelease string `json:"yearOfRelease"`
}

func (h handler) AddSong(c *gin.Context) { // Хэндлер добавления песни
	var body AddSongRequestBody // Переменная для тела запроса

	if err := c.BindJSON(&body); err != nil { // Биндим JSON из запроса в неё
		c.AbortWithError(401, err) // Обрабатываем ошибки
		return
	}

	var song models.Song // Переменная для самой песни

	MappingBodyToStruct(body, &song) // Запись данных из тела запроса в структуру песни

	if result := h.DB.Create(&song); result.Error != nil { // Создание в бд созданной песни
		c.AbortWithError(401, result.Error) // Обработка ошибок
	}

	c.JSON(200, &song) // Отправка успешного статус-кода и структуры песни в виде JSONа
}

func MappingBodyToStruct(body AddSongRequestBody, song *models.Song) { // Функция записи тела запроса в объект структуры песни
	song.Name = body.Name // Просто приравниваем все параметры 
	song.Band = body.Band // И song тут это указатель на объект структуры, так что он изменяется
	song.Album = body.Album
	song.Genre = body.Genre
	song.YearOfRelease = body.YearOfRelease
}
