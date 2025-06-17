package decoders

import (
	"fmt"
	"time"

	"github.com/ab36245/go-model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// type reader struct {
// 	get func() (any, error)
// }

type reader func() (any, error)

func (r reader) getArray() (*ArrayDecoder, error) {
	value, err := getAs[bson.A](r)
	if err != nil {
		return nil, err
	}
	return NewArrayDecoder(value), nil
}

func (r reader) getBool() (bool, error) {
	return getAs[bool](r)
}

func (r reader) getBytes() ([]byte, error) {
	return getAs[[]byte](r)
}

func (r reader) getFloat() (float64, error) {
	return getAs[float64](r)
}

func (r reader) getInt() (int, error) {
	value, err := getAs[int32](r)
	if err != nil {
		return 0, err
	}
	return int(value), nil
}

func (r reader) getMap() (*MapDecoder, error) {
	value, err := getAs[bson.D](r)
	if err != nil {
		return nil, err
	}
	return NewMapDecoder(value), nil
}

func (r reader) getObject() (*ObjectDecoder, error) {
	value, err := getAs[bson.M](r)
	if err != nil {
		return nil, err
	}
	return NewObjectDecoder(value), nil
}

func (r reader) getRef() (model.Ref, error) {
	value, err := getAs[bson.ObjectID](r)
	if err != nil {
		return model.Ref(""), nil
	}
	return model.Ref(value.Hex()), nil
}

func (r reader) getString() (string, error) {
	return getAs[string](r)
}

func (r reader) getTime() (time.Time, error) {
	value, err := getAs[bson.DateTime](r)
	if err != nil {
		return time.Time{}, err
	}
	return value.Time(), nil
}

func getAs[T any](getter func() (any, error)) (T, error) {
	value, err := getter()
	if err != nil {
		return *new(T), err
	}
	switch value := value.(type) {
	case T:
		return value, nil
	default:
		return *new(T), fmt.Errorf("expected %T, got %T", *new(T), value)
	}
}
