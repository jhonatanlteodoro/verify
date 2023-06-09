package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jhonatanlteodoro/verify/app/models"
	"github.com/jhonatanlteodoro/verify/app/sqlite_connector"
)

func TestBasicBehaviorOfGetUserByIdHandler(t *testing.T) {
	wait := 1
	retry := 0
	db, err := sqlite_connector.GetConnection("test-local.sqlite", wait, retry)
	if err != nil {
		t.Error("fail connecting database")
	}
	models.RunMigrations(db)

	request, request_err := http.NewRequest(http.MethodGet, "/users/1", nil)
	if request_err != nil {
		t.Error(request_err)
	}

	responseRecorder := httptest.NewRecorder()
	router := mux.NewRouter()

	router.HandleFunc("/users/{id}", GetUserById(db))
	router.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected 200 as status code but got: %d", responseRecorder.Code)
	}

	contentType := responseRecorder.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected application/json content-type but got: %s", contentType)
	}
}

func TestBasicBehaviorOfGetAllHandler(t *testing.T) {
	wait := 1
	retry := 0
	db, err := sqlite_connector.GetConnection("test-local.sqlite", wait, retry)
	if err != nil {
		t.Error("fail connecting database")
	}
	models.RunMigrations(db)

	request, request_err := http.NewRequest(http.MethodGet, "/users/", nil)
	if request_err != nil {
		t.Error(request_err)
	}

	responseRecorder := httptest.NewRecorder()
	router := mux.NewRouter()

	router.HandleFunc("/users/", GetAll(db))
	router.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected 200 as status code but got: %d", responseRecorder.Code)
	}

	contentType := responseRecorder.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected application/json content-type but got: %s", contentType)
	}
}
