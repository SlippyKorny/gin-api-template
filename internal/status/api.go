package status

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupStatusEndpoints(eng *gin.Engine) {
	c := NewStatusController()
	eng.GET("/status", c.GetStatus)
}

type StatusController struct {
}

func NewStatusController() StatusController {
	// TODO: Inject some config service into controller
	return StatusController{}
}

func (sc StatusController) GetStatus(c *gin.Context) {
	// TODO: Read from config file
	c.JSON(http.StatusOK, gin.H{
		"status":  "SomeStatusMessage",
		"env":     "environmentName",
		"version": "versionName",
	})
}
