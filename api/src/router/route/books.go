package route

import (
	"api/src/controllers"
	"net/http"
)

var bookRouters = []Route{
	{
		URI:          "/books",
		Method:       http.MethodPost,
		FunctionName: controllers.CreateBook,
	},
	{
		URI:          "/books",
		Method:       http.MethodGet,
		FunctionName: controllers.FindAllBooks,
	},
	{
		URI:          "/books/{bookId}",
		Method:       http.MethodGet,
		FunctionName: controllers.FindBook,
	},
	{
		URI:          "/books/{bookId}",
		Method:       http.MethodPut,
		FunctionName: controllers.UpdateBook,
	},
	{
		URI:          "/books/{bookId}",
		Method:       http.MethodDelete,
		FunctionName: controllers.DeleteBook,
	},
}
