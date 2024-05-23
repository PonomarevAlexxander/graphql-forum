package model

import (
	"fmt"
	"io"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type Rfc3339Date time.Time

func MarshalRfc3339Date(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, fmt.Sprintf(`"%s"`, t.Format(time.RFC3339Nano)))
	})
}

func UnmarshalRfc3339Date(v interface{}) (time.Time, error) {
	value, err := graphql.UnmarshalString(v)
	if err != nil {
		return time.Time{}, err
	}

	return time.Parse(time.RFC3339Nano, value)
}
