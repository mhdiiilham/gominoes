package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDBRepo struct
type MongoDBRepo struct {
	Collection *mongo.Collection
	Ctx        context.Context
}

// NewMongoDBRepository function
func NewMongoDBRepository(collection *mongo.Collection) *MongoDBRepo {
	ctx := context.Background()
	return &MongoDBRepo{
		Collection: collection,
		Ctx:        ctx,
	}
}

// Register user function
func (r *MongoDBRepo) Register(e User) string {
	res, err := r.Collection.InsertOne(r.Ctx, e)
	if err != nil {
		return ""
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return ""
	}
	return oid.Hex()
}

// FindOne user function
func (r *MongoDBRepo) FindOne(email string) (*User, error) {
	var user User
	return &user, nil
}
