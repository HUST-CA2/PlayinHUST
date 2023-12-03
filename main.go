package main

import (
	"PlayinHUST/common"
	"PlayinHUST/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	playinhust := gin.Default()

	playinhust.LoadHTMLGlob("./template/*")

	playinhust = routes.CollectRoutes(playinhust)

	err := playinhust.Run(common.Port)
	if err != nil {
		log.Fatal(err)
	}

}
