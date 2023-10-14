package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	client *mongo.Client
	db     string
}

func NewStore(uri, db string) (*Store, error) {
	client, err := newMongoClient(uri)
	if err != nil {
		return nil, err
	}

	return &Store{
		client: client,
		db:     db,
	}, nil
}

func (s Store) Disconnect() error {
	return s.client.Disconnect(context.Background())
}

func (s Store) Col(name string) *mongo.Collection {
	return s.client.Database(s.db).Collection(name)
}

func (s Store) Insert(col *mongo.Collection, data interface{}) error {
	return insert(col, data)
}

func (s Store) InsertMany(col *mongo.Collection, data []interface{}) error {
	return insertMany(col, data)
}

func (s Store) Update(col *mongo.Collection, filter M, data interface{}, out interface{}) error {
	return update(col, filter, data, out)
}

func (s Store) UpdateMany(col *mongo.Collection, filter M, data []interface{}) error {
	return updateMany(col, filter, data)
}

func (s Store) Upsert(col *mongo.Collection, filter M, data interface{}, out interface{}) error {
	return upsert(col, filter, data, out)
}

func (s Store) Delete(col *mongo.Collection, filter M, out interface{}) error {
	return delete(col, filter, out)
}

func (s Store) DeleteMany(col *mongo.Collection, filter M) error {
	return deleteMany(col, filter)
}

func (s Store) Aggregate(col *mongo.Collection, pipe []M, out interface{}) error {
	return aggregate(col, pipe, out)
}

func (s Store) FindOne(col *mongo.Collection, filter M, out interface{}) error {
	return findOne(col, filter, out)
}

func (s Store) Find(col *mongo.Collection, filter M, options *options.FindOptions, out interface{}) error {
	return find(col, filter, options, out)
}
