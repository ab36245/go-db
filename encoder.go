package db

import "github.com/ab36245/go-db/encoders"

func newEncoder() *encoder {
	return &encoder{
		encoders.NewObjectEncoder(),
	}
}

type encoder struct {
	*encoders.ObjectEncoder
}
