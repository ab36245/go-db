package source

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/ab36245/go-codec"
)

func Object(mongo bson.M) *objectSource {
	return &objectSource{
		mongo: mongo,
	}
}

type objectSource struct {
	mongo bson.M
}

func (s *objectSource) GetArray(name string) (codec.ArraySource, error) {
	return decodeArray(s.getValue(name))
}

func (s *objectSource) GetDate(name string) (time.Time, error) {
	return decodeDate(s.getValue(name))
}

func (s *objectSource) GetInt(name string) (int, error) {
	return decodeInt(s.getValue(name))
}

func (s *objectSource) GetObject(name string) (codec.ObjectSource, error) {
	return decodeObject(s.getValue(name))
}

func (s *objectSource) GetString(name string) (string, error) {
	return decodeString(s.getValue(name))
}

func (s *objectSource) getValue(name string) (any, error) {
	val, ok := s.mongo[name]
	if !ok {
		return nil, fmt.Errorf("field '%s': unknown", name)
	}
	return val, nil
}
