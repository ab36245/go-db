package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	codec "github.com/ab36245/go-model"
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
	var records []T
	for cursor.Next(ctx) {
		var data bson.M
		if err := cursor.Decode(&data); err != nil {
			return nil, fmt.Errorf("raw decode failed: %w", err)
		}
		record, err := t.decode(data)
		if err != nil {
			return nil, fmt.Errorf("type decode failed: %w", err)
		}
		records = append(records, record)
	}
	return records, nil
}

func (t *Table[T]) Insert(record T) error {
	ctx := context.TODO()
	data, err := t.encode(record)
	if err != nil {
		return err
	}
	res, err := t.mongo.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	_ = res
	return nil
}

func (t *Table[T]) decode(data bson.M) (T, error) {
	decoder := newDecoder(data)
	return t.codec.Decode(decoder)
}

func (t *Table[T]) encode(record T) (bson.M, error) {
	encoder := newEncoder()
	if err := t.codec.Encode(encoder, record); err != nil {
		return nil, err
	}
	return encoder.Value(), nil
}
