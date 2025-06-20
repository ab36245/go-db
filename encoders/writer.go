package encoders

import (
	"fmt"
	"time"

	"github.com/ab36245/go-model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type writer func(any) error

func (w writer) putArray(length int) (model.ArrayEncoder, error) {
	e := NewArrayEncoder(length)
	if err := w(e.mongo); err != nil {
		return nil, err
	}
	return e, nil
}

func (w writer) putBool(value bool) error {
	return w(value)
}

func (w writer) putBytes(value []byte) error {
	return w(value)
}

func (w writer) putFloat(value float64) error {
	return w(value)
}

func (w writer) putInt(value int) error {
	return w(value)
}

func (w writer) putMap(length int) (model.MapEncoder, error) {
	e := NewMapEncoder(length)
	if err := w(e.mongo); err != nil {
		return nil, err
	}
	return e, nil
}

func (w writer) putObject() (model.ObjectEncoder, error) {
	e := NewObjectEncoder()
	if err := w(e.mongo); err != nil {
		return nil, err
	}
	return e, nil
}

func (w writer) putRef(value model.Ref) error {
	if value == "" {
		return nil
	}
	id, err := bson.ObjectIDFromHex(string(value))
	if err != nil {
		return fmt.Errorf("can't encode %s as ObjectID: %w", value, err)
	}
	return w(id)
}

func (w writer) putString(value string) error {
	return w(value)
}

func (w writer) putTime(value time.Time) error {
	return w(bson.NewDateTimeFromTime(value))
}
