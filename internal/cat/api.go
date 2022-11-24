package cat

import (
	"fmt"
	"net/http"

	"github.com/SlippyKorny/gin-api-template/pkg/pagination"
	"github.com/gin-gonic/gin"
)

func SetupCatEndpoints(eng *gin.Engine) {
	c := NewCatController() // is it gonna get cleaned by the GC?
	eng.GET("/all-cats", c.GetCats)
	eng.GET("/paginated-cats", c.GetCatsPaged)
}

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

func (cc CatController) GetCatsPaged(c *gin.Context) {
	pr := pagination.NewPageRequestFromURL(c)
	fmt.Println(pr)
}
