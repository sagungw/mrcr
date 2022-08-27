package datastore

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LoginHistory struct {
	ID     primitive.ObjectID  `bson:"_id"`
	UserID primitive.ObjectID  `bson:"user_id"`
	Time   primitive.Timestamp `bson:"timestamp"`
}

type HistoryDatastore interface {
	SaveLoginHistory(ctx context.Context, userID string, t time.Time) error
	GetLoginHistory(ctx context.Context, userID string, page int64) ([]*LoginHistory, error)
}

type loginHistoryDs struct {
	coll     *mongo.Collection
	pageSize int64
}

func NewLoginHistoryDatastore() *loginHistoryDs {
	return &loginHistoryDs{
		coll:     db.Collection("login_history"),
		pageSize: 10,
	}
}

func (l *loginHistoryDs) SaveLoginHistory(ctx context.Context, userID string, t time.Time) error {
	userObjectID, _ := primitive.ObjectIDFromHex(userID)
	_, err := l.coll.InsertOne(ctx, bson.D{{"user_id", userObjectID}, {"timestamp", primitive.Timestamp{T: uint32(t.Unix())}}})
	if err != nil {
		return errors.Wrap(err, "datastore: error saving log")
	}

	return nil
}

func (l *loginHistoryDs) GetLoginHistory(ctx context.Context, userID string, page int64) ([]*LoginHistory, error) {
	userObjectID, _ := primitive.ObjectIDFromHex(userID)
	skip := (page - 1) * l.pageSize
	cursor, err := l.coll.Find(ctx, bson.D{{"user_id", userObjectID}}, options.Find().SetSort(bson.D{{"timestamp", 1}}), options.Find().SetSkip(skip), options.Find().SetLimit(l.pageSize))
	if err != nil {
		return nil, errors.Wrap(err, "datastore: error fetching logs")
	}

	result := []*LoginHistory{}
	for cursor.Next(ctx) {
		r := &LoginHistory{}
		if err := cursor.Decode(r); err != nil {
			return nil, errors.Wrap(err, "datastore: error fetching logs")
		}

		result = append(result, r)
	}

	return result, nil
}
