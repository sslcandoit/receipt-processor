package main

import (
    "encoding/json"
    "net/http"
    "sync"
	"strings"
	"log"
)

var (
    receiptStore = make(map[string]Receipt)//// Stores the complete Receipt objects by receipt ID
    pointsStore  = make(map[string]int)//// Stores the points associated with each receipt ID
    mu           sync.Mutex
)

//POST /receipts/process: Accepts a receipt JSON payload, generates a unique ID, 
//calculates points based on the receipt, stores it in memory, and returns the ID.
func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing a new receipt...")
    var receipt Receipt
    if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Generate a unique ID for the receipt
    id := generateID()

    // Calculate points based on receipt data
    points := calculatePoints(receipt)

    // Store receipt and points in memory
    mu.Lock()
    receiptStore[id] = receipt
    pointsStore[id] = points
    mu.Unlock()

    // Respond with the ID
    response := ReceiptResponse{ID: id}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

//GET /receipts/{id}/points: Looks up the stored points for a specific receipt ID and returns the points in JSON format.
func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
    // Extract the ID from the URL path
    id := strings.TrimPrefix(r.URL.Path, "/receipts/")
    id = strings.TrimSuffix(id, "/points")

	log.Println("Fetching points for receipt ID:", id)

    // Retrieve the points for the given ID
    mu.Lock()
    points, exists := pointsStore[id]
    mu.Unlock()

    if !exists {
        http.Error(w, "Receipt ID not found", http.StatusNotFound)
        return
    }

    // Respond with the points in JSON format
    response := map[string]int{"points": points}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
