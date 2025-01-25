package songs

import (
	"log"
	"mime/multipart"

	"github.com/Araks1255/libraryofsongs/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) AddSong(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.AbortWithError(401, err)
		log.Println(err)
		return
	}

	if err := CreateSong(form, h); err != nil {
		c.AbortWithError(401, err)
	}

	c.String(201, "До сюда не дойдёт")
}

func CreateSong(form *multipart.Form, h handler) error {
	var genre models.Genre
	genre.Name = form.Value["genre"][0]

	var band models.Band
	band.Name = form.Value["band"][0]

	var album models.Album
	album.Name = form.Value["album"][0]

	var song models.Song
	song.Name = form.Value["song"][0]

	var createdGenre models.Genre
	if result := h.DB.Create(&genre); result.Error != nil {
		return result.Error
	}
	if result := h.DB.Raw("SELECT * FROM genres WHERE name = ?", genre.Name).Scan(&createdGenre); result.Error != nil {
		log.Println(result.Error)
		return result.Error
	} else {
		band.GenreID = createdGenre.ID
		log.Println("вах")
	}

	var createdBand models.Band
	if result := h.DB.Create(&band); result.Error != nil {
		return result.Error
	}
	if result := h.DB.Raw("SELECT * FROM bands WHERE name = ?", band.Name).Scan(&createdBand); result.Error != nil {
		log.Println(result.Error)
		return result.Error
	} else {
		album.BandID = createdBand.ID
		log.Println("вахх")
	}

	var createdAlbum models.Album
	if result := h.DB.Create(&album); result.Error != nil {
		return result.Error
	}
	if result := h.DB.Raw("SELECT * FROM albums WHERE name = ?", album.Name).Scan(&createdAlbum); result.Error != nil {
		log.Println(result.Error)
		return result.Error
	} else {
		song.AlbumID = createdAlbum.ID
		log.Println("ваххх")
	}

	if result := h.DB.Create(&song); result.Error != nil {
		return result.Error
	}

	return nil
}
