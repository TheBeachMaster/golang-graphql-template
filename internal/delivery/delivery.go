package delivery

import "com.thebeachmaster/golang-graphql-template/internal/songs"

type Delivery struct {
	Songs interface{ songs.SongDelivery }
}
