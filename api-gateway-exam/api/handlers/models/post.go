package models

type Post struct {
	PostId       string `json:"post_id"`
	UserId       string `json:"user_id"`
	Content      string `json:"content"`
	Title        string `json:"title"`
	CreatedAt    string `json:"created_at"`
	Likes        int64  `json:"likes"`
	Dislikes     int64  `json:"dislikes"`
	Views        int64  `json:"views"`
	MediaUrl     string `json:"media_url"`
	AccesToken   string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type GetPostByIdResponse struct {
	PostId       string     `json:"post_id"`
	UserId       string     `json:"user_id"`
	Content      string     `json:"content"`
	Title        string     `json:"title"`
	CreatedAt    string     `json:"created_at"`
	Likes        int64      `json:"likes"`
	Dislikes     int64      `json:"dislikes"`
	Views        int64      `json:"views"`
	MediaUrl     string     `json:"media_url"`
	RefreshToken string     `json:"refresh_token"`
	Comments     []*Comment `json:"comments"`
}

type PostsByUserId struct {
	Posts []*Post
}

type UpdateContentResponse struct {
	Success bool
}

type DeletePostResponse struct {
	Success bool
}
