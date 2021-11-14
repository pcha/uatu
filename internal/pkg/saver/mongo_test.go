package saver

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//func TestMongoSaver_Initialize(t *testing.T) {
//	type TestCase struct {
//
//	}
//}

func TestMongoSaver_Save(t *testing.T) {
	type TestCase struct {
		mockedClient       *mongo.Client
		mockedClientErr    error
		mockedInsertResult *mongo.InsertOneResult
		mockedInsertErr    error
		wantedErr          error
		assertInsert       bool
	}

	clientErr := errors.New("error on client")
	insertionErr := errors.New("error on insertion")
	tests := map[string]TestCase{
		"Sunny case": {
			mockedClient:       &mongo.Client{},
			mockedInsertResult: &mongo.InsertOneResult{InsertedID: "1234"},
			assertInsert:       true,
		},
		"Error on client": {
			mockedClientErr: clientErr,
			wantedErr:       clientErr,
		},
		"Error on insertion": {
			mockedClient:    &mongo.Client{},
			mockedInsertErr: insertionErr,
			wantedErr:       insertionErr,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			mongoURI := "mongodb://test"
			dbName := "testdb"
			fact := &Fact{
				"message": "some mesasage",
				"attributes": map[string]string{
					"attr1": "val1",
					"attr2": "val2",
				},
				"tags": []string{
					"tag1",
					"tag2",
				},
			}
			var insertedDocs []interface{}
			monkey.Patch(mongo.Connect, func(ctx context.Context, opts ...*options.ClientOptions) (*mongo.Client, error) {
				return tc.mockedClient, tc.mockedClientErr
			})
			monkey.PatchInstanceMethod(reflect.TypeOf(&mongo.Collection{}), "InsertOne", func(coll *mongo.Collection, ctx context.Context, document interface{},
				opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
				insertedDocs = append(insertedDocs, document)
				return tc.mockedInsertResult, tc.mockedInsertErr
			})

			saver := &MongoSaver{
				uri:      mongoURI,
				database: dbName,
			}

			bucket := "buckettest"
			err := saver.Save(fact, bucket)
			assert.Equal(t, tc.wantedErr, err)
			if tc.assertInsert {
				assert.Equal(t, insertedDocs, []interface{}{fact})
			}
		})
	}
}

func TestNewMongoSaver(t *testing.T) {
	type TestCase struct {
		params      map[string]string
		wantedSaver *MongoSaver
		wantsErr    bool
	}

	tests := map[string]TestCase{
		"Sunny case": {
			params: map[string]string{
				"uri":      "mongodb://test",
				"database": "testdb",
			},
			wantedSaver: &MongoSaver{
				uri:      "mongodb://test",
				database: "testdb",
			},
		},
		"Missing uri": {
			params: map[string]string{
				"database": "testdb",
			},
			wantsErr: true,
		},
		"Missing database": {
			params: map[string]string{
				"uri": "mongodb://test",
			},
			wantsErr: true,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			saver, err := NewMongoSaver(tc.params)
			if tc.wantsErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tc.wantedSaver, saver)
		})
	}
}
