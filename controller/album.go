package controller

import (
	"encoding/json"
	// "fmt"
	"net/http"
	// "strconv"

	// "github.com/gorilla/mux"	// path用のライブラリ
	// "github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/service"
)

type albumController struct {
	service service.AlbumService
}

func NewAlbumController(s service.AlbumService) *albumController {
	return &albumController{service: s}
}

// GET /albumss のハンドラー
func (c *albumController) GetAlbumListHandler(w http.ResponseWriter, r *http.Request) {
	albums, err := c.service.GetAlbumListService(r.Context())
	if err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(albums)
}


