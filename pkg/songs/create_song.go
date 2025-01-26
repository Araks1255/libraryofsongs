package songs

import (
	"log"
	"mime/multipart"

	"github.com/Araks1255/libraryofsongs/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) CreateSong(c *gin.Context) { // Хэндлер создания песни
	form, err := c.MultipartForm() // Получение мультипарт формы из запроса
	if err != nil { // Проверка ошибок (вдркг кто-то JSON отправил)
		c.AbortWithStatusJSON(401, err) // Если такой нашелся, то выкидываем его со статусом 401 и ошибкой в формате JSON
		log.Println(err) // И в лог выводим
		return // И завершаем выполение функции
	}

	if err := CreateSong(form, h); err != nil { // Создаём песню при помощи самописной функции, в которую передаем форму и handler для взаимодействия с бд
		c.AbortWithStatusJSON(401, err) // Проверяем ошибки, и если есть, возвращаем её пользователю в виде JSONа со статусом 401
		return // Завершаем выполнение
	}

	c.String(201, "Песня успешно создана") // Отправляем строку об успешном создании с кодом 201
}

func CreateSong(form *multipart.Form, h handler) error { // Функция создания песни. Принимает мультипарт форму и хэндлер
	var genre models.Genre // Переменная для жанра
	genre.Name = form.Value["genre"][0] // Название жанра берём из формы под ключом genre

	var band models.Band // Группа
	band.Name = form.Value["band"][0] // Название из формы

	var album models.Album // Альбом
	album.Name = form.Value["album"][0] // Название из формы

	var song models.Song // Песня
	song.Name = form.Value["song"][0] // Название из формы

	var genreID uint // Переменная для айди жанра
	if result := h.DB.Create(&genre); result.Error != nil && !IsRecordExists(result.Error, "genres") { // Создаём жанр (проверяем наличие ошибки и существование жанра в бд, если ошибка есть, и она вызвана не тем, что жанр уже сущвествует)
		return result.Error // Возвращаем ошибку
	} // Если оштбка есть, и она вызвана тем, что жанр уже существует, то всё норм
	if result := h.DB.Raw("SELECT id FROM genres WHERE name = ?", genre.Name).Scan(&genreID); result.Error != nil { // Получаем айди только чот созданного жанра, проводя поиск по его названию
		return result.Error // Возвращаем ошибку, если таковая имеется
	} else { // А если не имеется
		band.GenreID = genreID  // То записываем айди жанра в поле genreID объекта band
	}

	var bandID uint // Айди группы
	if result := h.DB.Create(&band); result.Error != nil && !IsRecordExists(result.Error, "bands") { // Аналогично создаём группу (уже с айди жанра)
		return result.Error // Возврашаем ошибку, если есть
	}
	if result := h.DB.Raw("SELECT id FROM bands WHERE name = ?", band.Name).Scan(&bandID); result.Error != nil { // Ищем айди созданной грцппы
		return result.Error // Возвращаем ошибку, если возникла
	} else { // А если нет
		album.BandID = bandID // Записываем найденный айди в айди группы у альбома
	}

	var albumID uint // Айди альбома
	if result := h.DB.Create(&album); result.Error != nil && !IsRecordExists(result.Error, "albums") { // Создаём альбом
		return result.Error // Возвращаем возникшкую ошибку
	}
	if result := h.DB.Raw("SELECT id FROM albums WHERE name = ?", album.Name).Scan(&albumID); result.Error != nil { // Ищем айди созданного аольбома
		return result.Error // Ошибка
	} else { // Если нет ошибка
		song.AlbumID = albumID // Записываем айди альбома в айди альбома песни
	}

	if result := h.DB.Create(&song); result.Error != nil { // Создаём песню
		return result.Error // Оббрабатываем ошибки
	}

	return nil // Если нигде раньше ошибок не возникло, возвращаем ничего
}

func IsRecordExists(err error, table string) bool { // Проверка существования чего-нибудь в таблие, принимает ошибку и имя таблицы
	if err.Error() == "ОШИБКА: повторяющееся значение ключа нарушает ограничение уникальности \"uni_"+table+"_name\" (SQLSTATE 23505)" { // Если строковое значение ошибки равно нужной нам ошибке в указанной таблице
		return true // Возвращаем тру, то есть - запись существует
	} else { // Иначе
		return false // Фолс
	}
}
