package storage

import (
	"code-service-exam/post-service-exam/storage/postgres"
	"code-service-exam/post-service-exam/storage/repo"

	"github.com/jmoiron/sqlx"
)

// IStorage ...
type IStorage interface {
	Post() repo.PostStorageI
}

type Pg struct {
	db       *sqlx.DB
	postRepo repo.PostStorageI
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *Pg {
	return &Pg{
		db:       db,
		postRepo: postgres.NewpostRepo(db),
	}
}

func (s Pg) Post() repo.PostStorageI {
	return s.postRepo
}
