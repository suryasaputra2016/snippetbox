package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/suryasaputra2016/snippetbox/internal/assert"
)

func TestCommonHeader(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	commonHeaders(next).ServeHTTP(rr, r)

	rs := rr.Result()

	tests := []struct {
		name      string
		headerKey string
		want      string
	}{
		{
			name:      "Content-Security-Policy",
			headerKey: "Content-Security-Policy",
			want:      "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com",
		},
		{
			name:      "Referer-Policy",
			headerKey: "Referer-Policy",
			want:      "origin-when-cross-origin",
		},
		{
			name:      "X-Content-Type-Options",
			headerKey: "X-Content-Type-Options",
			want:      "nosniff",
		},
		{
			name:      "X-Frame-Options",
			headerKey: "X-Frame-Options",
			want:      "deny",
		},
		{
			name:      "X-XSS-Protection",
			headerKey: "X-XSS-Protection",
			want:      "0",
		},
		{
			name:      "Server",
			headerKey: "Server",
			want:      "Go",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, rs.Header.Get(tt.headerKey), tt.want)
		})
	}

	assert.Equal(t, rs.StatusCode, http.StatusOK)

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)
	assert.Equal(t, string(body), "OK")
}
