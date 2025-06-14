package encoders

import (
	"fmt"
	"time"

	"github.com/ab36245/go-model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func encodeArray(length int, handler func(model.ArrayEncoder)) bson.A {
	ae := NewArrayEncoder(length)
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
	oe := NewObjectEncoder()
	handler(oe)
	return oe.mongo
}

func encodeMap(length int, handler func(model.MapEncoder)) bson.M {
	me := NewMapEncoder(length)
	handler(me)
	return me.Value()
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
