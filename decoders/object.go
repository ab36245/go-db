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
	return decodeArray(d.getValue(name))
}

func (d *ObjectDecoder) GetDate(name string) (time.Time, error) {
	return decodeDate(d.getValue(name))
}

func (d *ObjectDecoder) GetInt(name string) (int, error) {
	return decodeInt(d.getValue(name))
}

func (d *ObjectDecoder) GetMap(name string) (model.MapDecoder, error) {
	return decodeMap(d.getValue(name))
}

func (d *ObjectDecoder) GetObject(name string) (model.ObjectDecoder, error) {
	return decodeObject(d.getValue(name))
}

func (d *ObjectDecoder) GetRef(name string) (model.Ref, error) {
	return decodeRef(d.getValue(name))
}

func (d *ObjectDecoder) GetString(name string) (string, error) {
	return decodeString(d.getValue(name))
}

func (d *ObjectDecoder) getValue(name string) (any, error) {
	val, ok := d.mongo[name]
	if !ok {
		return nil, fmt.Errorf("field '%s': unknown", name)
	}
	return val, nil
}
