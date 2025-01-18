package songs

import (
	"github.com/Araks1255/libraryofsongs/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddSongRequestBody struct {
	Name          string `json:"name"`
	Band          string `json:"band"`
	Album         string `json:"album"`
	Genre         string `json:"genre"`
	YearOfRelease string `json:"yearOfRelease"`
}

func (h handler) AddSong(c *gin.Context) {
	var body AddSongRequestBody

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(401, err)
		return
	}

	var song models.Song

	MappingBodyToStruct(body, &song)

	if result := h.DB.Create(&song); result.Error != nil {
		c.AbortWithError(401, result.Error)
	}

	c.JSON(200, &song)
}

func MappingBodyToStruct(body AddSongRequestBody, song *models.Song) {
	song.Name = body.Name
	song.Band = body.Band
	song.Album = body.Album
	song.Genre = body.Genre
	song.YearOfRelease = body.YearOfRelease
}
