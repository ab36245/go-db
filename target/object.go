package target

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/ab36245/go-codec"
)

func Object() *objectTarget {
	return &objectTarget{
		mongo: make(bson.M),
	}
}

type objectTarget struct {
	mongo bson.M
}

func (t *objectTarget) PutArray(name string, length int, f func(codec.ArrayTarget)) {
	array := Array(length)
	f(array)
	t.putValue(name, array.mongo)
}

func (t *objectTarget) PutDate(name string, value time.Time) {
	t.putValue(name, encodeDate(value))
}

func (t *objectTarget) PutInt(name string, value int) {
	t.putValue(name, encodeInt(value))
}

func (t *objectTarget) PutObject(name string, f func(codec.ObjectTarget)) {
	object := Object()
	f(object)
	t.putValue(name, object.mongo)
}

func (t *objectTarget) PutString(name string, value string) {
	t.putValue(name, encodeString(value))
}

func (t *objectTarget) Value() bson.M {
	return t.mongo
}

func (t *objectTarget) putValue(name string, value any) {
	t.mongo[name] = value
}
