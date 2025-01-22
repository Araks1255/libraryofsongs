package songs

import (
	"log"
	"mime/multipart"
	"os"

	"github.com/Araks1255/libraryofsongs/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddSongRequestBody struct { // Структура JSON тела POST запроса
	Name          string `json:"name"`
	Band          string `json:"band"`
	Album         string `json:"album"`
	Genre         string `json:"genre"`
	YearOfRelease string `json:"yearOfRelease"`
}

func (h handler) AddSong(c *gin.Context) { // Хэндлер добавления песни
	form, err := c.MultipartForm()
	if err != nil {
		c.AbortWithError(401, err)
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithError(401, err)
		return
	}

	var song models.Song // Переменная для самой песни

	MappingBodyToStruct(form, &song) // Запись данных из тела запроса в структуру песни

	if path, err := CreateSongFile(song, file); err != nil {
		c.AbortWithError(401, err)
	} else {
		c.SaveUploadedFile(file, path)
	}

	if result := h.DB.Create(&song); result.Error != nil { // Создание в бд созданной песни
		c.AbortWithError(401, result.Error) // Обработка ошибок
	}

	c.JSON(200, &song) // Отправка успешного статус-кода и структуры песни в виде JSONа
}

func MappingBodyToStruct(form *multipart.Form, song *models.Song) { // Функция записи тела запроса в объект структуры песни
	song.Name = form.Value["name"][0]
	song.Band = form.Value["band"][0]
	song.Album = form.Value["album"][0]
	song.Genre = form.Value["genre"][0]
	song.YearOfRelease = form.Value["yearOfRelease"][0]
}

func CreateSongFile(song models.Song, file *multipart.FileHeader) (pathToEmptyFile string, err1 error) {
	path := "H:/Мой диск/Проект Гоевый/Gin/libraryofsongs/songs/genres/" + song.Genre + "/" + song.Band + "/" + song.Album

	if err := os.MkdirAll(path, 0755); err != nil {
		log.Println(err)
		return "", err
	}

	_, err := os.Create(path + "/" + song.Name + ".mp3")
	if err != nil {
		return "", err
	}

	return path + "/" + song.Name + ".mp3", nil
}
