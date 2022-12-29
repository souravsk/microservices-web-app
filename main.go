package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	details "github.com/souravsk/microservices-web-app/details"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("serving the homepage")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is up and running")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking application health")
	response := map[string]string{
		"status":    "UP",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}

func detailHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Featching the details")
	hostname, err := details.GetHostname()
	if err != nil {
		panic(err)
	}
	IP, _ := details.GetIp()
	fmt.Println(IP)
	fmt.Print(hostname)

	response := map[string]string{
		"hostname": "hostname",
		"ip":       IP.String(),
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/detail", detailHandler)
	log.Println("server has started!!!")
	log.Fatal(http.ListenAndServe(":80", r))
}
