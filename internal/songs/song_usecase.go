package songs

import (
	"context"

	"com.thebeachmaster/golang-graphql-template/internal/songs/models"
	"github.com/google/uuid"
)

type SongUsecase interface {
	AddSong(ctx context.Context, info *models.SongInfo) (*models.SongInfo, error)
	FetchSong(ctx context.Context, Id uuid.UUID) (*models.SongInfo, error)
	FetchSongs(ctx context.Context, Ids []uuid.UUID) ([]*models.SongInfo, error)
	DeleteSong(ctx context.Context, Id uuid.UUID) error
}
