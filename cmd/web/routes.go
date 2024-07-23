package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (app *application) routes(staticDir string) http.Handler {
	// create a middleware chain
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable, noSurf, app.authenticate)

	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/snippet/create", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.createSnippet))
	mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.showSnippet))

	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.logoutUser))

	mux.Get("/health", dynamicMiddleware.ThenFunc(app.healthCheckHandler))
	mux.Get("/swagger/", httpSwagger.WrapHandler)

	fileserver := http.FileServer(http.Dir(staticDir))
	mux.Get("/static/", http.StripPrefix("/static", fileserver))

	return standardMiddleware.Then(mux)

	/*
		/users
			GET /users: получить список всех пользователей.
			POST /users: создать нового пользователя.
			GET /users/{id}: получить данные конкретного пользователя.
			PUT /users/{id}: обновить данные конкретного пользователя.
			DELETE /users/{id}: удалить конкретного пользователя.
			GET /users/{id}/tasks: получить список задач конкретного пользователя.
			GET /users/search?name={name}: найти пользователей по имени.
			GET /users/search?email={email}: найти пользователей по электронной почте.
		/tasks
			GET /tasks: получить список всех задач.
			POST /tasks: создать новую задачу.
			GET /tasks/{id}: получить данные конкретной задачи.
			PUT /tasks/{id}: обновить данные конкретной задачи.
			DELETE /tasks/{id}: удалить конкретную задачу.
			GET /tasks/search?title={title}: найти задачи по названию.
			GET /tasks/search?status={status}: найти задачи по состоянию.
			GET /tasks/search?priority={priority}: найти задачи по приоритету.
			GET /tasks/search?assignee={userId}: найти задачи по идентификатору ответственного.
			GET /tasks/search?project={projectId}: найти задачи по идентификатору проекта.
		/projects
			GET /projects: получить список всех проектов.
			POST /projects: создать новый проект.
			GET /projects/{id}: получить данные конкретного проекта.
			PUT /projects/{id}: обновить данные конкретного проекта.
			DELETE /projects/{id}: удалить конкретный проект.
			GET /projects/{id}/tasks: получить список задач в проекте.
			GET /projects/search?title={title}: найти проекты по названию.
			GET /projects/search?manager={userId}: найти проекты по идентификатору менеджера.

	*/
}
