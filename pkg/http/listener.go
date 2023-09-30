package http

import (
	"encoding/json"
	"log"
	"net/http"

	"go01/pkg/models"
)

type Config struct {
	Addr  string
	Route string
}

func ListenForHotdogs(cfg Config, receiver func(*models.Hotdog) bool) {
	http.HandleFunc(cfg.Route, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var hotdog models.Hotdog
		err := json.NewDecoder(r.Body).Decode(&hotdog)
		if err != nil {
			log.Printf("error parsing a hotdog: %s\n", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ok := receiver(&hotdog)
		if ok {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	log.Printf("starting HTTP hotdog listener on %s%s (POST)", cfg.Addr, cfg.Route)
	log.Panic(http.ListenAndServe(cfg.Addr, nil))
}
