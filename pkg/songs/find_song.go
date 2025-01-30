package songs

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h handler) FindSong(c *gin.Context) { // Хэндлер поиска песни
	desiredSong := strings.ToLower(c.Param("desiredSong")) // Получаем из ссылки название искомой песни и приводим его к нижнему регистру

	var song, album, band, genre string // Объявляем переменные для хранения названий песни, альбома, группы и жанра

	row := h.DB.Raw("SELECT songs.name, albums.name, bands.name, genres.name FROM songs "+ // Делаем сырой SQL запрос, в котором получаем названия всего
		"INNER JOIN albums ON songs.album_id = albums.id "+ // Используя INNER JOIN
		"INNER JOIN bands ON albums.band_id = bands.id "+
		"INNER JOIN genres ON bands.genre_id = genres.id "+
		"WHERE songs.name = ?", desiredSong).Row() // И в конце получаем ряд значений (ну и в переменную записываем его)

	if err := row.Scan(&song, &album, &band, &genre); err != nil { // Достаём из ряда все нужные значения, и записываем в нужные переменные
		log.Println(err)                // Если есть ошибка, то выводим её в лог
		c.AbortWithStatusJSON(404, err) // И выкидываем пользователя с кодом 404 и ошибкой в виде JSONа
		return                          // Завершаем выполнение функции
	}

	response := СomposeSong(song, album, band, genre) // Создаём мапу с песней
	c.JSON(200, response)                             // Отправляем её в виде JSONа
}

func СomposeSong(song, album, band, genre string) (songMap map[string]string) { // Функция составления песни из названий всего к ней причастного
	songMap = make(map[string]string) // Инициализируем мапу
	songMap["song"] = song            // Записываем под ключом song значение аргумента song
	songMap["album"] = album          // Ну и так далее
	songMap["band"] = band
	songMap["genre"] = genre

	return songMap // Возвращаем получившуюся мапу
}
