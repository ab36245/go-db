package db

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/ab36245/go-model"
)

func newObjectDecoder(mongo bson.M) *objectDecoder {
	return &objectDecoder{
		mongo: mongo,
	}
}

type objectDecoder struct {
	mongo bson.M
}

func (d *objectDecoder) GetArray(name string) (model.ArrayDecoder, error) {
	return decodeArray(d.getValue(name))
}

func (d *objectDecoder) GetDate(name string) (time.Time, error) {
	return decodeDate(d.getValue(name))
}

func (d *objectDecoder) GetInt(name string) (int, error) {
	return decodeInt(d.getValue(name))
}

func (d *objectDecoder) GetMap(name string) (model.MapDecoder, error) {
	return decodeMap(d.getValue(name))
}

func (d *objectDecoder) GetObject(name string) (model.ObjectDecoder, error) {
	return decodeObject(d.getValue(name))
}

func (d *objectDecoder) GetRef(name string) (model.Ref, error) {
	return decodeRef(d.getValue(name))
}

func (d *objectDecoder) GetString(name string) (string, error) {
	return decodeString(d.getValue(name))
}

func (d *objectDecoder) getValue(name string) (any, error) {
	val, ok := d.mongo[name]
	if !ok {
		return nil, fmt.Errorf("field '%s': unknown", name)
	}
	return val, nil
}

func newArrayDecoder(mongo bson.A) *arrayDecoder {
	return &arrayDecoder{
		mongo: mongo,
		index: 0,
	}
}

type arrayDecoder struct {
	mongo bson.A
	index int
}

func (d *arrayDecoder) GetArray() (model.ArrayDecoder, error) {
	return decodeArray(d.getValue())
}

func (d *arrayDecoder) GetDate() (time.Time, error) {
	return decodeDate(d.getValue())
}

func (d *arrayDecoder) GetInt() (int, error) {
	return decodeInt(d.getValue())
}

func (d *arrayDecoder) GetMap() (model.MapDecoder, error) {
	return decodeMap(d.getValue())
}

func (d *arrayDecoder) GetObject() (model.ObjectDecoder, error) {
	return decodeObject(d.getValue())
}

func (d *arrayDecoder) GetRef() (model.Ref, error) {
	return decodeRef(d.getValue())
}

func (d *arrayDecoder) GetString() (string, error) {
	return decodeString(d.getValue())
}

func (d *arrayDecoder) Length() int {
	return len(d.mongo)
}

func (d *arrayDecoder) getValue() (any, error) {
	if d.index >= d.Length() {
		return nil, fmt.Errorf("index %d exceeds length %d", d.index, d.Length())
	}
	val := d.mongo[d.index]
	d.index++
	return val, nil
}

func newMapDecoder(mongo bson.M) *mapDecoder {
	var items []any
	for k, v := range mongo {
		items = append(items, k, v)
	}
	return &mapDecoder{
		items: items,
		mongo: mongo,
		index: 0,
	}
}

type mapDecoder struct {
	items []any
	mongo bson.M
	index int
}

func (d *mapDecoder) GetArray() (model.ArrayDecoder, error) {
	return decodeArray(d.getValue())
}

func (d *mapDecoder) GetDate() (time.Time, error) {
	return decodeDate(d.getValue())
}

func (d *mapDecoder) GetInt() (int, error) {
	return decodeInt(d.getValue())
}

func (d *mapDecoder) GetKey() (string, error) {
	return decodeString(d.getValue())
}

func (d *mapDecoder) GetMap() (model.MapDecoder, error) {
	return decodeMap(d.getValue())
}

func (d *mapDecoder) GetObject() (model.ObjectDecoder, error) {
	return decodeObject(d.getValue())
}

func (d *mapDecoder) GetRef() (model.Ref, error) {
	return decodeRef(d.getValue())
}

func (d *mapDecoder) GetString() (string, error) {
	return decodeString(d.getValue())
}

func (d *mapDecoder) Length() int {
	return len(d.items)
}

func (d *mapDecoder) getValue() (any, error) {
	if d.index >= d.Length() {
		return nil, fmt.Errorf("index %d exceeds length %d", d.index, d.Length())
	}
	val := d.items[d.index]
	d.index++
	return val, nil
}

func decodeArray(from any, err error) (*arrayDecoder, error) {
	cast, err := castAs[bson.A](from, err)
	if err != nil {
		return nil, err
	}
	return newArrayDecoder(cast), nil
}

func decodeDate(from any, err error) (time.Time, error) {
	cast, err := castAs[bson.DateTime](from, err)
	if err != nil {
		return time.Time{}, err
	}
	return cast.Time(), nil
}

func decodeInt(from any, err error) (int, error) {
	cast, err := castAs[int32](from, err)
	if err != nil {
		return 0, err
	}
	return int(cast), nil
}

func decodeMap(from any, err error) (*mapDecoder, error) {
	cast, err := castAs[bson.M](from, err)
	if err != nil {
		return nil, err
	}
	return newMapDecoder(cast), nil
}

func decodeObject(from any, err error) (*objectDecoder, error) {
	cast, err := castAs[bson.M](from, err)
	if err != nil {
		return nil, err
	}
	return newObjectDecoder(cast), nil
}

func decodeRef(from any, err error) (model.Ref, error) {
	cast, err := castAs[bson.ObjectID](from, err)
	if err != nil {
		return model.Ref(""), nil
	}
	return model.Ref(cast.Hex()), nil
}

func decodeString(from any, err error) (string, error) {
	cast, err := castAs[string](from, err)
	if err != nil {
		return "", nil
	}
	return cast, nil
}

func castAs[T any](from any, err error) (T, error) {
	cast := *new(T)
	if err != nil {
		return cast, nil
	}
	cast, ok := from.(T)
	if !ok {
		return cast, fmt.Errorf("expected %T, got %T", cast, from)
	}
	return cast, nil
}
