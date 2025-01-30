package songs

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func (h handler) GetSongFile(c *gin.Context) { // Получение файла с песней
	genre := strings.ToLower(c.Param("genre")) // Берём из ссылки всё что надо, приводим к нижнему регистру
	band := strings.ToLower(c.Param("band")) // И записываем в переменные
	album := strings.ToLower(c.Param("album"))
	song := strings.ToLower(c.Param("song"))

	path := "H:/Мой диск/Проект Гоевый/Gin/libraryofsongs/list_of_songs/" + genre + "/" + band + "/" + album + "/" + song + ".mp3" // Составляем путь к файлу

	c.File(path) // И отправляем файл по созданному пути
}
