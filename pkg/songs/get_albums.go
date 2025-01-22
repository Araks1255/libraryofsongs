package songs

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h handler) GetAlbumsOfGroup(c *gin.Context) {
	band := c.Param("band")

	var albums []string

	if result := h.DB.Raw("SELECT album FROM songs WHERE band = ?", band).Scan(&albums); result.Error != nil {
		c.AbortWithError(400, result.Error)
		log.Println(result.Error)
	}

	response := MakeResponse(albums)

	c.JSON(200, response)
}

func MakeResponse(albums []string) (response map[int]string) {
	// var resp string
	// for i := 0; i < len(albums); i++ {
	// 	resp += albums[i] + "\n"
	// }
	// return resp
	resp := make(map[int]string)

	for i := 0; i < len(albums); i++ {
		resp[i+1] = albums[i]
	}

	return resp
}
