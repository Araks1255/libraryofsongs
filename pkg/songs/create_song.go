package songs

import (
	"errors"
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

	file, err := c.FormFile("file") // Получение файла из запроса
	if err != nil {                 // Проверка на ошибки
		log.Println(err) // Обработка ошибок
		c.AbortWithStatusJSON(422, err)
	}

	const numberOfGorutines uint8 = 2              // Переменная для хранения количества горутин (uint8 потому-что их не может быть меньше нуля и наврядли будет больше 255)
	errChan := make(chan error, numberOfGorutines) // Создание буферизированного канала для ошибок, емкостью количества горутин (каждая горутина возвращает одну ошибку)

	go func() { // Запуск горутины в виде анонимной функции
		err := CreateSong(form, h) // Которая вызывает функцию создания песни
		errChan <- err             // И записывает возвращаемую ей ошибку в канал
	}()

	go func() { // Запуск второй горутины в виде анонимной функции
		err := CreateSongFile(form, file, c) // Вызов в ней функции создания файла с песней
		errChan <- err                       // И запись ошибки в канал
	}()

	var errs [numberOfGorutines]error       // Создания массива ошибок длиной количества горутин
	var i uint8                             // Просто переменная для перебора в цикле, создана отдельно, чтобы могла иметь тип uint8, а не int (опять же, отрицательной она быть не должна, и больше 255 не будет, потому-что увеличивается до тех пор, пока меньше числа горутин)
	for i = 0; i < numberOfGorutines; i++ { // Цикл, делающий столько иттераций, сколько есть горутин
		errs[i] = <-errChan // Чтение из канала в элемент i массива
		if errs[i] != nil { // Если прочитанное значение не равно nil (в случае отсутсвия ошибок при вызове функций в горутинах, в канал возвращается nil)
			log.Println(errs[i]) // Вывод ошибки, прочитанной из канала
			c.AbortWithStatusJSON(422, gin.H{
				"error": errs[i].Error(),
			}) // Разрыв соеденения с той же самой ошибкой
			return // Выход из функции
		}
	}

	c.String(201, "Песня успешно создана") // Если все функции вернули nil, то отправляем уведомление об успешном создании песни
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

	if genreID := IsRecordExists(h, "genres", genre.Name); genreID != 0 { // Проверяем наличие жанра в бд самописной функцией
		band.GenreID = genreID // Если найденный айди не равен 0 (то есть, запись существует), присваиваем полю айди жанра у объекта группы найденный айди
	} else { // Если найденный айди равен нулю (запись не существует)
		result := h.DB.Create(&genre) // Создаём жанр по ранее заполненному объекту структуры
		if result.Error != nil {      // Обрабатываем ошибки
			return result.Error
		}
		band.GenreID = genre.ID // Записываем айди созданного жанра в поле айди жанра у объекта группы
	}

	if bandID := IsRecordExists(h, "bands", band.Name); bandID != 0 { // Тут всё аналогично жанру
		album.BandID = bandID
	} else {
		result := h.DB.Create(&band)
		if result.Error != nil {
			return result.Error
		}
		album.BandID = band.ID
	}

	if albumID := IsRecordExists(h, "albums", album.Name); albumID != 0 { // Тут тоже
		song.AlbumID = albumID
	} else {
		result := h.DB.Create(&album)
		if result.Error != nil {
			return result.Error
		}
		song.AlbumID = album.ID
	}

	if songID := IsRecordExists(h, "songs", song.Name); songID != 0 { // Проверяем наличие песни в бд
		return errors.New("Песня уже существует") // Если существует, то возвращаем ошибку с уведомлением об этом
	} else { // Если нет
		result := h.DB.Create(&song) // Создаём
		if result.Error != nil {     // Обрабатываем ошибки
			return result.Error
		}
	}

	return nil // Если нигде раньше ошибок не возникло, возвращаем ничего
}

func IsRecordExists(h handler, table, name string) (id uint) { // Функция для проверки существования записи в бд. Возвращает id записи, если она существует, 0 - если нет
	var foundID uint                                                                                             // Переменная для хранения найденного айди
	if result := h.DB.Raw("SELECT id FROM "+table+" WHERE name = ?", name).Scan(&foundID); result.Error != nil { // Поиск айди в таблице из аргументов по названию из аргументов (если записи не существует, вернётся 0)
		return 0 // Если возникла ошибка, возвращаем тоже 0, ведь в любом случае это значит то, что записиси не существует
	}
	return foundID // Возвращаем переменную с найденным айди
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

	fileWithSong, err := os.Create(path + song + ".mp3") // Создаём пустой файл для песни
	if err != nil {                                      // Обрабатываем ошибки
		return err
	}

	if err := c.SaveUploadedFile(file, path+song+".mp3"); err != nil { // Сохраняем файл (методом к контексту из аргументов, файл тоже из аргументов) (этот метод открывает файл, но не закрывает его)
		return err // Обработка ошибок
	}
	defer fileWithSong.Close() // Закрываем файл, открытый методом выше

	return nil // Если не было ошибок, то возвращаем nil
}
