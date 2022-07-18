// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"

	"github.com/google/uuid"
)

type Song struct {
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"Title"`
	Artist       []string  `json:"Artist"`
	AlbumArt     []string  `json:"AlbumArt"`
	SampleLyrics *string   `json:"SampleLyrics"`
}

type SongInfo struct {
	Title        string   `json:"Title"`
	Artist       []string `json:"Artist"`
	AlbumArt     []string `json:"AlbumArt"`
	SampleLyrics *string  `json:"SampleLyrics"`
}

type DeleteResult string

const (
	DeleteResultNotFound DeleteResult = "NOT_FOUND"
	DeleteResultDone     DeleteResult = "DONE"
	DeleteResultErr      DeleteResult = "ERR"
)

var AllDeleteResult = []DeleteResult{
	DeleteResultNotFound,
	DeleteResultDone,
	DeleteResultErr,
}

func (e DeleteResult) IsValid() bool {
	switch e {
	case DeleteResultNotFound, DeleteResultDone, DeleteResultErr:
		return true
	}
	return false
}

func (e DeleteResult) String() string {
	return string(e)
}

func (e *DeleteResult) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DeleteResult(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DELETE_RESULT", str)
	}
	return nil
}

func (e DeleteResult) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
