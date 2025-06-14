package encoders

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/ab36245/go-model"
)

func NewObjectEncoder() *ObjectEncoder {
	return &ObjectEncoder{
		mongo: make(bson.M),
	}
}

type ObjectEncoder struct {
	mongo bson.M
}

func (e *ObjectEncoder) PutArray(name string, length int, f func(model.ArrayEncoder)) {
	e.putValue(name, encodeArray(length, f))
}

func (e *ObjectEncoder) PutDate(name string, value time.Time) {
	e.putValue(name, encodeDate(value))
}

func (e *ObjectEncoder) PutInt(name string, value int) {
	e.putValue(name, encodeInt(value))
}

func (e *ObjectEncoder) PutMap(name string, length int, f func(model.MapEncoder)) {
	e.putValue(name, encodeMap(length, f))
}

func (e *ObjectEncoder) PutObject(name string, f func(model.ObjectEncoder)) {
	e.putValue(name, encodeObject(f))
}

func (e *ObjectEncoder) PutRef(name string, value model.Ref) {
	e.putValue(name, encodeRef(value))
}

func (e *ObjectEncoder) PutString(name string, value string) {
	e.putValue(name, encodeString(value))
}

func (e *ObjectEncoder) Value() bson.M {
	return e.mongo
}

func (e *ObjectEncoder) putValue(name string, value any) {
	e.mongo[name] = value
}
