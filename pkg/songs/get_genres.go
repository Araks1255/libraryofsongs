package songs

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h handler) GetGenres(c *gin.Context) {
	var genres []string

	if result := h.DB.Raw("SELECT name FROM genres").Scan(&genres); result.Error != nil {
		log.Println(result.Error)
		c.AbortWithStatusJSON(404, result.Error)
		return
	}

	response := MakeResponse(genres)
	c.JSON(200, response)
}
