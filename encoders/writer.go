package encoders

import (
	"fmt"
	"time"

	"github.com/ab36245/go-model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// type writer struct {
// 	put func(any) error
// }

type writer func(any) error

func (w writer) putArray(length int, handler model.ArrayHandler) error {
	encoder := NewArrayEncoder(length)
	err := handler(encoder)
	if err != nil {
		return err
	}
	return w(encoder.mongo)
}

func (w writer) putDate(value time.Time) error {
	return w(bson.NewDateTimeFromTime(value))
}

func (w writer) putInt(value int) error {
	return w(value)
}

func (w writer) putMap(length int, handler model.MapHandler) error {
	encoder := NewMapEncoder(length)
	err := handler(encoder)
	if err != nil {
		return err
	}
	return w(encoder.Value())
}

func (w writer) putObject(handler model.ObjectHandler) error {
	encoder := NewObjectEncoder()
	err := handler(encoder)
	if err != nil {
		return err
	}
	return w(encoder.Value())
}

func (w writer) putRef(value model.Ref) error {
	id, err := bson.ObjectIDFromHex(string(value))
	if err != nil {
		return fmt.Errorf("can't encode %s as ObjectID: %w", value, err)
	}
	return w(id)
}

func (w writer) putString(value string) error {
	return w(value)
}
