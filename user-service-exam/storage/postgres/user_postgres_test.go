package postgres

import (
	"testing"

	"code-service-exam/user-service-exam/config"
	pbu "code-service-exam/user-service-exam/genproto/user-proto"
	"code-service-exam/user-service-exam/pkg/db"

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

	repo := NewUserRepo(db)
	user := &pbu.User{
		Id:        uuid.NewString(),
		Username:  "test username",
		Email:     "test email",
		Password:  "test password",
		FirstName: "test first_name",
		LastName:  "test last_name",
		Bio:       "test bio",
		Website:   "test website",
	}
	// test for create user
	createdUser, err := repo.CreateUser(&pbu.User{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Bio:       user.Bio,
		Website:   user.Website,
	})
	assert.NoError(t, err)
	assert.Equal(t, user, createdUser)

	// test for update user
	updatedUser, err := repo.UpdateUser(&pbu.UpdateUserReq{
		Email: user.Email,
	})
	assert.NoError(t, err)
	assert.Equal(t, true, updatedUser.Success)

	// test for delete user
	deletedUser, err := repo.DeleteUser(&pbu.DeleteUserReq{
		Email: user.Email,
	})
	assert.NoError(t, err)
	assert.Equal(t, true, deletedUser.Success)
}
