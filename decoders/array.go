package decoders

import (
	"fmt"
	"time"

	"github.com/ab36245/go-model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func NewArrayDecoder(mongo bson.A) *ArrayDecoder {
	return &ArrayDecoder{
		mongo: mongo,
		index: 0,
	}
}

type ArrayDecoder struct {
	mongo bson.A
	index int
}

func (d *ArrayDecoder) GetArray() (model.ArrayDecoder, error) {
	return decodeArray(d.getValue())
}

func (d *ArrayDecoder) GetDate() (time.Time, error) {
	return decodeDate(d.getValue())
}

func (d *ArrayDecoder) GetInt() (int, error) {
	return decodeInt(d.getValue())
}

func (d *ArrayDecoder) GetMap() (model.MapDecoder, error) {
	return decodeMap(d.getValue())
}

func (d *ArrayDecoder) GetObject() (model.ObjectDecoder, error) {
	return decodeObject(d.getValue())
}

func (d *ArrayDecoder) GetRef() (model.Ref, error) {
	return decodeRef(d.getValue())
}

func (d *ArrayDecoder) GetString() (string, error) {
	return decodeString(d.getValue())
}

func (d *ArrayDecoder) Length() int {
	return len(d.mongo)
}

func (d *ArrayDecoder) getValue() (any, error) {
	if d.index >= d.Length() {
		return nil, fmt.Errorf("index %d exceeds length %d", d.index, d.Length())
	}
	val := d.mongo[d.index]
	d.index++
	return val, nil
}
