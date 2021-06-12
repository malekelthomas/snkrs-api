package rest

import (
	"log"
	"math/rand"
	"time"
)

func paginate(n int64) (limit int64, offset int64) {
	limit = 40
	if n <= 1 {
		offset = 0
	} else {
		offset = n * 40
	}

	return limit, offset
}

func paginateRandom(count int64) (limit int64, offset int64) {

	limit = 40
	rand.Seed(time.Now().UnixNano())
	offset = rand.Int63n(count)
	if count-offset < limit { //set offset to always return 40 items
		offset -= (limit - (count - offset))
	}
	log.Println(offset)
	return limit, offset
}
