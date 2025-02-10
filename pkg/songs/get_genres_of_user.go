package songs

import "github.com/gin-gonic/gin"

func (h handler) GetGenresOfUser(c *gin.Context) {
	user := c.Param("user")

	var genres []string
	h.DB.Raw("SELECT genres.name FROM genres "+
		"INNER JOIN user_genres ON genres.id = user_genres.genre_id "+
		"INNER JOIN users ON user_genres.user_id = users.id "+
		"WHERE users.name = ?", user).Scan(&genres)

	response := ConvertToMap(genres)
	c.JSON(200, response)
}