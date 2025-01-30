package songs

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h handler) GetSongsFromAlbum(c *gin.Context) {
	album := strings.ToLower(c.Param("album"))

	var albumID uint

	if result := h.DB.Raw("SELECT id FROM albums WHERE name = ?", album).Scan(&albumID); result.Error != nil {
		log.Println(result.Error)
		c.AbortWithStatusJSON(404, result.Error)
		return
	}

	var songs []string

	if result := h.DB.Raw("SELECT name FROM songs WHERE album_id = ?", albumID).Scan(&songs); result.Error != nil {
		log.Println(result.Error)
		c.AbortWithStatusJSON(404, result.Error)
		return
	}

	response := MakeResponse(songs)
	c.JSON(200, response)
}

func MakeResponse(names []string) (response map[int]string) {
	response = make(map[int]string)
	for i := 0; i < len(names); i++ {
		response[i+1] = names[i]
	}
	return response
}
