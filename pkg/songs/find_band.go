package songs

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h handler) FindBand(c *gin.Context) {
	desiredBand := strings.ToLower(c.Param("desiredBand"))

	var band, genre string

	row := h.DB.Raw("SELECT bands.name, genres.name FROM bands "+
		"INNER JOIN genres ON bands.genre_id = genres.id "+
		"WHERE bands.name = ?", desiredBand).Row()

	if err := row.Scan(&band, &genre); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(404, err)
		return
	}

	response := ComposeBand(band, genre)
	c.JSON(200, response)
}

func ComposeBand(band, genre string) (bandMap map[string]string) {
	bandMap = make(map[string]string)
	bandMap["band"] = band
	bandMap["genre"] = genre
	return bandMap
}