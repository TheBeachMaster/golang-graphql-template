package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"com.thebeachmaster/golang-graphql-template/config"
	songs "com.thebeachmaster/golang-graphql-template/internal/songs"
	"com.thebeachmaster/golang-graphql-template/internal/songs/models"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type songRepository struct {
	redis *redis.Client
	cfg   *config.Config
}

func NewSongRepository(r *redis.Client, c *config.Config) songs.SongRepository {
	return &songRepository{redis: r, cfg: c}
}

func (d *songRepository) AddSong(ctx context.Context, info *models.SongInfo) (*models.SongInfo, error) {
	songId := info.Id.String()
	_key := d.generateKey(d.cfg.Cache.SongPrefix, songId)
	if err := d.redis.HSet(ctx, _key, info).Err(); err != nil {
		return nil, fmt.Errorf("[REDIS] Set Error %s", err.Error())
	}

	return info, nil
}

func (d *songRepository) FetchSong(ctx context.Context, Id uuid.UUID) (*models.SongInfo, error) {
	songId := Id.String()
	_key := d.generateKey(d.cfg.Cache.SongPrefix, songId)
	songResult, err := d.redis.Get(ctx, _key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil // This just shows that there's nothing
	}
	if err != nil {
		return nil, fmt.Errorf("[REDIS] Query error: %s", err.Error())
	}
	song := &models.SongInfo{}
	if err := song.UnMarshalBinary([]byte(songResult)); err != nil {
		return nil, fmt.Errorf("[REDIS] unable to deserialize result due to: %s", err.Error())
	}

	return song, nil
}

func (d *songRepository) FetchSongs(ctx context.Context) ([]*models.SongInfo, error) {
	songs := make([]*models.SongInfo, 0)
	it := d.redis.Scan(ctx, 0, (d.cfg.Cache.SongPrefix + ":*"), 0).Iterator()
	for it.Next(ctx) {
		res := it.Val()
		song := &models.SongInfo{}
		if err := song.UnMarshalBinary([]byte(res)); err != nil {
			return nil, fmt.Errorf("[REDIS] unable to deserialize result due to: %s", err.Error())
		}
		songs = append(songs, song)
	}
	if err := it.Err(); err != nil {
		return nil, fmt.Errorf("[REDIS] Error scanning results")
	}
	return songs, nil
}

func (d *songRepository) DeleteSong(ctx context.Context, Id uuid.UUID) error {
	songId := Id.String()
	_key := d.generateKey(d.cfg.Cache.SongPrefix, songId)
	return d.redis.Del(ctx, _key).Err()
}

func (d *songRepository) generateKey(prefix string, key string) string {
	keyBuilder := strings.Builder{}
	keyBuilder.WriteString(fmt.Sprintf("%s:%s", prefix, key))
	return keyBuilder.String()
}
