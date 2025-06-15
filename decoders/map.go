package decoders

import (
	"fmt"
	"time"

	"github.com/ab36245/go-model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func NewMapDecoder(mongo bson.D) *MapDecoder {
	length := len(mongo)
	var items []any
	for _, e := range mongo {
		items = append(items, e.Key, e.Value)
	}
	return &MapDecoder{
		index:  0,
		items:  items,
		length: length,
	}
}

type MapDecoder struct {
	index  int
	items  []any
	length int
}

func (d *MapDecoder) GetArray() (model.ArrayDecoder, error) {
	return d.reader().getArray()
}

func (d *MapDecoder) GetDate() (time.Time, error) {
	return d.reader().getDate()
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

func (d *MapDecoder) Length() int {
	return d.length
}

func (d *MapDecoder) reader() reader {
	return func() (any, error) {
		if d.index >= len(d.items) {
			return nil, fmt.Errorf("index %d outside total map size %d", d.index, len(d.items))
		}
		value := d.items[d.index]
		d.index++
		return value, nil
	}
}
