package songs

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h handler) GetAlbumsOfGroup(c *gin.Context) { // Хэндлер получения всех альбомов группы
	band := c.Param("band") // Параметр названия группы из ссылки

	var albums []string // Переменная для хранения найденных альбомов (массив стрингов)

	if result := h.DB.Raw("SELECT album FROM songs WHERE band = ?", band).Scan(&albums); result.Error != nil { // Поиск всех значений столбца album таблицы songs в которых группа равна группе из сылки
		c.AbortWithError(400, result.Error) // Если есть ошибки, обрываем соеденение с ошибкой
		log.Println(result.Error)           // И в лог выводим
		return
	}

	response := MakeResponse(albums) // Создаём ответ с помощью самописной функции

	c.JSON(200, response) // Отправляем его
}

func MakeResponse(albums []string) (response map[int]string) { // Функия создания ответа
	resp := make(map[int]string) // Иницализиркем мапу с ключами числами (номерами альбомов) и стрингами (названиями альбомов)

	for i := 0; i < len(albums); i++ { // Через цикл перебираем массив albums
		resp[i+1] = albums[i] // И записываем в каждый номер (начиная с 1) по элементу массива альбомов
	}

	return resp // Возвращаем мапу
}
