package decoders

import (
	"fmt"
	"time"

	"github.com/ab36245/go-model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func NewMapDecoder(mongo bson.M) *MapDecoder {
	var items []any
	for k, v := range mongo {
		items = append(items, k, v)
	}
	return &MapDecoder{
		items: items,
		index: 0,
	}
}

type MapDecoder struct {
	items []any
	index int
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
	return len(d.items)
}

func (d *MapDecoder) reader() reader {
	return func() (any, error) {
		if d.index >= d.Length() {
			return nil, fmt.Errorf("index %d exceeds map length %d", d.index, d.Length())
		}
		value := d.items[d.index]
		d.index++
		return value, nil
	}
}
