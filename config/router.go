package config

import (
	"database/sql"

	userDeliver "goweb1/user/delivery/http"
	userRepo "goweb1/user/repository/psql"
	userCase "goweb1/user/usecase"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Router func
func Router(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))

	addUserRoutes(r, db)
	return r
}

func addUserRoutes(r *chi.Mux, db *sql.DB) {
	repo := userRepo.NewUserRepository(db)
	uc := userCase.NewUserUsecase(repo)
	userDeliver.NewUserController(r, uc)
}
