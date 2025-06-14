package db

import (
	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/ab36245/go-db/decoders"
)

func newDecoder(mongo bson.M) *decoder {
	return &decoder{
		decoders.NewObjectDecoder(mongo),
	}
}

type decoder struct {
	*decoders.ObjectDecoder
}
