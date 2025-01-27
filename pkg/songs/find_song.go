package songs

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h handler) FindSong(c *gin.Context) {
	desiredSong := c.Param("desiredSong")

	var song string
	var album string
	var band string
	var genre string

	row := h.DB.Raw("SELECT songs.name, albums.name, bands.name, genres.name FROM songs "+
		"INNER JOIN albums ON songs.album_id = albums.id "+
		"INNER JOIN bands ON albums.band_id = bands.id "+
		"INNER JOIN genres ON bands.genre_id = genres.id "+
		"WHERE songs.name = ?", desiredSong).Row()

	if err := row.Scan(&song, &album, &band, &genre); err != nil {
		log.Println(song, album, band, genre)
	}

	response := composeSong(song, album, band, genre)
	c.JSON(200, response)
}

func composeSong(song, album, band, genre string) (songMap map[string]string) {
	songMap = make(map[string]string)
	songMap["song"] = song
	songMap["album"] = album
	songMap["band"] = band
	songMap["genre"] = genre
	return songMap
}
