package usecase

import (
	"context"

	songs "com.thebeachmaster/golang-graphql-template/internal/songs"
	"com.thebeachmaster/golang-graphql-template/internal/songs/models"
	"github.com/google/uuid"
)

type songUsecase struct {
	repo songs.SongRepository
}

func NewSongUsecase(r songs.SongRepository) songs.SongUsecase {
	return &songUsecase{repo: r}
}

func (d *songUsecase) AddSong(ctx context.Context, info *models.SongInfo) (*models.SongInfo, error) {
	return d.repo.AddSong(ctx, info)
}

func (d *songUsecase) FetchSong(ctx context.Context, Id uuid.UUID) (*models.SongInfo, error) {
	return d.repo.FetchSong(ctx, Id)
}

func (d *songUsecase) FetchSongs(ctx context.Context, Ids []uuid.UUID) ([]*models.SongInfo, error) {
	return d.repo.FetchSongs(ctx, Ids)
}

func (d *songUsecase) DeleteSong(ctx context.Context, Id uuid.UUID) error {
	return d.repo.DeleteSong(ctx, Id)
}
