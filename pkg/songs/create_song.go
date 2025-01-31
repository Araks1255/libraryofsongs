package songs

import (
	"log"
	"mime/multipart"
	"os"
	"strings"

	"github.com/Araks1255/libraryofsongs/pkg/common/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var pathToList string // Глобальная переменная с путём до списка песен (то место, где будут песни полученные храниться)

func (h handler) CreateSong(c *gin.Context) { // Хэндлер создания песни
	viper.SetConfigFile("./pkg/common/envs/.env")   // Устанавливаем конфиг файл с переменными окружения
	viper.ReadInConfig()                            // Считываем его
	pathToList = viper.Get("PATH_TO_LIST").(string) // Получаем строковое значение переменной окружения PATH_TO_LIST, и записываем его в глобальную переменную

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

	if result := h.DB.Create(&genre); result.Error != nil && !IsRecordExists(result.Error, "genres") { // Создаём жанр (проверяем наличие ошибки и существование жанра в бд, если ошибка есть, и она вызвана не тем, что жанр уже сущвествует)
		return result.Error // Возвращаем ошибку
	} // Если ошибка есть, и она вызвана тем, что жанр уже существует, то всё норм

	band.GenreID = genre.ID                                                                          // Добавляем в поле айди жанра у альбома значения поля айди у созданного жанра (при создании через Create(), переданный объект структуры дополняется первичным ключом оказывается)
	if result := h.DB.Create(&band); result.Error != nil && !IsRecordExists(result.Error, "bands") { // Аналогично создаём группу (уже с айди жанра)
		return result.Error // Возврашаем ошибку, если есть
	}

	album.BandID = band.ID                                                                             // Добавляем к альбому айди группы, к которой он относится
	if result := h.DB.Create(&album); result.Error != nil && !IsRecordExists(result.Error, "albums") { // Создаём альбом
		return result.Error // Возвращаем возникшкую ошибку
	}

	song.AlbumID = album.ID                                // Добавляем айди альбома к песне
	if result := h.DB.Create(&song); result.Error != nil { // Создаём песню
		return result.Error // Оббрабатываем ошибки
	}

	return nil // Если нигде раньше ошибок не возникло, возвращаем ничего
}

func IsRecordExists(err error, table string) bool { // Проверка существования чего-нибудь в таблице, принимает ошибку и имя таблицы
	if err.Error() == "ОШИБКА: повторяющееся значение ключа нарушает ограничение уникальности \"uni_"+table+"_name\" (SQLSTATE 23505)" { // Если строковое значение ошибки равно нужной нам ошибке в указанной таблице
		return true // Возвращаем тру, то есть - запись существует
	} else { // Иначе
		return false // Фолс
	}
}

func CreateSongFile(form *multipart.Form, file *multipart.FileHeader, c *gin.Context) error { // Создание файла с песней
	genre := strings.ToLower(form.Value["genre"][0]) // Получаем из формы все названия и приводим к нижнему регистру
	band := strings.ToLower(form.Value["band"][0])   // Ну и в переменные записываем
	album := strings.ToLower(form.Value["album"][0])
	song := strings.ToLower(form.Value["song"][0])

	path := pathToList + genre + "/" + band + "/" + album + "/" // Путь (начало из переменной окружения, остальное из песни)

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
