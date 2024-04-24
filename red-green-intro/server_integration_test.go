package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T)  {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "padul"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	res := httptest.NewRecorder()
	server.ServeHTTP(res, newGetScoreRequest(player))

	assertResponseStatus(t, res.Code, http.StatusOK)
	assertResponseBody(t, res.Body.String(), "3")
}
