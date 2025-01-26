package helper

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomValueFromRange(max, min int) string {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	id := rng.Intn(max-min) + min

	return strconv.Itoa(id)
}
