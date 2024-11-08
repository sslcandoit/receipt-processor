package main

import (
    "strconv"
    "strings"
	"math"
	"time"
)

func calculatePoints(receipt Receipt) int {
    points := 0

    // 1. Points for alphanumeric characters in the retailer name
    for _, c := range receipt.Retailer {
        if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
            points++
        }
    }

    // 2. 50 points if the total is a round dollar amount with no cents
    if total, err := strconv.ParseFloat(receipt.Total, 64); err == nil {
        if total == math.Floor(total) {
            points += 50
        }
    }

    // 3. 25 points if the total is a multiple of 0.25
    if total, err := strconv.ParseFloat(receipt.Total, 64); err == nil {
        if math.Mod(total, 0.25) == 0 {
            points += 25
        }
    }

	// 4. 5 points for every two items on the receipt
    points += (len(receipt.Items) / 2) * 5

    // 5. Points based on item descriptions
    for _, item := range receipt.Items {
        trimmedDesc := strings.TrimSpace(item.ShortDescription)
        if len(trimmedDesc)%3 == 0 {
            if price, err := strconv.ParseFloat(item.Price, 64); err == nil {
                itemPoints := int(math.Ceil(price * 0.2))
                points += itemPoints
            }
        }
    }

	// 6. 6 points if the day in the purchase date is odd
    if date, err := time.Parse("2006-01-02", receipt.PurchaseDate); err == nil {
        if date.Day()%2 != 0 {
            points += 6
        }
    }

    // 7. 10 points if the time of purchase is after 2:00pm and before 4:00pm
    if t, err := time.Parse("15:04", receipt.PurchaseTime); err == nil {
        if t.Hour() == 14 { // 2:00pm is 14:00 in 24-hour format
            points += 10
        }
    }
    return points
}
