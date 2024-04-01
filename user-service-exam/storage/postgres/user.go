package postgres

import (
	pb "code-service-exam/user-service-exam/genproto/user-proto"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

// NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

// create a new user
func (r *UserRepo) CreateUser(user *pb.User) (*pb.User, error) {
	var newUser pb.User

	query := `INSERT INTO users(id, username, email, password, first_name, last_name, bio, website)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, username, email, password, created_at, first_name, last_name, bio, website`

	err := r.db.QueryRow(
		query,
		user.Id,
		user.Username,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.Bio,
		user.Website,
	).Scan(
		&newUser.Id,
		&newUser.Username,
		&newUser.Email,
		&newUser.Password,
		&newUser.CreatedAt,
		&newUser.FirstName,
		&newUser.LastName,
		&newUser.Bio,
		&newUser.Website,
	)

	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

// get a user by email
func (r *UserRepo) GetUser(req *pb.GetUserReq) (*pb.User, error) {
	var getUser pb.User

	query := `SELECT id, username, email, password, created_at, updated_at, first_name, last_name, bio, website
	FROM users WHERE email = $1`

	err := r.db.QueryRow(query, req.Email).Scan(
		&getUser.Id,
		&getUser.Username,
		&getUser.Email,
		&getUser.Password,
		&getUser.CreatedAt,
		&getUser.UpdatedAt,
		&getUser.FirstName,
		&getUser.LastName,
		&getUser.Bio,
		&getUser.Website,
	)

	if err != nil {
		return nil, err
	}

	return &getUser, nil
}

// get a list of users according to page and limit
func (r *UserRepo) ListUsers(req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	query := `SELECT id, username, email, password, created_at, updated_at, first_name, last_name, bio, website
	FROM users LIMIT $1 OFFSET $2;`

	offset := (req.Page - 1) * req.Limit

	rows, err := r.db.Query(query, req.Limit, offset)
	if err != nil {
		return nil, err
	}

	var users []*pb.User

	for rows.Next() {
		var user pb.User

		if err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.FirstName,
			&user.LastName,
			&user.Bio,
			&user.Website,
		); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return &pb.ListUsersResponse{
		Users: users,
	}, nil
}

// update user username by email
func (r *UserRepo) UpdateUser(req *pb.UpdateUserReq) (*pb.UpdateUserRes, error) {
	query := `UPDATE users
        SET username = 'striker', updated_at = CURRENT_TIMESTAMP
        WHERE email = $1`

	_, err := r.db.Exec(query, req.Email)
	if err != nil {
		return &pb.UpdateUserRes{
			Success: false,
		}, err
	}
	return &pb.UpdateUserRes{
		Success: true,
	}, nil
}

// delete a user by email
func (r *UserRepo) DeleteUser(req *pb.DeleteUserReq) (*pb.DeleteUserRes, error) {
	query := `DELETE FROM users WHERE email = $1`

	_, err := r.db.Exec(query, req.Email)
	if err != nil {
		return &pb.DeleteUserRes{
			Success: false,
		}, err
	}
	return &pb.DeleteUserRes{
		Success: true,
	}, nil
}
