package decoders

import (
	"fmt"
	"time"

	"github.com/ab36245/go-model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func NewMapDecoder(mongo bson.D) *MapDecoder {
	return &MapDecoder{
		index:   0,
		keyNext: true,
		mongo:   mongo,
	}
}

type MapDecoder struct {
	index   int
	keyNext bool
	mongo   bson.D
}

func (d *MapDecoder) GetArray() (model.ArrayDecoder, error) {
	return d.reader().getArray()
}

func (d *MapDecoder) GetBool() (bool, error) {
	return d.reader().getBool()
}

func (d *MapDecoder) GetBytes() ([]byte, error) {
	return d.reader().getBytes()
}

func (d *MapDecoder) GetFloat() (float64, error) {
	return d.reader().getFloat()
}

func (d *MapDecoder) GetInt() (int, error) {
	return d.reader().getInt()
}

func (d *MapDecoder) GetMap() (model.MapDecoder, error) {
	return d.reader().getMap()
}

func (d *MapDecoder) GetObject() (model.ObjectDecoder, error) {
	return d.reader().getObject()
}

func (d *MapDecoder) GetRef() (model.Ref, error) {
	return d.reader().getRef()
}

func (d *MapDecoder) GetString() (string, error) {
	return d.reader().getString()
}

func (d *MapDecoder) GetTime() (time.Time, error) {
	return d.reader().getTime()
}

func (d *MapDecoder) Length() int {
	return len(d.mongo)
}

func (d *MapDecoder) reader() reader {
	return func() (any, error) {
		max := len(d.mongo)
		if d.index >= max {
			return nil, fmt.Errorf("trying to read more than max (%d) values", max)
		}
		var value any
		if d.keyNext {
			value = d.mongo[d.index].Key
			d.keyNext = false
		} else {
			value = d.mongo[d.index].Value
			d.index++
			d.keyNext = true
		}
		return value, nil
	}
}
