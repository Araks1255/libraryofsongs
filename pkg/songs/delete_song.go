package songs

import (
	"github.com/Araks1255/libraryofsongs/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) DeleteSong(c *gin.Context) {
	var song models.Song
	
	id := c.Param("id")

	if result := h.DB.First(&song, id); result.Error != nil {
		c.AbortWithError(404, result.Error)
		return
	}

	h.DB.Delete(&song)

	c.Status(200)
}