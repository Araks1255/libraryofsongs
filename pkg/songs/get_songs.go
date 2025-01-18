package songs

import (
	"github.com/Araks1255/libraryofsongs/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetSongs(c *gin.Context) {
	var songs []models.Song

	if result := h.DB.Find(&songs); result.Error != nil {
		c.AbortWithError(404, result.Error)
		return
	}

	c.JSON(200, &songs)
}
