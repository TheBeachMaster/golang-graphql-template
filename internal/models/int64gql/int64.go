package int64gql

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalInt64(t int64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.FormatInt(t, 10))
	})
}

// Unmarshal{Typename} is only required if the scalar appears as an input. The raw values have already been decoded
// from json into int/float64/bool/nil/map[string]interface/[]interface
func UnmarshalInt64(v interface{}) (int64, error) {
	if res, ok := v.(json.Number); ok {
		return res.Int64()
	}
	if res, ok := v.(string); ok {
		return json.Number(res).Int64()
	}
	if res, ok := v.(int64); ok {
		return res, nil
	}
	if res, ok := v.(*int64); ok {
		return *res, nil
	}
	return 0, fmt.Errorf("could not convert %v of type %T to Int64", v, v)
}
