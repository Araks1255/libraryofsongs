package songs

import (
	"github.com/Araks1255/libraryofsongs/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetSong (c *gin.Context) {
	id := c.Param("id")

	var song models.Song

	if result := h.DB.First(&song, id); result.Error != nil {
		c.AbortWithError(404, result.Error)
		return
	}

	c.JSON(200, &song)
}