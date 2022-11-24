package cat

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CatController struct {
	catService CatService
}

func NewCatController() CatController {
	return CatController{
		catService: NewCatService(),
	}
}

func (cc CatController) GetCats(c *gin.Context) {
	cats, err := cc.catService.GetAllCats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "TODO: create a json error response",
		})
	}

	c.JSON(http.StatusOK, cats)
}
