package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/smonkeymonkey/marathon_runners/models"
	"github.com/smonkeymonkey/marathon_runners/repositories"
	"github.com/smonkeymonkey/marathon_runners/services"
	"github.com/stretchr/testify/assert"
)

func TestGetRunnersResponse(t *testing.T) {
	dbHandler, mock, _ := sqlmock.New()
	defer dbHandler.Close()

	columnsUsers := []string{"user_role"}
	mock.ExpectQuery("SELECT user_role").WillReturnRows(
		sqlmock.NewRows(columnsUsers).AddRow("runner"),
	)

	columns := []string{"id", "first_name", "last_name", "age", "is_active", "country", "personal_best", "season_best"}

	mock.ExpectQuery("SELECT *").WillReturnRows(
		sqlmock.NewRows(columns).AddRow("1", "John", "Smith", 30, true, "US", "02:00:41", "2:19:32").
			AddRow("2", "Elena", "Smith", "30", true, "UK", "01:58:99", "01:58:99"))

	router := initTestRouter(dbHandler)
	request, _ := http.NewRequest("GET", "/runner", nil)
	request.Header.Set("token", "token")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)

	var runners []*models.Runner
	json.Unmarshal(recorder.Body.Bytes(), &runners)
	assert.NotEmpty(t, runners)
	assert.Equal(t, 2, len(runners))
}

func initTestRouter(dbHandler *sql.DB) *gin.Engine {
	runnersRepository := repositories.NewRunnersRepository(dbHandler)
	usersRepository := repositories.NewUsersRepository(dbHandler)

	runnersService := services.NewRunnersService(runnersRepository, nil)
	usersService := services.NewUsersService(usersRepository)

	runnersController := NewRunnersController(runnersService, usersService)

	router := gin.Default()
	router.GET("/runner", runnersController.GetRunnersBatch)
	return router
}
