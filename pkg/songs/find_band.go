package songs

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h handler) FindBand(c *gin.Context) { // Хэндлер поиска группы
	desiredBand := strings.ToLower(c.Param("desiredBand")) // Получение группы из ссылки и приведение её к нижнему регистру

	var band, genre string // Переменные для хранения названий группы и жанра

	row := h.DB.Raw("SELECT bands.name, genres.name FROM bands "+ // Получаем ряд из сырого SQL запроса на поиск имени группы и жанра
		"INNER JOIN genres ON bands.genre_id = genres.id "+ // С помощью inner join
		"WHERE bands.name = ?", desiredBand).Row()

	if err := row.Scan(&band, &genre); err != nil { // Читаем ряд, и записываем его значения в переменные
		log.Println(err)                // Выводим ошибку в лог
		c.AbortWithStatusJSON(404, err) // Отправляем уведомление об ошибке
		return                          // Завершаем функцию
	} // Если конечно ошибки имеются

	response := ComposeBand(band, genre) // Составляем мапу из имени группы и жанра
	c.JSON(200, response)                // Отправляем её в виде JSONа вместе со статус-кодом 200
}

func ComposeBand(band, genre string) (bandMap map[string]string) { // Функция составления мапы из значений
	bandMap = make(map[string]string) // Инициализируем мапу с ключами стрингами и значениями стрингами
	bandMap["band"] = band            // Записываем значение аргумента band в мапу под ключом band
	bandMap["genre"] = genre          // Также с жанром
	return bandMap                    // Возвращаем мапу
}
