package main

import (
	"net/http"

	"os"

	"goweb1/infrastructure"

	"github.com/garyburd/redigo/redis"
	"github.com/go-chi/chi"
)

func main() {
	mux := chi.NewRouter()

	// sql new.
	sqlHandler := infrastructure.NewSQL()
	// cache new.
	cacheHandler := infrastructure.NewCache()
	// logger new.
	loggerHandler := infrastructure.NewLogger()
	// translation new.
	translationHandler := infrastructure.NewTranslation()

	r := &Router{mux: mux, sqlHandler: sqlHandler, cacheHandler: cacheHandler, loggerHandler: loggerHandler, translationHandler: translationHandler}

	r.InitializeRouter()
	r.SetupHandler()

	// after process
	defer closeLogger(r.loggerHandler.Logfile)
	defer closeRedis(r.cacheHandler.Conn)

	http.ListenAndServe(":8080", mux)
}

// after process
func closeLogger(logfile *os.File) {
	// close file.
	if logfile != nil {
		logfile.Close()
	}
}
func closeRedis(conn *redis.Conn) {
	// close redis connection.
	if conn != nil {
		(*conn).Close()
	}
}
