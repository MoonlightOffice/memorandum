package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	Server()
}

func Server() {
	http.ListenAndServe("0.0.0.0:8888", Middleware(Router()))
}

func Router() http.Handler {
	server := http.NewServeMux()

	server.HandleFunc("/health", HealthHandler)
	server.HandleFunc("/sse", SSEHandler)

	return server
}

func SSEHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Accel-Buffering", "no")
	w.Header().Add("Content-Type", "text/event-stream")
	w.Header().Add("Cache-Control", "no-cache")

	flusher := w.(http.Flusher)

	go func() {
		defer func() {
			recover()
		}()

		for c := range 1000 {
			if c%5 == 0 {
				w.Write([]byte(FormatKeepAlive()))
			} else {
				w.Write([]byte(Format(UserData{Name: "Mizuki Kanzaki", Count: c})))
			}

			flusher.Flush()

			time.Sleep(time.Millisecond * 300)
		}
	}()

	<-r.Context().Done()
}

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodOptions:
			w.WriteHeader(http.StatusNoContent)
			return
		case http.MethodPost:
			handler.ServeHTTP(w, r)
		default:
			w.WriteHeader(404)
			return
		}
	})
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	WriteResp(w, H{"message": "ok"})
}

type H map[string]any

func WriteResp(w http.ResponseWriter, h H) {
	b, err := json.Marshal(h)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(200)
	w.Write(b)
}
