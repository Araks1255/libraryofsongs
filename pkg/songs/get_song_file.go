package songs

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (h handler) GetSongFile(c *gin.Context) { // Получение файла с песней
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()
	pathToList := viper.Get("PATH_TO_LIST").(string)

	genre := strings.ToLower(c.Param("genre")) // Берём из ссылки всё что надо, приводим к нижнему регистру
	band := strings.ToLower(c.Param("band"))   // И записываем в переменные
	album := strings.ToLower(c.Param("album"))
	song := strings.ToLower(c.Param("song"))

	path := pathToList + genre + "/" + band + "/" + album + "/" + song + ".mp3" // Составляем путь к файлу

	c.File(path) // И отправляем файл по созданному пути
}
