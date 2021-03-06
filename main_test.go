package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeRoute(t *testing.T) {
	t.Run("Return home route message", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		Home(response, request)

		got := response.Body.String()
		want := "EduNet Foundation"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
