package postgres

import (
	"code-service-exam/comment-service-exam/config"
	pbc "code-service-exam/comment-service-exam/genproto/comment-proto"
	"code-service-exam/comment-service-exam/pkg/db"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentPostgres(t *testing.T) {
	// Connect to database
	cfg := config.Load()
	db, err, _ := db.ConnectToDB(cfg)
	if err != nil {
		return
	}

	repo := NewCommentRepo(db)

//			TESTS

	// test for create comment
	t.Run("TestCreateComment", func(t *testing.T) {
		comment := &pbc.Comment{
			CommentId: "540e8400-e29b-41d4-a716-446655440000",
			PostId:    "6ca7b810-9dad-11d1-80b4-00c04fd430c8",
			UserId:    "9a77e6fc-8d2e-4a02-8390-7be2f2a118e7",
			Content:   "This is a test comment.",
		}

		createdComment, err := repo.CreateComment(comment)

		assert.NoError(t, err)
		assert.NotNil(t, createdComment)
		assert.Equal(t, comment.CommentId, createdComment.CommentId)
		assert.Equal(t, comment.PostId, createdComment.PostId)
		assert.Equal(t, comment.UserId, createdComment.UserId)
		assert.Equal(t, comment.Content, createdComment.Content)
	})

	// test for update comment
	t.Run("TestUpdateComment", func(t *testing.T) {
		updateCommentRequest := &pbc.UpdateCommentRequest{
			CommentId: "550e8400-e29b-41d4-a716-446655440000",
		}

		updateCommentResponse, err := repo.UpdateComment(updateCommentRequest)

		assert.NoError(t, err)
		assert.NotNil(t, updateCommentResponse)
		assert.True(t, updateCommentResponse.Success)
	})

	// test for delete comment
	t.Run("TestDeleteComment", func(t *testing.T) {
		deleteCommentRequest := &pbc.DeleteCommentRequest{
			CommentId: "550e8400-e29b-41d4-a716-446655440000",
		}

		deleteCommentResponse, err := repo.DeleteComment(deleteCommentRequest)

		assert.NoError(t, err)
		assert.NotNil(t, deleteCommentResponse)
		assert.True(t, deleteCommentResponse.Success)
	})
}
