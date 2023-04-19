package handler

// // delivery, controller

// import (
// 	"sesi_8/helper"
// 	"sesi_8/model"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// func (h HttpServer) CreateEmployee(c *gin.Context) {
// 	in := model.Employee{}

// 	err := c.BindJSON(&in)
// 	if err != nil {
// 		helper.BadRequest(c, err.Error())
// 		return
// 	}

// 	err = in.Validation()
// 	if err != nil {
// 		helper.BadRequest(c, err.Error())
// 		return
// 	}

// 	// call service
// 	res, err := h.app.CreateEmployee(in)
// 	if err != nil {
// 		helper.InternalServerError(c, err.Error())
// 		return
// 	}

// 	helper.Ok(c, res)
// }

// func (h HttpServer) GetEmployeeByID(c *gin.Context) {
// 	id := c.Param("id")

// 	idInt, err := strconv.Atoi(id)
// 	if err != nil {
// 		helper.BadRequest(c, err.Error())
// 		return
// 	}

// 	// call service
// 	res, err := h.app.GetEmployeeByID(int64(idInt))
// 	if err != nil {
// 		if err.Error() == helper.ErrNotFound {
// 			helper.NotFound(c, err.Error())
// 			return
// 		}
// 		helper.InternalServerError(c, err.Error())
// 		return
// 	}

// 	helper.Ok(c, res)
// }

// func (h HttpServer) UpdateEmployee(c *gin.Context) {
// 	id := c.Param("id")

// 	idInt, err := strconv.Atoi(id)
// 	if err != nil {
// 		helper.BadRequest(c, err.Error())
// 		return
// 	}

// 	in := model.Employee{}

// 	err = c.BindJSON(&in)
// 	if err != nil {
// 		helper.BadRequest(c, err.Error())
// 		return
// 	}

// 	in.ID = idInt
// 	// call service
// 	res, err := h.app.UpdateEmployee(in)
// 	if err != nil {
// 		helper.InternalServerError(c, err.Error())
// 		return
// 	}

// 	helper.Ok(c, res)
// }

// func (h HttpServer) DeleteEmployee(c *gin.Context) {
// 	id := c.Param("id")

// 	idInt, err := strconv.Atoi(id)
// 	if err != nil {
// 		helper.BadRequest(c, err.Error())
// 		return
// 	}

// 	// call service
// 	err = h.app.DeleteEmployee(int64(idInt))
// 	if err != nil {
// 		helper.InternalServerError(c, err.Error())
// 		return
// 	}

// 	helper.OkWithMessage(c, "Successfully deleted employee")
// }

// // func (h HttpServer) GetEmployeeByID(c *gin.Context) {
// // 	id := c.Param("id")

// // 	// call service
// // 	res, err := h.app.GetEmployeeByID(id)
// // 	if err != nil {
// // 		helper.InternalServerError(c, err.Error())
// // 		return
// // 	}

// // 	helper.Ok(c, res)
// // }
