package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux" // path用のライブラリ
	"github.com/pulse227/server-recruit-challenge-sample/model"
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

// GET /albums/{id} のハンドラー
func (c *albumController) GetAlbumDetailHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	albumID, err := strconv.Atoi(idString)
	if err != nil {
		err = fmt.Errorf("invalid path param: %w", err)
		errorHandler(w, r, 400, err.Error())
		return
	}

	album, err := c.service.GetAlbumService(r.Context(), model.AlbumID(albumID))
	if err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(album)
}

// POST /albums のハンドラー
func (c *albumController) PostAlbumHandler(w http.ResponseWriter, r *http.Request) {
	var album *model.Album
	if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
		err = fmt.Errorf("invalid body param: %w", err)
		errorHandler(w, r, 400, err.Error())
		return
	}

	if err := c.service.PostAlbumService(r.Context(), album); err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(album)
}

// DELETE /albums/{id} のハンドラー
func (c *albumController) DeleteAlbumHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	albumID, err := strconv.Atoi(idString)
	if err != nil {
		err = fmt.Errorf("invalid path param: %w", err)
		errorHandler(w, r, 400, err.Error())
		return
	}

	if err := c.service.DeleteAlbumService(r.Context(), model.AlbumID(albumID)); err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	w.WriteHeader(204)
}

