package encoders

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/ab36245/go-model"
)

func NewMapEncoder(length int) *MapEncoder {
	return &MapEncoder{
		index:   0,
		keyNext: true,
		mongo:   make(bson.D, length),
	}
}

type MapEncoder struct {
	index   int
	keyNext bool
	mongo   bson.D
}

func (e *MapEncoder) PutArray(length int) (model.ArrayEncoder, error) {
	return e.writer().putArray(length)
}

func (e *MapEncoder) PutDate(value time.Time) error {
	return e.writer().putDate(value)
}

func (e *MapEncoder) PutInt(value int) error {
	return e.writer().putInt(value)
}

func (e *MapEncoder) PutMap(length int) (model.MapEncoder, error) {
	return e.writer().putMap(length)
}

func (e *MapEncoder) PutObject() (model.ObjectEncoder, error) {
	return e.writer().putObject()
}

func (e *MapEncoder) PutRef(value model.Ref) error {
	return e.writer().putRef(value)
}

func (e *MapEncoder) PutString(value string) error {
	return e.writer().putString(value)
}

func (e *MapEncoder) writer() writer {
	return func(value any) error {
		max := len(e.mongo)
		if e.index >= max {
			return fmt.Errorf("trying to write more than max (%d) values", max)
		}
		if e.keyNext {
			if value, ok := value.(string); ok {
				e.mongo[e.index].Key = value
				e.keyNext = false
			} else {
				return fmt.Errorf("key must be a string, got %T", value)
			}
		} else {
			e.mongo[e.index].Value = value
			e.index++
			e.keyNext = true
		}
		return nil
	}
}
