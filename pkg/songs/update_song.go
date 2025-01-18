package songs

import (
	"github.com/Araks1255/libraryofsongs/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type UpdateSongRequestBody struct {
	Name          string `json:"name"`
	Band          string `json:"band"`
	Album         string `json:"album"`
	Genre         string `json:"genre"`
	YearOfRelease string `json:"yearOfRelease"`
}

func (h handler) UpdateSong(c *gin.Context) {
	var body UpdateSongRequestBody

	id := c.Param("id")

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(404, err)
	}

	var song models.Song

	if result := h.DB.First(&song, id); result.Error != nil {
		c.AbortWithError(404, result.Error)
		return
	}

	UpdateSongData(&song, body)

	h.DB.Save(&song)

	c.JSON(200, &song)
}
func UpdateSongData(song *models.Song, body UpdateSongRequestBody) {
	song.Name = body.Name
	song.Band = body.Band
	song.Album = body.Album
	song.Genre = body.Genre
	song.YearOfRelease = body.YearOfRelease
}
