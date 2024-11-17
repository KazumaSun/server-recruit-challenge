package mysqldb

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/repository"
)

func NewAlbumRepository(db *sql.DB) *albumRepository {
	return &albumRepository{
		db: db,
	}
}

type albumRepository struct {
	db *sql.DB
}

var _ repository.AlbumRepository = (*albumRepository)(nil)

func (r *albumRepository) GetAll(ctx context.Context) ([]*model.AlbumDetail, error) {
	albums := []*model.AlbumDetail{}
	var singerJSON string
	query := "SELECT albums.id AS id, albums.title AS title, JSON_OBJECT( 'id', singers.id, 'name', singers.name) AS singer FROM albums INNER JOIN singers ON albums.singer_id = singers.id ORDER BY id ASC"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		album := &model.AlbumDetail{}
		if err := rows.Scan(&album.ID, &album.Title, &singerJSON); err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(singerJSON), &album.Singer)
		if err != nil {
			log.Fatal(err)
		}
		if album.ID != 0 {
			albums = append(albums, album)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return albums, nil
}

func (r *albumRepository) Get(ctx context.Context, id model.AlbumID) (*model.AlbumDetail, error) {
	album := &model.AlbumDetail{}
	var singerJSON string
	query := "SELECT albums.id AS id, albums.title AS title, JSON_OBJECT( 'id', singers.id, 'name', singers.name) AS singer FROM albums INNER JOIN singers ON albums.singer_id = singers.id WHERE albums.id = ?"
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&album.ID, &album.Title, &singerJSON); err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(singerJSON), &album.Singer)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	if album.ID == 0 {
		return nil, model.ErrNotFound
	}
	return album, nil
}

func (r *albumRepository) Add(ctx context.Context, album *model.Album) error {
	query := "INSERT INTO albums (id, title, singer_id) VALUES (?, ?, ?)"
	if _, err := r.db.ExecContext(ctx, query, album.ID, album.Title, album.SingerID); err != nil {
		return err
	}
	return nil
}

func (r *albumRepository) Delete(ctx context.Context, id model.AlbumID) error {
	query := "DELETE FROM albums WHERE id = ?"
	if _, err := r.db.ExecContext(ctx, query, id); err != nil {
		return err
	}
	return nil
}
