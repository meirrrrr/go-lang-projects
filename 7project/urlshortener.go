package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var urlMappings = make(map[string]string)

const baseURL = "http://localhost:8080/"

func isValidURL(input string) bool {
	_, err := url.ParseRequestURI(input)
	return err == nil
}

func generateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	shortKey := make([]byte, 5)
	for i := 0; i < 5; i++ {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}

func shortenURL(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		LongURL string `json:"long_url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	longURL := requestData.LongURL

	if !isValidURL(longURL) {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	shortKey := generateShortKey()
	urlMappings[shortKey] = longURL

	shortURL := baseURL + shortKey
	response := struct {
		ShortURL string `json:"short_url"`
	}{
		ShortURL: shortURL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectToURL(w http.ResponseWriter, r *http.Request) {
	shortKey := strings.TrimPrefix(r.URL.Path, "/")

	longURL, exists := urlMappings[shortKey]
	if !exists {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

func main() {
	http.HandleFunc("/shorten", shortenURL)        // POST /shorten
	http.HandleFunc("/", redirectToURL)           // GET /{short_key}

	fmt.Println("URL shortener server running on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
