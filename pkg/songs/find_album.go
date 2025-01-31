package songs

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h handler) FindAlbum(c *gin.Context) { // Поиск альбома
	desiredAlbum := strings.ToLower(c.Param("desiredAlbum")) // Искомый альбом (из ссылки, приведенный к нижнему регистру)

	var album, band, genre string // Переменные для хранения альбома, группы и жанра соответственно

	row := h.DB.Raw("SELECT albums.name, bands.name, genres.name FROM albums "+ // Выбираем имена альбома, группы и жанра
		"INNER JOIN bands ON albums.band_id = bands.id "+ // С помощью джоина таблицы альбомов к остальным вышестоящим
		"INNER JOIN genres ON bands.genre_id = genres.id "+
		"WHERE albums.name = ?", desiredAlbum).Row() // И берём только те значения, где название альбома равно тому, что было в ссылке. Получаем ряд, возвращенный в ответ на запрос

	if err := row.Scan(&album, &band, &genre); err != nil { // Сканим полученный ряд, и записываем всё в свои переменные
		log.Println(err) // Обрабатываем ошибки
		c.AbortWithStatusJSON(404, err)
		return
	}

	response := ComposeAlbum(album, band, genre) // Составляем мапу с альбомом самописной функцией
	c.JSON(200, response)                        // Отправляем
}

func ComposeAlbum(album, band, genre string) (albumMap map[string]string) { // Фунция составления мапы с альбомом
	albumMap = make(map[string]string) // Инициализируеум мапу со стринговым всем
	albumMap["album"] = album          // Сохраняем значение переменной album из аргументов под ключом album
	albumMap["band"] = band            // Аналогично
	albumMap["genre"] = genre
	return albumMap // Возвращаем эту мапу
}
