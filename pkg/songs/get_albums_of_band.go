package songs

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h handler) GetAlbumsOfBand(c *gin.Context) {
	band := c.Param("band")

	var bandID uint

	if result := h.DB.Raw("SELECT id FROM bands WHERE name = ?", band).Scan(&bandID); result.Error != nil {
		log.Println(result.Error)
		c.AbortWithStatusJSON(404, result.Error)
		return
	}

	var albums []string

	if result := h.DB.Raw("SELECT name FROM albums WHERE band_id = ?", bandID).Scan(&albums); result.Error != nil {
		log.Println(result.Error)
		c.AbortWithStatusJSON(404, result.Error)
		return
	}

	response := MakeResponse(albums)
	c.JSON(200, response)
}
