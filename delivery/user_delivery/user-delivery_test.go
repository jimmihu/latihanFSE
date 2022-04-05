package user_delivery

import (
	"bytes"
	"encoding/json"
	"io"
	"latihanFSE/models/dto"
	"latihanFSE/models/entity"
	"latihanFSE/usecase/user_usecase"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func prepareGinAndRecorder() (*httptest.ResponseRecorder, *gin.Context) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}

	return w, c
}

func MockJsonBodyRequest(t *testing.T, c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	var jsonbytes []byte
	var err error

	if content != nil {
		c.Request.Header.Set("Content-Type", "application/json")
		jsonbytes, err = json.Marshal(content)
		if err != nil {
			t.Fatal("Error marshalling json body")
		}
	} else {
		c.Request.Header.Set("Content-Type", "none")
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

}

func getDelivery() UserDeliveryInterface {
	usecaseMock := new(user_usecase.UserUsecaseMock)
	delivery := CreateUserDelivery(usecaseMock)
	return delivery
}

func getErrorResults(w *httptest.ResponseRecorder) (*http.Response, dto.HttpResponse) {
	result := w.Result()
	body, _ := io.ReadAll(result.Body)

	var bodyResp dto.HttpResponse

	json.Unmarshal(body, &bodyResp)

	return result, bodyResp
}

var reqUser = entity.User{}

func TestGivenEmptyJSON_WhenCreateUser_ShouldProduceBadRequestError(t *testing.T) {

	w, c := prepareGinAndRecorder()
	MockJsonBodyRequest(t, c, reqUser)
	delivery := getDelivery()
	delivery.CreateUser(c)
	result, bodyResp := getErrorResults(w)
	assert.EqualValues(t, http.StatusBadRequest, result.StatusCode)
	assert.EqualValues(t, http.StatusBadRequest, bodyResp.StatusCode)
}

func TestGivenNilBody_WhenCreateUser_ShouldProduceBadRequestError(t *testing.T) {

	w, c := prepareGinAndRecorder()
	MockJsonBodyRequest(t, c, nil)
	delivery := getDelivery()
	delivery.CreateUser(c)
	result, bodyResp := getErrorResults(w)
	assert.EqualValues(t, http.StatusBadRequest, result.StatusCode)
	assert.EqualValues(t, http.StatusBadRequest, bodyResp.StatusCode)
}

func TestGivenEmptyJSON_WhenLoginUser_ShouldProduceBadRequestError(t *testing.T) {

	w, c := prepareGinAndRecorder()
	MockJsonBodyRequest(t, c, reqUser)
	delivery := getDelivery()
	delivery.LoginUser(c)
	result, bodyResp := getErrorResults(w)

	assert.EqualValues(t, http.StatusBadRequest, result.StatusCode)
	assert.EqualValues(t, http.StatusBadRequest, bodyResp.StatusCode)
}

func TestGivenNilBody_WhenLoginUser_ShouldProduceBadRequestError(t *testing.T) {

	w, c := prepareGinAndRecorder()
	MockJsonBodyRequest(t, c, nil)
	delivery := getDelivery()
	delivery.LoginUser(c)
	result, bodyResp := getErrorResults(w)
	assert.EqualValues(t, http.StatusBadRequest, result.StatusCode)
	assert.EqualValues(t, http.StatusBadRequest, bodyResp.StatusCode)
}
