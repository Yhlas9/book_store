package service

import (
	"github.com/gin-gonic/gin"
	config2 "github.com/Yhlas9/book_store/config"
	"github.com/Yhlas9/book_store/service/handler"
)

func InitHandlers(app *gin.Engine) {
	config := config2.ParseConfig()

	appRoutes := handler.InitHandler(config)

	app.GET("/books", appRoutes.GetBooks)
	app.POST("/books", appRoutes.StoreBook)
	app.DELETE("/books/:id", appRoutes.DeleteBook)
	app.PUT("/books/:id", appRoutes.UpdateBook)
	app.GET("/books/:id", appRoutes.GetBook)
}
