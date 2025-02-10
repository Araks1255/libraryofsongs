package songs

import "github.com/gin-gonic/gin"

func (h handler) GetAlbumsOfUser(c *gin.Context) {
	user := c.Param("user")

	var albums []string
	h.DB.Raw("SELECT albums.name FROM albums " +
		"INNER JOIN user_albums ON albums.id = user_albums.album_id "+
		"INNER JOIN users ON user_albums.user_id = users.id "+
		"WHERE users.name = ?", user).Scan(&albums)

	response := ConvertToMap(albums)
	c.JSON(200, response)
}