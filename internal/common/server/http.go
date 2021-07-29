package server

import (
	"net/http"
	"os"
	"logs"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

func RunHTTPServer(createHandler func(router chi.Router) http.Handler) {
	apiRouter := chi.NewRouter()
	setMiddlwares(apiRouter)

	rootRouter := chi.NewRouter()

	rootRouter.Mount("/api", createHandler(apiRouter))

	logrus.Info("Starting HTTP server")
	http.ListenAndServe(":"+os.Getenv("PORT"), rootRouter)
}

func setMiddlwares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(logs.NewStructuredLogger(logrus.StandardLogger()))
	router.Use(middleware.Recoverer)

	addCorsMiddleware(router)
	addAuthMiddleware(router)

	router.Use(
		middleware.SetHeader("X-Content_Type_Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	router.Use(middleware.NoCache)
}