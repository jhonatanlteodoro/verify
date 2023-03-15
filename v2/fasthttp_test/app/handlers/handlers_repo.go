package handlers

import (
	"github.com/jhonatanlteodoro/fasthttp_test/app/config"
	"github.com/jhonatanlteodoro/fasthttp_test/app/repository"
	"github.com/jhonatanlteodoro/fasthttp_test/app/repository/db_repos"
	"gorm.io/gorm"
)

var HanlderRepo *Repository

type Repository struct {
	App *config.Config
	DB  repository.DatabaseRepo
}

// creates a new repository
func NewRepo(a *config.Config, db *gorm.DB) *Repository {
	return &Repository{
		App: a,
		DB:  db_repos.NewMysqlRepo(db),
	}
}

// creates a new test repository
func NewTestRepo(a *config.Config, db *gorm.DB) *Repository {
	return &Repository{
		App: a,
		DB:  db_repos.NewTestingRepo(db),
	}
}

// sets the repository for the handlers
func SetHandlerRepository(r *Repository) {
	HanlderRepo = r
}
