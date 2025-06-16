package encoders

import (
	"fmt"
	"time"

	"github.com/ab36245/go-model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func NewArrayEncoder(length int) *ArrayEncoder {
	return &ArrayEncoder{
		index: 0,
		mongo: make(bson.A, length, length),
	}
}

type ArrayEncoder struct {
	index int
	mongo bson.A
}

func (e *ArrayEncoder) PutArray(length int) (model.ArrayEncoder, error) {
	return e.writer().putArray(length)
}

func (e *ArrayEncoder) PutDate(value time.Time) error {
	return e.writer().putDate(value)
}

func (e *ArrayEncoder) PutInt(value int) error {
	return e.writer().putInt(value)
}

func (e *ArrayEncoder) PutMap(length int) (model.MapEncoder, error) {
	return e.writer().putMap(length)
}

func (e *ArrayEncoder) PutObject() (model.ObjectEncoder, error) {
	return e.writer().putObject()
}

func (e *ArrayEncoder) PutRef(value model.Ref) error {
	return e.writer().putRef(value)
}

func (e *ArrayEncoder) PutString(value string) error {
	return e.writer().putString(value)
}

func (e *ArrayEncoder) writer() writer {
	return func(value any) error {
		max := len(e.mongo)
		if e.index >= max {
			return fmt.Errorf("trying to write more than max (%d) values", max)
		}
		e.mongo[e.index] = value
		e.index++
		return nil
	}
}
