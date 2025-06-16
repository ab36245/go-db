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

func (e *ObjectEncoder) PutArray(name string, length int) (model.ArrayEncoder, error) {
	return e.writer(name).putArray(length)
}

func (e *ObjectEncoder) PutDate(name string, value time.Time) error {
	return e.writer(name).putDate(value)
}

func (e *ObjectEncoder) PutInt(name string, value int) error {
	return e.writer(name).putInt(value)
}

func (e *ObjectEncoder) PutMap(name string, length int) (model.MapEncoder, error) {
	return e.writer(name).putMap(length)
}

func (e *ObjectEncoder) PutObject(name string) (model.ObjectEncoder, error) {
	return e.writer(name).putObject()
}

func (e *ObjectEncoder) PutRef(name string, value model.Ref) error {
	return e.writer(name).putRef(value)
}

func (e *ObjectEncoder) PutString(name string, value string) error {
	return e.writer(name).putString(value)
}

func (e *ObjectEncoder) Value() bson.M {
	return e.mongo
}

func (e *ObjectEncoder) writer(name string) writer {
	return func(value any) error {
		e.mongo[name] = value
		return nil
	}
}
