package songs

import "github.com/gin-gonic/gin"

func (h handler) GetSongsOfUser(c *gin.Context) {
	user := c.Param("user")

	var songs []string
	h.DB.Raw("SELECT songs.name FROM songs "+
		"INNER JOIN user_songs ON songs.id = user_songs.song_id "+
		"INNER JOIN users ON user_songs.user_id = users.id "+
		"WHERE users.name = ?", user).Scan(&songs)

	response := ConvertToMap(songs)
	c.JSON(200, response)
}
