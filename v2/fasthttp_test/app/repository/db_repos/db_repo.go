package db_repos

import (
	"github.com/jhonatanlteodoro/fasthttp_test/app/repository"
	"gorm.io/gorm"
)

type mysqlDBRepo struct {
	// App interface{}
	DB *gorm.DB
}

type testDBRepo struct {
	// App interface{}
	DB *gorm.DB
}

func NewMysqlRepo(conn *gorm.DB) repository.DatabaseRepo {
	return &mysqlDBRepo{
		// App: app,
		DB: conn,
	}
}

func NewTestingRepo(conn *gorm.DB) repository.DatabaseRepo {
	return &testDBRepo{
		DB: conn,
	}
}
