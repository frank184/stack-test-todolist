package router

import (
	"stack-test-todolist/controllers/tasks"

	"github.com/julienschmidt/httprouter"
)

func New() *httprouter.Router {
	router := httprouter.New()
	router.GET("/tasks", tasks.Index)
	router.POST("/tasks", tasks.Create)
	return router
}
