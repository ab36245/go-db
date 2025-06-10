package target

import (
	"time"

	"github.com/ab36245/go-codec"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Array(length int) *arrayTarget {
	return &arrayTarget{
		mongo: make(bson.A, 0, length),
	}
}

type arrayTarget struct {
	mongo bson.A
}

func (t *arrayTarget) PutArray(length int, f func(codec.ArrayTarget)) {
	array := Array(length)
	f(array)
	t.putValue(array.mongo)
}

func (t *arrayTarget) PutDate(value time.Time) {
	t.putValue(encodeDate(value))
}

func (t *arrayTarget) PutInt(value int) {
	t.putValue(encodeInt(value))
}

func (t *arrayTarget) PutObject(f func(codec.ObjectTarget)) {
	object := Object()
	f(object)
	t.putValue(object.mongo)
}

func (t *arrayTarget) PutString(value string) {
	t.putValue(encodeString(value))
}

func (t *arrayTarget) Value() bson.A {
	return t.mongo
}

func (t *arrayTarget) putValue(value any) {
	t.mongo = append(t.mongo, value)
}
