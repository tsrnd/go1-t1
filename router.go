package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"fr-circle-api/article"
	"fr-circle-api/gist"
	"fr-circle-api/infrastructure"
	"fr-circle-api/shared/handler"
	mMiddleware "fr-circle-api/shared/middleware"
)

// Router is application struct hold mux and db connection
type Router struct {
	mux                *chi.Mux
	sqlHandler         *infrastructure.SQL
	cacheHandler       *infrastructure.Cache
	loggerHandler      *infrastructure.Logger
	translationHandler *infrastructure.Translation
}

// InitializeRouter initializes mux and middleware
func (r *Router) InitializeRouter() {
	r.mux.Use(middleware.RequestID)
	r.mux.Use(middleware.RealIP)

	// Custom middleware(Translation)
	r.mux.Use(r.translationHandler.Middleware.Middleware)
	// Custom middleware(Logger)
	r.mux.Use(mMiddleware.Logger(r.loggerHandler))
	// Custom middleware(Header)
	r.mux.Use(mMiddleware.Header(r.loggerHandler))
}

// SetupHandler set database and redis and usecase.
func (r *Router) SetupHandler() {
	// error handler set.
	eh := handler.NewHTTPErrorHandler(r.loggerHandler.Log)
	r.mux.NotFound(eh.StatusNotFound)
	r.mux.MethodNotAllowed(eh.StatusMethodNotAllowed)

	// base set.
	bh := handler.NewBaseHTTPHandler(r.loggerHandler.Log)

	// article set.
	ah := article.NewHTTPHandler(bh, r.sqlHandler, r.cacheHandler)
	r.mux.Get("/article", ah.Get)
	r.mux.Get("/article/{id}", ah.GetID)
	r.mux.Post("/article/add", ah.PostAdd)
	r.mux.Delete("/article/{id}", ah.DeleteID)
	r.mux.Get("/article/count", ah.GetCount)
	r.mux.Post("/article/count", ah.PostCount)
	r.mux.Post("/article/visenze/discoversearch", ah.PostVisenzeDiscoversearch)

	// gist set.
	gh := gist.NewGistHTTPHandler(bh)
	r.mux.Get("/gist", gh.ListGists)
}
