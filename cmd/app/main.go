package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Yhlas9/book_store/service"
)

func main() {
	app := gin.Default()

	service.InitHandlers(app)

	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}
