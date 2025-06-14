package decoders

import (
	"fmt"
	"time"

	"github.com/ab36245/go-model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func decodeAs[T any](from any, err error) (T, error) {
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

func decodeArray(from any, err error) (*ArrayDecoder, error) {
	cast, err := decodeAs[bson.A](from, err)
	if err != nil {
		return nil, err
	}
	return NewArrayDecoder(cast), nil
}

func decodeDate(from any, err error) (time.Time, error) {
	cast, err := decodeAs[bson.DateTime](from, err)
	if err != nil {
		return time.Time{}, err
	}
	return cast.Time(), nil
}

func decodeInt(from any, err error) (int, error) {
	cast, err := decodeAs[int32](from, err)
	if err != nil {
		return 0, err
	}
	return int(cast), nil
}

func decodeMap(from any, err error) (*MapDecoder, error) {
	cast, err := decodeAs[bson.M](from, err)
	if err != nil {
		return nil, err
	}
	return NewMapDecoder(cast), nil
}

func decodeObject(from any, err error) (*ObjectDecoder, error) {
	cast, err := decodeAs[bson.M](from, err)
	if err != nil {
		return nil, err
	}
	return NewObjectDecoder(cast), nil
}

func decodeRef(from any, err error) (model.Ref, error) {
	cast, err := decodeAs[bson.ObjectID](from, err)
	if err != nil {
		return model.Ref(""), nil
	}
	return model.Ref(cast.Hex()), nil
}

func decodeString(from any, err error) (string, error) {
	cast, err := decodeAs[string](from, err)
	if err != nil {
		return "", nil
	}
	return cast, nil
}
