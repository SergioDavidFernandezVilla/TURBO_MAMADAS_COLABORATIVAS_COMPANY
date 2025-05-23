package main

import (
	"github.com/SergioDavidFernandezVilla/TURBO_MAMADAS_COLABORATIVAS_COMPANY/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routers.Router(router)

	router.Run(":8080")
}
