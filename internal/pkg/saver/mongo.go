package saver

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoSaver struct {
	uri      string
	database string
}

func NewMongoSaver(params map[string]string) (*MongoSaver, error) {
	ms := new(MongoSaver)
	err := ms.Initialize(params)
	if err != nil {
		return nil, err
	}
	return ms, nil
}

func (m *MongoSaver) Initialize(params map[string]string) error {
	uri, ok := params["uri"]
	if !ok {
		return errors.New("uri not found")
	}
	m.uri = uri
	db, ok := params["database"]
	if !ok {
		return errors.New("database not found")
	}
	m.database = db
	return nil
}

func (m *MongoSaver) connect(ctx context.Context, opts ...*options.ClientOptions) (*mongo.Client, error) {
	cli, err := mongo.Connect(ctx, opts...)
	return cli, err
}

func (m *MongoSaver) Save(fact *Fact, bucket string) error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(m.uri))
	if err != nil {
		return err
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database(m.database, []*options.DatabaseOptions{}...).Collection(bucket)
	_, err = coll.InsertOne(context.Background(), fact)
	return err
}
