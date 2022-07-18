package registry

import (
	"com.thebeachmaster/golang-graphql-template/internal/songs"
	repository "com.thebeachmaster/golang-graphql-template/internal/songs/repository"
	usecase "com.thebeachmaster/golang-graphql-template/internal/songs/usecase"

	delivery "com.thebeachmaster/golang-graphql-template/internal/songs/usecase"
)

func (r *registry) NewSongDelivery() songs.SongDelivery {
	repo := repository.NewSongRepository(r.redis)
	_usecase := usecase.NewSongUsecase(repo)

	return delivery.NewSongUsecase(_usecase)
}
