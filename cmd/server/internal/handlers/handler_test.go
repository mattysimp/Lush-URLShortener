package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/mock"
)

// Mocks the database using mock library and map
type MockRepo struct {
	Mock mock.Mock
	urls map[string]*URL
}

func (mock *MockRepo) SetURL(setURL *URL) (err error) {
	if _, ok := mock.urls[setURL.URLCode]; ok {
		return errors.New("In DB")
	}
	mock.urls[setURL.URLCode] = setURL
	return nil
}

func (mock *MockRepo) GetURL(getURL string) (*URL, error) {
	if val, ok := mock.urls[getURL]; ok {
		return val, nil
	}
	return nil, errors.New("url not found")

}

func TestCreateShortURL(t *testing.T) {
	mockRepo := &MockRepo{mock.Mock{}, make(map[string]*URL)}
	config := &Config{BaseURL: "localhost:8080/"}

	// mockRepo.Mock.On("GetURL").Return(&URL{URLCode: "ytoij90", ShortURL: "localhost:8080/ytoij90", LongURL: "https://github.com/mattysimp/Lush-URLShortener"}, nil)

	mockRepo.Mock.On("GetURL").Return(nil, errors.New("url not found"))
	mockRepo.Mock.On("SetURL").Return(nil)

	Router := Routes(mockRepo, config)

	jsonStr := []byte(`{"url":"https://github.com/mattysimp/Lush-URLShortener"}`)
	response := request(t, "POST", "/", bytes.NewBuffer(jsonStr), Router, http.StatusOK)

	var u URL
	json.Unmarshal(response.Body.Bytes(), &u)

	if u.URLCode != "ytoij90" {
		t.Errorf("Expected id to be 'ytoij90'. Got %s", u.URLCode)
	}
	if u.ShortURL != "localhost:8080/ytoij90" {
		t.Errorf("Expected id to be 'localhost:8080/ytoij90'. Got %s", u.ShortURL)
	}
	if u.LongURL != "https://github.com/mattysimp/Lush-URLShortener" {
		t.Errorf("Expected id to be 'https://github.com/mattysimp/Lush-URLShortener'. Got %s", u.LongURL)
	}
}
func TestCreateBadURL(t *testing.T) {
	mockRepo := &MockRepo{mock.Mock{}, make(map[string]*URL)}
	config := &Config{BaseURL: "localhost:8080/"}

	Router := Routes(mockRepo, config)

	jsonStr := []byte(`{"url":"https/shortener"}`)
	request(t, "POST", "/", bytes.NewBuffer(jsonStr), Router, http.StatusBadRequest)
}

func TestCreateBadJSON(t *testing.T) {
	mockRepo := &MockRepo{mock.Mock{}, make(map[string]*URL)}
	config := &Config{BaseURL: "localhost:8080/"}

	Router := Routes(mockRepo, config)

	jsonStr := []byte(`{"pie":"meat and potato"}`)
	request(t, "POST", "/", bytes.NewBuffer(jsonStr), Router, http.StatusBadRequest)
}

func TestCreateNoJSON(t *testing.T) {
	mockRepo := &MockRepo{mock.Mock{}, make(map[string]*URL)}
	config := &Config{BaseURL: "localhost:8080/"}

	Router := Routes(mockRepo, config)

	request(t, "POST", "/", nil, Router, http.StatusBadRequest)
}
func TestCreateExistingURL(t *testing.T) {
	mockRepo := &MockRepo{mock.Mock{}, make(map[string]*URL)}
	config := &Config{BaseURL: "localhost:8080/"}

	existingURL := &URL{URLCode: "ytoij90", ShortURL: "localhost:8080/ytoij90", LongURL: "https://github.com/mattysimp/Lush-URLShortener"}
	mockRepo.urls["ytoij90"] = existingURL

	mockRepo.Mock.On("GetURL").Return(existingURL, nil)

	Router := Routes(mockRepo, config)

	jsonStr := []byte(`{"url":"https://github.com/mattysimp/Lush-URLShortener"}`)
	response := request(t, "POST", "/", bytes.NewBuffer(jsonStr), Router, http.StatusOK)

	var u URL
	json.Unmarshal(response.Body.Bytes(), &u)

	if u.URLCode != "ytoij90" {
		t.Errorf("Expected id to be 'ytoij90'. Got %s", u.URLCode)
	}
	if u.ShortURL != "localhost:8080/ytoij90" {
		t.Errorf("Expected id to be 'localhost:8080/ytoij90'. Got %s", u.ShortURL)
	}
	if u.LongURL != "https://github.com/mattysimp/Lush-URLShortener" {
		t.Errorf("Expected id to be 'https://github.com/mattysimp/Lush-URLShortener'. Got %s", u.LongURL)
	}
}

func TestCreateHashConflict(t *testing.T) {
	mockRepo := &MockRepo{mock.Mock{}, make(map[string]*URL)}
	config := &Config{BaseURL: "localhost:8080/"}

	existingURL := &URL{URLCode: "ytoij90", ShortURL: "localhost:8080/ytoij90", LongURL: "https://google.com"}
	mockRepo.urls["ytoij90"] = existingURL

	mockRepo.Mock.On("GetURL").Return(existingURL, nil)

	Router := Routes(mockRepo, config)

	jsonStr := []byte(`{"url":"https://github.com/mattysimp/Lush-URLShortener"}`)
	response := request(t, "POST", "/", bytes.NewBuffer(jsonStr), Router, http.StatusOK)

	var u URL
	json.Unmarshal(response.Body.Bytes(), &u)

	if u.URLCode != "ytoij91" {
		t.Errorf("Expected id to be 'ytoij91'. Got %s", u.URLCode)
	}
	if u.ShortURL != "localhost:8080/ytoij91" {
		t.Errorf("Expected id to be 'localhost:8080/ytoij91'. Got %s", u.ShortURL)
	}
	if u.LongURL != "https://github.com/mattysimp/Lush-URLShortener" {
		t.Errorf("Expected id to be 'https://github.com/mattysimp/Lush-URLShortener'. Got %s", u.LongURL)
	}
}

func TestGetURL(t *testing.T) {
	mockRepo := &MockRepo{mock.Mock{}, make(map[string]*URL)}
	config := &Config{BaseURL: "localhost:8080/"}

	existingURL := &URL{URLCode: "ytoij90", ShortURL: "localhost:8080/ytoij90", LongURL: "https://github.com/mattysimp/Lush-URLShortener"}
	mockRepo.urls["ytoij90"] = existingURL

	mockRepo.Mock.On("GetURL").Return(existingURL, nil)

	Router := Routes(mockRepo, config)

	response := request(t, "GET", "/ytoij90", nil, Router, http.StatusOK)

	var u URL
	json.Unmarshal(response.Body.Bytes(), &u)

	if u.URLCode != "ytoij90" {
		t.Errorf("Expected id to be 'ytoij90'. Got %s", u.URLCode)
	}
	if u.ShortURL != "localhost:8080/ytoij90" {
		t.Errorf("Expected id to be 'localhost:8080/ytoij91'. Got %s", u.ShortURL)
	}
	if u.LongURL != "https://github.com/mattysimp/Lush-URLShortener" {
		t.Errorf("Expected id to be 'https://github.com/mattysimp/Lush-URLShortener'. Got %s", u.LongURL)
	}
}

func TestGetBlankURL(t *testing.T) {
	mockRepo := &MockRepo{mock.Mock{}, make(map[string]*URL)}
	config := &Config{BaseURL: "localhost:8080/"}

	Router := Routes(mockRepo, config)

	request(t, "GET", "/", nil, Router, http.StatusMethodNotAllowed)
}

func TestGetNonExistingURL(t *testing.T) {
	mockRepo := &MockRepo{mock.Mock{}, make(map[string]*URL)}
	config := &Config{BaseURL: "localhost:8080/"}

	Router := Routes(mockRepo, config)

	request(t, "GET", "/y2k", nil, Router, http.StatusBadRequest)
}

func request(t *testing.T, rest string, url string, contents io.Reader, router *chi.Mux, expectedResponseCode int) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(rest, url, contents)
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req, router)
	checkResponseCode(t, expectedResponseCode, response.Code)
	return response
}

func executeRequest(req *http.Request, Router *chi.Mux) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	Router.ServeHTTP(responseRecorder, req)
	return responseRecorder
}

func checkResponseCode(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d", expected, actual)
	}
}
