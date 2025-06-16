package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.Handle("GET /{$}", app.sessionManager.LoadAndSave(http.HandlerFunc(app.home)))
	dynamicMux := http.NewServeMux()
	dynamicMux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	dynamicMux.HandleFunc("GET /snippet/create", app.snippetCreate)
	dynamicMux.HandleFunc("POST /snippet/create", app.snippetCreatePost)
	mux.Handle("/snippet/", app.sessionManager.LoadAndSave(dynamicMux))

	mux.HandleFunc("GET /user/signup", app.userSignup)
	mux.HandleFunc("POST /user/signup", app.userSignupPost)
	mux.HandleFunc("GET /user/login", app.userLogin)
	mux.HandleFunc("POST /user/login", app.userLoginPost)
	mux.HandleFunc("POST /user/logout", app.userLogout)

	return app.recoverPanic(app.logRequest(commonHeaders(mux)))
}
