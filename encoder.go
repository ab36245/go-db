package db

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/ab36245/go-codec"
)

func newObjectEncoder() *objectEncoder {
	return &objectEncoder{
		mongo: make(bson.M),
	}
}

type objectEncoder struct {
	mongo bson.M
}

func (e *objectEncoder) PutArray(name string, length int, f func(codec.ArrayEncoder)) {
	e.putValue(name, encodeArray(length, f))
}

func (e *objectEncoder) PutDate(name string, value time.Time) {
	e.putValue(name, encodeDate(value))
}

func (e *objectEncoder) PutInt(name string, value int) {
	e.putValue(name, encodeInt(value))
}

func (e *objectEncoder) PutObject(name string, f func(codec.ObjectEncoder)) {
	e.putValue(name, encodeObject(f))
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

func (e *arrayEncoder) PutArray(length int, f func(codec.ArrayEncoder)) {
	e.putValue(encodeArray(length, f))
}

func (e *arrayEncoder) PutDate(value time.Time) {
	e.putValue(encodeDate(value))
}

func (e *arrayEncoder) PutInt(value int) {
	e.putValue(encodeInt(value))
}

func (e *arrayEncoder) PutObject(f func(codec.ObjectEncoder)) {
	e.putValue(encodeObject(f))
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

func encodeArray(length int, handler func(codec.ArrayEncoder)) any {
	ae := newArrayEncoder(length)
	handler(ae)
	return ae.mongo
}

func encodeDate(value time.Time) any {
	return bson.NewDateTimeFromTime(value)
}

func encodeInt(value int) any {
	return value
}

func encodeObject(handler func(codec.ObjectEncoder)) any {
	oe := newObjectEncoder()
	handler(oe)
	return oe.mongo
}

func encodeString(value string) any {
	return value
}
