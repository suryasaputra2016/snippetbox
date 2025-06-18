package main

import (
	"net/http"

	"github.com/suryasaputra2016/snippetbox/ui"
)

func (app *application) routes() http.Handler {
	dMux := http.NewServeMux()
	dMux.HandleFunc("GET /{$}", app.home)
	dMux.HandleFunc("GET /snippet/view/", app.snippetView)
	dMux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	dMux.HandleFunc("GET /user/signup", app.userSignup)
	dMux.HandleFunc("POST /user/signup", app.userSignupPost)
	dMux.HandleFunc("GET /user/login", app.userLogin)
	dMux.HandleFunc("POST /user/login", app.userLoginPost)

	pdMux := http.NewServeMux()
	pdMux.HandleFunc("GET /snippet/create", app.snippetCreate)
	pdMux.HandleFunc("POST /snippet/create", app.snippetCreatePost)
	pdMux.HandleFunc("POST /user/logout", app.userLogout)

	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.FileServerFS(ui.Files))
	mux.Handle("/", noSurf(app.sessionManager.LoadAndSave(app.authenticate(dMux))))
	dMux.Handle("/", app.requireAuthentication(pdMux))

	mux.HandleFunc("GET /ping", ping)

	return app.recoverPanic(app.logRequest(commonHeaders(mux)))
}
