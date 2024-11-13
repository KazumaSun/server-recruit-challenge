package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pulse227/server-recruit-challenge-sample/api/middleware"
	"github.com/pulse227/server-recruit-challenge-sample/controller"
	"github.com/pulse227/server-recruit-challenge-sample/infra/mysqldb"
	"github.com/pulse227/server-recruit-challenge-sample/service"
)

func NewRouter(
	dbUser, dbPass, dbHost, dbName string,
) (http.Handler, error) {
	dbClient, err := mysqldb.Initialize(dbUser, dbPass, dbHost, dbName)
	if err != nil {
		return nil, err
	}
	// 接続確認
	if err := dbClient.Ping(); err != nil {
		return nil, err
	}

	singerRepo := mysqldb.NewSingerRepository(dbClient)
	singerService := service.NewSingerService(singerRepo)
	singerController := controller.NewSingerController(singerService)

	mux := mux.NewRouter()
	mux.HandleFunc("/singers", singerController.GetSingerListHandler).Methods("GET")
	mux.HandleFunc("/singers/{id}", singerController.GetSingerDetailHandler).Methods("GET")
	mux.HandleFunc("/singers", singerController.PostSingerHandler).Methods("POST")
	mux.HandleFunc("/singers/{id}", singerController.DeleteSingerHandler).Methods("DELETE")

	wrappedMux := middleware.LoggingMiddleware(mux)

	return wrappedMux, nil
}
