package songs

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h handler) GetSongsFromAlbum(c *gin.Context) { // Получение всех песен из альбома
	album := strings.ToLower(c.Param("album")) // Берем из ссылки название альбома, приводим к нижнему регистру и записываем в переменную

	var albumID uint // Объявляем переменную для хранения id альбома

	if result := h.DB.Raw("SELECT id FROM albums WHERE name = ?", album).Scan(&albumID); result.Error != nil { // Ищем айди альбома с именем из ссылки и сканим в переменную
		log.Println(result.Error) // Обработка ошибок
		c.AbortWithStatusJSON(404, result.Error)
		return
	}

	var songs []string // Переменная для хранения названий найденных песен

	if result := h.DB.Raw("SELECT name FROM songs WHERE album_id = ?", albumID).Scan(&songs); result.Error != nil { // Берём из таблицы песен все названия, где айди альбома равен найденному ранее
		log.Println(result.Error) // Обрабатываем ошибки
		c.AbortWithStatusJSON(404, result.Error)
		return
	}

	response := MakeResponse(songs) // Составляем ответ самописной функцией
	c.JSON(200, response) // Отправляем его
}

func MakeResponse(names []string) (response map[int]string) { // Функция создания ответа (в нескольких хэндлерах используется)
	response = make(map[int]string) // Инициализируем мапу с целочисленными ключами и стринговыми значениями
	for i := 0; i < len(names); i++ { // В цикле, зависящем от длины среза из аргументов
		response[i+1] = names[i] // Приравниваем порядковый номер из ключа к значению под индексом массива из аргументов i
	}
	return response // Возвращаем созданную мапу
}
