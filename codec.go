package db

import "go.mongodb.org/mongo-driver/v2/bson"

type A = bson.A
type D = bson.D
type E = bson.E
type M = bson.M

type Codec[T any] struct {
	Decode func(M) (T, error)
	Encode func(T) (M, error)
}
