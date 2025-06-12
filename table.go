package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	codec "github.com/ab36245/go-defs"
)

func NewTable[T any](db *Database, name string, codec codec.Codec[T]) *Table[T] {
	mongo := db.mongo.Collection(name)
	return &Table[T]{
		mongo: mongo,
		codec: codec,
	}
}

type Table[T any] struct {
	mongo *mongo.Collection
	codec codec.Codec[T]
}

func (t *Table[T]) Drop() error {
	ctx := context.TODO()
	err := t.mongo.Drop(ctx)
	return err
}

func (t *Table[T]) Find() ([]T, error) {
	ctx := context.TODO()
	filter := bson.M{}
	cursor, err := t.mongo.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("table '%s' search failed: %w", t.mongo.Name(), err)
	}
	defer cursor.Close(ctx)
	var list []T
	for cursor.Next(ctx) {
		var raw bson.M
		if err := cursor.Decode(&raw); err != nil {
			return nil, fmt.Errorf("raw decode failed: %w", err)
		}
		decoder := newObjectDecoder(raw)
		item, err := t.decode(decoder)
		if err != nil {
			return nil, fmt.Errorf("type decode failed: %w", err)
		}
		list = append(list, item)
	}
	return list, nil
}

func (t *Table[T]) Insert(record T) error {
	encoder := newObjectEncoder()
	t.encode(encoder, record)
	ctx := context.TODO()
	res, err := t.mongo.InsertOne(ctx, encoder.Value())
	if err != nil {
		return err
	}
	_ = res
	return nil
}

func (t *Table[T]) decode(decoder *objectDecoder) (T, error) {
	return t.codec.Decode(decoder)
}

func (t *Table[T]) encode(encoder *objectEncoder, record T) {
	t.codec.Encode(encoder, record)
}
