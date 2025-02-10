package songs

import "github.com/gin-gonic/gin"

func (h handler) GetBandsOfUser(c *gin.Context) {
	user := c.Param("user")

	var bands []string
	h.DB.Raw("SELECT bands.name FROM bands "+
		"INNER JOIN user_bands ON bands.id = user_bands.band_id "+
		"INNER JOIN users ON user_bands.user_id = users.id "+
		"WHERE users.name = ?", user).Scan(&bands)

	response := ConvertToMap(bands)
	c.JSON(200, response)
}