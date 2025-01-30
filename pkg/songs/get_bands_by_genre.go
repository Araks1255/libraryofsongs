package songs

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h handler) GetBandsByGenre(c *gin.Context) {
	genre := strings.ToLower(c.Param("genre"))

	var genreID uint

	if result := h.DB.Raw("SELECT id FROM genres WHERE name = ?", genre).Scan(&genreID); result.Error != nil {
		log.Println(result.Error)
		c.AbortWithStatusJSON(404, result.Error)
		return
	}

	var bands []string

	if result := h.DB.Raw("SELECT name FROM bands WHERE genre_id = ?", genreID).Scan(&bands); result.Error != nil {
		log.Println(result.Error)
		c.AbortWithStatusJSON(404, result.Error)
		return
	}

	response := MakeResponse(bands)
	c.JSON(200, response)
}
