package main

import "github.com/onymax/Modern-Business-Soft-Tutor/internal/common/server"

func main() {

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return HandlerFromMux(HttpServer{firebaseDB, trainerClientm usersClient}, router)
	})
}