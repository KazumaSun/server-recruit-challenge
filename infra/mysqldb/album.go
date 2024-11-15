package mysqldb

import (
	"context"
	"database/sql"

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

func (r *albumRepository) GetAll(ctx context.Context) ([]*model.Album, error) {
	albums := []*model.Album{}
	query := "SELECT id, title, singer_id FROM albums ORDER BY id ASC"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		album := &model.Album{}
		if err := rows.Scan(&album.ID, &album.Title, &album.SingerID); err != nil {
			return nil, err
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