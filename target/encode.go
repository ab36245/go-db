package target

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func encodeDate(value time.Time) any {
	return bson.NewDateTimeFromTime(value)
}

func encodeInt(value int) any {
	return value
}

func encodeString(value string) any {
	return value
}
