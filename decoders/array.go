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
	return d.reader().getArray()
}

func (d *ArrayDecoder) GetDate() (time.Time, error) {
	return d.reader().getDate()
}

func (d *ArrayDecoder) GetInt() (int, error) {
	return d.reader().getInt()
}

func (d *ArrayDecoder) GetMap() (model.MapDecoder, error) {
	return d.reader().getMap()
}

func (d *ArrayDecoder) GetObject() (model.ObjectDecoder, error) {
	return d.reader().getObject()
}

func (d *ArrayDecoder) GetRef() (model.Ref, error) {
	return d.reader().getRef()
}

func (d *ArrayDecoder) GetString() (string, error) {
	return d.reader().getString()
}

func (d *ArrayDecoder) Length() int {
	return len(d.mongo)
}

func (d *ArrayDecoder) reader() reader {
	return func() (any, error) {
		max := len(d.mongo)
		if d.index >= max {
			return nil, fmt.Errorf("trying to read more than max (%d) values", max)
		}
		value := d.mongo[d.index]
		d.index++
		return value, nil
	}
}
