package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"com.thebeachmaster/golang-graphql-template/graphql/models"
	"github.com/google/uuid"
)

// AddSong is the resolver for the addSong field.
func (r *mutationResolver) AddSong(ctx context.Context, input models.SongInfo) (*models.Song, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteSong is the resolver for the deleteSong field.
func (r *mutationResolver) DeleteSong(ctx context.Context, id uuid.UUID) (models.DeleteResult, error) {
	panic(fmt.Errorf("not implemented"))
}

// FetchSongByID is the resolver for the fetchSongById field.
func (r *queryResolver) FetchSongByID(ctx context.Context, id uuid.UUID) (*models.Song, error) {
	panic(fmt.Errorf("not implemented"))
}

// FetchSongs is the resolver for the fetchSongs field.
func (r *queryResolver) FetchSongs(ctx context.Context, ids []uuid.UUID) ([]*models.Song, error) {
	panic(fmt.Errorf("not implemented"))
}
