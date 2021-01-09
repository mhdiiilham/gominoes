package user

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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
func (r *MongoDBRepo) Register(e User) (*User, error) {

	// Check if user exists or not
	_, err := r.FindOne(e.Email)
	if err == nil {
		return nil, errors.New("DUPLICATE EMAIL")
	}

	res, err := r.Collection.InsertOne(r.Ctx, e)
	if err != nil {
		fmt.Println("Error->", err.Error())
		return nil, err
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, err

	}
	return &User{
		ID:       oid,
		Fullname: e.Fullname,
		Email:    e.Email,
		Password: e.Password,
	}, nil
	// return nil
}

// FindOne user function
func (r *MongoDBRepo) FindOne(email string) (*User, error) {
	var user User

	err := r.Collection.FindOne(r.Ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		log.Printf("Error on User Repository FindOne. Error: %v", err)
		return nil, err
	}

	return &user, nil
}
