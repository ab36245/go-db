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

func (e *ArrayEncoder) PutArray(length int, handler model.ArrayHandler) error {
	return e.writer().putArray(length, handler)
}

func (e *ArrayEncoder) PutDate(value time.Time) error {
	return e.writer().putDate(value)
}

func (e *ArrayEncoder) PutInt(value int) error {
	return e.writer().putInt(value)
}

func (e *ArrayEncoder) PutMap(length int, handler model.MapHandler) error {
	return e.writer().putMap(length, handler)
}

func (e *ArrayEncoder) PutObject(handler model.ObjectHandler) error {
	return e.writer().putObject(handler)
}

func (e *ArrayEncoder) PutRef(value model.Ref) error {
	return e.writer().putRef(value)
}

func (e *ArrayEncoder) PutString(value string) error {
	return e.writer().putString(value)
}

func (e *ArrayEncoder) Value() bson.A {
	return e.mongo
}

func (e *ArrayEncoder) writer() writer {
	return func(value any) error {
		e.mongo = append(e.mongo, value)
		return nil
	}
}
