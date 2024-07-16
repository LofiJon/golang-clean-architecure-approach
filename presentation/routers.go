package presentation

//func NewRouter(
//	createTaskController *task.CreateTaskController,
//	getByIdTaskController *task.GetByIdTaskController,
//) *mux.Router {
//	router := mux.NewRouter().StrictSlash(true)
//
//	taskRouter := NewRouter(createTaskController)
//	taskGetByIdRouterRouter := NewRouter(getByIdTaskController)
//	router.PathPrefix("/").Handler(taskRouter)
//	router.PathPrefix("/{id}").Handler(taskGetByIdRouterRouter)
//	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
//
//	return router
//}
