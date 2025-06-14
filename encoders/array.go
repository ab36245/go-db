package encoders

import (
	"time"

	"github.com/ab36245/go-model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func NewArrayEncoder(length int) *ArrayEncoder {
	return &ArrayEncoder{
		mongo: make(bson.A, 0, length),
	}
}

type ArrayEncoder struct {
	mongo bson.A
}

func (e *ArrayEncoder) PutArray(length int, f func(model.ArrayEncoder)) {
	e.putValue(encodeArray(length, f))
}

func (e *ArrayEncoder) PutDate(value time.Time) {
	e.putValue(encodeDate(value))
}

func (e *ArrayEncoder) PutInt(value int) {
	e.putValue(encodeInt(value))
}

func (e *ArrayEncoder) PutMap(length int, f func(model.MapEncoder)) {
	e.putValue(encodeMap(length, f))
}

func (e *ArrayEncoder) PutObject(f func(model.ObjectEncoder)) {
	e.putValue(encodeObject(f))
}

func (e *ArrayEncoder) PutRef(value model.Ref) {
	e.putValue(encodeRef(value))
}

func (e *ArrayEncoder) PutString(value string) {
	e.putValue(encodeString(value))
}

func (e *ArrayEncoder) Value() bson.A {
	return e.mongo
}

func (e *ArrayEncoder) putValue(value any) {
	e.mongo = append(e.mongo, value)
}
