package mongo

import (
	"GoNews/pkg/storage"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	db *mongo.Client
}

func New(dbURL string) (*Mongo, error) {
	mongoOpts := options.Client().ApplyURI(dbURL)
	db, err := mongo.Connect(context.Background(), mongoOpts)
	if err != nil {
		return nil, err
	}
	return &Mongo{db: db}, nil
}

func (m *Mongo) AddPost(p storage.Post) error {
	db := fmt.Sprintf("%s", viper.Get("mongo.database"))
	_, err := m.db.Database(db).Collection("Posts").InsertOne(context.Background(), p)
	if err != nil {
		return err
	}
	return nil
}

func (m *Mongo) UpdatePost(p storage.Post) error {
	db := fmt.Sprintf("%s", viper.Get("mongo.database"))
	filter := bson.D{{"ID", p.ID}}
	update := bson.D{{"$set", bson.D{{"Author ID", p.AuthorID}, {"Title", p.Title}, {"Content", p.Content}}}}
	res, err := m.db.Database(db).Collection("Posts").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		err = fmt.Errorf("post with id#%d not found", p.ID)
		return err
	}
	return nil
}

func (m *Mongo) DeletePost(p storage.Post) error {
	db := fmt.Sprintf("%s", viper.Get("mongo.database"))
	filter := bson.D{{"ID", p.ID}}
	res, err := m.db.Database(db).Collection("Posts").DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		err = fmt.Errorf("post with id#%d not found", p.ID)
		return err
	}
	return nil
}

func (m *Mongo) Posts() ([]storage.Post, error) {
	db := fmt.Sprintf("%s", viper.Get("mongo.database"))
	collection := m.db.Database(db).Collection("Posts")
	filter := bson.D{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
		}
	}(cur, context.Background())
	var data []storage.Post
	for cur.Next(context.Background()) {
		var p storage.Post
		err := cur.Decode(&p)
		if err != nil {
			return nil, err
		}
		data = append(data, p)
	}
	return data, cur.Err()
}
