package db

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/ab36245/go-model"
)

func newObjectEncoder() *objectEncoder {
	return &objectEncoder{
		mongo: make(bson.M),
	}
}

type objectEncoder struct {
	mongo bson.M
}

func (e *objectEncoder) PutArray(name string, length int, f func(model.ArrayEncoder)) {
	e.putValue(name, encodeArray(length, f))
}

func (e *objectEncoder) PutDate(name string, value time.Time) {
	e.putValue(name, encodeDate(value))
}

func (e *objectEncoder) PutInt(name string, value int) {
	e.putValue(name, encodeInt(value))
}

func (e *objectEncoder) PutObject(name string, f func(model.ObjectEncoder)) {
	e.putValue(name, encodeObject(f))
}

func (e *objectEncoder) PutRef(name string, value model.Ref) {
	e.putValue(name, encodeRef(value))
}

func (e *objectEncoder) PutString(name string, value string) {
	e.putValue(name, encodeString(value))
}

func (e *objectEncoder) Value() bson.M {
	return e.mongo
}

func (e *objectEncoder) putValue(name string, value any) {
	e.mongo[name] = value
}

func newArrayEncoder(length int) *arrayEncoder {
	return &arrayEncoder{
		mongo: make(bson.A, 0, length),
	}
}

type arrayEncoder struct {
	mongo bson.A
}

func (e *arrayEncoder) PutArray(length int, f func(model.ArrayEncoder)) {
	e.putValue(encodeArray(length, f))
}

func (e *arrayEncoder) PutDate(value time.Time) {
	e.putValue(encodeDate(value))
}

func (e *arrayEncoder) PutInt(value int) {
	e.putValue(encodeInt(value))
}

func (e *arrayEncoder) PutObject(f func(model.ObjectEncoder)) {
	e.putValue(encodeObject(f))
}

func (e *arrayEncoder) PutRef(value model.Ref) {
	e.putValue(encodeRef(value))
}

func (e *arrayEncoder) PutString(value string) {
	e.putValue(encodeString(value))
}

func (e *arrayEncoder) Value() bson.A {
	return e.mongo
}

func (e *arrayEncoder) putValue(value any) {
	e.mongo = append(e.mongo, value)
}

func encodeArray(length int, handler func(model.ArrayEncoder)) bson.A {
	ae := newArrayEncoder(length)
	handler(ae)
	return ae.mongo
}

func encodeDate(value time.Time) bson.DateTime {
	return bson.NewDateTimeFromTime(value)
}

func encodeInt(value int) int32 {
	return int32(value)
}

func encodeObject(handler func(model.ObjectEncoder)) bson.M {
	oe := newObjectEncoder()
	handler(oe)
	return oe.mongo
}

func encodeRef(value model.Ref) bson.ObjectID {
	id, err := bson.ObjectIDFromHex(string(value))
	if err != nil {
		// TODO is panic the right thing to do here?
		panic(fmt.Sprintf("can't encode %s as ObjectID: %s", value, err))
	}
	return id
}

func encodeString(value string) string {
	return value
}
