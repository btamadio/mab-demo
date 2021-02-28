package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var campaignRewards map[int][]struct{ Alpha, Beta float64 }

func init() {
	campaignRewards = make(map[int][]struct{ Alpha, Beta float64 })

	campaignRewards[1] = []struct{ Alpha, Beta float64 }{
		{10, 125},
		{4, 130},
		{16, 80},
		{25, 99},
	}

	campaignRewards[2] = []struct{ Alpha, Beta float64 }{
		{25, 125},
		{5, 50},
		{7, 90},
		{13, 200},
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	var req rewardRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rewards, ok := campaignRewards[req.CampaignID]
	if !ok {
		http.Error(w, fmt.Sprintf("no rewards for campaign ID %d", req.CampaignID), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(rewards)
}

type rewardRequest struct {
	CampaignID int `json:"campaign_id"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/rewards", handler).Methods("POST")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"alive": true}`)
	})

	server := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:80",
	}

	log.Fatal(server.ListenAndServe())
}
