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
	litIndex, err := loadIndexedList("lists/literature.json", litColumns)
	if err != nil {
		fmt.Printf("Error loading the Literature list %v", err)
	}

	mvColumns := []string{"firstName", "lastName", "aliases", "dob"}
	mvIndex, err := loadIndexedList("lists/marvel.json", mvColumns)
	if err != nil {
		fmt.Printf("Error loading the Marvel list %v", err)
	}

	ghColumns := []string{"firstName", "lastName", "aliases"}
	ghIndex, err := loadIndexedList("lists/ghibli.json", ghColumns)
	if err != nil {
		fmt.Printf("Error loading the Ghibli list %v", err)
	}

	rcColumns := []string{"firstName", "lastName", "aliases", "dob"}
	rcIndex, err := loadIndexedList("lists/raccoon_city.json", rcColumns)
	if err != nil {
		fmt.Printf("Error loading the Raccoon City list %v", err)
	}

	controller := &ScreeningController{
		LitIndex: litIndex,
		MvIndex:  mvIndex,
		GhIndex:  ghIndex,
		RcIndex:  rcIndex,
	}

	http.HandleFunc("/wlm/screen", controller.HandleScreening)
	fmt.Println("Server running on http://localhost:9090")
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Error starting the server")
	}
}
