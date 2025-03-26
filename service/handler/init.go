package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/Yhlas9/book_store/config"
	"github.com/Yhlas9/book_store/service/service"
	"github.com/Yhlas9/book_store/service/structure"
	"strconv"
)

type Handler struct {
	Service *service.Service
}

func InitHandler(config config.AppConfig) *Handler {
	return &Handler{
		Service: service.NewService(config),
	}
}

func (h *Handler) GetBooks(ctx *gin.Context) {
	books := h.Service.GetBooks(ctx)
	SuccessResponse(ctx, gin.H{"books": books})
}

func (h *Handler) StoreBook(ctx *gin.Context) {
	var request structure.StoreBookRequest
	if err := ctx.ShouldBind(&request); err != nil {
		BadRequest(ctx, err)
		return
	}

	err := h.Service.StoreBook(ctx, structure.Book{
		Title:  request.Title,
		Author: request.Author,
	})

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	SuccessResponseWithMessage(ctx, "Book successfully stored")
}

func (h *Handler) DeleteBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		BadRequest(ctx, err)
		return
	}

	if err = h.Service.DeleteBook(ctx, id); err != nil {
		InternalServerError(ctx, err)
		return
	}

	SuccessResponseWithMessage(ctx, "Book successfully deleted")
}

func (h *Handler) UpdateBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		BadRequest(ctx, err)
		return
	}

	var request structure.UpdateBookRequest
	if err = ctx.ShouldBind(&request); err != nil {
		BadRequest(ctx, err)
		return
	}

	err = h.Service.UpdateBook(ctx, id, structure.Book{
		Title:  request.Title,
		Author: request.Author,
	})

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	SuccessResponseWithMessage(ctx, "Book successfully updated")
}

func (h *Handler) GetBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		BadRequest(ctx, err)
		return
	}

	book, err := h.Service.GetBook(ctx, id)
	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	SuccessResponse(ctx, gin.H{"book": book})
}
