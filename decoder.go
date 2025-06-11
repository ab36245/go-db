package db

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/ab36245/go-codec"
)

func newObjectDecoder(mongo bson.M) *objectDecoder {
	return &objectDecoder{
		mongo: mongo,
	}
}

type objectDecoder struct {
	mongo bson.M
}

func (d *objectDecoder) GetArray(name string) (codec.ArrayDecoder, error) {
	return decodeArray(d.getValue(name))
}

func (d *objectDecoder) GetDate(name string) (time.Time, error) {
	return decodeDate(d.getValue(name))
}

func (d *objectDecoder) GetInt(name string) (int, error) {
	return decodeInt(d.getValue(name))
}

func (d *objectDecoder) GetObject(name string) (codec.ObjectDecoder, error) {
	return decodeObject(d.getValue(name))
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

func (d *arrayDecoder) GetArray() (codec.ArrayDecoder, error) {
	return decodeArray(d.getValue())
}

func (d *arrayDecoder) GetDate() (time.Time, error) {
	return decodeDate(d.getValue())
}

func (d *arrayDecoder) GetInt() (int, error) {
	return decodeInt(d.getValue())
}

func (d *arrayDecoder) GetString() (string, error) {
	return decodeString(d.getValue())
}

func (d *arrayDecoder) GetObject() (codec.ObjectDecoder, error) {
	return decodeObject(d.getValue())
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

func decodeArray(value any, err error) (*arrayDecoder, error) {
	if err == nil {
		if value, ok := value.(bson.A); ok {
			return newArrayDecoder(value), nil
		}
		err = fmt.Errorf("expected bson.A, got %T", value)
	}
	return nil, err
}

func decodeDate(value any, err error) (time.Time, error) {
	if err == nil {
		if value, ok := value.(bson.DateTime); ok {
			return value.Time(), nil
		}
		err = fmt.Errorf("expected bson.DateTime, got %T", value)
	}
	return time.Time{}, err
}

func decodeInt(value any, err error) (int, error) {
	if err == nil {
		if value, ok := value.(int32); ok {
			return int(value), nil
		}
		err = fmt.Errorf("expected int32, got %T", value)
	}
	return 0, err
}

func decodeObject(value any, err error) (*objectDecoder, error) {
	if err == nil {
		if value, ok := value.(bson.M); ok {
			return newObjectDecoder(value), nil
		}
		err = fmt.Errorf("expected bson.M, got %T", value)
	}
	return nil, err
}

func decodeString(value any, err error) (string, error) {
	if err == nil {
		if value, ok := value.(string); ok {
			return value, nil
		}
		err = fmt.Errorf("expected string, got %T", value)
	}
	return "", err
}
