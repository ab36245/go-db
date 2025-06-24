package db

import "go.mongodb.org/mongo-driver/v2/bson"

type Codec[T any] struct {
	Decode func(bson.M) (T, error)
	Encode func(T) (bson.M, error)
}
