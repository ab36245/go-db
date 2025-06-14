package encoders

import (
	"time"

	"github.com/ab36245/go-model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func NewMapEncoder(length int) *MapEncoder {
	return &MapEncoder{
		keys:   make([]string, 0, length),
		values: make([]any, 0, length),
	}
}

type MapEncoder struct {
	keys   []string
	values []any
}

func (e *MapEncoder) PutArray(length int, f func(model.ArrayEncoder)) {
	e.putValue(encodeArray(length, f))
}

func (e *MapEncoder) PutDate(value time.Time) {
	e.putValue(encodeDate(value))
}

func (e *MapEncoder) PutInt(value int) {
	e.putValue(encodeInt(value))
}

func (e *MapEncoder) PutKey(value string) {
	e.putKey(encodeString(value))
}

func (e *MapEncoder) PutMap(length int, f func(model.MapEncoder)) {
	e.putValue(encodeMap(length, f))
}

func (e *MapEncoder) PutObject(f func(model.ObjectEncoder)) {
	e.putValue(encodeObject(f))
}

func (e *MapEncoder) PutRef(value model.Ref) {
	e.putValue(encodeRef(value))
}

func (e *MapEncoder) PutString(value string) {
	e.putValue(encodeString(value))
}

func (e *MapEncoder) Value() bson.M {
	value := make(bson.M, len(e.keys))
	for i := range e.keys {
		k := e.keys[i]
		v := e.values[i]
		if v != nil {
			value[k] = v
		}
	}
	return value
}

func (e *MapEncoder) putKey(value string) {
	e.keys = append(e.keys, value)
}

func (e *MapEncoder) putValue(value any) {
	for len(e.values) < len(e.keys)-1 {
		e.values = append(e.values, nil)
	}
	e.values = append(e.values, value)
}
