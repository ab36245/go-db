package decoders

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/ab36245/go-model"
)

func NewObjectDecoder(mongo bson.M) *ObjectDecoder {
	return &ObjectDecoder{
		mongo: mongo,
	}
}

type ObjectDecoder struct {
	mongo bson.M
}

func (d *ObjectDecoder) GetArray(name string) (model.ArrayDecoder, error) {
	return d.reader(name).getArray()
}

func (d *ObjectDecoder) GetBool(name string) (bool, error) {
	return d.reader(name).getBool()
}

func (d *ObjectDecoder) GetFloat(name string) (float64, error) {
	return d.reader(name).getFloat()
}

func (d *ObjectDecoder) GetInt(name string) (int, error) {
	return d.reader(name).getInt()
}

func (d *ObjectDecoder) GetMap(name string) (model.MapDecoder, error) {
	return d.reader(name).getMap()
}

func (d *ObjectDecoder) GetObject(name string) (model.ObjectDecoder, error) {
	return d.reader(name).getObject()
}

func (d *ObjectDecoder) GetRef(name string) (model.Ref, error) {
	return d.reader(name).getRef()
}

func (d *ObjectDecoder) GetString(name string) (string, error) {
	return d.reader(name).getString()
}

func (d *ObjectDecoder) GetTime(name string) (time.Time, error) {
	return d.reader(name).getTime()
}

func (d *ObjectDecoder) reader(name string) reader {
	return func() (any, error) {
		value, ok := d.mongo[name]
		if !ok {
			return nil, fmt.Errorf("field '%s': unknown", name)
		}
		return value, nil
	}
}
