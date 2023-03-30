package handler

import (
	"sesi_8/helper"
	"sesi_8/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h HttpServer) CreateBook(c *gin.Context) {
	in := model.Book{}

	err := c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	err = in.Validation()
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	//call service
	res, err := h.app.CreateBook(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) GetAllBook(c *gin.Context) {
	res, err := h.app.GetAllBook()
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}
	helper.Ok(c, res)
}

func (h HttpServer) GetBookById(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	res, err := h.app.GetBookById(bookID)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}
	helper.Ok(c, res)
}

func (h HttpServer) UpdateBook(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	in := model.Book{}

	err := c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}
	err = in.Validation()
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}
	res, err := h.app.UpdateBook(bookID, in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}
	helper.OkWithMessage(c, res)
}

func (h HttpServer) DeleteBook(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	res, err := h.app.DeleteBook(bookID)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}
	helper.OkWithMessage(c, res)
}
