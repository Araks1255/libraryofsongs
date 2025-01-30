package songs

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func (h handler) GetSongFile(c *gin.Context) {
	genre := strings.ToLower(c.Param("genre"))
	band := strings.ToLower(c.Param("band"))
	album := strings.ToLower(c.Param("album"))
	song := strings.ToLower(c.Param("song"))

	path := "H:/Мой диск/Проект Гоевый/Gin/libraryofsongs/list_of_songs/" + genre + "/" + band + "/" + album + "/" + song + ".mp3"

	c.File(path)
}
