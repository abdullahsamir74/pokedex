package pokeapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/abdullahsamir74/pokedex/internal/pokecache"
)

func TestGetCachesResponses(t *testing.T) {
	var calls int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"name":"pikachu","base_experience":112}`))
	}))
	defer server.Close()

	client := &Client{
		cache:      pokecache.NewCache(time.Hour),
		httpClient: server.Client(),
	}

	var first Pokemon
	if err := client.get(server.URL, &first); err != nil {
		t.Fatalf("first get: %v", err)
	}
	if first.Name != "pikachu" || first.BaseExperience != 112 {
		t.Fatalf("decoded unexpected pokemon: %+v", first)
	}

	var second Pokemon
	if err := client.get(server.URL, &second); err != nil {
		t.Fatalf("second get: %v", err)
	}
	if second.Name != "pikachu" {
		t.Fatalf("cached response decoded unexpected pokemon: %+v", second)
	}

	if calls != 1 {
		t.Fatalf("expected 1 HTTP request, got %d", calls)
	}
}

func TestGetReturnsErrorOnNonOKStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)
	}))
	defer server.Close()

	client := &Client{
		cache:      pokecache.NewCache(time.Hour),
		httpClient: server.Client(),
	}

	var target Pokemon
	if err := client.get(server.URL, &target); err == nil {
		t.Fatal("expected an error for a 404 response")
	}
}

func TestGetReturnsErrorOnInvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`not json`))
	}))
	defer server.Close()

	client := &Client{
		cache:      pokecache.NewCache(time.Hour),
		httpClient: server.Client(),
	}

	var target Pokemon
	if err := client.get(server.URL, &target); err == nil {
		t.Fatal("expected an error for an invalid JSON body")
	}
}
