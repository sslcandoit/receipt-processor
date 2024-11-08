package main

import (
    "log"
    "net/http"
)

func main() {
    // Register routes
    http.HandleFunc("/receipts/process", ProcessReceiptHandler)
    http.HandleFunc("/receipts/", GetPointsHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the Receipt Processor API! Use /receipts/process and /receipts/{id}/points endpoints."))
	})

    // Start the server
    log.Println("Starting server on :8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// curl -X POST http://localhost:8080/receipts/process \
// -H "Content-Type: application/json" \
//  -d '{
//   "retailer": "Target",
//   "purchaseDate": "2022-01-01",
//   "purchaseTime": "13:01",
//   "items": [
//     {
//       "shortDescription": "Mountain Dew 12PK",
//       "price": "6.49"
//     },{
//       "shortDescription": "Emils Cheese Pizza",
//       "price": "12.25"
//     },{
//       "shortDescription": "Knorr Creamy Chicken",
//       "price": "1.26"
//     },{
//       "shortDescription": "Doritos Nacho Cheese",
//       "price": "3.35"
//     },{
//       "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
//       "price": "12.00"
//     }
//   ],
//   "total": "35.35"
// }'

// curl -X POST http://localhost:8080/receipts/process \
// -H "Content-Type: application/json" \
//  -d '{
// 	"retailer": "M&M Corner Market",
// 	"purchaseDate": "2022-03-20",
// 	"purchaseTime": "14:33",
// 	"items": [
// 	  {
// 		"shortDescription": "Gatorade",
// 		"price": "2.25"
// 	  },{
// 		"shortDescription": "Gatorade",
// 		"price": "2.25"
// 	  },{
// 		"shortDescription": "Gatorade",
// 		"price": "2.25"
// 	  },{
// 		"shortDescription": "Gatorade",
// 		"price": "2.25"
// 	  }
// 	],
// 	"total": "9.00"
//   }'
// curl http://localhost:8080/receipts/aa7c0efa-a06b-4451-99f6-cbd1f01dd743/points





