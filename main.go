package main

import (
	"encoding/json"
	"net/http"
	"github.com/bradfitz/gomemcache"
)

type Response struct {
	Message string `json:"message"`
}

var mc *memcache.Client

func init() {
	mc_host := os.Getenv("MEMCACHED_HOST")
	mc_port := os.Getenv("MEMCACHED_PORT")
	mc = memcache.New(mc_host+":"+mc_port)
	if mc == nil {
		log.Fatalf("Failed to connect to Memcached server")
	}

	defaultMessage := "Hello, World from Golang API!"
	if err := mc.Set(&memcache.Item{Key: "message", Value: []byte(defaultMessage)}); err != nil {
		log.Fatalf("Failed to set default message: %v", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	item, err := mc.Get("message")
	if err != nil {
		http.Error(w, "Failed to retrieve message from Memcached", http.StatusInternalServerError)
		return
	}
	response := Response{Message: string(item.Value)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/api/hello", helloHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
