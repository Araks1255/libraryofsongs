package songs

import (
	"log"
	"mime/multipart"
	"os"
	"strings"

	"github.com/Araks1255/libraryofsongs/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) CreateSong(c *gin.Context) { // Хэндлер создания песни
	form, err := c.MultipartForm() // Получение мультипарт формы из запроса
	if err != nil {                // Проверка ошибок (вдркг кто-то JSON отправил)
		c.AbortWithStatusJSON(401, err) // Если такой нашелся, то выкидываем его со статусом 401 и ошибкой в формате JSON
		log.Println(err)                // И в лог выводим
		return                          // И завершаем выполение функции
	}

	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(422, err)
	}

	if err := CreateSong(form, h); err != nil { // Создаём песню при помощи самописной функции, в которую передаем форму и handler для взаимодействия с бд
		c.AbortWithStatusJSON(422, err) // Проверяем ошибки, и если есть, возвращаем её пользователю в виде JSONа со статусом 401
		return                          // Завершаем выполнение
	}

	if err := CreateSongFile(form, file, c); err != nil { // Создаём файл с песней с помощью самописной функции
		log.Println(err) // Обработка ошибок
		return
	}

	c.String(201, "Песня успешно создана") // Отправляем строку об успешном создании с кодом 201
}

func CreateSong(form *multipart.Form, h handler) error { // Функция создания песни. Принимает мультипарт форму и хэндлер
	var genre models.Genre                               // Переменная для жанра
	genre.Name = strings.ToLower(form.Value["genre"][0]) // Название жанра берём из формы под ключом genre

	var band models.Band                               // Группа
	band.Name = strings.ToLower(form.Value["band"][0]) // Название из формы

	var album models.Album                               // Альбом
	album.Name = strings.ToLower(form.Value["album"][0]) // Название из формы

	var song models.Song                               // Песня
	song.Name = strings.ToLower(form.Value["song"][0]) // Название из формы

	var genreID uint                                                                                   // Переменная для айди жанра
	if result := h.DB.Create(&genre); result.Error != nil && !IsRecordExists(result.Error, "genres") { // Создаём жанр (проверяем наличие ошибки и существование жанра в бд, если ошибка есть, и она вызвана не тем, что жанр уже сущвествует)
		return result.Error // Возвращаем ошибку
	} // Если оштбка есть, и она вызвана тем, что жанр уже существует, то всё норм
	if result := h.DB.Raw("SELECT id FROM genres WHERE name = ?", genre.Name).Scan(&genreID); result.Error != nil { // Получаем айди только чот созданного жанра, проводя поиск по его названию
		return result.Error // Возвращаем ошибку, если таковая имеется
	} else { // А если не имеется
		band.GenreID = genreID // То записываем айди жанра в поле genreID объекта band
	}

	var bandID uint                                                                                  // Айди группы
	if result := h.DB.Create(&band); result.Error != nil && !IsRecordExists(result.Error, "bands") { // Аналогично создаём группу (уже с айди жанра)
		return result.Error // Возврашаем ошибку, если есть
	}
	if result := h.DB.Raw("SELECT id FROM bands WHERE name = ?", band.Name).Scan(&bandID); result.Error != nil { // Ищем айди созданной грцппы
		return result.Error // Возвращаем ошибку, если возникла
	} else { // А если нет
		album.BandID = bandID // Записываем найденный айди в айди группы у альбома
	}

	var albumID uint                                                                                   // Айди альбома
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

func CreateSongFile(form *multipart.Form, file *multipart.FileHeader, c *gin.Context) error { // Создание файла с песней 
	genre := strings.ToLower(form.Value["genre"][0]) // Получаем из формы все названия и приводим к нижнему регистру
	band := strings.ToLower(form.Value["band"][0]) // Ну и в переменные записываем
	album := strings.ToLower(form.Value["album"][0])
	song := strings.ToLower(form.Value["song"][0])

	path := "H:/Мой диск/Проект Гоевый/Gin/libraryofsongs/list_of_songs/" + genre + "/" + band + "/" + album + "/" // Путь

	if err := os.MkdirAll(path, 0755); err != nil { // Создание папок со всеми элементами путя
		return err // Обработка ошибок
	}

	if _, err := os.Create(path + song + ".mp3"); err != nil { // Создаём пустой файл с названием песни в формате mp3 в нужной папке
		return err // Обработываем ошибки
	}

	if err := c.SaveUploadedFile(file, path+song+".mp3"); err != nil { // Сохраняем файл (методом к контексту из аргументов, файл тоже из аргументов)
		return err // Обработка ошибок
	}

	return nil // Если не было ошибок, то возвращаем nil
}
