package models

import (
	"encoding/json"

	"github.com/google/uuid"
)

type SongInfo struct {
	Id           uuid.UUID
	Title        string
	Artist       []string
	AlbumArt     string // FIXME: URL.Url
	SampleLyrics string
}

func (s *SongInfo) MarshalBinary() ([]byte, error) {
	bytes, err := json.Marshal(s)
	return bytes, err
}
func (s *SongInfo) UnMarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return nil
}
