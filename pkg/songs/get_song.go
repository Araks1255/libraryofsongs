package songs

import (
	"log"

	"github.com/Araks1255/libraryofsongs/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetSong(c *gin.Context) { // Хэндлер получения одной песни по айди
	id := c.Param("id") // Получение айди из ссылки

	var song models.Song // Переменная для песни

	if result := h.DB.First(&song, id); result.Error != nil { // Поиск песни по айди
		log.Println(result.Error)
		c.AbortWithError(404, result.Error) // Обработка ошибок
		return
	}

	c.File(FindSongFile(song))
}

func FindSongFile(song models.Song) string {
	path := "H:/Мой диск/Проект Гоевый/Gin/libraryofsongs/songs/genres/" + song.Genre + "/" + song.Band + "/" + song.Album + "/" + song.Name + ".mp3"
	return path
}
