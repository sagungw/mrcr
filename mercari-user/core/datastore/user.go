package datastore

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

type UserDatastore interface {
	Register(ctx context.Context, email, password string) error
	Login(ctx context.Context, email, password string) (*User, error)
}

type userds struct {
	coll *mongo.Collection
}

func NewUserDatastore() *userds {
	return &userds{
		coll: db.Collection("users"),
	}
}

func (u *userds) Register(ctx context.Context, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "datastore: error registering user")
	}

	_, err = u.coll.InsertOne(ctx, bson.D{{"email", email}, {"password", hashedPassword}})
	if err != nil {
		return errors.Wrap(err, "datastore: error registering user")
	}

	return nil
}

func (u *userds) Login(ctx context.Context, email, password string) (*User, error) {
	user := &User{}
	result := u.coll.FindOne(ctx, bson.D{{"email", email}}, options.FindOne().SetProjection(bson.D{{"_id", 1}, {"email", 1}, {"password", 1}}))
	err := result.Decode(user)
	if err != nil {
		return nil, errors.Wrap(err, "datastore: login error")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, InvalidPassword{}
	}

	return user, nil
}

type InvalidPassword struct{}

func (e InvalidPassword) Error() string {
	return "invalid password"
}
