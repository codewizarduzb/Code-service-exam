package postgres

import (
	pbp "code-service-exam/post-service-exam/genproto/post-proto"
	"log"

	"github.com/jmoiron/sqlx"
)

type PostRepo struct {
	db *sqlx.DB
}

// NewpostRepo ...
func NewpostRepo(db *sqlx.DB) *PostRepo {
	return &PostRepo{db: db}
}

// create a new post to the store
func (r *PostRepo) CreatePost(post *pbp.Post) (*pbp.Post, error) {
	var newpost pbp.Post

	query := `INSERT INTO posts(post_id, user_id, content, title, likes, dislikes, views, media_url, refresh_token)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING post_id, user_id, content, title, created_at, likes, dislikes, views, media_url, refresh_token`

	err := r.db.QueryRow(
		query,
		post.PostId,
		post.UserId,
		post.Content,
		post.Title,
		post.Likes,
		post.Dislikes,
		post.Views,
		post.MediaUrl,
		post.RefreshToken,
	).Scan(
		&newpost.PostId,
		&newpost.UserId,
		&newpost.Content,
		&newpost.Title,
		&newpost.CreatedAt,
		&newpost.Likes,
		&newpost.Dislikes,
		&newpost.Views,
		&newpost.MediaUrl,
		&newpost.RefreshToken,
	)

	newpost.Comments = nil

	if err != nil {
		return nil, err
	}

	return &newpost, nil
}

// get a post by post_id from the store
func (r *PostRepo) GetPostById(post *pbp.GetPostByIdReq) (*pbp.GetPostByIdRes, error) {
	var getPost pbp.Post

	query := `SELECT post_id, user_id, content, title, created_at, likes, dislikes, views, media_url, refresh_token
	FROM posts WHERE post_id = $1`

	err := r.db.QueryRow(query, post.PostId).Scan(
		&getPost.PostId,
		&getPost.UserId,
		&getPost.Content,
		&getPost.Title,
		&getPost.CreatedAt,
		&getPost.Likes,
		&getPost.Dislikes,
		&getPost.Views,
		&getPost.MediaUrl,
		&getPost.RefreshToken,
	)
	if err != nil {
		return nil, err
	}

	query = `SELECT comment_id, post_id, user_id, content FROM comments WHERE post_id = $1`
	rows, err := r.db.Query(query, post.PostId)
	if err != nil {
		// if there are no comments, it returns response with nil comments
		return &pbp.GetPostByIdRes{
			Post: &getPost,
		}, nil
	}

	var comments []*pbp.Comment

	for rows.Next() {
		var comment pbp.Comment

		err := rows.Scan(&comment.CommentId, &comment.PostId, &comment.UserId, &comment.Content)
		if err != nil {
			// if there is an error in getting one comment, it sets this response to empty string
			comment.CommentId, comment.PostId, comment.UserId, comment.Content = "", "", "", ""
		}

		// append a comment to comments
		comments = append(comments, &comment)
	}

	// set response comments
	getPost.Comments = comments

	return &pbp.GetPostByIdRes{
		Post: &getPost,
	}, nil
}

// get posts by category from the store
func (r *PostRepo) GetPostsByUserId(req *pbp.GetPostsByUserIdReq) (*pbp.GetPostsByUserIdRes, error) {
	var posts []*pbp.Post

	query := `SELECT post_id, user_id, content, title, created_at, likes, dislikes, views, media_url, refresh_token
	FROM posts WHERE user_id = $1`

	rows, err := r.db.Query(query, req.UserId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post pbp.Post

		err := rows.Scan(
			&post.PostId,
			&post.UserId,
			&post.Content,
			&post.Title,
			&post.CreatedAt,
			&post.Likes,
			&post.Dislikes,
			&post.Views,
			&post.MediaUrl,
			&post.RefreshToken)
		if err != nil {
			log.Println(err)
		}

		query := `SELECT comment_id, post_id, user_id, content FROM comments WHERE post_id = $1`

		rows, err := r.db.Query(query, post.PostId)
		if err != nil {
			return nil, err
		}

		var comments []*pbp.Comment

		for rows.Next() {
			var comment pbp.Comment

			err := rows.Scan(&comment.CommentId, &comment.PostId, &comment.UserId, &comment.Content)
			if err != nil {
				comment.CommentId, comment.PostId, comment.UserId, comment.Content = "", "", "", ""
			}

			comments = append(comments, &comment)
		}

		// set comments to response post comments
		post.Comments = comments

		posts = append(posts, &post)
	}

	return &pbp.GetPostsByUserIdRes{
		Posts: posts,
	}, nil
}

// update a post by post_id
func (r *PostRepo) UpdatePost(req *pbp.UpdatePostReq) (*pbp.UpdatePostRes, error) {
	query := `UPDATE posts
        SET content = $1,
        updated_at = CURRENT_TIMESTAMP
        WHERE post_id = $2`

	_, err := r.db.Exec(query, req.Content, req.PostId)
	if err != nil {
		return &pbp.UpdatePostRes{
			Success: false,
		}, err
	}
	return &pbp.UpdatePostRes{
		Success: true,
	}, nil
}

// delete a post by post_id from the store
func (r *PostRepo) DeletePost(req *pbp.DeletePostReq) (*pbp.DeletePostRes, error) {
	query := `DELETE FROM posts WHERE post_id = $1`

	_, err := r.db.Exec(query, req.PostId)
	if err != nil {
		return &pbp.DeletePostRes{
			Success: false,
		}, err
	}
	return &pbp.DeletePostRes{
		Success: true,
	}, nil
}
