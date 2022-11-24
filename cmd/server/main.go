package main

import (
	"flag"

	"github.com/SlippyKorny/gin-api-template/internal/cat"
	"github.com/gin-gonic/gin"
)

var Version = "1.0.0"

var configFlag = flag.String("config", "./config/local.yaml", "path to the configuration file")

func main() {
	flag.Parse()

	// TODO: Load application configuration

	// Setup gin
	eng := gin.Default()
	buildHandlers(eng)

	eng.Run()
}

func buildHandlers(eng *gin.Engine) {
	cat.SetupCatEndpoints(eng)
}
