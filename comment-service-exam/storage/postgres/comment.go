package postgres

import (
	pbc "code-service-exam/comment-service-exam/genproto/comment-proto"
	"log"

	"github.com/jmoiron/sqlx"
)

type CommentRepo struct {
	db *sqlx.DB
}

// NewCommentRepo ...
func NewCommentRepo(db *sqlx.DB) *CommentRepo {
	return &CommentRepo{db: db}
}

// create a new comment
func (r *CommentRepo) CreateComment(comment *pbc.Comment) (*pbc.Comment, error) {
	var newComment pbc.Comment

	query := `INSERT INTO comments(comment_id, post_id, user_id, content)
	VALUES($1, $2, $3, $4)
	RETURNING comment_id, post_id, user_id, content, created_at`

	if err := r.db.QueryRow(
		query,
		comment.CommentId,
		comment.PostId,
		comment.UserId,
		comment.Content,
	).Scan(
		&newComment.CommentId,
		&newComment.PostId,
		&newComment.UserId,
		&newComment.Content,
		&newComment.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &newComment, nil
}

// get a comment by comment_id
func (r *CommentRepo) GetComment(comment *pbc.GetCommentRequest) (*pbc.GetCommentResponse, error) {
	var getComment pbc.Comment

	query := `SELECT comment_id, post_id, user_id, content, created_at
	FROM comments WHERE comment_id = $1`

	if err := r.db.QueryRow(query, comment.CommentId).Scan(
		&getComment.CommentId,
		&getComment.PostId,
		&getComment.UserId,
		&getComment.Content,
		&getComment.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &pbc.GetCommentResponse{
		Comment: &getComment,
	}, nil
}

// update a comment content by comment_id
func (r *CommentRepo) UpdateComment(req *pbc.UpdateCommentRequest) (*pbc.UpdateCommentResponse, error) {
	query := `UPDATE comments
        SET content = 'updated comment',
		updated_at = CURRENT_TIMESTAMP
        WHERE comment_id = $1`

	_, err := r.db.Exec(query, req.CommentId)
	if err != nil {
		return &pbc.UpdateCommentResponse{
			Success: false,
		}, err
	}
	return &pbc.UpdateCommentResponse{
		Success: true,
	}, nil
}

// delete a comment by comment_id
func (r *CommentRepo) DeleteComment(req *pbc.DeleteCommentRequest) (*pbc.DeleteCommentResponse, error) {
	query := `DELETE FROM comments WHERE comment_id = $1`

	_, err := r.db.Exec(query, req.CommentId)
	if err != nil {
		return &pbc.DeleteCommentResponse{
			Success: false,
		}, err
	}
	return &pbc.DeleteCommentResponse{
		Success: true,
	}, nil
}

// get a list of comments by post_id
func (r *CommentRepo) ListComments(req *pbc.ListCommentsRequest) (*pbc.ListCommentsResponse, error) {
	var comments []*pbc.Comment

	query := `SELECT comment_id, post_id, user_id, content, created_at
	FROM comments WHERE post_id = $1`

	rows, err := r.db.Query(query, req.PostId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var comment pbc.Comment

		if err := rows.Scan(
			&comment.CommentId,
			&comment.PostId,
			&comment.UserId,
			&comment.Content,
			&comment.CreatedAt,
		); err != nil {
			log.Println(err)
		}

		comments = append(comments, &comment)
	}

	return &pbc.ListCommentsResponse{
		Comments: comments,
	}, nil
}
