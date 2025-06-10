package source

import (
	"fmt"
	"time"

	"github.com/ab36245/go-codec"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Array(mongo bson.A) *arraySource {
	return &arraySource{
		mongo: mongo,
		index: 0,
	}
}

type arraySource struct {
	mongo bson.A
	index int
}

func (s *arraySource) GetArray() (codec.ArraySource, error) {
	return decodeArray(s.getValue())
}

func (s *arraySource) GetDate() (time.Time, error) {
	return decodeDate(s.getValue())
}

func (s *arraySource) GetInt() (int, error) {
	return decodeInt(s.getValue())
}

func (s *arraySource) GetString() (string, error) {
	return decodeString(s.getValue())
}

func (s *arraySource) GetObject() (codec.ObjectSource, error) {
	return decodeObject(s.getValue())
}

func (s *arraySource) Length() int {
	return len(s.mongo)
}

func (s *arraySource) getValue() (any, error) {
	if s.index >= s.Length() {
		return nil, fmt.Errorf("index %d exceeds length %d", s.index, s.Length())
	}
	val := s.mongo[s.index]
	s.index++
	return val, nil
}
