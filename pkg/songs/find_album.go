package songs

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h handler) FindAlbum(c *gin.Context) {
	desiredAlbum := strings.ToLower(c.Param("desiredAlbum"))

	var album, band, genre string

	row := h.DB.Raw("SELECT albums.name, bands.name, genres.name FROM albums "+
		"INNER JOIN bands ON albums.band_id = bands.id "+
		"INNER JOIN genres ON bands.genre_id = genres.id "+
		"WHERE albums.name = ?", desiredAlbum).Row()

	if err := row.Scan(&album, &band, &genre); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(404, err)
		return
	}

	response := ComposeAlbum(album, band, genre)
	c.JSON(200, response)
}

func ComposeAlbum(album, band, genre string) (albumMap map[string]string) {
	albumMap = make(map[string]string)
	albumMap["album"] = album
	albumMap["band"] = band
	albumMap["genre"] = genre
	return albumMap
}