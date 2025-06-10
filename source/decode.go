package source

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/ab36245/go-codec"
)

func decodeArray(value any, err error) (codec.ArraySource, error) {
	if err == nil {
		if value, ok := value.(bson.A); ok {
			return Array(value), nil
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

func decodeObject(value any, err error) (codec.ObjectSource, error) {
	if err == nil {
		if value, ok := value.(bson.M); ok {
			return Object(value), nil
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
