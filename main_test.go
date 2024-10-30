package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMainHandlerStatusOk Проверка корректности сформированного запроса
func TestMainHandlerStatusOk(t *testing.T) {
	req := httptest.NewRequest("GET", fmt.Sprintf("/cafe?count=%d&city=moscow", 4), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body)
}

// TestMainHandlerInvalidNameCity Проверка поддержки сервером указанного города
func TestMainHandlerInvalidNameCity(t *testing.T) {
	req := httptest.NewRequest("GET", fmt.Sprintf("/cafe?count=%d&city=ustyg", 4), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	expected := "wrong city value"
	require.Equal(t, expected, responseRecorder.Body.String())
}

// TestMainHandlerTotal Проверка корректного вывода, когда параметр count > есть всего
func TestMainHandlerTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", fmt.Sprintf("/cafe?count=%d&city=moscow", 8), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Len(t, list, totalCount)
}
