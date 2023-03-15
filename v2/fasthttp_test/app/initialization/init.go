// Package responsible to manage the app initialization.
// Only init, in you want to add more things you must create an abstract method
// in order to only call the method here and nothing more.
package initalization

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/jhonatanlteodoro/fasthttp_test/app/config"
	"github.com/jhonatanlteodoro/fasthttp_test/app/handlers"
	"github.com/jhonatanlteodoro/fasthttp_test/app/models"
	"github.com/jhonatanlteodoro/fasthttp_test/app/routes"
	"github.com/valyala/fasthttp"

	mysql_connector "github.com/jhonatanlteodoro/fasthttp_test/app/db_connectors/mysql"
	sqlite_connector "github.com/jhonatanlteodoro/fasthttp_test/app/db_connectors/sqlite"

	routing "github.com/qiangxue/fasthttp-routing"
	"gorm.io/gorm"
)

type App struct {
	cfg *config.Config

	DBConn *gorm.DB
	Router *routing.Router
}

func (a *App) loadURI() string {
	mysqlUri := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		a.cfg.MYSQL_USERNAME, a.cfg.MYSQL_PASSWORD,
		a.cfg.MYSQL_HOST, a.cfg.MYSQL_HOST_PORT,
		a.cfg.MYSQL_DB_NAME,
	)
	return mysqlUri
}

func (a *App) LoadSqliteFilePath() string {
	sqliteFile, err := filepath.Abs(fmt.Sprintf("./%s", a.cfg.SQLITE_FILENAME))
	if err != nil {
		log.Println("Fail while load sqlite path")
		log.Fatal(err)
	}
	return sqliteFile
}

func (a *App) InitilizeDB() {
	waitSecondsCaseError := 5
	retry_case_error := 5

	var conn *gorm.DB
	var err error

	if a.cfg.DB_DRIVER == "mysql" {
		uri := a.loadURI()
		conn = mysql_connector.GetConnection(uri, waitSecondsCaseError, retry_case_error)
		log.Println("using mysql database")
	} else {
		sqliteFile := a.LoadSqliteFilePath()
		conn = sqlite_connector.GetConnection(sqliteFile, waitSecondsCaseError, retry_case_error)
		log.Println("using sqlite database")
	}

	if err != nil {
		log.Fatal(err)
	}

	a.DBConn = conn
}

func (a *App) MakeMigrations() {
	log.Println("running migrations...")
	models.RunMigrations(a.DBConn)
	log.Println("migrations completed!")
}

func (a *App) InitilizeRoutes() {
	a.Router = routing.New()
	routes.RegistryRoutes(a.Router)
	log.Println("routers registred!")
}

func (a *App) InitializeHandlers() {
	handlers_repo := handlers.NewTestRepo(a.cfg, a.DBConn)
	handlers.SetHandlerRepository(handlers_repo)
}

func (a *App) Initilize() {
	a.cfg = config.InitConfig()
	a.InitilizeDB()
	a.MakeMigrations()
	a.InitializeHandlers()
	a.InitilizeRoutes()
}

func (a *App) Run(host string, port string) {
	runningAt := fmt.Sprintf("%s:%s", host, port)
	log.Printf("Running server... %s", runningAt)
	log.Fatal(
		fasthttp.ListenAndServe(runningAt, a.Router.HandleRequest),
	)
}
