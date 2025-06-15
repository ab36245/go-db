package encoders

import (
	"fmt"
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

func (e *MapEncoder) PutArray(length int, handler model.ArrayHandler) error {
	return e.writer().putArray(length, handler)
}

func (e *MapEncoder) PutDate(value time.Time) error {
	return e.writer().putDate(value)
}

func (e *MapEncoder) PutInt(value int) error {
	return e.writer().putInt(value)
}

func (e *MapEncoder) PutMap(length int, handler model.MapHandler) error {
	return e.writer().putMap(length, handler)
}

func (e *MapEncoder) PutObject(handler model.ObjectHandler) error {
	return e.writer().putObject(handler)
}

func (e *MapEncoder) PutRef(value model.Ref) error {
	return e.writer().putRef(value)
}

func (e *MapEncoder) PutString(value string) error {
	return e.writer().putString(value)
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

func (e *MapEncoder) writer() writer {
	switch {
	case len(e.keys) == len(e.values):
		// expecting a key
		return func(key any) error {
			switch key := key.(type) {
			case string:
				e.keys = append(e.keys, key)
				return nil
			default:
				return fmt.Errorf("key must be a string, got %T", key)
			}
		}

	case len(e.keys) == len(e.values)+1:
		// expecting a value
		return func(value any) error {
			e.values = append(e.values, value)
			return nil
		}

	default:
		panic("weird map state!")
	}
}
