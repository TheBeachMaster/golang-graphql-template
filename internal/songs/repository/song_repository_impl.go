package repository

import (
	"context"

	songs "com.thebeachmaster/golang-graphql-template/internal/songs"
	"com.thebeachmaster/golang-graphql-template/internal/songs/models"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type songRepository struct {
	redis *redis.Client
}

func NewSongRepository(r *redis.Client) songs.SongRepository {
	return &songRepository{redis: r}
}

func (d *songRepository) AddSong(ctx context.Context, info *models.SongInfo) (*models.SongInfo, error) {

}

func (d *songRepository) FetchSong(ctx context.Context, Id uuid.UUID) (*models.SongInfo, error) {}

func (d *songRepository) FetchSongs(ctx context.Context, Ids []uuid.UUID) ([]*models.SongInfo, error) {
}

func (d *songRepository) DeleteSong(ctx context.Context, Id uuid.UUID) error
