package uuidgql

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

func MarshalUUID(u uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(writer io.Writer) {
		_, _ = io.WriteString(writer, strconv.Quote(u.String()))
	})
}

func UnmarshalUUID(v interface{}) (u uuid.UUID, err error) {
	s, ok := v.(string)
	if !ok {
		return u, fmt.Errorf("INVALID TYPE %T, EXPECT STRING", v)
	}

	return uuid.Parse(s)
}
