package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCalculateAverage_1(t *testing.T) {
	// Setup
	e := echo.New()

	q := make(url.Values)
	q.Set("mark_1", "12")
	q.Set("mark_2", "10")
	q.Set("mark_3", "5")
	q.Set("mark_4", "20")
	q.Set("mark_5", "78")

	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	desiredResponse := `{"error":false,"average":25,"string":"Your average mark is 25"}`

	// Assertions
	if assert.NoError(t, calculateAverage(c)) {
		actualResponse := strings.Trim(rec.Body.String(), "\n")

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, desiredResponse, actualResponse)
	}
}

func TestCalculateAverage_2(t *testing.T) {
	// Setup
	e := echo.New()

	q := make(url.Values)
	q.Set("mark_1", "12")
	q.Set("mark_2", "10")
	q.Set("mark_4", "20")
	q.Set("mark_5", "78")

	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	desiredResponse := `{"error":true,"average":null,"string":"Mark 3 value is missing"}`

	// Assertions
	if assert.NoError(t, calculateAverage(c)) {
		actualResponse := strings.Trim(rec.Body.String(), "\n")

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, desiredResponse, actualResponse)
	}
}

func TestCalculateAverage_3(t *testing.T) {
	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, calculateAverage(c)) {
		actualResponse := strings.Trim(rec.Body.String(), "\n")
		containsError := strings.Contains(actualResponse, "value is missing")

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.True(t, containsError)
	}
}

func TestCalculateAverage_4(t *testing.T) {
	// Setup
	e := echo.New()

	q := make(url.Values)
	q.Set("mark_1", "12")
	q.Set("mark_2", "asdasdasda")
	q.Set("mark_3", "5")
	q.Set("mark_4", "20")
	q.Set("mark_5", "78")

	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	desiredResponse := `{"error":true,"average":null,"string":"You must provide a valid integer for Mark 2"}`

	// Assertions
	if assert.NoError(t, calculateAverage(c)) {
		actualResponse := strings.Trim(rec.Body.String(), "\n")

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, desiredResponse, actualResponse)
	}
}

func TestCalculateAverage_5(t *testing.T) {
	// Setup
	e := echo.New()

	q := make(url.Values)
	q.Set("mark_1", "12")
	q.Set("mark_2", "10")
	q.Set("mark_3", "5")
	q.Set("mark_4", "20")
	q.Set("mark_5", "-78")

	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	desiredResponse := `{"error":true,"average":null,"string":"You must provide a non-negative integer for Mark 5"}`

	// Assertions
	if assert.NoError(t, calculateAverage(c)) {
		actualResponse := strings.Trim(rec.Body.String(), "\n")

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, desiredResponse, actualResponse)
	}
}

func TestCalculateAverage_6(t *testing.T) {
	// Setup
	e := echo.New()

	q := make(url.Values)
	q.Set("mark_1", "12")
	q.Set("mark_2", "10")
	q.Set("mark_3", "520")
	q.Set("mark_4", "20")
	q.Set("mark_5", "78")

	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	desiredResponse := `{"error":true,"average":null,"string":"You cannot exceed 100 marks for Mark 3"}`

	// Assertions
	if assert.NoError(t, calculateAverage(c)) {
		actualResponse := strings.Trim(rec.Body.String(), "\n")

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, desiredResponse, actualResponse)
	}
}
