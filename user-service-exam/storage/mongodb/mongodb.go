package mongodb

import (
	"context"
	"errors"
	"time"

	pb "code-service-exam/user-service-exam/genproto/user-proto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepo struct {
	collection *mongo.Collection
}

func NewUserRepo(database *mongo.Database) *UserRepo {
	return &UserRepo{
		collection: database.Collection("users"),
	}
}

// create user
func (r *UserRepo) CreateUser(user *pb.User) (*pb.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user.CreatedAt = time.Now().Format(time.RFC3339)
	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// get user
func (r *UserRepo) GetUser(req *pb.GetUserReq) (*pb.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user pb.User
	err := r.collection.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// list users
func (r *UserRepo) ListUsers(req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	limit := int64(req.Limit)
	offset := int64((req.Page - 1) * req.Limit)

	cursor, err := r.collection.Find(ctx, bson.M{}, &options.FindOptions{
		Limit: &limit,
		Skip:  &offset,
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*pb.User
	for cursor.Next(ctx) {
		var user pb.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.ListUsersResponse{
		Users: users,
	}, nil
}

// update user
func (r *UserRepo) UpdateUser(req *pb.UpdateUserReq) (*pb.UpdateUserRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"email": req.Email}
	update := bson.M{"$set": bson.M{"username": "striker", "updated_at": time.Now()}}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return &pb.UpdateUserRes{Success: false}, err
	}

	return &pb.UpdateUserRes{Success: true}, nil
}

// delete user
func (r *UserRepo) DeleteUser(req *pb.DeleteUserReq) (*pb.DeleteUserRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.collection.DeleteOne(ctx, bson.M{"email": req.Email})
	if err != nil {
		return &pb.DeleteUserRes{Success: false}, err
	}
	if result.DeletedCount == 0 {
		return &pb.DeleteUserRes{Success: false}, errors.New("no user deleted")
	}

	return &pb.DeleteUserRes{Success: true}, nil
}
