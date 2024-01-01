package utils

import (
	"github.com/oklog/ulid/v2"
	"math/rand"
	"time"
)

func GenerateULID() ulid.ULID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id, err := ulid.New(ulid.Timestamp(t), entropy)
	if err != nil {
		panic(err)
	}
	return id
}
