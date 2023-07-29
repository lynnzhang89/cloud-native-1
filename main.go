package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handle(writer http.ResponseWriter, request *http.Request) {
	headers := request.Header
	for key, value := range headers {
		writer.Header().Set(key, value[0])
	}
	version := os.Getenv("VERSION")
	log.Printf("version: %s", version)
	if version != "" {
		writer.Header().Set("VERSION", version)
	}

	log.Printf("client ip: %s, status code: %d", request.RemoteAddr, http.StatusOK)
}

func healthz(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("http start failed err:", err)
		return
	}
}
