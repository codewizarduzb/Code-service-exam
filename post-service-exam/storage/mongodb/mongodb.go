package mongodb

import (
	"context"
	"errors"
	"time"

	pbp "code-service-exam/post-service-exam/genproto/post-proto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepo struct {
	collection *mongo.Collection
}

func NewPostRepo(database *mongo.Database) *PostRepo {
	return &PostRepo{
		collection: database.Collection("posts"),
	}
}

func (r *PostRepo) CreatePost(post *pbp.Post) (*pbp.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	post.CreatedAt = time.Now().Format(time.RFC3339)
	_, err := r.collection.InsertOne(ctx, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *PostRepo) GetPostById(req *pbp.GetPostByIdReq) (*pbp.GetPostByIdRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var result pbp.Post
	err := r.collection.FindOne(ctx, bson.M{"post_id": req.PostId}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &pbp.GetPostByIdRes{Post: &result}, nil
}

func (r *PostRepo) GetPostsByUserId(req *pbp.GetPostsByUserIdReq) (*pbp.GetPostsByUserIdRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"user_id": req.UserId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var posts []*pbp.Post
	for cursor.Next(ctx) {
		var post pbp.Post
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}

		posts = append(posts, &post)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pbp.GetPostsByUserIdRes{Posts: posts}, nil
}

func (r *PostRepo) UpdatePost(req *pbp.UpdatePostReq) (*pbp.UpdatePostRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"post_id": req.PostId}
	update := bson.M{"$set": bson.M{"content": req.Content, "updated_at": time.Now().Format(time.RFC3339)}}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	if result.ModifiedCount == 0 {
		return nil, errors.New("no post updated")
	}

	return &pbp.UpdatePostRes{Success: true}, nil
}

func (r *PostRepo) DeletePost(req *pbp.DeletePostReq) (*pbp.DeletePostRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.collection.DeleteOne(ctx, bson.M{"post_id": req.PostId})
	if err != nil {
		return nil, err
	}
	if result.DeletedCount == 0 {
		return nil, errors.New("no post deleted")
	}

	return &pbp.DeletePostRes{Success: true}, nil
}
