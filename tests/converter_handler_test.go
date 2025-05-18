package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mbrunoon/go-unit-converter/routes"
	"github.com/stretchr/testify/assert"
)

func TestHomeHandler(t *testing.T) {
	assert := assert.New(t)

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	router := routes.NewRouter()
	router.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(http.StatusOK, res.StatusCode, "should return 200")
}

func TestSuccessConverterHandler(t *testing.T) {
	assert := assert.New(t)

	body := map[string]interface{}{
		"value": 10.0,
		"from":  "kilometers",
		"to":    "meters",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/converter", bytes.NewBuffer(jsonBody))
	rec := httptest.NewRecorder()

	router := routes.NewRouter()
	router.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(res.StatusCode, http.StatusOK, "should return 200")

	var responseBody map[string]float64
	err := json.NewDecoder(res.Body).Decode(&responseBody)

	assert.NoError(err)
	assert.Equal(10.*1000, responseBody["value"])
}

func TestErrorInvalidsUnitsConverterHandler(t *testing.T) {
	assert := assert.New(t)

	body := map[string]interface{}{
		"value": 1.,
		"from":  "invalid",
		"to":    "invalid",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/converter", bytes.NewBuffer(jsonBody))
	rec := httptest.NewRecorder()

	router := routes.NewRouter()
	router.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(res.StatusCode, http.StatusBadRequest, "shoud return bad request")
}
