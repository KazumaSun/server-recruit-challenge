package repository

import (
	"context"

	"github.com/pulse227/server-recruit-challenge-sample/model"
)

type AlbumRepository interface {
	GetAll(ctx context.Context) ([]*model.AlbumDetail, error)
	Get(ctx context.Context, id model.AlbumID) (*model.AlbumDetail, error)
	Add(ctx context.Context, singer *model.Album) error
	Delete(ctx context.Context, id model.AlbumID) error
}
