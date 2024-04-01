package postgres

import (
	"fmt"
	"log"
	"testing"

	"code-service-exam/post-service-exam/config"
	pbp "code-service-exam/post-service-exam/genproto/post-proto"
	"code-service-exam/post-service-exam/pkg/db"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserPostgres(t *testing.T) {
	// Connect to database
	cfg := config.Load()
	db, err, _ := db.ConnectToDB(cfg)
	if err != nil {
		return
	}

	repo := NewpostRepo(db)
	// declaration of new post

	post_id := uuid.New().String()
	user_id := uuid.New().String()

	post := &pbp.Post{
		PostId:       post_id,
		UserId:       user_id,
		Content:      "Test content",
		Title:        "Test title",
		Likes:        100,
		Dislikes:     10,
		Views:        1000,
		MediaUrl:     "231423132.jpg",
		RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJzdHJpbmcifQ.MRSgGgVl5hlW0Al0865oIC8gxXIEBPI7S0CUiHNdMTI",
	}

	fmt.Println(post)
	// test for creating post
	createdpost, err := repo.CreatePost(post)
	assert.NoError(t, err)
	assert.Equal(t, post.Content, createdpost.Content)

	// test for updating post content
	updatedpost, err := repo.UpdatePost(&pbp.UpdatePostReq{
		PostId:  post.PostId,
		Content: post.Content,
	})
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, true, updatedpost.Success)

	// test for deleting post
	deletedpost, err := repo.DeletePost(&pbp.DeletePostReq{
		PostId: post.PostId,
	})
	assert.NoError(t, err)
	assert.Equal(t, true, deletedpost.Success)
}
