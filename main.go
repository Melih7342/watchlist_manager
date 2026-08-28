package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ScreeningController struct {
	LitIndex IndexedWatchlist
	MvIndex  IndexedWatchlist
	GhIndex  IndexedWatchlist
	RcIndex  IndexedWatchlist
}

func (c *ScreeningController) HandleScreening(w http.ResponseWriter, r *http.Request) {
	// Only Post allowed
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}
	// Convert http request into struct
	var req ScreeningRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	result := RunScreening(req, c.LitIndex, c.MvIndex, c.GhIndex, c.RcIndex)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, "Error creating the JSON body", http.StatusInternalServerError)
		return
	}
}

func main() {
	litColumns := []string{"firstName", "lastName", "aliases"}
	litIndex, _ := loadIndexedList("lists/literature.json", litColumns)

	mvColumns := []string{"firstName", "lastName", "aliases", "DOB"}
	mvIndex, _ := loadIndexedList("lists/marvel.json", mvColumns)

	ghColumns := []string{"firstName", "lastName", "aliases"}
	ghIndex, _ := loadIndexedList("lists/ghibli.json", ghColumns)

	rcColumns := []string{"firstName", "lastName", "aliases", "DOB"}
	rcIndex, _ := loadIndexedList("lists/raccoon_city.json", rcColumns)

	controller := &ScreeningController{
		LitIndex: litIndex,
		MvIndex:  mvIndex,
		GhIndex:  ghIndex,
		RcIndex:  rcIndex,
	}

	http.HandleFunc("wlm/screen", controller.HandleScreening)
	fmt.Println("Server running on http://localhost:9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Error starting the server")
	}
}
