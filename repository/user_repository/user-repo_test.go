package user_repository

import (
	"database/sql"
	"net/http"
	"testing"

	"latihanFSE/config"
	"latihanFSE/models/dto"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

var repo UserRepoInterface
var mockSql sqlmock.Sqlmock
var db *sql.DB
var querySelect = "^SELECT (.+)"

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	var err error
	db, mockSql, err = sqlmock.New()
	if err != nil {
		panic("error sql mock")
	}
	defer db.Close()

	connectionString := config.CONFIG["MYSQL_USER"] + ":" + config.CONFIG["MYSQL_PASS"] + "@tcp(" + config.CONFIG["MYSQL_HOST"] + ":" + config.CONFIG["MYSQL_PORT"] + ")/" + config.CONFIG["MYSQL_SCHEMA"] + "?parseTime=true"
	dbMy, errMy := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if errMy != nil {
		panic("error gorm")
	}
	repo = GetUserRepo(dbMy, dbMy)

	m.Run()
}

func TestGetUserDetailDataNotFound(t *testing.T) {
	rows := sqlmock.NewRows([]string{"id", "personal_number", "password", "email", "name", "role_id", "active"})
	uuID, _ := uuid.Parse("falseid")
	mockSql.ExpectQuery(querySelect).WithArgs(uuID).WillReturnRows(rows)

	_, result := repo.GetUserDetail(uuID)
	respons := dto.HttpResponse{}
	if result.Error != nil {
		respons = dto.DBErrorResponse(result.Error)
	} else {
		respons = dto.HttpResponse{
			StatusCode: http.StatusOK,
			Status:     "ok",
			Error:      nil,
			Data:       nil,
		}
	}

	assert.NotNil(t, result.Error, "error should not be nil")
	assert.EqualValues(t, http.StatusInternalServerError, respons.StatusCode)
}
