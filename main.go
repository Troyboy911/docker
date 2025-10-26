package main

import (
	"encoding/json"
	"net/http"
	"os"
	"runtime"
	"time"
)

type Health struct {
	Status    string `json:"status"`
	Time      string `json:"time"`
	GoVersion string `json:"go_version"`
	Hostname  string `json:"hostname"`
	UptimeSec int64  `json:"uptime_seconds"`
}

var start = time.Now()

func healthHandler(w http.ResponseWriter, r *http.Request) {
	h := Health{
		Status:    "ok",
		Time:      time.Now().UTC().Format(time.RFC3339),
		GoVersion: runtime.Version(),
		Hostname:  getHostname(),
		UptimeSec: int64(time.Since(start).Seconds()),
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(h)
}

func getHostname() string {
	if hn, err := os.Hostname(); err == nil {
		return hn
	}
	return "unknown"
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":8080", nil)
}
