package main

import (
	"PlayinHUST/common"
	"PlayinHUST/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	playinhust := gin.Default()

	DB := common.GetDB()
	defer DB.Close()

	playinhust.LoadHTMLGlob("./template/*")

	playinhust = routes.CollectRoutes(playinhust)

	err := playinhust.Run(common.Port)
	if err != nil {
		log.Fatal(err)
	}

}
