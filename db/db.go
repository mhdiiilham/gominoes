package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewMongoDBConnection create new MongoDB Connection
func NewMongoDBConnection(user, pass, db string) (*mongo.Client, error) {
	mongoURI := fmt.Sprintf(
		"mongodb+srv://%s:%s@cluster0.ub9ns.mongodb.net/%s?retryWrites=true&w=majority",
		user,
		pass,
		db,
	)

	client, _ := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err := client.Connect(context.TODO()); err != nil {
		return nil, err
	}
	return client, nil
}
