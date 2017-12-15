package config

import (
	"database/sql"

	"goweb1/services/cache"
	userRepo "goweb1/user/repository"
	userCase "goweb1/user/usecase"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Router func
func Router(db *sql.DB, c cache.Cache) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	addUserRoutes(r, db, c)
	addProductRoutes(r, db, c)
}

func addUserRoutes(r *chi.Mux, db *sql.DB, c cache.Cache) {
	repo := userRepo.NewUserRepository(db)
	uc := userCase.NewUserUsecase(repo)
	userDeliver.NewUserController(r, uc, c)
}

func addProductRoutes(r *chi.Mux, db *sql.DB, c cache.Cache) {
	repo := productRepo.NewUserRepository(db)
	uc := productCase.NewUserUsecase(repo)
	productDeliver.NewProductController(r, uc, c)
}
