package delivery

import (
	"context"

	songs "com.thebeachmaster/golang-graphql-template/internal/songs"
	"com.thebeachmaster/golang-graphql-template/internal/songs/models"
	"github.com/google/uuid"
)

type songDelivery struct {
	usecase songs.SongUsecase
}

func NewSongDelivery(u songs.SongUsecase) songs.SongDelivery {
	return &songDelivery{usecase: u}
}

func (d *songDelivery) AddSong(ctx context.Context, info *models.SongInfo) (*models.SongInfo, error) {
	return d.usecase.AddSong(ctx, info)
}

func (d *songDelivery) FetchSong(ctx context.Context, Id uuid.UUID) (*models.SongInfo, error) {
	return d.usecase.FetchSong(ctx, Id)
}

func (d *songDelivery) FetchSongs(ctx context.Context, Ids []uuid.UUID) ([]*models.SongInfo, error) {
	return d.usecase.FetchSongs(ctx, Ids)
}

func (d *songDelivery) DeleteSong(ctx context.Context, Id uuid.UUID) error {
	return d.usecase.DeleteSong(ctx, Id)
}
