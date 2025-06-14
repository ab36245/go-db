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
	return decodeArray(d.getValue())
}

func (d *MapDecoder) GetDate() (time.Time, error) {
	return decodeDate(d.getValue())
}

func (d *MapDecoder) GetInt() (int, error) {
	return decodeInt(d.getValue())
}

func (d *MapDecoder) GetKey() (string, error) {
	return decodeString(d.getValue())
}

func (d *MapDecoder) GetMap() (model.MapDecoder, error) {
	return decodeMap(d.getValue())
}

func (d *MapDecoder) GetObject() (model.ObjectDecoder, error) {
	return decodeObject(d.getValue())
}

func (d *MapDecoder) GetRef() (model.Ref, error) {
	return decodeRef(d.getValue())
}

func (d *MapDecoder) GetString() (string, error) {
	return decodeString(d.getValue())
}

func (d *MapDecoder) Length() int {
	return len(d.items)
}

func (d *MapDecoder) getValue() (any, error) {
	if d.index >= d.Length() {
		return nil, fmt.Errorf("index %d exceeds length %d", d.index, d.Length())
	}
	val := d.items[d.index]
	d.index++
	return val, nil
}
