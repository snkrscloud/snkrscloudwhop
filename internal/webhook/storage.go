package webhook

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// how the user is stored in the database
type UserDB struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username"`
	Email     string             `bson:"email"`
	AvatarURL string             `bson:"avatar_url"`
	Valid     bool               `bson:"valid"`
}

type WebhookStorage struct {
	db *mongo.Database
}

func NewWebhookStorage(db *mongo.Database) *WebhookStorage {
	return &WebhookStorage{
		db: db,
	}
}

func (s *WebhookStorage) updateUser(id, username, email, avatarUrl string, valid bool, ctx context.Context) error {
	collection := s.db.Collection("users")

	// convert the id string to an ObjectID
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// create the filter
	filter := bson.M{"_id": oid}

	// create the update
	update := bson.M{
		"$set": bson.M{
			"username":   username,
			"email":      email,
			"avatar_url": avatarUrl,
			"valid":      valid,
		},
	}

	// update the user
	_, err = collection.UpdateOne(ctx, filter, update)
	return err
}
