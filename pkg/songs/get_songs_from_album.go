package songs

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h handler) GetSongsFromAlbum(c *gin.Context) { // Получение всех песен из альбома
	album := strings.ToLower(c.Param("album")) // Берем из ссылки название альбома, приводим к нижнему регистру и записываем в переменную

	var songs []string // Переменная для хранения найденных песен

	if result := h.DB.Raw("SELECT songs.name FROM songs INNER JOIN albums ON songs.album_id = albums.id WHERE albums.name = ?", album).Scan(&songs); result.Error != nil { // Получаем все имена песен из таблицы песен, совмещаем с таблицей альбомов, и берём только те песни, айди жанра которых совпадает с айди жанра из ссылки
		log.Println(result.Error) // Обрабатываем ошибки
		c.AbortWithError(404, result.Error)
		return
	}

	response := ConvertToMap(songs) // Конвертируем срез найденных песен в мапу, где ключами будет порядковый номер в виде числа
	c.JSON(200, response)           // Отправляем его
}

func ConvertToMap(slice []string) map[int]string { // Самописная функция преобразования среза в мапу, создающая ключи в виде порядкового номера каждому элементу
	resultingMap := make(map[int]string) // Инициализируем мапу

	for i := 0; i < len(slice); i++ { // Цикл, длящийся столько, сколько элементов в срезе
		resultingMap[i+1] = slice[i] // И записывающий каждый i элемент массива в мапу под ключом i+1
	}

	return resultingMap // Возвращаем мапу
}
