package songs

import (
	"github.com/Araks1255/libraryofsongs/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type UpdateSongRequestBody struct { // Структура JSON тела запроса
	Name          string `json:"name"`
	Band          string `json:"band"`
	Album         string `json:"album"`
	Genre         string `json:"genre"`
	YearOfRelease string `json:"yearOfRelease"`
}

func (h handler) UpdateSong(c *gin.Context) { // Хэндлер обновления песни
	var body UpdateSongRequestBody // Переменная для тела запроса

	id := c.Param("id") // Айди из ссылки

	if err := c.BindJSON(&body); err != nil { // Биндинг тела запроса в переменную
		c.AbortWithError(404, err) // Обработка ошибок
	}

	var song models.Song // Переменная для песни

	if result := h.DB.First(&song, id); result.Error != nil { // Поиск записи в бд по айди из ссылки
		c.AbortWithError(404, result.Error) // Обработка ошибок
		return
	}

	UpdateSongData(&song, body) // Обновление данных песни из бд данными из запроса

	h.DB.Save(&song) // Сохранение изменений

	c.JSON(200, &song) // Статус код и изменённая песня в ответе
}

func UpdateSongData(song *models.Song, body UpdateSongRequestBody) { // Функция Обновления записи в бд данными из тела запроса. Принимает запись, которую надо обновить и забинженное в структуру тело запроса
	song.Name = body.Name // Просто тупая замена данных
	song.Band = body.Band // Также по ссылке на запись в бд
	song.Album = body.Album
	song.Genre = body.Genre
	song.YearOfRelease = body.YearOfRelease
}
