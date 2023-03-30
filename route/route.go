package route

import (
	"sesi_8/handler"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, server handler.HttpServer) {
	// api := r.Group("/employee") // prefix
	// {
	// 	api.POST("", server.CreateEmployee)       // /employees
	// 	api.GET("/:id", server.GetEmployeeByID)   // /employee/:id
	// 	api.PUT("/:id", server.UpdateEmployee)    // /employee/:id
	// 	api.DELETE("/:id", server.DeleteEmployee) // /employee/:id
	// }
	book := r.Group("/book")
	{
		book.POST("", server.CreateBook)
		book.GET("", server.GetAllBook)
		book.GET("/:id", server.GetBookById)
		book.PUT("/:id", server.UpdateBook)
		book.DELETE("/:id", server.DeleteBook)
	}
}
