package main

import (
    "github.com/google/uuid" // Make sure to import this package for UUID generation
)

func generateID() string {
    return uuid.New().String()
}

