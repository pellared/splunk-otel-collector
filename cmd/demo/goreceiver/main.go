package main

import (
	"encoding/json"
	"net/http"
)

type Service struct {
	Port    string `json:"port"`
	State   string `json:"state"`
	Service string `json:"service"`
	Version string `json:"version"`
}

func main() {
	services := []Service{
		{
			Port:    "3306/tcp",
			State:   "open",
			Service: "mysql",
			Version: "MySQL 5.5.64-MariaDB-1~trusty",
		},
		{
			Port:    "8080/tcp",
			State:   "open",
			Service: "http",
			Version: "Apache httpd 2.4.57 ((Debian))",
		},
		{
			Port:    "8081/tcp",
			State:   "open",
			Service: "http",
			Version: "Apache httpd 2.4.62 ((Debian))",
		},
	}

	http.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(services); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	port := "8090"
	http.ListenAndServe(":"+port, nil)
}
